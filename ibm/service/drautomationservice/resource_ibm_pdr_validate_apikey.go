// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
 */

package drautomationservice

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"

	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func ResourceIBMPdrValidateApikey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPdrValidateApikeyCreate,
		ReadContext:   resourceIBMPdrValidateApikeyRead,
		UpdateContext: resourceIBMPdrValidateApikeyUpdate,
		DeleteContext: resourceIBMPdrValidateApikeyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "instance id of instance to provision.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "api key",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Validation result message.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the API key.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of the API key.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMPdrValidateApikeyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createApikeyOptions := &drautomationservicev1.CreateApikeyOptions{}

	createApikeyOptions.SetInstanceID(d.Get("instance_id").(string))
	createApikeyOptions.SetAPIKey(d.Get("api_key").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		createApikeyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	_, response, err := drAutomationServiceClient.CreateApikeyWithContext(context, createApikeyOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("CreateApikeyWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"CreateApikeyWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_validate_apikey", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s", *createApikeyOptions.InstanceID))

	return resourceIBMPdrValidateApikeyRead(context, d, meta)
}

func resourceIBMPdrValidateApikeyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getApikeyOptions := &drautomationservicev1.GetApikeyOptions{}

	getApikeyOptions.SetInstanceID(d.Id())
	// if _, ok := d.GetOk("accept_language"); ok {
	// 	getApikeyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	// }

	validationKeyResponse, response, err := drAutomationServiceClient.GetApikeyWithContext(context, getApikeyOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetApikeyWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetApikeyWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		if response.StatusCode == 404 {
			d.SetId("")
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_validate_apikey", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	if !core.IsNil(validationKeyResponse.Description) {
		if err = d.Set("description", validationKeyResponse.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(validationKeyResponse.Status) {
		if err = d.Set("status", validationKeyResponse.Status); err != nil {
			err = fmt.Errorf("Error setting status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "read", "set-status").GetDiag()
		}
	}
	if !core.IsNil(validationKeyResponse.ID) {
		if err = d.Set("instance_id", d.Id()); err != nil {
			err = fmt.Errorf("Error setting instance_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "read", "set-instance_id").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_pdr_validate_apikey", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMPdrValidateApikeyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateApikeyOptions := &drautomationservicev1.UpdateApikeyOptions{}

	updateApikeyOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		updateApikeyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	updateApikeyOptions.SetAPIKey(d.Get("api_key").(string))

	_, response, err := drAutomationServiceClient.UpdateApikeyWithContext(context, updateApikeyOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("UpdateApikeyWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"UpdateApikeyWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		if response.StatusCode == 404 {
			d.SetId("")
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_validate_apikey", "update")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	return resourceIBMPdrValidateApikeyRead(context, d, meta)
}

func resourceIBMPdrValidateApikeyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
