// This allows iam_trusted_profile_link data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_trusted_profile_link" {
  value       = ibm_iam_trusted_profile_link.iam_trusted_profile_link_instance
  description = "iam_trusted_profile_link resource instance"
}
