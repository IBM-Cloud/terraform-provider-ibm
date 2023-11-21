// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfigurationevaluation

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	acLib "github.com/IBM/appconfiguration-go-sdk/lib"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAppConfigEvaluateProperty() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppConfigurationPropertyRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "App Configuration instance id or guid.",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the environment created in App Configuration instance under the Environments section.",
			},
			"collection_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the collection created in App Configuration instance under the Collections section.",
			},
			"property_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Property id required to be evaluated.",
			},
			"entity_id": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Id of the Entity. This will be a string identifier related to the Entity against which" +
					"the property is evaluated. For example, an entity might be an instance of an app that runs on a mobile device," +
					"a microservice that runs on the cloud, or a component of infrastructure that runs that microservice." +
					"For any entity to interact with App Configuration, it must provide a unique entity ID.",
			},
			"entity_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: "Key value pair consisting of the attribute name and their values that defines the specified entity. " +
					"This is an optional parameter if the property is not configured with any targeting definition. If the " +
					"targeting is configured, then entityAttributes should be provided for the rule evaluation. An attribute is " +
					"a parameter that is used to define a segment. The SDK uses the attribute values to determine if the specified entity " +
					"satisfies the targeting rules, and returns the appropriate property value.",
				Sensitive: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"secret": {
				Type:        schema.TypeBool,
				Description: "Flag to check secret",
				Default:     false,
				Optional:    true,
			},
			"result_boolean": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Contains the evaluated value of the STRING type property only.",
			},
			"result_string": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Contains the evaluated value of the STRING type property only.",
			},
			"result_numeric": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Contains the evaluated value of the STRING type property only.",
			},
		},
	}
}

func DataSourceIBMAppConfigEvaluatePropertyValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "guid",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
		})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "environment_id",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
		})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "collection_id",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
		})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "property_id",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
		})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "entity_id",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
		})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "entity_attributes",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.ValueType(schema.TypeMap),
			Required:                   false,
		})

	ibmAppConfigEvaluatePropertyDataSourceValidator := validate.ResourceValidator{
		ResourceName: "ibm_app_config_evaluate_property",
		Schema:       validateSchema,
	}
	return &ibmAppConfigEvaluatePropertyDataSourceValidator
}

func dataSourceAppConfigurationPropertyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ac := acLib.GetInstance()
	sess, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}
	region := sess.Config.Region
	apiKey := sess.Config.BluemixAPIKey
	guid := d.Get("guid").(string)
	collectionId := d.Get("collection_id").(string)
	environmentId := d.Get("environment_id").(string)

	ac.Init(region, guid, apiKey)
	ac.SetContext(collectionId, environmentId)

	propertyId := d.Get("property_id").(string)
	entityId := d.Get("entity_id").(string)
	entityAttributes := make(map[string]interface{})
	if d.Get("entity_attributes") != nil {
		entityAttributes = d.Get("entity_attributes").(map[string]interface{})
	}

	property, err := ac.GetProperty(propertyId)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to retrieve property: %w", err))
	}

	d.SetId(guid + "_" + propertyId)

	result := property.GetCurrentValue(entityId, entityAttributes)

	switch result.(type) {
	case string:
		if err = d.Set("result_string", result.(string)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting the result: %s", err))
		}
	case float64:
		if err = d.Set("result_numeric", result); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting the result: %s", err))
		}
	case bool:
		if err = d.Set("result_boolean", result.(bool)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting the result: %s", err))
		}
	default:

		var resultMap map[string]interface{}
		var ok bool

		if resultMap, ok = result.(map[string]interface{}); ok {
			key := "secret_type"

			// Check if the key exists in the map
			if _, exists := resultMap[key]; exists {
				d.Set("secret", true)
			}
		}
		resultVal, err := json.Marshal(resultMap)
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error with json marshal : %w", err))
		}
		stringRes := strings.Replace(string(resultVal), "\"", "'", -1)

		if err = d.Set("result_string", stringRes); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting the result: %s", err))
		}
	}

	return nil
}
