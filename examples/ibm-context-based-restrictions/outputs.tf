// This allows cbr_zone data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cbr_zone" {
  value       = ibm_cbr_zone.cbr_zone_instance
  description = "cbr_zone resource instance"
}
// This allows cbr_rule data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cbr_rule" {
  value       = ibm_cbr_rule.cbr_rule_instance
  description = "cbr_rule resource instance"
}
