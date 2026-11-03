// This output allows sm_destination data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_sm_destination" {
  value       = ibm_sm_destination.sm_destination_instance
  description = "sm_destination resource instance"
}
