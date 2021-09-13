
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

variable "pri_region" {
  description = "Provisioning Region for primary instance"
  type        = string
  default     = "us-east"
}

variable "dr_region" {
  description = "Provisioning Region for DR instance"
  type        = string
  default     = "us-east"
}

variable "pri_instance_name" {
  description = "Name of the cloudant instance for primary"
  type        = string
}

variable "dr_instance_name" {
  description = "Name of the cloudant instance for DR"
  type        = string
}

variable "pri_resource_key" {
  description = "Name of the resource key of the primary instance"
  type        = string
}

variable "dr_resource_key" {
  description = "Name of the resource key of the DR"
  type        = string
}

variable "pri_rg_name" {
  type        = string
  description = "Enter resource group name for primary instance"
}

variable "dr_rg_name" {
  type        = string
  description = "Enter resource group name for disaster recovery"
}

variable "legacy_credentials" {
  description = "Legacy authentication method for cloudant"
  type        = bool
  default     = null
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

variable "is_dr_provision" {
  type        = bool
  description = "Would you like to provision a DR cloudant instance (true/false)"
  default     = true
}

variable "db_name" {
  type        = string
  description = "Database name"
}

variable "is_partitioned" {
  description = "To set whether the database is partitioned"
  default     = false
}

variable "cloudant_database_q" {
  description = "The number of shards in the database. Each shard is a partition of the hash value range. Default is 8, unless overridden in the `cluster config`."
  type        = number
  default     = 0
}

#####################################################
# IBMCLOUD Cloudant Database Variables
#####################################################

variable "cloudant_replication_doc_id" {
  description = "Path parameter to specify the document ID."
  type        = string
  default     = "doc_id"
}

variable "create_target" {
  description = "Creates the target database. Requires administrator privileges on target server."
  type        = bool
  default     = false
}

variable "continuous" {
  description = "Configure the replication to be continuous."
  type        = bool
  default     = true
}