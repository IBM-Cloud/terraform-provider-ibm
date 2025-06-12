---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_workspace"
sidebar_current: "docs-ibm-resource-schematics-workspace"
description: |-
  Manages Schematics workspace.
---

# ibm_schematics_workspace
Create, read, update, and delete operations of Schematics workspace. Other Schematics operations such as plan, apply, destroy that are available through Schematics console are currently not supported. For more information, about IBM Cloud Schematics workspace, refer to [setting up workspaces](https://cloud.ibm.com/docs/schematics?topic=schematics-workspace-setup).


## Example usage

```terraform
resource "ibm_schematics_workspace" "schematics_workspace" {
  name = "<workspace_name>"
  description = "<workspace_description>"
  location = "us-east"
  resource_group = "default"
  template_type = "terraform_v0.13.5"
}
```


## Argument reference

Review the argument reference that you can specify for your resource.

* `applied_shareddata_ids` - (Optional, List) List of applied shared dataset ID.
* `catalog_ref` - (Optional, List) Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only. MaxItems:1.
Nested scheme for **catalog_ref**:
	* `dry_run` - (Optional, Boolean) Dry run.
	* `owning_account` - (Optional, String) Owning account ID of the catalog.
	* `item_icon_url` - (Optional, String) The URL to the icon of the software template in the IBM Cloud catalog.
	* `item_id` - (Optional, String) The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics.
	* `item_name` - (Optional, String) The name of the software that you chose to install from the IBM Cloud catalog.
	* `item_readme_url` - (Optional, String) The URL to the readme file of the software template in the IBM Cloud catalog.
	* `item_url` - (Optional, String) The URL to the software template in the IBM Cloud catalog.
	* `launch_url` - (Optional, String) The URL to the dashboard to access your software.
	* `offering_version` - (Optional, String) The version of the software template that you chose to install from the IBM Cloud catalog.
* `description` - (Optional, String) The description of the workspace.
* `location` - (Optional, String) The location where you want to create your Schematics workspace and run the Schematics jobs. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability.
* `name` - (Required, String) The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name.
* `resource_group` - (Optional, String) The ID of the resource group where you want to provision the workspace.
* `shared_data` - (Optional, List) Information about the Target used by the templates originating from the  IBM Cloud catalog offerings. This information is not relevant for workspace created using your own Terraform template. MaxItems:1.
Nested scheme for **shared_data**:
	* `cluster_created_on` - (Optional, String) Cluster created on.
	* `cluster_id` - (Optional, String) The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	* `cluster_name` - (Optional, String) The cluster name.
	* `cluster_type` - (Optional, String) The cluster type.
	* `entitlement_keys` - (Optional, List) The entitlement key that you want to use to install IBM Cloud entitled software.
	* `namespace` - (Optional, String) The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.
	* `region` - (Optional, String) The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	* `resource_group_id` - (Optional, String) The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	* `worker_count` - (Optional, Integer) The cluster worker count.
	* `worker_machine_type` - (Optional, String) The cluster worker type.
* `tags` - (Optional, List) A list of tags that are associated with the workspace.
* `template_env_settings` - (Optional, List) A list of environment variables that you want to apply during the execution of a bash script or Terraform job. This field must be provided as a list of key-value pairs, for example, **TF_LOG=debug**. Each entry will be a map with one entry where `key is the environment variable name and value is value`. You can define environment variables for IBM Cloud catalog offerings that are provisioned by using a bash script. See [example to use special environment variable](https://cloud.ibm.com/docs/schematics?topic=schematics-set-parallelism#parallelism-example)  that are supported by Schematics.
* `template_git_folder` - (Optional, String) The subfolder in your GitHub or GitLab repository where your Terraform template is stored.
* `template_init_state_file` - (Optional, String) The content of an existing Terraform statefile that you want to import in to your workspace. To get the content of a Terraform statefile for a specific Terraform template in an existing workspace, run `ibmcloud schematics state pull --id <workspace_id> --template <template_id>`.
* `template_type` - (Required, String) The Terraform version that you want to use to run your Terraform code. Enter `terraform_v0.12` to use Terraform version 0.12, and `terraform_v0.11` to use Terraform version 0.11. The Terraform config files are run with Terraform version 0.11. This is a required variable. Make sure that your Terraform config files are compatible with the Terraform version that you select. See [terraform versions](https://cloud.ibm.com/docs/schematics?topic=schematics-migrating-terraform-version) that are supported by Schematics.
* `template_uninstall_script_name` - (Optional, String) Uninstall script name.
* `template_values` - (Optional, String) A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `"autoscaling: enabled: true minReplicas: 2"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.
* `template_values_metadata` - (Optional, List) List of values metadata.
Nested scheme for **template_values_metadata**:
	* `aliases` - (Optional, List) The list of aliases for the variable name.
	* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
	* `default` - (Optional, String) Default value for the variable only if the override value is not specified.
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
	* `name` - (String) Name of the variable.
	* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
	* `position` - (Optional, Integer) The relative position of this variable in a list.
	* `required` - (Optional, Boolean) If the variable required?.
	* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
	* `source` - (Optional, String) The source of this meta-data.
	* `type` - (Optional, String) Type of the variable.
		* Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
* `template_inputs` - (Optional, List) VariablesRequest -.
Nested scheme for **variablestore**:
	* `description` - (Optional, String) The description of your input variable.
	* `name` - (Required, String) The name of the variable.
	* `secure` - (Optional, Boolean) If set to `true`, the value of your input variable is protected and not returned in your API response.
	* `type` - (Required, String) `Terraform v0.11` supports `string`, `list`, `map` data type. For more information, about the syntax, see [Configuring input variables](https://www.terraform.io/docs/configuration-0-11/variables.html).<br> `Terraform v0.12` additionally, supports `bool`, `number` and complex data types such as `list(type)`, `map(type)`,`object({attribute name=type,..})`, `set(type)`, `tuple([type])`. For more information, about the syntax to use the complex data type, see [Configuring variables](https://www.terraform.io/docs/configuration/variables.html#type-constraints).
	* `use_default` - (Optional, Boolean) Variable uses default value; and is not over-ridden.
	* `value` - (Required, String) Enter the value as a string for the primitive types such as `bool`, `number`, `string`, and `HCL` format for the complex variables, as you provide in a `.tfvars` file. **You need to enter escaped string of `HCL` format for the complex variable value**. For more information, about how to declare variables in a terraform configuration file and provide value to schematics, see [Providing values for the declared variables](https://cloud.ibm.com/docs/schematics?topic=schematics-create-tf-config#declare-variable).
* `template_ref` - (Optional, String) Workspace template ref.
* `template_git_branch` - (Optional, String) The repository branch.
* `template_git_release` - (Optional, String) The repository release.
* `template_git_repo_sha_value` - (Optional, String) The repository SHA value.
* `template_git_repo_url` - (Optional, String) The repository URL.
* `template_git_url` - (Optional, String) The source URL.
* `frozen` - (Optional, Boolean) If set to true, the workspace is frozen and changes to the workspace are disabled.
* `frozen_at` - (Optional, String) The timestamp when the workspace was frozen.
* `frozen_by` - (Optional, String) The user ID that froze the workspace.
* `locked` - (Optional, Boolean) If set to true, the workspace is locked and disabled for changes.
* `locked_by` - (Optional, String) The user ID that initiated a resource-related job, such as applying or destroying resources, that locked the workspace.
* `locked_time` - (Optional, String) The timestamp when the workspace was locked.
* `x_github_token` - (Optional, String) The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_workspace.
* `created_at` - (String) The timestamp when the workspace was created.
* `created_by` - (String) The user ID that created the workspace.
* `crn` - (Optional, String) The workspace CRN.
* `last_health_check_at` - (String) The timestamp when the last health check was performed by Schematics.
* `runtime_data` - (Optional, List) Information about the provisioning engine, state file, and runtime logs.
Nested scheme for **runtime_data**:
	* `engine_cmd` - (Optional, String) The command that was used to apply the Terraform template or IBM Cloud catalog software template.
	* `engine_name` - (Optional, String) The provisioning engine that was used to apply the Terraform template or IBM Cloud catalog software template.
	* `engine_version` - (Optional, String) The version of the provisioning engine that was used.
	* `id` - (Optional, String) The ID that was assigned to your Terraform template or IBM Cloud catalog software template.
	* `log_store_url` - (Optional, String) The URL to access the logs that were created during the creation, update, or deletion of your IBM Cloud resources.
	* `output_values` - (Optional, List) List of Output values.
	* `resources` - (Optional, List) List of resources.
	* `state_store_url` - (Optional, String) The URL where the Terraform statefile (`terraform.tfstate`) is stored. You can use the statefile to find an overview of IBM Cloud resources that were created by Schematics. Schematics uses the statefile as an inventory list to determine future create, update, or deletion jobs.
* `status` - (String) The status of the workspace.   **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.   **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.   **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.   **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.   **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply jobs to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.   **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.   **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.   **Stopped**: The Schematics plan, apply, or destroy job was cancelled manually.   **Template Error**: The Schematics template contains errors and cannot be processed.
* `updated_at` - (String) The timestamp when the workspace was last updated.
* `updated_by` - (String) The user ID that updated the workspace.
* `status_code` - (String) The success or error code that was returned for the last plan, apply, or destroy job that ran against your workspace.
* `status_msg` - (String) The success or error message that was returned for the last plan, apply, or destroy job that ran against your workspace.

## Import

You can import the `ibm_schematics_workspace` resource by using `id`. The unique identifier of the workspace.

# Syntax
```
$ terraform import ibm_schematics_workspace.schematics_workspace <id>
```
**Note**

This resource can perform create, read, update and delete operations of schematics workspace. Other schematics operations such as `Plan` and `Apply` which are available in IBM Cloud Schematics console are currently not supported.


