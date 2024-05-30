// This output allows common_source_registration_request data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_common_source_registration_request" {
  value       = ibm_common_source_registration_request.common_source_registration_request_instance
  description = "common_source_registration_request resource instance"
}
