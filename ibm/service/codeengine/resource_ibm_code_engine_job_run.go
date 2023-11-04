// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIbmCodeEngineJobRun() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmCodeEngineJobRunCreate,
		ReadContext:   resourceIbmCodeEngineJobRunRead,
		DeleteContext: resourceIbmCodeEngineJobRunDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "project_id"),
				Description:  "The ID of the project.",
			},
			"image_reference": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// ForceNew:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "image_reference"),
				Description:  "The name of the image that is used for this job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`.",
			},
			"image_secret": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "image_secret"),
				Description:  "The name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the job / job runs will be created but submitted job runs will fail, until this property is provided, too. This property must not be set on a job run, which references a job template.",
			},
			"job_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "job_name"),
				Description:  "Optional name of the job on which this job run is based on. If specified, the job run will inherit the configuration of the referenced job.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "name"),
				Description:  "The name of the job run. Use a name that is unique within the project.",
			},
			"run_arguments": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MinItems:    0,
				Description: "Set arguments for the job that are passed to start job run containers. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"run_as_user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// Default:     0,
				Description: "The user ID (UID) to run the application (e.g., 1001).",
			},
			"run_commands": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MinItems:    0,
				Description: "Set commands for the job that are passed to start job run containers. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"run_env_variables": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MinItems:    0,
				Description: "Optional references to config maps, secrets or a literal values.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The key to reference as environment variable.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the environment variable.",
						},
						"prefix": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A prefix that can be added to all keys of a full secret or config map reference.",
						},
						"reference": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the secret or config map.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "literal",
							Description: "Specify the type of the environment variable.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The literal value of the environment variable.",
						},
					},
				},
			},
			"run_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// Default:      "task",
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "run_mode"),
				Description:  "The mode for runs of the job. Valid values are `task` and `daemon`. In `task` mode, the `max_execution_time` and `retry_limit` properties apply. In `daemon` mode, since there is no timeout and failed instances are restarted indefinitely, the `max_execution_time` and `retry_limit` properties are not allowed.",
			},
			"run_service_account": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// Default:      "default",
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "run_service_account"),
				Description:  "The name of the service account. For built-in service accounts, you can use the shortened names `manager`, `none`, `reader`, and `writer`. This property must not be set on a job run, which references a job template.",
			},
			"run_volume_mounts": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MinItems:    0,
				Description: "Optional mounts of config maps or a secrets.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mount_path": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The path that should be mounted.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional name of the mount. If not set, it will be generated based on the `ref` and a random ID. In case the `ref` is longer than 58 characters, it will be cut off.",
						},
						"reference": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the referenced secret or config map.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specify the type of the volume mount. Allowed types are: 'config_map', 'secret'.",
						},
					},
				},
			},
			"scale_array_spec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// Default:      "0",
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "scale_array_spec"),
				Description:  "Define a custom set of array indices as comma-separated list containing single values and hyphen-separated ranges like `5,12-14,23,27`. Each instance can pick up its array index via environment variable `JOB_INDEX`. The number of unique array indices specified here determines the number of job instances to run.",
			},
			"scale_cpu_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// Default:      "1",
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "scale_cpu_limit"),
				Description:  "Optional amount of CPU set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo).",
			},
			"scale_ephemeral_storage_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// Default:      "400M",
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "scale_ephemeral_storage_limit"),
				Description:  "Optional amount of ephemeral storage to set for the instance of the job. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).",
			},
			"scale_max_execution_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// Default:     7200,
				Description: "The maximum execution time in seconds for runs of the job. This property can only be specified if `run_mode` is `task`.",
			},
			"scale_memory_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// Default:      "4G",
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_job_run", "scale_memory_limit"),
				Description:  "Optional amount of memory set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).",
			},
			"scale_retry_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// Default:     3,
				Description: "The number of times to rerun an instance of the job before the job is marked as failed. This property can only be specified if `run_mode` is `task`.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the resource was created.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When you provision a new job,  a URL is created identifying the location of the instance.",
			},
			"job_run_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier of the resource.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the job.",
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

func ResourceIbmCodeEngineJobRunValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "project_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "image_reference",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z0-9][a-z0-9\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\-_.\/]+[a-z0-9](:[\w][\w.\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$`,
			MinValueLength:             1,
			MaxValueLength:             256,
		},
		validate.ValidateSchema{
			Identifier:                 "image_secret",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$`,
			MinValueLength:             1,
			MaxValueLength:             253,
		},
		validate.ValidateSchema{
			Identifier:                 "job_name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "run_mode",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "daemon, task",
			Regexp:                     `^(task|daemon)$`,
			MinValueLength:             0,
		},
		validate.ValidateSchema{
			Identifier:                 "run_service_account",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "default, manager, none, reader, writer",
			Regexp:                     `^(manager|reader|writer|none|default)$`,
			MinValueLength:             0,
		},
		validate.ValidateSchema{
			Identifier:                 "scale_array_spec",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(?:[1-9]\d\d\d\d\d\d|[1-9]\d\d\d\d\d|[1-9]\d\d\d\d|[1-9]\d\d\d|[1-9]\d\d|[1-9]?\d)(?:-(?:[1-9]\d\d\d\d\d\d|[1-9]\d\d\d\d\d|[1-9]\d\d\d\d|[1-9]\d\d\d|[1-9]\d\d|[1-9]?\d))?(?:,(?:[1-9]\d\d\d\d\d\d|[1-9]\d\d\d\d\d|[1-9]\d\d\d\d|[1-9]\d\d\d|[1-9]\d\d|[1-9]?\d)(?:-(?:[1-9]\d\d\d\d\d\d|[1-9]\d\d\d\d\d|[1-9]\d\d\d\d|[1-9]\d\d\d|[1-9]\d\d|[1-9]?\d))?)*$`,
			MinValueLength:             1,
			MaxValueLength:             253,
		},
		validate.ValidateSchema{
			Identifier:                 "scale_cpu_limit",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([0-9.]+)([eEinumkKMGTPB]*)$`,
			MinValueLength:             0,
			MaxValueLength:             10,
		},
		validate.ValidateSchema{
			Identifier:                 "scale_ephemeral_storage_limit",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([0-9.]+)([eEinumkKMGTPB]*)$`,
			MinValueLength:             0,
			MaxValueLength:             10,
		},
		validate.ValidateSchema{
			Identifier:                 "scale_memory_limit",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([0-9.]+)([eEinumkKMGTPB]*)$`,
			MinValueLength:             0,
			MaxValueLength:             10,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_code_engine_job_run", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmCodeEngineJobRunCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createJobRunOptions := &codeenginev2.CreateJobRunOptions{}

	createJobRunOptions.SetProjectID(d.Get("project_id").(string))
	createJobRunOptions.SetName(d.Get("name").(string))
	createJobRunOptions.SetJobName(d.Get("job_name").(string))

	if _, ok := d.GetOk("image_reference"); ok {
		createJobRunOptions.SetImageReference(d.Get("image_reference").(string))
	}
	if _, ok := d.GetOk("image_secret"); ok {
		createJobRunOptions.SetImageSecret(d.Get("image_secret").(string))
	}
	if _, ok := d.GetOk("run_arguments"); ok {
		var runArguments []string
		for _, v := range d.Get("run_arguments").([]interface{}) {
			runArgumentsItem := v.(string)
			runArguments = append(runArguments, runArgumentsItem)
		}
		createJobRunOptions.SetRunArguments(runArguments)
	}
	if _, ok := d.GetOk("run_as_user"); ok {
		createJobRunOptions.SetRunAsUser(int64(d.Get("run_as_user").(int)))
	}
	if _, ok := d.GetOk("run_commands"); ok {
		var runCommands []string
		for _, v := range d.Get("run_commands").([]interface{}) {
			runCommandsItem := v.(string)
			runCommands = append(runCommands, runCommandsItem)
		}
		createJobRunOptions.SetRunCommands(runCommands)
	}
	if _, ok := d.GetOk("run_env_variables"); ok {
		var runEnvVariables []codeenginev2.EnvVarPrototype
		for _, v := range d.Get("run_env_variables").([]interface{}) {
			value := v.(map[string]interface{})
			runEnvVariablesItem, err := resourceIbmCodeEngineJobMapToEnvVarPrototype(value)
			if err != nil {
				return diag.FromErr(err)
			}
			runEnvVariables = append(runEnvVariables, *runEnvVariablesItem)
		}
		createJobRunOptions.SetRunEnvVariables(runEnvVariables)
	}
	if _, ok := d.GetOk("run_mode"); ok {
		createJobRunOptions.SetRunMode(d.Get("run_mode").(string))
	}
	if _, ok := d.GetOk("run_service_account"); ok {
		createJobRunOptions.SetRunServiceAccount(d.Get("run_service_account").(string))
	}
	if _, ok := d.GetOk("run_volume_mounts"); ok {
		var runVolumeMounts []codeenginev2.VolumeMountPrototype
		for _, v := range d.Get("run_volume_mounts").([]interface{}) {
			value := v.(map[string]interface{})
			runVolumeMountsItem, err := resourceIbmCodeEngineJobMapToVolumeMountPrototype(value)
			if err != nil {
				return diag.FromErr(err)
			}
			runVolumeMounts = append(runVolumeMounts, *runVolumeMountsItem)
		}
		createJobRunOptions.SetRunVolumeMounts(runVolumeMounts)
	}
	if _, ok := d.GetOk("scale_array_spec"); ok {
		createJobRunOptions.SetScaleArraySpec(d.Get("scale_array_spec").(string))
	}
	if _, ok := d.GetOk("scale_cpu_limit"); ok {
		createJobRunOptions.SetScaleCpuLimit(d.Get("scale_cpu_limit").(string))
	}
	if _, ok := d.GetOk("scale_ephemeral_storage_limit"); ok {
		createJobRunOptions.SetScaleEphemeralStorageLimit(d.Get("scale_ephemeral_storage_limit").(string))
	}
	if _, ok := d.GetOk("scale_max_execution_time"); ok {
		createJobRunOptions.SetScaleMaxExecutionTime(int64(d.Get("scale_max_execution_time").(int)))
	}
	if _, ok := d.GetOk("scale_memory_limit"); ok {
		createJobRunOptions.SetScaleMemoryLimit(d.Get("scale_memory_limit").(string))
	}
	if _, ok := d.GetOk("scale_retry_limit"); ok {
		createJobRunOptions.SetScaleRetryLimit(int64(d.Get("scale_retry_limit").(int)))
	}

	jobRun, response, err := codeEngineClient.CreateJobRunWithContext(context, createJobRunOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateJobRunWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateJobRunWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createJobRunOptions.ProjectID, *jobRun.Name))

	_, err = waitForIbmCodeEngineJonRunCreate(context, d, meta)
	if err != nil {
		return diag.FromErr(fmt.Errorf(
			"Error waiting for resource IbmCodeEngineJobRun (%s) to be created: %s", d.Id(), err))
	}

	return resourceIbmCodeEngineJobRunRead(context, d, meta)
}

func waitForIbmCodeEngineJonRunCreate(context context.Context, d *schema.ResourceData, meta interface{}) (interface{}, error) {
	if d.Get("run_mode") == "daemon" {
		return true, nil
	}

	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return false, err
	}

	getJobRunOptions := &codeenginev2.GetJobRunOptions{}
	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return false, err
	}

	getJobRunOptions.SetProjectID(parts[0])
	getJobRunOptions.SetName(parts[1])

	stateConf := &resource.StateChangeConf{
		Pending: []string{"running"},
		Target:  []string{"failed", "completed"},
		Refresh: func() (interface{}, string, error) {
			stateObj, response, err := codeEngineClient.GetJobRun(getJobRunOptions)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The instance %s does not exist anymore: %s\n%s", "getJobRunOptions", err, response)
				}
				return nil, "", err
			}
			failStates := map[string]bool{"failed": true}
			if failStates[*stateObj.Status] {
				return stateObj, *stateObj.Status, fmt.Errorf("The instance %s failed: %s\n%s", "getJobRunOptions", err, response)
			}
			return stateObj, *stateObj.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      20 * time.Second,
		MinTimeout: 20 * time.Second,
	}

	return stateConf.WaitForStateContext(context)
}

func resourceIbmCodeEngineJobRunRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getJobRunOptions := &codeenginev2.GetJobRunOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getJobRunOptions.SetProjectID(parts[0])
	getJobRunOptions.SetName(parts[1])

	jobRun, response, err := codeEngineClient.GetJobRunWithContext(context, getJobRunOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetJobRunWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetJobRunWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("project_id", jobRun.ProjectID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting project_id: %s", err))
	}
	if err = d.Set("image_reference", jobRun.ImageReference); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting image_reference: %s", err))
	}
	if err = d.Set("name", jobRun.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("job_name", jobRun.JobName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if !core.IsNil(jobRun.ImageSecret) {
		if err = d.Set("image_secret", jobRun.ImageSecret); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting image_secret: %s", err))
		}
	}
	if !core.IsNil(jobRun.RunArguments) {
		if err = d.Set("run_arguments", jobRun.RunArguments); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting run_arguments: %s", err))
		}
	}
	if !core.IsNil(jobRun.RunAsUser) {
		if err = d.Set("run_as_user", flex.IntValue(jobRun.RunAsUser)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting run_as_user: %s", err))
		}
	}
	if !core.IsNil(jobRun.RunCommands) {
		if err = d.Set("run_commands", jobRun.RunCommands); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting run_commands: %s", err))
		}
	}
	if !core.IsNil(jobRun.RunEnvVariables) {
		runEnvVariables := []map[string]interface{}{}
		for _, runEnvVariablesItem := range jobRun.RunEnvVariables {
			runEnvVariablesItemMap, err := resourceIbmCodeEngineJobEnvVarToMap(&runEnvVariablesItem)
			if err != nil {
				return diag.FromErr(err)
			}
			runEnvVariables = append(runEnvVariables, runEnvVariablesItemMap)
		}
		if err = d.Set("run_env_variables", runEnvVariables); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting run_env_variables: %s", err))
		}
	}
	if !core.IsNil(jobRun.RunMode) {
		if err = d.Set("run_mode", jobRun.RunMode); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting run_mode: %s", err))
		}
	}
	if !core.IsNil(jobRun.RunServiceAccount) {
		if err = d.Set("run_service_account", jobRun.RunServiceAccount); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting run_service_account: %s", err))
		}
	}
	if !core.IsNil(jobRun.RunVolumeMounts) {
		runVolumeMounts := []map[string]interface{}{}
		for _, runVolumeMountsItem := range jobRun.RunVolumeMounts {
			runVolumeMountsItemMap, err := resourceIbmCodeEngineJobVolumeMountToMap(&runVolumeMountsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			runVolumeMounts = append(runVolumeMounts, runVolumeMountsItemMap)
		}
		if err = d.Set("run_volume_mounts", runVolumeMounts); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting run_volume_mounts: %s", err))
		}
	}
	if !core.IsNil(jobRun.ScaleArraySpec) {
		if err = d.Set("scale_array_spec", jobRun.ScaleArraySpec); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scale_array_spec: %s", err))
		}
	}
	if !core.IsNil(jobRun.ScaleCpuLimit) {
		if err = d.Set("scale_cpu_limit", jobRun.ScaleCpuLimit); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scale_cpu_limit: %s", err))
		}
	}
	if !core.IsNil(jobRun.ScaleEphemeralStorageLimit) {
		if err = d.Set("scale_ephemeral_storage_limit", jobRun.ScaleEphemeralStorageLimit); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scale_ephemeral_storage_limit: %s", err))
		}
	}
	if !core.IsNil(jobRun.ScaleMaxExecutionTime) {
		if err = d.Set("scale_max_execution_time", flex.IntValue(jobRun.ScaleMaxExecutionTime)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scale_max_execution_time: %s", err))
		}
	}
	if !core.IsNil(jobRun.ScaleMemoryLimit) {
		if err = d.Set("scale_memory_limit", jobRun.ScaleMemoryLimit); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scale_memory_limit: %s", err))
		}
	}
	if !core.IsNil(jobRun.ScaleRetryLimit) {
		if err = d.Set("scale_retry_limit", flex.IntValue(jobRun.ScaleRetryLimit)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scale_retry_limit: %s", err))
		}
	}
	if !core.IsNil(jobRun.CreatedAt) {
		if err = d.Set("created_at", jobRun.CreatedAt); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
		}
	}
	if !core.IsNil(jobRun.Href) {
		if err = d.Set("href", jobRun.Href); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
		}
	}
	if !core.IsNil(jobRun.ID) {
		if err = d.Set("job_run_id", jobRun.ID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting job_run_id: %s", err))
		}
	}
	if !core.IsNil(jobRun.ResourceType) {
		if err = d.Set("resource_type", jobRun.ResourceType); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
		}
	}
	if !core.IsNil(jobRun.Status) {
		if err = d.Set("status", jobRun.Status); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
		}
	}
	if !core.IsNil(jobRun.StatusDetails) {
		statusDetailsMap, err := resourceIbmCodeEngineJobRunDetailedStatusToMap(jobRun.StatusDetails)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("status_details", []map[string]interface{}{statusDetailsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting status_details: %s", err))
		}
	}

	return nil
}

func resourceIbmCodeEngineJobRunDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteJobRunOptions := &codeenginev2.DeleteJobRunOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteJobRunOptions.SetProjectID(parts[0])
	deleteJobRunOptions.SetName(parts[1])

	response, err := codeEngineClient.DeleteJobRunWithContext(context, deleteJobRunOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteJobRunWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteJobRunWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmCodeEngineJobRunMapToEnvVarPrototype(modelMap map[string]interface{}) (*codeenginev2.EnvVarPrototype, error) {
	model := &codeenginev2.EnvVarPrototype{}
	if modelMap["key"] != nil && modelMap["key"].(string) != "" {
		model.Key = core.StringPtr(modelMap["key"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["prefix"] != nil && modelMap["prefix"].(string) != "" {
		model.Prefix = core.StringPtr(modelMap["prefix"].(string))
	}
	if modelMap["reference"] != nil && modelMap["reference"].(string) != "" {
		model.Reference = core.StringPtr(modelMap["reference"].(string))
	}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	return model, nil
}

func resourceIbmCodeEngineJobRunDetailedStatusToMap(model *codeenginev2.JobRunStatus) (map[string]interface{}, error) {
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
