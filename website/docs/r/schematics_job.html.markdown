---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_job"
sidebar_current: "docs-ibm-resource-schematics-job"
description: |-
  Manages Schematics job.
---

# ibm_schematics_job
Create, update, and delete `ibm_schematics_job`. For more information, about IBM Cloud Schematics job, refer to [setting up jobs](https://cloud.ibm.com/docs/schematics?topic=schematics-action-setup#action-jobs).

## Example usage

```terraform
resource "ibm_schematics_job" "schematics_job" {
  command_object = "action"
  command_object_id = "<action_id>"
  command_name = "ansible_playbook_run | ansible_playbook_check"
  command_parameter = "<yml_file_name>"
  location = "us-east"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `bastion`- (Optional, List) Complete target details with the user inputs and the system generated data.

  Nested scheme for `bastion`:
  - `credential`- (Optional, String) Override the credential for each resource. Reference to credentials values, used by all the resources.
  - `created_at`- (Optional, TypeString) The targets creation time.
  - `created_by`- (Optional, String) The Email address of the user who created the targets.
  - `description`- (Optional, String) The target description.
  - `id`- (Optional, String) The target ID.
  - `name`- (Optional, String) The target name.
  - `resource_query`- (Optional, String) The resource selection query string.
  - `resource_ids`- (Optional, []interface{}) An array of the resource IDs.
  - `sys_lock`- (Optional, SystemLock) The system lock status.
  - `type`- (Optional, String) The target type such as `cluster`, `vsi`, `icd`, `vpc`.
  - `updated_at`- (Optional, TypeString) The targets update time.
  - `updated_by`- (Optional, String) The Email address of user who updated the targets.
- `command_object`- (Optional, String) The name of the Schematics automation resource.
- `command_object_id`- (Optional, String) The job command object ID. Supported values are `workspace-id`, `action-id`, or `control-id`.
- `command_name`- (Optional, String) The Schematics job command name.
- `command_parameter`- (Optional, String) The Schematics job command parameter. Supported values are  `playbook-name`, `capsule-name`, or `flow-name`.
- `command_options`- (Optional, List) The command line options for the command.
- `data`- (Optional, List) The job data.

  Nested scheme for `data`:
  - `action_job_data`- (Optional, String) Action Job data.
  - `job_type` - (Required, String)Type of the job.
- `inputs`- (Optional, List) The job inputs used by an action.

  Nested scheme for `inputs`:
  - `link`- (Optional, String) The reference link to the variable value By default the expression will point to self.value. 
  - `metadata`- `VariableMetadata` - Optional - User editable metadata for the variables.
  - `name`- (Optional, String) The name of the variable.
  - `value`- (Optional, String) The value for the variable or reference to the value.
- `location`- (Optional, String) List of action locations supported by Schematics service. **Note** this does not limit the location of the resources provisioned by using Schematics.
- `log_summary`- (Optional, List) The job log summary record.

  Nested scheme for `log_summary`:
  - `action_job`- (Optional, String) The job log summary flow.
  - `elapsed_time`- (Optional, Float64) The job log elapsed time `log_analyzed_till`, `log_start_at`.
  - `job_id`- (Optional, String) The workspace ID.
  - `job_type`- (Optional, String) The type of Job.
  - `log_start_at`- `TypeString` - Optional - The job log start timestamp.
  - `log_analyzed_till`- (Optional, TypeString) The job log update timestamp.
  - `log_errors`- (Optional, []Interface{})  The job log errors.
  - `repo_download_job`- (Optional, String)  The repository download job log summary.
- `settings`- (Optional, List) Environment variables used by the job while performing an action.

  Nested scheme for `settings`:
  - `link`- (Optional, String) Reference link to the variable value. By default the expression will point to `self.value`.
  - `metadata`- (Optional, String) - User editable metadata for the variables.
  - `name`- (Optional, String) Name of the variable.
  - `value`- (Optional, String) Value for the variable or reference to the value.
- `status`- (Optional, List) The job status.

  Nested scheme for `status`:
  - `action_job_status`- (Optional, String) The action job status.
- `tags`- (Optional, List) User defined tags, while running the job.
- `x_github_token`- (Optional, String) Creates and launches the job record.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `description` - (String) The job description derived from the related action.
- `duration` - (String) The duration of the job execution, for example, `40 seconds`.
- `end_at` - (String) The job end time.
- `id` - (String) The unique identifier of the Schematics job.
- `log_store_url` - (String) The job log store URL.
- `name` - (String) The job name, uniquely derived from the related action.
- `resource_group` - (String) The resource group name derived from the related action.
- `results_url` - (String) The job results store URL.
- `submitted_at` - (String) The job submission time.
- `submitted_by` - (String) The Email address of the user who submitted the job.
- `start_at` - (String) The job start time.
- `state_store_url` - (String) The job state store URL.
- `targets_ini` - (String) An inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost] 172.22.192.6 [dbhost] 172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#inventory-host-grps).
- `updated_at` - (Timestamp) The job status update timestamp.
