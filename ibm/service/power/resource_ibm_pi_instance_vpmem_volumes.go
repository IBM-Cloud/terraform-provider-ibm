// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMPIInstanceVpmemVolumes() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIInstanceVpmemVolumesCreate,
		ReadContext:   resourceIBMPIInstanceVpmemVolumesRead,
		UpdateContext: resourceIBMPIInstanceVpmemVolumesUpdate,
		DeleteContext: resourceIBMPIInstanceVpmemVolumesDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v any) error {
				return flex.ResourcePowerUserTagsCustomizeDiff(diff)
			},

			// When volumes are renamed, propagate only the name change into the
			// computed Attr_Volumes so Terraform shows a precise diff (old→new)
			// rather than marking every volume as (known after apply).
			func(_ context.Context, diff *schema.ResourceDiff, v any) error {
				if !diff.HasChange(Arg_VPMEMVolumes) {
					return nil
				}
				old, new := diff.GetChange(Arg_VPMEMVolumes)
				oldList := old.([]any)
				newList := new.([]any)

				// if adding then set volumes.
				if len(newList) > len(oldList) {
					return diff.SetNewComputed(Attr_Volumes)
				}
				// This when removing. I try to figure out how to only have that volume remove.
				// All I have is a place change on apply.
				if len(newList) < len(oldList) {
					newNameSet := make(map[string]bool)
					for _, v := range newList {
						newNameSet[v.(map[string]any)[Attr_Name].(string)] = true
					}
					oldVolumes, _ := diff.GetChange(Attr_Volumes)
					oldSet := oldVolumes.(*schema.Set)
					filtered := make([]map[string]any, 0, len(newList))
					for _, elem := range oldSet.List() {
						vol := elem.(map[string]any)
						if name, ok := vol[Attr_Name].(string); ok && newNameSet[name] {
							filtered = append(filtered, vol)
						}
					}
					return diff.SetNew(Attr_Volumes, filtered)
				}
				// TypeList preserves index order, so old[i] always corresponds to new[i].
				renameMap := make(map[string]string) // old name -> new name
				for i, v := range newList {
					oldName := oldList[i].(map[string]any)[Attr_Name].(string)
					newName := v.(map[string]any)[Attr_Name].(string)
					if oldName != newName {
						renameMap[oldName] = newName
					}
				}
				if len(renameMap) == 0 {
					return nil
				}

				// Apply renames to the current Attr_Volumes state.
				currentSet := diff.Get(Attr_Volumes).(*schema.Set)
				updated := make([]map[string]any, 0, currentSet.Len())
				for _, elem := range currentSet.List() {
					vol := elem.(map[string]any)
					vpmem := make(map[string]any, len(vol))
					maps.Copy(vpmem, vol)
					if name, ok := vol[Attr_Name].(string); ok {
						if newName, ok := renameMap[name]; ok {
							vpmem[Attr_Name] = newName
						}
					}
					updated = append(updated, vpmem)
				}
				return diff.SetNew(Attr_Volumes, updated)
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
							Description: "Volume size (GiB).",
							Required:    true,
							Type:        schema.TypeInt,
						},
						Attr_VolumeID: {
							Computed:    true,
							Description: "Volume ID.",
							Type:        schema.TypeString,
						},
					},
				},
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

func resourceIBMPIInstanceVpmemVolumesCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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

	var vpmemList []any
	if v, ok := d.GetOk(Arg_VPMEMVolumes); ok {
		vpmemList = v.([]any)
	}

	var vpmemVolumes []*models.VPMemVolumeCreate
	for _, v := range vpmemList {
		vol := v.(map[string]any)
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

func resourceIBMPIInstanceVpmemVolumesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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
	volIDMap := make(map[string]string)
	if vpmemVolumes.Volumes != nil {
		for _, vol := range vpmemVolumes.Volumes {
			if vol.Name != nil && vol.UUID != nil {
				volIDMap[*vol.Name] = *vol.UUID
			}
		}
	}
	vpmemList := d.Get(Arg_VPMEMVolumes).([]any)
	updatedVpmem := make([]map[string]any, 0, len(vpmemList))
	for _, v := range vpmemList {
		vol := v.(map[string]any)
		vpmem := map[string]any{
			Attr_Name: vol[Attr_Name],
			Attr_Size: vol[Attr_Size],
		}
		if id, ok := volIDMap[vol[Attr_Name].(string)]; ok {
			vpmem[Attr_VolumeID] = id
		}
		updatedVpmem = append(updatedVpmem, vpmem)
	}
	d.Set(Arg_VPMEMVolumes, updatedVpmem)

	volumes := []map[string]any{}
	if vpmemVolumes.Volumes != nil {
		for _, volume := range vpmemVolumes.Volumes {
			volumes = append(volumes, dataSourceIBMPIVPMEMVolumeToMap(volume, meta))
		}
	}
	d.Set(Attr_Volumes, volumes)

	return nil
}

func resourceIBMPIInstanceVpmemVolumesUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SepIdParts failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	client := instance.NewIBMPIVPMEMClient(ctx, sess, parts[0])

	if d.HasChange(Arg_VPMEMVolumes) {
		old, new := d.GetChange(Arg_VPMEMVolumes)
		oldList := old.([]any)
		newList := new.([]any)

		toCreate, deletedIDs, diagErr := processVpmemVolumeUpdates(ctx, client, parts[1], oldList, newList, d.Timeout(schema.TimeoutUpdate), "ibm_pi_instance_vpmem_volumes")
		if diagErr != nil {
			return diagErr
		}
		// This block fixes the ID of the resource which should be the
		// vpmem volume UUIDs stapled together
		if len(deletedIDs) > 0 {
			newParts := parts[:2]
			for _, p := range parts[2:] {
				if !deletedIDs[p] {
					newParts = append(newParts, p)
				}
			}
			d.SetId(strings.Join(newParts, "/"))
		}

		// Create brand-new volumes.
		for _, newVol := range toCreate {
			created, err := client.CreatePvmVpmemVolumes(parts[1], &models.VPMemVolumeAttach{
				VpmemVolumes: []*models.VPMemVolumeCreate{
					resourceIBMPIInstanceVpmemVolumesMapToVpMemVolumeCreate(newVol),
				},
			})
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreatePvmVpmemVolumes failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			for _, vol := range created.Volumes {
				if _, err = isWaitForVpmemAvailable(ctx, client, parts[1], *vol.UUID, d.Timeout(schema.TimeoutUpdate)); err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVpmemAvailable failed: %s", err.Error()), "ibm_pi_instance_vpmem_volumes", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				d.SetId(d.Id() + "/" + *vol.UUID)
			}
		}
	}

	return resourceIBMPIInstanceVpmemVolumesRead(ctx, d, meta)
}

func resourceIBMPIInstanceVpmemVolumesDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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

func resourceIBMPIInstanceVpmemVolumesMapToVpMemVolumeCreate(modelMap map[string]any) *models.VPMemVolumeCreate {
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
	return func() (any, string, error) {

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

func isWaitForVpmemUpdated(ctx context.Context, client *instance.IBMPIVPMEMClient, instanceID string, newNames []string, timeout time.Duration) (any, error) {

	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Updating},
		Target:     []string{State_Completed},
		Refresh:    isVpmemUpdateRefreshFunc(client, instanceID, newNames),
		Delay:      Timeout_Delay,
		MinTimeout: Retry_Delay,
		Timeout:    timeout,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isVpmemUpdateRefreshFunc(client *instance.IBMPIVPMEMClient, instanceID string, newNames []string) retry.StateRefreshFunc {
	return func() (any, string, error) {
		vpmemVolumes, err := client.GetAllPvmVpmemVolumes(instanceID)
		if err != nil {
			return nil, "", flex.FmtErrorf("[ERROR] error getting vpmem volumes: %s", err)
		}
		numFound := 0
		for _, vpmemVolume := range vpmemVolumes.Volumes {
			if slices.Contains(newNames, *vpmemVolume.Name) {
				numFound++
			}
		}
		if numFound == len(newNames) {
			return vpmemVolumes, State_Completed, nil
		}
		return vpmemVolumes, State_Updating, nil
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

// processVpmemVolumeUpdates handles the common logic for updating vPMEM volumes
// It processes renames, deletions, and creations based on the old and new configurations
func processVpmemVolumeUpdates(ctx context.Context, client *instance.IBMPIVPMEMClient, instanceID string, oldList, newList []any, timeout time.Duration, resourceName string) (toCreate []map[string]any, deletedIDs map[string]bool, diagErr diag.Diagnostics) {
	// Build lookup tables from old state.
	oldNameToID := make(map[string]string) // name -> volume_id
	oldNameToSize := make(map[string]int)  // name -> size
	oldIDToSize := make(map[string]int)    // volume_id -> size
	for _, v := range oldList {
		vol := v.(map[string]any)
		name := vol[Attr_Name].(string)
		id := vol[Attr_VolumeID].(string)
		size := vol[Attr_Size].(int)
		oldNameToID[name] = id
		oldNameToSize[name] = size
		if id != "" {
			oldIDToSize[id] = size
		}
	}

	newNameSet := make(map[string]bool)
	for _, v := range newList {
		newNameSet[v.(map[string]any)[Attr_Name].(string)] = true
	}

	// Volumes whose name is not in the new config are candidates for deletion or rename.
	// Build the set keyed by volume_id.
	toDelete := make(map[string]bool)
	for oldName, oldID := range oldNameToID {
		if !newNameSet[oldName] && oldID != "" {
			toDelete[oldID] = true
		}
	}

	// Process each new entry.
	var updatedNames []string
	toCreate = []map[string]any{}
	for i, v := range newList {
		newVol := v.(map[string]any)
		newName := newVol[Attr_Name].(string)
		newSize := newVol[Attr_Size].(int)

		if oldSize, kept := oldNameToSize[newName]; kept {
			// Existing volume — only size changes are rejected.
			if oldSize != newSize {
				opErr := flex.FmtErrorf("%s cannot be updated", Attr_Size)
				tfErr := flex.TerraformErrorf(opErr, fmt.Sprintf("operation failed: %s", opErr.Error()), resourceName, "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return nil, nil, tfErr.GetDiag()
			}
			continue
		}

		// New name: either a rename or a brand-new volume.
		// Use the index-based volume_id as the primary rename
		// if the volume that was at this index is being deleted and has the same size,
		// it is being renamed to newName.
		var renameID string
		if i < len(oldList) {
			idAtIndex := oldList[i].(map[string]any)[Attr_VolumeID].(string)
			if toDelete[idAtIndex] && oldIDToSize[idAtIndex] == newSize {
				renameID = idAtIndex
				delete(toDelete, idAtIndex)
			}
		}
		// Fall back to size-based matching among remaining deletion candidates
		// handles removal-from-middle where index pointed to the wrong volume.
		if renameID == "" {
			for delID := range toDelete {
				if oldIDToSize[delID] == newSize {
					renameID = delID
					delete(toDelete, delID)
					break
				}
			}
		}

		if renameID != "" {
			err := client.UpdatePvmVpmemVolume(instanceID, renameID, &models.VPMemVolumeUpdate{
				Name: flex.PtrToString(newName),
			})
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdatePvmVpmemVolume failed: %s", err.Error()), resourceName, "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return nil, nil, tfErr.GetDiag()
			}
			updatedNames = append(updatedNames, newName)
		} else {
			toCreate = append(toCreate, newVol)
		}
	}

	// Wait for all renames to propagate.
	if len(updatedNames) > 0 {
		if _, err := isWaitForVpmemUpdated(ctx, client, instanceID, updatedNames, timeout); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVpmemUpdated failed: %s", err.Error()), resourceName, "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return nil, nil, tfErr.GetDiag()
		}
	}

	// Delete volumes that were removed and not matched to a rename.
	deletedIDs = make(map[string]bool)
	for volID := range toDelete {
		err := client.DeletePvmVpmemVolume(instanceID, volID)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePvmVpmemVolume failed: %s", err.Error()), resourceName, "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return nil, nil, tfErr.GetDiag()
		}
		if _, err := isWaitForVpmemDeleted(ctx, client, instanceID, volID, timeout); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVpmemDeleted failed: %s", err.Error()), resourceName, "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return nil, nil, tfErr.GetDiag()
		}
		deletedIDs[volID] = true
	}

	return toCreate, deletedIDs, nil
}
