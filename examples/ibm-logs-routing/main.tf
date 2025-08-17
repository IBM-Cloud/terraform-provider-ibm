provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "us-east"
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "eu-de"
  alias = "provider-eu-de"
}

// Provision logs_router_tenant resource instance
resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
  name = var.logs_router_tenant_name
  region = "us-east"
    targets {
    log_sink_crn = "crn:v1:bluemix:public:logs:us-east:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-cloud-logs-target"
    parameters {
      host = "www.example-2.com"
      port = 443
    }
  }
}

resource "ibm_logs_router_tenant" "logs_router_tenant_instance_eu_de" {
  provider = ibm.provider-eu-de
  name = "eu-de-tenant"
  region = "eu-de"
  targets {
    log_sink_crn = "crn:v1:bluemix:public:logs:eu-de:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-ad52-af34-af34-34278391cbc8::"
    name = "my-logs-target"
    parameters {
      host = "www.example-1.com"
      port = 443
    }
  }
}

// Create logs_router_tenants data source
data "ibm_logs_router_tenants" "logs_router_tenants_instance" {
  name = ibm_logs_router_tenant.logs_router_tenant_instance.name
  region = ibm_logs_router_tenant.logs_router_tenant_instance.region
}

// Create logs_router_targets data source
data "ibm_logs_router_targets" "logs_router_targets_instance" {
  tenant_id = ibm_logs_router_tenant.logs_router_tenant_instance.id
  region = ibm_logs_router_tenant.logs_router_tenant_instance.region
}