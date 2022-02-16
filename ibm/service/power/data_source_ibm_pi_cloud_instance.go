// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const (
	// Arguments
	PICloudInstanceID = "pi_cloud_instance_id"

	// Attributes
	CloudInstanceCapabilities                 = "capabilities"
	CloudInstanceEnabled                      = "enabled"
	CloudInstancePVMInstances                 = "pvm_instances"
	CloudInstancePVMCreationDate              = "creation_date"
	CloudInstancePVMHref                      = "href"
	CloudInstancePVMID                        = "id"
	CloudInstancePVMName                      = "name"
	CloudInstancePVMStatus                    = "status"
	CloudInstancePVMSystype                   = "systype"
	CloudInstanceRegion                       = "region"
	CloudInstanceTenantID                     = "tenant_id"
	CloudInstanceTotalInstances               = "total_instances"
	CloudInstanceTotalMemoryConsumed          = "total_memory_consumed"
	CloudInstanceTotalProcessorsConsumed      = "total_processors_consumed"
	CloudInstanceTotalSSDStorageConsumed      = "total_ssd_storage_consumed"
	CloudInstanceTotalStandardStorageConsumed = "total_standard_storage_consumed"
)

func DataSourceIBMPICloudInstance() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPICloudInstanceRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Start of Computed Attributes
			CloudInstanceEnabled: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			CloudInstanceTenantID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			CloudInstanceRegion: {
				Type:     schema.TypeString,
				Computed: true,
			},
			CloudInstanceCapabilities: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			CloudInstanceTotalProcessorsConsumed: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			CloudInstanceTotalInstances: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			CloudInstanceTotalMemoryConsumed: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			CloudInstanceTotalSSDStorageConsumed: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			CloudInstanceTotalStandardStorageConsumed: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			CloudInstancePVMInstances: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						CloudInstancePVMID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudInstancePVMName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudInstancePVMHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudInstancePVMStatus: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudInstancePVMSystype: {
							Type:     schema.TypeString,
							Computed: true,
						},
						CloudInstancePVMCreationDate: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPICloudInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	cloud_instance := instance.NewIBMPICloudInstanceClient(ctx, sess, cloudInstanceID)
	cloud_instance_data, err := cloud_instance.Get(cloudInstanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*cloud_instance_data.CloudInstanceID)
	d.Set(CloudInstanceTenantID, (cloud_instance_data.TenantID))
	d.Set(CloudInstanceEnabled, cloud_instance_data.Enabled)
	d.Set(CloudInstanceRegion, cloud_instance_data.Region)
	d.Set(CloudInstanceCapabilities, cloud_instance_data.Capabilities)
	d.Set(CloudInstancePVMInstances, flattenpvminstances(cloud_instance_data.PvmInstances))
	d.Set(CloudInstanceTotalSSDStorageConsumed, cloud_instance_data.Usage.StorageSSD)
	d.Set(CloudInstanceTotalInstances, cloud_instance_data.Usage.Instances)
	d.Set(CloudInstanceTotalStandardStorageConsumed, cloud_instance_data.Usage.StorageStandard)
	d.Set(CloudInstanceTotalProcessorsConsumed, cloud_instance_data.Usage.Processors)
	d.Set(CloudInstanceTotalMemoryConsumed, cloud_instance_data.Usage.Memory)

	return nil

}

func flattenpvminstances(list []*models.PVMInstanceReference) []map[string]interface{} {
	pvms := make([]map[string]interface{}, 0)
	for _, lpars := range list {

		l := map[string]interface{}{
			CloudInstancePVMID:           *lpars.PvmInstanceID,
			CloudInstancePVMName:         *lpars.ServerName,
			CloudInstancePVMHref:         *lpars.Href,
			CloudInstancePVMStatus:       *lpars.Status,
			CloudInstancePVMSystype:      lpars.SysType,
			CloudInstancePVMCreationDate: lpars.CreationDate.String(),
		}
		pvms = append(pvms, l)

	}
	return pvms
}
