// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouter

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/logs-router-go-sdk/ibmlogsrouteropenapi30v0"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

func DataSourceIbmLogsRouterTenant() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsRouterTenantRead,

		Schema: map[string]*schema.Schema{
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance ID of the tenant.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account ID the tenant belongs to.",
			},
			"target_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of log-sink.",
			},
			"target_host": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host name of log-sink.",
			},
			"target_port": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Network port of log sink.",
			},
			"target_instance_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud resource name of the log-sink target instance.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time stamp the tenant was originally created.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "time stamp the tenant was last updated.",
			},
		},
	}
}

func dataSourceIbmLogsRouterTenantRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmLogsRouterOpenApi30Client, err := meta.(conns.ClientSession).IbmLogsRouterOpenApi30V0()
	if err != nil {
		return diag.FromErr(err)
	}

	getTenantDetailOptions := &ibmlogsrouteropenapi30v0.GetTenantDetailOptions{}

	getTenantDetailOptions.SetTenantID(core.UUIDPtr(strfmt.UUID(d.Get("tenant_id").(string))))

	tenantDetailsResponse, response, err := ibmLogsRouterOpenApi30Client.GetTenantDetailWithContext(context, getTenantDetailOptions)
	if err != nil {
		log.Printf("[DEBUG] GetTenantDetailWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTenantDetailWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getTenantDetailOptions.TenantID))

	if err = d.Set("account_id", tenantDetailsResponse.AccountID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
	}

	if err = d.Set("target_type", tenantDetailsResponse.TargetType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_type: %s", err))
	}

	if err = d.Set("target_host", tenantDetailsResponse.TargetHost); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_host: %s", err))
	}

	if err = d.Set("target_port", flex.IntValue(tenantDetailsResponse.TargetPort)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_port: %s", err))
	}

	if err = d.Set("target_instance_crn", tenantDetailsResponse.TargetInstanceCrn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_instance_crn: %s", err))
	}

	if err = d.Set("created_at", tenantDetailsResponse.CreatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	if err = d.Set("updated_at", tenantDetailsResponse.UpdatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	return nil
}
