// This output allows pdr_managedr data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_pdr_managedr" {
  value       = ibm_pdr_managedr.pdr_managedr_instance
  description = "pdr_managedr resource instance"
}
// This output allows pdr_validate_apikey data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_pdr_validate_apikey" {
  value       = ibm_pdr_validate_apikey.pdr_validate_apikey_instance
  description = "pdr_validate_apikey resource instance"
}
