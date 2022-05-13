// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIBMPICloudInstance() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPICloudInstanceRead,
		Schema: map[string]*schema.Schema{
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Start of Computed Attributes
			Attr_CloudInstanceEnabled: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			Attr_CloudInstanceTenant: {
				Type:     schema.TypeString,
				Computed: true,
			},
			Attr_CloudInstanceRegion: {
				Type:     schema.TypeString,
				Computed: true,
			},
			Attr_CloudInstanceCapabilities: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			Attr_CloudInstanceTotalProcessors: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			Attr_CloudInstanceTotalInstances: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			Attr_CloudInstanceTotalMemory: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			Attr_CloudInstanceTotalSSD: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			Attr_CloudInstanceTotalStorage: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			Attr_CloudInstanceInstances: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_InstanceID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_InstanceName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_InstanceHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_InstanceStatus: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_InstanceSysType: {
							Type:     schema.TypeString,
							Computed: true,
						},
						Attr_InstanceCreationDate: {
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

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)

	cloud_instance := instance.NewIBMPICloudInstanceClient(ctx, sess, cloudInstanceID)
	cloud_instance_data, err := cloud_instance.Get(cloudInstanceID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*cloud_instance_data.CloudInstanceID)
	d.Set(Attr_CloudInstanceTenant, (cloud_instance_data.TenantID))
	d.Set(Attr_CloudInstanceEnabled, cloud_instance_data.Enabled)
	d.Set(Attr_CloudInstanceRegion, cloud_instance_data.Region)
	d.Set(Attr_CloudInstanceCapabilities, cloud_instance_data.Capabilities)
	d.Set(Attr_CloudInstanceInstances, flattenpvminstances(cloud_instance_data.PvmInstances))
	d.Set(Attr_CloudInstanceTotalSSD, cloud_instance_data.Usage.StorageSSD)
	d.Set(Attr_CloudInstanceTotalInstances, cloud_instance_data.Usage.Instances)
	d.Set(Attr_CloudInstanceTotalStorage, cloud_instance_data.Usage.StorageStandard)
	d.Set(Attr_CloudInstanceTotalProcessors, cloud_instance_data.Usage.Processors)
	d.Set(Attr_CloudInstanceTotalMemory, cloud_instance_data.Usage.Memory)

	return nil

}

func flattenpvminstances(list []*models.PVMInstanceReference) []map[string]interface{} {
	pvms := make([]map[string]interface{}, 0)
	for _, lpars := range list {

		l := map[string]interface{}{
			Attr_InstanceID:           *lpars.PvmInstanceID,
			Attr_InstanceName:         *lpars.ServerName,
			Attr_InstanceHref:         *lpars.Href,
			Attr_InstanceStatus:       *lpars.Status,
			Attr_InstanceSysType:      lpars.SysType,
			Attr_InstanceCreationDate: lpars.CreationDate.String(),
		}
		pvms = append(pvms, l)

	}
	return pvms
}
