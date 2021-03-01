// This allows iam_policy data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_policy" {
  value       = ibm_iam_policy.iam_policy_instance
  description = "iam_policy resource instance"
}
// This allows iam_custom_role data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_custom_role" {
  value       = ibm_iam_custom_role.iam_custom_role_instance
  description = "iam_custom_role resource instance"
}
