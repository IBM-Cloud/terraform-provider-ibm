// This output allows iam_trusted_profile_identities data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_trusted_profile_identities" {
  value       = ibm_iam_trusted_profile_identities.iam_trusted_profile_identities_instance
  description = "iam_trusted_profile_identities resource instance"
}
