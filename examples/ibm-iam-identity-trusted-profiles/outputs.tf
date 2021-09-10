// This allows iam_trusted_profiles data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_trusted_profiles" {
  value       = ibm_iam_trusted_profiles.iam_trusted_profiles_instance
  description = "iam_trusted_profiles resource instance"
}
