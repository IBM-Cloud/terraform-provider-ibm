// This allows iam_trusted_profile_identity data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed

// for trusted profile identity instance output
output "ibm_iam_trusted_profile_identity_instance" {
  value       = ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance
  description = "iam_trusted_profile_identity instance"
}

// for trusted profile identity data output
output "ibm_iam_trusted_profile_identity_data" {
  value       = data.ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_data
  description = "iam_trusted_profile_identity data"
}

// for trusted profile identity list operation
output "ibm_iam_trusted_profile_identities_data" {
  value       = data.ibm_iam_trusted_profile_identities.iam_trusted_profile_identities_data
  description = "iam_trusted_profiles identities list"
}
