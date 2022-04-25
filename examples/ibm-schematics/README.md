# Example for SchematicsV1

This example illustrates how to use the SchematicsV1

These types of resources are supported:

* schematics_workspace
* schematics_action
* schematics_job
* schematics_inventory
* schematics_resource_query

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## SchematicsV1 resources

schematics_workspace resource:

```hcl
resource "schematics_workspace" "schematics_workspace_instance" {
  applied_shareddata_ids = var.schematics_workspace_applied_shareddata_ids
  catalog_ref = var.schematics_workspace_catalog_ref
  dependencies = var.schematics_workspace_dependencies
  description = var.schematics_workspace_description
  location = var.schematics_workspace_location
  name = var.schematics_workspace_name
  resource_group = var.schematics_workspace_resource_group
  shared_data = var.schematics_workspace_shared_data
  tags = var.schematics_workspace_tags
  template_data = var.schematics_workspace_template_data
  template_ref = var.schematics_workspace_template_ref
  template_repo = var.schematics_workspace_template_repo
  type = var.schematics_workspace_type
  workspace_status = var.schematics_workspace_workspace_status
  x_github_token = var.schematics_workspace_x_github_token
}
```
schematics_action resource:

```hcl
resource "schematics_action" "schematics_action_instance" {
  name = var.schematics_action_name
  description = var.schematics_action_description
  location = var.schematics_action_location
  resource_group = var.schematics_action_resource_group
  bastion_connection_type = var.schematics_action_bastion_connection_type
  inventory_connection_type = var.schematics_action_inventory_connection_type
  tags = var.schematics_action_tags
  user_state = var.schematics_action_user_state
  source_readme_url = var.schematics_action_source_readme_url
  source = var.schematics_action_source
  source_type = var.schematics_action_source_type
  command_parameter = var.schematics_action_command_parameter
  inventory = var.schematics_action_inventory
  credentials = var.schematics_action_credentials
  bastion = var.schematics_action_bastion
  bastion_credential = var.schematics_action_bastion_credential
  targets_ini = var.schematics_action_targets_ini
  action_inputs = var.schematics_action_action_inputs
  action_outputs = var.schematics_action_action_outputs
  settings = var.schematics_action_settings
  state = var.schematics_action_state
  sys_lock = var.schematics_action_sys_lock
  x_github_token = var.schematics_action_x_github_token
}
```
schematics_job resource:

```hcl
resource "schematics_job" "schematics_job_instance" {
  refresh_token = var.schematics_job_refresh_token
  command_object = var.schematics_job_command_object
  command_object_id = var.schematics_job_command_object_id
  command_name = var.schematics_job_command_name
  command_parameter = var.schematics_job_command_parameter
  command_options = var.schematics_job_command_options
  job_inputs = var.schematics_job_job_inputs
  job_env_settings = var.schematics_job_job_env_settings
  tags = var.schematics_job_tags
  location = var.schematics_job_location
  status = var.schematics_job_status
  data = var.schematics_job_data
  bastion = var.schematics_job_bastion
  log_summary = var.schematics_job_log_summary
}
```
schematics_inventory resource:

```hcl
resource "schematics_inventory" "schematics_inventory_instance" {
  name = var.schematics_inventory_name
  description = var.schematics_inventory_description
  location = var.schematics_inventory_location
  resource_group = var.schematics_inventory_resource_group
  inventories_ini = var.schematics_inventory_inventories_ini
  resource_queries = var.schematics_inventory_resource_queries
}
```
schematics_resource_query resource:

```hcl
resource "schematics_resource_query" "schematics_resource_query_instance" {
  type = var.schematics_resource_query_type
  name = var.schematics_resource_query_name
  queries = var.schematics_resource_query_queries
}
```

## SchematicsV1 Data sources

schematics_workspace data source:

```hcl
data "schematics_workspace" "schematics_workspace_instance" {
  workspace_id = var.schematics_workspace_workspace_id
}
```
schematics_action data source:

```hcl
data "schematics_action" "schematics_action_instance" {
  action_id = var.schematics_action_action_id
}
```
schematics_job data source:

```hcl
data "schematics_job" "schematics_job_instance" {
  job_id = var.schematics_job_job_id
}
```
schematics_inventory data source:

```hcl
data "schematics_inventory" "schematics_inventory_instance" {
  inventory_id = var.schematics_inventory_inventory_id
}
```
schematics_resource_query data source:

```hcl
data "schematics_resource_query" "schematics_resource_query_instance" {
  query_id = var.schematics_resource_query_query_id
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| applied_shareddata_ids | List of applied shared dataset ID. | `list(string)` | false |
| catalog_ref | Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only. | `` | false |
| dependencies | Workspace dependencies. | `` | false |
| description | The description of the workspace. | `string` | false |
| location | The location where you want to create your Schematics workspace and run the Schematics jobs. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability. | `string` | false |
| name | The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name. | `string` | false |
| resource_group | The ID of the resource group where you want to provision the workspace. | `string` | false |
| shared_data | Information about the Target used by the templates originating from the  IBM Cloud catalog offerings. This information is not relevant for workspace created using your own Terraform template. | `` | false |
| tags | A list of tags that are associated with the workspace. | `list(string)` | false |
| template_data | Input data for the Template. | `list()` | false |
| template_ref | Workspace template ref. | `string` | false |
| template_repo | Input variables for the Template repoository, while creating a workspace. | `` | false |
| type | List of Workspace type. | `list(string)` | false |
| workspace_status | WorkspaceStatusRequest -. | `` | false |
| x_github_token | The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template. | `string` | false |
| name | The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action. | `string` | false |
| description | Action description. | `string` | false |
| location | List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics. | `string` | false |
| resource_group | Resource-group name for an action. By default, an action is created in `Default` resource group. | `string` | false |
| bastion_connection_type | Type of connection to be used when connecting to bastion host. If the `inventory_connection_type=winrm`, then `bastion_connection_type` is not supported. | `string` | false |
| inventory_connection_type | Type of connection to be used when connecting to remote host. **Note** Currently, WinRM supports only Windows system with the public IPs and do not support Bastion host. | `string` | false |
| tags | Action tags. | `list(string)` | false |
| user_state | User defined status of the Schematics object. | `` | false |
| source_readme_url | URL of the `README` file, for the source URL. | `string` | false |
| source | Source of templates, playbooks, or controls. | `` | false |
| source_type | Type of source for the Template. | `string` | false |
| command_parameter | Schematics job command parameter (playbook-name). | `string` | false |
| inventory | Target inventory record ID, used by the action or ansible playbook. | `string` | false |
| credentials | credentials of the Action. | `list()` | false |
| bastion | Describes a bastion resource. | `` | false |
| bastion_credential | User editable credential variable data and system generated reference to the value. | `` | false |
| targets_ini | Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost] 172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps). | `string` | false |
| action_inputs | Input variables for the Action. | `list()` | false |
| action_outputs | Output variables for the Action. | `list()` | false |
| settings | Environment variables for the Action. | `list()` | false |
| state | Computed state of the Action. | `` | false |
| sys_lock | System lock status. | `` | false |
| x_github_token | The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template. | `string` | false |
| refresh_token | The IAM refresh token for the user or service identity.  **Retrieving refresh token**:   * Use `export IBMCLOUD_API_KEY=<ibmcloud_api_key>`, and execute `curl -X POST "https://iam.cloud.ibm.com/identity/token" -H "Content-Type: application/x-www-form-urlencoded" -d "grant_type=urn:ibm:params:oauth:grant-type:apikey&apikey=$IBMCLOUD_API_KEY" -u bx:bx`.   * For more information, about creating IAM access token and API Docs, refer, [IAM access token](/apidocs/iam-identity-token-api#gettoken-password) and [Create API key](/apidocs/iam-identity-token-api#create-api-key).    **Limitation**:   * If the token is expired, you can use `refresh token` to get a new IAM access token.   * The `refresh_token` parameter cannot be used to retrieve a new IAM access token.   * When the IAM access token is about to expire, use the API key to create a new access token. | `string` | true |
| command_object | Name of the Schematics automation resource. | `string` | false |
| command_object_id | Job command object id (workspace-id, action-id). | `string` | false |
| command_name | Schematics job command name. | `string` | false |
| command_parameter | Schematics job command parameter (playbook-name). | `string` | false |
| command_options | Command line options for the command. | `list(string)` | false |
| job_inputs | Job inputs used by Action or Workspace. | `list()` | false |
| job_env_settings | Environment variables used by the Job while performing Action or Workspace. | `list()` | false |
| tags | User defined tags, while running the job. | `list(string)` | false |
| location | List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics. | `string` | false |
| status | Job Status. | `` | false |
| data | Job data. | `` | false |
| bastion | Describes a bastion resource. | `` | false |
| log_summary | Job log summary record. | `` | false |
| name | The unique name of your Inventory definition. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. | `string` | false |
| description | The description of your Inventory definition. The description can be up to 2048 characters long in size. | `string` | false |
| location | List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics. | `string` | false |
| resource_group | Resource-group name for the Inventory definition.   By default, Inventory definition will be created in Default Resource Group. | `string` | false |
| inventories_ini | Input inventory of host and host group for the playbook, in the `.ini` file format. | `string` | false |
| resource_queries | Input resource query definitions that is used to dynamically generate the inventory of host and host group for the playbook. | `list(string)` | false |
| type | Resource type (cluster, vsi, icd, vpc). | `string` | false |
| name | Resource query name. | `string` | false |
| queries |  | `list()` | false |
| workspace_id | The ID of the workspace.  To find the workspace ID, use the `GET /v1/workspaces` API. | `string` | true |
| action_id | Action Id.  Use GET /actions API to look up the Action Ids in your IBM Cloud account. | `string` | true |
| job_id | Job Id. Use `GET /v2/jobs` API to look up the Job Ids in your IBM Cloud account. | `string` | true |
| inventory_id | Resource Inventory Id.  Use `GET /v2/inventories` API to look up the Resource Inventory definition Ids  in your IBM Cloud account. | `string` | true |
| query_id | Resource query Id.  Use `GET /v2/resource_query` API to look up the Resource query definition Ids  in your IBM Cloud account. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| schematics_workspace | schematics_workspace object |
| schematics_action | schematics_action object |
| schematics_job | schematics_job object |
| schematics_inventory | schematics_inventory object |
| schematics_resource_query | schematics_resource_query object |
| schematics_workspace | schematics_workspace object |
| schematics_action | schematics_action object |
| schematics_job | schematics_job object |
| schematics_inventory | schematics_inventory object |
| schematics_resource_query | schematics_resource_query object |
