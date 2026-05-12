// This output allows logs_router_target data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_router_target" {
  value       = ibm_logs_router_target.logs_router_target_instance
  description = "logs_router_target resource instance"
}
// This output allows logs_router_route data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_router_route" {
  value       = ibm_logs_router_route.logs_router_route_instance
  description = "logs_router_route resource instance"
}
// This output allows logs_router_settings data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_router_settings" {
  value       = ibm_logs_router_settings.logs_router_settings_instance
  description = "logs_router_settings resource instance"
}
