// This allows scc_template data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_scc_template" {
  value       = ibm_scc_template.scc_template_instance
  description = "scc_template resource instance"
}
// This allows scc_template_attachment data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_scc_template_attachment" {
  value       = ibm_scc_template_attachment.scc_template_attachment_instance
  description = "scc_template_attachment resource instance"
} 
// This allows scc_rule data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_scc_rule" {
  value       = ibm_scc_rule.scc_rule_instance
  description = "scc_rule resource instance"
}
// This allows scc_rule_attachment data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_scc_rule_attachment" {
  value       = ibm_scc_rule_attachment.scc_rule_attachment_instance
  description = "scc_rule_attachment resource instance"
}
