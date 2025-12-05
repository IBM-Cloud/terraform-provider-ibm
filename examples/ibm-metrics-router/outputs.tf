// This output allows metrics_router_target data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_metrics_router_target" {
  value       = ibm_metrics_router_target.metrics_router_target_instance
  description = "metrics_router_target resource instance"
}
// This output allows metrics_router_route data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_metrics_router_route" {
  value       = ibm_metrics_router_route.metrics_router_route_instance
  description = "metrics_router_route resource instance"
}
// This output allows metrics_router_settings data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_metrics_router_settings" {
  value       = ibm_metrics_router_settings.metrics_router_settings_instance
  description = "metrics_router_settings resource instance"
}
