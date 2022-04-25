variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for schematics_workspace
variable "schematics_workspace_applied_shareddata_ids" {
  description = "List of applied shared dataset ID."
  type        = list(string)
  default     = [ "applied_shareddata_ids" ]
}
variable "schematics_workspace_description" {
  description = "The description of the workspace."
  type        = string
  default     = "description"
}
variable "schematics_workspace_location" {
  description = "The location where you want to create your Schematics workspace and run the Schematics jobs. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability."
  type        = string
  default     = "location"
}
variable "schematics_workspace_name" {
  description = "The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name."
  type        = string
  default     = "name"
}
variable "schematics_workspace_resource_group" {
  description = "The ID of the resource group where you want to provision the workspace."
  type        = string
  default     = "resource_group"
}
variable "schematics_workspace_tags" {
  description = "A list of tags that are associated with the workspace."
  type        = list(string)
  default     = [ "tags" ]
}
variable "schematics_workspace_template_ref" {
  description = "Workspace template ref."
  type        = string
  default     = "template_ref"
}
variable "schematics_workspace_type" {
  description = "List of Workspace type."
  type        = list(string)
  default     = [ "type" ]
}
variable "schematics_workspace_x_github_token" {
  description = "The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template."
  type        = string
  default     = "x_github_token"
}

// Resource arguments for schematics_action
variable "schematics_action_name" {
  description = "The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action."
  type        = string
  default     = "Stop Action"
}
variable "schematics_action_description" {
  description = "Action description."
  type        = string
  default     = "The description of your action. The description can be up to 2048 characters long in size. **Example** you can use the description to stop the targets."
}
variable "schematics_action_location" {
  description = "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics."
  type        = string
  default     = "us-south"
}
variable "schematics_action_resource_group" {
  description = "Resource-group name for an action. By default, an action is created in `Default` resource group."
  type        = string
  default     = "resource_group"
}
variable "schematics_action_bastion_connection_type" {
  description = "Type of connection to be used when connecting to bastion host. If the `inventory_connection_type=winrm`, then `bastion_connection_type` is not supported."
  type        = string
  default     = "ssh"
}
variable "schematics_action_inventory_connection_type" {
  description = "Type of connection to be used when connecting to remote host. **Note** Currently, WinRM supports only Windows system with the public IPs and do not support Bastion host."
  type        = string
  default     = "ssh"
}
variable "schematics_action_tags" {
  description = "Action tags."
  type        = list(string)
  default     = [ "tags" ]
}
variable "schematics_action_source_readme_url" {
  description = "URL of the `README` file, for the source URL."
  type        = string
  default     = "source_readme_url"
}
variable "schematics_action_source_type" {
  description = "Type of source for the Template."
  type        = string
  default     = "local"
}
variable "schematics_action_command_parameter" {
  description = "Schematics job command parameter (playbook-name)."
  type        = string
  default     = "command_parameter"
}
variable "schematics_action_inventory" {
  description = "Target inventory record ID, used by the action or ansible playbook."
  type        = string
  default     = "inventory"
}
variable "schematics_action_targets_ini" {
  description = "Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost] 172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps)."
  type        = string
  default     = "targets_ini"
}
variable "schematics_action_x_github_token" {
  description = "The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template."
  type        = string
  default     = "x_github_token"
}

// Resource arguments for schematics_job
variable "schematics_job_refresh_token" {
  description = "The IAM refresh token for the user or service identity.  **Retrieving refresh token**:   * Use `export IBMCLOUD_API_KEY=<ibmcloud_api_key>`, and execute `curl -X POST "https://iam.cloud.ibm.com/identity/token" -H "Content-Type: application/x-www-form-urlencoded" -d "grant_type=urn:ibm:params:oauth:grant-type:apikey&apikey=$IBMCLOUD_API_KEY" -u bx:bx`.   * For more information, about creating IAM access token and API Docs, refer, [IAM access token](/apidocs/iam-identity-token-api#gettoken-password) and [Create API key](/apidocs/iam-identity-token-api#create-api-key).    **Limitation**:   * If the token is expired, you can use `refresh token` to get a new IAM access token.   * The `refresh_token` parameter cannot be used to retrieve a new IAM access token.   * When the IAM access token is about to expire, use the API key to create a new access token."
  type        = string
  default     = "refresh_token"
}
variable "schematics_job_command_object" {
  description = "Name of the Schematics automation resource."
  type        = string
  default     = "workspace"
}
variable "schematics_job_command_object_id" {
  description = "Job command object id (workspace-id, action-id)."
  type        = string
  default     = "command_object_id"
}
variable "schematics_job_command_name" {
  description = "Schematics job command name."
  type        = string
  default     = "workspace_plan"
}
variable "schematics_job_command_parameter" {
  description = "Schematics job command parameter (playbook-name)."
  type        = string
  default     = "command_parameter"
}
variable "schematics_job_command_options" {
  description = "Command line options for the command."
  type        = list(string)
  default     = [ "command_options" ]
}
variable "schematics_job_tags" {
  description = "User defined tags, while running the job."
  type        = list(string)
  default     = [ "tags" ]
}
variable "schematics_job_location" {
  description = "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics."
  type        = string
  default     = "us-south"
}

// Resource arguments for schematics_inventory
variable "schematics_inventory_name" {
  description = "The unique name of your Inventory definition. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores."
  type        = string
  default     = "name"
}
variable "schematics_inventory_description" {
  description = "The description of your Inventory definition. The description can be up to 2048 characters long in size."
  type        = string
  default     = "description"
}
variable "schematics_inventory_location" {
  description = "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics."
  type        = string
  default     = "us-south"
}
variable "schematics_inventory_resource_group" {
  description = "Resource-group name for the Inventory definition.   By default, Inventory definition will be created in Default Resource Group."
  type        = string
  default     = "resource_group"
}
variable "schematics_inventory_inventories_ini" {
  description = "Input inventory of host and host group for the playbook, in the `.ini` file format."
  type        = string
  default     = "inventories_ini"
}
variable "schematics_inventory_resource_queries" {
  description = "Input resource query definitions that is used to dynamically generate the inventory of host and host group for the playbook."
  type        = list(string)
  default     = [ "resource_queries" ]
}

// Resource arguments for schematics_resource_query
variable "schematics_resource_query_type" {
  description = "Resource type (cluster, vsi, icd, vpc)."
  type        = string
  default     = "vsi"
}
variable "schematics_resource_query_name" {
  description = "Resource query name."
  type        = string
  default     = "name"
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
  description = "The ID of the workspace.  To find the workspace ID, use the `GET /v1/workspaces` API."
  type        = string
  default     = "workspace_id"
}

// Data source arguments for schematics_action
variable "schematics_action_action_id" {
  description = "Action Id.  Use GET /actions API to look up the Action Ids in your IBM Cloud account."
  type        = string
  default     = "action_id"
}

// Data source arguments for schematics_job
variable "schematics_job_job_id" {
  description = "Job Id. Use `GET /v2/jobs` API to look up the Job Ids in your IBM Cloud account."
  type        = string
  default     = "job_id"
}

// Data source arguments for schematics_inventory
variable "schematics_inventory_inventory_id" {
  description = "Resource Inventory Id.  Use `GET /v2/inventories` API to look up the Resource Inventory definition Ids  in your IBM Cloud account."
  type        = string
  default     = "inventory_id"
}

// Data source arguments for schematics_resource_query
variable "schematics_resource_query_query_id" {
  description = "Resource query Id.  Use `GET /v2/resource_query` API to look up the Resource query definition Ids  in your IBM Cloud account."
  type        = string
  default     = "query_id"
}
