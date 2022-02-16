// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const (
	// Arguments
	PIInstanceName                  = "pi_instance_name"
	PIInstanceAffinityInstance      = "pi_affinity_instance"
	PIInstanceAffinityPolicy        = "pi_affinity_policy"
	PIInstanceAffinityVolume        = "pi_affinity_volume"
	PIInstanceAntiAffinityInstances = "pi_anti_affinity_instances"
	PIInstanceAntiAffinityVolumes   = "pi_anti_affinity_volumes"
	PIInstanceHealth                = "pi_health_status"
	PIInstanceImageID               = "pi_image_id"
	PIInstanceKey                   = "pi_key_pair_name"
	PIInstanceLRC                   = "pi_license_repository_capacity"
	PIInstanceMemory                = "pi_memory"
	PIInstanceMigratable            = "pi_migratable"
	PIInstanceNetwork               = "pi_network"
	PIInstanceNetworkID             = "network_id"
	PIInstanceNetworkIP             = "ip_address"
	PIInstancePinPolicy             = "pi_pin_policy"
	PIInstancePlacementGroup        = "pi_placement_group_id"
	PIInstanceProcessors            = "pi_processors"
	PIInstanceProcType              = "pi_proc_type"
	PIInstanceReplicants            = "pi_replicants"
	PIInstanceReplicationPolicy     = "pi_replication_policy"
	PIInstanceReplicationScheme     = "pi_replication_scheme"
	PIInstanceSapProfileID          = "pi_sap_profile_id"
	PIInstanceStoragePool           = "pi_storage_pool"
	PIInstanceStoragePoolAffinity   = "pi_storage_pool_affinity"
	PIInstanceStorageType           = "pi_storage_type"
	PIInstanceStorageConnection     = "pi_storage_connection"
	PIInstanceSysType               = "pi_sys_type"
	PIInstanceUserData              = "pi_user_data"
	PIInstanceVirtualCoresAssigned  = "pi_virtual_cores_assigned"
	PIInstanceVolumeIDs             = "pi_volume_ids"
	PIInstanceNetworkName           = "pi_network_name"

	// Attributes
	InstanceInstanceID          = "instance_id"
	InstanceHealthStatus        = "health_status"
	InstanceMemory              = "memory"
	InstanceMinProc             = "minproc"
	InstanceMaxProc             = "maxproc"
	InstanceMaxVirtualCores     = "max_virtual_cores"
	InstanceMinMem              = "minmem"
	InstanceMaxMem              = "maxmem"
	InstanceMinVirtualCores     = "min_virtual_cores"
	InstanceLRC                 = "license_repository_capacity"
	InstanceAddresses           = "addresses"
	InstanceNetworks            = "networks"
	InstanceNetworksIP          = "ip"
	InstanceNetworkExternalIP   = "external_ip"
	InstanceNetworksMAC         = "macaddress"
	InstanceNetworkID           = "network_id"
	InstanceNetworkName         = "network_name"
	InstanceNetworkType         = "type"
	InstancePlacementGroup      = "placement_group_id"
	InstanceProcessors          = "processors"
	InstanceProcType            = "proctype"
	InstanceStatus              = "status"
	InstanceStoragePool         = "storage_pool"
	InstanceStoragePoolAffinity = "storage_pool_affinity"
	InstanceStorageType         = "storage_type"
	InstanceVirtualCores        = "virtual_cores_assigned"
	InstanceVolumes             = "volumes"
	InstancePinPolicy           = "pin_policy"
	InstanceOSType              = "os_type"
	InstanceOperatingSystem     = "operating_system"
	InstanceProgress            = "progress"
	InstanceIpOctet             = "ipoctet"

	// Attributes to fix later
	InstanceMaxProcessors = "max_processors"
	InstanceMinProcessors = "min_processors"
	InstanceMinMemory     = "min_memory"
	InstanceMaxMemory     = "max_memory"
	InstanceNetwork       = "pi_network"
	InstanceNetworkIP     = "ip_address"
	InstanceNetworkMAC    = "mac_address"
	InstanceExternalIP    = "external_ip"
	Instances             = "pvm_instances"
	InstancesInstanceID   = "pvm_instance_id"
)

func DataSourceIBMPIInstance() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstancesRead,
		Schema: map[string]*schema.Schema{

			PIInstanceName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Server Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},

			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
			InstanceVolumes: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			InstanceMemory: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			InstanceProcessors: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			InstanceHealthStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceAddresses: {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: "This field is deprecated, use networks instead",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						InstanceNetworksIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworksMAC: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkExternalIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			InstanceNetworks: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						InstanceNetworksIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworksMAC: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						InstanceNetworkExternalIP: {
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
			InstanceProcType: {
				Type:     schema.TypeString,
				Computed: true,
			},

			InstanceStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			InstanceMinProc: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			InstanceMinMem: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			InstanceMaxProc: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			InstanceMaxMem: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			InstancePinPolicy: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceVirtualCores: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			InstanceMaxVirtualCores: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			InstanceMinVirtualCores: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			InstanceStorageType: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceStoragePool: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceStoragePoolAffinity: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			InstanceLRC: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			InstancePlacementGroup: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIInstancesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := meta.(conns.ClientSession).IBMPISession()

	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	powerC := instance.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
	powervmdata, err := powerC.Get(d.Get(PIInstanceName).(string))

	if err != nil {
		return diag.FromErr(err)
	}

	pvminstanceid := *powervmdata.PvmInstanceID
	d.SetId(pvminstanceid)
	d.Set(InstanceMemory, powervmdata.Memory)
	d.Set(InstanceProcessors, powervmdata.Processors)
	d.Set(InstanceStatus, powervmdata.Status)
	d.Set(InstanceProcType, powervmdata.ProcType)
	d.Set(InstanceVolumes, powervmdata.VolumeIDs)
	d.Set(InstanceMinProc, powervmdata.Minproc)
	d.Set(InstanceMinMem, powervmdata.Minmem)
	d.Set(InstanceMaxProc, powervmdata.Maxproc)
	d.Set(InstanceMaxMem, powervmdata.Maxmem)
	d.Set(InstancePinPolicy, powervmdata.PinPolicy)
	d.Set(InstanceVirtualCores, powervmdata.VirtualCores.Assigned)
	d.Set(InstanceMaxVirtualCores, powervmdata.VirtualCores.Max)
	d.Set(InstanceMinVirtualCores, powervmdata.VirtualCores.Min)
	d.Set(InstanceStorageType, powervmdata.StorageType)
	d.Set(InstanceStoragePool, powervmdata.StoragePool)
	d.Set(InstanceStoragePoolAffinity, powervmdata.StoragePoolAffinity)
	d.Set(InstanceLRC, powervmdata.LicenseRepositoryCapacity)
	d.Set(InstanceNetworks, flattenPvmInstanceNetworks(powervmdata.Networks))
	if *powervmdata.PlacementGroup != "none" {
		d.Set(InstancePlacementGroup, powervmdata.PlacementGroup)
	}

	if powervmdata.Addresses != nil {
		pvmaddress := make([]map[string]interface{}, len(powervmdata.Addresses))
		for i, pvmip := range powervmdata.Addresses {

			p := make(map[string]interface{})
			p[InstanceNetworksIP] = pvmip.IP
			p[InstanceNetworkName] = pvmip.NetworkName
			p[InstanceNetworkID] = pvmip.NetworkID
			p[InstanceNetworksMAC] = pvmip.MacAddress
			p[InstanceNetworkType] = pvmip.Type
			p[InstanceNetworkExternalIP] = pvmip.ExternalIP
			pvmaddress[i] = p
		}
		d.Set(InstanceAddresses, pvmaddress)

	}

	if powervmdata.Health != nil {

		d.Set(InstanceHealthStatus, powervmdata.Health.Status)

	}

	return nil
}
