// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func ResourceIBMMetricsRouterSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMMetricsRouterSettingsCreate,
		ReadContext:   resourceIBMMetricsRouterSettingsRead,
		UpdateContext: resourceIBMMetricsRouterSettingsUpdate,
		DeleteContext: resourceIBMMetricsRouterSettingsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"metadata_region_primary": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_metrics_router_settings", "metadata_region_primary"),
				Description:  "To store all your meta data in a single region.",
			},
			"private_api_endpoint_only": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "If you set this true then you cannot access api through public network.",
			},
			"default_targets": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The target ID List. In the event that no routing rule causes the metrics to be sent to a target, these targets will receive the metrics.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"permitted_target_regions": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "If present then only these regions may be used to define a target.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"api_version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The lowest API version of targets or routes that customer might have under his or her account.",
			},
		},
	}
}

func ResourceIBMMetricsRouterSettingsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "metadata_region_primary",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 -_]`,
			MinValueLength:             3,
			MaxValueLength:             256,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_metrics_router_settings", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMMetricsRouterSettingsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceSettingsOptions := &metricsrouterv3.ReplaceSettingsOptions{}

	replaceSettingsOptions.SetMetadataRegionPrimary(d.Get("metadata_region_primary").(string))
	replaceSettingsOptions.SetPrivateAPIEndpointOnly(d.Get("private_api_endpoint_only").(bool))
	if _, ok := d.GetOk("default_targets"); ok {
		replaceSettingsOptions.SetDefaultTargets(resourceInterfaceToStringArray(d.Get("default_targets").([]interface{})))
	}
	if _, ok := d.GetOk("permitted_target_regions"); ok {
		replaceSettingsOptions.SetPermittedTargetRegions(resourceInterfaceToStringArray(d.Get("permitted_target_regions").([]interface{})))
	}

	settings, response, err := metricsRouterClient.ReplaceSettingsWithContext(context, replaceSettingsOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ReplaceSettingsWithContext failed %s\n%s", err, response))
	}

	d.SetId(*settings.MetadataRegionPrimary)

	return resourceIBMMetricsRouterSettingsRead(context, d, meta)
}

func resourceIBMMetricsRouterSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	getSettingsOptions := &metricsrouterv3.GetSettingsOptions{}

	settings, response, err := metricsRouterClient.GetSettingsWithContext(context, getSettingsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetSettingsWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("metadata_region_primary", settings.MetadataRegionPrimary); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting metadata_region_primary: %s", err))
	}
	if err = d.Set("private_api_endpoint_only", settings.PrivateAPIEndpointOnly); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting private_api_endpoint_only: %s", err))
	}
	if settings.DefaultTargets != nil {
		if err = d.Set("default_targets", settings.DefaultTargets); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting default_targets: %s", err))
		}
	}
	if settings.PermittedTargetRegions != nil {
		if err = d.Set("permitted_target_regions", settings.PermittedTargetRegions); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting permitted_target_regions: %s", err))
		}
	}
	if err = d.Set("api_version", flex.IntValue(settings.APIVersion)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting api_version: %s", err))
	}

	return nil
}

func resourceInterfaceToStringArray(resources []interface{}) (result []string) {
	result = make([]string, 0)
	for _, item := range resources {
		if item != nil {
			result = append(result, item.(string))
		} else {
			result = append(result, "")
		}
	}
	return result
}

func resourceIBMMetricsRouterSettingsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceSettingsOptions := &metricsrouterv3.ReplaceSettingsOptions{}

	hasChange := false
	newMetaDataRegionPrimary := d.Get("metadata_region_primary").(string)
	replaceSettingsOptions.SetMetadataRegionPrimary(newMetaDataRegionPrimary)
	replaceSettingsOptions.SetPrivateAPIEndpointOnly(d.Get("private_api_endpoint_only").(bool))

	hasChange = hasChange || d.HasChange("metadata_region_primary") || d.HasChange("private_api_endpoint_only") || d.HasChange("api_version") || d.HasChange("permitted_target_regions") || d.HasChange("default_targets")

	if d.HasChange("metadata_region_primary") {
		d.SetId(newMetaDataRegionPrimary)
	}

	replaceSettingsOptions.DefaultTargets = resourceInterfaceToStringArray(d.Get("default_targets").([]interface{}))

	replaceSettingsOptions.PermittedTargetRegions = resourceInterfaceToStringArray(d.Get("permitted_target_regions").([]interface{}))

	if hasChange {
		setting, response, err := metricsRouterClient.ReplaceSettingsWithContext(context, replaceSettingsOptions)
		if err != nil {
			log.Printf("[DEBUG] ReplaceSettingsWithContext failed %s\n%s", err, response)
			log.Printf("[DEBUG] ReplaceSettingsWithContext failed %v\n", replaceSettingsOptions)
			return diag.FromErr(fmt.Errorf("ReplaceSettingsWithContext failed %s\n%s", err, response))
		}
		d.SetId(*setting.MetadataRegionPrimary)
	}

	return resourceIBMMetricsRouterSettingsRead(context, d, meta)
}

func resourceIBMMetricsRouterSettingsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	// Retrieve old settings and put them for required fields.  Remove all other fields
	settings, getResponse, err := metricsRouterClient.GetSettingsWithContext(context, &metricsrouterv3.GetSettingsOptions{})
	if err != nil {
		log.Printf("[DEBUG] PutSettingsWithContext with GetSettingsWithContext failed %s\n%s", err, getResponse)
		return diag.FromErr(fmt.Errorf("GetSettingsWithContext failed %s\n%s", err, getResponse))
	}

	replaceSettingsOptions := &metricsrouterv3.ReplaceSettingsOptions{}

	replaceSettingsOptions.MetadataRegionPrimary = settings.MetadataRegionPrimary
	replaceSettingsOptions.PrivateAPIEndpointOnly = settings.PrivateAPIEndpointOnly
	replaceSettingsOptions.PermittedTargetRegions = []string{}
	replaceSettingsOptions.DefaultTargets = []string{}

	_, response, err := metricsRouterClient.ReplaceSettingsWithContext(context, replaceSettingsOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ReplaceSettingsWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}
