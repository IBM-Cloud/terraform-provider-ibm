// This output allows logs_router_tenant data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_logs_router_tenant" {
  value       = ibm_logs_router_tenant.logs_router_tenant_instance
  description = "logs_router_tenant resource instance"
}

output "ibm_logs_router_tenants" {
  value       = data.ibm_logs_router_tenants.logs_router_tenants_instance
  description = "logs_router_tenants"
}

