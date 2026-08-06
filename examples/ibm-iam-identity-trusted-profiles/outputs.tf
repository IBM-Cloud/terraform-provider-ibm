// This allows iam_trusted_profile data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed

output "ibm_iam_trusted_profile" {
  value       = data.ibm_iam_trusted_profile.iam_trusted_profile_instance_data
  description = "iam_trusted_profile resource instance"
}

// for trusted profile list operation
output "ibm_iam_trusted_profiles_list" {
  value       = data.ibm_iam_trusted_profiles.iam_trusted_profiles_list_data
  description = "iam_trusted_profiles list"
}
