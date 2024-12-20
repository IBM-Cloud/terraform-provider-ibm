// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func ResourceIbmSccScope() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMSccScopeCreate,
		ReadContext:   resourceIBMSccScopeRead,
		UpdateContext: resourceIBMSccScopeUpdate,
		DeleteContext: resourceIBMSccScopeDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_scc_scope", "instance_id"),
				Description:  "The ID of the Security and Compliance Center instance.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_scc_scope", "name"),
				Description:  "The scope name.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_scc_scope", "description"),
				Description:  "The scope description.",
			},
			"environment": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_scc_scope", "environment"),
				Description:  "The scope environment. This value details what cloud provider the scope targets.",
			},
			// Manual Change: change name and value for scope_type and scope_id
			"properties": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "A list of scopes/targets to exclude from a scope.",
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "account_id": &schema.Schema{
              Type:        schema.TypeString,
              Optional:    true,
              Description: "The ID of the account to target",
              ConflictsWith: []string{
                "properties.0.enterprise_id",
                "properties.0.resource_group_id",
                "properties.0.account_group_id"},
            },
            "enterprise_id": &schema.Schema{
              Type:        schema.TypeString,
              Optional:    true,
              Description: "The ID of the enterprise to target",
              ConflictsWith: []string{
                "properties.0.account_id",
                "properties.0.resource_group_id",
                "properties.0.account_group_id"},
            },
            "resource_group_id": &schema.Schema{
              Type:        schema.TypeString,
              Optional:    true,
              Description: "The ID of the resource group to target",
              ConflictsWith: []string{
                "properties.0.account_id",
                "properties.0.enterprise_id",
                "properties.0.account_group_id"},
            },
            "account_group_id": &schema.Schema{
              Type:        schema.TypeString,
              Optional:    true,
              Description: "The ID of the account group to target",
              ConflictsWith: []string{
                "properties.0.account_id",
                "properties.0.enterprise_id",
                "properties.0.resource_group_id"},
            },
          },
        },
			},
			"exclusions": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of scopes/targets to exclude from a scope.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the account to exclude. Only works with the attribute properties.0.account_group_id set",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the resource group to exclude. Only works with the attribute properties.0.account_id set",
						},
						"account_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the account group to exclude. Only works with the attribue properties.0.enterprise_id set",
						},
					},
				},
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the account associated with the scope.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier of the account or service ID who created the scope.",
			},
			"created_on": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the scope was created.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the user or service ID who updated the scope.",
			},
			"updated_on": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when the scope was updated.",
			},
			"attachment_count": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "The number of attachments tied to the scope.",
			},
			"scope_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the scope.",
			},
		},
	}
}

func ResourceIBMSccScopeValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "instance_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9_,'\s\-\.]*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9_,'\s\-\.]*$`,
			MinValueLength:             0,
			MaxValueLength:             512,
		},
		validate.ValidateSchema{
			Identifier:                 "environment",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9_,'\s\-\.]*$`,
			MinValueLength:             0,
			MaxValueLength:             128,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_scc_scope", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMSccScopeCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	securityAndComplianceCenterClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	createScopeOptions := &securityandcompliancecenterapiv3.CreateScopeOptions{}

	createScopeOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("name"); ok {
		createScopeOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createScopeOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("environment"); ok {
		createScopeOptions.SetEnvironment(d.Get("environment").(string))
	}
	props := []interface{}{}
	if _, ok := d.GetOk("properties"); ok {
		// Manual Change for scope properties
    
    if p, ok := d.Get("properties").(map[string]interface{}); ok {
      scopeProps, err := scopeTypeValueToProperties(p)
      if err != nil {
        return diag.FromErr(fmt.Errorf("Scope validation failed %s", err))
      }
      props = append(props, scopeProps)
    } else {
      return diag.FromErr(fmt.Errorf("Cannot convert scope properties to map[string]interface{}"))
    }
		// Removal
		// var properties []securityandcompliancecenterv3.ScopePropertyIntf
		// for _, v := range d.Get("properties").([]interface{}) {
		// 	value := v.(map[string]interface{})
		// 	propertiesItem, err := resourceIBMSccScopeMapToScopeProperty(value)
		// 	if err != nil {
		// 		return diag.FromErr(err)
		// 	}
		// 	properties = append(properties, propertiesItem)
		// }
	}
	if _, ok := d.GetOk("exclusions"); ok {
		exclusions := []securityandcompliancecenterapiv3.ScopePropertyExclusionItem{}
		for _, exclusionsItem := range d.Get("exclusions").([]interface{}) {
			exclusionsItemModel, err := resourceIBMSccScopeMapToScopePropertyExclusionItem(exclusionsItem.(map[string]interface{}))
			if err != nil {
				return diag.FromErr(fmt.Errorf("Scope validation failed %s", err))
			}
			exclusions = append(exclusions, *exclusionsItemModel)
		}
		props = append(props, exclusions)
	}
  var properties []securityandcompliancecenterapiv3.ScopePropertyIntf
  for _, v := range d.Get("properties").([]interface{}) {
    value := v.(map[string]interface{})
    propertiesItem, err := resourceIBMSccScopeMapToScopeProperty(value)
    if err != nil {
      return diag.FromErr(err)
    }
    properties = append(properties, propertiesItem)
  }
	createScopeOptions.SetProperties(properties)

	scope, response, err := securityAndComplianceCenterClient.CreateScopeWithContext(context, createScopeOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateScopeWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateScopeWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createScopeOptions.InstanceID, *scope.ID))

	return resourceIBMSccScopeRead(context, d, meta)
}

func resourceIBMSccScopeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	securityAndComplianceCenterClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	getScopeOptions := &securityandcompliancecenterapiv3.GetScopeOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getScopeOptions.SetInstanceID(parts[0])
	getScopeOptions.SetScopeID(parts[1])

	scope, response, err := securityAndComplianceCenterClient.GetScopeWithContext(context, getScopeOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetScopeWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetScopeWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("instance_id", scope.InstanceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instance_id: %s", err))
	}
	if !core.IsNil(scope.Name) {
		if err = d.Set("name", scope.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
	}
	if !core.IsNil(scope.Description) {
		if err = d.Set("description", scope.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
	}
	if !core.IsNil(scope.Environment) {
		if err = d.Set("environment", scope.Environment); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting environment: %s", err))
		}
	}
	// if !core.IsNil(scope.Properties) {
	// 	properties := []map[string]interface{}{}
	// 	for _, propertiesItem := range scope.Properties {
	// 		propertiesItemMap, err := resourceIBMSccScopeScopePropertyToMap(propertiesItem)
	// 		if err != nil {
	// 			return diag.FromErr(err)
	// 		}
	// 		properties = append(properties, propertiesItemMap)
	// 	}
	// 	if err = d.Set("properties", properties); err != nil {
	// 		return diag.FromErr(fmt.Errorf("Error setting properties: %s", err))
	// 	}
	// }
	if err = d.Set("account_id", scope.AccountID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
	}
	if err = d.Set("created_by", scope.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("created_on", flex.DateTimeToString(scope.CreatedOn)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_on: %s", err))
	}
	if err = d.Set("updated_by", scope.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}
	if err = d.Set("updated_on", flex.DateTimeToString(scope.UpdatedOn)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_on: %s", err))
	}
	if err = d.Set("attachment_count", scope.AttachmentCount); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting attachment_count: %s", err))
	}
	if err = d.Set("scope_id", scope.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting scope_id: %s", err))
	}

	return nil
}

func resourceIBMSccScopeUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	securityAndComplianceCenterClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	updateScopeOptions := &securityandcompliancecenterapiv3.UpdateScopeOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	updateScopeOptions.SetInstanceID(parts[0])
	updateScopeOptions.SetScopeID(parts[1])

	hasChange := false

	if d.HasChange("instance_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "instance_id"))
	}
	if d.HasChange("name") {
		updateScopeOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		updateScopeOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}

	if hasChange {
		_, response, err := securityAndComplianceCenterClient.UpdateScopeWithContext(context, updateScopeOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateScopeWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateScopeWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIBMSccScopeRead(context, d, meta)
}

func resourceIBMSccScopeDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	securityAndComplianceCenterClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteScopeOptions := &securityandcompliancecenterapiv3.DeleteScopeOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteScopeOptions.SetInstanceID(parts[0])
	deleteScopeOptions.SetScopeID(parts[1])

	response, err := securityAndComplianceCenterClient.DeleteScopeWithContext(context, deleteScopeOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteScopeWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteScopeWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

// scopeTypeValueToProperties will convert the custom terraform modification to accomadate to the API
func scopeTypeValueToProperties(modelMap map[string]interface{}) ([]interface{}, error) {
  
	properties := []interface{}{}
	scopeType := map[string]string{
		"name": "scope_type",
	}
	scopeId := map[string]string{
		"name": "scope_id",
	}
	if id, ok := modelMap["account_id"]; ok {
		scopeType["value"] = "account"
		scopeId["value"] = id.(string)
	} else if id, ok := modelMap["enterprise_id"]; ok {
		scopeType["value"] = "enterprise"
		scopeId["value"] = id.(string)
	} else if id, ok := modelMap["resource_group_id"]; ok {
		scopeType["value"] = "resource_group"
		scopeId["value"] = id.(string)
	} else if id, ok := modelMap["account_group_id"]; ok {
		scopeType["value"] = "account_group"
		scopeId["value"] = id.(string)
	} else {
		err := errors.New("unsupported property type")
		return properties, err
	}
	properties = append(properties, scopeType)
	properties = append(properties, scopeId)
	return properties, nil
}

func resourceIBMSccScopeMapToScopeProperty(modelMap map[string]interface{}) (securityandcompliancecenterapiv3.ScopePropertyIntf, error) {
	model := &securityandcompliancecenterapiv3.ScopeProperty{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil {
		model.Value = modelMap["value"].(string)
	}
	if modelMap["exclusions"] != nil {
		exclusions := []securityandcompliancecenterapiv3.ScopePropertyExclusionItem{}
		for _, exclusionsItem := range modelMap["exclusions"].([]interface{}) {
			exclusionsItemModel, err := resourceIBMSccScopeMapToScopePropertyExclusionItem(exclusionsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			exclusions = append(exclusions, *exclusionsItemModel)
		}
		model.Exclusions = exclusions
	}
	return model, nil
}

func resourceIBMSccScopeMapToScopePropertyExclusionItem(modelMap map[string]interface{}) (*securityandcompliancecenterapiv3.ScopePropertyExclusionItem, error) {
	model := &securityandcompliancecenterapiv3.ScopePropertyExclusionItem{}
	if modelMap["scope_id"] != nil && modelMap["scope_id"].(string) != "" {
		model.ScopeID = core.StringPtr(modelMap["scope_id"].(string))
	}
	if modelMap["scope_type"] != nil && modelMap["scope_type"].(string) != "" {
		model.ScopeType = core.StringPtr(modelMap["scope_type"].(string))
	}
	return model, nil
}

func resourceIBMSccScopeMapToScopePropertyScopeID(modelMap map[string]interface{}) (*securityandcompliancecenterapiv3.ScopePropertyScopeID, error) {
	model := &securityandcompliancecenterapiv3.ScopePropertyScopeID{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil {
		model.Value = modelMap["value"].(string)
	}
	return model, nil
}

func resourceIBMSccScopeMapToScopePropertyScopeType(modelMap map[string]interface{}) (*securityandcompliancecenterapiv3.ScopePropertyScopeType, error) {
	model := &securityandcompliancecenterapiv3.ScopePropertyScopeType{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	return model, nil
}

func resourceIBMSccScopeMapToScopePropertyExclusions(modelMap map[string]interface{}) (*securityandcompliancecenterapiv3.ScopePropertyExclusions, error) {
	model := &securityandcompliancecenterapiv3.ScopePropertyExclusions{}
	if modelMap["exclusions"] != nil {
		exclusions := []securityandcompliancecenterapiv3.ScopePropertyExclusionItem{}
		for _, exclusionsItem := range modelMap["exclusions"].([]interface{}) {
			exclusionsItemModel, err := resourceIBMSccScopeMapToScopePropertyExclusionItem(exclusionsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			exclusions = append(exclusions, *exclusionsItemModel)
		}
		model.Exclusions = exclusions
	}
	return model, nil
}

func resourceIBMSccScopeScopePropertyToMap(model securityandcompliancecenterapiv3.ScopePropertyIntf) (map[string]interface{}, error) {
	if _, ok := model.(*securityandcompliancecenterapiv3.ScopePropertyScopeID); ok {
		return resourceIBMSccScopeScopePropertyScopeIDToMap(model.(*securityandcompliancecenterapiv3.ScopePropertyScopeID))
	} else if _, ok := model.(*securityandcompliancecenterapiv3.ScopePropertyScopeType); ok {
		return resourceIBMSccScopeScopePropertyScopeTypeToMap(model.(*securityandcompliancecenterapiv3.ScopePropertyScopeType))
	} else if _, ok := model.(*securityandcompliancecenterapiv3.ScopePropertyExclusions); ok {
		return resourceIBMSccScopeScopePropertyExclusionsToMap(model.(*securityandcompliancecenterapiv3.ScopePropertyExclusions))
	} else if _, ok := model.(*securityandcompliancecenterapiv3.ScopeProperty); ok {
		modelMap := make(map[string]interface{})
		model := model.(*securityandcompliancecenterapiv3.ScopeProperty)
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.Value != nil {
			modelMap["value"] = model.Value
		}
		if model.Exclusions != nil {
			exclusions := []map[string]interface{}{}
			for _, exclusionsItem := range model.Exclusions {
				exclusionsItemMap, err := resourceIBMSccScopeScopePropertyExclusionItemToMap(&exclusionsItem)
				if err != nil {
					return modelMap, err
				}
				exclusions = append(exclusions, exclusionsItemMap)
			}
			modelMap["exclusions"] = exclusions
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized securityandcompliancecenterv3.ScopePropertyIntf subtype encountered")
	}
}

func resourceIBMSccScopeScopePropertyExclusionItemToMap(model *securityandcompliancecenterapiv3.ScopePropertyExclusionItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ScopeID != nil {
		modelMap["scope_id"] = model.ScopeID
	}
	if model.ScopeType != nil {
		modelMap["scope_type"] = model.ScopeType
	}
	return modelMap, nil
}

func resourceIBMSccScopeScopePropertyScopeIDToMap(model *securityandcompliancecenterapiv3.ScopePropertyScopeID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func resourceIBMSccScopeScopePropertyScopeTypeToMap(model *securityandcompliancecenterapiv3.ScopePropertyScopeType) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func resourceIBMSccScopeScopePropertyExclusionsToMap(model *securityandcompliancecenterapiv3.ScopePropertyExclusions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Exclusions != nil {
		exclusions := []map[string]interface{}{}
		for _, exclusionsItem := range model.Exclusions {
			exclusionsItemMap, err := resourceIBMSccScopeScopePropertyExclusionItemToMap(&exclusionsItem)
			if err != nil {
				return modelMap, err
			}
			exclusions = append(exclusions, exclusionsItemMap)
		}
		modelMap["exclusions"] = exclusions
	}
	return modelMap, nil
}
