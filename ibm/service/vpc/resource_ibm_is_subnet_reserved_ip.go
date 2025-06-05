// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isReservedIPProvisioning     = "provisioning"
	isReservedIPProvisioningDone = "done"
	isReservedIP                 = "reserved_ip"
	isReservedIPTarget           = "target"
	isReservedIPTargetCrn        = "target_crn"
	isReservedIPLifecycleState   = "lifecycle_state"
)

func ResourceIBMISReservedIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISReservedIPCreate,
		ReadContext:   resourceIBMISReservedIPRead,
		UpdateContext: resourceIBMISReservedIPUpdate,
		DeleteContext: resourceIBMISReservedIPDelete,
		Exists:        resourceIBMISReservedIPExists,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			/*
				Request Parameters
				==================
				These are mandatory req parameters
			*/
			isSubNetID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet identifier.",
			},
			isReservedIPAutoDelete: {
				Type:        schema.TypeBool,
				Default:     nil,
				Computed:    true,
				Optional:    true,
				Description: "If set to true, this reserved IP will be automatically deleted",
			},
			isReservedIPName: {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_subnet_reserved_ip", isReservedIPName),
				Description:  "The user-defined or system-provided name for this reserved IP.",
			},
			isReservedIPTarget: {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "The unique identifier for target.",
			},
			isReservedIPTargetCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "The crn for target.",
			},
			isReservedIPLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the reserved IP",
			},
			/*
				Response Parameters
				===================
				All of these are computed and an user doesn't need to provide
				these from outside.
			*/

			isReservedIPAddress: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The address for this reserved IP.",
			},
			isReservedIP: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the reserved IP.",
			},
			isReservedIPCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the reserved IP was created.",
			},
			isReservedIPhref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this reserved IP.",
			},
			isReservedIPOwner: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The owner of a reserved IP, defining whether it is managed by the user or the provider.",
			},
			isReservedIPType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
		},
	}
}
func ResourceIBMISSubnetReservedIPValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isReservedIPName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	ibmISSubnetReservedIPCResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_subnet_reserved_ip", Schema: validateSchema}
	return &ibmISSubnetReservedIPCResourceValidator
}

// resourceIBMISReservedIPCreate Creates a reserved IP given a subnet ID
func resourceIBMISReservedIPCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()

	}

	subnetID := d.Get(isSubNetID).(string)
	options := sess.NewCreateSubnetReservedIPOptions(subnetID)

	nameStr := ""
	if name, ok := d.GetOk(isReservedIPName); ok {
		nameStr = name.(string)
	}
	if nameStr != "" {
		options.Name = &nameStr
	}
	addStr := ""
	if address, ok := d.GetOk(isReservedIPAddress); ok {
		addStr = address.(string)
	}
	if addStr != "" {
		options.Address = &addStr
	}

	autoDeleteBool := d.Get(isReservedIPAutoDelete).(bool)
	options.AutoDelete = &autoDeleteBool
	if t, ok := d.GetOk(isReservedIPTarget); ok {
		targetId := t.(string)
		options.Target = &vpcv1.ReservedIPTargetPrototype{
			ID: &targetId,
		}
	}
	rip, response, err := sess.CreateSubnetReservedIPWithContext(context, options)
	if err != nil || response == nil || rip == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Set id for the reserved IP as combination of subnet ID and reserved IP ID
	d.SetId(fmt.Sprintf("%s/%s", subnetID, *rip.ID))
	_, err = isWaitForReservedIpAvailable(sess, subnetID, *rip.ID, d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForReservedIpAvailable failed: %s", err.Error()), "ibm_is_subnet_reserved_ip", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMISReservedIPRead(context, d, meta)
}

func resourceIBMISReservedIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	reservedIP, diagErr := getReservedIpWithContext(context, d, meta)
	if diagErr != nil {
		return diagErr
	}

	allIDs, err := flex.IdParts(d.Id())
	if err != nil {
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "sep-id-parts").GetDiag()
		}
	}
	subnetID := allIDs[0]

	if reservedIP != nil {
		if !core.IsNil(reservedIP.Address) {
			if err = d.Set("address", reservedIP.Address); err != nil {
				err = fmt.Errorf("Error setting address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-address").GetDiag()
			}
		}
		if err = d.Set(isReservedIP, *reservedIP.ID); err != nil {
			err = fmt.Errorf("Error setting reserved_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-reserved_ip").GetDiag()
		}
		if err = d.Set(isSubNetID, subnetID); err != nil {
			err = fmt.Errorf("Error setting subnet: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-subnet").GetDiag()
		}
		if err = d.Set("lifecycle_state", reservedIP.LifecycleState); err != nil {
			err = fmt.Errorf("Error setting lifecycle_state: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-lifecycle_state").GetDiag()
		}
		if !core.IsNil(reservedIP.AutoDelete) {
			if err = d.Set("auto_delete", reservedIP.AutoDelete); err != nil {
				err = fmt.Errorf("Error setting auto_delete: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-auto_delete").GetDiag()
			}
		}
		if err = d.Set("created_at", flex.DateTimeToString(reservedIP.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-created_at").GetDiag()
		}
		if err = d.Set("href", reservedIP.Href); err != nil {
			err = fmt.Errorf("Error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-href").GetDiag()
		}
		if !core.IsNil(reservedIP.Name) {
			if err = d.Set("name", reservedIP.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-name").GetDiag()
			}
		}
		if err = d.Set("owner", reservedIP.Owner); err != nil {
			err = fmt.Errorf("Error setting owner: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-owner").GetDiag()
		}
		if err = d.Set("resource_type", reservedIP.ResourceType); err != nil {
			err = fmt.Errorf("Error setting resource_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-resource_type").GetDiag()
		}
		if reservedIP.Target != nil {
			targetIntf := reservedIP.Target
			switch reflect.TypeOf(targetIntf).String() {
			case "*vpcv1.ReservedIPTargetEndpointGatewayReference":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetEndpointGatewayReference)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target").GetDiag()
					}
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target_crn").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTargetGenericResourceReference":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetGenericResourceReference)
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target_crn").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTargetNetworkInterfaceReferenceTargetContext":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetNetworkInterfaceReferenceTargetContext)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTargetLoadBalancerReference":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetLoadBalancerReference)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target").GetDiag()
					}
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target_crn").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTargetVPNGatewayReference":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetVPNGatewayReference)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target").GetDiag()
					}
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target_crn").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTarget":
				{
					target := targetIntf.(*vpcv1.ReservedIPTarget)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target").GetDiag()
					}
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "set-target_crn").GetDiag()
					}
				}
			}
		}
	}
	return nil
}

func resourceIBMISReservedIPUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// For updating the name
	nameChanged := d.HasChange(isReservedIPName)
	autoDeleteChanged := d.HasChange(isReservedIPAutoDelete)

	if nameChanged || autoDeleteChanged {
		sess, err := vpcClient(meta)
		if err != nil {
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "update", "initialize-client")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		allIDs, err := flex.IdParts(d.Id())
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "update", "sep-id-parts").GetDiag()
		}
		subnetID := allIDs[0]
		reservedIPID := allIDs[1]

		options := &vpcv1.UpdateSubnetReservedIPOptions{
			SubnetID: &subnetID,
			ID:       &reservedIPID,
		}

		patch := new(vpcv1.ReservedIPPatch)

		if nameChanged {
			name := d.Get(isReservedIPName).(string)
			patch.Name = core.StringPtr(name)
		}

		if autoDeleteChanged {
			autoDelete := d.Get(isReservedIPAutoDelete).(bool)
			patch.AutoDelete = core.BoolPtr(autoDelete)
		}

		reservedIPPatch, err := patch.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("patch.AsPatch() failed: %s", err.Error()), "ibm_is_subnet_reserved_ip", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		options.ReservedIPPatch = reservedIPPatch

		_, _, err = sess.UpdateSubnetReservedIPWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForReservedIpAvailable(sess, subnetID, reservedIPID, d.Timeout(schema.TimeoutCreate), d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForReservedIpAvailable failed: %s", err.Error()), "ibm_is_subnet_reserved_ip", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISReservedIPRead(context, d, meta)
}

func resourceIBMISReservedIPDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	rip, diagErr := getReservedIpWithContext(context, d, meta)
	if diagErr != nil {
		return diagErr
	}
	if diagErr == nil && rip == nil {
		// If there is no such reserved IP, it can not be deleted
		return nil
	}

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allIDs, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "delete", "sep-id-parts").GetDiag()
	}
	subnetID := allIDs[0]
	reservedIPID := allIDs[1]
	deleteOptions := sess.NewDeleteSubnetReservedIPOptions(subnetID, reservedIPID)
	response, err := sess.DeleteSubnetReservedIPWithContext(context, deleteOptions)
	if err != nil || response == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISReservedIPExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rip, err := getReservedIpWithoutContext(d, meta)
	if err != nil {
		return false, err
	}
	if err == nil && rip == nil {
		return false, nil
	}
	return true, nil
}

// get is a generic function that gets the reserved ip given subnet id and reserved ip
func getReservedIpWithContext(context context.Context, d *schema.ResourceData, meta interface{}) (*vpcv1.ReservedIP, diag.Diagnostics) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return nil, tfErr.GetDiag()
	}
	allIDs, err := flex.IdParts(d.Id())
	if err != nil {
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "sep-id-parts").GetDiag()
	}
	subnetID := allIDs[0]
	reservedIPID := allIDs[1]
	options := sess.NewGetSubnetReservedIPOptions(subnetID, reservedIPID)
	reservedIP, response, err := sess.GetSubnetReservedIPWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return nil, tfErr.GetDiag()
	}
	return reservedIP, nil
}
func getReservedIpWithoutContext(d *schema.ResourceData, meta interface{}) (*vpcv1.ReservedIP, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return nil, tfErr
	}
	allIDs, err := flex.IdParts(d.Id())
	if err != nil {
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip", "read", "sep-id-parts")
	}
	subnetID := allIDs[0]
	reservedIPID := allIDs[1]
	options := sess.NewGetSubnetReservedIPOptions(subnetID, reservedIPID)
	reservedIP, response, err := sess.GetSubnetReservedIP(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return nil, tfErr
	}
	return reservedIP, nil
}

func isWaitForReservedIpAvailable(sess *vpcv1.VpcV1, subnetid, id string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for reseved ip (%s/%s) to be available.", subnetid, id)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"pending"},
		Target:     []string{"done", "failed", ""},
		Refresh:    isReserveIpRefreshFunc(sess, subnetid, id, d),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isReserveIpRefreshFunc(sess *vpcv1.VpcV1, subnetid, id string, d *schema.ResourceData) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getreservedipOptions := &vpcv1.GetSubnetReservedIPOptions{
			ID:       &id,
			SubnetID: &subnetid,
		}
		rsip, response, err := sess.GetSubnetReservedIP(getreservedipOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting reserved ip(%s/%s) : %s\n%s", subnetid, id, err, response)
		}
		if rsip.LifecycleState != nil {
			d.Set(isReservedIPLifecycleState, *rsip.LifecycleState)
		}
		d.Set(isReservedIPAddress, *rsip.Address)

		if rsip.LifecycleState != nil && *rsip.LifecycleState == "failed" {
			return rsip, "failed", fmt.Errorf("[ERROR] Error Reserved ip(%s/%s) creation failed : %s\n%s", subnetid, id, err, response)
		}
		if rsip.LifecycleState != nil && *rsip.LifecycleState == "stable" {
			return rsip, "done", nil
		}
		return rsip, "pending", nil
	}
}
