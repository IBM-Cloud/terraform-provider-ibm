// This allows iam_trusted_profile_claim_rule data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed

// for claim rule CRUD operations 
output "ibm_iam_trusted_profile_claim_rule" {
  value       = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule_instance
  description = "iam_trusted_profile_claim_rule resource instance"
}

// for claim rule list operation
output "ibm_iam_trusted_profile_claim_rules" {
  value       = data.ibm_iam_trusted_profile_claim_rules.iam_trusted_profile_claim_rules_instance
  description = "iam_trusted_profile_claim_rules resource instance"
}