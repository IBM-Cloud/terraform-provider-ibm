// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSecurityGroupTargetID     = "target"
	isSecurityGroupResourceType = "resource_type"
)

func ResourceIBMISSecurityGroupTarget() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceIBMISSecurityGroupTargetCreate,
		ReadContext:   resourceIBMISSecurityGroupTargetRead,
		DeleteContext: resourceIBMISSecurityGroupTargetDelete,
		Exists:        resourceIBMISSecurityGroupTargetExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"security_group": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Security group id",
			},

			isSecurityGroupTargetID: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "security group target identifier",
				ValidateFunc: validate.InvokeValidator("ibm_is_security_group_target", isSecurityGroupTargetID),
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Security group target name",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this Security group target",
			},

			isSecurityGroupResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource Type",
			},
		},
	}
}

func ResourceIBMISSecurityGroupTargetValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupTargetID,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64},
		validate.ValidateSchema{
			Identifier:                 "security_group",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64})

	ibmISSecurityGroupResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_security_group_target", Schema: validateSchema}
	return &ibmISSecurityGroupResourceValidator
}

func resourceIBMISSecurityGroupTargetCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	securityGroupID := d.Get("security_group").(string)
	targetID := d.Get(isSecurityGroupTargetID).(string)

	createSecurityGroupTargetBindingOptions := &vpcv1.CreateSecurityGroupTargetBindingOptions{}
	createSecurityGroupTargetBindingOptions.SecurityGroupID = &securityGroupID
	createSecurityGroupTargetBindingOptions.ID = &targetID
	isSGTargetPrefixKey := "security_group_key_" + targetID
	conns.IbmMutexKV.Lock(isSGTargetPrefixKey)
	defer conns.IbmMutexKV.Unlock(isSGTargetPrefixKey)

	sg, _, err := sess.CreateSecurityGroupTargetBindingWithContext(context, createSecurityGroupTargetBindingOptions)
	if err != nil || sg == nil {

		if strings.Contains(strings.ToLower(err.Error()), "load balancer") && ((strings.Contains(strings.ToUpper(err.Error()), "UPDATE_PENDING")) || (strings.Contains(strings.ToUpper(err.Error()), "CREATE_PENDING"))) {
			log.Printf("[INFO] Load balancer with ID '%s' is in UPDATE_PENDING state. Waiting for it to become available before retrying...", targetID)

			// Wait for the load balancer to become available
			_, waitErr := isWaitForSGTargetLBAvailable(sess, targetID, d.Timeout(schema.TimeoutCreate))
			if waitErr != nil {
				err = fmt.Errorf("isWaitForSGTargetLBAvailable failed: waiting for load balancer to become available while creating Security Group Target Binding: %s", waitErr)
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLbSgTargetCreateAvailable failed: %s", err.Error()), "ibm_is_security_group_target", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}

			sg, _, err = sess.CreateSecurityGroupTargetBindingWithContext(context, createSecurityGroupTargetBindingOptions)
			// Check for errors after the initial attempt or retry
			if err != nil || sg == nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_security_group_target", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		} else {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_security_group_target", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	sgtarget := sg.(*vpcv1.SecurityGroupTargetReference)
	d.SetId(fmt.Sprintf("%s/%s", securityGroupID, *sgtarget.ID))
	crn := sgtarget.CRN
	if crn != nil && *crn != "" && strings.Contains(*crn, "load-balancer") {
		lbid := sgtarget.ID
		_, errsgt := isWaitForLbSgTargetCreateAvailable(sess, *lbid, d.Timeout(schema.TimeoutCreate))
		if errsgt != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLbSgTargetCreateAvailable failed: %s", errsgt.Error()), "ibm_is_security_group_target", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else if crn != nil && *crn != "" && (strings.Contains(*crn, "virtual-network-interface") || strings.Contains(*crn, "virtual-network-interfaces")) {
		vniId := sgtarget.ID
		_, errsgt := isWaitForVNISgTargetCreateAvailable(sess, *vniId, d.Timeout(schema.TimeoutCreate))
		if errsgt != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVNISgTargetCreateAvailable failed: %s", err.Error()), "ibm_is_security_group_target", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMISSecurityGroupTargetRead(context, d, meta)
}

func resourceIBMISSecurityGroupTargetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "read", "sep-id-parts").GetDiag()
	}
	securityGroupID := parts[0]
	securityGroupTargetID := parts[1]

	if err = d.Set("security_group", securityGroupID); err != nil {
		err = fmt.Errorf("Error setting security_group: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "read", "set-security_group").GetDiag()
	}

	if err = d.Set(isSecurityGroupTargetID, securityGroupTargetID); err != nil {
		err = fmt.Errorf("Error setting target: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "read", "set-target").GetDiag()
	}

	getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
		SecurityGroupID: &securityGroupID,
		ID:              &securityGroupTargetID,
	}

	data, response, err := sess.GetSecurityGroupTargetWithContext(context, getSecurityGroupTargetOptions)
	if err != nil || data == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupTarget failed: %s", err.Error()), "ibm_is_security_group_target", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	target := data.(*vpcv1.SecurityGroupTargetReference)

	if err = d.Set("name", *target.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "read", "set-name").GetDiag()
	}

	if err = d.Set("crn", target.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "read", "set-crn").GetDiag()
	}
	if target.ResourceType != nil && *target.ResourceType != "" {

		if err = d.Set(isSecurityGroupResourceType, *target.ResourceType); err != nil {
			err = fmt.Errorf("Error setting resource_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "read", "set-resource_type").GetDiag()
		}
	}

	return nil
}

func resourceIBMISSecurityGroupTargetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "delete", "sep-id-parts").GetDiag()
	}
	securityGroupID := parts[0]
	securityGroupTargetID := parts[1]

	getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
		SecurityGroupID: &securityGroupID,
		ID:              &securityGroupTargetID,
	}
	sgt, response, err := sess.GetSecurityGroupTargetWithContext(context, getSecurityGroupTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupTargetWithContext failed: %s", err.Error()), "ibm_is_security_group_target", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// Acquire a lock based on the target ID to prevent simultaneous delete on same target
	isSGTargetPrefixKey := "security_group_key_" + securityGroupTargetID
	conns.IbmMutexKV.Lock(isSGTargetPrefixKey)
	defer conns.IbmMutexKV.Unlock(isSGTargetPrefixKey)

	deleteSecurityGroupTargetBindingOptions := sess.NewDeleteSecurityGroupTargetBindingOptions(securityGroupID, securityGroupTargetID)
	response, err = sess.DeleteSecurityGroupTargetBindingWithContext(context, deleteSecurityGroupTargetBindingOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_security_group_target", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	securityGroupTargetReference := sgt.(*vpcv1.SecurityGroupTargetReference)
	crn := securityGroupTargetReference.CRN
	if crn != nil && *crn != "" && strings.Contains(*crn, "load-balancer") {
		lbid := securityGroupTargetReference.ID
		_, errsgt := isWaitForLBRemoveAvailable(sess, sgt, *lbid, securityGroupID, securityGroupTargetID, d.Timeout(schema.TimeoutDelete))
		if errsgt != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLBRemoveAvailable failed: %s", err.Error()), "ibm_is_security_group_target", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	d.SetId("")
	return nil
}

func resourceIBMISSecurityGroupTargetExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, err
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group_target", "exists", "sep-id-parts")
	}
	securityGroupID := parts[0]
	securityGroupTargetID := parts[1]

	getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
		SecurityGroupID: &securityGroupID,
		ID:              &securityGroupTargetID,
	}

	_, response, err := sess.GetSecurityGroupTarget(getSecurityGroupTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupTarget failed: %s", err.Error()), "ibm_is_security_group_target", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil

}

func isWaitForLBRemoveAvailable(sess *vpcv1.VpcV1, sgt vpcv1.SecurityGroupTargetReferenceIntf, lbId, securityGroupID, securityGroupTargetID string, timeout time.Duration) (interface{}, error) {
	log.Printf("[INFO] Waiting for load balancer binding (%s) to be removed.", lbId)

	stateConf := &retry.StateChangeConf{
		Pending:        []string{isLBProvisioning},
		Target:         []string{isLBProvisioningDone},
		Refresh:        isLBRemoveRefreshFunc(sess, sgt, lbId, securityGroupID, securityGroupTargetID),
		Timeout:        timeout,
		Delay:          10 * time.Second,
		MinTimeout:     10 * time.Second,
		NotFoundChecks: 1,
	}

	return stateConf.WaitForState()
}

func isLBRemoveRefreshFunc(sess *vpcv1.VpcV1, sgt vpcv1.SecurityGroupTargetReferenceIntf, lbId, securityGroupID, securityGroupTargetID string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
			SecurityGroupID: &securityGroupID,
			ID:              &securityGroupTargetID,
		}
		_, response, err := sess.GetSecurityGroupTarget(getSecurityGroupTargetOptions)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				getlboptions := &vpcv1.GetLoadBalancerOptions{
					ID: &lbId,
				}
				lb, response, err := sess.GetLoadBalancer(getlboptions)
				if err != nil {
					return nil, "", fmt.Errorf("[ERROR] Error Getting Load Balancer : %s\n%s", err, response)
				}

				if *lb.ProvisioningStatus == "active" || *lb.ProvisioningStatus == "failed" {
					return sgt, isLBProvisioningDone, nil
				} else {
					return sgt, isLBProvisioning, nil
				}
			}
			return nil, isLBProvisioningDone, fmt.Errorf("[ERROR] Error getting Security Group Target : %s\n%s", err, response)
		}
		return sgt, isLBProvisioning, nil
	}
}

func isWaitForLbSgTargetCreateAvailable(sess *vpcv1.VpcV1, lbId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer (%s) to be available.", lbId)

	stateConf := &retry.StateChangeConf{
		Pending:    []string{"retry", isLBProvisioning, "update_pending"},
		Target:     []string{isLBProvisioningDone, ""},
		Refresh:    isLBSgTargetRefreshFunc(sess, lbId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBSgTargetRefreshFunc(sess *vpcv1.VpcV1, lbId string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getlboptions := &vpcv1.GetLoadBalancerOptions{
			ID: &lbId,
		}
		lb, response, err := sess.GetLoadBalancer(getlboptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Load Balancer : %s\n%s", err, response)
		}

		if *lb.ProvisioningStatus == "active" || *lb.ProvisioningStatus == "failed" {
			return lb, isLBProvisioningDone, nil
		}

		return lb, isLBProvisioning, nil
	}
}

func isWaitForVNISgTargetCreateAvailable(sess *vpcv1.VpcV1, vniId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for virtual network interface (%s) to be available.", vniId)

	stateConf := &retry.StateChangeConf{
		Pending:    []string{"pending", "updating", "waiting"},
		Target:     []string{isLBProvisioningDone, "", "stable"},
		Refresh:    isVNISgTargetRefreshFunc(sess, vniId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVNISgTargetRefreshFunc(vpcClient *vpcv1.VpcV1, vniId string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getVNIOptions := &vpcv1.GetVirtualNetworkInterfaceOptions{
			ID: &vniId,
		}
		vni, response, err := vpcClient.GetVirtualNetworkInterface(getVNIOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting virtual network interface : %s\n%s", err, response)
		}

		if *vni.LifecycleState == "failed" {
			return vni, *vni.LifecycleState, fmt.Errorf("Virtual Network Interface creating failed with status %s ", *vni.LifecycleState)
		}
		return vni, *vni.LifecycleState, nil
	}
}

// New function to wait for a load balancer to become available before attaching a security group
func isWaitForSGTargetLBAvailable(sess *vpcv1.VpcV1, lbId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer (%s) to be available before attaching security group.", lbId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBProvisioning, "update_pending"},
		Target:     []string{isLBProvisioningDone, ""},
		Refresh:    isSGTargetLBRefreshFunc(sess, lbId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

// Refresh function for checking load balancer status before security group attachment
func isSGTargetLBRefreshFunc(sess *vpcv1.VpcV1, lbId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getlboptions := &vpcv1.GetLoadBalancerOptions{
			ID: &lbId,
		}
		lb, response, err := sess.GetLoadBalancer(getlboptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Load Balancer : %s\n%s", err, response)
		}

		if *lb.ProvisioningStatus == "active" || *lb.ProvisioningStatus == "failed" {
			return lb, isLBProvisioningDone, nil
		}

		return lb, isLBProvisioning, nil
	}
}
