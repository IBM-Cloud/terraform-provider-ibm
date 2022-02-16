// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	// Arguments

	// Attributes
	TenantCreationDate    = "creation_date"
	TenantCloudInstances  = "cloud_instances"
	TenantCloudInstanceID = "cloud_instance_id"
	TenantRegion          = "region"
	TenantEnabled         = "enabled"
	TenantName            = "tenant_name"
)

func DataSourceIBMPITenant() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPITenantRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
			TenantCreationDate: {
				Type:     schema.TypeString,
				Computed: true,
			},
			TenantEnabled: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			TenantName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			TenantCloudInstances: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						TenantCloudInstanceID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						TenantRegion: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPITenantRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	//tenantid := d.Get("tenantid").(string)

	tenantC := instance.NewIBMPITenantClient(ctx, sess, cloudInstanceID)
	tenantData, err := tenantC.GetSelfTenant()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*tenantData.TenantID)
	d.Set(TenantCreationDate, tenantData.CreationDate.String())
	d.Set(TenantEnabled, tenantData.Enabled)

	if tenantData.CloudInstances != nil {
		d.Set(TenantName, tenantData.CloudInstances[0].Name)
	}

	if tenantData.CloudInstances != nil {
		tenants := make([]map[string]interface{}, len(tenantData.CloudInstances))
		for i, cloudinstance := range tenantData.CloudInstances {
			j := make(map[string]interface{})
			j[TenantRegion] = cloudinstance.Region
			j[TenantCloudInstanceID] = cloudinstance.CloudInstanceID
			tenants[i] = j
		}

		d.Set(TenantCloudInstances, tenants)
	}

	return nil
}
