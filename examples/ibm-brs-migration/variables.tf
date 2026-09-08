variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for brs_migration
variable "brs_migration_name" {
  description = "Human-readable name for this migration project."
  type        = string
  default     = "prod-classic-to-vpc"
}
variable "brs_migration_brs_crn" {
  description = "CRN of the IBM Cloud Backup and Recovery instance backing this migration."
  type        = string
  default     = "crn:v1:bluemix:public:backup-recovery:us-south:a/123456:abcdef01-2345-6789-abcd-ef0123456789::"
}
variable "brs_migration_description" {
  description = "Optional human-readable description."
  type        = string
  default     = "Migrate production Classic workloads to VPC"
}

// Resource arguments for brs_migration_host
variable "brs_migration_host_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_host_type" {
  description = "Whether the host is a Virtual Server Instance or bare metal server."
  type        = string
  default     = "virtual_server"
}
variable "brs_migration_host_env" {
  description = "Infrastructure environment this host belongs to."
  type        = string
  default     = "classic"
}

// Resource arguments for brs_migration_volume
variable "brs_migration_volume_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_volume_env" {
  description = "Infrastructure environment this volume belongs to."
  type        = string
  default     = "vpc"
}
variable "brs_migration_volume_storage_type" {
  description = "Storage type of the volume."
  type        = string
  default     = "block"
}

// Resource arguments for brs_migration_workload
variable "brs_migration_workload_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_workload_name" {
  description = "Human-readable name for this workload."
  type        = string
  default     = "prod-to-vpc-migration"
}

// Resource arguments for brs_migration_workload_run
variable "brs_migration_workload_run_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_workload_run_workload_id" {
  description = "The migration service workload ID (wl-{uuid4} format)."
  type        = string
  default     = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}

// Resource arguments for brs_migration_discover
variable "brs_migration_discover_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_discover_env" {
  description = "Infrastructure environment being discovered."
  type        = string
  default     = "classic"
}

// Data source arguments for brs_migration
variable "data_brs_migration_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}

// Data source arguments for brs_migrations
variable "brs_migrations_state" {
  description = "Filter by migration state."
  type        = string
  default     = "placeholder"
}

// Data source arguments for brs_migration_host
variable "data_brs_migration_host_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "data_brs_migration_host_host_id" {
  description = "The migration service host ID (host-{uuid4} format)."
  type        = string
  default     = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}

// Data source arguments for brs_migration_hosts
variable "brs_migration_hosts_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_hosts_env" {
  description = "Filter by infrastructure environment."
  type        = string
  default     = "placeholder"
}
variable "brs_migration_hosts_type" {
  description = "Filter by compute type."
  type        = string
  default     = "placeholder"
}
variable "brs_migration_hosts_migrated" {
  description = "Filter by migration status."
  type        = bool
  default     = false
}

// Data source arguments for brs_migration_volume
variable "data_brs_migration_volume_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "data_brs_migration_volume_volume_id" {
  description = "The migration service volume ID (vol-{uuid4} format)."
  type        = string
  default     = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
}

// Data source arguments for brs_migration_volumes
variable "brs_migration_volumes_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_volumes_env" {
  description = "Filter by infrastructure environment."
  type        = string
  default     = "placeholder"
}
variable "brs_migration_volumes_storage_type" {
  description = "Filter by storage type."
  type        = string
  default     = "placeholder"
}
variable "brs_migration_volumes_migrated" {
  description = "Filter by migration status."
  type        = bool
  default     = false
}

// Data source arguments for brs_migration_workload
variable "data_brs_migration_workload_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "data_brs_migration_workload_workload_id" {
  description = "The migration service workload ID (wl-{uuid4} format)."
  type        = string
  default     = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}

// Data source arguments for brs_migration_workloads
variable "brs_migration_workloads_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}

// Data source arguments for brs_migration_workload_run
variable "data_brs_migration_workload_run_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "data_brs_migration_workload_run_workload_id" {
  description = "The migration service workload ID (wl-{uuid4} format)."
  type        = string
  default     = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
variable "data_brs_migration_workload_run_run_id" {
  description = "The migration service run ID (run-{uuid4} format)."
  type        = string
  default     = "run-e5f6a7b8-c9d0-1234-ef01-234567890123"
}

// Data source arguments for brs_migration_workload_runs
variable "brs_migration_workload_runs_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_workload_runs_workload_id" {
  description = "The migration service workload ID (wl-{uuid4} format)."
  type        = string
  default     = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
variable "brs_migration_workload_runs_status" {
  description = "Filter by run status."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "brs_migration_workload_runs_run_type" {
  description = "Filter by run type (scheduled or on-demand)."
  type        = string
  default     = "placeholder"
}

// Data source arguments for brs_migration_workload_history
variable "brs_migration_workload_history_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_workload_history_workload_id" {
  description = "The migration service workload ID (wl-{uuid4} format)."
  type        = string
  default     = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}

// Data source arguments for brs_migration_discover
variable "data_brs_migration_discover_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "data_brs_migration_discover_job_id" {
  description = "The unique ID of the discovery job (job-{uuid4} format)."
  type        = string
  default     = "job-12345678-abcd-ef01-2345-678901234567"
}

// Data source arguments for brs_migration_discovers
variable "brs_migration_discovers_migration_id" {
  description = "The migration project ID (mgr-{uuid4} format)."
  type        = string
  default     = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
variable "brs_migration_discovers_env" {
  description = "Filter discovery jobs by infrastructure environment."
  type        = string
  default     = "placeholder"
}
variable "brs_migration_discovers_state" {
  description = "Filter discovery jobs by current state."
  type        = string
  default     = "placeholder"
}
