// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

func ResourceIBMPIInstanceVpmemVolumes() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIInstanceVpmemVolumesCreate,
		ReadContext:   resourceIBMPIInstanceVpmemVolumesRead,
		DeleteContext: resourceIBMPIInstanceVpmemVolumesDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return flex.ResourcePowerUserTagsCustomizeDiff(diff)
			},
		),
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description: "This is the Power Instance id that is assigned to the account",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			Arg_PVMInstanceID: {
				Description: "PCloud PVM Instance ID.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			Arg_UserTags: {
				Description: "List of user tags.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    true,
				Optional:    true,
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},
			Arg_VPMEMVolumes: {
				Description: "Description of volume(s) to create.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Name: {
							Description: "Volume base name.",
							Required:    true,
							Type:        schema.TypeString,
						},
						Attr_Size: {
							Description: "Volume size (GB).",
							Required:    true,
							Type:        schema.TypeInt,
						},
					},
				},
				ForceNew: true,
				MaxItems: 4,
				MinItems: 1,
				Required: true,
				Type:     schema.TypeList,
			},

			// Attributes
			Attr_Volumes: vpmemVolumeSchema(),
		},
	}
}

func resourceIBMPIInstanceVpmemVolumesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	pvmInstanceID := d.Get(Arg_PVMInstanceID).(string)
	client := instance.NewIBMPIVPMEMClient(ctx, sess, cloudInstanceID)
	var body = &models.VPMemVolumeAttach{}
	if tags, ok := d.GetOk(Arg_UserTags); ok {
		body.UserTags = flex.FlattenSet(tags.(*schema.Set))
	}
	var vpmemVolumes []*models.VPMemVolumeCreate
	for _, v := range d.Get(Arg_VPMEMVolumes).([]interface{}) {
		vol := v.(map[string]interface{})
		vpmemVolume := resourceIBMPIInstanceVpmemVolumesMapToVpMemVolumeCreate(vol)
		vpmemVolumes = append(vpmemVolumes, vpmemVolume)
	}

	body.VpmemVolumes = vpmemVolumes
	volumes, err := client.CreatePvmVpmemVolumes(pvmInstanceID, body)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreatePvmVpmemVolumes failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := fmt.Sprintf("%s/%s", cloudInstanceID, pvmInstanceID)
	for _, vol := range volumes.Volumes {
		id += "/" + *vol.UUID
		_, err = isWaitForVpmemAvailable(ctx, client, pvmInstanceID, *vol.UUID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVpmemAvailable failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	d.SetId(id)

	return resourceIBMPIInstanceVpmemVolumesRead(ctx, d, meta)
}

func resourceIBMPIInstanceVpmemVolumesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SepIdParts failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	client := instance.NewIBMPIVPMEMClient(ctx, sess, parts[0])
	vpmemVolumes, err := client.GetAllPvmVpmemVolumes(parts[1])
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), NotFound) {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAllPvmVpmemVolumes failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	volumes := []map[string]any{}
	if vpmemVolumes.Volumes != nil {
		for _, volume := range vpmemVolumes.Volumes {
			vpmemVol := dataSourceIBMPIVPMEMVolumeToMap(volume, meta)
			volumes = append(volumes, vpmemVol)
		}
	}
	d.Set(Attr_Volumes, volumes)

	return nil
}

func resourceIBMPIInstanceVpmemVolumesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SepIdParts failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	client := instance.NewIBMPIVPMEMClient(ctx, sess, parts[0])
	for i := 2; i < len(parts); i++ {
		err := client.DeletePvmVpmemVolume(parts[1], parts[i])
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePvmVpmemVolume failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForVpmemDeleted(ctx, client, parts[1], parts[i], d.Timeout(schema.TimeoutDelete))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVpmemDeleted failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId("")

	return nil
}

func resourceIBMPIInstanceVpmemVolumesMapToVpMemVolumeCreate(modelMap map[string]interface{}) *models.VPMemVolumeCreate {
	model := &models.VPMemVolumeCreate{}
	model.Name = core.StringPtr(modelMap[Attr_Name].(string))
	model.Size = core.Int64Ptr(int64(modelMap[Attr_Size].(int)))
	return model
}

func isWaitForVpmemAvailable(ctx context.Context, client *instance.IBMPIVPMEMClient, instanceID, volID string, timeout time.Duration) (any, error) {
	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Configuring},
		Target:     []string{State_Active, State_Error},
		Refresh:    isVpmemRefreshFunc(client, instanceID, volID),
		Delay:      Timeout_Delay,
		MinTimeout: Retry_Delay,
		Timeout:    timeout,
	}

	return stateConf.WaitForStateContext(ctx)
}
func isVpmemRefreshFunc(client *instance.IBMPIVPMEMClient, instanceID, volID string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {

		vpmemVol, err := client.GetPvmVpmemVolume(instanceID, volID)
		if err != nil {
			return nil, "", flex.FmtErrorf("[ERROR] error getting vpmem %s", err)
		}

		if strings.ToLower(*vpmemVol.Status) == State_Active {
			return vpmemVol, State_Active, nil
		}
		if strings.ToLower(*vpmemVol.Status) == State_Error {
			return vpmemVol, *vpmemVol.Status, flex.FmtErrorf("[ERROR] vpmem is in error state: %s", err)
		}

		return vpmemVol, State_Configuring, nil
	}
}
func isWaitForVpmemDeleted(ctx context.Context, client *instance.IBMPIVPMEMClient, instanceID, volID string, timeout time.Duration) (any, error) {

	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Retry, State_Deleting},
		Target:     []string{State_NotFound},
		Refresh:    isVpmemDeleteRefreshFunc(client, instanceID, volID),
		Delay:      Timeout_Delay,
		MinTimeout: Retry_Delay,
		Timeout:    timeout,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isVpmemDeleteRefreshFunc(client *instance.IBMPIVPMEMClient, instanceID, volID string) retry.StateRefreshFunc {
	return func() (any, string, error) {
		vpmemVol, err := client.GetPvmVpmemVolume(instanceID, volID)
		if err != nil {
			return vpmemVol, State_NotFound, nil
		}
		return vpmemVol, State_Deleting, nil
	}
}
