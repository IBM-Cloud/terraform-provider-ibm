// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerRouteReportJobs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerRouteReportJobsRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"sort": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "-created_at",
				Description: "Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. For example, the value `-created_at` sorts the collection by the `created_at` property in descending order, and the value `name` sorts it by the `name` property in ascending order.",
			},
			"route_report_jobs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of dynamic route server route report jobs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"completed_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this route report job completed. If absent, the route report job has not yet completed.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this dynamic route server route report job was created.",
						},
						"format": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The format of the route report: `json` - JSON format.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dynamic route server route report job.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dynamic route server route report job.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this dynamic route server route report job. The name is unique across all route report jobs for the dynamic route server.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"schema_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The schema version of the route report content: `v1` - The initial route report content schema.",
						},
						"started_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this route report job started running. If absent, the route report job has not yet started.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the route report job: `deleting`, `failed`, `queued`, `running`, `succeeded`.",
						},
						"status_reasons": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reasons for the current dynamic route server route report status (if any).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A reason code for this route report status.",
									},
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the status reason.",
									},
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A link to documentation about this status reason.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIsDynamicRouteServerRouteReportJobsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listDynamicRouteServerRouteReportJobsOptions := &vpcv1.ListDynamicRouteServerRouteReportJobsOptions{}

	listDynamicRouteServerRouteReportJobsOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	if _, ok := d.GetOk("sort"); ok {
		listDynamicRouteServerRouteReportJobsOptions.SetSort(d.Get("sort").(string))
	}

	pager, err := vpcClient.NewDynamicRouteServerRouteReportJobsPager(listDynamicRouteServerRouteReportJobsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAllWithContext(context)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerRouteReportJobsPager.GetAll() failed %s", err), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsDynamicRouteServerRouteReportJobsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(&modelItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read", "route_report_jobs-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("route_report_jobs", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_report_jobs %s", err), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read", "route_report_jobs-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsDynamicRouteServerRouteReportJobsID returns a reasonable ID for the list.
func dataSourceIBMIsDynamicRouteServerRouteReportJobsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(model *vpcv1.DynamicRouteServerRouteReportJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.CompletedAt != nil {
		modelMap["completed_at"] = flex.DateTimeToString(model.CompletedAt)
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = flex.DateTimeToString(model.CreatedAt)
	}
	if model.Format != nil {
		modelMap["format"] = *model.Format
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	if model.SchemaVersion != nil {
		modelMap["schema_version"] = *model.SchemaVersion
	}
	if model.StartedAt != nil {
		modelMap["started_at"] = flex.DateTimeToString(model.StartedAt)
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range model.StatusReasons {
		statusReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	modelMap["status_reasons"] = statusReasons
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportStatusReasonToMap(model *vpcv1.DynamicRouteServerRouteReportStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, nil
	}
	if model.Code != nil {
		modelMap["code"] = *model.Code
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}
