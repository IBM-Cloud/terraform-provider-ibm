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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMISReservedIPPatch() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISReservedIPPatchCreate,
		ReadContext:   resourceIBMISReservedIPPatchRead,
		UpdateContext: resourceIBMISReservedIPPatchUpdate,
		DeleteContext: resourceIBMISReservedIPPatchDelete,
		Exists:        resourceIBMISReservedIPPatchExists,
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
				Type:         schema.TypeBool,
				Default:      nil,
				AtLeastOneOf: []string{isReservedIPAutoDelete, isReservedIPName},
				Computed:     true,
				Optional:     true,
				Description:  "If set to true, this reserved IP will be automatically deleted",
			},
			isReservedIPName: {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{isReservedIPAutoDelete, isReservedIPName},
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_subnet_reserved_ip_patch", isReservedIPName),
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
			isReservedIPAddress: {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The address for this reserved IP.",
			},
			isReservedIP: {
				Type:        schema.TypeString,
				Required:    true,
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

// resourceIBMISReservedIPCreate Creates a reserved IP given a subnet ID
func resourceIBMISReservedIPPatchCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	subnetID := d.Get(isSubNetID).(string)
	reservedIPID := d.Get(isReservedIP).(string)
	name := d.Get(isReservedIPName).(string)
	reservedIPPatchModel := &vpcv1.ReservedIPPatch{}
	if name != "" {
		reservedIPPatchModel.Name = &name
	}
	if autoDeleteBoolOk, ok := d.GetOkExists(isReservedIPAutoDelete); ok {
		autoDeleteBool := autoDeleteBoolOk.(bool)
		reservedIPPatchModel.AutoDelete = &autoDeleteBool
	}
	reservedIPPatch, err := reservedIPPatchModel.AsPatch()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("reservedIPPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_subnet_reserved_ip_patch", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := sess.NewUpdateSubnetReservedIPOptions(subnetID, reservedIPID, reservedIPPatch)

	rip, response, err := sess.UpdateSubnetReservedIPWithContext(context, options)
	if err != nil || response == nil || rip == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip_patch", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Set id for the reserved IP as combination of subnet ID and reserved IP ID
	d.SetId(fmt.Sprintf("%s/%s", subnetID, *rip.ID))
	return resourceIBMISReservedIPPatchRead(context, d, meta)
}

func resourceIBMISReservedIPPatchRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	reservedIP, diagErr := getReservedIpPatchWithContext(context, d, meta)
	if diagErr != nil {
		return diagErr
	}

	allIDs, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "sep-id-parts").GetDiag()
	}
	subnetID := allIDs[0]

	if reservedIP != nil {
		if !core.IsNil(reservedIP.Address) {
			if err = d.Set("address", reservedIP.Address); err != nil {
				err = fmt.Errorf("Error setting address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-address").GetDiag()
			}
		}
		if err = d.Set(isReservedIP, *reservedIP.ID); err != nil {
			err = fmt.Errorf("Error setting reserved_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-reserved_ip").GetDiag()
		}
		if err = d.Set(isSubNetID, subnetID); err != nil {
			err = fmt.Errorf("Error setting subnet: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-subnet").GetDiag()
		}
		if err = d.Set("lifecycle_state", reservedIP.LifecycleState); err != nil {
			err = fmt.Errorf("Error setting lifecycle_state: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-lifecycle_state").GetDiag()
		}
		if !core.IsNil(reservedIP.AutoDelete) {
			if err = d.Set("auto_delete", reservedIP.AutoDelete); err != nil {
				err = fmt.Errorf("Error setting auto_delete: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-auto_delete").GetDiag()
			}
		}
		if err = d.Set("created_at", flex.DateTimeToString(reservedIP.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-created_at").GetDiag()
		}
		if err = d.Set("href", reservedIP.Href); err != nil {
			err = fmt.Errorf("Error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-href").GetDiag()
		}
		if !core.IsNil(reservedIP.Name) {
			if err = d.Set("name", reservedIP.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-name").GetDiag()
			}
		}
		if err = d.Set("owner", reservedIP.Owner); err != nil {
			err = fmt.Errorf("Error setting owner: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-owner").GetDiag()
		}
		if err = d.Set("resource_type", reservedIP.ResourceType); err != nil {
			err = fmt.Errorf("Error setting resource_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-resource_type").GetDiag()
		}
		if reservedIP.Target != nil {
			targetIntf := reservedIP.Target
			switch reflect.TypeOf(targetIntf).String() {
			case "*vpcv1.ReservedIPTargetEndpointGatewayReference":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetEndpointGatewayReference)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target").GetDiag()
					}
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target_crn").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTargetGenericResourceReference":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetGenericResourceReference)
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target_crn").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTargetNetworkInterfaceReferenceTargetContext":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetNetworkInterfaceReferenceTargetContext)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTargetLoadBalancerReference":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetLoadBalancerReference)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target").GetDiag()
					}
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target_crn").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTargetVPNGatewayReference":
				{
					target := targetIntf.(*vpcv1.ReservedIPTargetVPNGatewayReference)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target").GetDiag()
					}
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target_crn").GetDiag()
					}
				}
			case "*vpcv1.ReservedIPTarget":
				{
					target := targetIntf.(*vpcv1.ReservedIPTarget)
					if err = d.Set(isReservedIPTarget, target.ID); err != nil {
						err = fmt.Errorf("Error setting target: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target").GetDiag()
					}
					if err = d.Set(isReservedIPTargetCrn, target.CRN); err != nil {
						err = fmt.Errorf("Error setting target_crn: %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "set-target_crn").GetDiag()
					}
				}
			}
		}
	}
	return nil
}

func resourceIBMISReservedIPPatchUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// For updating the name
	nameChanged := d.HasChange(isReservedIPName)
	autoDeleteChanged := d.HasChange(isReservedIPAutoDelete)

	if nameChanged || autoDeleteChanged {
		sess, err := vpcClient(meta)
		if err != nil {
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "update", "initialize-client")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		allIDs, err := flex.IdParts(d.Id())
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "update", "sep-id-parts").GetDiag()
		}
		subnetID := allIDs[0]
		reservedIPID := allIDs[1]

		options := &vpcv1.UpdateSubnetReservedIPOptions{
			SubnetID: &subnetID,
			ID:       &reservedIPID,
		}

		reservedIPPatchModel := new(vpcv1.ReservedIPPatch)

		if nameChanged {
			name := d.Get(isReservedIPName).(string)
			reservedIPPatchModel.Name = core.StringPtr(name)
		}

		if autoDeleteChanged {
			autoDelete := d.Get(isReservedIPAutoDelete).(bool)
			reservedIPPatchModel.AutoDelete = core.BoolPtr(autoDelete)
		}

		reservedIPPatch, err := reservedIPPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("reservedIPPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_subnet_reserved_ip_patch", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		options.ReservedIPPatch = reservedIPPatch

		_, _, err = sess.UpdateSubnetReservedIPWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip_patch", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISReservedIPPatchRead(context, d, meta)
}

func resourceIBMISReservedIPPatchDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func resourceIBMISReservedIPPatchExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rip, err := getReservedIpPatchWithoutContext(d, meta)
	if err != nil {
		return false, err
	}
	if err == nil && rip == nil {
		return false, nil
	}
	return true, nil
}

// get is a generic function that gets the reserved ip given subnet id and reserved ip
func getReservedIpPatchWithContext(context context.Context, d *schema.ResourceData, meta interface{}) (*vpcv1.ReservedIP, diag.Diagnostics) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return nil, tfErr.GetDiag()
	}
	allIDs, err := flex.IdParts(d.Id())
	if err != nil {
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "sep-id-parts").GetDiag()
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
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip_patch", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return nil, tfErr.GetDiag()
	}
	return reservedIP, nil
}

func getReservedIpPatchWithoutContext(d *schema.ResourceData, meta interface{}) (*vpcv1.ReservedIP, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return nil, tfErr
	}
	allIDs, err := flex.IdParts(d.Id())
	if err != nil {
		return nil, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_reserved_ip_patch", "read", "sep-id-parts")
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
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_reserved_ip_patch", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return nil, tfErr
	}
	return reservedIP, nil
}
