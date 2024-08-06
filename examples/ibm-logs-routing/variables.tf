variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for logs_router_tenant
variable "logs_router_tenant_name" {
  description = "The name for this tenant. The name is regionally unique across all tenants in the account."
  type        = string
  default     = "my-logging-tenant"
}
