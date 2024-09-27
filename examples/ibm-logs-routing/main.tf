provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "us-south"
}

// Provision logs_router_tenant resource instance
resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
  name = var.logs_router_tenant_name
  targets {
    log_sink_crn = "crn:v1:bluemix:public:logdna:eu-de:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-logdna-target"
    parameters {
      host = "www.example-1.com"
      port = 80
      access_credential = "new-cred"
    }
  }
    targets {
    log_sink_crn = "crn:v1:bluemix:public:logs:eu-de:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-cloud-logs-target"
    parameters {
      host = "www.example-2.com"
      port = 80
    }
  }
}

// Create logs_router_tenants data source
data "ibm_logs_router_tenants" "logs_router_tenants_instance" {
  name = ibm_logs_router_tenant.logs_router_tenant_instance_both.name
}

// Create logs_router_targets data source
data "ibm_logs_router_targets" "logs_router_targets_instance" {
  tenant_id = ibm_logs_router_tenant.logs_router_tenant_instance_both.id
}