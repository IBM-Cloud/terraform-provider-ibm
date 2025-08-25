// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func DataSourceIBMEnPreDefinedTemplates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMEnPreDefinedTemplatesRead,

		Schema: map[string]*schema.Schema{
			"instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"search_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter the template by name or type.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of templates.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of template.",
			},
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of source.",
			},
			"templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of templates.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template ID.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "description of the template.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of template.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of source.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Updated at.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMEnPreDefinedTemplatesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_en_pre_defined_templates", "list")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &en.ListPreDefinedTemplatesOptions{}

	options.SetInstanceID(d.Get("instance_guid").(string))

	if _, ok := d.GetOk("search_key"); ok {
		options.SetSearch(d.Get("search_key").(string))
	}

	if _, ok := d.GetOk("source"); ok {
		options.SetSource(d.Get("source").(string))
	}
	if _, ok := d.GetOk("type"); ok {
		options.SetType(d.Get("type").(string))
	}
	var templateList *en.PredefinedTemplatesList

	finalList := []en.PredefinedTemplate{}

	var offset int64 = 0
	var limit int64 = 100

	options.SetLimit(limit)

	for {
		options.SetOffset(offset)

		result, _, err := enClient.ListPreDefinedTemplatesWithContext(context, options)

		templateList = result

		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListPreDefinedTemplatesWithContext failed: %s", err.Error()), "(Data) ibm_en_pre_defined_templates", "list")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		offset = offset + limit

		finalList = append(finalList, result.Templates...)

		if offset > *result.TotalCount {
			break
		}
	}

	templateList.Templates = finalList

	d.SetId(fmt.Sprintf("Templates/%s", *options.InstanceID))

	if err = d.Set("total_count", flex.IntValue(templateList.TotalCount)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting total_count: %s", err), "(Data) ibm_en_pre_defined_templates", "list")
		return tfErr.GetDiag()
	}

	if templateList.Templates != nil {
		if err = d.Set("templates", enFlattenpredefinedtemplatesList(templateList.Templates)); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting Templates: %s", err), "(Data) ibm_en_pre_defined_templates", "list")
			return tfErr.GetDiag()
		}
	}

	return nil
}

func enFlattenpredefinedtemplatesList(result []en.PredefinedTemplate) (templates []map[string]interface{}) {
	for _, templateItem := range result {
		templates = append(templates, enpredefinedTemplateListToMap(templateItem))
	}

	return templates
}

func enpredefinedTemplateListToMap(templateItem en.PredefinedTemplate) (template map[string]interface{}) {
	template = map[string]interface{}{}

	if templateItem.ID != nil {
		template["id"] = templateItem.ID
	}
	if templateItem.Name != nil {
		template["name"] = templateItem.Name
	}
	if templateItem.Description != nil {
		template["description"] = templateItem.Description
	}
	if templateItem.Type != nil {
		template["type"] = templateItem.Type
	}
	if templateItem.Source != nil {
		template["source"] = templateItem.Source
	}
	if templateItem.UpdatedAt != nil {
		template["updated_at"] = flex.DateTimeToString(templateItem.UpdatedAt)
	}

	return template
}
