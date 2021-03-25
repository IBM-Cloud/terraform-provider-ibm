---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_workspace"
sidebar_current: "docs-ibm-resource-schematics-workspace"
description: |-
  Manages schematics_workspace.
---

# ibm\_schematics_workspace

Provides a resource for ibm_schematics_workspace. This allows ibm_schematics_workspace to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_workspace" "schematics_workspace" {
  name = "<workspace_name>"
  description = "<workspace_description>"
  location = "us-east"
  resource_group = "default"
  template_type = "terraform_v0.13.5"
}
```

## Argument Reference

The following arguments are supported:

* `applied_shareddata_ids` - (Optional, List) List of applied shared dataset id.
* `catalog_ref` - (Optional, List) Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only.
  * `dry_run` - (Optional, bool) Dry run.
  * `item_icon_url` - (Optional, string) The URL to the icon of the software template in the IBM Cloud catalog.
  * `item_id` - (Optional, string) The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics.
  * `item_name` - (Optional, string) The name of the software that you chose to install from the IBM Cloud catalog.
  * `item_readme_url` - (Optional, string) The URL to the readme file of the software template in the IBM Cloud catalog.
  * `item_url` - (Optional, string) The URL to the software template in the IBM Cloud catalog.
  * `launch_url` - (Optional, string) The URL to the dashboard to access your software.
  * `offering_version` - (Optional, string) The version of the software template that you chose to install from the IBM Cloud catalog.
* `description` - (Optional, string) The description of the workspace.
* `location` - (Optional, string) The location where you want to create your Schematics workspace and run Schematics actions. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability.
* `name` - (Optional, string) The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name.
* `resource_group` - (Optional, string) The ID of the resource group where you want to provision the workspace.
* `shared_data` - (Optional, List) Information that is shared across templates in IBM Cloud catalog offerings. This information is not provided when you create a workspace from your own Terraform template.
  * `cluster_created_on` - (Optional, string) Cluster created on.
  * `cluster_id` - (Optional, string) The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.
  * `cluster_name` - (Optional, string) Cluster name.
  * `cluster_type` - (Optional, string) Cluster type.
  * `entitlement_keys` - (Optional, []interface{}) The entitlement key that you want to use to install IBM Cloud entitled software.
  * `namespace` - (Optional, string) The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.
  * `region` - (Optional, string) The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
  * `resource_group_id` - (Optional, string) The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
  * `worker_count` - (Optional, int) Cluster worker count.
  * `worker_machine_type` - (Optional, string) Cluster worker type.
* `tags` - (Optional, List) A list of tags that are associated with the workspace.
* `template_data` - (Optional, List) TemplateData -.
  * `env_values` - (Optional, []interface{}) A list of environment variables that you want to apply during the execution of a bash script or Terraform action. This field must be provided as a list of key-value pairs, for example, **TF_LOG=debug**. Each entry will be a map with one entry where `key is the environment variable name and value is value`. You can define environment variables for IBM Cloud catalog offerings that are provisioned by using a bash script.
  * `folder` - (Optional, string) The subfolder in your GitHub or GitLab repository where your Terraform template is stored.
  * `init_state_file` - (Optional, string) The content of an existing Terraform statefile that you want to import in to your workspace. To get the content of a Terraform statefile for a specific Terraform template in an existing workspace, run `ibmcloud terraform state pull --id <workspace_id> --template <template_id>`.
  * `type` - (Optional, string) The Terraform version that you want to use to run your Terraform code. Enter `terraform_v0.12` to use Terraform version 0.12, and `terraform_v0.11` to use Terraform version 0.11. If no value is specified, the Terraform config files are run with Terraform version 0.11. Make sure that your Terraform config files are compatible with the Terraform version that you select.
  * `uninstall_script_name` - (Optional, string) Uninstall script name.
  * `values` - (Optional, string) A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `"autoscaling:  enabled: true  minReplicas: 2"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.
  * `values_metadata` - (Optional, []interface{}) List of values metadata.
  * `variablestore` - (Optional, []interface{}) VariablesRequest -.
* `template_ref` - (Optional, string) Workspace template ref.
* `template_repo` - (Optional, List) Input parameter to specify the source repository where your Schematics template is stored.
  * `branch` - (Optional, string) The branch in GitHub where your Terraform template is stored.
  * `release` - (Optional, string) The release tag in GitHub of your Terraform template.
  * `repo_sha_value` - (Optional, string) Repo SHA value.
  * `repo_url` - (Optional, string) The URL to the repository where the IBM Cloud catalog software template is stored.
  * `url` - (Optional, string) The URL to the GitHub or GitLab repository where your Terraform and public bit bucket template is stored. For more information of the environment variable syntax, see [Create workspace new](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-workspace-new).
* `type` - (Optional, List) The Terraform version that you want to use to run your Terraform code. Enter `terraform_v0.12` to use Terraform version 0.12, and `terraform_v0.11` to use Terraform version 0.11. If no value is specified, the Terraform config files are run with Terraform version 0.11. Make sure that your Terraform config files are compatible with the Terraform version that you select.
* `workspace_status` - (Optional, List) WorkspaceStatusRequest -.
  * `frozen` - (Optional, bool) If set to true, the workspace is frozen and changes to the workspace are disabled.
  * `frozen_at` - (Optional, TypeString) The timestamp when the workspace was frozen.
  * `frozen_by` - (Optional, string) The user ID that froze the workspace.
  * `locked` - (Optional, bool) If set to true, the workspace is locked and disabled for changes.
  * `locked_by` - (Optional, string) The user ID that initiated a resource-related action, such as applying or destroying resources, that locked the workspace.
  * `locked_time` - (Optional, TypeString) The timestamp when the workspace was locked.
* `x_github_token` - (Optional, string) The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the schematics_workspace.
* `created_at` - The timestamp when the workspace was created.
* `created_by` - The user ID that created the workspace.
* `crn` - Workspace CRN.
* `last_health_check_at` - The timestamp when the last health check was performed by Schematics.
* `runtime_data` - Information about the provisioning engine, state file, and runtime logs.
* `status` - The status of the workspace.  **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.  **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.  **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.  **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.  **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply actions to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.  **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.  **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.  **Stopped**: The Schematics plan, apply, or destroy action was cancelled manually.  **Template Error**: The Schematics template contains errors and cannot be processed.
* `updated_at` - The timestamp when the workspace was last updated.
* `updated_by` - The user ID that updated the workspace.
* `workspace_status_msg` - Information about the last action that ran against the workspace.

## Note
*  Please note that this resource can be only used to perform Create, Read, Update and Delete operations of schematics workspace. Other schematics operations such as `Plan` and `Apply` which are available via IBM Cloud Schematics Console are currently not supported.
