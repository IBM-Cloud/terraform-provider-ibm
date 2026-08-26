// This output allows brs_migration data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_brs_migration" {
  value       = ibm_brs_migration.brs_migration_instance
  description = "brs_migration resource instance"
}
// This output allows brs_migration_host data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_brs_migration_host" {
  value       = ibm_brs_migration_host.brs_migration_host_instance
  description = "brs_migration_host resource instance"
}
// This output allows brs_migration_volume data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_brs_migration_volume" {
  value       = ibm_brs_migration_volume.brs_migration_volume_instance
  description = "brs_migration_volume resource instance"
}
// This output allows brs_migration_workload data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_brs_migration_workload" {
  value       = ibm_brs_migration_workload.brs_migration_workload_instance
  description = "brs_migration_workload resource instance"
}
// This output allows brs_migration_workload_run data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_brs_migration_workload_run" {
  value       = ibm_brs_migration_workload_run.brs_migration_workload_run_instance
  description = "brs_migration_workload_run resource instance"
}
// This output allows brs_migration_discover data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_brs_migration_discover" {
  value       = ibm_brs_migration_discover.brs_migration_discover_instance
  description = "brs_migration_discover resource instance"
}
