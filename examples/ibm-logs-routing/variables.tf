variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for logs-router_tenant
variable "logs-router_tenant_ibm_api_version" {
  description = "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version."
  type        = string
  default     = "ibm_api_version"
}
variable "logs-router_tenant_name" {
  description = "The name for this tenant. The name is regionally unique across all tenants in the account."
  type        = string
  default     = "my-logging-tenant"
}

// Data source arguments for logs-router_tenant
variable "data_logs-router_tenant_ibm_api_version" {
  description = "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version."
  type        = string
  default     = "ibm_api_version"
}
variable "data_logs-router_tenant_tenant_id" {
  description = "The instance ID of the tenant."
  type        = string
  default     = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}

// Data source arguments for logs-router_tenants
variable "logs-router_tenants_ibm_api_version" {
  description = "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version."
  type        = string
  default     = "ibm_api_version"
}
variable "logs-router_tenants_name" {
  description = "Optional: The name of a tenant."
  type        = string
  default     = "placeholder"
}

// Data source arguments for logs-router_targets
variable "logs-router_targets_ibm_api_version" {
  description = "Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version."
  type        = string
  default     = "ibm_api_version"
}
variable "logs-router_targets_tenant_id" {
  description = "The instance ID of the tenant."
  type        = string
  default     = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}
variable "logs-router_targets_name" {
  description = "Optional: Name of the tenant target."
  type        = string
  default     = "placeholder"
}
