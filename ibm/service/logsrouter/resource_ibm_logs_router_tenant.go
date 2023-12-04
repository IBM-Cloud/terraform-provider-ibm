// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouter

import (
	"context"
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/ibmlogsrouteropenapi30v0"
)

func ResourceIbmLogsRouterTenant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmLogsRouterTenantCreate,
		ReadContext:   resourceIbmLogsRouterTenantRead,
		UpdateContext: resourceIbmLogsRouterTenantUpdate,
		DeleteContext: resourceIbmLogsRouterTenantDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"target_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_tenant", "target_type"),
				Description:  "Type of log-sink.",
			},
			"target_host": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_tenant", "target_host"),
				Description:  "Host name of log-sink.",
			},
			"target_port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Network port of log sink.",
			},
			"target_instance_crn": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_tenant", "target_instance_crn"),
				Description:  "Cloud resource name of the log-sink target instance.",
			},
			"access_credential": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_tenant", "access_credential"),
				Description:  "Secret to connect to the log-sink",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account ID the tenant belongs to.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time stamp the tenant was originally created.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "time stamp the tenant was last updated.",
			},
		},
	}
}

func ResourceIbmLogsRouterTenantValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "target_type",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `[a-z,A-Z,0-9,-]`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "target_host",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `[a-z,A-Z,0-9,-,.]`,
			MinValueLength:             1,
			MaxValueLength:             256,
		},
		validate.ValidateSchema{
			Identifier:                 "target_instance_crn",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `[a-z,A-Z,0-9,:,-]`,
			MinValueLength:             1,
			MaxValueLength:             256,
		},
		validate.ValidateSchema{
			Identifier:                 "access_credential",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `[a-z,A-Z,0-9,-\\.;!?]`,
			MinValueLength:             10,
			MaxValueLength:             50,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs_router_tenant", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmLogsRouterTenantCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmLogsRouterOpenApi30Client, err := meta.(conns.ClientSession).IbmLogsRouterOpenApi30V0()
	if err != nil {
		return diag.FromErr(err)
	}

	createTenantOptions := &ibmlogsrouteropenapi30v0.CreateTenantOptions{}

	createTenantOptions.SetTargetType(d.Get("target_type").(string))
	createTenantOptions.SetTargetHost(d.Get("target_host").(string))
	createTenantOptions.SetTargetPort(int64(d.Get("target_port").(int)))
	createTenantOptions.SetAccessCredential(d.Get("access_credential").(string))
	createTenantOptions.SetTargetInstanceCrn(d.Get("target_instance_crn").(string))

	tenantDetailsResponse, response, err := ibmLogsRouterOpenApi30Client.CreateTenantWithContext(context, createTenantOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTenantWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateTenantWithContext failed %s\n%s", err, response))
	}

	d.SetId(tenantDetailsResponse.ID.String())

	return resourceIbmLogsRouterTenantRead(context, d, meta)
}

func resourceIbmLogsRouterTenantRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmLogsRouterOpenApi30Client, err := meta.(conns.ClientSession).IbmLogsRouterOpenApi30V0()
	if err != nil {
		return diag.FromErr(err)
	}

	getTenantDetailOptions := &ibmlogsrouteropenapi30v0.GetTenantDetailOptions{}
	tenantId := strfmt.UUID(d.Id())
	getTenantDetailOptions.SetTenantID(&tenantId)

	tenantDetailsResponse, response, err := ibmLogsRouterOpenApi30Client.GetTenantDetailWithContext(context, getTenantDetailOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTenantDetailWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTenantDetailWithContext failed %s\n%s", err, response))
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
	if !core.IsNil(tenantDetailsResponse.AccountID) {
		if err = d.Set("account_id", tenantDetailsResponse.AccountID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
		}
	}
	if !core.IsNil(tenantDetailsResponse.CreatedAt) {
		if err = d.Set("created_at", tenantDetailsResponse.CreatedAt); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
		}
	}
	if !core.IsNil(tenantDetailsResponse.UpdatedAt) {
		if err = d.Set("updated_at", tenantDetailsResponse.UpdatedAt); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
		}
	}

	return nil
}

func resourceIbmLogsRouterTenantUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmLogsRouterOpenApi30Client, err := meta.(conns.ClientSession).IbmLogsRouterOpenApi30V0()
	if err != nil {
		return diag.FromErr(err)
	}

	updateTargetOptions := &ibmlogsrouteropenapi30v0.UpdateTargetOptions{}
	tenantId := strfmt.UUID(d.Id())
	updateTargetOptions.SetTenantID(&tenantId)

	hasChange := false

	patchVals := &ibmlogsrouteropenapi30v0.TenantDetailsResponsePatch{}
	if d.HasChange("target_host") {
		newTargetHost := d.Get("target_host").(string)
		patchVals.TargetHost = &newTargetHost
		hasChange = true
	}
	if d.HasChange("target_port") {
		newTargetPort := int64(d.Get("target_port").(int))
		patchVals.TargetPort = &newTargetPort
		hasChange = true
	}
	if d.HasChange("access_credential") {
		newAccessCredential := d.Get("access_credential").(string)
		patchVals.AccessCredential = &newAccessCredential
		hasChange = true
	}
	if d.HasChange("target_instance_crn") {
		newTargetInstanceCrn := d.Get("target_instance_crn").(string)
		patchVals.TargetInstanceCrn = &newTargetInstanceCrn
		hasChange = true
	}

	if hasChange {
		updateTargetOptions.TenantDetailsResponsePatch, _ = patchVals.AsPatch()
		_, response, err := ibmLogsRouterOpenApi30Client.UpdateTargetWithContext(context, updateTargetOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateTargetWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateTargetWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmLogsRouterTenantRead(context, d, meta)
}

func resourceIbmLogsRouterTenantDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ibmLogsRouterOpenApi30Client, err := meta.(conns.ClientSession).IbmLogsRouterOpenApi30V0()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTenantOptions := &ibmlogsrouteropenapi30v0.DeleteTenantOptions{}
	tenantId := strfmt.UUID(d.Id())
	deleteTenantOptions.SetTenantID(&tenantId)

	_, response, err := ibmLogsRouterOpenApi30Client.DeleteTenantWithContext(context, deleteTenantOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTenantWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTenantWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}
