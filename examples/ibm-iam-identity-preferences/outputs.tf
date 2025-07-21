// This output allows iam_identity_preference data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_identity_preference" {
  value       = ibm_iam_identity_preference.iam_identity_preference_instance
  description = "iam_identity_preference resource instance"
}
