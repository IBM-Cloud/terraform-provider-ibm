// This output allows iam_access_group_template data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_access_group_template" {
  value       = ibm_iam_access_group_template.iam_access_group_template_instance
  description = "iam_access_group_template resource instance"
}
// This output allows iam_access_group_template_version data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_access_group_template_version" {
  value       = ibm_iam_access_group_template_version.iam_access_group_template_version_instance
  description = "iam_access_group_template_version resource instance"
}
// This output allows iam_access_group_template_assignment data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_access_group_template_assignment" {
  value       = ibm_iam_access_group_template_assignment.iam_access_group_template_assignment_instance
  description = "iam_access_group_template_assignment resource instance"
}
