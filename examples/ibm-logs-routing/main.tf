provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region = "us-south"
}

// Provision logs-router_tenant resource instance
resource "ibm_logs-router_tenant" "logs-router_tenant_instance" {
  ibm_api_version = var.logs-router_tenant_ibm_api_version
  name = var.logs-router_tenant_name
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