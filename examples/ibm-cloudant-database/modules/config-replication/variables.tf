#####################################################
# IBMCLOUD Cloudant Replication Variables
#####################################################

variable "cloudant_guid" {
  description = "Cloudant instance GUID."
  type        = string
}

variable "db_name" {
  type        = string
  description = "Database name"
}

variable "cloudant_database_partitioned" {
  description = "Query parameter to specify whether to enable database partitions when creating a database."
  type        = bool
}

variable "cloudant_database_q" {
  description = "The number of shards in the database. Each shard is a partition of the hash value range. Default is 8, unless overridden in the `cluster config`."
  type        = number
}

variable "cloudant_replication_doc_id" {
  description = "Path parameter to specify the document ID."
  type        = string
}

variable "source_api_key" {
  description = "HTTP request body for replication operations."
  type        = string
}

variable "target_api_key" {
  description = "HTTP request body for replication operations."
  type        = string
}

variable "source_host" {
  description = "HTTP request body for replication operations."
  type        = string
}

variable "target_host" {
  description = "HTTP request body for replication operations."
  type        = string
}

variable "create_target" {
  description = "Creates the target database. Requires administrator privileges on target server."
}

variable "continuous" {
  description = "Configure the replication to be continuous."
  type        = bool
}
