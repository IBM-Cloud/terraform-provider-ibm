// This allows iam_trusted_profile_claim_rule data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_trusted_profile_claim_rule" {
  value       = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule_instance
  description = "iam_trusted_profile_claim_rule resource instance"
}
