// This output allows iam_identity_preference data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_identity_preference" {
  value       = data.ibm_iam_identity_preference.iam_identity_preference_instance_data
  description = "iam_identity_preference data"
}

output "ibm_iam_identity_preferences" {
  value       = data.ibm_iam_identity_preferences.iam_identity_preferences_instance_list
  description = "iam_identity_preferences list"
}