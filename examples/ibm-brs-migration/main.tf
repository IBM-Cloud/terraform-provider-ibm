provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision brs_migration resource instance
resource "ibm_brs_migration" "brs_migration_instance" {
  name = var.brs_migration_name
  brs_crn = var.brs_migration_brs_crn
  description = var.brs_migration_description
}

// Provision brs_migration_host resource instance
resource "ibm_brs_migration_host" "brs_migration_host_instance" {
  migration_id = var.brs_migration_host_migration_id
  type = var.brs_migration_host_type
  env = var.brs_migration_host_env
}

// Provision brs_migration_volume resource instance
resource "ibm_brs_migration_volume" "brs_migration_volume_instance" {
  migration_id = var.brs_migration_volume_migration_id
  env = var.brs_migration_volume_env
  storage_type = var.brs_migration_volume_storage_type
}

// Provision brs_migration_workload resource instance
resource "ibm_brs_migration_workload" "brs_migration_workload_instance" {
  migration_id = var.brs_migration_workload_migration_id
  name = var.brs_migration_workload_name
  payloads {
    id = "pl-c3d4e5f6-a7b8-9012-cdef-012345678901"
    source_host_id = "host-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
    destination_host_id = "host-b2c3d4e5-f6a7-8901-bcde-f01234567890"
    data_specs {
      data_format = "raw"
      source {
        volume_id = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
        type = "ext4"
        path = "/mnt/data"
      }
      destination {
        volume_id = "vol-b1c2d3e4-f5a6-7890-bcde-f01234567890"
        type = "ext4"
        path = "/mnt/data"
      }
    }
  }
}

// Provision brs_migration_workload_run resource instance
resource "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
  migration_id = var.brs_migration_workload_run_migration_id
  workload_id = var.brs_migration_workload_run_workload_id
}

// Provision brs_migration_discover resource instance
resource "ibm_brs_migration_discover" "brs_migration_discover_instance" {
  migration_id = var.brs_migration_discover_migration_id
  env = var.brs_migration_discover_env
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration data source
data "ibm_brs_migration" "brs_migration_instance" {
  migration_id = var.data_brs_migration_migration_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migrations data source
data "ibm_brs_migrations" "brs_migrations_instance" {
  state = var.brs_migrations_state
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_host data source
data "ibm_brs_migration_host" "brs_migration_host_instance" {
  migration_id = var.data_brs_migration_host_migration_id
  host_id = var.data_brs_migration_host_host_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_hosts data source
data "ibm_brs_migration_hosts" "brs_migration_hosts_instance" {
  migration_id = var.brs_migration_hosts_migration_id
  env = var.brs_migration_hosts_env
  type = var.brs_migration_hosts_type
  migrated = var.brs_migration_hosts_migrated
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_volume data source
data "ibm_brs_migration_volume" "brs_migration_volume_instance" {
  migration_id = var.data_brs_migration_volume_migration_id
  volume_id = var.data_brs_migration_volume_volume_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_volumes data source
data "ibm_brs_migration_volumes" "brs_migration_volumes_instance" {
  migration_id = var.brs_migration_volumes_migration_id
  env = var.brs_migration_volumes_env
  storage_type = var.brs_migration_volumes_storage_type
  migrated = var.brs_migration_volumes_migrated
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_workload data source
data "ibm_brs_migration_workload" "brs_migration_workload_instance" {
  migration_id = var.data_brs_migration_workload_migration_id
  workload_id = var.data_brs_migration_workload_workload_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_workloads data source
data "ibm_brs_migration_workloads" "brs_migration_workloads_instance" {
  migration_id = var.brs_migration_workloads_migration_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_workload_run data source
data "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
  migration_id = var.data_brs_migration_workload_run_migration_id
  workload_id = var.data_brs_migration_workload_run_workload_id
  run_id = var.data_brs_migration_workload_run_run_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_workload_runs data source
data "ibm_brs_migration_workload_runs" "brs_migration_workload_runs_instance" {
  migration_id = var.brs_migration_workload_runs_migration_id
  workload_id = var.brs_migration_workload_runs_workload_id
  status = var.brs_migration_workload_runs_status
  run_type = var.brs_migration_workload_runs_run_type
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_workload_history data source
data "ibm_brs_migration_workload_history" "brs_migration_workload_history_instance" {
  migration_id = var.brs_migration_workload_history_migration_id
  workload_id = var.brs_migration_workload_history_workload_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_discover data source
data "ibm_brs_migration_discover" "brs_migration_discover_instance" {
  migration_id = var.data_brs_migration_discover_migration_id
  job_id = var.data_brs_migration_discover_job_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create brs_migration_discovers data source
data "ibm_brs_migration_discovers" "brs_migration_discovers_instance" {
  migration_id = var.brs_migration_discovers_migration_id
  env = var.brs_migration_discovers_env
  state = var.brs_migration_discovers_state
}
*/
