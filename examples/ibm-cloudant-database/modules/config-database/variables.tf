

variable "cloudant_instance_crn" {
  description = "The cloudant instance CRN"
  type        = string
}

variable "db_name" {
  type        = string
  description = "Database name"
}

variable "cloudant_database_partitioned" {
  description = "Query parameter to specify whether to enable database partitions when creating a database."
  type        = bool
  default     = false
}

variable "cloudant_database_shards" {
  description = "The number of shards in the database. Each shard is a partition of the hash value range. When omitted the default is set by the server."
  type        = number
  default     = null
}