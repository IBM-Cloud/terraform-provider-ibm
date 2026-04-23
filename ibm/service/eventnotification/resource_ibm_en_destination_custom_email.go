// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMEnCustomEmailDestination() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMEnCustomEmailDestinationCreate,
		ReadContext:   resourceIBMEnCustomEmailDestinationRead,
		UpdateContext: resourceIBMEnCustomEmailDestinationUpdate,
		DeleteContext: resourceIBMEnCustomEmailDestinationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Destintion name.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of Destination type smtp_custom.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Destination description.",
			},
			"collect_failed_events": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to collect the failed event in Cloud Object Storage bucket",
			},
			"is_sandbox": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set to true to create a sandbox custom email destination. Once upgraded to production (is_sandbox=false), it cannot be downgraded back to sandbox.",
			},
			"verification_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_en_destination_custom_email", "verification_type"),
				Description:  "Verification Method spf/dkim.",
			},
			"config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Payload describing a destination configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"params": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Domain for the Custom Domain Email Destination",
									},
									"dkim": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The DKIM attributes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"public_key": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "dkim public key.",
												},
												"selector": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "dkim selector.",
												},
												"verification": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "dkim verification.",
												},
											},
										},
									},
									"spf": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The SPF attributes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"txt_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "spf text name.",
												},
												"txt_value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "spf text value.",
												},
												"verification": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "spf verification.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"destination_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Destination ID",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last updated time.",
			},
			"subscription_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of subscriptions.",
			},
			"subscription_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of subscriptions.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceIBMEnEmailDestinationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "verification_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "spf,dkim",
			MinValueLength:             1,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_en_destination_custom_email", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMEnCustomEmailDestinationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_destination_custom_email", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &en.CreateDestinationOptions{}

	options.SetInstanceID(d.Get("instance_guid").(string))
	options.SetName(d.Get("name").(string))

	// Determine destination type based on is_sandbox parameter
	isSandbox := d.Get("is_sandbox").(bool)
	if isSandbox {
		options.SetType("smtp_custom_sandbox")
		log.Printf("[DEBUG] Creating sandbox custom email destination")
	} else {
		options.SetType(d.Get("type").(string))
		log.Printf("[DEBUG] Creating production custom email destination")
	}

	options.SetCollectFailedEvents(d.Get("collect_failed_events").(bool))

	if _, ok := d.GetOk("description"); ok {
		options.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("config"); ok {
		destinationType := options.Type
		config, err := CustomEmailDestinationMapToDestinationConfig(d.Get("config.0").(map[string]interface{}), *destinationType)
		if err != nil {
			return diag.FromErr(err)
		}
		options.SetConfig(&config)
	}

	result, _, err := enClient.CreateDestinationWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDestinationWithContext failed: %s", err.Error()), "ibm_en_destination_custom_email", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *options.InstanceID, *result.ID))

	return resourceIBMEnCustomEmailDestinationRead(context, d, meta)
}

func resourceIBMEnCustomEmailDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_destination_custom_email", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &en.GetDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_destination_custom_email", "read")
		return tfErr.GetDiag()
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])

	result, response, err := enClient.GetDestinationWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDestinationWithContext failed: %s", err.Error()), "ibm_en_destination_custom_email", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("instance_guid", options.InstanceID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting instance_guid: %s", err))
	}

	if err = d.Set("destination_id", options.ID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting destination_id: %s", err))
	}

	if err = d.Set("name", result.Name); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting name: %s", err))
	}

	if err = d.Set("type", result.Type); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting type: %s", err))
	}

	if err = d.Set("collect_failed_events", result.CollectFailedEvents); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting CollectFailedEvents: %s", err))
	}

	// Set is_sandbox based on destination type
	isSandbox := false
	if result.Type != nil && *result.Type == "smtp_custom_sandbox" {
		isSandbox = true
	}
	if err = d.Set("is_sandbox", isSandbox); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting is_sandbox: %s", err))
	}

	if err = d.Set("description", result.Description); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting description: %s", err))
	}

	if result.Config != nil {
		err = d.Set("config", enCustomEmailDestinationFlattenConfig(*result.Config))
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting config %s", err))
		}
	}

	if err = d.Set("updated_at", flex.DateTimeToString(result.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting updated_at: %s", err))
	}

	if err = d.Set("subscription_count", flex.IntValue(result.SubscriptionCount)); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting subscription_count: %s", err))
	}

	if err = d.Set("subscription_names", result.SubscriptionNames); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting subscription_names: %s", err))
	}

	return nil
}

func resourceIBMEnCustomEmailDestinationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_destination_custom_email", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_destination_custom_email", "update")
		return tfErr.GetDiag()
	}

	// Check if is_sandbox is being changed
	if d.HasChange("is_sandbox") {
		oldVal, newVal := d.GetChange("is_sandbox")
		oldSandbox := oldVal.(bool)
		newSandbox := newVal.(bool)

		// Prevent downgrade from production to sandbox
		if !oldSandbox && newSandbox {
			return diag.FromErr(fmt.Errorf("[ERROR] Cannot downgrade from production (is_sandbox=false) to sandbox (is_sandbox=true). Once upgraded to production, the destination cannot be converted back to sandbox"))
		}

		// Upgrade from sandbox to production
		if oldSandbox && !newSandbox {
			log.Printf("[DEBUG] Upgrading sandbox destination to production")

			// Get the domain from config
			if _, ok := d.GetOk("config"); !ok {
				return diag.FromErr(fmt.Errorf("[ERROR] Config with domain is required to upgrade sandbox destination to production"))
			}

			configMap := d.Get("config.0").(map[string]interface{})
			if configMap["params"] == nil || len(configMap["params"].([]interface{})) == 0 {
				return diag.FromErr(fmt.Errorf("[ERROR] Config params with domain is required to upgrade sandbox destination"))
			}

			paramsMap := configMap["params"].([]interface{})[0].(map[string]interface{})
			domain, ok := paramsMap["domain"].(string)
			if !ok || domain == "" {
				return diag.FromErr(fmt.Errorf("[ERROR] Domain is required to upgrade sandbox destination to production"))
			}

			upgradeOptions := &en.UpdateEmailSandboxDestinationOptions{}
			upgradeOptions.SetInstanceID(parts[0])
			upgradeOptions.SetID(parts[1])
			upgradeOptions.SetDomain(domain)

			log.Printf("[DEBUG] Calling UpdateSandboxDestination to upgrade with domain: %s", domain)
			_, response, err := enClient.UpdateEmailSandboxDestinationWithContext(context, upgradeOptions)
			if err != nil {
				log.Printf("[DEBUG] UpdateSandboxDestination failed. Response: %v", response)
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSandboxDestination failed: %s", err.Error()), "ibm_en_destination_custom_email", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			log.Printf("[DEBUG] Successfully upgraded sandbox destination to production")

			// After upgrade, continue with normal update flow for other fields
		}
	}

	// Normal update flow for non-sandbox destinations or after upgrade
	updateDestinationOptions := &en.UpdateDestinationOptions{}
	updateDestinationOptions.SetInstanceID(parts[0])
	updateDestinationOptions.SetID(parts[1])

	hasChange := false

	if d.HasChange("name") {
		newName := d.Get("name").(string)
		log.Printf("[DEBUG] Updating name to: %s", newName)
		updateDestinationOptions.SetName(newName)
		hasChange = true
	}
	if d.HasChange("description") {
		newDesc := d.Get("description").(string)
		log.Printf("[DEBUG] Updating description to: %s", newDesc)
		updateDestinationOptions.SetDescription(newDesc)
		hasChange = true
	}
	if d.HasChange("collect_failed_events") {
		newCollect := d.Get("collect_failed_events").(bool)
		log.Printf("[DEBUG] Updating collect_failed_events to: %v", newCollect)
		updateDestinationOptions.SetCollectFailedEvents(newCollect)
		hasChange = true
	}
	// Always check and send config if it exists, not just on change
	// This ensures domain updates are properly sent to the API
	if _, ok := d.GetOk("config"); ok {
		configChanged := d.HasChange("config")
		log.Printf("[DEBUG] Config exists, changed: %v", configChanged)

		// Determine current destination type
		currentType := d.Get("type").(string)
		isSandbox := d.Get("is_sandbox").(bool)
		if isSandbox {
			currentType = "smtp_custom_sandbox"
		}

		log.Printf("[DEBUG] Current destination type: %s", currentType)

		configData := d.Get("config.0")
		log.Printf("[DEBUG] Config data from state: %+v", configData)

		if configData != nil {
			config, err := CustomEmailDestinationMapToDestinationConfig(configData.(map[string]interface{}), currentType)
			if err != nil {
				log.Printf("[DEBUG] Error mapping config: %s", err.Error())
				return diag.FromErr(err)
			}

			// Log the config details
			if config.Params != nil {
				switch params := config.Params.(type) {
				case *en.DestinationConfigOneOfCustomDomainEmailDestinationConfig:
					if params.Domain != nil {
						log.Printf("[DEBUG] Setting config with domain: %s", *params.Domain)
						// Only set config if domain is actually present
						updateDestinationOptions.SetConfig(&config)
						hasChange = true
					} else {
						log.Printf("[DEBUG] No domain found in custom domain config")
					}
				case *en.DestinationConfigOneOfCustomEmailSandboxDestinationConfig:
					if params.Domain != nil {
						log.Printf("[DEBUG] Setting sandbox config with domain: %s", *params.Domain)
						updateDestinationOptions.SetConfig(&config)
						hasChange = true
					} else {
						log.Printf("[DEBUG] No domain found in sandbox config")
					}
				default:
					log.Printf("[DEBUG] Unknown config params type: %T", params)
				}
			} else {
				log.Printf("[DEBUG] Config.Params is nil")
			}
		}
	}

	if hasChange {
		log.Printf("[DEBUG] Calling UpdateDestinationWithContext for instance: %s, destination: %s", parts[0], parts[1])
		_, response, err := enClient.UpdateDestinationWithContext(context, updateDestinationOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateDestinationWithContext failed. Response: %v", response)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDestinationWithContext failed: %s", err.Error()), "ibm_en_destination_custom_email", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		log.Printf("[DEBUG] UpdateDestinationWithContext succeeded")
	}

	// Check if verification type needs to be updated
	if d.HasChange("verification_type") {
		log.Printf("[DEBUG] Verification type has changed")
		verifyOptions := &en.UpdateVerifyDestinationOptions{}
		verifyOptions.SetInstanceID(parts[0])
		verifyOptions.SetID(parts[1])
		verifyOptions.SetType(d.Get("verification_type").(string))

		_, _, err = enClient.UpdateVerifyDestinationWithContext(context, verifyOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVerifyDestinationWithContext failed: %s", err.Error()), "ibm_en_destination_custom_email", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		log.Printf("[DEBUG] UpdateVerifyDestinationWithContext succeeded")
	}

	log.Printf("[DEBUG] Calling Read to refresh state")
	return resourceIBMEnCustomEmailDestinationRead(context, d, meta)
}

func resourceIBMEnCustomEmailDestinationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_destination_custom_email", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &en.DeleteDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_destination_custom_email", "delete")
		return tfErr.GetDiag()
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])

	response, err := enClient.DeleteDestinationWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDestinationWithContext failed: %s", err.Error()), "ibm_en_destination_custom_email", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func CustomEmailDestinationMapToDestinationConfig(configMap map[string]interface{}, destinationType string) (en.DestinationConfig, error) {
	destinationConfig := en.DestinationConfig{}

	if configMap["params"] != nil && len(configMap["params"].([]interface{})) > 0 {
		paramsMap := configMap["params"].([]interface{})[0].(map[string]interface{})

		// Check destination type to determine which config struct to use
		if destinationType == "smtp_custom" {
			params := &en.DestinationConfigOneOfCustomDomainEmailDestinationConfig{}

			// Only send domain - DKIM and SPF are computed fields returned by the API
			if paramsMap["domain"] != nil && paramsMap["domain"].(string) != "" {
				params.Domain = core.StringPtr(paramsMap["domain"].(string))
			}

			destinationConfig.Params = params
		} else if destinationType == "smtp_custom_sandbox" {
			// Handle sandbox destination
			params := &en.DestinationConfigOneOfCustomEmailSandboxDestinationConfig{}

			// Only send domain - DKIM and SPF are computed fields returned by the API
			if paramsMap["domain"] != nil && paramsMap["domain"].(string) != "" {
				params.Domain = core.StringPtr(paramsMap["domain"].(string))
			}

			destinationConfig.Params = params
		}
	}

	return destinationConfig, nil
}
