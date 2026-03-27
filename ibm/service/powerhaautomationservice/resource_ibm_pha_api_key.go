// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func ResourceIBMPhaAPIKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPhaAPIKeyCreate,
		ReadContext:   resourceIBMPhaAPIKeyRead,
		DeleteContext: resourceIBMPhaAPIKeyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pha_instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_api_key", "pha_instance_id"),
				Description:  "instance id of instance to provision.",
			},
			"accept_language": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_api_key", "accept_language"),
				Description:  "The language requested for the return document.",
			},
			"if_none_match": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_api_key", "if_none_match"),
				Description:  "ETag for conditional requests (optional).",
			},
			"api_key": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Sensitive:    true,
				ValidateFunc: validate.InvokeValidator("ibm_pha_api_key", "api_key"),
				Description:  "The API key associated with the request.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier for the API key record.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMPhaAPIKeyValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "pha_instance_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9-]+$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		validate.ValidateSchema{
			Identifier:                 "accept_language",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_,;=.*]+$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		validate.ValidateSchema{
			Identifier:                 "if_none_match",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_,;=.*]+$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		validate.ValidateSchema{
			Identifier:                 "api_key",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9._:-]+$`,
			MinValueLength:             1,
			MaxValueLength:             2048,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_pha_api_key", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMPhaAPIKeyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_api_key", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createAPIKeyOptions := &powerhaautomationservicev1.CreateAPIKeyOptions{}

	createAPIKeyOptions.SetPhaInstanceID(d.Get("pha_instance_id").(string))
	if _, ok := d.GetOk("api_key"); ok {
		createAPIKeyOptions.SetAPIKey(d.Get("api_key").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		createAPIKeyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		createAPIKeyOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	_, response, err := powerhaAutomationServiceClient.CreateAPIKeyWithContext(context, createAPIKeyOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("CreateAPIKeyWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"CreateAPIKeyWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_api_key", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
		// tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateAPIKeyWithContext failed: %s", err.Error()), "ibm_pha_api_key", "create")
		// log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s", *createAPIKeyOptions.PhaInstanceID))

	return resourceIBMPhaAPIKeyRead(context, d, meta)
}

func resourceIBMPhaAPIKeyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_api_key", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAPIKeyOptions := &powerhaautomationservicev1.GetAPIKeyOptions{}

	instanceID := d.Id()

	getAPIKeyOptions.SetPhaInstanceID(instanceID)
	if _, ok := d.GetOk("accept_language"); ok {
		getAPIKeyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getAPIKeyOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	apiKeyResponse, response, err := powerhaAutomationServiceClient.GetAPIKeyWithContext(context, getAPIKeyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}

		detailedMsg := fmt.Sprintf("GetAPIKeyWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetAPIKeyWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_api_key", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
		// tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAPIKeyWithContext failed: %s", err.Error()), "ibm_pha_api_key", "read")
		// log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// return tfErr.GetDiag()
	}

	// if !core.IsNil(apiKeyResponse.APIKey) {
	// 	if err = d.Set("api_key", apiKeyResponse.APIKey); err != nil {
	// 		err = fmt.Errorf("Error setting api_key: %s", err)
	// 		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_api_key", "read", "set-api_key").GetDiag()
	// 	}
	// }
	if !core.IsNil(apiKeyResponse.ID) {
		if err = d.Set("pha_instance_id", extractInstanceIDFromCRN(*apiKeyResponse.ID)); err != nil {
			err = fmt.Errorf("Error setting pha_instance_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pha_api_key", "read", "set-pha_instance_id").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_pha_api_key", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMPhaAPIKeyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
