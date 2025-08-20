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
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func ResourceIbmPdrValidateApikey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmPdrValidateApikeyCreate,
		ReadContext:   resourceIbmPdrValidateApikeyRead,
		UpdateContext: resourceIbmPdrValidateApikeyUpdate,
		DeleteContext: resourceIbmPdrValidateApikeyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "instance id of instance to provision.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
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
			// "instance_id": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Computed:    true,
			// },
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIbmPdrValidateApikeyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	serviceInstanceValidateKeyOptions := &drautomationservicev1.ServiceInstanceValidateKeyOptions{}

	serviceInstanceValidateKeyOptions.SetInstanceID(d.Get("instance_id").(string))
	serviceInstanceValidateKeyOptions.SetApiKey(d.Get("api_key").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		serviceInstanceValidateKeyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		serviceInstanceValidateKeyOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	validationKeyResponse, _, err := drAutomationServiceClient.ServiceInstanceValidateKeyWithContext(context, serviceInstanceValidateKeyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ServiceInstanceValidateKeyWithContext failed: %s", err.Error()), "ibm_pdr_validate_apikey", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *serviceInstanceValidateKeyOptions.InstanceID, *validationKeyResponse.ID))

	return resourceIbmPdrValidateApikeyRead(context, d, meta)
}

func resourceIbmPdrValidateApikeyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	serviceInstanceGetKeyV1Options := &drautomationservicev1.ServiceInstanceGetKeyV1Options{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "read", "sep-id-parts").GetDiag()
	}

	serviceInstanceGetKeyV1Options.SetInstanceID(parts[0])
	serviceInstanceGetKeyV1Options.SetInstanceID(parts[1])
	if _, ok := d.GetOk("accept_language"); ok {
		serviceInstanceGetKeyV1Options.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		serviceInstanceGetKeyV1Options.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	validationKeyResponse, response, err := drAutomationServiceClient.ServiceInstanceGetKeyV1WithContext(context, serviceInstanceGetKeyV1Options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ServiceInstanceGetKeyV1WithContext failed: %s", err.Error()), "ibm_pdr_validate_apikey", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
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
		if err = d.Set("instance_id", validationKeyResponse.ID); err != nil {
			err = fmt.Errorf("Error setting instance_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "read", "set-instance_id").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_pdr_validate_apikey", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIbmPdrValidateApikeyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	serviceInstanceUpdateApiKeyOptions := &drautomationservicev1.ServiceInstanceUpdateApiKeyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_validate_apikey", "update", "sep-id-parts").GetDiag()
	}

	serviceInstanceUpdateApiKeyOptions.SetInstanceID(parts[0])
	serviceInstanceUpdateApiKeyOptions.SetInstanceID(parts[1])
	serviceInstanceUpdateApiKeyOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		serviceInstanceUpdateApiKeyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		serviceInstanceUpdateApiKeyOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}
	serviceInstanceUpdateApiKeyOptions.SetApiKey(d.Get("api_key").(string))

	_, _, err = drAutomationServiceClient.ServiceInstanceUpdateApiKeyWithContext(context, serviceInstanceUpdateApiKeyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ServiceInstanceUpdateApiKeyWithContext failed: %s", err.Error()), "ibm_pdr_validate_apikey", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIbmPdrValidateApikeyRead(context, d, meta)
}

func resourceIbmPdrValidateApikeyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
