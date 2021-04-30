---

subcategory: "Schematics"
layout: "ibm"
page_title: "IBM: ibm_schematics_workspace"
sidebar_current: "docs-ibm-datasource-schematics-workspace"
description: |-
  Get information about schematics_workspace
---

# ibm\_schematics_workspace

Provides a read-only data source for ibm_schematics_workspace. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_workspace" "schematics_workspace" {
	workspace_id = "workspace_id"
}
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required, string) The ID of the workspace for which you want to retrieve detailed information. To find the workspace ID, use the `GET /v1/workspaces` API.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the schematics_workspace.
* `applied_shareddata_ids` - List of applied shared dataset id.

* `catalog_ref` - Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only. Nested `catalog_ref` blocks have the following structure:
	* `dry_run` - Dry run.
	* `item_icon_url` - The URL to the icon of the software template in the IBM Cloud catalog.
	* `item_id` - The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics.
	* `item_name` - The name of the software that you chose to install from the IBM Cloud catalog.
	* `item_readme_url` - The URL to the readme file of the software template in the IBM Cloud catalog.
	* `item_url` - The URL to the software template in the IBM Cloud catalog.
	* `launch_url` - The URL to the dashboard to access your software.
	* `offering_version` - The version of the software template that you chose to install from the IBM Cloud catalog.

* `created_at` - The timestamp when the workspace was created.

* `created_by` - The user ID that created the workspace.

* `crn` - Workspace CRN.

* `description` - The description of the workspace.

* `id` - The unique identifier of the workspace.

* `last_health_check_at` - The timestamp when the last health check was performed by Schematics.

* `location` - The IBM Cloud location where your workspace was provisioned.

* `name` - The name of the workspace.

* `resource_group` - The resource group the workspace was provisioned in.

* `runtime_data` - Information about the provisioning engine, state file, and runtime logs. Nested `runtime_data` blocks have the following structure:
	* `engine_cmd` - The command that was used to apply the Terraform template or IBM Cloud catalog software template.
	* `engine_name` - The provisioning engine that was used to apply the Terraform template or IBM Cloud catalog software template.
	* `engine_version` - The version of the provisioning engine that was used.
	* `id` - The ID that was assigned to your Terraform template or IBM Cloud catalog software template.
	* `log_store_url` - The URL to access the logs that were created during the creation, update, or deletion of your IBM Cloud resources.
	* `output_values` - List of Output values.
	* `resources` - List of resources.
	* `state_store_url` - The URL where the Terraform statefile (`terraform.tfstate`) is stored. You can use the statefile to find an overview of IBM Cloud resources that were created by Schematics. Schematics uses the statefile as an inventory list to determine future create, update, or deletion actions.

* `shared_data` - Information that is shared across templates in IBM Cloud catalog offerings. This information is not provided when you create a workspace from your own Terraform template. Nested `shared_data` blocks have the following structure:
	* `cluster_id` - The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	* `cluster_name` - Target cluster name.
	* `entitlement_keys` - The entitlement key that you want to use to install IBM Cloud entitled software.
	* `namespace` - The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.
	* `region` - The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	* `resource_group_id` - The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.

* `status` - The status of the workspace.  **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.  **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.  **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.  **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.  **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply actions to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.  **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.  **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.  **Stopped**: The Schematics plan, apply, or destroy action was cancelled manually.  **Template Error**: The Schematics template contains errors and cannot be processed.

* `tags` - A list of tags that are associated with the workspace.

* `template_data` - Information about the Terraform or IBM Cloud software template that you want to use. Nested `template_data` blocks have the following structure:
	* `env_values` - List of environment values. Nested `env_values` blocks have the following structure:
		* `hidden` - Environment variable is hidden.
		* `name` - Environment variable name.
		* `secure` - Environment variable is secure.
		* `value` - Value for environment variable.
	* `folder` - The subfolder in your GitHub or GitLab repository where your Terraform template is stored. If your template is stored in the root directory, `.` is returned.
	* `has_githubtoken` - Has github token.
	* `id` - The ID that was assigned to your Terraform template or IBM Cloud catalog software template.
	* `template_type` - The Terraform version that was used to run your Terraform code.
	* `uninstall_script_name` - Uninstall script name.
	* `values` - A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `""autoscaling:  enabled: true  minReplicas: 2"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.
	* `values_metadata` - A list of input variables that are associated with the workspace.
	* `values_url` - The API endpoint to access the input variables that you defined for your template.
	* `variablestore` - Information about the input variables that your template uses. Nested `variablestore` blocks have the following structure:
		* `description` - The description of your input variable.
		* `name` - The name of the variable.
		* `secure` - If set to `true`, the value of your input variable is protected and not returned in your API response.
		* `type` - `Terraform v0.11` supports `string`, `list`, `map` data type. For more information, about the syntax, see [Configuring input variables](https://www.terraform.io/docs/configuration-0-11/variables.html). <br> `Terraform v0.12` additionally, supports `bool`, `number` and complex data types such as `list(type)`, `map(type)`, `object({attribute name=type,..})`, `set(type)`, `tuple([type])`. For more information, about the syntax to use the complex data type, see [Configuring variables](https://www.terraform.io/docs/configuration/variables.html#type-constraints).
		* `value` - Enter the value as a string for the primitive types such as `bool`, `number`, `string`, and `HCL` format for the complex variables, as you provide in a `.tfvars` file. **You need to enter escaped string of `HCL` format for the complex variable value**. For more information, about how to declare variables in a terraform configuration file and provide value to schematics, see [Providing values for the declared variables](/docs/schematics?topic=schematics-create-tf-config#declare-variable).

* `template_ref` - Workspace template ref.

* `template_repo` - Information about the Terraform template that your workspace points to. Nested `template_repo` blocks have the following structure:
	* `branch` - The branch in GitHub where your Terraform template is stored.
	* `full_url` - Full repo URL.
	* `has_uploadedgitrepotar` - Has uploaded git repo tar.
	* `release` - The release tag in GitHub of your Terraform template.
	* `repo_sha_value` - Repo SHA value.
	* `repo_url` - The URL to the repository where the IBM Cloud catalog software template is stored.
	* `url` - The URL to the GitHub or GitLab repository where your Terraform template is stored.

* `type` - The Terraform version that was used to run your Terraform code.

* `updated_at` - The timestamp when the workspace was last updated.

* `updated_by` - The user ID that updated the workspace.

* `workspace_status` - Response parameter that indicate if a workspace is frozen or locked. Nested `workspace_status` blocks have the following structure:
	* `frozen` - If set to true, the workspace is frozen and changes to the workspace are disabled.
	* `frozen_at` - The timestamp when the workspace was frozen.
	* `frozen_by` - The user ID that froze the workspace.
	* `locked` - If set to true, the workspace is locked and disabled for changes.
	* `locked_by` - The user ID that initiated a resource-related action, such as applying or destroying resources, that locked the workspace.
	* `locked_time` - The timestamp when the workspace was locked.

* `workspace_status_msg` - Information about the last action that ran against the workspace. Nested `workspace_status_msg` blocks have the following structure:
	* `status_code` - The success or error code that was returned for the last plan, apply, or destroy action that ran against your workspace.
	* `status_msg` - The success or error message that was returned for the last plan, apply, or destroy action that ran against your workspace.

