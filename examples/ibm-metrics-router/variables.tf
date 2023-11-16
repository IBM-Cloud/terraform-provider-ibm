variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for metrics_router_target
variable "metrics_router_target_name" {
  description = "The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names."
  type        = string
  default     = "my-mr-target"
}
variable "metrics_router_target_destination_crn" {
  description = "The CRN of a destination service instance or resource. Ensure you have a service authorization between IBM Cloud Metrics Routing and your Cloud resource. Read [S2S authorization](https://cloud.ibm.com/docs/metrics-router?topic=metrics-router-target-monitoring&interface=ui#target-monitoring-ui) for details."
  type        = string
  default     = "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
}
variable "metrics_router_target_region" {
  description = "Include this optional field if you want to create a target in a different region other than the one you are connected."
  type        = string
  default     = "us-south"
}

// Resource arguments for metrics_router_route
variable "metrics_router_route_name" {
  description = "The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names."
  type        = string
  default     = "my-route"
}

// Resource arguments for metrics_router_settings
variable "metrics_router_settings_permitted_target_regions" {
  description = "If present then only these regions may be used to define a target."
  type        = list(string)
  default     = [ "us-south" ]
}
variable "metrics_router_settings_primary_metadata_region" {
  description = "To store all your meta data in a single region."
  type        = string
  default     = "us-south"
}
variable "metrics_router_settings_backup_metadata_region" {
  description = "To backup all your meta data in a different region."
  type        = string
  default     = "us-east"
}
variable "metrics_router_settings_private_api_endpoint_only" {
  description = "If you set this true then you cannot access api through public network."
  type        = bool
  default     = false
}

// Data source arguments for metrics_router_targets
variable "metrics_router_targets_name" {
  description = "The name of the target resource."
  type        = string
  default     = "a-mr-target-us-south"
}

// Data source arguments for metrics_router_routes
variable "metrics_router_routes_name" {
  description = "The name of the route."
  type        = string
  default     = "my-route"
}
