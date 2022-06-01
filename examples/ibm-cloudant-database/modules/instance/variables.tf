#####################################################
# Cloudant Instance
# Copyright 2021 IBM
#####################################################

variable "provision" {
  description = "Enable this to provision the cloudant instance (true/false)"
  type        = bool
  default     = true
}

variable "provision_resource_key" {
  description = "Enable this to bind key to cloudant instance (true/false)"
  type        = bool
  default     = true
}

variable "instance_name" {
  description = "Name of the cloudant instance"
  type        = string
}

variable "plan" {
  description = "plan type (standard and lite)"
  type        = string
  default     = "standard"
}

variable "create_timeout" {
  type        = string
  description = "Timeout duration for create."
  default     = null
}

variable "update_timeout" {
  type        = string
  description = "Timeout duration for update."
  default     = null
}

variable "delete_timeout" {
  type        = string
  description = "Timeout duration for delete."
  default     = null
}

variable "resource_group_id" {
  description = "Enter resource group name"
  type        = string
}

variable "region" {
  description = "Provisioning Region"
  type        = string
}

variable "service_endpoints" {
  description = "Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'."
  type        = string
  default     = null
}

variable "resource_key_name" {
  description = "Name of the resource key"
  type        = string
}

variable "tags" {
  type        = list(string)
  description = "Tags that should be applied to the service"
  default     = null
}

variable "legacy_credentials" {
  description = "Legacy authentication method for cloudant"
  type        = bool
  default     = false
}

variable "role" {
  description = "Resource key role"
  type        = string
}

variable "resource_key_tags" {
  type        = list(string)
  description = "Tags that should be applied to the service"
  default     = null
}

#####################################################
# Service Policy Configuration
#####################################################
variable "service_policy_provision" {
  description = "Enable this to provision the service policy (true/false)"
  type        = bool
}

variable "service_name" {
  description = "Name of the service ID"
  type        = string
}

variable "description" {
  description = "Description to service ID"
  type        = string
  default     = null
}

variable "roles" {
  description = "service policy roles"
  type        = list(string)
}
