# Reference DNS registration
/*data "ibm_dns_domain_registration" "web_domain" {
  name = "dnstestdomain.com"
}

# Set DNS name servers for CIS
resource "ibm_dns_domain_registration_nameservers" "web_domain" {
  name_servers        = ibm_cis_domain.web_domain.name_servers
  dns_registration_id = data.ibm_dns_domain_registration.web_domain.id
}
*/

# IBM Cloud Resource Group the CIS instance will be created under
data "ibm_resource_group" "web_group" {
  name = var.resource_group
}
#IBM CLOUD CIS instance resource
resource "ibm_cis" "web_domain" {
  name              = "web_domain"
  resource_group_id = data.ibm_resource_group.web_group.id
  plan              = "standard"
  location          = "global"
}

#Domain settings for IBM CIS instance
resource "ibm_cis_domain_settings" "web_domain" {
  cis_id          = ibm_cis.web_domain.id
  domain_id       = ibm_cis_domain.web_domain.id
  waf             = "on"
  ssl             = "full"
  min_tls_version = "1.2"
}

#Adding valid Domain for IBM CIS instance
resource "ibm_cis_domain" "web_domain" {
  cis_id = ibm_cis.web_domain.id
  domain = var.domain
}

# CIS GLB Monitor|HealthCheck
resource "ibm_cis_healthcheck" "root" {
  cis_id         = ibm_cis.web_domain.id
  description    = "Websiteroot"
  expected_body  = ""
  expected_codes = "200"
  path           = "/"
}

# CIS GLB Origin Pool
resource "ibm_cis_origin_pool" "lon" {
  cis_id        = ibm_cis.web_domain.id
  name          = var.datacenter1
  check_regions = ["WEU"]

  monitor = ibm_cis_healthcheck.root.id

  origins {
    name    = var.datacenter1
    address = "192.0.2.1"
    enabled = true
  }

  description = "LON pool"
  enabled     = true
}

resource "ibm_cis_origin_pool" "ams" {
  cis_id        = ibm_cis.web_domain.id
  name          = var.datacenter2
  check_regions = ["WEU"]

  monitor = ibm_cis_healthcheck.root.id

  origins {
    name    = var.datacenter2
    address = "192.0.2.2"
    enabled = true
  }

  description = "AMS pool"
  enabled     = true
}

# GLB name - name advertised by DNS for the website: prefix + domain
resource "ibm_cis_global_load_balancer" "web_domain" {
  cis_id           = ibm_cis.web_domain.id
  domain_id        = ibm_cis_domain.web_domain.id
  name             = "${var.dns_name}${var.domain}"
  fallback_pool_id = ibm_cis_origin_pool.lon.id
  default_pool_ids = [ibm_cis_origin_pool.lon.id, ibm_cis_origin_pool.ams.id]
  description      = "Load balancer"
  proxied          = true
  session_affinity = "cookie"
}

# CIS DNS Record
resource "ibm_cis_dns_record" "example" {
  cis_id    = ibm_cis.web_domain.id
  domain_id = ibm_cis_domain.web_domain.id
  name      = var.record_name
  type      = var.record_type
  content   = var.record_content
  proxied   = true

}

# CIS Firewall - Present resource supports only lockdown
resource "ibm_cis_firewall" "lockdown" {
  cis_id        = ibm_cis.web_domain.id
  domain_id     = ibm_cis_domain.web_domain.id
  firewall_type = var.firewall_type

  lockdown {
    paused = "true"
    urls   = [var.lockdown_url]

    configurations {
      target = var.lockdown_target
      value  = var.lockdown_value
    }
  }
}

#CIS Rate Limit
resource "ibm_cis_rate_limit" "ratelimit" {
  cis_id    = data.ibm_cis.web_domain.id
  domain_id = data.ibm_cis_domain.web_domain.id
  threshold = var.threshold
  period    = var.period
  match {
    request {
      url     = var.match_request_url
      schemes = var.match_request_schemes
      methods = var.match_request_methods
    }
    response {
      status         = var.match_response_status
      origin_traffic = var.match_response_traffic
      header {
        name  = var.header1_name
        op    = var.header1_op
        value = var.hearder1_value
      }
    }
  }
  action {
    mode    = var.action_mode
    timeout = var.action_timeout
    response {
      content_type = var.action_response_content_type
      body         = var.action_response_body
    }
  }
  correlate {
    by = var.correlate_by
  }
  disabled    = var.disabled
  description = var.description
  bypass {
    name  = var.bypass1_name
    value = var.bypass1_value
  }
}

# CIS Edge Functions action
resource "ibm_cis_edge_functions_action" "test_action" {
  cis_id      = data.ibm_cis.cis.id
  domain_id   = data.ibm_cis_domain.cis_domain.domain_id
  script_name = "sample-script"
  script      = file("./script.js")
}

# CIS Edge Functions trigger
resource "ibm_cis_edge_functions_trigger" "test_trigger" {
  cis_id    = ibm_cis_edge_functions_action.test_action.cis_id
  domain_id = ibm_cis_edge_functions_action.test_action.domain_id
  script    = ibm_cis_edge_functions_action.test_action.script_name
  pattern   = "example.domain.com/*"
}

# CIS Edge Functions action data source
data "ibm_cis_edge_functions_actions" "test_actions" {
  cis_id    = ibm_cis_edge_functions_trigger.test_trigger.cis_id
  domain_id = ibm_cis_edge_functions_trigger.test_trigger.domain_id
}

# CIS Edge Functions trigger data source
data "ibm_cis_edge_functions_triggers" "test_triggers" {
  cis_id    = ibm_cis_edge_functions_trigger.test_trigger.cis_id
  domain_id = ibm_cis_edge_functions_trigger.test_trigger.domain_id
}
