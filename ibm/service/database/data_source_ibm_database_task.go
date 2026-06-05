// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

type dataSourceIBMDatabaseTaskBackend interface {
	Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
}

func pickDataSourceTaskBackend(d *schema.ResourceData, meta interface{}) (dataSourceIBMDatabaseTaskBackend, error) {
	taskID := d.Get("task_id").(string)

	// First, try to get the task to extract the deployment ID
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return nil, fmt.Errorf("error getting database client: %s", err)
	}

	getTaskOptions := &clouddatabasesv5.GetTaskOptions{
		ID: &taskID,
	}

	task, _, err := cloudDatabasesClient.GetTask(getTaskOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %s", err)
	}

	if task.Task == nil || task.Task.DeploymentID == nil {
		return nil, fmt.Errorf("task or deployment_id is nil")
	}

	deploymentID := *task.Task.DeploymentID

	// Get the resource controller client to fetch instance details using deployment ID
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, err
	}

	// Get the instance to check its plan
	instance, _, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{
		ID: &deploymentID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get resource instance: %s", err)
	}

	plan := *instance.ResourcePlanID
	if isGen2Plan(plan) {
		return newDataSourceIBMDatabaseTaskGen2Backend(), nil
	}
	return newDataSourceIBMDatabaseTaskClassicBackend(), nil
}

func DataSourceIBMDatabaseTask() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMDatabaseTaskRead,

		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Task ID.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable description of the task.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the task.",
			},
			"deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the deployment the task is being performed on.",
			},
			"progress_percent": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Indicator as percentage of progress of the task.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date and time when the task was created.",
			},
		},
	}
}

func dataSourceIBMDatabaseTaskRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	b, err := pickDataSourceTaskBackend(d, meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_task", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return b.Read(context, d, meta)
}

type dataSourceIBMDatabaseTaskClassicBackend struct{}

func newDataSourceIBMDatabaseTaskClassicBackend() dataSourceIBMDatabaseTaskBackend {
	return &dataSourceIBMDatabaseTaskClassicBackend{}
}

func (c *dataSourceIBMDatabaseTaskClassicBackend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	getTaskOptions := &clouddatabasesv5.GetTaskOptions{}
	getTaskOptions.SetID(d.Get("task_id").(string))

	task, response, err := cloudDatabasesClient.GetTaskWithContext(context, getTaskOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetTaskWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	d.SetId(*task.Task.ID)

	if err = d.Set("task_id", task.Task.ID); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting task_id: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	if task.Task.Description != nil {
		if err = d.Set("description", task.Task.Description); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_database_task", "read")
			return tfErr.GetDiag()
		}
	}

	if task.Task.Status != nil {
		if err = d.Set("status", task.Task.Status); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_database_task", "read")
			return tfErr.GetDiag()
		}
	}

	if task.Task.DeploymentID != nil {
		if err = d.Set("deployment_id", task.Task.DeploymentID); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting deployment_id: %s", err), "(Data) ibm_database_task", "read")
			return tfErr.GetDiag()
		}
	}

	if task.Task.ProgressPercent != nil {
		if err = d.Set("progress_percent", task.Task.ProgressPercent); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting progress_percent: %s", err), "(Data) ibm_database_task", "read")
			return tfErr.GetDiag()
		}
	}

	if task.Task.CreatedAt != nil {
		if err = d.Set("created_at", flex.DateTimeToString(task.Task.CreatedAt)); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_database_task", "read")
			return tfErr.GetDiag()
		}
	}

	return nil
}
