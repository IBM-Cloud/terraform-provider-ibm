// This output allows logs_alert data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_alert" {
  value       = ibm_logs_alert.logs_alert_instance
  description = "logs_alert resource instance"
}
// This output allows logs_rule_group data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_rule_group" {
  value       = ibm_logs_rule_group.logs_rule_group_instance
  description = "logs_rule_group resource instance"
}
// This output allows logs_outgoing_webhook data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_outgoing_webhook" {
  value       = ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance
  description = "logs_outgoing_webhook resource instance"
}
// This output allows logs_policy data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_policy" {
  value       = ibm_logs_policy.logs_policy_instance
  description = "logs_policy resource instance"
}
// This output allows logs_dashboard data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_dashboard" {
  value       = ibm_logs_dashboard.logs_dashboard_instance
  description = "logs_dashboard resource instance"
}
// This output allows logs_dashboard_folder data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_dashboard_folder" {
  value       = ibm_logs_dashboard_folder.logs_dashboard_folder_instance
  description = "logs_dashboard_folder resource instance"
}
// This output allows logs_e2m data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_e2m" {
  value       = ibm_logs_e2m.logs_e2m_instance
  description = "logs_e2m resource instance"
}
// This output allows logs_view data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_view" {
  value       = ibm_logs_view.logs_view_instance
  description = "logs_view resource instance"
}
// This output allows logs_view_folder data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_view_folder" {
  value       = ibm_logs_view_folder.logs_view_folder_instance
  description = "logs_view_folder resource instance"
}
// This output allows logs_data_access_rule data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_data_access_rule" {
  value       = ibm_logs_data_access_rule.logs_data_access_rule_instance
  description = "logs_data_access_rule resource instance"
}
// This output allows logs_enrichment data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_enrichment" {
  value       = ibm_logs_enrichment.logs_enrichment_instance
  description = "logs_enrichment resource instance"
}
// This output allows logs_data_usage_metrics data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_data_usage_metrics" {
  value       = ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance
  description = "logs_data_usage_metrics resource instance"
}
// This output allows logs_stream data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_stream" {
  value       = ibm_logs_stream.logs_stream_instance
  description = "logs_stream resource instance"
}
// This output allows logs_alert_definition data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_alert_definition" {
  value       = ibm_logs_alert_definition.logs_alert_definition_instance
  description = "logs_alert_definition resource instance"
}
