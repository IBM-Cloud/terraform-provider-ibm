// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerRouteReportJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerRouteReportJobRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"dynamic_route_server_route_report_job_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"dynamic_route_server_route_report_job_id", "name"},
				Description:  "The dynamic route server route report job identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"dynamic_route_server_route_report_job_id", "name"},
				Description:  "The name for this dynamic route server route report job. The name is unique across all route report jobs for the dynamic route server.",
			},
			"completed_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this route report job completed. If absent, the route report job has not yet completed.",
			},
			"content": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The route report content as a JSON string. Only present once the job status is `succeeded`.",
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
	}
}

func dataSourceIBMIsDynamicRouteServerRouteReportJobRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)
	var dynamicRouteServerRouteReportJob *vpcv1.DynamicRouteServerRouteReportJob

	if id, ok := d.GetOk("dynamic_route_server_route_report_job_id"); ok {
		getDynamicRouteServerRouteReportJobOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{}
		getDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getDynamicRouteServerRouteReportJobOptions.SetID(id.(string))

		job, _, err := vpcClient.GetDynamicRouteServerRouteReportJobWithContext(context, getDynamicRouteServerRouteReportJobOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_route_report_job", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		dynamicRouteServerRouteReportJob = job
	} else {
		name := d.Get("name").(string)
		listDynamicRouteServerRouteReportJobsOptions := &vpcv1.ListDynamicRouteServerRouteReportJobsOptions{}
		listDynamicRouteServerRouteReportJobsOptions.SetDynamicRouteServerID(dynamicRouteServerID)

		pager, err := vpcClient.NewDynamicRouteServerRouteReportJobsPager(listDynamicRouteServerRouteReportJobsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		allItems, err := pager.GetAllWithContext(context)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerRouteReportJobsPager.GetAll() failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_route_report_job", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		for i := range allItems {
			if allItems[i].Name != nil && *allItems[i].Name == name {
				dynamicRouteServerRouteReportJob = &allItems[i]
				break
			}
		}
		if dynamicRouteServerRouteReportJob == nil {
			err = fmt.Errorf("No dynamic route server route report job found with name %s for dynamic route server %s", name, dynamicRouteServerID)
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId(fmt.Sprintf("%s/%s", dynamicRouteServerID, *dynamicRouteServerRouteReportJob.ID))

	if err = d.Set("dynamic_route_server_route_report_job_id", dynamicRouteServerRouteReportJob.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_route_report_job_id: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-dynamic_route_server_route_report_job_id").GetDiag()
	}
	if err = d.Set("name", dynamicRouteServerRouteReportJob.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-name").GetDiag()
	}
	completedAt := ""
	if !core.IsNil(dynamicRouteServerRouteReportJob.CompletedAt) {
		completedAt = flex.DateTimeToString(dynamicRouteServerRouteReportJob.CompletedAt)
	}
	if err = d.Set("completed_at", completedAt); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting completed_at: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-completed_at").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServerRouteReportJob.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("format", dynamicRouteServerRouteReportJob.Format); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting format: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-format").GetDiag()
	}
	if err = d.Set("href", dynamicRouteServerRouteReportJob.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-href").GetDiag()
	}
	if err = d.Set("resource_type", dynamicRouteServerRouteReportJob.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("schema_version", dynamicRouteServerRouteReportJob.SchemaVersion); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting schema_version: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-schema_version").GetDiag()
	}
	startedAt := ""
	if !core.IsNil(dynamicRouteServerRouteReportJob.StartedAt) {
		startedAt = flex.DateTimeToString(dynamicRouteServerRouteReportJob.StartedAt)
	}
	if err = d.Set("started_at", startedAt); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting started_at: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-started_at").GetDiag()
	}
	if err = d.Set("status", dynamicRouteServerRouteReportJob.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-status").GetDiag()
	}
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range dynamicRouteServerRouteReportJob.StatusReasons {
		statusReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "status_reasons-to-map").GetDiag()
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-status_reasons").GetDiag()
	}

	content, err := getDynamicRouteServerRouteReportJobContentJSON(context, vpcClient, dynamicRouteServerID, *dynamicRouteServerRouteReportJob.ID, dynamicRouteServerRouteReportJob.Status)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set("content", content); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting content: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-content").GetDiag()
	}

	return nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(model *vpcv1.DynamicRouteServerRouteReportStatusReason) (map[string]interface{}, error) {
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
