variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for pdr_managedr
variable "pdr_managedr_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_managedr_stand_by_redeploy" {
  description = "Flag to indicate if standby should be redeployed (must be \"true\" or \"false\")."
  type        = string
  default     = "true"
}
variable "pdr_managedr_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "accept_language"
}
variable "pdr_managedr_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "if_none_match"
}
variable "pdr_managedr_accepts_incomplete" {
  description = "A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning."
  type        = bool
  default     = true
}

// Resource arguments for pdr_validate_apikey
variable "pdr_validate_apikey_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_validate_apikey_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "accept_language"
}
variable "pdr_validate_apikey_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "if_none_match"
}

// Data source arguments for pdr_get_deployment_status
variable "pdr_get_deployment_status_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_get_deployment_status_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for pdr_get_event
variable "pdr_get_event_provision_id" {
  description = "provision id."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_get_event_event_id" {
  description = "Event ID."
  type        = string
  default     = "00116b2a-9326-4024-839e-fb5364b76898"
}
variable "pdr_get_event_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "placeholder"
}
variable "pdr_get_event_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for pdr_get_events
variable "pdr_get_events_provision_id" {
  description = "provision id."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_get_events_time" {
  description = "(deprecated - use from_time) A time in either ISO 8601 or unix epoch format."
  type        = string
  default     = "2025-06-19T23:59:59Z"
}
variable "pdr_get_events_from_time" {
  description = "A from query time in either ISO 8601 or unix epoch format."
  type        = string
  default     = "2025-06-19T00:00:00Z"
}
variable "pdr_get_events_to_time" {
  description = "A to query time in either ISO 8601 or unix epoch format."
  type        = string
  default     = "2025-06-19T23:59:59Z"
}
variable "pdr_get_events_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "placeholder"
}
variable "pdr_get_events_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for pdr_get_machine_types
variable "pdr_get_machine_types_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_get_machine_types_primary_workspace_name" {
  description = "Primary Workspace Name."
  type        = string
  default     = "Test-workspace-wdc06"
}
variable "pdr_get_machine_types_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "placeholder"
}
variable "pdr_get_machine_types_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}
variable "pdr_get_machine_types_standby_workspace_name" {
  description = "Standby Workspace Name."
  type        = string
  default     = "Test-workspace-wdc07"
}

// Data source arguments for pdr_get_managed_vm_list
variable "pdr_get_managed_vm_list_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_get_managed_vm_list_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "placeholder"
}
variable "pdr_get_managed_vm_list_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for pdr_last_operation
variable "pdr_last_operation_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_last_operation_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "placeholder"
}
variable "pdr_last_operation_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for pdr_validate_clustertype
variable "pdr_validate_clustertype_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_validate_clustertype_orchestrator_cluster_type" {
  description = "orchestrator cluster type value."
  type        = string
  default     = "on-premises"
}
variable "pdr_validate_clustertype_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "placeholder"
}
variable "pdr_validate_clustertype_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for pdr_validate_proxyip
variable "pdr_validate_proxyip_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_validate_proxyip_proxyip" {
  description = "proxyip value."
  type        = string
  default     = "10.30.40.5:3128"
}
variable "pdr_validate_proxyip_vpc_location" {
  description = "vpc location value."
  type        = string
  default     = "us-south"
}
variable "pdr_validate_proxyip_vpc_id" {
  description = "vpc id value."
  type        = string
  default     = "r006-2f3b3ab9-2149-49cc-83a1-30a5d93d59b2"
}
variable "pdr_validate_proxyip_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for pdr_validate_workspace
variable "pdr_validate_workspace_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_validate_workspace_workspace_id" {
  description = "standBy workspaceID value."
  type        = string
  default     = "75cbf05b-78f6-406e-afe7-a904f646d798"
}
variable "pdr_validate_workspace_crn" {
  description = "crn value."
  type        = string
  default     = "crn:v1:bluemix:public:power-iaas:dal10:a/094f4214c75941f991da601b001df1fe:75cbf05b-78f6-406e-afe7-a904f646d798::"
}
variable "pdr_validate_workspace_location_url" {
  description = "schematic_workspace_id value."
  type        = string
  default     = "https://us-south.power-iaas.cloud.ibm.com"
}
variable "pdr_validate_workspace_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for pdr_get_dr_summary_response
variable "pdr_get_dr_summary_response_instance_id" {
  description = "instance id of instance to provision."
  type        = string
  default     = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
variable "pdr_get_dr_summary_response_accept_language" {
  description = "The language requested for the return document."
  type        = string
  default     = "placeholder"
}
variable "pdr_get_dr_summary_response_if_none_match" {
  description = "ETag for conditional requests (optional)."
  type        = string
  default     = "placeholder"
}
