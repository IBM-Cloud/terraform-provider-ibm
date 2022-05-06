---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_workspace"
sidebar_current: "docs-ibm-resource-schematics-workspace"
description: |-
  Manages schematics_workspace.
subcategory: "Schematics Service API"
---

# ibm_schematics_workspace

Provides a resource for schematics_workspace. This allows schematics_workspace to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_workspace" "schematics_workspace" {
  catalog_ref {
		dry_run = true
		owning_account = "owning_account"
		item_icon_url = "item_icon_url"
		item_id = "item_id"
		item_name = "item_name"
		item_readme_url = "item_readme_url"
		item_url = "item_url"
		launch_url = "launch_url"
		offering_version = "offering_version"
  }
  dependencies {
		parents = [ "parents" ]
		children = [ "children" ]
  }
  shared_data {
		cluster_created_on = "cluster_created_on"
		cluster_id = "cluster_id"
		cluster_name = "cluster_name"
		cluster_type = "cluster_type"
		entitlement_keys = [ null ]
		namespace = "namespace"
		region = "region"
		resource_group_id = "resource_group_id"
		worker_count = 1
		worker_machine_type = "worker_machine_type"
  }
  template_data {
		env_values = [ null ]
		env_values_metadata {
			hidden = true
			name = "name"
			secure = true
		}
		folder = "folder"
		compact = true
		init_state_file = "init_state_file"
		injectors {
			tft_git_url = "tft_git_url"
			tft_git_token = "tft_git_token"
			tft_prefix = "tft_prefix"
			injection_type = "injection_type"
			tft_name = "tft_name"
			tft_parameters {
				name = "name"
				value = "value"
			}
		}
		type = "type"
		uninstall_script_name = "uninstall_script_name"
		values = "values"
		values_metadata = [ null ]
		variablestore {
			description = "description"
			name = "name"
			secure = true
			type = "type"
			use_default = true
			value = "value"
		}
  }
  template_repo {
		branch = "branch"
		release = "release"
		repo_sha_value = "repo_sha_value"
		repo_url = "repo_url"
		url = "url"
  }
  workspace_status {
		frozen = true
		frozen_at = "2021-01-31T09:44:12Z"
		frozen_by = "frozen_by"
		locked = true
		locked_by = "locked_by"
		locked_time = "2021-01-31T09:44:12Z"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `applied_shareddata_ids` - (Optional, List) List of applied shared dataset ID.
* `catalog_ref` - (Optional, List) Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only.
Nested scheme for **catalog_ref**:
	* `dry_run` - (Optional, Boolean) Dry run.
	* `item_icon_url` - (Optional, String) The URL to the icon of the software template in the IBM Cloud catalog.
	* `item_id` - (Optional, String) The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics.
	* `item_name` - (Optional, String) The name of the software that you chose to install from the IBM Cloud catalog.
	* `item_readme_url` - (Optional, String) The URL to the readme file of the software template in the IBM Cloud catalog.
	* `item_url` - (Optional, String) The URL to the software template in the IBM Cloud catalog.
	* `launch_url` - (Optional, String) The URL to the dashboard to access your software.
	* `offering_version` - (Optional, String) The version of the software template that you chose to install from the IBM Cloud catalog.
	* `owning_account` - (Optional, String) Owning account ID of the catalog.
* `dependencies` - (Optional, List) Workspace dependencies.
Nested scheme for **dependencies**:
	* `children` - (Optional, List) List of workspace children CRN identifiers.
	* `parents` - (Optional, List) List of workspace parents CRN identifiers.
* `description` - (Optional, String) The description of the workspace.
* `location` - (Optional, String) The location where you want to create your Schematics workspace and run the Schematics jobs. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability.
* `name` - (Optional, String) The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name.
* `resource_group` - (Optional, String) The ID of the resource group where you want to provision the workspace.
* `shared_data` - (Optional, List) Information about the Target used by the templates originating from the  IBM Cloud catalog offerings. This information is not relevant for workspace created using your own Terraform template.
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
* `template_data` - (Optional, List) Input data for the Template.
Nested scheme for **template_data**:
	* `compact` - (Optional, Boolean) True, to use the files from the specified folder & subfolder in your GitHub or GitLab repository and ignore the other folders in the repository. For more information, see [Compact download for Schematics workspace](https://cloud.ibm.com/docs/schematics?topic=schematics-compact-download&interface=ui).
	* `env_values` - (Optional, List) A list of environment variables that you want to apply during the execution of a bash script or Terraform job. This field must be provided as a list of key-value pairs, for example, **TF_LOG=debug**. Each entry will be a map with one entry where `key is the environment variable name and value is value`. You can define environment variables for IBM Cloud catalog offerings that are provisioned by using a bash script. See [example to use special environment variable](https://cloud.ibm.com/docs/schematics?topic=schematics-set-parallelism#parallelism-example)  that are supported by Schematics.
	* `env_values_metadata` - (Optional, List) Environment variables metadata.
	Nested scheme for **env_values_metadata**:
		* `hidden` - (Optional, Boolean) Environment variable is hidden.
		* `name` - (Optional, String) Environment variable name.
		* `secure` - (Optional, Boolean) Environment variable is secure.
	* `folder` - (Optional, String) The subfolder in your GitHub or GitLab repository where your Terraform template is stored.
	* `init_state_file` - (Optional, String) The content of an existing Terraform statefile that you want to import in to your workspace. To get the content of a Terraform statefile for a specific Terraform template in an existing workspace, run `ibmcloud terraform state pull --id <workspace_id> --template <template_id>`.
	* `injectors` - (Optional, List) Array of injectable terraform blocks.
	Nested scheme for **injectors**:
		* `injection_type` - (Optional, String) Injection type. Default is 'override'.
		* `tft_git_token` - (Optional, String) Token to access the git repository (Optional).
		* `tft_git_url` - (Optional, String) Git repo url hosting terraform template files.
		* `tft_name` - (Optional, String) Terraform template name. Maps to folder name in git repo.
		* `tft_parameters` - (Optional, List)
		Nested scheme for **tft_parameters**:
			* `name` - (Optional, String) Key name to replace.
			* `value` - (Optional, String) Value to replace.
		* `tft_prefix` - (Optional, String) Optional prefix word to append to files (Optional).
	* `type` - (Optional, String) The Terraform version that you want to use to run your Terraform code. Enter `terraform_v1.1` to use Terraform version 1.1, and `terraform_v1.0` to use Terraform version 1.0. This is a required variable. Make sure that your Terraform config files are compatible with the Terraform version that you select. For more information, refer to [Terraform version](https://cloud.ibm.com/docs/schematics?topic=schematics-workspace-setup&interface=ui#create-workspace_ui).
	* `uninstall_script_name` - (Optional, String) Uninstall script name.
	* `values` - (Optional, String) A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `"autoscaling: enabled: true minReplicas: 2"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.
	* `values_metadata` - (Optional, List) List of values metadata.
	* `variablestore` - (Optional, List) VariablesRequest -.
	Nested scheme for **variablestore**:
		* `description` - (Optional, String) The description of your input variable.
		* `name` - (Optional, String) The name of the variable.
		* `secure` - (Optional, Boolean) If set to `true`, the value of your input variable is protected and not returned in your API response.
		* `type` - (Optional, String) `Terraform v0.11` supports `string`, `list`, `map` data type. For more information, about the syntax, see [Configuring input variables](https://www.terraform.io/docs/configuration-0-11/variables.html).<br> `Terraform v0.12` additionally, supports `bool`, `number` and complex data types such as `list(type)`, `map(type)`,`object({attribute name=type,..})`, `set(type)`, `tuple([type])`. For more information, about the syntax to use the complex data type, see [Configuring variables](https://www.terraform.io/docs/configuration/variables.html#type-constraints).
		* `use_default` - (Optional, Boolean) Variable uses default value; and is not over-ridden.
		* `value` - (Optional, String) Enter the value as a string for the primitive types such as `bool`, `number`, `string`, and `HCL` format for the complex variables, as you provide in a `.tfvars` file. **You need to enter escaped string of `HCL` format for the complex variable value**. For more information, about how to declare variables in a terraform configuration file and provide value to schematics, see [Providing values for the declared variables](https://cloud.ibm.com/docs/schematics?topic=schematics-create-tf-config#declare-variable).
* `template_ref` - (Optional, String) Workspace template ref.
* `template_repo` - (Optional, List) Input variables for the Template repoository, while creating a workspace.
Nested scheme for **template_repo**:
	* `branch` - (Optional, String) The repository branch.
	* `release` - (Optional, String) The repository release.
	* `repo_sha_value` - (Optional, String) The repository SHA value.
	* `repo_url` - (Optional, String) The repository URL.
	* `url` - (Optional, String) The source URL.
* `type` - (Optional, List) List of Workspace type.
* `workspace_status` - (Optional, List) WorkspaceStatusRequest -.
Nested scheme for **workspace_status**:
	* `frozen` - (Optional, Boolean) If set to true, the workspace is frozen and changes to the workspace are disabled.
	* `frozen_at` - (Optional, String) The timestamp when the workspace was frozen.
	* `frozen_by` - (Optional, String) The user ID that froze the workspace.
	* `locked` - (Optional, Boolean) If set to true, the workspace is locked and disabled for changes.
	* `locked_by` - (Optional, String) The user ID that initiated a resource-related job, such as applying or destroying resources, that locked the workspace.
	* `locked_time` - (Optional, String) The timestamp when the workspace was locked.
* `x_github_token` - (Optional, String) The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_workspace.
* `cart_id` - (Optional, String) The associate cart order ID.
* `created_at` - (Optional, String) The timestamp when the workspace was created.
* `created_by` - (Optional, String) The user ID that created the workspace.
* `crn` - (Optional, String) The workspace CRN.
* `last_action_name` - (Optional, String) Name of the last Action performed on workspace.
* `last_activity_id` - (Optional, String) ID of last Activity performed.
* `last_health_check_at` - (Optional, String) The timestamp when the last health check was performed by Schematics.
* `last_job` - (Optional, List) Last job details.
Nested scheme for **last_job**:
	* `job_id` - (Optional, String) ID of last job.
	* `job_name` - (Optional, String) Name of the last job.
	* `job_status` - (Optional, String) Status of the last job.
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
* `status` - (Optional, String) The status of the workspace.   **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.   **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.   **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.   **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.   **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply jobs to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.   **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.   **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.   **Stopped**: The Schematics plan, apply, or destroy job was cancelled manually.   **Template Error**: The Schematics template contains errors and cannot be processed.
* `updated_at` - (Optional, String) The timestamp when the workspace was last updated.
* `updated_by` - (Optional, String) The user ID that updated the workspace.
* `workspace_status_msg` - (Optional, List) Information about the last job that ran against the workspace. -.
Nested scheme for **workspace_status_msg**:
	* `status_code` - (Optional, String) The success or error code that was returned for the last plan, apply, or destroy job that ran against your workspace.
	* `status_msg` - (Optional, String) The success or error message that was returned for the last plan, apply, or destroy job that ran against your workspace.

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

You can import the `ibm_schematics_workspace` resource by using `id`. The unique identifier of the workspace.

# Syntax
```
$ terraform import ibm_schematics_workspace.schematics_workspace <id>
```
