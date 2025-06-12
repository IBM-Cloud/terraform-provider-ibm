// This output allows trusted_profile_template data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_trusted_profile_template" {
  value       = ibm_iam_trusted_profile_template.trusted_profile_template_instance
  description = "trusted_profile_template resource instance"
}

output "ibm_iam_trusted_profile_template_version" {
  value       = ibm_iam_trusted_profile_template.trusted_profile_template_version
  description = "trusted_profile_template resource instance"
}

