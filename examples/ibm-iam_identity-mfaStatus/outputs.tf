// This allows iam_user_mfa_enrollments data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed



// for user mfa enrollments list operation
output "ibm_iam_user_mfa_enrollments" {
  value       = data.ibm_iam_user_mfa_enrollments.iam_user_mfa_enrollments_instance
  description = "iam_user_mfa_enrollments datasource instance"
}