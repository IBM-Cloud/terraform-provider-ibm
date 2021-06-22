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
Review the argument references that you can specify for your resource. 

- `applied_shareddata_ids` - (Optional, List) List of applied shared data set ID.
- `catalog_ref` - (Optional, List) The software template that you select from the IBM Cloudcatalog and returned for IBM Cloud catalog offerings only.

  Nested scheme for `catalog_ref`:
  - `dry_run`- (Optional, Bool) Dry run.
  - `item_icon_url`- (Optional, String) The icon URL of the software template in the IBM Cloud catalog.
  - `item_id`- (Optional, String) The ID of the software template that you select to install from the IBM Cloud catalog. This software is provisioned with Schematics.
  - `item_name` - (Optional, String) The name of the software that you select to install from the IBM Cloud catalog.
  - `item_readme_url`- (Optional, String) The URL to the readme file of the software template in the IBM Cloud catalog.
  - `item_url` - (Optional, String) The URL to the software template in the IBM Cloud catalog.
  - `launch_url` - (Optional, String) The dashboard URL to access the software.
  - `offering_version` - (Optional, String) The version of the software template that you select to install from the IBM Cloud catalog.
- `description` - (Optional, String) The description of the workspace.
- `location` - (Optional, String) The location where you want to create your Schematics workspace and run Schematics actions. The location that you enter should match the API endpoint that you use. For example, if you use the `Frankfurt API endpoint`, you should specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability.
- `name` - (Optional, String) The name of your workspace. The name can be up to **128** characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name.
- `resource_group` - (Optional, String) The ID of the resource group where you want to provision the workspace.
- `shared_data` - (Optional, String)  Information that is shared across templates in IBM Cloud catalog offerings. This information is not provided when you create a workspace from your own Terraform template.

  Nested scheme for `shared_data`:
  - `cluster_created_on` - (Optional, String) The cluster created on.
  - `cluster_id` - (Optional, String) The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.
  - `cluster_name` - (Optional, String) The cluster name.
  - `cluster_type` - (Optional, String) The cluster type.
  - `entitlement_keys` - `[]interface{}]` - Optional - The entitlement key that you want to use to install IBM Cloud entitled software.
  - `namespace` - (Optional, String) The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.
  - `region` - (Optional, String) The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
  - `resource_group_id` - (Optional, String)   The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
  - `worker_count`  - (Optional, Integer)  Cluster worker count.
  - `worker_machine_type` - (Optional, String)  Cluster worker type.
- `tags` - (Optional, List) A list of tags that are associated with the workspace.
- `template_data` - (Optional, List) List of template data.

  Nested scheme for `template_data`:
  - `env_values`- (Optional, []interface{}) A list of environment variables that you want to apply during the execution of a bash script or Terraform action. This field contains the list of key-value pairs, for example, `TF_LOG=debug`. Each entry is map with one entry where key is the environment variable name and value is value. You can define environment variables for IBM Cloud catalog offerings that are provisioned by using a bash script.
  - `folder` - (Optional, String) The subfolder in your GitHub or GitLab repository where your Terraform template is stored.
  - `init_state_file` - (Optional, String) The content of an existing Terraform statefile that you want to import to your workspace. To get the content of a Terraform statefile for a specific Terraform template in an existing workspace, run `ibmcloud terraform state pull id <workspace_id> template <template_id>`.
  - `type` - (Optional, String) The Terraform version that you want to run your Terraform code. Enter `terraform_v0.12` to use Terraform version 0.12, and `terraform_v0.11` to use Terraform version 0.11. If value is not specified, the Terraform config files are run with Terraform version 0.11. Make sure that your Terraform config files are compatible with the Terraform version that you select.
  - `uninstall_script_name` - (Optional, String) The script name to uninstall.
  - `values` - (Optional, String)  A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `autoscaling: enabled: true minReplicas: 2`. The values that you define here overrides the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.
  - `values_metadata` - (Optional, []interface{})  The list of values metadata.
  - `variablestore` - (Optional, []interface{})  The variable request.
- `template_ref` - (Optional, String) The workspace template reference.
- `template_repo` - (Optional, List) The input parameter to specify the source repository where your Schematics template is stored.

  Nested scheme for `template_repo`:
  - `branch` - (Optional, String) The branch in GitHub where your Terraform template is stored.
  - `release` - (Optional, String) The release tag in GitHub of your Terraform template.
  - `repo_sha_value` - (Optional, String) The SHA value from the repository.
  - `repo_url` - (Optional, String) The URL to the repository where the IBM Cloud catalog software template is stored.
  - `url` - (Optional, String) The GitHub or GitLab repository URL where your Terraform and public bit bucket template is stored. For more information of the environment variable syntax, see [Create workspace new](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-workspace-new).
- `type` - (Optional, List) The Terraform version that you want to use to run your Terraform code. Enter terraform_v0.12 to use Terraform version 0.12, and terraform_v0.11 to use Terraform version 0.11. If no value is specified, the Terraform config files are run with Terraform version 0.11. Make sure that your Terraform config files are compatible with the Terraform version that you select.
- `workspace_status` - (Optional, List) The Workspace status request.

  Nested scheme for `workspace_status`:
  - `frozen`- (Optional, Bool) If set to **true**, the workspace is frozen and changes to the workspace are disabled.
  - `frozen_at` - (Optional, TypeString) The timestamp when the workspace was frozen.
  - `frozen_by` - (Optional, String) The user ID that froze the workspace.
  - `locked` - (Optional, Bool)  If set to **true**, the workspace is locked and disabled for changes.
  - `locked_by` - (Optional, String) The user ID that initiated a resource-related action, such as applying or destroying resources, that locked the workspace.
  - `locked_time` - (Optional, TypeString) The timestamp when the workspace was locked.
- `x_github_token` - (Optional, String) The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the Schematics workspace.
- `created_at` - (Timestamp) The timestamp when the workspace is created.
- `created_by ` - (String) The user ID that created the workspace.
- `crn` - (String) The workspace CRN.
- `last_health_check_at` - (String) The timestamp when the last health check is executed by Schematics.
- `runtime_data` - (String) The information about the provisioning engine, state file, and runtime logs.
- `status`-  (String) The status of the workspace. For more information, see [workspace status](https://cloud.ibm.com/docs/schematics?topic=schematics-workspace-setup#wks-state).
- `updated_at`-  (String) The timestamp when the workspace last updated.
- `updated_by`-  (String) The user ID that updated the workspace.
- `workspace_status_msg` -String - The last action information that workspace run.

**Note**

This resource can perform create, read, update and delete operations of schematics workspace. Other schematics operations such as `Plan` and `Apply` which are available in IBM Cloud Schematics console are currently not supported.


