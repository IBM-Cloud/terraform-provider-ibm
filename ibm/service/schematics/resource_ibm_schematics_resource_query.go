// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics

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
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func ResourceIBMSchematicsResourceQuery() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMSchematicsResourceQueryCreate,
		ReadContext:   ResourceIBMSchematicsResourceQueryRead,
		UpdateContext: ResourceIBMSchematicsResourceQueryUpdate,
		DeleteContext: ResourceIBMSchematicsResourceQueryDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_resource_query", "type"),
				Description:  "Resource type (cluster, vsi, icd, vpc).",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Resource query name.",
			},
			"queries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type of the query(workspaces).",
						},
						"query_condition": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the resource query param.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Value of the resource query param.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Description of resource query param variable.",
									},
								},
							},
						},
						"query_select": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of query selection parameters.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource query creation time.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email address of user who created the Resource query.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource query updation time.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email address of user who updated the Resource query.",
			},
		},
	}
}

func ResourceIBMSchematicsResourceQueryValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "vsi",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_schematics_resource_query", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMSchematicsResourceQueryCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createResourceQueryOptions := &schematicsv1.CreateResourceQueryOptions{}

	if _, ok := d.GetOk("type"); ok {
		createResourceQueryOptions.SetType(d.Get("type").(string))
	}
	if _, ok := d.GetOk("name"); ok {
		createResourceQueryOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("queries"); ok {
		var queries []schematicsv1.ResourceQuery
		for _, e := range d.Get("queries").([]interface{}) {
			value := e.(map[string]interface{})
			queriesItem, err := ResourceIBMSchematicsResourceQueryMapToResourceQuery(value)
			if err != nil {
				return diag.FromErr(err)
			}
			queries = append(queries, *queriesItem)
		}
		createResourceQueryOptions.SetQueries(queries)
	}

	resourceQueryRecord, response, err := schematicsClient.CreateResourceQueryWithContext(context, createResourceQueryOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateResourceQueryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateResourceQueryWithContext failed %s\n%s", err, response))
	}

	d.SetId(*resourceQueryRecord.ID)

	return ResourceIBMSchematicsResourceQueryRead(context, d, meta)
}

func ResourceIBMSchematicsResourceQueryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getResourcesQueryOptions := &schematicsv1.GetResourcesQueryOptions{}

	getResourcesQueryOptions.SetQueryID(d.Id())

	resourceQueryRecord, response, err := schematicsClient.GetResourcesQueryWithContext(context, getResourcesQueryOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetResourcesQueryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetResourcesQueryWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("type", resourceQueryRecord.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	if err = d.Set("name", resourceQueryRecord.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	queries := []map[string]interface{}{}
	if resourceQueryRecord.Queries != nil {
		for _, queriesItem := range resourceQueryRecord.Queries {
			queriesItemMap, err := ResourceIBMSchematicsResourceQueryResourceQueryToMap(&queriesItem)
			if err != nil {
				return diag.FromErr(err)
			}
			queries = append(queries, queriesItemMap)
		}
	}
	if err = d.Set("queries", queries); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting queries: %s", err))
	}
	if err = d.Set("created_at", flex.DateTimeToString(resourceQueryRecord.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("created_by", resourceQueryRecord.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(resourceQueryRecord.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("updated_by", resourceQueryRecord.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}

	return nil
}

func ResourceIBMSchematicsResourceQueryUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceResourcesQueryOptions := &schematicsv1.ReplaceResourcesQueryOptions{}

	replaceResourcesQueryOptions.SetQueryID(d.Id())
	if _, ok := d.GetOk("type"); ok {
		replaceResourcesQueryOptions.SetType(d.Get("type").(string))
	}
	if _, ok := d.GetOk("name"); ok {
		replaceResourcesQueryOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("queries"); ok {
		var queries []schematicsv1.ResourceQuery
		for _, e := range d.Get("queries").([]interface{}) {
			value := e.(map[string]interface{})
			queriesItem, err := ResourceIBMSchematicsResourceQueryMapToResourceQuery(value)
			if err != nil {
				return diag.FromErr(err)
			}
			queries = append(queries, *queriesItem)
		}
		replaceResourcesQueryOptions.SetQueries(queries)
	}

	_, response, err := schematicsClient.ReplaceResourcesQueryWithContext(context, replaceResourcesQueryOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceResourcesQueryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ReplaceResourcesQueryWithContext failed %s\n%s", err, response))
	}

	return ResourceIBMSchematicsResourceQueryRead(context, d, meta)
}

func ResourceIBMSchematicsResourceQueryDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteResourcesQueryOptions := &schematicsv1.DeleteResourcesQueryOptions{}

	deleteResourcesQueryOptions.SetQueryID(d.Id())

	response, err := schematicsClient.DeleteResourcesQueryWithContext(context, deleteResourcesQueryOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteResourcesQueryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteResourcesQueryWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMSchematicsResourceQueryMapToResourceQuery(modelMap map[string]interface{}) (*schematicsv1.ResourceQuery, error) {
	model := &schematicsv1.ResourceQuery{}
	if modelMap["query_type"] != nil && modelMap["query_type"].(string) != "" {
		model.QueryType = core.StringPtr(modelMap["query_type"].(string))
	}
	if modelMap["query_condition"] != nil {
		queryCondition := []schematicsv1.ResourceQueryParam{}
		for _, queryConditionItem := range modelMap["query_condition"].([]interface{}) {
			queryConditionItemModel, err := ResourceIBMSchematicsResourceQueryMapToResourceQueryParam(queryConditionItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			queryCondition = append(queryCondition, *queryConditionItemModel)
		}
		model.QueryCondition = queryCondition
	}
	if modelMap["query_select"] != nil {
		querySelect := []string{}
		for _, querySelectItem := range modelMap["query_select"].([]interface{}) {
			querySelect = append(querySelect, querySelectItem.(string))
		}
		model.QuerySelect = querySelect
	}
	return model, nil
}

func ResourceIBMSchematicsResourceQueryMapToResourceQueryParam(modelMap map[string]interface{}) (*schematicsv1.ResourceQueryParam, error) {
	model := &schematicsv1.ResourceQueryParam{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsResourceQueryResourceQueryToMap(model *schematicsv1.ResourceQuery) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.QueryType != nil {
		modelMap["query_type"] = model.QueryType
	}
	if model.QueryCondition != nil {
		queryCondition := []map[string]interface{}{}
		for _, queryConditionItem := range model.QueryCondition {
			queryConditionItemMap, err := ResourceIBMSchematicsResourceQueryResourceQueryParamToMap(&queryConditionItem)
			if err != nil {
				return modelMap, err
			}
			queryCondition = append(queryCondition, queryConditionItemMap)
		}
		modelMap["query_condition"] = queryCondition
	}
	if model.QuerySelect != nil {
		modelMap["query_select"] = model.QuerySelect
	}
	return modelMap, nil
}

func ResourceIBMSchematicsResourceQueryResourceQueryParamToMap(model *schematicsv1.ResourceQueryParam) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	return modelMap, nil
}
