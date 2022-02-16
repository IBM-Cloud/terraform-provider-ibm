// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

const (
	//Added timeout values for warning  and active status
	warningTimeOut = 60 * time.Second
	activeTimeOut  = 2 * time.Minute
)

// Attributes and Arguments defined in data_source_ibm_pi_instance.go
func ResourceIBMPIInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIInstanceCreate,
		ReadContext:   resourceIBMPIInstanceRead,
		UpdateContext: resourceIBMPIInstanceUpdate,
		DeleteContext: resourceIBMPIInstanceDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(120 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			PICloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the Power Instance id that is assigned to the account",
			},
			PIInstanceLRC: {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The VTL license repository capacity TB value",
			},
			InstanceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PI instance status",
			},
			PIInstanceMigratable: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "set to true to enable migration of the PI instance",
			},
			InstanceMinProcessors: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Minimum number of the CPUs",
			},
			InstanceMinMemory: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Minimum memory",
			},
			InstanceMaxProcessors: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Maximum number of processors",
			},
			InstanceMaxMemory: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Maximum memory size",
			},
			PIInstanceVolumeIDs: {
				Type:             schema.TypeSet,
				Optional:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "List of PI volumes",
			},

			PIInstanceUserData: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base64 encoded data to be passed in for invoking a cloud init script",
			},

			PIInstanceStorageType: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Storage type for server deployment",
			},
			PIInstanceStoragePool: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Storage Pool for server deployment; if provided then pi_affinity_policy and pi_storage_type will be ignored",
			},
			PIInstanceAffinityPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Affinity policy for pvm instance being created; ignored if pi_storage_pool provided; for policy affinity requires one of pi_affinity_instance or pi_affinity_volume to be specified; for policy anti-affinity requires one of pi_anti_affinity_instances or pi_anti_affinity_volumes to be specified",
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"affinity", "anti-affinity"}),
			},
			PIInstanceAffinityVolume: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Volume (ID or Name) to base storage affinity policy against; required if requesting affinity and pi_affinity_instance is not provided",
				ConflictsWith: []string{PIInstanceAffinityInstance},
			},
			PIInstanceAffinityInstance: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "PVM Instance (ID or Name) to base storage affinity policy against; required if requesting storage affinity and pi_affinity_volume is not provided",
				ConflictsWith: []string{PIInstanceAffinityVolume},
			},
			PIInstanceAntiAffinityVolumes: {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of volumes to base storage anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_instances is not provided",
				ConflictsWith: []string{PIInstanceAntiAffinityInstances},
			},
			PIInstanceAntiAffinityInstances: {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of pvmInstances to base storage anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_volumes is not provided",
				ConflictsWith: []string{PIInstanceAntiAffinityVolumes},
			},
			PIInstanceStorageConnection: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"vSCSI"}),
				Description:  "Storage Connectivity Group for server deployment",
			},
			PIInstanceStoragePoolAffinity: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicates if all volumes attached to the server must reside in the same storage pool",
			},
			PIInstanceNetwork: {
				Type:             schema.TypeList,
				Required:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "List of one or more networks to attach to the instance",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PIInstanceNetworkIP: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						InstanceNetworkMAC: {
							Type:     schema.TypeString,
							Computed: true,
						},
						PIInstanceNetworkID: {
							Type:     schema.TypeString,
							Required: true,
						},
						InstanceNetworkName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceExternalIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			PIInstancePlacementGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Placement group ID",
			},
			InstanceHealthStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PI Instance health status",
			},
			InstanceInstanceID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance ID",
			},
			InstancePinPolicy: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PIN Policy of the Instance",
			},
			PIInstanceImageID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI instance image id",
			},
			PIInstanceProcessors: {
				Type:          schema.TypeFloat,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{PIInstanceSapProfileID},
				Description:   "Processors count",
			},
			PIInstanceName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI Instance name",
			},
			PIInstanceProcType: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ValidateFunc:  validate.ValidateAllowedStringValues([]string{"dedicated", "shared", "capped"}),
				ConflictsWith: []string{PIInstanceSapProfileID},
				Description:   "Instance processor type",
			},
			PIInstanceKey: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SSH key name",
			},
			PIInstanceMemory: {
				Type:          schema.TypeFloat,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{PIInstanceSapProfileID},
				Description:   "Memory size",
			},
			PIInstanceSapProfileID: {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{PIInstanceProcessors, PIInstanceMemory, PIInstanceProcType},
				Description:   "SAP Profile ID for the amount of cores and memory",
			},
			PIInstanceSysType: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "PI Instance system type",
			},
			PIInstanceReplicants: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "PI Instance replicas count",
			},
			PIInstanceReplicationPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"affinity", "anti-affinity", "none"}),
				Default:      "none",
				Description:  "Replication policy for the PI Instance",
			},
			PIInstanceReplicationScheme: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"prefix", "suffix"}),
				Default:      "suffix",
				Description:  "Replication scheme",
			},
			InstanceProgress: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Progress of the operation",
			},
			PIInstancePinPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Pin Policy of the instance",
				Default:      "none",
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"none", "soft", "hard"}),
			},

			// "reboot_for_resource_change": {
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Flag to be passed for CPU/Memory changes that require a reboot to take effect",
			// },
			InstanceOperatingSystem: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operating System",
			},
			InstanceOSType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OS Type",
			},
			PIInstanceHealth: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"OK", "WARNING"}),
				Default:      "OK",
				Description:  "Allow the user to set the status of the lpar so that they can connect to it faster",
			},
			PIInstanceVirtualCoresAssigned: {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Virtual Cores Assigned to the PVMInstance",
			},
			InstanceMaxVirtualCores: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum Virtual Cores Assigned to the PVMInstance",
			},
			InstanceMinVirtualCores: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum Virtual Cores Assigned to the PVMInstance",
			},
		},
	}
}

func resourceIBMPIInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("Now in the PowerVMCreate")
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	client := st.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
	sapClient := st.NewIBMPISAPInstanceClient(ctx, sess, cloudInstanceID)
	imageClient := st.NewIBMPIImageClient(ctx, sess, cloudInstanceID)

	var pvmList *models.PVMInstanceList
	if _, ok := d.GetOk(PIInstanceSapProfileID); ok {
		pvmList, err = createSAPInstance(d, sapClient)
	} else {
		pvmList, err = createPVMInstance(d, client, imageClient)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	var instanceReadyStatus string
	if r, ok := d.GetOk(PIInstanceHealth); ok {
		instanceReadyStatus = r.(string)
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *(*pvmList)[0].PvmInstanceID))

	for _, s := range *pvmList {
		_, err = isWaitForPIInstanceAvailable(ctx, client, *s.PvmInstanceID, instanceReadyStatus)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// If Storage Pool Affinity is given as false we need to update the vm instance.
	// Default value is true which indicates that all volumes attached to the server
	// must reside in the same storage pool.
	storagePoolAffinity := d.Get(PIInstanceStoragePoolAffinity).(bool)
	if !storagePoolAffinity {
		for _, s := range *pvmList {
			body := &models.PVMInstanceUpdate{
				StoragePoolAffinity: &storagePoolAffinity,
			}
			// This is a synchronous process hence no need to check for health status
			_, err = client.Update(*s.PvmInstanceID, body)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return resourceIBMPIInstanceRead(ctx, d, meta)

}

func resourceIBMPIInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, instanceID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
	powervmdata, err := client.Get(instanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set(PIInstanceMemory, powervmdata.Memory)
	d.Set(PIInstanceProcessors, powervmdata.Processors)
	if powervmdata.Status != nil {
		d.Set(InstanceStatus, powervmdata.Status)
	}
	d.Set(PIInstanceProcType, powervmdata.ProcType)
	if powervmdata.Migratable != nil {
		d.Set(PIInstanceMigratable, powervmdata.Migratable)
	}
	d.Set(InstanceMinProcessors, powervmdata.Minproc)
	d.Set(InstanceProgress, powervmdata.Progress)
	if powervmdata.StorageType != nil {
		d.Set(PIInstanceStorageType, powervmdata.StorageType)
	}
	d.Set(PIInstanceStoragePool, powervmdata.StoragePool)
	d.Set(PIInstanceStoragePoolAffinity, powervmdata.StoragePoolAffinity)
	d.Set(PICloudInstanceID, cloudInstanceID)
	d.Set(InstanceInstanceID, powervmdata.PvmInstanceID)
	d.Set(PIInstanceName, powervmdata.ServerName)
	d.Set(PIInstanceImageID, powervmdata.ImageID)
	if *powervmdata.PlacementGroup != "none" {
		d.Set(PIInstancePlacementGroup, powervmdata.PlacementGroup)
	}

	networksMap := []map[string]interface{}{}
	if powervmdata.Networks != nil {
		for _, n := range powervmdata.Networks {
			if n != nil {
				v := map[string]interface{}{
					InstanceNetworkIP:   n.IPAddress,
					InstanceNetworkMAC:  n.MacAddress,
					InstanceNetworkID:   n.NetworkID,
					InstanceNetworkName: n.NetworkName,
					InstanceNetworkType: n.Type,
					InstanceExternalIP:  n.ExternalIP,
				}
				networksMap = append(networksMap, v)
			}
		}
	}
	d.Set(InstanceNetwork, networksMap)

	if powervmdata.SapProfile != nil && powervmdata.SapProfile.ProfileID != nil {
		d.Set(PIInstanceSapProfileID, powervmdata.SapProfile.ProfileID)
	}
	d.Set(PIInstanceSysType, powervmdata.SysType)
	d.Set(InstanceMinMemory, powervmdata.Minmem)
	d.Set(InstanceMaxProcessors, powervmdata.Maxproc)
	d.Set(InstanceMaxMemory, powervmdata.Maxmem)
	d.Set(InstancePinPolicy, powervmdata.PinPolicy)
	d.Set(InstanceOperatingSystem, powervmdata.OperatingSystem)
	if powervmdata.OsType != nil {
		d.Set(InstanceOSType, powervmdata.OsType)
	}

	if powervmdata.Health != nil {
		d.Set(InstanceHealthStatus, powervmdata.Health.Status)
	}
	if powervmdata.VirtualCores != nil {
		d.Set(PIInstanceVirtualCoresAssigned, powervmdata.VirtualCores.Assigned)
		d.Set(InstanceMaxVirtualCores, powervmdata.VirtualCores.Max)
		d.Set(InstanceMinVirtualCores, powervmdata.VirtualCores.Min)
	}
	d.Set(PIInstanceLRC, powervmdata.LicenseRepositoryCapacity)

	return nil
}

func resourceIBMPIInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	name := d.Get(PIInstanceName).(string)
	mem := d.Get(PIInstanceMemory).(float64)
	procs := d.Get(PIInstanceProcessors).(float64)
	processortype := d.Get(PIInstanceProcType).(string)
	assignedVirtualCores := int64(d.Get(PIInstanceVirtualCoresAssigned).(int))

	if d.Get(InstanceHealthStatus) == "WARNING" {
		return diag.Errorf("the operation cannot be performed when the lpar health in the WARNING State")
	}

	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.Errorf("failed to get the session from the IBM Cloud Service")
	}

	cloudInstanceID, instanceID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)

	// Check if cloud instance is capable of changing virtual cores
	cloudInstanceClient := st.NewIBMPICloudInstanceClient(ctx, sess, cloudInstanceID)
	cloudInstance, err := cloudInstanceClient.Get(cloudInstanceID)
	if err != nil {
		return diag.FromErr(err)
	}
	cores_enabled := checkCloudInstanceCapability(cloudInstance, "custom-virtualcores")

	if d.HasChange(PIInstanceName) {
		body := &models.PVMInstanceUpdate{
			ServerName: name,
		}
		_, err = client.Update(instanceID, body)
		if err != nil {
			return diag.Errorf("failed to update the lpar with the change for name: %v", err)
		}
		_, err = isWaitForPIInstanceAvailable(ctx, client, instanceID, "OK")
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange(PIInstanceProcType) {

		// Stop the lpar
		if d.Get(InstanceStatus) == "SHUTOFF" {
			log.Printf("the lpar is in the shutoff state. Nothing to do . Moving on ")
		} else {
			err := stopLparForResourceChange(ctx, client, instanceID)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		// Modify
		log.Printf("At this point the lpar should be off. Executing the Processor Update Change")
		updatebody := &models.PVMInstanceUpdate{ProcType: processortype}
		if cores_enabled {
			log.Printf("support for %s is enabled", "custom-virtualcores")
			updatebody.VirtualCores = &models.VirtualCores{Assigned: &assignedVirtualCores}
		} else {
			log.Printf("no virtual cores support enabled for this customer..")
		}
		_, err = client.Update(instanceID, updatebody)
		if err != nil {
			return diag.FromErr(err)
		}
		_, err = isWaitForPIInstanceStopped(ctx, client, instanceID)
		if err != nil {
			return diag.FromErr(err)
		}

		// Start the lpar
		err := startLparAfterResourceChange(ctx, client, instanceID)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Virtual core will be updated only if service instance capability is enabled
	if d.HasChange(PIInstanceVirtualCoresAssigned) {
		body := &models.PVMInstanceUpdate{
			VirtualCores: &models.VirtualCores{Assigned: &assignedVirtualCores},
		}
		_, err = client.Update(instanceID, body)
		if err != nil {
			return diag.Errorf("failed to update the lpar with the change for virtual cores: %v", err)
		}
		_, err = isWaitForPIInstanceAvailable(ctx, client, instanceID, "OK")
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Start of the change for Memory and Processors
	if d.HasChange(PIInstanceMemory) || d.HasChange(PIInstanceProcessors) || d.HasChange(PIInstanceMigratable) {

		maxMemLpar := d.Get(InstanceMaxMemory).(float64)
		maxCPULpar := d.Get(InstanceMaxProcessors).(float64)
		//log.Printf("the required memory is set to [%d] and current max memory is set to  [%d] ", int(mem), int(maxMemLpar))

		if mem > maxMemLpar || procs > maxCPULpar {
			log.Printf("Will require a shutdown to perform the change")
		} else {
			log.Printf("maxMemLpar is set to %f", maxMemLpar)
			log.Printf("maxCPULpar is set to %f", maxCPULpar)
		}

		//if d.GetOkExists("reboot_for_resource_change")

		if mem > maxMemLpar || procs > maxCPULpar {

			err = performChangeAndReboot(ctx, client, instanceID, cloudInstanceID, mem, procs)
			if err != nil {
				return diag.FromErr(err)
			}

		} else {

			body := &models.PVMInstanceUpdate{
				Memory:     mem,
				Processors: procs,
			}
			if m, ok := d.GetOk(PIInstanceMigratable); ok {
				migratable := m.(bool)
				body.Migratable = &migratable
			}
			if cores_enabled {
				log.Printf("support for %s is enabled", "custom-virtualcores")
				body.VirtualCores = &models.VirtualCores{Assigned: &assignedVirtualCores}
			} else {
				log.Printf("no virtual cores support enabled for this customer..")
			}

			_, err = client.Update(instanceID, body)
			if err != nil {
				return diag.Errorf("failed to update the lpar with the change %v", err)
			}
			_, err = isWaitForPIInstanceAvailable(ctx, client, instanceID, "OK")
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	// License repository capacity will be updated only if service instance is a vtl instance
	// might need to check if lrc was set
	if d.HasChange(PIInstanceLRC) {

		lrc := d.Get(PIInstanceLRC).(int64)
		body := &models.PVMInstanceUpdate{
			LicenseRepositoryCapacity: lrc,
		}
		_, err = client.Update(instanceID, body)
		if err != nil {
			return diag.Errorf("failed to update the lpar with the change for license repository capacity %s", err)
		}
		_, err = isWaitForPIInstanceAvailable(ctx, client, instanceID, "OK")
		if err != nil {
			diag.FromErr(err)
		}
	}

	if d.HasChange(PIInstanceSapProfileID) {
		// Stop the lpar
		if d.Get(InstanceStatus) == "SHUTOFF" {
			log.Printf("the lpar is in the shutoff state. Nothing to do... Moving on ")
		} else {
			err := stopLparForResourceChange(ctx, client, instanceID)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		// Update the profile id
		profileID := d.Get(PIInstanceSapProfileID).(string)
		body := &models.PVMInstanceUpdate{
			SapProfileID: profileID,
		}
		_, err = client.Update(instanceID, body)
		if err != nil {
			return diag.Errorf("failed to update the lpar with the change for sap profile: %v", err)
		}

		// Wait for the resize to complete and status to reset
		_, err = isWaitForPIInstanceStopped(ctx, client, instanceID)
		if err != nil {
			return diag.FromErr(err)
		}

		// Start the lpar
		err := startLparAfterResourceChange(ctx, client, instanceID)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange(PIInstanceStoragePoolAffinity) {
		storagePoolAffinity := d.Get(PIInstanceStoragePoolAffinity).(bool)
		body := &models.PVMInstanceUpdate{
			StoragePoolAffinity: &storagePoolAffinity,
		}
		// This is a synchronous process hence no need to check for health status
		_, err = client.Update(instanceID, body)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange(PIInstancePlacementGroup) {

		pgClient := st.NewIBMPIPlacementGroupClient(ctx, sess, cloudInstanceID)

		oldRaw, newRaw := d.GetChange(PIInstancePlacementGroup)
		old := oldRaw.(string)
		new := newRaw.(string)

		if len(strings.TrimSpace(old)) > 0 {
			placementGroupID := old
			//remove server from old placement group
			body := &models.PlacementGroupServer{
				ID: &instanceID,
			}
			_, err := pgClient.DeleteMember(placementGroupID, body)
			if err != nil {
				// ignore delete member error where the server is already not in the PG
				if !strings.Contains(err.Error(), "is not part of placement-group") {
					return diag.FromErr(err)
				}
			}
		}

		if len(strings.TrimSpace(new)) > 0 {
			placementGroupID := new
			// add server to a new placement group
			body := &models.PlacementGroupServer{
				ID: &instanceID,
			}
			_, err := pgClient.AddMember(placementGroupID, body)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return resourceIBMPIInstanceRead(ctx, d, meta)

}

func resourceIBMPIInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID, instanceID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := st.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
	err = client.Delete(instanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = isWaitForPIInstanceDeleted(ctx, client, instanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func isWaitForPIInstanceDeleted(ctx context.Context, client *st.IBMPIInstanceClient, id string) (interface{}, error) {

	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", "DELETING"},
		Target:     []string{"Not Found"},
		Refresh:    isPIInstanceDeleteRefreshFunc(client, id),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
		Timeout:    10 * time.Minute,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isPIInstanceDeleteRefreshFunc(client *st.IBMPIInstanceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pvm, err := client.Get(id)
		if err != nil {
			log.Printf("The power vm does not exist")
			return pvm, "Not Found", nil
		}
		return pvm, "DELETING", nil
	}
}

func isWaitForPIInstanceAvailable(ctx context.Context, client *st.IBMPIInstanceClient, id string, instanceReadyStatus string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be available and active ", id)

	queryTimeOut := activeTimeOut
	if instanceReadyStatus == "WARNING" {
		queryTimeOut = warningTimeOut
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING", "BUILD", "WARNING"},
		Target:     []string{"ACTIVE", "OK", "ERROR", ""},
		Refresh:    isPIInstanceRefreshFunc(client, id, instanceReadyStatus),
		Delay:      30 * time.Second,
		MinTimeout: queryTimeOut,
		Timeout:    120 * time.Minute,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isPIInstanceRefreshFunc(client *st.IBMPIInstanceClient, id, instanceReadyStatus string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}
		// Check for `instanceReadyStatus` health status and also the final health status "OK"
		if *pvm.Status == "ACTIVE" && (pvm.Health.Status == instanceReadyStatus || pvm.Health.Status == "OK") {
			return pvm, "ACTIVE", nil
		}
		if *pvm.Status == "ERROR" {
			return pvm, *pvm.Status, fmt.Errorf("failed to create the lpar")
		}

		return pvm, "BUILD", nil
	}
}

func checkBase64(input string) error {
	_, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return fmt.Errorf("failed to check if input is base64 %s", err)
	}
	return err
}

func isWaitForPIInstanceStopped(ctx context.Context, client *st.IBMPIInstanceClient, id string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be stopped and powered off ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"STOPPING", "RESIZE", "VERIFY_RESIZE", "WARNING"},
		Target:     []string{"OK", "SHUTOFF"},
		Refresh:    isPIInstanceRefreshFuncOff(client, id),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute, // This is the time that the client will execute to check the status of the request
		Timeout:    30 * time.Minute,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isPIInstanceRefreshFuncOff(client *st.IBMPIInstanceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		log.Printf("Calling the check Refresh status of the pvm instance %s", id)
		pvm, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}
		if *pvm.Status == "SHUTOFF" && pvm.Health.Status == "OK" {
			return pvm, "SHUTOFF", nil
		}
		return pvm, "STOPPING", nil
	}
}

func stopLparForResourceChange(ctx context.Context, client *st.IBMPIInstanceClient, id string) error {
	body := &models.PVMInstanceAction{
		//Action: flex.PtrToString("stop"),
		Action: flex.PtrToString("immediate-shutdown"),
	}
	err := client.Action(id, body)
	if err != nil {
		return fmt.Errorf("failed to perform the stop action on the pvm instance %v", err)
	}

	_, err = isWaitForPIInstanceStopped(ctx, client, id)

	return err
}

// Start the lpar

func startLparAfterResourceChange(ctx context.Context, client *st.IBMPIInstanceClient, id string) error {
	body := &models.PVMInstanceAction{
		Action: flex.PtrToString("start"),
	}
	err := client.Action(id, body)
	if err != nil {
		return fmt.Errorf("failed to perform the start action on the pvm instance %v", err)
	}

	_, err = isWaitForPIInstanceAvailable(ctx, client, id, "OK")

	return err
}

// Stop / Modify / Start only when the lpar is off limits

func performChangeAndReboot(ctx context.Context, client *st.IBMPIInstanceClient, id, cloudInstanceID string, mem, procs float64) error {
	/*
		These are the steps
		1. Stop the lpar - Check if the lpar is SHUTOFF
		2. Once the lpar is SHUTOFF - Make the cpu / memory change - DUring this time , you can check for RESIZE and VERIFY_RESIZE as the transition states
		3. If the change is successful , the lpar state will be back in SHUTOFF
		4. Once the LPAR state is SHUTOFF , initiate the start again and check for ACTIVE + OK
	*/
	//Execute the stop

	log.Printf("Calling the stop lpar for Resource Change code ..")
	err := stopLparForResourceChange(ctx, client, id)
	if err != nil {
		return err
	}

	body := &models.PVMInstanceUpdate{
		Memory:     mem,
		Processors: procs,
	}

	_, updateErr := client.Update(id, body)
	if updateErr != nil {
		return fmt.Errorf("failed to update the lpar with the change, %s", updateErr)
	}

	_, err = isWaitforPIInstanceUpdate(ctx, client, id)
	if err != nil {
		return fmt.Errorf("failed to get an update from the Service after the resource change, %s", err)
	}

	// Now we can start the lpar
	log.Printf("Calling the start lpar After the  Resource Change code ..")
	err = startLparAfterResourceChange(ctx, client, id)
	if err != nil {
		return err
	}

	return nil

}

func isWaitforPIInstanceUpdate(ctx context.Context, client *st.IBMPIInstanceClient, id string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be SHUTOFF AFTER THE RESIZE Due to DLPAR Operation ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"RESIZE", "VERIFY_RESIZE"},
		Target:     []string{"ACTIVE", "SHUTOFF", "OK"},
		Refresh:    isPIInstanceShutAfterResourceChange(client, id),
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Minute,
		Timeout:    60 * time.Minute,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isPIInstanceShutAfterResourceChange(client *st.IBMPIInstanceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if *pvm.Status == "SHUTOFF" && pvm.Health.Status == "OK" {
			log.Printf("The lpar is now off after the resource change...")
			return pvm, "SHUTOFF", nil
		}

		return pvm, "RESIZE", nil
	}
}

func expandPVMNetworks(networks []interface{}) []*models.PVMInstanceAddNetwork {
	pvmNetworks := make([]*models.PVMInstanceAddNetwork, 0, len(networks))
	for _, v := range networks {
		network := v.(map[string]interface{})
		pvmInstanceNetwork := &models.PVMInstanceAddNetwork{
			IPAddress: network[PIInstanceNetworkIP].(string),
			NetworkID: flex.PtrToString(network[PIInstanceNetworkID].(string)),
		}
		pvmNetworks = append(pvmNetworks, pvmInstanceNetwork)
	}
	return pvmNetworks
}

func checkCloudInstanceCapability(cloudInstance *models.CloudInstance, custom_capability string) bool {
	log.Printf("Checking for the following capability %s", custom_capability)
	log.Printf("the instance features are %s", cloudInstance.Capabilities)
	for _, v := range cloudInstance.Capabilities {
		if v == custom_capability {
			return true
		}
	}
	return false
}
func createSAPInstance(d *schema.ResourceData, sapClient *st.IBMPISAPInstanceClient) (*models.PVMInstanceList, error) {

	name := d.Get(PIInstanceName).(string)
	profileID := d.Get(PIInstanceSapProfileID).(string)
	imageid := d.Get(PIInstanceImageID).(string)

	pvmNetworks := expandPVMNetworks(d.Get(PIInstanceNetwork).([]interface{}))

	var replicants int64
	if r, ok := d.GetOk(PIInstanceReplicants); ok {
		replicants = int64(r.(int))
	}
	var replicationpolicy string
	if r, ok := d.GetOk(PIInstanceReplicationPolicy); ok {
		replicationpolicy = r.(string)
	}
	var replicationNamingScheme string
	if r, ok := d.GetOk(PIInstanceReplicationScheme); ok {
		replicationNamingScheme = r.(string)
	}
	instances := &models.PVMInstanceMultiCreate{
		AffinityPolicy: &replicationpolicy,
		Count:          replicants,
		Numerical:      &replicationNamingScheme,
	}

	body := &models.SAPCreate{
		ImageID:   &imageid,
		Instances: instances,
		Name:      &name,
		Networks:  pvmNetworks,
		ProfileID: &profileID,
	}

	if v, ok := d.GetOk(PIInstanceVolumeIDs); ok {
		volids := flex.ExpandStringList((v.(*schema.Set)).List())
		if len(volids) > 0 {
			body.VolumeIDs = volids
		}
	}
	if p, ok := d.GetOk(PIInstancePinPolicy); ok {
		pinpolicy := p.(string)
		if d.Get(PIInstancePinPolicy) == "soft" || d.Get(PIInstancePinPolicy) == "hard" {
			body.PinPolicy = models.PinPolicy(pinpolicy)
		}
	}

	if v, ok := d.GetOk(PIInstanceKey); ok {
		sshkey := v.(string)
		body.SSHKeyName = sshkey
	}
	if u, ok := d.GetOk(PIInstanceUserData); ok {
		userData := u.(string)
		err := checkBase64(userData)
		if err != nil {
			log.Printf("Data is not base64 encoded")
			return nil, err
		}
		body.UserData = userData
	}
	if sys, ok := d.GetOk(PIInstanceSysType); ok {
		body.SysType = sys.(string)
	}

	if st, ok := d.GetOk(PIInstanceStorageType); ok {
		body.StorageType = st.(string)
	}
	if sp, ok := d.GetOk(PIInstanceStoragePool); ok {
		body.StoragePool = sp.(string)
	}

	if ap, ok := d.GetOk(PIInstanceAffinityPolicy); ok {
		policy := ap.(string)
		affinity := &models.StorageAffinity{
			AffinityPolicy: &policy,
		}

		if policy == "affinity" {
			if av, ok := d.GetOk(PIInstanceAffinityVolume); ok {
				afvol := av.(string)
				affinity.AffinityVolume = &afvol
			}
			if ai, ok := d.GetOk(PIInstanceAffinityInstance); ok {
				afins := ai.(string)
				affinity.AffinityPVMInstance = &afins
			}
		} else {
			if avs, ok := d.GetOk(PIInstanceAntiAffinityVolumes); ok {
				afvols := flex.ExpandStringList(avs.([]interface{}))
				affinity.AntiAffinityVolumes = afvols
			}
			if ais, ok := d.GetOk(PIInstanceAntiAffinityInstances); ok {
				afinss := flex.ExpandStringList(ais.([]interface{}))
				affinity.AntiAffinityPVMInstances = afinss
			}
		}
		body.StorageAffinity = affinity
	}

	pvmList, err := sapClient.Create(body)
	if err != nil {
		return nil, fmt.Errorf("failed to provision: %v", err)
	}
	if pvmList == nil {
		return nil, fmt.Errorf("failed to provision")
	}

	return pvmList, nil
}
func createPVMInstance(d *schema.ResourceData, client *st.IBMPIInstanceClient, imageClient *st.IBMPIImageClient) (*models.PVMInstanceList, error) {

	name := d.Get(PIInstanceName).(string)
	imageid := d.Get(PIInstanceImageID).(string)

	var mem, procs float64
	var systype, processortype string
	if v, ok := d.GetOk(PIInstanceMemory); ok {
		mem = v.(float64)
	} else {
		return nil, fmt.Errorf("%s is required for creating pvm instances", PIInstanceMemory)
	}
	if v, ok := d.GetOk(PIInstanceProcessors); ok {
		procs = v.(float64)
	} else {
		return nil, fmt.Errorf("%s is required for creating pvm instances", PIInstanceProcessors)
	}
	if v, ok := d.GetOk(PIInstanceSysType); ok {
		systype = v.(string)
	} else {
		return nil, fmt.Errorf("%s is required for creating pvm instances", PIInstanceSysType)
	}
	if v, ok := d.GetOk(PIInstanceProcType); ok {
		processortype = v.(string)
	} else {
		return nil, fmt.Errorf("%s is required for creating pvm instances", PIInstanceProcType)
	}

	pvmNetworks := expandPVMNetworks(d.Get(PIInstanceNetwork).([]interface{}))

	var volids []string
	if v, ok := d.GetOk(PIInstanceVolumeIDs); ok {
		volids = flex.ExpandStringList((v.(*schema.Set)).List())
	}
	var replicants float64
	if r, ok := d.GetOk(PIInstanceReplicants); ok {
		replicants = float64(r.(int))
	}
	var replicationpolicy string
	if r, ok := d.GetOk(PIInstanceReplicationPolicy); ok {
		replicationpolicy = r.(string)
	}
	var replicationNamingScheme string
	if r, ok := d.GetOk(PIInstanceReplicationScheme); ok {
		replicationNamingScheme = r.(string)
	}
	var migratable bool
	if m, ok := d.GetOk(PIInstanceMigratable); ok {
		migratable = m.(bool)
	}

	var pinpolicy string
	if p, ok := d.GetOk(PIInstancePinPolicy); ok {
		pinpolicy = p.(string)
		if pinpolicy == "" {
			pinpolicy = "none"
		}
	}

	var userData string
	if u, ok := d.GetOk(PIInstanceUserData); ok {
		userData = u.(string)
	}
	err := checkBase64(userData)
	if err != nil {
		log.Printf("Data is not base64 encoded")
		return nil, err
	}

	//publicinterface := d.Get(helpers.PIInstancePublicNetwork).(bool)
	body := &models.PVMInstanceCreate{
		//NetworkIds: networks,
		Processors:              &procs,
		Memory:                  &mem,
		ServerName:              flex.PtrToString(name),
		SysType:                 systype,
		ImageID:                 flex.PtrToString(imageid),
		ProcType:                flex.PtrToString(processortype),
		Replicants:              replicants,
		UserData:                userData,
		ReplicantNamingScheme:   flex.PtrToString(replicationNamingScheme),
		ReplicantAffinityPolicy: flex.PtrToString(replicationpolicy),
		Networks:                pvmNetworks,
		Migratable:              &migratable,
	}
	if s, ok := d.GetOk(PIInstanceKey); ok {
		sshkey := s.(string)
		body.KeyPairName = sshkey
	}
	if len(volids) > 0 {
		body.VolumeIDs = volids
	}
	if d.Get(PIInstancePinPolicy) == "soft" || d.Get(PIInstancePinPolicy) == "hard" {
		body.PinPolicy = models.PinPolicy(pinpolicy)
	}

	var assignedVirtualCores int64
	if a, ok := d.GetOk(PIInstanceVirtualCoresAssigned); ok {
		assignedVirtualCores = int64(a.(int))
		body.VirtualCores = &models.VirtualCores{Assigned: &assignedVirtualCores}
	}

	if st, ok := d.GetOk(PIInstanceStorageType); ok {
		body.StorageType = st.(string)
	}
	if sp, ok := d.GetOk(PIInstanceStoragePool); ok {
		body.StoragePool = sp.(string)
	}

	if ap, ok := d.GetOk(PIInstanceAffinityPolicy); ok {
		policy := ap.(string)
		affinity := &models.StorageAffinity{
			AffinityPolicy: &policy,
		}

		if policy == "affinity" {
			if av, ok := d.GetOk(PIInstanceAffinityVolume); ok {
				afvol := av.(string)
				affinity.AffinityVolume = &afvol
			}
			if ai, ok := d.GetOk(PIInstanceAffinityInstance); ok {
				afins := ai.(string)
				affinity.AffinityPVMInstance = &afins
			}
		} else {
			if avs, ok := d.GetOk(PIInstanceAntiAffinityVolumes); ok {
				afvols := flex.ExpandStringList(avs.([]interface{}))
				affinity.AntiAffinityVolumes = afvols
			}
			if ais, ok := d.GetOk(PIInstanceAntiAffinityInstances); ok {
				afinss := flex.ExpandStringList(ais.([]interface{}))
				affinity.AntiAffinityPVMInstances = afinss
			}
		}
		body.StorageAffinity = affinity
	}

	if sc, ok := d.GetOk(PIInstanceStorageConnection); ok {
		body.StorageConnection = sc.(string)
	}

	if pg, ok := d.GetOk(PIInstancePlacementGroup); ok {
		body.PlacementGroup = pg.(string)
	}

	if lrc, ok := d.GetOk(PIInstanceLRC); ok {
		// check if using vtl image
		// check if vtl image is stock image
		imageData, err := imageClient.GetStockImage(imageid)
		if err != nil {
			// check if vtl image is cloud instance image
			imageData, err = imageClient.Get(imageid)
			if err != nil {
				return nil, fmt.Errorf("image doesn't exist. %e", err)
			}
		}

		if imageData.Specifications.ImageType == "stock-vtl" {
			body.LicenseRepositoryCapacity = int64(lrc.(int))
		} else {
			return nil, fmt.Errorf("pi_license_repository_capacity should only be used when creating VTL instances. %e", err)
		}
	}

	pvmList, err := client.Create(body)

	if err != nil {
		return nil, fmt.Errorf("failed to provision: %v", err)
	}
	if pvmList == nil {
		return nil, fmt.Errorf("failed to provision")
	}

	return pvmList, nil
}

func splitID(id string) (id1, id2 string, err error) {
	parts, err := flex.IdParts(id)
	if err != nil {
		return
	}
	id1 = parts[0]
	id2 = parts[1]
	return
}
