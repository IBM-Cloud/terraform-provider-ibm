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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isInstanceActionAvailable = "available"
	isInstanceActionPending   = "pending"
	isInstanceActionFailed    = "failed"
	isInstanceStopType        = "stop_type"
	isInstanceID              = "instance"
	isInstanceActionForce     = "force_action"
)

func ResourceIBMISInstanceAction() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISInstanceActionCreate,
		ReadContext:   resourceIBMISInstanceActionRead,
		UpdateContext: resourceIBMISInstanceActionUpdate,
		DeleteContext: resourceIBMISInstanceActionDelete,
		Exists:        resourceIBMISInstanceActionExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance identifier",
			},
			isInstanceAction: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_action", isInstanceAction),
				Description:  "This restart/start/stops an instance.",
			},
			isInstanceActionForce: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If set to true, the action will be forced immediately, and all queued actions deleted. Ignored for the start action.",
			},
			isInstanceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance status",
			},

			isInstanceStatusReasons: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isInstanceStatusReasonsCode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason",
						},

						isInstanceStatusReasonsMessage: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason",
						},

						isInstanceStatusReasonsMoreInfo: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about this status reason",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMISInstanceActionValidator() *validate.ResourceValidator {

	instanceActions := "start, reboot, stop"
	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isInstanceAction,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              instanceActions})
	ibmISInstanceActionResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_action", Schema: validateSchema}
	return &ibmISInstanceActionResourceValidator
}

func resourceIBMISInstanceActionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_action", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	instanceId := ""
	if insId, ok := d.GetOk(isInstanceID); ok {
		instanceId = insId.(string)
	}

	actiontypeIntf := d.Get(isInstanceAction)
	actiontype := actiontypeIntf.(string)

	getinsOptions := &vpcv1.GetInstanceOptions{
		ID: &instanceId,
	}
	instance, response, err := sess.GetInstanceWithContext(context, getinsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_action", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if (actiontype == "stop" || actiontype == "reboot") && *instance.Status != isInstanceStatusRunning {
		d.Set(isInstanceAction, nil)
		err = fmt.Errorf("[ERROR] Error with stop/reboot action: Cannot invoke stop/reboot action while instance is not in running state")
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_action", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	} else if actiontype == "start" && *instance.Status != isInstanceActionStatusStopped {
		d.Set(isInstanceAction, nil)
		err = fmt.Errorf("[ERROR] Error with start action: Cannot invoke start action while instance is not in stopped state")
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_action", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	createinsactoptions := &vpcv1.CreateInstanceActionOptions{
		InstanceID: &instanceId,
		Type:       &actiontype,
	}
	if instanceActionForceIntf, ok := d.GetOk(isInstanceActionForce); ok {
		force := instanceActionForceIntf.(bool)
		createinsactoptions.Force = &force
	}
	_, response, err = sess.CreateInstanceActionWithContext(context, createinsactoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceActionWithContext failed: %s", err.Error()), "ibm_is_instance_action", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if actiontype == "stop" {
		_, err = isWaitForInstanceActionStop(sess, d.Timeout(schema.TimeoutUpdate), instanceId, d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceActionStop failed: %s", err.Error()), "ibm_is_instance_action", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else if actiontype == "start" || actiontype == "reboot" {
		_, err = isWaitForInstanceActionStart(sess, d.Timeout(schema.TimeoutUpdate), instanceId, d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceActionStart failed: %s", err.Error()), "ibm_is_instance_action", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId(instanceId)
	return resourceIBMISInstanceActionRead(context, d, meta)
}

func resourceIBMISInstanceActionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_action", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := d.Id()

	options := &vpcv1.GetInstanceOptions{
		ID: &id,
	}
	instance, response, err := sess.GetInstanceWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(isInstanceStatus, *instance.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance", "read", "set-status").GetDiag()
	}
	statusReasonsList := make([]map[string]interface{}, 0)
	if instance.StatusReasons != nil {
		for _, sr := range instance.StatusReasons {
			currentSR := map[string]interface{}{}
			if sr.Code != nil && sr.Message != nil {
				currentSR[isInstanceStatusReasonsCode] = *sr.Code
				currentSR[isInstanceStatusReasonsMessage] = *sr.Message
				if sr.MoreInfo != nil {
					currentSR[isInstanceStatusReasonsMoreInfo] = *sr.MoreInfo
				}
				statusReasonsList = append(statusReasonsList, currentSR)
			}
		}
	}
	if err = d.Set(isInstanceStatusReasons, statusReasonsList); err != nil {
		err = fmt.Errorf("Error setting status_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance", "read", "set-status_reasons").GetDiag()
	}
	return nil
}

func resourceIBMISInstanceActionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_action", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, actiontypeIntf := d.GetChange(isInstanceAction)
	actiontype := actiontypeIntf.(string)
	id := d.Id()

	getinsOptions := &vpcv1.GetInstanceOptions{
		ID: &id,
	}
	instance, response, err := sess.GetInstanceWithContext(context, getinsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_action", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if (actiontype == "stop" || actiontype == "reboot") && *instance.Status != isInstanceStatusRunning {
		d.Set(isInstanceAction, nil)
		err = fmt.Errorf("Error with stop/reboot action: Cannot invoke stop/reboot action while instance is not in running state")
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_action", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	} else if actiontype == "start" && *instance.Status != isInstanceActionStatusStopped {
		d.Set(isInstanceAction, nil)
		err = fmt.Errorf("Error with start action: Cannot invoke start action while instance is not in stopped state")
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "ibm_is_instance_action", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	createinsactoptions := &vpcv1.CreateInstanceActionOptions{
		InstanceID: &id,
		Type:       &actiontype,
	}
	_, response, err = sess.CreateInstanceActionWithContext(context, createinsactoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceActionWithContext failed: %s", err.Error()), "ibm_is_instance_action", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if actiontype == "stop" {
		_, err = isWaitForInstanceActionStop(sess, d.Timeout(schema.TimeoutUpdate), id, d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceActionStop failed: %s", err.Error()), "ibm_is_instance_action", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else if actiontype == "start" || actiontype == "reboot" {
		_, err = isWaitForInstanceActionStart(sess, d.Timeout(schema.TimeoutUpdate), id, d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForInstanceActionStart failed: %s", err.Error()), "ibm_is_instance_action", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return nil
}

func resourceIBMISInstanceActionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func resourceIBMISInstanceActionExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_action", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	id := d.Id()
	getInstanceOptions := &vpcv1.GetInstanceOptions{
		ID: &id,
	}
	_, response, err := sess.GetInstance(getInstanceOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstance failed: %s", err.Error()), "ibm_is_instance_action", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, fmt.Errorf("[ERROR] Error getting instance : %s\n%s", err, response)
	}
	return true, err
}
