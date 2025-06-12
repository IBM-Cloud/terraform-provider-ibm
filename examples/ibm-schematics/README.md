# Example for SchematicsV1

This example illustrates how to use the SchematicsV1

These types of resources are supported:

* schematics_workspace
* schematics_action
* schematics_job
* schematics_policy
* schematics_agent
* schematics_agent_prs
* schematics_agent_deploy
* schematics_agent_health

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
  tags = var.schematics_action_tags
  user_state = var.schematics_action_user_state
  source_readme_url = var.schematics_action_source_readme_url
  source = var.schematics_action_source
  source_type = var.schematics_action_source_type
  command_parameter = var.schematics_action_command_parameter
  bastion = var.schematics_action_bastion
  targets_ini = var.schematics_action_targets_ini
  credentials = var.schematics_action_credentials
  inputs = var.schematics_action_inputs
  outputs = var.schematics_action_outputs
  settings = var.schematics_action_settings
  trigger_record_id = var.schematics_action_trigger_record_id
  state = var.schematics_action_state
  sys_lock = var.schematics_action_sys_lock
  x_github_token = var.schematics_action_x_github_token
}
```
schematics_job resource:

```hcl
resource "schematics_job" "schematics_job_instance" {
  command_object = var.schematics_job_command_object
  command_object_id = var.schematics_job_command_object_id
  command_name = var.schematics_job_command_name
  command_parameter = var.schematics_job_command_parameter
  command_options = var.schematics_job_command_options
  inputs = var.schematics_job_inputs
  settings = var.schematics_job_settings
  tags = var.schematics_job_tags
  location = var.schematics_job_location
  status = var.schematics_job_status
  data = var.schematics_job_data
  bastion = var.schematics_job_bastion
  log_summary = var.schematics_job_log_summary
  x_github_token = var.schematics_job_x_github_token
}
```
schematics_policy resource:

```hcl
resource "schematics_policy" "schematics_policy_instance" {
  name = var.schematics_policy_name
  description = var.schematics_policy_description
  resource_group = var.schematics_policy_resource_group
  tags = var.schematics_policy_tags
  location = var.schematics_policy_location
  state = var.schematics_policy_state
  kind = var.schematics_policy_kind
  target = var.schematics_policy_target
  parameter = var.schematics_policy_parameter
  scoped_resources = var.schematics_policy_scoped_resources
}
```
schematics_agent resource:

```hcl
resource "schematics_agent" "schematics_agent_instance" {
  name = var.schematics_agent_name
  resource_group = var.schematics_agent_resource_group
  version = var.schematics_agent_version
  schematics_location = var.schematics_agent_schematics_location
  agent_location = var.schematics_agent_agent_location
  agent_infrastructure = var.schematics_agent_agent_infrastructure
  description = var.schematics_agent_description
  tags = var.schematics_agent_tags
  agent_metadata = var.schematics_agent_agent_metadata
  agent_inputs = var.schematics_agent_agent_inputs
  user_state = var.schematics_agent_user_state
  agent_kpi = var.schematics_agent_agent_kpi
}
```
schematics_agent_prs resource:

```hcl
resource "schematics_agent_prs" "schematics_agent_prs_instance" {
  agent_id = var.schematics_agent_prs_agent_id
  force = var.schematics_agent_prs_force
}
```
schematics_agent_deploy resource:

```hcl
resource "schematics_agent_deploy" "schematics_agent_deploy_instance" {
  agent_id = var.schematics_agent_deploy_agent_id
  force = var.schematics_agent_deploy_force
}
```
schematics_agent_health resource:

```hcl
resource "schematics_agent_health" "schematics_agent_health_instance" {
  agent_id = var.schematics_agent_health_agent_id
  force = var.schematics_agent_health_force
}
```

## SchematicsV1 Data sources

schematics_output data source:

```hcl
data "schematics_output" "schematics_output_instance" {
  workspace_id = var.schematics_output_workspace_id
}
```
schematics_state data source:

```hcl
data "schematics_state" "schematics_state_instance" {
  workspace_id = var.schematics_state_workspace_id
  template_id = var.schematics_state_template_id
}
```
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
schematics_policies data source:

```hcl
data "schematics_policies" "schematics_policies_instance" {
  policy_kind = var.schematics_policies_policy_kind
}
```
schematics_policy data source:

```hcl
data "schematics_policy" "schematics_policy_instance" {
  policy_id = var.schematics_policy_policy_id
}
```
schematics_agents data source:

```hcl
data "schematics_agents" "schematics_agents_instance" {
}
```
schematics_agent data source:

```hcl
data "schematics_agent" "schematics_agent_instance" {
  agent_id = var.schematics_agent_agent_id
}
```
schematics_agent_prs data source:

```hcl
data "schematics_agent_prs" "schematics_agent_prs_instance" {
  agent_id = var.schematics_agent_prs_agent_id
}
```
schematics_agent_deploy data source:

```hcl
data "schematics_agent_deploy" "schematics_agent_deploy_instance" {
  agent_id = var.schematics_agent_deploy_agent_id
}
```
schematics_agent_health data source:

```hcl
data "schematics_agent_health" "schematics_agent_health_instance" {
  agent_id = var.schematics_agent_health_agent_id
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 1.5 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.58.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| applied_shareddata_ids | List of applied shared dataset id. | `list(string)` | false |
| catalog_ref | Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only. | `` | false |
| description | The description of the workspace. | `string` | false |
| location | The location where you want to create your Schematics workspace and run Schematics actions. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability. | `string` | false |
| name | The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name. | `string` | false |
| resource_group | The ID of the resource group where you want to provision the workspace. | `string` | false |
| shared_data | Information that is shared across templates in IBM Cloud catalog offerings. This information is not provided when you create a workspace from your own Terraform template. | `` | false |
| tags | A list of tags that are associated with the workspace. | `list(string)` | false |
| template_data | TemplateData -. | `list()` | false |
| template_ref | Workspace template ref. | `string` | false |
| template_repo | Input parameter to specify the source repository where your Schematics template is stored. | `` | false |
| type | The Terraform version that you want to use to run your Terraform code. Enter `terraform_v0.12` to use Terraform version 0.12, and `terraform_v0.11` to use Terraform version 0.11. If no value is specified, the Terraform config files are run with Terraform version 0.11. Make sure that your Terraform config files are compatible with the Terraform version that you select. | `list(string)` | false |
| workspace_status | WorkspaceStatusRequest -. | `` | false |
| x_github_token | The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template. | `string` | false |
| name | Action name (unique for an account). | `string` | false |
| description | Action description. | `string` | false |
| location | List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics. | `string` | false |
| resource_group | Resource-group name for an action.  By default, action is created in default resource group. | `string` | false |
| tags | Action tags. | `list(string)` | false |
| user_state | User defined status of the Schematics object. | `` | false |
| source_readme_url | URL of the `README` file, for the source. | `string` | false |
| source | Source of templates, playbooks, or controls. | `` | false |
| source_type | Type of source for the Template. | `string` | false |
| command_parameter | Schematics job command parameter (playbook-name, capsule-name or flow-name). | `string` | false |
| bastion | Complete target details with the user inputs and the system generated data. | `` | false |
| targets_ini | Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps). | `string` | false |
| inputs | Input variables for an action. | `list()` | false |
| outputs | Output variables for an action. | `list()` | false |
| settings | Environment variables for an action. | `list()` | false |
| trigger_record_id | ID to the trigger. | `string` | false |
| state | Computed state of an action. | `` | false |
| sys_lock | System lock status. | `` | false |
| x_github_token | The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template. | `string` | false |
| command_object | Name of the Schematics automation resource. | `string` | false |
| command_object_id | Job command object ID (`workspace-id, action-id or control-id`). | `string` | false |
| command_name | Schematics job command name. | `string` | false |
| command_parameter | Schematics job command parameter (`playbook-name, capsule-name or flow-name`). | `string` | false |
| command_options | Command line options for the command. | `list(string)` | false |
| inputs | Job inputs used by an action. | `list()` | false |
| settings | Environment variables used by the job while performing an action. | `list()` | false |
| tags | User defined tags, while running the job. | `list(string)` | false |
| location | List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics. | `string` | false |
| status | Job Status. | `` | false |
| data | Job data. | `` | false |
| bastion | Complete target details with the user inputs and the system generated data. | `` | false |
| log_summary | Job log summary record. | `` | false |
| x_github_token | Create a job record and launch the job. | `string` | false |
| workspace_id | The ID of the workspace for which you want to retrieve output values. To find the workspace ID, use the `GET /workspaces` API. | `string` | true |
| workspace_id | The ID of the workspace for which you want to retrieve the Terraform statefile. To find the workspace ID, use the `GET /v1/workspaces` API. | `string` | true |
| template_id | The ID of the Terraform template for which you want to retrieve the Terraform statefile. When you create a workspace, the Terraform template that your workspace points to is assigned a unique ID. To find this ID, use the `GET /v1/workspaces` API and review the `template_data.id` value. | `string` | true |
| workspace_id | The ID of the workspace for which you want to retrieve detailed information. To find the workspace ID, use the `GET /v1/workspaces` API. | `string` | true |
| action_id | Use GET or actions API to look up the action IDs in your IBM Cloud account. | `string` | true |
| job_id | Use GET jobs API to look up the Job IDs in your IBM Cloud account. | `string` | true |
| schematics_policy_name | Name of Schematics customization policy. | `string` | true |
| schematics_policy_description | The description of Schematics customization policy. | `string` | false |
| schematics_policy_resource_group | The resource group name for the policy.  By default, Policy will be created in `default` Resource Group. | `string` | false |
| schematics_policy_tags | Tags for the Schematics customization policy. | `list(string)` | false |
| schematics_policy_location | List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics. | `string` | false |
| schematics_policy_kind | Policy kind or categories for managing and deriving policy decision  * `agent_assignment_policy` Agent assignment policy for job execution. | `string` | false |
| target | The objects for the Schematics policy. | `` | false |
| parameter | The parameter to tune the Schematics policy. | `` | false |
| scoped_resources | List of scoped Schematics resources targeted by the policy. | `list()` | false |
| schematics_agent_name | The name of the agent (must be unique, for an account). | `string` | true |
| schematics_agent_resource_group | The resource-group name for the agent.  By default, agent will be registered in Default Resource Group. | `string` | true |
| schematics_agent_version | Agent version. | `string` | true |
| schematics_agent_schematics_location | List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit thejegccnhbrgfulbgbrnftiuudlikviubvnnhlbkcfrhl
e location of the IBM Cloud resources, provisioned using Schematics. | `string` | true |
| schematics_agent_agent_location | The location where agent is deployed in the user environment. | `string` | true |
| agent_infrastructure | The infrastructure parameters used by the agent. | `` | true |
| schematics_agent_description | Agent description. | `string` | false |
| schematics_agent_tags | Tags for the agent. | `list(string)` | false |
| agent_metadata | The metadata of an agent. | `list()` | false |
| agent_inputs | Additional input variables for the agent. | `list()` | false |
| user_state | User defined status of the agent. | `` | false |
| agent_kpi | Schematics Agent key performance indicators. | `` | false |
| schematics_agent_prs_agent_id | Agent ID to get the details of agent. | `string` | true |
| schematics_agent_prs_force | Equivalent to -force options in the command line, default is false. | `bool` | false |
| schematics_agent_deploy_agent_id | Agent ID to get the details of agent. | `string` | true |
| schematics_agent_deploy_force | Equivalent to -force options in the command line, default is false. | `bool` | false |
| schematics_agent_health_agent_id | Agent ID to get the details of agent. | `string` | true |
| schematics_agent_health_force | Equivalent to -force options in the command line, default is false. | `bool` | false |
| schematics_policies_policy_kind | Policy kind or categories for managing and deriving policy decision  * `agent_assignment_policy` Agent assignment policy for job execution. | `string` | false |
| schematics_policy_policy_id | ID to get the details of policy. | `string` | true |
| schematics_agent_agent_id | Agent ID to get the details of agent. | `string` | true |
## Outputs

| Name | Description |
|------|-------------|
| schematics_workspace | schematics_workspace object |
| schematics_action | schematics_action object |
| schematics_job | schematics_job object |
| schematics_output | schematics_output object |
| schematics_state | schematics_state object |
| ibm_schematics_policy | schematics_policy resource instance |
| ibm_schematics_agent | schematics_agent resource instance |
| ibm_schematics_agent_prs | schematics_agent_prs resource instance |
| ibm_schematics_agent_deploy | schematics_agent_deploy resource instance |
| ibm_schematics_agent_health | schematics_agent_health resource instance |
