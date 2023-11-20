variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for schematics_workspace
variable "schematics_workspace_applied_shareddata_ids" {
  description = "List of applied shared dataset id."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "schematics_workspace_catalog_ref" {
  description = "Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_workspace_description" {
  description = "The description of the workspace."
  type        = string
  default     = "placeholder"
}
variable "schematics_workspace_location" {
  description = "The location where you want to create your Schematics workspace and run Schematics actions. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability."
  type        = string
  default     = "placeholder"
}
variable "schematics_workspace_name" {
  description = "The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name."
  type        = string
  default     = "placeholder"
}
variable "schematics_workspace_resource_group" {
  description = "The ID of the resource group where you want to provision the workspace."
  type        = string
  default     = "placeholder"
}
variable "schematics_workspace_shared_data" {
  description = "Information that is shared across templates in IBM Cloud catalog offerings. This information is not provided when you create a workspace from your own Terraform template."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_workspace_tags" {
  description = "A list of tags that are associated with the workspace."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "schematics_workspace_template_data" {
  description = "TemplateData -."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_workspace_template_ref" {
  description = "Workspace template ref."
  type        = string
  default     = "placeholder"
}
variable "schematics_workspace_template_repo" {
  description = "Input parameter to specify the source repository where your Schematics template is stored."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_workspace_type" {
  description = "The Terraform version that you want to use to run your Terraform code. Enter `terraform_v0.12` to use Terraform version 0.12, and `terraform_v0.11` to use Terraform version 0.11. If no value is specified, the Terraform config files are run with Terraform version 0.11. Make sure that your Terraform config files are compatible with the Terraform version that you select."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "schematics_workspace_workspace_status" {
  description = "WorkspaceStatusRequest -."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_workspace_x_github_token" {
  description = "The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template."
  type        = string
  default     = "placeholder"
}

// Resource arguments for schematics_action
variable "schematics_action_name" {
  description = "Action name (unique for an account)."
  type        = string
  default     = "Stop Action"
}
variable "schematics_action_description" {
  description = "Action description."
  type        = string
  default     = "This Action can be used to Stop the targets."
}
variable "schematics_action_location" {
  description = "List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics."
  type        = string
  default     = "placeholder"
}
variable "schematics_action_resource_group" {
  description = "Resource-group name for an action.  By default, action is created in default resource group."
  type        = string
  default     = "placeholder"
}
variable "schematics_action_tags" {
  description = "Action tags."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "schematics_action_user_state" {
  description = "User defined status of the Schematics object."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_source_readme_url" {
  description = "URL of the `README` file, for the source."
  type        = string
  default     = "placeholder"
}
variable "schematics_action_source" {
  description = "Source of templates, playbooks, or controls."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_source_type" {
  description = "Type of source for the Template."
  type        = string
  default     = "placeholder"
}
variable "schematics_action_command_parameter" {
  description = "Schematics job command parameter (playbook-name, capsule-name or flow-name)."
  type        = string
  default     = "placeholder"
}
variable "schematics_action_bastion" {
  description = "Complete target details with the user inputs and the system generated data."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_targets_ini" {
  description = "Inventory of host and host group for the playbook in `INI` file format. For example, `'targets_ini': '[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5'`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps)."
  type        = string
  default     = "placeholder"
}
variable "schematics_action_credentials" {
  description = "credentials of the Action."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_inputs" {
  description = "Input variables for an action."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_outputs" {
  description = "Output variables for an action."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_settings" {
  description = "Environment variables for an action."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_trigger_record_id" {
  description = "ID to the trigger."
  type        = string
  default     = "placeholder"
}
variable "schematics_action_state" {
  description = "Computed state of an action."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_sys_lock" {
  description = "System lock status."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_action_x_github_token" {
  description = "The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template."
  type        = string
  default     = "placeholder"
}

variable "schematics_job_command_object" {
  description = "Name of the Schematics automation resource."
  type        = string
  default     = "placeholder"
}
variable "schematics_job_command_object_id" {
  description = "Job command object ID (`workspace-id, action-id or control-id`)."
  type        = string
  default     = "placeholder"
}
variable "schematics_job_command_name" {
  description = "Schematics job command name."
  type        = string
  default     = "placeholder"
}
variable "schematics_job_command_parameter" {
  description = "Schematics job command parameter (`playbook-name, capsule-name or flow-name`)."
  type        = string
  default     = "placeholder"
}
variable "schematics_job_command_options" {
  description = "Command line options for the command."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "schematics_job_inputs" {
  description = "Job inputs used by an action."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_job_settings" {
  description = "Environment variables used by the job while performing an action."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_job_tags" {
  description = "User defined tags, while running the job."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "schematics_job_location" {
  description = "List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics."
  type        = string
  default     = "placeholder"
}
variable "schematics_job_status" {
  description = "Job Status."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_job_data" {
  description = "Job data."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_job_bastion" {
  description = "Complete target details with the user inputs and the system generated data."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_job_log_summary" {
  description = "Job log summary record."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "schematics_job_x_github_token" {
  description = "Create a job record and launch the job."
  type        = string
  default     = "placeholder"
}

// Data source arguments for schematics_output
variable "schematics_output_workspace_id" {
  description = "The ID of the workspace for which you want to retrieve output values. To find the workspace ID, use the `GET /workspaces` API."
  type        = string
  default     = "workspace_id"
}

// Data source arguments for schematics_state
variable "schematics_state_workspace_id" {
  description = "The ID of the workspace for which you want to retrieve the Terraform statefile. To find the workspace ID, use the `GET /v1/workspaces` API."
  type        = string
  default     = "workspace_id"
}
variable "schematics_state_template_id" {
  description = "The ID of the Terraform template for which you want to retrieve the Terraform statefile. When you create a workspace, the Terraform template that your workspace points to is assigned a unique ID. To find this ID, use the `GET /v1/workspaces` API and review the `template_data.id` value."
  type        = string
  default     = "template_id"
}

// Data source arguments for schematics_workspace
variable "schematics_workspace_workspace_id" {
  description = "The ID of the workspace for which you want to retrieve detailed information. To find the workspace ID, use the `GET /v1/workspaces` API."
  type        = string
  default     = "workspace_id"
}

// Data source arguments for schematics_action
variable "schematics_action_action_id" {
  description = "Use GET or actions API to look up the action IDs in your IBM Cloud account."
  type        = string
  default     = "action_id"
}

// Data source arguments for schematics_job
variable "schematics_job_job_id" {
  description = "Use GET jobs API to look up the Job IDs in your IBM Cloud account."
  type        = string
  default     = "job_id"
}


// Resource arguments for schematics_policy
variable "schematics_policy_name" {
  description = "Name of Schematics customization policy."
  type        = string
  default     = "Agent1-DevWS"
}
variable "schematics_policy_description" {
  description = "The description of Schematics customization policy."
  type        = string
  default     = "Policy for job execution of secured workspaces on agent1"
}
variable "schematics_policy_resource_group" {
  description = "The resource group name for the policy.  By default, Policy will be created in `default` Resource Group."
  type        = string
  default     = "Default"
}
variable "schematics_policy_tags" {
  description = "Tags for the Schematics customization policy."
  type        = list(string)
  default     = ["policy:secured-job"]
}
variable "schematics_policy_location" {
  description = "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics."
  type        = string
  default     = "us-south"
}
variable "schematics_policy_kind" {
  description = "Policy kind or categories for managing and deriving policy decision  * `agent_assignment_policy` Agent assignment policy for job execution."
  type        = string
  default     = "agent_assignment_policy"
}

// Resource arguments for schematics_agent
variable "schematics_agent_name" {
  description = "The name of the agent (must be unique, for an account)."
  type        = string
  default     = "MyDevAgent"
}
variable "schematics_agent_resource_group" {
  description = "The resource-group name for the agent.  By default, agent will be registered in Default Resource Group."
  type        = string
  default     = "Default"
}
variable "schematics_agent_version" {
  description = "Agent version."
  type        = string
  default     = "1.0.0-beta2"
}
variable "schematics_agent_schematics_location" {
  description = "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics."
  type        = string
  default     = "us-south"
}
variable "schematics_agent_agent_location" {
  description = "The location where agent is deployed in the user environment."
  type        = string
  default     = "us-south"
}
variable "schematics_agent_description" {
  description = "Agent description."
  type        = string
  default     = "Create Agent"
}
variable "schematics_agent_tags" {
  description = "Tags for the agent."
  type        = list(string)
  default     = [ "tags" ]
}

// Resource arguments for schematics_agent_prs
variable "schematics_agent_prs_agent_id" {
  description = "Agent ID to get the details of agent."
  type        = string
  default     = "agent_id"
}
variable "schematics_agent_prs_force" {
  description = "Equivalent to -force options in the command line, default is false."
  type        = bool
  default     = true
}

// Resource arguments for schematics_agent_deploy
variable "schematics_agent_deploy_agent_id" {
  description = "Agent ID to get the details of agent."
  type        = string
  default     = "agent_id"
}
variable "schematics_agent_deploy_force" {
  description = "Equivalent to -force options in the command line, default is false."
  type        = bool
  default     = true
}

// Resource arguments for schematics_agent_health
variable "schematics_agent_health_agent_id" {
  description = "Agent ID to get the details of agent."
  type        = string
  default     = "agent_id"
}
variable "schematics_agent_health_force" {
  description = "Equivalent to -force options in the command line, default is false."
  type        = bool
  default     = true
}

// Data source arguments for schematics_policies
variable "schematics_policies_policy_kind" {
  description = "Policy kind or categories for managing and deriving policy decision  * `agent_assignment_policy` Agent assignment policy for job execution."
  type        = string
  default     = "agent_assignment_policy"
}

// Data source arguments for schematics_policy
variable "schematics_policy_policy_id" {
  description = "ID to get the details of policy."
  type        = string
  default     = "policy_id"
}

// Data source arguments for schematics_agent
variable "schematics_agent_agent_id" {
  description = "Agent ID to get the details of agent."
  type        = string
  default     = "agent_id"
}