// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
			Update: schema.DefaultTimeout(10 * time.Minute),
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

			"deleted": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"more_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about deleted resources.",
						},
					},
				},
			},

			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for Security group target.",
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
		return diag.FromErr(err)
	}

	securityGroupID := d.Get("security_group").(string)
	targetID := d.Get(isSecurityGroupTargetID).(string)

	createSecurityGroupTargetBindingOptions := &vpcv1.CreateSecurityGroupTargetBindingOptions{}
	createSecurityGroupTargetBindingOptions.SecurityGroupID = &securityGroupID
	createSecurityGroupTargetBindingOptions.ID = &targetID

	sg, response, err := sess.CreateSecurityGroupTargetBinding(createSecurityGroupTargetBindingOptions)
	if err != nil || sg == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while creating Security Group Target Binding %s\n%s", err, response))
	}
	sgtarget := sg.(*vpcv1.SecurityGroupTargetReference)
	d.SetId(fmt.Sprintf("%s/%s", securityGroupID, *sgtarget.ID))

	return resourceIBMISSecurityGroupTargetRead(context, d, meta)
}

func resourceIBMISSecurityGroupTargetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	securityGroupID := parts[0]
	securityGroupTargetID := parts[1]

	d.Set("security_group", securityGroupID)
	d.Set(isSecurityGroupTargetID, securityGroupTargetID)

	getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
		SecurityGroupID: &securityGroupID,
		ID:              &securityGroupTargetID,
	}

	data, response, err := sess.GetSecurityGroupTarget(getSecurityGroupTargetOptions)
	if err != nil || data == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting Security Group Target : %s\n%s", err, response))
	}

	target := data.(*vpcv1.SecurityGroupTargetReference)

	if target.ResourceType != nil && *target.ResourceType != "" && *target.ResourceType == "load_balancer" {
		lbid := target.ID
		_, errsgt := isWaitForLbSgTargetCreateAvailable(sess, *lbid, d.Timeout(schema.TimeoutUpdate))
		if errsgt != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while creating Security Group Target Binding %s\n", errsgt))
		}
	} else if target.ResourceType != nil && *target.ResourceType != "" && *target.ResourceType == "endpoint_gateway" {
		edid := target.ID
		_, errsgt := isWaitForVirtualEndpointGatewayAvailable(sess, *edid, d.Timeout(schema.TimeoutUpdate))
		if errsgt != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while creating Security Group Target Binding %s\n", errsgt))
		}
	} else if target.ResourceType != nil && *target.ResourceType != "" && *target.ResourceType == "vpn_server" {
		_, errsgt := isWaitForVPNServerStable(context, sess, d, d.Timeout(schema.TimeoutUpdate))
		if errsgt != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while creating Security Group Target Binding %s\n", errsgt))
		}
	}

	if target.Name != nil {
		if err = d.Set("name", *target.Name); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting name: %s", err))
		}
	}
	if target.Href != nil {
		if err = d.Set("href", *target.Href); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting href: %s", err))
		}
	}
	if target.CRN != nil {
		if err = d.Set("crn", *target.CRN); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting crn: %s", err))
		}
	}

	if target.Deleted != nil {
		targetDeletedMap := map[string]interface{}{}
		targetDeletedMap["more_info"] = target.Deleted.MoreInfo
		if err = d.Set("deleted", []map[string]interface{}{targetDeletedMap}); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting crn: %s", err))
		}
	}

	if target.ResourceType != nil && *target.ResourceType != "" {
		if err = d.Set(isSecurityGroupResourceType, *target.ResourceType); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting resource type: %s", err))
		}
	}

	return nil
}

func resourceIBMISSecurityGroupTargetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	securityGroupID := parts[0]
	securityGroupTargetID := parts[1]

	getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
		SecurityGroupID: &securityGroupID,
		ID:              &securityGroupTargetID,
	}
	sgt, response, err := sess.GetSecurityGroupTarget(getSecurityGroupTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error Getting Security Group Targets (%s): %s\n%s", securityGroupID, err, response))
	}

	deleteSecurityGroupTargetBindingOptions := sess.NewDeleteSecurityGroupTargetBindingOptions(securityGroupID, securityGroupTargetID)
	response, err = sess.DeleteSecurityGroupTargetBinding(deleteSecurityGroupTargetBindingOptions)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error Deleting Security Group Targets : %s\n%s", err, response))
	}
	securityGroupTargetReference := sgt.(*vpcv1.SecurityGroupTargetReference)
	if securityGroupTargetReference.ResourceType != nil && *securityGroupTargetReference.ResourceType != "" && *securityGroupTargetReference.ResourceType == "load_balancer" {
		lbid := securityGroupTargetReference.ID
		_, errsgt := isWaitForLBRemoveAvailable(sess, sgt, *lbid, securityGroupID, securityGroupTargetID, d.Timeout(schema.TimeoutDelete))
		if errsgt != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while deleting Security Group Target Binding %s\n", errsgt))
		}
	} else if securityGroupTargetReference.ResourceType != nil && *securityGroupTargetReference.ResourceType != "" && *securityGroupTargetReference.ResourceType == "endpoint_gateway" {
		edid := securityGroupTargetReference.ID
		_, errsgt := isWaitForVPNServerRemoveAvailable(sess, sgt, *edid, securityGroupID, securityGroupTargetID, d.Timeout(schema.TimeoutDelete))
		if errsgt != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while creating Security Group Target Binding %s\n", errsgt))
		}
	} else if securityGroupTargetReference.ResourceType != nil && *securityGroupTargetReference.ResourceType != "" && *securityGroupTargetReference.ResourceType == "vpn_server" {
		vpnServerId := securityGroupTargetReference.ID
		_, errsgt := isWaitForVPNServerRemoveAvailable(sess, sgt, *vpnServerId, securityGroupID, securityGroupTargetID, d.Timeout(schema.TimeoutDelete))
		if errsgt != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error while creating Security Group Target Binding %s\n", errsgt))
		}
	}
	d.SetId("")
	return nil
}

func resourceIBMISSecurityGroupTargetExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, err
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
		return false, fmt.Errorf("[ERROR] Error getting Security Group Target : %s\n%s", err, response)
	}
	return true, nil

}

func isWaitForLBRemoveAvailable(sess *vpcv1.VpcV1, sgt vpcv1.SecurityGroupTargetReferenceIntf, lbId, securityGroupID, securityGroupTargetID string, timeout time.Duration) (interface{}, error) {
	log.Printf("[INFO] Waiting for load balancer binding (%s) to be removed.", lbId)

	stateConf := &resource.StateChangeConf{
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

func isWaitForVPNServerRemoveAvailable(sess *vpcv1.VpcV1, sgt vpcv1.SecurityGroupTargetReferenceIntf, vpnServerId, securityGroupID, securityGroupTargetID string, timeout time.Duration) (interface{}, error) {
	log.Printf("[INFO] Waiting for vpn server binding (%s) to be removed.", vpnServerId)

	stateConf := &resource.StateChangeConf{
		Pending: []string{isVPNServerStatusPending},
		Target:  []string{isVPNServerStatusStable},
		Refresh: func() (interface{}, string, error) {

			getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
				SecurityGroupID: &securityGroupID,
				ID:              &securityGroupTargetID,
			}
			_, response, err := sess.GetSecurityGroupTarget(getSecurityGroupTargetOptions)

			if err != nil {
				if response != nil && response.StatusCode == 404 {
					getVPNServerOptions := &vpcv1.GetVPNServerOptions{}

					getVPNServerOptions.SetID(vpnServerId)

					vpnServer, response, err := sess.GetVPNServer(getVPNServerOptions)
					if err != nil {
						log.Printf("[DEBUG] GetVPNServerWithContext failed %s\n%s", err, response)
						return vpnServer, "", fmt.Errorf("Error Getting VPC Server: %s\n%s", err, response)
					}

					if *vpnServer.LifecycleState == "stable" || *vpnServer.LifecycleState == "failed" {
						return vpnServer, *vpnServer.LifecycleState, nil
					}
					return vpnServer, *vpnServer.LifecycleState, nil
				}
				return nil, isVPNServerStatusStable, fmt.Errorf("[ERROR] Error getting Security Group Target : %s\n%s", err, response)
			}
			return sgt, isVPNServerStatusPending, nil
		},
		Timeout:        timeout,
		Delay:          10 * time.Second,
		MinTimeout:     10 * time.Second,
		NotFoundChecks: 1,
	}

	return stateConf.WaitForState()
}

func isWaitForEndpointGatewayRemoveAvailable(sess *vpcv1.VpcV1, sgt vpcv1.SecurityGroupTargetReferenceIntf, endpointGatewayId, securityGroupID, securityGroupTargetID string, timeout time.Duration) (interface{}, error) {
	log.Printf("[INFO] Waiting for endpoint gateway binding (%s) to be removed.", endpointGatewayId)

	stateConf := &resource.StateChangeConf{
		Pending: []string{"waiting", "pending", "updating"},
		Target:  []string{"stable"},
		Refresh: func() (interface{}, string, error) {

			getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
				SecurityGroupID: &securityGroupID,
				ID:              &securityGroupTargetID,
			}
			_, response, err := sess.GetSecurityGroupTarget(getSecurityGroupTargetOptions)

			if err != nil {
				if response != nil && response.StatusCode == 404 {

					opt := sess.NewGetEndpointGatewayOptions(endpointGatewayId)
					result, response, err := sess.GetEndpointGateway(opt)
					if err != nil {
						if response != nil && response.StatusCode == 404 {
							return result, "", fmt.Errorf("Error Getting Virtual Endpoint Gateway : %s\n%s", err, response)
						}
					}
					if *result.LifecycleState == "stable" || *result.LifecycleState == "failed" {
						return result, *result.LifecycleState, nil
					}
					return result, *result.LifecycleState, nil
				}
				return nil, "stable", fmt.Errorf("[ERROR] Error getting Security Group Target : %s\n%s", err, response)
			}
			return sgt, "pending", nil
		},
		Timeout:        timeout,
		Delay:          10 * time.Second,
		MinTimeout:     10 * time.Second,
		NotFoundChecks: 1,
	}

	return stateConf.WaitForState()
}

func isLBRemoveRefreshFunc(sess *vpcv1.VpcV1, sgt vpcv1.SecurityGroupTargetReferenceIntf, lbId, securityGroupID, securityGroupTargetID string) resource.StateRefreshFunc {
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

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBProvisioning, "update_pending"},
		Target:     []string{isLBProvisioningDone, ""},
		Refresh:    isLBSgTargetRefreshFunc(sess, lbId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBSgTargetRefreshFunc(sess *vpcv1.VpcV1, lbId string) resource.StateRefreshFunc {
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
