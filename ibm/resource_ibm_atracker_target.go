// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
)

func resourceIBMAtrackerTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMAtrackerTargetCreate,
		ReadContext:   resourceIBMAtrackerTargetRead,
		UpdateContext: resourceIBMAtrackerTargetUpdate,
		DeleteContext: resourceIBMAtrackerTargetDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the target. Must be 256 characters or less.",
			},
			"target_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: InvokeValidator("ibm_atracker_target", "target_type"),
				Description:  "The type of the target.",
			},
			"cos_endpoint": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "Property values for a Cloud Object Storage Endpoint.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The host name of this COS endpoint.",
						},
						"target_crn": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN of this COS instance.",
						},
						"bucket": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The bucket name under this COS instance.",
						},
						"api_key": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The IAM Api key that have writer access to this cos instance. This credential will be masked in the response.",
						},
					},
				},
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of this target type resource.",
			},
			"encrypt_key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption key used to encrypt events before ATracker services buffer them on storage. This credential will be masked in the response.",
			},
		},
	}
}

func resourceIBMAtrackerTargetValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "target_type",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              "cloud_object_storage",
		},
	)

	resourceValidator := ResourceValidator{ResourceName: "ibm_atracker_target", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMAtrackerTargetCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createTargetOptions := &atrackerv1.CreateTargetOptions{}

	createTargetOptions.SetName(d.Get("name").(string))
	createTargetOptions.SetTargetType(d.Get("target_type").(string))
	cosEndpoint := resourceIBMAtrackerTargetMapToCosEndpoint(d.Get("cos_endpoint.0").(map[string]interface{}))
	createTargetOptions.SetCosEndpoint(&cosEndpoint)

	target, response, err := atrackerClient.CreateTargetWithContext(context, createTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*target.ID)

	return resourceIBMAtrackerTargetRead(context, d, meta)
}

func resourceIBMAtrackerTargetMapToCosEndpoint(cosEndpointMap map[string]interface{}) atrackerv1.CosEndpoint {
	cosEndpoint := atrackerv1.CosEndpoint{}

	cosEndpoint.Endpoint = core.StringPtr(cosEndpointMap["endpoint"].(string))
	cosEndpoint.TargetCRN = core.StringPtr(cosEndpointMap["target_crn"].(string))
	cosEndpoint.Bucket = core.StringPtr(cosEndpointMap["bucket"].(string))
	cosEndpoint.APIKey = core.StringPtr(cosEndpointMap["api_key"].(string))

	return cosEndpoint
}

func resourceIBMAtrackerTargetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getTargetOptions := &atrackerv1.GetTargetOptions{}

	getTargetOptions.SetID(d.Id())

	target, response, err := atrackerClient.GetTargetWithContext(context, getTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	if err = d.Set("name", target.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("target_type", target.TargetType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_type: %s", err))
	}
	cosEndpointMap := resourceIBMAtrackerTargetCosEndpointToMap(*target.CosEndpoint)
	// This line is a workaround for api_key, which comes back as "REDACTED" from the service.
	// This causes havok in the tests (in legacy testing framework) so we store the original
	// api_key value into the state.  This is the least bad solution I could come up with.
	cosEndpointMap["api_key"] = d.Get("cos_endpoint.0.api_key")
	if err = d.Set("cos_endpoint", []map[string]interface{}{cosEndpointMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cos_endpoint: %s", err))
	}
	if err = d.Set("crn", target.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("encrypt_key", target.EncryptKey); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting encrypt_key: %s", err))
	}

	return nil
}

func resourceIBMAtrackerTargetCosEndpointToMap(cosEndpoint atrackerv1.CosEndpoint) map[string]interface{} {
	cosEndpointMap := map[string]interface{}{}

	cosEndpointMap["endpoint"] = cosEndpoint.Endpoint
	cosEndpointMap["target_crn"] = cosEndpoint.TargetCRN
	cosEndpointMap["bucket"] = cosEndpoint.Bucket
	cosEndpointMap["api_key"] = cosEndpoint.APIKey

	return cosEndpointMap
}

func resourceIBMAtrackerTargetUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceTargetOptions := &atrackerv1.ReplaceTargetOptions{}

	replaceTargetOptions.SetID(d.Id())
	replaceTargetOptions.SetName(d.Get("name").(string))
	replaceTargetOptions.SetTargetType(d.Get("target_type").(string))
	cosEndpoint := resourceIBMAtrackerTargetMapToCosEndpoint(d.Get("cos_endpoint.0").(map[string]interface{}))
	replaceTargetOptions.SetCosEndpoint(&cosEndpoint)

	_, response, err := atrackerClient.ReplaceTargetWithContext(context, replaceTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	return resourceIBMAtrackerTargetRead(context, d, meta)
}

func resourceIBMAtrackerTargetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTargetOptions := &atrackerv1.DeleteTargetOptions{}

	deleteTargetOptions.SetID(d.Id())

	response, err := atrackerClient.DeleteTargetWithContext(context, deleteTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
