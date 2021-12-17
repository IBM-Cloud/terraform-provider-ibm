// This allows iam_trusted_profile data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed



// for trusted profile list operation
output "ibm_iam_trusted_profiles" {
  value       = data.ibm_iam_trusted_profiles.iam_trusted_profiles_instance
  description = "iam_trusted_profiles resource instance"
}
