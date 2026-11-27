// This output allows sm_destination data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sm_destination" {
  value       = ibm_sm_destination.sm_destination_instance
  description = "sm_destination resource instance"
}
// This output allows sm_hsm_destination data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sm_hsm_destination" {
  value       = ibm_sm_hsm_destination.sm_hsm_destination_instance
  description = "sm_hsm_destination resource instance"
}
// This output allows sm_hsm_ha_group data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sm_hsm_ha_group" {
  value       = ibm_sm_hsm_ha_group.sm_hsm_ha_group_instance
  description = "sm_hsm_ha_group resource instance"
}
// This output allows sm_hsm_path data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sm_hsm_path" {
  value       = ibm_sm_hsm_path.sm_hsm_path_instance
  description = "sm_hsm_path resource instance"
}
// This output allows sm_hsm_path_build data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sm_hsm_path_build" {
  value       = ibm_sm_hsm_path_build.sm_hsm_path_build_instance
  description = "sm_hsm_path_build resource instance"
}
// This output allows sm_hsm_server_ca data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sm_hsm_server_ca" {
  value       = ibm_sm_hsm_server_ca.sm_hsm_server_ca_instance
  description = "sm_hsm_server_ca resource instance"
}
// This output allows sm_hsm_client_certificate data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sm_hsm_client_certificate" {
  value       = ibm_sm_hsm_client_certificate.sm_hsm_client_certificate_instance
  description = "sm_hsm_client_certificate resource instance"
}
