variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for logs_router_target
variable "logs_router_target_name" {
  description = "The name of the target resource."
  type        = string
  default     = "my-lr-target"
}
variable "logs_router_target_destination_crn" {
  description = "Cloud Resource Name (CRN) of the destination resource. Ensure you have a service authorization between IBM Cloud Logs Routing and your Cloud resource. See [service-to-service authorization](https://cloud.ibm.com/docs/logs-router?topic=logs-router-target-monitoring&interface=ui#target-monitoring-ui) for details."
  type        = string
  default     = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
}
variable "logs_router_target_region" {
  description = "Include this optional field if you used it to create a target in a different region other than the one you are connected."
  type        = string
  default     = "us-south"
}
variable "logs_router_target_managed_by" {
  description = "Present when the target is enterprise-managed (`managed_by: enterprise`). For account-managed targets this field is omitted."
  type        = string
  default     = "enterprise"
}

// Resource arguments for logs_router_route
variable "logs_router_route_name" {
  description = "The name of the route."
  type        = string
  default     = "my-route"
}
variable "logs_router_route_managed_by" {
  description = "Present when the route is enterprise-managed (`managed_by: enterprise`)."
  type        = string
  default     = "enterprise"
}

// Resource arguments for logs_router_settings
variable "logs_router_settings_permitted_target_regions" {
  description = "If present then only these regions may be used to define a target."
  type        = list(string)
  default     = [ "us-south" ]
}
variable "logs_router_settings_primary_metadata_region" {
  description = "To store all your meta data in a single region."
  type        = string
  default     = "us-south"
}
variable "logs_router_settings_backup_metadata_region" {
  description = "To backup all your meta data in a different region."
  type        = string
  default     = "us-east"
}
variable "logs_router_settings_private_api_endpoint_only" {
  description = "If you set this true then you cannot access api through public network."
  type        = bool
  default     = false
}

// Data source arguments for logs_router_targets
variable "logs_router_targets_name" {
  description = "The name of the target resource."
  type        = string
  default     = "a-lr-target-us-south"
}

// Data source arguments for logs_router_routes
variable "logs_router_routes_name" {
  description = "The name of the route."
  type        = string
  default     = "my-route"
}
