

variable "cloudant_guid" {
  description = "The cloudant instance GUID"
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

variable "cloudant_database_q" {
  description = "The number of shards in the database. Each shard is a partition of the hash value range. Default is 8, unless overridden in the `cluster config`."
  type        = number
  default     = 0
}