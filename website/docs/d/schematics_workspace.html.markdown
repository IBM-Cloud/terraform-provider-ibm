---

subcategory: "Schematics"
layout: "ibm"
page_title: "IBM: ibm_schematics_workspace"
sidebar_current: "docs-ibm-datasource-schematics-workspace"
description: |-
  Get information about Schematics workspace.
---

# ibm_schematics_workspace
Retrieve information about a Schematics workspace. For more details about the Schematics and Schematics workspace, see [setting up workspaces](https://cloud.ibm.com/docs/schematics?topic=schematics-getting-started).

## Example usage

```terraform
data "ibm_schematics_workspace" "schematics_workspace" {
	workspace_id = "workspace_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `workspace_id` - (Required, Forces new resource, String) The ID of the workspace.  To find the workspace ID, use the `GET /v1/workspaces` API.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `applied_shareddata_ids` - (Optional, List) List of applied shared dataset ID.

* `catalog_ref` - (Optional, List) Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only.
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

* `created_at` - (Optional, String) The timestamp when the workspace was created.

* `created_by` - (Optional, String) The user ID that created the workspace.

* `crn` - (Optional, String) The workspace CRN.

* `description` - (Optional, String) The description of the workspace.

* `id` - (Optional, String) The unique identifier of the workspace.

* `last_health_check_at` - (Optional, String) The timestamp when the last health check was performed by Schematics.

* `location` - (Optional, String) The IBM Cloud location where your workspace was provisioned.

* `name` - (Optional, String) The name of the workspace.

* `resource_group` - (Optional, String) The resource group the workspace was provisioned in.

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

* `shared_data` - (Optional, List) Information about the Target used by the templates originating from IBM Cloud catalog offerings. This information is not relevant when you create a workspace from your own Terraform template.
Nested scheme for **shared_data**:
	* `cluster_id` - (Optional, String) The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	* `cluster_name` - (Optional, String) Target cluster name.
	* `entitlement_keys` - (Optional, List) The entitlement key that you want to use to install IBM Cloud entitled software.
	* `namespace` - (Optional, String) The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.
	* `region` - (Optional, String) The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	* `resource_group_id` - (Optional, String) The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.

* `status` - (Optional, String) The status of the workspace.   **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.   **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.   **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.   **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.   **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply jobs to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.   **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.   **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.   **Stopped**: The Schematics plan, apply, or destroy job was cancelled manually.   **Template Error**: The Schematics template contains errors and cannot be processed.

* `tags` - (Optional, List) A list of tags that are associated with the workspace.

* `template_env_settings` - (Optional, List) List of environment values.
Nested scheme for **env_values**:
	* `hidden` - (Optional, Boolean) Environment variable is hidden.
	* `name` - (Optional, String) Environment variable name.
	* `secure` - (Optional, Boolean) Environment variable is secure.
	* `value` - (Optional, String) Value for environment variable.

* `template_git_folder` - (Optional, String) The subfolder in your GitHub or GitLab repository where your Terraform template is stored. If your template is stored in the root directory, `.` is returned.

* `template_type` - (Optional, String) The Terraform version that was used to run your Terraform code.

* `template_uninstall_script_name` - (Optional, String) Uninstall script name.

* `template_values` - (Optional, String) A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `"autoscaling: enabled: true minReplicas: 2"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.

* `template_values_metadata` - (Optional, List) A list of input variables that are associated with the workspace.

* `template_inputs` - (Optional, List) Information about the input variables that your template uses.
Nested scheme for **variablestore**:
	* `description` - (Optional, String) The description of your input variable.
	* `name` - (Optional, String) The name of the variable.
	* `secure` - (Optional, Boolean) If set to `true`, the value of your input variable is protected and not returned in your API response.
	* `type` - (Optional, String) `Terraform v0.11` supports `string`, `list`, `map` data type. For more information, about the syntax, see [Configuring input variables](https://www.terraform.io/docs/configuration-0-11/variables.html).<br> `Terraform v0.12` additionally, supports `bool`, `number` and complex data types such as `list(type)`, `map(type)`,`object({attribute name=type,..})`, `set(type)`, `tuple([type])`. For more information, about the syntax to use the complex data type, see [Configuring variables](https://www.terraform.io/docs/configuration/variables.html#type-constraints).
	* `value` - (Optional, String) Enter the value as a string for the primitive types such as `bool`, `number`, `string`, and `HCL` format for the complex variables, as you provide in a `.tfvars` file. **You need to enter escaped string of `HCL` format for the complex variable value**. For more information, about how to declare variables in a terraform configuration file and provide value to schematics, see [Providing values for the declared variables](https://cloud.ibm.com/docs/schematics?topic=schematics-create-tf-config#declare-variable).

* `template_ref` - (Optional, String) Workspace template reference.

* `template_git_branch` - (Optional, String) The repository branch.

* `template_git_full_url` - (Optional, String) Full repository URL.

* `template_git_has_uploadedgitrepotar` - (Optional, Boolean) Has uploaded Git repository tar.

* `template_git_release` - (Optional, String) The repository release.

* `template_git_repo_sha_value` - (Optional, String) The repository SHA value.

* `template_git_repo_url` - (Optional, String) The repository URL.

* `template_git_url` - (Optional, String) The source URL.

* `type` - (Optional, List) The Terraform version that was used to run your Terraform code.

* `updated_at` - (Optional, String) The timestamp when the workspace was last updated.

* `updated_by` - (Optional, String) The user ID that updated the workspace.

* `frozen` - (Optional, Boolean) If set to true, the workspace is frozen and changes to the workspace are disabled.
	
* `frozen_at` - (Optional, String) The timestamp when the workspace was frozen.
	
* `frozen_by` - (Optional, String) The user ID that froze the workspace.
	
* `locked` - (Optional, Boolean) If set to true, the workspace is locked and disabled for changes.

* `locked_by` - (Optional, String) The user ID that initiated a resource-related job, such as applying or destroying resources, that locked the workspace.

* `locked_time` - (Optional, String) The timestamp when the workspace was locked.

* `status_code` - (Optional, String) The success or error code that was returned for the last plan, apply, or destroy job that ran against your workspace.

* `status_msg` - (Optional, String) The success or error message that was returned for the last plan, apply, or destroy job that ran against your workspace.