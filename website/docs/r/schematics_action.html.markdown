---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_action"
sidebar_current: "docs-ibm-resource-schematics-action"
description: |-
  Manages schematics_action.
subcategory: "Schematics Service API"
---

# ibm_schematics_action

Provides a resource for schematics_action. This allows schematics_action to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_action" "schematics_action" {
  action_inputs {
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
  action_outputs {
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
  bastion {
		name = "name"
		host = "host"
  }
  bastion_credential {
		name = "name"
		value = "value"
		use_default = true
		metadata {
			type = "string"
			aliases = [ "aliases" ]
			description = "description"
			cloud_data_type = "cloud_data_type"
			default_value = "default_value"
			link_status = "normal"
			immutable = true
			hidden = true
			required = true
			position = 1
			group_by = "group_by"
			source = "source"
		}
		link = "link"
  }
  credentials {
		name = "name"
		value = "value"
		use_default = true
		metadata {
			type = "string"
			aliases = [ "aliases" ]
			description = "description"
			cloud_data_type = "cloud_data_type"
			default_value = "default_value"
			link_status = "normal"
			immutable = true
			hidden = true
			required = true
			position = 1
			group_by = "group_by"
			source = "source"
		}
		link = "link"
  }
  description = "The description of your action. The description can be up to 2048 characters long in size. **Example** you can use the description to stop the targets."
  name = "Stop Action"
  settings {
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
  source {
		source_type = "local"
		git {
			computed_git_repo_url = "computed_git_repo_url"
			git_repo_url = "git_repo_url"
			git_token = "git_token"
			git_repo_folder = "git_repo_folder"
			git_release = "git_release"
			git_branch = "git_branch"
		}
		catalog {
			catalog_name = "catalog_name"
			offering_name = "offering_name"
			offering_version = "offering_version"
			offering_kind = "offering_kind"
			offering_id = "offering_id"
			offering_version_id = "offering_version_id"
			offering_repo_url = "offering_repo_url"
		}
  }
  state {
		status_code = "normal"
		status_job_id = "status_job_id"
		status_message = "status_message"
  }
  sys_lock {
		sys_locked = true
		sys_locked_by = "sys_locked_by"
		sys_locked_at = "2021-01-31T09:44:12Z"
  }
  user_state {
		state = "draft"
		set_by = "set_by"
		set_at = "2021-01-31T09:44:12Z"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `action_inputs` - (Optional, List) Input variables for the Action.
Nested scheme for **action_inputs**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
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
* `action_outputs` - (Optional, List) Output variables for the Action.
Nested scheme for **action_outputs**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
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
* `bastion` - (Optional, List) Describes a bastion resource.
Nested scheme for **bastion**:
	* `host` - (Optional, String) Reference to the Inventory resource definition.
	* `name` - (Optional, String) Bastion Name(Unique).
* `bastion_connection_type` - (Optional, String) Type of connection to be used when connecting to bastion host. If the `inventory_connection_type=winrm`, then `bastion_connection_type` is not supported.
  * Constraints: Allowable values are: `ssh`.
* `bastion_credential` - (Optional, List) User editable credential variable data and system generated reference to the value.
Nested scheme for **bastion_credential**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
	* `metadata` - (Optional, List) An user editable metadata for the credential variables.
	Nested scheme for **metadata**:
		* `aliases` - (Optional, List) The list of aliases for the variable name.
		* `cloud_data_type` - (Optional, String) Cloud data type of the credential variable. eg. api_key, iam_token, profile_id.
		* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
		* `description` - (Optional, String) The description of the meta data.
		* `group_by` - (Optional, String) The display name of the group this variable belongs to.
		* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `link_status` - (Optional, String) The status of the link.
		  * Constraints: Allowable values are: `normal`, `broken`.
		* `position` - (Optional, Integer) The relative position of this variable in a list.
		* `required` - (Optional, Boolean) If the variable required?.
		* `source` - (Optional, String) The source of this meta-data.
		* `type` - (Optional, String) Type of the variable.
		  * Constraints: Allowable values are: `string`, `link`.
	* `name` - (Optional, String) The name of the credential variable.
	* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
	* `value` - (Optional, String) The credential value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
* `command_parameter` - (Optional, String) Schematics job command parameter (playbook-name).
* `credentials` - (Optional, List) credentials of the Action.
Nested scheme for **credentials**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
	* `metadata` - (Optional, List) An user editable metadata for the credential variables.
	Nested scheme for **metadata**:
		* `aliases` - (Optional, List) The list of aliases for the variable name.
		* `cloud_data_type` - (Optional, String) Cloud data type of the credential variable. eg. api_key, iam_token, profile_id.
		* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
		* `description` - (Optional, String) The description of the meta data.
		* `group_by` - (Optional, String) The display name of the group this variable belongs to.
		* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `link_status` - (Optional, String) The status of the link.
		  * Constraints: Allowable values are: `normal`, `broken`.
		* `position` - (Optional, Integer) The relative position of this variable in a list.
		* `required` - (Optional, Boolean) If the variable required?.
		* `source` - (Optional, String) The source of this meta-data.
		* `type` - (Optional, String) Type of the variable.
		  * Constraints: Allowable values are: `string`, `link`.
	* `name` - (Optional, String) The name of the credential variable.
	* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
	* `value` - (Optional, String) The credential value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
* `description` - (Optional, String) Action description.
* `inventory` - (Optional, String) Target inventory record ID, used by the action or ansible playbook.
* `inventory_connection_type` - (Optional, String) Type of connection to be used when connecting to remote host. **Note** Currently, WinRM supports only Windows system with the public IPs and do not support Bastion host.
  * Constraints: Allowable values are: `ssh`, `winrm`.
* `location` - (Optional, String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.
* `name` - (Optional, String) The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action.
* `resource_group` - (Optional, String) Resource-group name for an action. By default, an action is created in `Default` resource group.
* `settings` - (Optional, List) Environment variables for the Action.
Nested scheme for **settings**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
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
* `source` - (Optional, List) Source of templates, playbooks, or controls.
Nested scheme for **source**:
	* `catalog` - (Optional, List) The connection details to the IBM Cloud Catalog source.
	Nested scheme for **catalog**:
		* `catalog_name` - (Optional, String) The name of the private catalog.
		* `offering_id` - (Optional, String) The ID of an offering in the IBM Cloud Catalog.
		* `offering_kind` - (Optional, String) The type of an offering, in the IBM Cloud Catalog.
		* `offering_name` - (Optional, String) The name of an offering in the IBM Cloud Catalog.
		* `offering_repo_url` - (Optional, String) The repository URL of an offering, in the IBM Cloud Catalog.
		* `offering_version` - (Optional, String) The version string of an offering in the IBM Cloud Catalog.
		* `offering_version_id` - (Optional, String) The ID of an offering version the IBM Cloud Catalog.
	* `git` - (Optional, List) The connection details to the Git source repository.
	Nested scheme for **git**:
		* `computed_git_repo_url` - (Optional, String) The complete URL which is computed by the **git_repo_url**, **git_repo_folder**, and **branch**.
		* `git_branch` - (Optional, String) The name of the branch that are used to fetch the Git repository.
		* `git_release` - (Optional, String) The name of the release tag that are used to fetch the Git repository.
		* `git_repo_folder` - (Optional, String) The name of the folder in the Git repository, that contains the template.
		* `git_repo_url` - (Optional, String) The URL to the Git repository that can be used to clone the template.
		* `git_token` - (Optional, String) The Personal Access Token (PAT) to connect to the Git URLs.
	* `source_type` - (Required, String) Type of source for the Template.
	  * Constraints: Allowable values are: `local`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_git_lab`, `ibm_cloud_catalog`, `external_scm`.
* `source_readme_url` - (Optional, String) URL of the `README` file, for the source URL.
* `source_type` - (Optional, String) Type of source for the Template.
  * Constraints: Allowable values are: `local`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_git_lab`, `ibm_cloud_catalog`, `external_scm`.
* `state` - (Optional, List) Computed state of the Action.
Nested scheme for **state**:
	* `status_code` - (Optional, String) Status of automation (workspace or action).
	  * Constraints: Allowable values are: `normal`, `pending`, `disabled`, `critical`.
	* `status_job_id` - (Optional, String) Job id reference for this status.
	* `status_message` - (Optional, String) Automation status message - to be displayed along with the status_code.
* `sys_lock` - (Optional, List) System lock status.
Nested scheme for **sys_lock**:
	* `sys_locked` - (Optional, Boolean) Is the automation locked by a Schematic job ?.
	* `sys_locked_at` - (Optional, String) When the User performed the job that lead to locking of the automation ?.
	* `sys_locked_by` - (Optional, String) Name of the User who performed the job, that lead to the locking of the automation.
* `tags` - (Optional, List) Action tags.
* `targets_ini` - (Optional, String) Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost] 172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).
* `user_state` - (Optional, List) User defined status of the Schematics object.
Nested scheme for **user_state**:
	* `set_at` - (Optional, String) When the User who set the state of the Object.
	* `set_by` - (Optional, String) Name of the User who set the state of the Object.
	* `state` - (Optional, String) User-defined states  * `draft` Object can be modified; can be used by Jobs run by the author, during execution  * `live` Object can be modified; can be used by Jobs during execution  * `locked` Object cannot be modified; can be used by Jobs during execution  * `disable` Object can be modified. cannot be used by Jobs during execution.
	  * Constraints: Allowable values are: `draft`, `live`, `locked`, `disable`.
* `x_github_token` - (Optional, String) The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_action.
* `account` - (Optional, String) Action account ID.
* `created_at` - (Optional, String) Action creation time.
* `created_by` - (Optional, String) E-mail address of the user who created an action.
* `crn` - (Optional, String) Action Cloud Resource Name.
* `playbook_names` - (Optional, List) Playbook names retrieved from the repository.
* `source_created_at` - (Optional, String) Action Playbook Source creation time.
* `source_created_by` - (Optional, String) E-mail address of user who created the Action Playbook Source.
* `source_updated_at` - (Optional, String) The action playbook updation time.
* `source_updated_by` - (Optional, String) E-mail address of user who updated the action playbook source.
* `updated_at` - (Optional, String) Action updation time.
* `updated_by` - (Optional, String) E-mail address of the user who updated an action.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_schematics_action` resource by using `id`. Action ID.

# Syntax
```
$ terraform import ibm_schematics_action.schematics_action <id>
```
