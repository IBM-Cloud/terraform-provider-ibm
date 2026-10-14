// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isDynamicRouteServerRouteReportJobDeleted = "deleted"
	isDynamicRouteServerRouteReportJobFailed  = "failed"
)

func ResourceIBMIsDynamicRouteServerRouteReportJob() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerRouteReportJobCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerRouteReportJobRead,
		UpdateContext: resourceIBMIsDynamicRouteServerRouteReportJobUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerRouteReportJobDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_route_report_job", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_route_report_job", "name"),
				Description:  "The name for this dynamic route server route report job. The name is unique across all route report jobs for the dynamic route server. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"schema_version": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_route_report_job", "schema_version"),
				Description:  "The schema version to use when generating the route report content. If not specified, defaults to `v1`.",
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
			"dynamic_route_server_route_report_job_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server route report job.",
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerRouteReportJobValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "dynamic_route_server_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "schema_version",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "v1",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_route_report_job", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerRouteReportJobCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDynamicRouteServerRouteReportJobOptions := &vpcv1.CreateDynamicRouteServerRouteReportJobOptions{}

	createDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	if _, ok := d.GetOk("name"); ok {
		createDynamicRouteServerRouteReportJobOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("schema_version"); ok {
		createDynamicRouteServerRouteReportJobOptions.SetSchemaVersion(d.Get("schema_version").(string))
	}

	dynamicRouteServerRouteReportJob, _, err := vpcClient.CreateDynamicRouteServerRouteReportJobWithContext(context, createDynamicRouteServerRouteReportJobOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createDynamicRouteServerRouteReportJobOptions.DynamicRouteServerID, *dynamicRouteServerRouteReportJob.ID))

	_, err = isWaitForDynamicRouteServerRouteReportJobDone(context, vpcClient, *createDynamicRouteServerRouteReportJobOptions.DynamicRouteServerID, *dynamicRouteServerRouteReportJob.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error waiting for dynamic route server route report job to complete: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIsDynamicRouteServerRouteReportJobRead(context, d, meta)
}

func isDynamicRouteServerRouteReportJobRefreshFunc(ctx context.Context, sess *vpcv1.VpcV1, dynamicRouteServerID, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{
			DynamicRouteServerID: &dynamicRouteServerID,
			ID:                   &id,
		}
		job, response, err := sess.GetDynamicRouteServerRouteReportJobWithContext(ctx, getOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting dynamic route server route report job (%s): %s\n%s", id, err, response)
		}
		status := ""
		if job.Status != nil {
			status = *job.Status
		}
		if status == vpcv1.DynamicRouteServerRouteReportJobStatusFailedConst {
			reasons := ""
			for _, reason := range job.StatusReasons {
				if reason.Code != nil && reason.Message != nil {
					reasons = fmt.Sprintf("%s [%s: %s]", reasons, *reason.Code, *reason.Message)
				}
			}
			return job, status, fmt.Errorf("[ERROR] Dynamic route server route report job (%s) failed%s", id, reasons)
		}
		return job, status, nil
	}
}

func isWaitForDynamicRouteServerRouteReportJobDone(ctx context.Context, sess *vpcv1.VpcV1, dynamicRouteServerID, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server route report job (%s) to complete.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{vpcv1.DynamicRouteServerRouteReportJobStatusQueuedConst, vpcv1.DynamicRouteServerRouteReportJobStatusRunningConst},
		Target:     []string{vpcv1.DynamicRouteServerRouteReportJobStatusSucceededConst, vpcv1.DynamicRouteServerRouteReportJobStatusFailedConst},
		Refresh:    isDynamicRouteServerRouteReportJobRefreshFunc(ctx, sess, dynamicRouteServerID, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isDynamicRouteServerRouteReportJobDeleteRefreshFunc(ctx context.Context, sess *vpcv1.VpcV1, dynamicRouteServerID, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{
			DynamicRouteServerID: &dynamicRouteServerID,
			ID:                   &id,
		}
		job, response, err := sess.GetDynamicRouteServerRouteReportJobWithContext(ctx, getOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return job, isDynamicRouteServerRouteReportJobDeleted, nil
			}
			return nil, isDynamicRouteServerRouteReportJobFailed, fmt.Errorf("[ERROR] The dynamic route server route report job %s failed to delete: %s\n%s", id, err, response)
		}
		status := vpcv1.DynamicRouteServerRouteReportJobStatusDeletingConst
		if job.Status != nil {
			status = *job.Status
		}
		return job, status, nil
	}
}

func isWaitForDynamicRouteServerRouteReportJobDeleted(ctx context.Context, sess *vpcv1.VpcV1, dynamicRouteServerID, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server route report job (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending: []string{
			vpcv1.DynamicRouteServerRouteReportJobStatusDeletingConst,
			vpcv1.DynamicRouteServerRouteReportJobStatusQueuedConst,
			vpcv1.DynamicRouteServerRouteReportJobStatusRunningConst,
			vpcv1.DynamicRouteServerRouteReportJobStatusSucceededConst,
			vpcv1.DynamicRouteServerRouteReportJobStatusFailedConst,
		},
		Target:     []string{isDynamicRouteServerRouteReportJobDeleted, isDynamicRouteServerRouteReportJobFailed},
		Refresh:    isDynamicRouteServerRouteReportJobDeleteRefreshFunc(ctx, sess, dynamicRouteServerID, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}

// getDynamicRouteServerRouteReportJobContentJSON returns the route report content as a JSON string.
// An empty string is returned when the job has not succeeded yet or the content is not available (404).
func getDynamicRouteServerRouteReportJobContentJSON(ctx context.Context, sess *vpcv1.VpcV1, dynamicRouteServerID, id string, status *string) (string, error) {
	if status == nil || *status != vpcv1.DynamicRouteServerRouteReportJobStatusSucceededConst {
		return "", nil
	}
	getContentOptions := &vpcv1.GetDynamicRouteServerRouteReportJobContentOptions{
		DynamicRouteServerID: &dynamicRouteServerID,
		ID:                   &id,
	}
	content, response, err := sess.GetDynamicRouteServerRouteReportJobContentWithContext(ctx, getContentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return "", nil
		}
		return "", fmt.Errorf("GetDynamicRouteServerRouteReportJobContentWithContext failed: %s", err.Error())
	}
	return ResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON(content)
}

// ResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON marshals the report content to a JSON string.
// A nil content yields an empty string.
func ResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON(content vpcv1.DynamicRouteServerRouteReportContentIntf) (string, error) {
	if core.IsNil(content) {
		return "", nil
	}
	b, err := json.Marshal(content)
	if err != nil {
		return "", fmt.Errorf("Error marshalling route report content: %s", err)
	}
	return string(b), nil
}

func resourceIBMIsDynamicRouteServerRouteReportJobRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "sep-id-parts").GetDiag()
	}
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		err = fmt.Errorf("Invalid ID %q: expected format <dynamic_route_server_id>/<route_report_job_id>", d.Id())
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "sep-id-parts").GetDiag()
	}

	getDynamicRouteServerRouteReportJobOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{}
	getDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
	getDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

	dynamicRouteServerRouteReportJob, response, err := vpcClient.GetDynamicRouteServerRouteReportJobWithContext(context, getDynamicRouteServerRouteReportJobOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("dynamic_route_server_id", parts[0]); err != nil {
		err = fmt.Errorf("Error setting dynamic_route_server_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-dynamic_route_server_id").GetDiag()
	}
	if err = d.Set("dynamic_route_server_route_report_job_id", dynamicRouteServerRouteReportJob.ID); err != nil {
		err = fmt.Errorf("Error setting dynamic_route_server_route_report_job_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-dynamic_route_server_route_report_job_id").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerRouteReportJob.Name) {
		if err = d.Set("name", dynamicRouteServerRouteReportJob.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerRouteReportJob.SchemaVersion) {
		if err = d.Set("schema_version", dynamicRouteServerRouteReportJob.SchemaVersion); err != nil {
			err = fmt.Errorf("Error setting schema_version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-schema_version").GetDiag()
		}
	}
	completedAt := ""
	if !core.IsNil(dynamicRouteServerRouteReportJob.CompletedAt) {
		completedAt = flex.DateTimeToString(dynamicRouteServerRouteReportJob.CompletedAt)
	}
	if err = d.Set("completed_at", completedAt); err != nil {
		err = fmt.Errorf("Error setting completed_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-completed_at").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServerRouteReportJob.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("format", dynamicRouteServerRouteReportJob.Format); err != nil {
		err = fmt.Errorf("Error setting format: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-format").GetDiag()
	}
	if err = d.Set("href", dynamicRouteServerRouteReportJob.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-href").GetDiag()
	}
	if err = d.Set("resource_type", dynamicRouteServerRouteReportJob.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-resource_type").GetDiag()
	}
	startedAt := ""
	if !core.IsNil(dynamicRouteServerRouteReportJob.StartedAt) {
		startedAt = flex.DateTimeToString(dynamicRouteServerRouteReportJob.StartedAt)
	}
	if err = d.Set("started_at", startedAt); err != nil {
		err = fmt.Errorf("Error setting started_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-started_at").GetDiag()
	}
	if err = d.Set("status", dynamicRouteServerRouteReportJob.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-status").GetDiag()
	}
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range dynamicRouteServerRouteReportJob.StatusReasons {
		statusReasonsItemMap, err := ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "status_reasons-to-map").GetDiag()
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		err = fmt.Errorf("Error setting status_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-status_reasons").GetDiag()
	}

	content, err := getDynamicRouteServerRouteReportJobContentJSON(context, vpcClient, parts[0], parts[1], dynamicRouteServerRouteReportJob.Status)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set("content", content); err != nil {
		err = fmt.Errorf("Error setting content: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-content").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerRouteReportJobUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "update", "sep-id-parts").GetDiag()
	}

	updateDynamicRouteServerRouteReportJobOptions := &vpcv1.UpdateDynamicRouteServerRouteReportJobOptions{}
	updateDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
	updateDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.DynamicRouteServerRouteReportJobPatch{}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}

	if hasChange {
		patch, err := patchVals.AsPatch()
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "update", "as-patch").GetDiag()
		}
		updateDynamicRouteServerRouteReportJobOptions.DynamicRouteServerRouteReportJobPatch = patch

		_, _, err = vpcClient.UpdateDynamicRouteServerRouteReportJobWithContext(context, updateDynamicRouteServerRouteReportJobOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerRouteReportJobRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerRouteReportJobDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "delete", "sep-id-parts").GetDiag()
	}

	deleteDynamicRouteServerRouteReportJobOptions := &vpcv1.DeleteDynamicRouteServerRouteReportJobOptions{}
	deleteDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
	deleteDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

	response, err := vpcClient.DeleteDynamicRouteServerRouteReportJobWithContext(context, deleteDynamicRouteServerRouteReportJobOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForDynamicRouteServerRouteReportJobDeleted(context, vpcClient, parts[0], parts[1], d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error waiting for dynamic route server route report job to be deleted: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(model *vpcv1.DynamicRouteServerRouteReportStatusReason) (map[string]interface{}, error) {
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
