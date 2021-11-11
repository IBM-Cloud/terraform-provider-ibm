// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

const (
	createTimeOut = 120 * time.Second
	updateTimeOut = 120 * time.Second
	postTimeOut   = 60 * time.Second
	getTimeOut    = 60 * time.Second
	deleteTimeOut = 60 * time.Second
	//Added timeout values for warning  and active status
	warningTimeOut = 30 * time.Second
	activeTimeOut  = 2 * time.Minute
	// power service instance capabilities
	CUSTOM_VIRTUAL_CORES  = "custom-virtualcores"
	PIInstanceNetwork     = "pi_network"
	PIInstanceStoragePool = "pi_storage_pool"
)

func resourceIBMPIInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPIInstanceCreate,
		Read:     resourceIBMPIInstanceRead,
		Update:   resourceIBMPIInstanceUpdate,
		Delete:   resourceIBMPIInstanceDelete,
		Exists:   resourceIBMPIInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(120 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the Power Instance id that is assigned to the account",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PI instance status",
			},
			"pi_migratable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "set to true to enable migration of the PI instance",
			},
			"migratable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "set to true to enable migration of the PI instance",
				Deprecated:  "This field is deprecated, Use pi_migratable instead.",
			},
			"min_processors": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Minimum number of the CPUs",
			},
			"min_memory": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Minimum memory",
			},
			"max_processors": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Maximum number of processors",
			},
			"max_memory": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Maximum memory size",
			},
			helpers.PIInstanceNetworkIds: {
				Type:             schema.TypeList,
				Optional:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Description:      "List of Networks that have been configured for the account",
				DiffSuppressFunc: applyOnce,
				Deprecated:       "Use pi_network instead",
				ConflictsWith:    []string{PIInstanceNetwork},
			},

			helpers.PIInstanceVolumeIds: {
				Type:             schema.TypeSet,
				Optional:         true,
				Computed:         true,
				Elem:             &schema.Schema{Type: schema.TypeString},
				Set:              schema.HashString,
				DiffSuppressFunc: applyOnce,
				Description:      "List of PI volumes",
			},

			helpers.PIInstanceUserData: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base64 encoded data to be passed in for invoking a cloud init script",
			},

			helpers.PIInstanceStorageType: {
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
			PIAffinityPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Affinity policy for pvm instance being created; ignored if pi_storage_pool provided; for policy affinity requires one of pi_affinity_instance or pi_affinity_volume to be specified; for policy anti-affinity requires one of pi_anti_affinity_instances or pi_anti_affinity_volumes to be specified",
				ValidateFunc: validateAllowedStringValue([]string{"affinity", "anti-affinity"}),
			},
			PIAffinityVolume: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Volume (ID or Name) to base storage affinity policy against; required if requesting affinity and pi_affinity_instance is not provided",
				ConflictsWith: []string{PIAffinityInstance},
			},
			PIAffinityInstance: {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "PVM Instance (ID or Name) to base storage affinity policy against; required if requesting storage affinity and pi_affinity_volume is not provided",
				ConflictsWith: []string{PIAffinityVolume},
			},
			PIAntiAffinityVolumes: {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of volumes to base storage anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_instances is not provided",
				ConflictsWith: []string{PIAntiAffinityInstances},
			},
			PIAntiAffinityInstances: {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   "List of pvmInstances to base storage anti-affinity policy against; required if requesting anti-affinity and pi_anti_affinity_volumes is not provided",
				ConflictsWith: []string{PIAntiAffinityVolumes},
			},

			helpers.PIInstanceStorageConnection: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"vSCSI"}),
				Description:  "Storage Connectivity Group for server deployment",
			},
			PIInstanceNetwork: {
				Type: schema.TypeList,
				// TODO: Once pi_network_ids is removed this will be a required field
				Optional:         true,
				ConflictsWith:    []string{helpers.PIInstanceNetworkIds},
				Computed:         true,
				DiffSuppressFunc: applyOnce,
				Description:      "List of one or more networks to attach to the instance",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"network_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"addresses": {
				Deprecated: "Use pi_network instead",
				Type:       schema.TypeList,
				Computed:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"macaddress": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						/*"version": {
							Type:     schema.TypeFloat,
							Computed: true,
						},*/
					},
				},
			},

			"health_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PI Instance health status",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance ID",
			},
			"pin_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PIN Policy of the Instance",
			},
			helpers.PIInstanceImageId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI instance image id",
			},
			helpers.PIInstanceProcessors: {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Processors count",
			},
			helpers.PIInstanceName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI Instance name",
			},
			helpers.PIInstanceProcType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"dedicated", "shared", "capped"}),
				Description:  "Instance processor type",
			},
			helpers.PIInstanceSSHKeyName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SSH key name",
			},
			helpers.PIInstanceMemory: {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Memory size",
			},
			helpers.PIInstanceSystemType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"s922", "e880", "e980"}),
				Description:  "PI Instance system type",
			},
			helpers.PIInstanceReplicants: {
				Type:        schema.TypeFloat,
				Optional:    true,
				Default:     1.0,
				Description: "PI Instance replicas count",
			},
			helpers.PIInstanceReplicationPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"affinity", "anti-affinity", "none"}),
				Default:      "none",
				Description:  "Replication policy for the PI Instance",
			},
			helpers.PIInstanceReplicationScheme: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"prefix", "suffix"}),
				Default:      "suffix",
				Description:  "Replication scheme",
			},
			helpers.PIInstanceProgress: {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Progress of the operation",
			},
			helpers.PIInstancePinPolicy: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Pin Policy of the instance",
				Default:      "none",
				ValidateFunc: validateAllowedStringValue([]string{"none", "soft", "hard"}),
			},

			// "reboot_for_resource_change": {
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// 	Description: "Flag to be passed for CPU/Memory changes that require a reboot to take effect",
			// },
			"operating_system": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operating System",
			},
			"os_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OS Type",
			},
			helpers.PIInstanceHealthStatus: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"OK", "WARNING"}),
				Default:      "OK",
				Description:  "Allow the user to set the status of the lpar so that they can connect to it faster",
			},
			helpers.PIVirtualCoresAssigned: {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Virtual Cores Assigned to the PVMInstance",
			},
			"max_virtual_cores": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum Virtual Cores Assigned to the PVMInstance",
			},
			"min_virtual_cores": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum Virtual Cores Assigned to the PVMInstance",
			},
		},
	}
}

func resourceIBMPIInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Now in the PowerVMCreate")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	name := d.Get(helpers.PIInstanceName).(string)
	sshkey := d.Get(helpers.PIInstanceSSHKeyName).(string)
	mem := d.Get(helpers.PIInstanceMemory).(float64)
	procs := d.Get(helpers.PIInstanceProcessors).(float64)
	systype := d.Get(helpers.PIInstanceSystemType).(string)

	var pvmNetworks []*models.PVMInstanceAddNetwork
	// Either pi_network_ids or pi_network is provided
	// TODO: Once pi_network_ids is removed pi_network will be a required field
	if v, ok := d.GetOk(helpers.PIInstanceNetworkIds); ok {
		networks := expandStringList(v.([]interface{}))
		pvmNetworks = buildPVMNetworks(networks)
	}
	if v, ok := d.GetOk(PIInstanceNetwork); ok {
		networksMap := v.([]interface{})
		pvmNetworks = expandPVMNetworks(networksMap)
	}
	if len(pvmNetworks) <= 0 {
		return fmt.Errorf("one of pi_network_ids or pi_network must be configured")
	}

	var volids []string
	if v, ok := d.GetOk(helpers.PIInstanceVolumeIds); ok {
		volids = expandStringList((v.(*schema.Set)).List())
	}
	var replicants float64
	if r, ok := d.GetOk(helpers.PIInstanceReplicants); ok {
		replicants = r.(float64)
	}
	var replicationpolicy string
	if r, ok := d.GetOk(helpers.PIInstanceReplicationPolicy); ok {
		replicationpolicy = r.(string)
	}
	var replicationNamingScheme string
	if r, ok := d.GetOk(helpers.PIInstanceReplicationScheme); ok {
		replicationNamingScheme = r.(string)
	}
	var migratable bool
	if m, ok := d.GetOk("pi_migratable"); ok {
		migratable = m.(bool)
	}
	imageid := d.Get(helpers.PIInstanceImageId).(string)
	processortype := d.Get(helpers.PIInstanceProcType).(string)

	var pinpolicy string
	if p, ok := d.GetOk(helpers.PIInstancePinPolicy); ok {
		pinpolicy = p.(string)
		if pinpolicy == "" {
			pinpolicy = "none"
		}
	}
	var instanceReadyStatus string
	if r, ok := d.GetOk(helpers.PIInstanceHealthStatus); ok {
		instanceReadyStatus = r.(string)
	}

	var userData string
	if u, ok := d.GetOk(helpers.PIInstanceUserData); ok {
		userData = u.(string)
	}
	err = checkBase64(userData)
	if err != nil {
		log.Printf("Data is not base64 encoded")
		return err
	}

	//publicinterface := d.Get(helpers.PIInstancePublicNetwork).(bool)
	body := &models.PVMInstanceCreate{
		//NetworkIds: networks,
		Processors:              &procs,
		Memory:                  &mem,
		ServerName:              ptrToString(name),
		SysType:                 systype,
		KeyPairName:             sshkey,
		ImageID:                 ptrToString(imageid),
		ProcType:                ptrToString(processortype),
		Replicants:              replicants,
		UserData:                userData,
		ReplicantNamingScheme:   ptrToString(replicationNamingScheme),
		ReplicantAffinityPolicy: ptrToString(replicationpolicy),
		Networks:                pvmNetworks,
		Migratable:              &migratable,
	}
	if len(volids) > 0 {
		body.VolumeIds = volids
	}
	if d.Get(helpers.PIInstancePinPolicy) == "soft" || d.Get(helpers.PIInstancePinPolicy) == "hard" {
		body.PinPolicy = models.PinPolicy(pinpolicy)
	}

	var assignedVirtualCores int64
	if a, ok := d.GetOk(helpers.PIVirtualCoresAssigned); ok {
		assignedVirtualCores = int64(a.(int))
		body.VirtualCores = &models.VirtualCores{Assigned: &assignedVirtualCores}
	}

	if st, ok := d.GetOk(helpers.PIInstanceStorageType); ok {
		body.StorageType = st.(string)
	}
	if sp, ok := d.GetOk(PIInstanceStoragePool); ok {
		body.StoragePool = sp.(string)
	}

	if ap, ok := d.GetOk(PIAffinityPolicy); ok {
		policy := ap.(string)
		affinity := &models.StorageAffinity{
			AffinityPolicy: &policy,
		}

		if policy == "affinity" {
			if av, ok := d.GetOk(PIAffinityVolume); ok {
				afvol := av.(string)
				affinity.AffinityVolume = &afvol
			}
			if ai, ok := d.GetOk(PIAffinityInstance); ok {
				afins := ai.(string)
				affinity.AffinityPVMInstance = &afins
			}
		} else {
			if avs, ok := d.GetOk(PIAntiAffinityVolumes); ok {
				afvols := expandStringList(avs.([]interface{}))
				affinity.AntiAffinityVolumes = afvols
			}
			if ais, ok := d.GetOk(PIAntiAffinityInstances); ok {
				afinss := expandStringList(ais.([]interface{}))
				affinity.AntiAffinityPVMInstances = afinss
			}
		}
		body.StorageAffinity = affinity
	}

	if sc, ok := d.GetOk(helpers.PIInstanceStorageConnection); ok {
		body.StorageConnection = sc.(string)
	}

	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	pvm, err := client.Create(&p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
		Body: body,
	}, powerinstanceid, createTimeOut)

	if err != nil {
		return fmt.Errorf("failed to provision %s", err)
	}

	var pvminstanceids []string
	if replicants > 1 {
		log.Printf("We are in a multi create mode")
		for i := 0; i < int(replicants); i++ {
			truepvmid := (*pvm)[i].PvmInstanceID
			pvminstanceids = append(pvminstanceids, fmt.Sprintf("%s", *truepvmid))
			d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, *truepvmid))
		}
		d.SetId(strings.Join(pvminstanceids, "/"))
	} else {
		truepvmid := (*pvm)[0].PvmInstanceID
		d.SetId(fmt.Sprintf("%s/%s", powerinstanceid, *truepvmid))
		pvminstanceids = append(pvminstanceids, *truepvmid)
	}

	for ids := range pvminstanceids {
		_, err = isWaitForPIInstanceAvailable(client, pvminstanceids[ids], d.Timeout(schema.TimeoutCreate), powerinstanceid, instanceReadyStatus)
		if err != nil {
			return err
		}
	}

	return resourceIBMPIInstanceRead(d, meta)

}

func resourceIBMPIInstanceRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	powerC := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	powervmdata, err := powerC.Get(parts[1], powerinstanceid, getTimeOut)
	if err != nil {
		return fmt.Errorf("failed to get the instance %v", err)
	}

	d.Set(helpers.PIInstanceMemory, powervmdata.Memory)
	d.Set(helpers.PIInstanceProcessors, powervmdata.Processors)
	if powervmdata.Status != nil {
		d.Set("status", powervmdata.Status)
	}
	d.Set(helpers.PIInstanceProcType, powervmdata.ProcType)
	if powervmdata.Migratable != nil {
		d.Set("pi_migratable", powervmdata.Migratable)
	}
	d.Set("min_processors", powervmdata.Minproc)
	d.Set(helpers.PIInstanceProgress, powervmdata.Progress)
	if powervmdata.StorageType != nil {
		d.Set(helpers.PIInstanceStorageType, powervmdata.StorageType)
	}
	d.Set(PIInstanceStoragePool, powervmdata.StoragePool)
	d.Set(helpers.PICloudInstanceId, powerinstanceid)
	d.Set("instance_id", powervmdata.PvmInstanceID)
	d.Set(helpers.PIInstanceName, powervmdata.ServerName)
	d.Set(helpers.PIInstanceImageId, powervmdata.ImageID)

	if _, ok := d.GetOk(helpers.PIInstanceNetworkIds); ok {
		var networks []string
		networks = make([]string, 0)
		if powervmdata.Networks != nil {
			for _, n := range powervmdata.Networks {
				if n != nil {
					networks = append(networks, n.NetworkID)
				}
			}
		}
		d.Set(helpers.PIInstanceNetworkIds, networks)
	} else {
		networksMap := []map[string]interface{}{}
		if powervmdata.Networks != nil {
			for _, n := range powervmdata.Networks {
				if n != nil {
					v := map[string]interface{}{
						"ip_address":   n.IPAddress,
						"mac_address":  n.MacAddress,
						"network_id":   n.NetworkID,
						"network_name": n.NetworkName,
						"type":         n.Type,
						"external_ip":  n.ExternalIP,
					}
					networksMap = append(networksMap, v)
				}
			}
		}
		d.Set(PIInstanceNetwork, networksMap)
	}

	if powervmdata.VolumeIds != nil {
		d.Set(helpers.PIInstanceVolumeIds, powervmdata.VolumeIds)
	}
	d.Set(helpers.PIInstanceSystemType, powervmdata.SysType)
	d.Set("min_memory", powervmdata.Minmem)
	d.Set("max_processors", powervmdata.Maxproc)
	d.Set("max_memory", powervmdata.Maxmem)
	d.Set("pin_policy", powervmdata.PinPolicy)
	d.Set("operating_system", powervmdata.OperatingSystem)
	if powervmdata.OsType != nil {
		d.Set("os_type", powervmdata.OsType)
	}

	if powervmdata.Addresses != nil {
		pvmaddress := make([]map[string]interface{}, len(powervmdata.Addresses))
		for i, pvmip := range powervmdata.Addresses {
			log.Printf("Now entering the powervm address space....")

			p := make(map[string]interface{})
			p["ip"] = pvmip.IP
			p["network_name"] = pvmip.NetworkName
			p["network_id"] = pvmip.NetworkID
			p["macaddress"] = pvmip.MacAddress
			p["type"] = pvmip.Type
			p["external_ip"] = pvmip.ExternalIP
			pvmaddress[i] = p
		}
		d.Set("addresses", pvmaddress)
	}

	if powervmdata.Health != nil {
		d.Set("health_status", powervmdata.Health.Status)
	}
	if powervmdata.VirtualCores.Assigned != nil {
		d.Set(helpers.PIVirtualCoresAssigned, powervmdata.VirtualCores.Assigned)
	}
	if &powervmdata.VirtualCores.Max != nil {
		d.Set("max_virtual_cores", powervmdata.VirtualCores.Max)
	}
	if &powervmdata.VirtualCores.Min != nil {
		d.Set("min_virtual_cores", powervmdata.VirtualCores.Min)
	}

	return nil

}

func resourceIBMPIInstanceUpdate(d *schema.ResourceData, meta interface{}) error {

	name := d.Get(helpers.PIInstanceName).(string)
	mem := d.Get(helpers.PIInstanceMemory).(float64)
	procs := d.Get(helpers.PIInstanceProcessors).(float64)
	processortype := d.Get(helpers.PIInstanceProcType).(string)
	assignedVirtualCores := int64(d.Get(helpers.PIVirtualCoresAssigned).(int))

	if d.Get("health_status") == "WARNING" {
		return fmt.Errorf("the operation cannot be performed when the lpar health in the WARNING State")
	}

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return fmt.Errorf("failed to get the session from the IBM Cloud Service")
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	// Check if cloud instance is capable of changing virtual cores
	cloudInstanceClient := st.NewIBMPICloudInstanceClient(sess, powerinstanceid)
	cloudInstance, err := cloudInstanceClient.Get(powerinstanceid)
	if err != nil {
		return fmt.Errorf("failed to get cloud instance %v", err)
	}
	cores_enabled := checkCloudInstanceCapability(cloudInstance, CUSTOM_VIRTUAL_CORES)

	if d.HasChange(helpers.PIInstanceName) {
		body := &models.PVMInstanceUpdate{
			ServerName: name,
		}
		_, err = client.Update(parts[1], powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body}, updateTimeOut)
		if err != nil {
			return fmt.Errorf("failed to update the lpar with the change for name %s", err)
		}
		_, err = isWaitForPIInstanceAvailable(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid, "OK")
		if err != nil {
			return err
		}
	}

	if d.HasChange(helpers.PIInstanceProcType) {

		// Stop the lpar
		if d.Get("status") == "SHUTOFF" {
			log.Printf("the lpar is in the shutoff state. Nothing to do . Moving on ")
		} else {
			err := stopLparForResourceChange(client, parts[1], powerinstanceid, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return err
			}
		}

		// Modify
		log.Printf("At this point the lpar should be off. Executing the Processor Update Change")
		updatebody := &models.PVMInstanceUpdate{ProcType: processortype}
		if cores_enabled == true {
			log.Printf("support for %s is enabled", CUSTOM_VIRTUAL_CORES)
			updatebody.VirtualCores = &models.VirtualCores{Assigned: &assignedVirtualCores}
		} else {
			log.Printf("no virtual cores support enabled for this customer..")
		}
		_, err = client.Update(parts[1], powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: updatebody}, updateTimeOut)
		if err != nil {
			return fmt.Errorf("failed to perform the modify operation on the pvm instance %v", err)
		}
		_, err = isWaitForPIInstanceStopped(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid)
		if err != nil {
			return err
		}

		// Start the lpar
		err := startLparAfterResourceChange(client, parts[1], powerinstanceid, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	// Virtual core will be updated only if service instance capability is enabled
	if d.HasChange(helpers.PIVirtualCoresAssigned) {
		body := &models.PVMInstanceUpdate{
			VirtualCores: &models.VirtualCores{Assigned: &assignedVirtualCores},
		}
		_, err = client.Update(parts[1], powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body}, updateTimeOut)
		if err != nil {
			return fmt.Errorf("failed to update the lpar with the change for virtual cores %s", err)
		}
		_, err = isWaitForPIInstanceAvailable(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid, "OK")
		if err != nil {
			return err
		}
	}

	// Start of the change for Memory and Processors
	if d.HasChange(helpers.PIInstanceMemory) || d.HasChange(helpers.PIInstanceProcessors) || d.HasChange("pi_migratable") {

		maxMemLpar := d.Get("max_memory").(float64)
		maxCPULpar := d.Get("max_processors").(float64)
		//log.Printf("the required memory is set to [%d] and current max memory is set to  [%d] ", int(mem), int(maxMemLpar))

		if mem > maxMemLpar || procs > maxCPULpar {
			log.Printf("Will require a shutdown to perform the change")
		} else {
			log.Printf("maxMemLpar is set to %f", maxMemLpar)
			log.Printf("maxCPULpar is set to %f", maxCPULpar)
		}

		//if d.GetOkExists("reboot_for_resource_change")

		if mem > maxMemLpar || procs > maxCPULpar {

			err = performChangeAndReboot(client, parts[1], powerinstanceid, mem, procs)
			if err != nil {
				return err
			}

		} else {

			body := &models.PVMInstanceUpdate{
				Memory:     mem,
				Processors: procs,
			}
			if m, ok := d.GetOk("pi_migratable"); ok {
				migratable := m.(bool)
				body.Migratable = &migratable
			}
			if cores_enabled {
				log.Printf("support for %s is enabled", CUSTOM_VIRTUAL_CORES)
				body.VirtualCores = &models.VirtualCores{Assigned: &assignedVirtualCores}
			} else {
				log.Printf("no virtual cores support enabled for this customer..")
			}

			_, err = client.Update(parts[1], powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body}, updateTimeOut)
			if err != nil {
				return fmt.Errorf("failed to update the lpar with the change %v", err)
			}
			_, err = isWaitForPIInstanceAvailable(client, parts[1], d.Timeout(schema.TimeoutUpdate), powerinstanceid, "OK")
			if err != nil {
				return err
			}
		}
	}

	return resourceIBMPIInstanceRead(d, meta)

}

func resourceIBMPIInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).IBMPISession()
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)
	err = client.Delete(parts[1], powerinstanceid, deleteTimeOut)
	if err != nil {
		return fmt.Errorf("failed to perform the delete action on the pvm instance %s", err)
	}

	_, err = isWaitForPIInstanceDeleted(client, parts[1], d.Timeout(schema.TimeoutDelete), powerinstanceid)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

// Exists

func resourceIBMPIInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	log.Printf("Calling the PowerInstance Exists method")
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	powerinstanceid := parts[0]
	client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

	instance, err := client.Get(parts[1], powerinstanceid, getTimeOut)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("error communicating with the API: %s", err)
	}

	truepvmid := *instance.PvmInstanceID
	return truepvmid == parts[1], nil
}

func isWaitForPIInstanceDeleted(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {

	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", helpers.PIInstanceDeleting},
		Target:     []string{helpers.PIInstanceNotFound},
		Refresh:    isPIInstanceDeleteRefreshFunc(client, id, powerinstanceid),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
		Timeout:    10 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isPIInstanceDeleteRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pvm, err := client.Get(id, powerinstanceid, getTimeOut)
		if err != nil {
			log.Printf("The power vm does not exist")
			return pvm, helpers.PIInstanceNotFound, nil

		}
		return pvm, helpers.PIInstanceDeleting, nil

	}
}

func isWaitForPIInstanceAvailable(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string, instanceReadyStatus string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be available and active ", id)

	var queryTimeOut time.Duration

	if instanceReadyStatus == "WARNING" {
		queryTimeOut = warningTimeOut
	} else {
		queryTimeOut = activeTimeOut
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING", helpers.PIInstanceBuilding, helpers.PIInstanceHealthWarning},
		Target:     []string{helpers.PIInstanceAvailable, helpers.PIInstanceHealthOk, "ERROR", ""},
		Refresh:    isPIInstanceRefreshFunc(client, id, powerinstanceid, instanceReadyStatus),
		Delay:      10 * time.Second,
		MinTimeout: queryTimeOut,
		Timeout:    120 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isPIInstanceRefreshFunc(client *st.IBMPIInstanceClient, id, powerinstanceid, instanceReadyStatus string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id, powerinstanceid, getTimeOut)
		if err != nil {
			return nil, "", err
		}
		allowableStatus := instanceReadyStatus
		if *pvm.Status == helpers.PIInstanceAvailable && (pvm.Health.Status == allowableStatus) {
			return pvm, helpers.PIInstanceAvailable, nil
		}
		if *pvm.Status == "ERROR" {
			return pvm, *pvm.Status, fmt.Errorf("Failed to create the lpar")
		}

		return pvm, helpers.PIInstanceBuilding, nil
	}
}

func checkBase64(input string) error {
	_, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return fmt.Errorf("Failed to check if input is base64 %s", err)
	}
	return err

}

func isWaitForPIInstanceStopped(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be stopped and powered off ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"STOPPING", "RESIZE", "VERIFY_RESIZE", helpers.PIInstanceHealthWarning},
		Target:     []string{"OK", "SHUTOFF"},
		Refresh:    isPIInstanceRefreshFuncOff(client, id, powerinstanceid),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute, // This is the time that the client will execute to check the status of the request
		Timeout:    30 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isPIInstanceRefreshFuncOff(client *st.IBMPIInstanceClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		log.Printf("Calling the check Refresh status of the pvm [%s] for cloud instance id [%s ]", id, powerinstanceid)
		pvm, err := client.Get(id, powerinstanceid, getTimeOut)
		if err != nil {
			return nil, "", err
		}
		if *pvm.Status == "SHUTOFF" && pvm.Health.Status == helpers.PIInstanceHealthOk {
			return pvm, "SHUTOFF", nil
		}
		return pvm, "STOPPING", nil
	}
}

func stopLparForResourceChange(client *st.IBMPIInstanceClient, id, powerinstanceid string, stopUpdateTimeOut time.Duration) error {
	body := &models.PVMInstanceAction{
		//Action: ptrToString("stop"),
		Action: ptrToString("immediate-shutdown"),
	}
	_, err := client.Action(&p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams{Body: body}, id, powerinstanceid, postTimeOut)
	if err != nil {
		return fmt.Errorf("failed to perform the stop action on the pvm instance %v", err)
	}

	_, err = isWaitForPIInstanceStopped(client, id, stopUpdateTimeOut, powerinstanceid)

	return err
}

// Start the lpar

func startLparAfterResourceChange(client *st.IBMPIInstanceClient, id, powerinstanceid string, startUpdateTimeOut time.Duration) error {
	body := &models.PVMInstanceAction{
		Action: ptrToString("start"),
	}
	_, err := client.Action(&p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams{Body: body}, id, powerinstanceid, postTimeOut)
	if err != nil {
		return fmt.Errorf("failed to perform the start action on the pvm instance %v", err)
	}

	_, err = isWaitForPIInstanceAvailable(client, id, startUpdateTimeOut, powerinstanceid, "OK")

	return err
}

// Stop / Modify / Start only when the lpar is off limits

func performChangeAndReboot(client *st.IBMPIInstanceClient, id, powerinstanceid string, mem, procs float64) error {
	/*
		These are the steps
		1. Stop the lpar - Check if the lpar is SHUTOFF
		2. Once the lpar is SHUTOFF - Make the cpu / memory change - DUring this time , you can check for RESIZE and VERIFY_RESIZE as the transition states
		3. If the change is successful , the lpar state will be back in SHUTOFF
		4. Once the LPAR state is SHUTOFF , initiate the start again and check for ACTIVE + OK
	*/
	//Execute the stop

	log.Printf("Callin the stop lpar for Resource Change code ..")
	err := stopLparForResourceChange(client, id, powerinstanceid, 30)
	if err != nil {
		return err
	}

	body := &models.PVMInstanceUpdate{
		Memory:     mem,
		Processors: procs,
	}

	_, updateErr := client.Update(id, powerinstanceid, &p_cloud_p_vm_instances.PcloudPvminstancesPutParams{Body: body}, updateTimeOut)
	if updateErr != nil {
		return fmt.Errorf("failed to update the lpar with the change, %s", updateErr)
	}

	_, err = isWaitforPIInstanceUpdate(client, id, 30, powerinstanceid)
	if err != nil {
		return fmt.Errorf("failed to get an update from the Service after the resource change, %s", err)
	}

	// Now we can start the lpar
	log.Printf("Calling the start lpar After the  Resource Change code ..")
	err = startLparAfterResourceChange(client, id, powerinstanceid, 30)
	if err != nil {
		return err
	}

	return nil

}

func isWaitforPIInstanceUpdate(client *st.IBMPIInstanceClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {
	log.Printf("Waiting for PIInstance (%s) to be SHUTOFF AFTER THE RESIZE Due to DLPAR Operation ", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"RESIZE", "VERIFY_RESIZE"},
		Target:     []string{"ACTIVE", "SHUTOFF", helpers.PIInstanceHealthOk},
		Refresh:    isPIInstanceShutAfterResourceChange(client, id, powerinstanceid),
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Minute,
		Timeout:    60 * time.Minute,
	}

	return stateConf.WaitForState()
}

func isPIInstanceShutAfterResourceChange(client *st.IBMPIInstanceClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		pvm, err := client.Get(id, powerinstanceid, getTimeOut)
		if err != nil {
			return nil, "", err
		}

		if *pvm.Status == "SHUTOFF" && pvm.Health.Status == helpers.PIInstanceHealthOk {
			log.Printf("The lpar is now off after the resource change...")
			return pvm, "SHUTOFF", nil
		}

		return pvm, "RESIZE", nil
	}
}

func buildPVMNetworks(networks []string) []*models.PVMInstanceAddNetwork {
	var pvmNetworks []*models.PVMInstanceAddNetwork

	for i := 0; i < len(networks); i++ {
		pvmInstanceNetwork := &models.PVMInstanceAddNetwork{
			//TODO : Enable the functionality to pass in ip address for the network
			IPAddress: "",
			NetworkID: ptrToString(string(networks[i])),
		}
		pvmNetworks = append(pvmNetworks, pvmInstanceNetwork)

	}
	return pvmNetworks
}

func expandPVMNetworks(networks []interface{}) []*models.PVMInstanceAddNetwork {
	pvmNetworks := make([]*models.PVMInstanceAddNetwork, 0, len(networks))
	for _, v := range networks {
		network := v.(map[string]interface{})
		pvmInstanceNetwork := &models.PVMInstanceAddNetwork{
			IPAddress: network["ip_address"].(string),
			NetworkID: ptrToString(network["network_id"].(string)),
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
