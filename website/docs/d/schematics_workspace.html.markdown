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

## Argument reference
Review the argument references that you can specify for your data source. 

- `workspace_id` - (Required, String) The workspace ID that you want to retrieve. To find the workspace ID, use the `GET /v1/workspaces API`.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `applied_shareddata_ids` - (String) List of applied shared data set ID.
- `catalog_ref`-  (String) Information about the software template that you select from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only. Nested `catalog_ref` blocks have the following structure.

  Nested scheme for `catalog_ref`:
  - `dry_run`- Bool - Dry run.
  - `item_icon_url`- (String) The icon URL of the software template in the IBM Cloud catalog.
  - `item_id`-  (String) The ID of the software template that you select to install from the IBM Cloud catalog. This software is provisioned with Schematics.
  - `item_name` - (String) The name of the software that you select to install from the IBM Cloud catalog.
  - `item_readme_url`-  (String) The URL to the readme file of the software template in the IBM Cloud catalog.
  - `item_url` - (String) The URL to the software template in the IBM Cloud catalog.
  - `launch_url` - (String) The dashboard URL to access the software.
  - `offering_version` - (String) The version of the software template that you select to install from the IBM Cloud catalog.
- `created_at` - (Timestamp) The workspace created timestamp.
- `created_by` - (String) The user ID that created the workspace.
- `crn` - (String) The workspace CRN.
- `description` - (String) The description of the workspace.
- `id` - (String) The unique ID of the Schematics workspace.
- `last_health_check_at` - (Timestampe) The timestamp when the last health check was performed by Schematics.
- `location` - (String) The IBM Cloud location where your workspace was provisioned.
- `name` - (String) The name of your workspace.
- `resource_group` - (String) The ID of the resource group where you want to provision the workspace.
- `runtime_data` - (String) Information about provisioning engine, state file, and runtime logs. Nested runtime_data blocks have the following structure.

  Nested scheme for `runtime_data`:
  - `engine_cmd`-  (String) The command used to apply the Terraform template or IBM Cloud catalog software template.
  - `engine_name`-  (String) The provisioning engine used to apply the Terraform template or IBM Cloud catalog software template.
  - `engine_version`-  (String) The version of the provisioning engine.
  - `id`-  (String) The ID that was assigned to your Terraform template or IBM Cloud catalog software template.
  - `log_store_url`-  (String) The URL to access the logs created during the creation, update, or deletion of your IBM Cloud resources.
  - `output_values`-  (String) The list of output values.
  - `resources`-  (String) The list of resources.
  - `state_store_url`-  (String) The URL where the Terraform statefile `terraform.tfstate` is stored. You can use the statefile to find an overview of IBM Cloud  resources created by Schematics. Schematics uses the statefile as an inventory list to determine future create, update, or deletion actions.
- `shared_data` - (String) Information that is shared across templates in IBM Cloud catalog offerings. This information is not provided when you create a workspace from your own Terraform template. Nested `shared_data` blocks have the following structure.

  Nested scheme for `shared_data`:
  - `cluster_id` - (String) The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.
  - `cluster_name` - (String) The target cluster name.
  - `entitlement_keys` - `String` - The entitlement key that you want to use to install IBM Cloud entitled software.
  - `namespace` - (String) The Kubernetes Service namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.
  - `region` - (String) The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
  - `resource_group_id` - (String) The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
- `status` - (String) The status of the workspace. For more information, about the Schematics workspace status, see [workspace states](https://cloud.ibm.com/docs/schematics?topic=schematics-workspace-setup#wks-state).
- `tags`- (List) A list of tags that are associated with the workspace.
- `template_data`- (List) Information about the Terraform or IBM Cloud software template to use. Nested `template_data` blocks have the following structure:

  Nested scheme for `template_data`:
  - `env_values` -  (String) List of environment values. Nested `env_values` blocks have the following structure:

    Nested scheme for `env_values`:
    - `hidden`-  (String) The environment variable is hidden.
    - `name`-  (String) The environment variable name.
    - `secure`-  (String) The environment variable is secure.
    - `value`-  (String)  Value for an environment variable.
  - `folder` - (String) The subfolder in your GitHub or GitLab repository where your Terraform template is stored.
  - `has_githubtoken` - (String) Has GitHub token.
  - `id` - (String) The ID assigned to your Terraform template or IBM Cloud catalog software template.
  - `template_type` - (String) The Terraform version used to run your Terraform code.
  - `uninstall_script_name` - (String) The script name to uninstall.
  - `values` - (String) A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `autoscaling: enabled: true minReplicas: 2`. The values that you define here overrides the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.
  - `values_metadata` - (String) A list of input variables that are associated with the workspace.
  - `values_url` - (String) The API endpoint to access the input variables that you defined for your template.
  - `variablestore` - (String) Information about the input variables that your template uses. Nested variable store blocks have the following structure.

    Nested scheme for `variablestore`:
    - `description` - (String) The description of your input variable.
    - `name` - (String) The name of your variable.
	- `secure` - (String) If set to **true**, the value of your input variable is protected and not returned in your API response.
	- `type` - (String) Terraform v0.11 supports string, list, map data type. For more information, about the syntax, see [configuring input variables](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-workspace-new). Terraform v0.12 additionally, supports bool, number and complex data types such as list(type), map(type), object({attribute name=type,.}), set(type), tuple([type]). For more information, about the syntax to use the complex data type, see [configuring variables](https://cloud.ibm.com/docs/schematics?topic=schematics-create-tf-config#configure-variables).
	- `value` - (String) Enter the value as a string for the primitive types such as bool, number, string, and HCL format for the complex variables, as you provide in a `.tfvars` file. You need to enter escaped string of HCL format for the complex variable value. For more information, about how to declare variables in a Terraform configuration file and provide value to schematics, see [providing values for the declared variables](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-workspace-new).
- `template_ref` - (String) The workspace template reference.
- `template_repo` - (List)  The input parameter to specify the source repository where your Schematics template is stored.

  Nested scheme for `template_repo`:
  - `branch` - (String) The branch in GitHub where your Terraform template is stored.
  - `full_url` - (String) The full URL repository.
  - `has_uploadedgitrepotar` - (String) Has uploaded the Git repository tap archive file.
  - `release` - (String) The release tag in GitHub of your Terraform template.
  - `repo_sha_value` - (String) The SHA value from the repository.
  - `repo_url` - (String) The URL to the repository where the IBM Cloud catalog software template is stored.
  - `url` - (String) The GitHub or GitLab repository URL where your Terraform template is stored.
  - `type`- (List) The Terraform version that you want to use to run your Terraform code.
  - `updated_at`- (Timestamp) The timestamp when the workspace updated.
  - `updated_by`- (Timestamp) The user ID of the workspace updated.
- `workspace_status`- (List) Response parameter that indicate if a workspace is frozen or locked. Nested `workspace_status` blocks have the following structure.

  Nested scheme for `workspace_status`:
  - `frozen`- (Optional, Bool) If set to **true**, the workspace is frozen and changes to the workspace are disabled.
  - `frozen_at` - (String) The timestamp when the workspace was frozen.
  - `frozen_by` - (String) The user ID that froze the workspace.
  - `locked` - (Bool) If set to **true**, the workspace is locked and disabled for changes.
  - `locked_by` - (String) The workspace locked for the user who initiates a resource-related action, such as applying or destroying resources.
  - `locked_time` - (Timestamp) The timestamp when the workspace is locked.
- `workspace_status_msg`-  (String) Information about the last action that ran against the workspace. Nested `workspace_status_msg` blocks have the following structure.

  Nested scheme for `workspace_status_msg`:
  - `status_code`-  (String) The success or error code that was returned for the last plan, apply, or destroy action that ran against your workspace.
  - `status_msg`-  (String) The success or error message returned for the last plan, apply, or destroy action of your workspace.
