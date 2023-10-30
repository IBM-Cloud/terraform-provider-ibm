---
layout: "ibm"
page_title: "IBM : ibm_schematics_agent"
description: |-
  Manages schematics_agent.
subcategory: "Schematics"
---

# ibm_schematics_agent

~> **Beta:** This resource is in Beta, and is subject to change.

Provides a resource for schematics_agent. This allows schematics_agent to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_agent" "schematics_agent_instance" {
  agent_infrastructure {
		infra_type = "ibm_kubernetes"
		cluster_id = "cluster_id"
		cluster_resource_group = "cluster_resource_group"
		cos_instance_name = "cos_instance_name"
		cos_bucket_name = "cos_bucket_name"
		cos_bucket_region = "cos_bucket_region"
  }
  agent_inputs {
		name = "name"
		value = "value"
		use_default = true
		metadata {
			type = "boolean"
			aliases = [ "aliases" ]
			description = "description"
			cloud_data_type = "cloud_data_type"
			default_value = "default_value"
			link_status = "normal"
			secure = true
			immutable = true
			hidden = true
			required = true
			options = [ "options" ]
			min_value = 1
			max_value = 1
			min_length = 1
			max_length = 1
			matches = "matches"
			position = 1
			group_by = "group_by"
			source = "source"
		}
		link = "link"
  }
  agent_kpi {
		availability_indicator = "available"
		lifecycle_indicator = "consistent"
		percent_usage_indicator = "percent_usage_indicator"
		application_indicators = [ null ]
		infra_indicators = [ null ]
  }
  agent_location = "us-south"
  agent_metadata {
		name = "purpose"
		value = ["git", "terraform", "ansible"]
  }
  description = "Create Agent"
  name = "MyDevAgent"
  resource_group = "Default"
  schematics_location = "us-south"
  user_state {
		state = "enable"
		set_by = "set_by"
		set_at = "2021-01-31T09:44:12Z"
  }
  version = "v1.0.0"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `agent_infrastructure` - (Required, List) The infrastructure parameters used by the agent.
Nested scheme for **agent_infrastructure**:
	* `cluster_id` - (Optional, String) The cluster ID where agent services will be running.
	* `cluster_resource_group` - (Optional, String) The resource group of the cluster (is it required?).
	* `cos_bucket_name` - (Optional, String) The COS bucket name used to store the logs.
	* `cos_bucket_region` - (Optional, String) The COS bucket region.
	* `cos_instance_name` - (Optional, String) The COS instance name to store the agent logs.
	* `infra_type` - (Optional, String) Type of target agent infrastructure.
	  * Constraints: Allowable values are: `ibm_kubernetes`, `ibm_openshift`, `ibm_satellite`.
* `agent_inputs` - (Optional, List) Additional input variables for the agent.
Nested scheme for **agent_inputs**:
	* `link` - (Computed, String) The reference link to the variable value By default the expression points to `$self.value`.
	* `metadata` - (Optional, List) An user editable metadata for the variables.
	Nested scheme for **metadata**:
		* `aliases` - (Optional, List) The list of aliases for the variable name.
		* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
		* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
		* `description` - (Optional, String) The description of the meta data.
		* `group_by` - (Optional, String) The display name of the group this variable belongs to.
		* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `link_status` - (Optional, String) The status of the link.
		  * Constraints: Allowable values are: `normal`, `broken`.
		* `matches` - (Optional, String) The regex for the variable value.
		* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
		* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
		* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
		* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
		* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
		* `position` - (Optional, Integer) The relative position of this variable in a list.
		* `required` - (Optional, Boolean) If the variable required?.
		* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
		* `source` - (Optional, String) The source of this meta-data.
		* `type` - (Optional, String) Type of the variable.
		  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
	* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
	* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
	* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
* `agent_kpi` - (Optional, List) Schematics Agent key performance indicators.
Nested scheme for **agent_kpi**:
	* `application_indicators` - (Optional, List) Agent application key performance indicators.
	* `availability_indicator` - (Optional, String) Overall availability indicator reported by the agent.
	  * Constraints: Allowable values are: `available`, `unavailable`, `error`.
	* `infra_indicators` - (Optional, List) Agent infrastructure key performance indicators.
	* `lifecycle_indicator` - (Optional, String) Overall lifecycle indicator reported by the agents.
	  * Constraints: Allowable values are: `consistent`, `inconsistent`, `obselete`.
	* `percent_usage_indicator` - (Optional, String) Percentage usage of the agent resources.
* `agent_location` - (Required, String) The location where agent is deployed in the user environment.
* `agent_metadata` - (Optional, List) The metadata of an agent.
Nested scheme for **agent_metadata**:
	* `name` - (Optional, String) Name of the metadata.
	* `value` - (Optional, List) Value of the metadata name.
* `description` - (Optional, String) Agent description.
* `name` - (Required, String) The name of the agent (must be unique, for an account).
* `resource_group` - (Required, String) The resource-group name for the agent.  By default, agent will be registered in Default Resource Group.
* `schematics_location` - (Required, String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.
* `tags` - (Optional, List) Tags for the agent.
* `user_state` - (Optional, List) User defined status of the agent.
Nested scheme for **user_state**:
	* `set_at` - (Computed, String) When the User who set the state of the Object.
	* `set_by` - (Computed, String) Name of the User who set the state of the Object.
	* `state` - (Optional, String) User-defined states  * `enable`  Agent is enabled by the user.  * `disable` Agent is disbaled by the user.
	  * Constraints: Allowable values are: `enable`, `disable`.
* `version` - (Required, String) Agent version.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_agent.
* `agent_crn` - (String) The agent crn, obtained from the Schematics agent deployment configuration.
* `created_at` - (String) The agent creation date-time.
* `creation_by` - (String) The email address of an user who created the agent.
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
* `system_state` - (List) Computed state of the agent.
Nested scheme for **system_state**:
	* `status_code` - (String) Agent Status.
	  * Constraints: Allowable values are: `error`, `normal`, `in_progress`, `pending`, `draft`.
	* `status_message` - (String) The agent status message.
* `updated_at` - (String) The agent registration updation time.
* `updated_by` - (String) Email address of user who updated the agent registration.

## Import

You can import the `ibm_schematics_agent` resource by using `id`. The agent resource id.

# Syntax
```
$ terraform import ibm_schematics_agent.schematics_agent <id>
```
