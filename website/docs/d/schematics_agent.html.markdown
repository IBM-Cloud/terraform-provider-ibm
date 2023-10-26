---
layout: "ibm"
page_title: "IBM : ibm_schematics_agent"
description: |-
  Get information about schematics_agent
subcategory: "Schematics"
---

# ibm_schematics_agent

Provides a read-only data source for schematics_agent. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_agent" "schematics_agent" {
	agent_id = "agent_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `agent_id` - (Required, Forces new resource, String) Agent ID to get the details of agent.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_agent.
* `agent_crn` - (String) The agent crn, obtained from the Schematics agent deployment configuration.

* `agent_infrastructure` - (List) The infrastructure parameters used by the agent.
Nested scheme for **agent_infrastructure**:
	* `cluster_id` - (String) The cluster ID where agent services will be running.
	* `cluster_resource_group` - (String) The resource group of the cluster (is it required?).
	* `cos_bucket_name` - (String) The COS bucket name used to store the logs.
	* `cos_bucket_region` - (String) The COS bucket region.
	* `cos_instance_name` - (String) The COS instance name to store the agent logs.
	* `infra_type` - (String) Type of target agent infrastructure.
	  * Constraints: Allowable values are: `ibm_kubernetes`, `ibm_openshift`, `ibm_satellite`.

* `agent_inputs` - (List) Additional input variables for the agent.
Nested scheme for **agent_inputs**:
	* `link` - (String) The reference link to the variable value By default the expression points to `$self.value`.
	* `metadata` - (List) An user editable metadata for the variables.
	Nested scheme for **metadata**:
		* `aliases` - (List) The list of aliases for the variable name.
		* `cloud_data_type` - (String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
		* `default_value` - (String) Default value for the variable only if the override value is not specified.
		* `description` - (String) The description of the meta data.
		* `group_by` - (String) The display name of the group this variable belongs to.
		* `hidden` - (Boolean) If **true**, the variable is not displayed on UI or Command line.
		* `immutable` - (Boolean) Is the variable readonly ?.
		* `link_status` - (String) The status of the link.
		  * Constraints: Allowable values are: `normal`, `broken`.
		* `matches` - (String) The regex for the variable value.
		* `max_length` - (Integer) The maximum length of the variable value. Applicable for the string type.
		* `max_value` - (Integer) The maximum value of the variable. Applicable for the integer type.
		* `min_length` - (Integer) The minimum length of the variable value. Applicable for the string type.
		* `min_value` - (Integer) The minimum value of the variable. Applicable for the integer type.
		* `options` - (List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
		* `position` - (Integer) The relative position of this variable in a list.
		* `required` - (Boolean) If the variable required?.
		* `secure` - (Boolean) Is the variable secure or sensitive ?.
		* `source` - (String) The source of this meta-data.
		* `type` - (String) Type of the variable.
		  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
	* `name` - (String) The name of the variable. For example, `name = "inventory username"`.
	* `use_default` - (Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
	* `value` - (String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.

* `agent_kpi` - (List) Schematics Agent key performance indicators.
Nested scheme for **agent_kpi**:
	* `application_indicators` - (List) Agent application key performance indicators.
	* `availability_indicator` - (String) Overall availability indicator reported by the agent.
	  * Constraints: Allowable values are: `available`, `unavailable`, `error`.
	* `infra_indicators` - (List) Agent infrastructure key performance indicators.
	* `lifecycle_indicator` - (String) Overall lifecycle indicator reported by the agents.
	  * Constraints: Allowable values are: `consistent`, `inconsistent`, `obselete`.
	* `percent_usage_indicator` - (String) Percentage usage of the agent resources.

* `agent_location` - (String) The location where agent is deployed in the user environment.

* `agent_metadata` - (List) The metadata of an agent.
Nested scheme for **agent_metadata**:
	* `name` - (String) Name of the metadata.
	* `value` - (List) Value of the metadata name.

* `created_at` - (String) The agent creation date-time.

* `creation_by` - (String) The email address of an user who created the agent.

* `description` - (String) Agent description.

* `id` - (String) The agent resource id.

* `name` - (String) The name of the agent (must be unique, for an account).

* `recent_deploy_job` - (List) Post-installations checks for Agent health.
Nested scheme for **recent_deploy_job**:
	* `agent_id` - (String) Id of the agent.
	* `agent_version` - (String) Agent version.
	* `is_redeployed` - (Boolean) True, when the same version of the agent was redeployed.
	* `job_id` - (String) Job Id.
	* `log_url` - (String) URL to the full agent deployment job logs.
	* `status_code` - (String) Final result of the agent deployment job.
	  * Constraints: Allowable values are: `pending`, `in-progress`, `success`, `failed`.
	* `status_message` - (String) The outcome of the agent deployment job, in a formatted log string.
	* `updated_at` - (String) The agent deploy job updation time.
	* `updated_by` - (String) Email address of user who ran the agent deploy job.

* `recent_health_job` - (List) Agent health check.
Nested scheme for **recent_health_job**:
	* `agent_id` - (String) Id of the agent.
	* `agent_version` - (String) Agent version.
	* `job_id` - (String) Job Id.
	* `log_url` - (String) URL to the full health-check job logs.
	* `status_code` - (String) Final result of the health-check job.
	  * Constraints: Allowable values are: `pending`, `in-progress`, `success`, `failed`.
	* `status_message` - (String) The outcome of the health-check job, in a formatted log string.
	* `updated_at` - (String) The agent health check job updation time.
	* `updated_by` - (String) Email address of user who ran the agent health check job.

* `recent_prs_job` - (List) Run a pre-requisite scanner for deploying agent.
Nested scheme for **recent_prs_job**:
	* `agent_id` - (String) Id of the agent.
	* `agent_version` - (String) Agent version.
	* `job_id` - (String) Job Id.
	* `log_url` - (String) URL to the full pre-requisite scanner job logs.
	* `status_code` - (String) Final result of the pre-requisite scanner job.
	  * Constraints: Allowable values are: `pending`, `in-progress`, `success`, `failed`.
	* `status_message` - (String) The outcome of the pre-requisite scanner job, in a formatted log string.
	* `updated_at` - (String) The agent prs job updation time.
	* `updated_by` - (String) Email address of user who ran the agent prs job.

* `resource_group` - (String) The resource-group name for the agent.  By default, agent will be registered in Default Resource Group.

* `schematics_location` - (String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.

* `system_state` - (List) Computed state of the agent.
Nested scheme for **system_state**:
	* `status_code` - (String) Agent Status.
	  * Constraints: Allowable values are: `error`, `normal`, `in_progress`, `pending`, `draft`.
	* `status_message` - (String) The agent status message.

* `tags` - (List) Tags for the agent.

* `updated_at` - (String) The agent registration updation time.

* `updated_by` - (String) Email address of user who updated the agent registration.

* `user_state` - (List) User defined status of the agent.
Nested scheme for **user_state**:
	* `set_at` - (String) When the User who set the state of the Object.
	* `set_by` - (String) Name of the User who set the state of the Object.
	* `state` - (String) User-defined states  * `enable`  Agent is enabled by the user.  * `disable` Agent is disbaled by the user.
	  * Constraints: Allowable values are: `enable`, `disable`.

* `version` - (String) Agent version.

