variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for logs_router_tenant
variable "logs_router_tenant_target_type" {
  description = "Type of log-sink."
  type        = string
  default     = "logdna"
}
variable "logs_router_tenant_target_host" {
  description = "Host name of log-sink."
  type        = string
  default     = "www.example.com"
}
variable "logs_router_tenant_target_port" {
  description = "Network port of log sink."
  type        = number
  default     = 10
}
variable "logs_router_tenant_target_instance_crn" {
  description = "Cloud resource name of the log-sink target instance."
  type        = string
  default     = "crn:v1:bluemix:public:logdna:us-east:a/36ff82794a734d7580b90c97b0327d28:f08aea7c-dde9-4452-b552-225af4b51eaa::"
}

// Data source arguments for logs_router_tenant
variable "logs_router_tenant_tenant_id" {
  description = "The instance ID of the tenant."
  type        = string
  default     = "f3a466c9-c4db-4eee-95cc-ba82db58e2b5"
}
