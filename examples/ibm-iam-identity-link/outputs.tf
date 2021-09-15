// This allows iam_trusted_profiles_link data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_trusted_profiles_link" {
  value       = ibm_iam_trusted_profiles_link.iam_trusted_profiles_link_instance
  description = "iam_trusted_profiles_link resource instance"
}
