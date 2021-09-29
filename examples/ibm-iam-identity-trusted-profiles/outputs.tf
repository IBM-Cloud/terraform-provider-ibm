// This allows iam_trusted_profile data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_trusted_profile" {
  value       = ibm_iam_trusted_profile.iam_trusted_profile_instance
  description = "iam_trusted_profile resource instance"
}
