provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "us-south"
}

// Provision logs_router_tenant resource instance
resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
  name = var.logs_router_tenant_name
  region = "us-east"
  targets {
    log_sink_crn = "crn:v1:bluemix:public:logdna:us-east:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-logdna-target"
    parameters {
      host = "www.example-1.com"
      port = 443
      access_credential = "new-cred"
    }
  }
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
  name = "eu-de-tenant"
  region = "eu-de"
  targets {
    log_sink_crn = "crn:v1:bluemix:public:logdna:eu-de:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-logdna-target"
    parameters {
      host = "www.example-1.com"
      port = 443
      access_credential = "new-cred"
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