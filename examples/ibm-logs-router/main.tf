provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision logs_router_tenant resource instance
resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
  target_type = var.logs_router_tenant_target_type
  target_host = var.logs_router_tenant_target_host
  target_port = var.logs_router_tenant_target_port
  target_instance_crn = var.logs_router_tenant_target_instance_crn
}

// Create logs_router_tenant data source
data "ibm_logs_router_tenant" "logs_router_tenant_instance" {
  tenant_id = ibm_logs_router_tenant.logs_router_tenant_instance.id
}

