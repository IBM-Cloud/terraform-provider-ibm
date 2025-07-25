// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPIVirtualSerialNumber() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVirtualSerialNumberCreate,
		ReadContext:   resourceIBMPIVirtualSerialNumberRead,
		UpdateContext: resourceIBMPIVirtualSerialNumberUpdate,
		DeleteContext: resourceIBMPIVirtualSerialNumberDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(45 * time.Minute),
			Update: schema.DefaultTimeout(45 * time.Minute),
			Delete: schema.DefaultTimeout(45 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description: "This is the Power Instance id that is assigned to the account",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			Arg_Description: {
				Computed:    true,
				Description: "Description of virtual serial number.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			Arg_InstanceID: {
				Description: "PVM Instance to attach VSN to.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			Arg_RetainVirtualSerialNumber: {
				Description:  "Indicates whether to retain virtual serial number after unassigning from PVM instance during deletion.",
				Optional:     true,
				RequiredWith: []string{Arg_InstanceID},
				Type:         schema.TypeBool,
			},
			Arg_Serial: {
				Description:      "Virtual serial number.",
				DiffSuppressFunc: supressVSNDiffAutoAssign,
				ForceNew:         true,
				Required:         true,
				Type:             schema.TypeString,
			},
			Arg_SoftwareTier: {
				Computed:     true,
				Description:  "Software tier for virtual serial number.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"P05", "P10", "P20", "P30"}, false),
			},
		},
	}
}

func resourceIBMPIVirtualSerialNumberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	client := instance.NewIBMPIVSNClient(ctx, sess, cloudInstanceID)

	vsnArg := d.Get(Arg_Serial).(string)
	if _, ok := d.GetOk(Arg_InstanceID); !ok {
		if vsnArg == AutoAssign {
			err := flex.FmtErrorf("cannot use '%s' unless '%s' is specified", AutoAssign, Arg_InstanceID)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("operation failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if _, ok := d.GetOk(Arg_SoftwareTier); ok {
			err := flex.FmtErrorf("cannot use '%s' unless '%s' is specified", Arg_SoftwareTier, Arg_InstanceID)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("operation failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	serialString := ""
	oldPvmInstanceId := ""
	if vsnArg != AutoAssign {
		vsn, err := client.Get(vsnArg)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Get failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		description := ""
		if v, ok := d.GetOk(Arg_Description); ok {
			description = v.(string)
		}
		if vsn.PvmInstanceID != nil {
			oldPvmInstanceId = *vsn.PvmInstanceID
			if v, ok := d.GetOk(Arg_InstanceID); ok && v.(string) == oldPvmInstanceId && description != "" && description != *vsn.Description {
				updateBody := &models.UpdateServerVirtualSerialNumber{
					Description: &description,
				}
				_, err := client.PVMInstanceUpdateVSN(oldPvmInstanceId, updateBody)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceUpdateVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			} else if v, ok := d.GetOk(Arg_InstanceID); !ok || (ok && v.(string) != oldPvmInstanceId) {
				err := flex.FmtErrorf("please unassign virtual serial number %s from current pvm instance %s or specify \"%s\" for %s", *vsn.Serial, oldPvmInstanceId, oldPvmInstanceId, Arg_InstanceID)
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("operation failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		} else {
			// Update description if not attaching to a new VM
			if _, ok := d.GetOk(Arg_InstanceID); !ok && description != "" && description != *vsn.Description {
				updateBody := &models.UpdateVirtualSerialNumber{
					Description: &description,
				}
				_, err := client.Update(vsnArg, updateBody)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Update failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}

		serialString = vsnArg
	}

	if pvmInstanceId, ok := d.GetOk(Arg_InstanceID); ok {
		pvmInstanceIdArg := pvmInstanceId.(string)
		instanceClient := instance.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
		restartInstance := false
		// if specified instance was not attached before
		if oldPvmInstanceId == "" {
			restartInstance, err = stopLparForVSNChange(ctx, instanceClient, pvmInstanceIdArg, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("stopLparForVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			serialNumber := d.Get(Arg_Serial).(string)
			addBody := &models.AddServerVirtualSerialNumber{
				Serial: &serialNumber,
			}
			if v, ok := d.GetOk(Arg_Description); ok {
				addBody.Description = v.(string)
			}
			_, err = client.PVMInstanceAttachVSN(pvmInstanceIdArg, addBody)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceAttachVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

				err = instanceRestartAfterVSNFailure(ctx, pvmInstanceIdArg, restartInstance, instanceClient, d, err)
				if err != nil {
					tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
					log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
				}

				return tfErr.GetDiag()
			}

			_, err = isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped(ctx, instanceClient, pvmInstanceIdArg, nil, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

				err = instanceRestartAfterVSNFailure(ctx, pvmInstanceIdArg, restartInstance, instanceClient, d, err)
				if err != nil {
					tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
					log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
				}

				return tfErr.GetDiag()
			}
		}

		if vsnArg == AutoAssign {
			vsns, err := client.GetAll(&pvmInstanceIdArg)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAll failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

				err = instanceRestartAfterVSNFailure(ctx, pvmInstanceIdArg, restartInstance, instanceClient, d, err)
				if err != nil {
					tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
					log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
				}

				return tfErr.GetDiag()
			}
			serialString = *vsns[0].Serial
		} else {
			serialString = vsnArg
		}

		id := cloudInstanceID + "/" + serialString
		d.SetId(id)

		if v, ok := d.GetOk(Arg_SoftwareTier); ok {
			// Need to get current vsn for software tier
			vsnObj, err := client.Get(serialString)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Get failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			softwareTierCurrent := vsnObj.SoftwareTier
			softwareTier := models.SoftwareTier(v.(string))
			if softwareTierCurrent != softwareTier {
				restartInstanceSoftwareTierUpdate, err := stopLparForVSNChange(ctx, instanceClient, pvmInstanceIdArg, d.Timeout(schema.TimeoutCreate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("stopLparForVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				restartInstance = restartInstance || restartInstanceSoftwareTierUpdate
				// Update software tier to match configuration
				updateBody := &models.UpdateServerVirtualSerialNumber{
					SoftwareTier: softwareTier,
				}
				_, err = client.PVMInstanceUpdateVSN(pvmInstanceIdArg, updateBody)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceUpdateVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

					err = instanceRestartAfterVSNFailure(ctx, pvmInstanceIdArg, restartInstance, instanceClient, d, err)
					if err != nil {
						tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
						log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
					}

					return tfErr.GetDiag()
				}
				_, err = isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped(ctx, instanceClient, pvmInstanceIdArg, updateBody, d.Timeout(schema.TimeoutCreate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

					err = instanceRestartAfterVSNFailure(ctx, pvmInstanceIdArg, restartInstance, instanceClient, d, err)
					if err != nil {
						tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
						log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
					}

					return tfErr.GetDiag()
				}
			}
		}

		if restartInstance {
			err = startLparAfterVSNChange(ctx, instanceClient, pvmInstanceIdArg, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("startLparAfterVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
	} else {
		// Set here incase above if statement block does not run
		id := cloudInstanceID + "/" + serialString
		d.SetId(id)
	}

	return resourceIBMPIVirtualSerialNumberRead(ctx, d, meta)
}

func resourceIBMPIVirtualSerialNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	idArr, err := flex.IdParts(d.Id())
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IdParts failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := idArr[0]
	serial := idArr[1]
	client := instance.NewIBMPIVSNClient(ctx, sess, cloudInstanceID)

	vsn, err := client.Get(serial)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), NotFound) {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Get failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.Set(Arg_Description, vsn.Description)
	d.Set(Arg_InstanceID, vsn.PvmInstanceID)
	d.Set(Arg_Serial, vsn.Serial)
	d.Set(Arg_SoftwareTier, vsn.SoftwareTier)

	return nil
}

func resourceIBMPIVirtualSerialNumberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	idArr, err := flex.IdParts(d.Id())
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IdParts failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	cloudInstanceID := idArr[0]
	client := instance.NewIBMPIVSNClient(ctx, sess, cloudInstanceID)

	if v, ok := d.GetOk(Arg_InstanceID); ok {
		pvmInstanceId := v.(string)
		instanceClient := instance.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
		restartInstance, err := stopLparForVSNChange(ctx, instanceClient, pvmInstanceId, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("stopLparForVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		retainVSN := false
		if v, ok := d.GetOk(Arg_RetainVirtualSerialNumber); ok {
			retainVSN = v.(bool)
		}
		deleteBody := &models.DeleteServerVirtualSerialNumber{
			RetainVSN: retainVSN,
		}
		err = client.PVMInstanceDeleteVSN(pvmInstanceId, deleteBody)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceDeleteVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

			err = instanceRestartAfterVSNFailure(ctx, pvmInstanceId, restartInstance, instanceClient, d, err)
			if err != nil {
				tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
				log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
			}

			return tfErr.GetDiag()
		}

		_, err = isWaitForPIInstanceVSNRemoved(ctx, instanceClient, pvmInstanceId, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForPIInstanceVSNRemoved failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

			err = instanceRestartAfterVSNFailure(ctx, pvmInstanceId, restartInstance, instanceClient, d, err)
			if err != nil {
				tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
				log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
			}

			return tfErr.GetDiag()
		}

		if restartInstance {
			err = startLparAfterVSNChange(ctx, instanceClient, pvmInstanceId, d.Timeout(schema.TimeoutDelete))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("startLparAfterVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}

	} else {
		// TODO: allow user to use pi_retain_virtual_serial_number to retain VSN instead of deleting it when not attached to PVM instance
		serialNumber := d.Get(Arg_Serial).(string)
		err = client.Delete(serialNumber)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Delete failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId("")

	return nil
}

func resourceIBMPIVirtualSerialNumberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	client := instance.NewIBMPIVSNClient(ctx, sess, cloudInstanceID)
	instanceClient := instance.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)

	if (d.HasChange(Arg_Description) || d.HasChange(Arg_SoftwareTier)) && !d.HasChange(Arg_InstanceID) {
		if _, ok := d.GetOk(Arg_InstanceID); !ok {
			if _, ok := d.GetOk(Arg_SoftwareTier); ok {
				// Set old software tier back into state incase it is not re-read on next apply
				oldSoftwareTier, _ := d.GetChange(Arg_SoftwareTier)
				d.Set(Arg_SoftwareTier, oldSoftwareTier)

				err = flex.FmtErrorf("cannot set '%s' unless '%s' is specified", Arg_SoftwareTier, Arg_InstanceID)
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("operation failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}

		// if instance ID is defined, need to use correct endpoints for assigned vsns
		if v, ok := d.GetOk(Arg_InstanceID); ok {
			// Cannot update software tier and description at the same time, so two separate update calls must be used
			pvmInstanceId := v.(string)
			restartInstance := false
			if d.HasChange(Arg_Description) {
				newDescription := d.Get(Arg_Description).(string)
				updateBody := &models.UpdateServerVirtualSerialNumber{
					Description: &newDescription,
				}
				_, err = client.PVMInstanceUpdateVSN(pvmInstanceId, updateBody)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceUpdateVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

					err = instanceRestartAfterVSNFailure(ctx, pvmInstanceId, restartInstance, instanceClient, d, err)
					if err != nil {
						tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
						log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
					}

					return tfErr.GetDiag()
				}
			}

			if d.HasChange(Arg_SoftwareTier) {
				restartInstanceSoftwareTier, err := stopLparForVSNChange(ctx, instanceClient, pvmInstanceId, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("stopLparForVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				restartInstance = restartInstance || restartInstanceSoftwareTier

				newSoftwareTier := models.SoftwareTier(d.Get(Arg_SoftwareTier).(string))
				updateBody := &models.UpdateServerVirtualSerialNumber{
					SoftwareTier: newSoftwareTier,
				}

				_, err = client.PVMInstanceUpdateVSN(pvmInstanceId, updateBody)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceUpdateVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

					err = instanceRestartAfterVSNFailure(ctx, pvmInstanceId, restartInstance, instanceClient, d, err)
					if err != nil {
						tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
						log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
					}

					return tfErr.GetDiag()
				}

				_, err = isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped(ctx, instanceClient, pvmInstanceId, updateBody, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

					err = instanceRestartAfterVSNFailure(ctx, pvmInstanceId, restartInstance, instanceClient, d, err)
					if err != nil {
						tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
						log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
					}

					return tfErr.GetDiag()
				}
			}

			if restartInstance {
				err = startLparAfterVSNChange(ctx, instanceClient, pvmInstanceId, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("startLparAfterVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
			// else, use endpoints for unassigned vsns
		} else if d.HasChange(Arg_Description) {
			newDescription := d.Get(Arg_Description).(string)
			updateBody := &models.UpdateVirtualSerialNumber{
				Description: &newDescription,
			}

			vsnArg := d.Get(Arg_Serial).(string)

			_, err = client.Update(vsnArg, updateBody)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Update failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
	}

	// if instance ID is changed, detach from old lpar (if present) and re-attach to new lpar (if present)
	if d.HasChange(Arg_InstanceID) {
		oldId, newId := d.GetChange(Arg_InstanceID)
		oldIdString, newIdString := oldId.(string), newId.(string)

		if _, ok := d.GetOk(Arg_SoftwareTier); ok && newIdString == "" {
			err = flex.FmtErrorf("cannot set '%s' unless '%s' is specified", Arg_SoftwareTier, Arg_InstanceID)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("operation failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		if oldIdString != "" {
			restartInstance, err := stopLparForVSNChange(ctx, instanceClient, oldIdString, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("stopLparForVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}

			detachBody := &models.DeleteServerVirtualSerialNumber{
				RetainVSN: true,
			}
			err = client.PVMInstanceDeleteVSN(oldIdString, detachBody)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceDeleteVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

				err = instanceRestartAfterVSNFailure(ctx, oldIdString, restartInstance, instanceClient, d, err)
				if err != nil {
					tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
				}

				return tfErr.GetDiag()
			}

			_, err = isWaitForPIInstanceVSNRemoved(ctx, instanceClient, oldIdString, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForPIInstanceVSNRemoved failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

				err = instanceRestartAfterVSNFailure(ctx, oldIdString, restartInstance, instanceClient, d, err)
				if err != nil {
					tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
				}

				return tfErr.GetDiag()
			}

			if restartInstance {
				err = startLparAfterVSNChange(ctx, instanceClient, oldIdString, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("startLparAfterVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}

		if newIdString != "" {
			restartInstance, err := stopLparForVSNChange(ctx, instanceClient, newIdString, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("stopLparForVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}

			serial := d.Get(Arg_Serial).(string)
			addBody := &models.AddServerVirtualSerialNumber{
				Serial: &serial,
			}
			if v, ok := d.GetOk(Arg_Description); ok {
				description := v.(string)
				addBody.Description = description
			}
			_, err = client.PVMInstanceAttachVSN(newIdString, addBody)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceAttachVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

				err = instanceRestartAfterVSNFailure(ctx, newIdString, restartInstance, instanceClient, d, err)
				if err != nil {
					tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
				}

				return tfErr.GetDiag()
			}

			pvm, err := isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped(ctx, instanceClient, newIdString, nil, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

				err = instanceRestartAfterVSNFailure(ctx, newIdString, restartInstance, instanceClient, d, err)
				if err != nil {
					tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
				}

				return tfErr.GetDiag()
			}

			// cannot specify software tier when attaching new VSN, update to match software tier in configuration
			if v, ok := d.GetOk(Arg_SoftwareTier); ok {
				pvmInstance := pvm.(*models.PVMInstance)
				softwareTier := models.SoftwareTier(v.(string))
				if softwareTier != pvmInstance.VirtualSerialNumber.SoftwareTier {
					updateBody := &models.UpdateServerVirtualSerialNumber{
						SoftwareTier: softwareTier,
					}
					_, err = client.PVMInstanceUpdateVSN(newIdString, updateBody)
					if err != nil {
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PVMInstanceUpdateVSN failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

						err = instanceRestartAfterVSNFailure(ctx, newIdString, restartInstance, instanceClient, d, err)
						if err != nil {
							tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
							log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
						}

						return tfErr.GetDiag()
					}
					_, err = isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped(ctx, instanceClient, newIdString, updateBody, d.Timeout(schema.TimeoutUpdate))
					if err != nil {
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForPIInstanceVSNAssignedOrUpdatedAndStopped failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())

						err = instanceRestartAfterVSNFailure(ctx, newIdString, restartInstance, instanceClient, d, err)
						if err != nil {
							tfErr2 := flex.TerraformErrorf(err, fmt.Sprintf("instanceRestartAfterVSNFailure failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
							log.Printf("[DEBUG]\n%s", tfErr2.GetDebugMessage())
						}

						return tfErr.GetDiag()
					}
				}
			}

			if restartInstance {
				err = startLparAfterVSNChange(ctx, instanceClient, newIdString, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("startLparAfterVSNChange failed: %s", err.Error()), "ibm_pi_virtual_serial_number", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
	}

	return resourceIBMPIVirtualSerialNumberRead(ctx, d, meta)
}

func startLparAfterVSNChange(ctx context.Context, client *instance.IBMPIInstanceClient, id string, timeout time.Duration) error {
	body := &models.PVMInstanceAction{
		Action: flex.PtrToString(Action_Start),
	}
	err := client.Action(id, body)
	if err != nil {
		return flex.FmtErrorf("failed to perform the start action on the pvm instance %v", err)
	}

	_, err = isWaitForPIInstanceAvailable(ctx, client, id, OK, timeout)

	return err
}

func stopLparForVSNChange(ctx context.Context, client *instance.IBMPIInstanceClient, id string, timeout time.Duration) (bool, error) {
	instanceRestart := false
	ins, err := client.Get(id)
	if err != nil {
		return false, flex.FmtErrorf("failed to get pvm instance (%s): %v", id, err)
	}
	status := *ins.Status
	if strings.ToLower(status) != State_Shutoff {
		body := &models.PVMInstanceAction{
			Action: flex.PtrToString(Action_ImmediateShutdown),
		}
		err := client.Action(id, body)
		if err != nil {
			return false, flex.FmtErrorf("failed to perform the stop action on the pvm instance %v", err)
		}
		instanceRestart = true
	}

	_, err = isWaitForPIInstanceStopped(ctx, client, id, timeout)

	return instanceRestart, err
}
