---
layout: "ibm"
page_title: "IBM : schematics_job"
sidebar_current: "docs-ibm-resource-schematics-job"
description: |-
  Manages schematics_job.
---

# ibm\_schematics_job

Provides a resource for schematics_job. This allows schematics_job to be created, updated and deleted.

## Example Usage

```hcl
resource "schematics_job" "schematics_job" {
}
```

## Argument Reference

The following arguments are supported:

* `command_object` - (Optional, string) Name of the Schematics automation resource.
* `command_object_id` - (Optional, string) Job command object ID (`workspace-id, action-id or control-id`).
* `command_name` - (Optional, string) Schematics job command name.
* `command_parameter` - (Optional, string) Schematics job command parameter (`playbook-name, capsule-name or flow-name`).
* `command_options` - (Optional, List) Command line options for the command.
* `inputs` - (Optional, List) Job inputs used by an action.
  * `name` - (Optional, string) Name of the variable.
  * `value` - (Optional, string) Value for the variable or reference to the value.
  * `metadata` - (Optional, VariableMetadata) User editable metadata for the variables.
  * `link` - (Optional, string) Reference link to the variable value By default the expression will point to self.value.
* `settings` - (Optional, List) Environment variables used by the job while performing an action.
  * `name` - (Optional, string) Name of the variable.
  * `value` - (Optional, string) Value for the variable or reference to the value.
  * `metadata` - (Optional, VariableMetadata) User editable metadata for the variables.
  * `link` - (Optional, string) Reference link to the variable value By default the expression will point to self.value.
* `tags` - (Optional, List) User defined tags, while running the job.
* `location` - (Optional, string) List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics.
* `status` - (Optional, List) Job Status.
  * `action_job_status` - (Optional, JobStatusAction) Action Job Status.
* `data` - (Optional, List) Job data.
  * `job_type` - (Required, string) Type of the job.
  * `action_job_data` - (Optional, JobDataAction) Action Job data.
* `bastion` - (Optional, List) Complete target details with the user inputs and the system generated data.
  * `name` - (Optional, string) Target name.
  * `type` - (Optional, string) Target type (`cluster`, `vsi`, `icd`, `vpc`).
  * `description` - (Optional, string) Target description.
  * `resource_query` - (Optional, string) Resource selection query string.
  * `credential` - (Optional, string) Override credential for each resource.  Reference to credentials values, used by all the resources.
  * `id` - (Optional, string) Target ID.
  * `created_at` - (Optional, TypeString) Targets creation time.
  * `created_by` - (Optional, string) E-mail address of the user who created the targets.
  * `updated_at` - (Optional, TypeString) Targets updation time.
  * `updated_by` - (Optional, string) E-mail address of user who updated the targets.
  * `sys_lock` - (Optional, SystemLock) System lock status.
  * `resource_ids` - (Optional, []interface{}) Array of the resource IDs.
* `log_summary` - (Optional, List) Job log summary record.
  * `job_id` - (Optional, string) Workspace ID.
  * `job_type` - (Optional, string) Type of Job.
  * `log_start_at` - (Optional, TypeString) Job log start timestamp.
  * `log_analyzed_till` - (Optional, TypeString) Job log update timestamp.
  * `elapsed_time` - (Optional, float64) Job log elapsed time (`log_analyzed_till - log_start_at`).
  * `log_errors` - (Optional, []interface{}) Job log errors.
  * `repo_download_job` - (Optional, JobLogSummaryRepoDownloadJob) Repo download Job log summary.
  * `action_job` - (Optional, JobLogSummaryActionJob) Flow Job log summary.
* `x_github_token` - (Optional, string) Create a job record and launch the job.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the schematics_job.
* `name` - Job name, uniquely derived from the related action.
* `description` - Job description derived from the related action.
* `resource_group` - Resource group name derived from the related action.
* `submitted_at` - Job submission time.
* `submitted_by` - E-mail address of the user who submitted the job.
* `start_at` - Job start time.
* `end_at` - Job end time.
* `duration` - Duration of job execution, for example, `40 sec`.
* `targets_ini` - Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).
* `log_store_url` - Job log store URL.
* `state_store_url` - Job state store URL.
* `results_url` - Job results store URL.
* `updated_at` - Job status updation timestamp.
