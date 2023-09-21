variable "scc_provider_type_id" {
  description = "The provider type ID."
  type        = string
  default     = "INSERT VALID INSTANCE ID"
}

variable "scc_provider_type_instance_name" {
  description = "The name of the provider type instance."
  type        = string
  default     = "workload-protection-instance-1"
}

variable "scc_provider_type_instance_attributes" {
  description = "The provider type instance attributes"
  type        = map
  default     = {}
}
