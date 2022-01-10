
#####################################################
# Cloudant - database
# Copyright 2021 IBM
#####################################################

#####################################################
# IBMCLOUD Cloudant Instance Variables
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

variable "region" {
  description = "Provisioning Region for the instance"
  type        = string
  default     = "us-south"
}

variable "instance_name" {
  description = "Name of the cloudant instance"
  type        = string
}

variable "resource_key" {
  description = "Name of the resource key of the instance"
  type        = string
}

variable "rg_name" {
  type        = string
  description = "Enter resource group name for the cloudant instance"
}

variable "legacy_credentials" {
  description = "Legacy authentication method for cloudant"
  type        = bool
  default     = false
}

variable "plan" {
  description = "plan type (standard and lite)"
  type        = string
  default     = "standard"
}

variable "service_endpoints" {
  description = "Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'."
  type        = string
  default     = null
}

variable "tags" {
  type        = list(string)
  description = "Tags that should be applied to the service"
  default     = null
}

variable "role" {
  type        = string
  description = "Resource key role"
  default     = "Writer"
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
  default     = true
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
  default     = ["Writer"]
}


#####################################################
# IBMCLOUD Cloudant Database Variables
#####################################################

variable "db_name" {
  type        = string
  description = "Database name"
}

variable "is_partitioned" {
  description = "To set whether the database is partitioned"
  default     = false
}

variable "cloudant_database_shards" {
  description = "The number of shards in the database. Each shard is a partition of the hash value range. When omitted the default is set by the server."
  type        = number
  default     = null
}
