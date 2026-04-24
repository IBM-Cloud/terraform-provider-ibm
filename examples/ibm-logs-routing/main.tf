# default public endpoints
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "us-east"
}

# alias provider with different region
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "eu-de"
  alias = "provider-eu-de"
}

# Private CSE endpoint
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "jp-osa"
  visibility = "private"
  alias = "provider-private-cse"
}

# Private with VPE endpoint
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "jp-tok"
  visibility = "private"
  private_endpoint_type = "vpe"
  alias = "provider-private-vpe"
}

// public endpoints
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

// public endpoints in eu-de
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

//Private CSE
resource "ibm_logs_router_tenant" "logs_router_tenant_private_cse" {
  provider = ibm.provider-private-cse
  name = "jp-osa-tenant-cse"
  region = "jp-osa"
  targets {
    log_sink_crn = "crn:v1:bluemix:public:logs:jp-osa:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-cse-target"
    parameters {
      host = "www.example-cse.com"
      port = 443
    }
  }
}

// Private VPE endpoints
resource "ibm_logs_router_tenant" "logs_router_tenant_private_vpe" {
  provider = ibm.provider-private-vpe
  name = "jp-tok-tenant-vpe"
  region = "jp-tok"
  targets {
    log_sink_crn = "crn:v1:bluemix:public:logs:jp-tok:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-vpe-target"
    parameters {
      host = "www.example-vpe.com"
      port = 443
    }
  }
}

data "ibm_logs_router_tenants" "logs_router_tenants_instance" {
  name = ibm_logs_router_tenant.logs_router_tenant_instance.name
  region = ibm_logs_router_tenant.logs_router_tenant_instance.region
}

data "ibm_logs_router_targets" "logs_router_targets_instance" {
  tenant_id = ibm_logs_router_tenant.logs_router_tenant_instance.id
  region = ibm_logs_router_tenant.logs_router_tenant_instance.region
}