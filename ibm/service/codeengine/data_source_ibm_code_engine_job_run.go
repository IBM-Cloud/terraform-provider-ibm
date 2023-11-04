// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

func DataSourceIbmCodeEngineJobRun() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmCodeEngineJobRunRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the project.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the job run.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the resource was created.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you provision a new job run, a URL is created identifying the location of the instance.",
			},
			"job_run_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier of the resource.",
			},
			"job_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional name of the job reference of this job run. If specified, the job run will inherit the configuration of the referenced job.",
			},
			"image_reference": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the image that is used for this job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`.",
			},
			"image_secret": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the job / job runs will be created but submitted job runs will fail, until this property is provided, too. This property must not be set on a job run, which references a job template.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the job run.",
			},
			"run_arguments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Set arguments for the job that are passed to start job run containers. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"run_as_user": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The user ID (UID) to run the job (e.g., 1001).",
			},
			"run_commands": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Set commands for the job that are passed to start job run containers. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"run_env_variables": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "References to config maps, secrets or literal values, which are exposed as environment variables in the job run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The key to reference as environment variable.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the environment variable.",
						},
						"prefix": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A prefix that can be added to all keys of a full secret or config map reference.",
						},
						"reference": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the secret or config map.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the type of the environment variable.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The literal value of the environment variable.",
						},
					},
				},
			},
			"run_mode": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The mode for runs of the job. Valid values are `task` and `daemon`. In `task` mode, the `max_execution_time` and `retry_limit` properties apply. In `daemon` mode, since there is no timeout and failed instances are restarted indefinitely, the `max_execution_time` and `retry_limit` properties are not allowed.",
			},
			"run_service_account": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the service account. For built-in service accounts, you can use the shortened names `manager`, `none`, `reader`, and `writer`. This property must not be set on a job run, which references a job template.",
			},
			"run_volume_mounts": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Optional mounts of config maps or a secrets.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mount_path": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The path that should be mounted.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the mount.",
						},
						"reference": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the referenced secret or config map.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the type of the volume mount. Allowed types are: 'config_map', 'secret'.",
						},
					},
				},
			},
			"scale_array_spec": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Define a custom set of array indices as comma-separated list containing single values and hyphen-separated ranges like `5,12-14,23,27`. Each instance can pick up its array index via environment variable `JOB_INDEX`. The number of unique array indices specified here determines the number of job instances to run.",
			},
			"scale_cpu_limit": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional amount of CPU set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo).",
			},
			"scale_ephemeral_storage_limit": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional amount of ephemeral storage to set for the instance of the job. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).",
			},
			"scale_max_execution_time": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum execution time in seconds for runs of the job. This property can only be specified if `run_mode` is `task`.",
			},
			"scale_memory_limit": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional amount of memory set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).",
			},
			"scale_retry_limit": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of times to rerun an instance of the job before the job is marked as failed. This property can only be specified if `run_mode` is `task`.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the job run. Possible values: [failed, completed, running]",
			},
			"status_details": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The detailed status of the job run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"completion_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time the job run completed.",
						},
						"failed": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of failed job run instances.",
						},
						"pending": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of pending job run instances.",
						},
						"requested": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of requested job run instances.",
						},
						"running": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of running job run instances.",
						},
						"start_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time the job run started.",
						},
						"succeeded": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of succeeded job run instances.",
						},
						"unknown": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of job run instances with unknown state.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmCodeEngineJobRunRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getJobRunOptions := &codeenginev2.GetJobRunOptions{}

	getJobRunOptions.SetProjectID(d.Get("project_id").(string))
	getJobRunOptions.SetName(d.Get("name").(string))

	jobRun, response, err := codeEngineClient.GetJobRunWithContext(context, getJobRunOptions)
	if err != nil {
		log.Printf("[DEBUG] GetJobRunWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetJobRunWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *getJobRunOptions.ProjectID, *getJobRunOptions.Name))

	if err = d.Set("created_at", jobRun.CreatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	if err = d.Set("href", jobRun.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}

	if err = d.Set("job_run_id", jobRun.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting job_run_id: %s", err))
	}

	if err = d.Set("job_name", jobRun.JobName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting job_name: %s", err))
	}

	if err = d.Set("image_reference", jobRun.ImageReference); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting image_reference: %s", err))
	}

	if err = d.Set("image_secret", jobRun.ImageSecret); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting image_secret: %s", err))
	}

	if err = d.Set("resource_type", jobRun.ResourceType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
	}

	if err = d.Set("run_as_user", flex.IntValue(jobRun.RunAsUser)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting run_as_user: %s", err))
	}

	runEnvVariables := []map[string]interface{}{}
	if jobRun.RunEnvVariables != nil {
		for _, modelItem := range jobRun.RunEnvVariables {
			modelMap, err := dataSourceIbmCodeEngineJobRunEnvVarToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			runEnvVariables = append(runEnvVariables, modelMap)
		}
	}
	if err = d.Set("run_env_variables", runEnvVariables); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting run_env_variables %s", err))
	}

	if err = d.Set("run_mode", jobRun.RunMode); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting run_mode: %s", err))
	}

	if err = d.Set("run_service_account", jobRun.RunServiceAccount); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting run_service_account: %s", err))
	}

	runVolumeMounts := []map[string]interface{}{}
	if jobRun.RunVolumeMounts != nil {
		for _, modelItem := range jobRun.RunVolumeMounts {
			modelMap, err := dataSourceIbmCodeEngineJobRunVolumeMountToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			runVolumeMounts = append(runVolumeMounts, modelMap)
		}
	}
	if err = d.Set("run_volume_mounts", runVolumeMounts); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting run_volume_mounts %s", err))
	}

	if err = d.Set("scale_array_spec", jobRun.ScaleArraySpec); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting scale_array_spec: %s", err))
	}

	if err = d.Set("scale_cpu_limit", jobRun.ScaleCpuLimit); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting scale_cpu_limit: %s", err))
	}

	if err = d.Set("scale_ephemeral_storage_limit", jobRun.ScaleEphemeralStorageLimit); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting scale_ephemeral_storage_limit: %s", err))
	}

	if err = d.Set("scale_max_execution_time", flex.IntValue(jobRun.ScaleMaxExecutionTime)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting scale_max_execution_time: %s", err))
	}

	if err = d.Set("scale_memory_limit", jobRun.ScaleMemoryLimit); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting scale_memory_limit: %s", err))
	}

	if err = d.Set("scale_retry_limit", flex.IntValue(jobRun.ScaleRetryLimit)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting scale_retry_limit: %s", err))
	}

	if err = d.Set("status", jobRun.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
	}

	statusDetailsMap, err := dataSourceIbmCodeEngineJobRunDetailedStatusToMap(jobRun.StatusDetails)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("status_details", []map[string]interface{}{statusDetailsMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status_details: %s", err))
	}

	return nil
}

func dataSourceIbmCodeEngineJobRunEnvVarToMap(model *codeenginev2.EnvVar) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Key != nil {
		modelMap["key"] = model.Key
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Prefix != nil {
		modelMap["prefix"] = model.Prefix
	}
	if model.Reference != nil {
		modelMap["reference"] = model.Reference
	}
	modelMap["type"] = model.Type
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func dataSourceIbmCodeEngineJobRunVolumeMountToMap(model *codeenginev2.VolumeMount) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["mount_path"] = model.MountPath
	modelMap["name"] = model.Name
	modelMap["reference"] = model.Reference
	modelMap["type"] = model.Type
	return modelMap, nil
}

func dataSourceIbmCodeEngineJobRunDetailedStatusToMap(model *codeenginev2.JobRunStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CompletionTime != nil {
		modelMap["completion_time"] = model.CompletionTime
	}
	if model.Failed != nil {
		modelMap["failed"] = model.Failed
	}
	if model.Pending != nil {
		modelMap["pending"] = model.Pending
	}
	if model.Requested != nil {
		modelMap["requested"] = model.Requested
	}
	if model.Running != nil {
		modelMap["running"] = model.Running
	}
	if model.StartTime != nil {
		modelMap["start_time"] = model.StartTime
	}
	if model.Succeeded != nil {
		modelMap["succeeded"] = model.Succeeded
	}
	if model.Unknown != nil {
		modelMap["unknown"] = model.Unknown
	}
	return modelMap, nil
}
