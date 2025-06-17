// This output allows sds_volume data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed

output "ibm_sds_volume" {
  value       = [ibm_sds_volume.sds_volume_instance_1, ibm_sds_volume.sds_volume_instance_2]
  description = "sds_volume resource instance"
}
// This output allows sds_host data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sds_host" {
  value       = ibm_sds_host.sds_host_instance
  description = "sds_host resource instance"
}
