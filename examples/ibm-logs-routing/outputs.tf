// This output allows logs-router_tenant data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs-router_tenant" {
  value       = ibm_logs-router_tenant.logs-router_tenant_instance
  description = "logs-router_tenant resource instance"
}
