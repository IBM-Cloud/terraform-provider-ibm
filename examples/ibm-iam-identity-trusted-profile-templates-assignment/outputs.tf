// This output allows trusted_profile_template_assignment data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_trusted_profile_template_assignment" {
  value       = ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance
  description = "trusted_profile_template_assignment resource instance"
}
