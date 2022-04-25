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
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func DataSourceIBMSchematicsResourceQuery() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMSchematicsResourceQueryRead,

		Schema: map[string]*schema.Schema{
			"query_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Resource query Id.  Use `GET /v2/resource_query` API to look up the Resource query definition Ids  in your IBM Cloud account.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource type (cluster, vsi, icd, vpc).",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource query name.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource Query id.",
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
			"queries": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the query(workspaces).",
						},
						"query_condition": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the resource query param.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of the resource query param.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Description of resource query param variable.",
									},
								},
							},
						},
						"query_select": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of query selection parameters.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMSchematicsResourceQueryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getResourcesQueryOptions := &schematicsv1.GetResourcesQueryOptions{}

	getResourcesQueryOptions.SetQueryID(d.Get("query_id").(string))

	resourceQueryRecord, response, err := schematicsClient.GetResourcesQueryWithContext(context, getResourcesQueryOptions)
	if err != nil {
		log.Printf("[DEBUG] GetResourcesQueryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetResourcesQueryWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getResourcesQueryOptions.QueryID))

	if err = d.Set("type", resourceQueryRecord.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}

	if err = d.Set("name", resourceQueryRecord.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("id", resourceQueryRecord.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
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

	queries := []map[string]interface{}{}
	if resourceQueryRecord.Queries != nil {
		for _, modelItem := range resourceQueryRecord.Queries {
			modelMap, err := DataSourceIBMSchematicsResourceQueryResourceQueryToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			queries = append(queries, modelMap)
		}
	}
	if err = d.Set("queries", queries); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting queries %s", err))
	}

	return nil
}

func DataSourceIBMSchematicsResourceQueryResourceQueryToMap(model *schematicsv1.ResourceQuery) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.QueryType != nil {
		modelMap["query_type"] = *model.QueryType
	}
	if model.QueryCondition != nil {
		queryCondition := []map[string]interface{}{}
		for _, queryConditionItem := range model.QueryCondition {
			queryConditionItemMap, err := DataSourceIBMSchematicsResourceQueryResourceQueryParamToMap(&queryConditionItem)
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

func DataSourceIBMSchematicsResourceQueryResourceQueryParamToMap(model *schematicsv1.ResourceQueryParam) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	return modelMap, nil
}
