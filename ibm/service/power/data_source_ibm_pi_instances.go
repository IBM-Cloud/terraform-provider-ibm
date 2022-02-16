// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

// Attributes and Arguments defined in data_source_ibm_pi_instance.go
func DataSourceIBMPIInstances() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstancesAllRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
			Instances: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						InstancesInstanceID: {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceIBMPIInstancesAllRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := meta.(conns.ClientSession).IBMPISession()

	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	powerC := instance.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)
	powervmdata, err := powerC.GetAll()

	if err != nil {
		return diag.FromErr(err)
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set(Instances, flattenPvmInstances(powervmdata.PvmInstances))

	return nil
}

func flattenPvmInstances(list []*models.PVMInstanceReference) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {

		l := map[string]interface{}{
			InstancesInstanceID:         *i.PvmInstanceID,
			InstanceMemory:              *i.Memory,
			InstanceProcessors:          *i.Processors,
			InstanceProcType:            *i.ProcType,
			InstanceStatus:              *i.Status,
			InstanceMinProc:             i.Minproc,
			InstanceMinMem:              i.Minmem,
			InstanceMaxProc:             i.Maxproc,
			InstanceMaxMem:              i.Maxmem,
			InstancePinPolicy:           i.PinPolicy,
			InstanceVirtualCores:        i.VirtualCores.Assigned,
			InstanceMaxVirtualCores:     i.VirtualCores.Max,
			InstanceMinVirtualCores:     i.VirtualCores.Min,
			InstanceStorageType:         i.StorageType,
			InstanceStoragePool:         i.StoragePool,
			InstanceStoragePoolAffinity: i.StoragePoolAffinity,
			InstanceLRC:                 i.LicenseRepositoryCapacity,
			InstancePlacementGroup:      i.PlacementGroup,
			InstanceNetworks:            flattenPvmInstanceNetworks(i.Networks),
		}

		if i.Health != nil {
			l[InstanceHealthStatus] = i.Health.Status
		}

		result = append(result, l)

	}
	return result
}

func flattenPvmInstanceNetworks(list []*models.PVMInstanceNetwork) (networks []map[string]interface{}) {
	if list != nil {
		networks = make([]map[string]interface{}, len(list))
		for i, pvmip := range list {

			p := make(map[string]interface{})
			p[InstanceNetworksIP] = pvmip.IP
			p[InstanceNetworkName] = pvmip.NetworkName
			p[InstanceNetworkID] = pvmip.NetworkID
			p[InstanceNetworksMAC] = pvmip.MacAddress
			p[InstanceNetworkType] = pvmip.Type
			p[InstanceNetworkExternalIP] = pvmip.ExternalIP
			networks[i] = p
		}
		return networks
	}
	return
}
