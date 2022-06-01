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
  brotli					= "on"
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

# CIS Firewall
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

# CIS Firewall data source
data "ibm_cis_firewall" "ua_rules" {
  cis_id        = ibm_cis.web_domain.id
  domain_id     = ibm_cis_domain.web_domain.id
  firewall_type = "ua_rules"
}
# CIS Filter
resource "ibm_cis_filter" "test" {
  cis_id          = ibm_cis.web_domain.id
  domain_id       = ibm_cis_domain.web_domain.id
  expression = "(ip.src eq 19.25.53.139 and http.request.uri.path eq \"^.*/wp-login[0-5].php$\") or (http.request.uri.path eq \"^.*/xmlrpc[[:xdigit:]].php$\")"
  paused =  false
  description = "Filter-creation"
}
# CIS Filter data source
data "ibm_cis_filters" "test" {
  cis_id    = data.ibm_cis_filters.test.cis_id
  domain_id = data.ibm_cis_filters.test.domain_id
}

# CIS Firewall Rules Resource
resource "ibm_cis_firewall_rule" "firewall_rules_instance" {
  cis_id = ibm_cis.web_domain.id
  domain_id = ibm_cis_domain.web_domain.id
  filter_id = ibm_cis_filter.test.filter_id
  action = "allow"
  priority = 5
  description = "Firewallrules-creation"

}
# CIS Firewall Rules data source
data "ibm_cis_firewall_rules" "test" {
  cis_id    = data.ibm_cis_firewall_rules.test.cis_id
  domain_id = data.ibm_cis_firewall_rules.test.domain_id
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
  action_name = "sample-script"
  script      = file("./script.js")
}

# CIS Edge Functions trigger
resource "ibm_cis_edge_functions_trigger" "test_trigger" {
  cis_id      = ibm_cis_edge_functions_action.test_action.cis_id
  domain_id   = ibm_cis_edge_functions_action.test_action.domain_id
  action_name = ibm_cis_edge_functions_action.test_action.action_name
  pattern_url = "example.com/*"
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

# CIS TLS Settings
resource "ibm_cis_tls_settings" "tls_settings" {
  cis_id          = data.ibm_cis.cis.id
  domain_id       = data.ibm_cis_domain.cis_domain.domain_id
  tls_1_3         = "off"
  min_tls_version = "1.2"
  universal_ssl   = true
}

# CIS Routing
resource "ibm_cis_routing" "routing" {
  cis_id        = data.ibm_cis.cis.id
  domain_id     = data.ibm_cis_domain.cis_domain.domain_id
  smart_routing = "on"
}

# CIS Cache Settings
resource "ibm_cis_cache_settings" "test" {
  cis_id             = var.cis_crn
  domain_id          = var.zone_id
  caching_level      = "aggressive"
  browser_expiration = 14400
  development_mode   = "off"
  query_string_sort  = "off"
  purge_all          = true
  serve_stale_content = "on"
}

# CIS Cache Settings data source
data "ibm_cis_cache_settings" "test" {
  cis_id    = data.ibm_cis_cache_settings.test.cis_id
  domain_id = data.ibm_cis_cache_settings.test.domain_id
}


# CIS Custom Page service
resource "ibm_cis_custom_page" "custom_page" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  page_id   = "basic_challenge"
  url       = "https://test.com/index.html"
}

# CIS Custom Page service data source
data "ibm_cis_custom_pages" "custom_pages" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
}

# CIS Page Rule service
resource "ibm_cis_page_rule" "page_rule" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  targets {
    target = "url"
    constraint {
      operator = "matches"
      value    = "example.com"
    }
  }
  actions {
    id    = "email_obfuscation"
    value = "on"
  }
}

# CIS Page Rule data source
data "ibm_cis_page_rules" "rules" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}

# CIS WAF Packages
resource "ibm_cis_waf_package" "test" {
  cis_id      = data.ibm_cis.cis.id
  domain_id   = data.ibm_cis_domain.cis_domain.domain_id
  package_id  = "c504870194831cd12c3fc0284f294abb"
  sensitivity = "low"
  action_mode = "block"
}

# CIS WAF Packages data source
data "ibm_cis_waf_packages" "packages" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
}

# CIS WAF Rule Group service
resource "ibm_cis_waf_group" "test" {
  cis_id     = data.ibm_cis.cis.id
  domain_id  = data.ibm_cis_domain.cis_domain.domain_id
  package_id = "c504870194831cd12c3fc0284f294abb"
  group_id   = "3d8fb0c18b5a6ba7682c80e94c7937b2"
  mode       = "on"
}

# CIS WAF Rule Groups data source
data "ibm_cis_waf_groups" "waf_groups" {
  cis_id     = data.ibm_cis.cis.id
  domain_id  = data.ibm_cis_domain.cis_domain.id
  package_id = "c504870194831cd12c3fc0284f294abb"
}

# CIS Rnage application service
resource "ibm_cis_range_app" "app" {
  cis_id         = data.ibm_cis.cis.id
  domain_id      = data.ibm_cis_domain.cis_domain.id
  protocol       = "tcp/22"
  dns_type       = "CNAME"
  dns            = "ssh.example.com"
  origin_direct  = ["tcp://12.1.1.1:22"]
  ip_firewall    = true
  proxy_protocol = "v1"
  traffic_type   = "direct"
  tls            = "off"
}

# CIS Range application data source
data "ibm_cis_range_apps" "test" {
  cis_id    = ibm_cis_range_app.app.cis_id
  domain_id = ibm_cis_range_app.app.domain_id
}

# CIS WAF Rule service
resource "ibm_cis_waf_rule" "test" {
  cis_id     = data.ibm_cis.cis.id
  domain_id  = data.ibm_cis_domain.cis_domain.id
  package_id = "c504870194831cd12c3fc0284f294abb"
  rule_id    = "100000356"
  mode       = "on"
}

# CIS WAF Rule data source
data "ibm_cis_waf_rules" "rules" {
  cis_id     = data.ibm_cis.cis.id
  domain_id  = data.ibm_cis_domain.cis_domain.id
  package_id = "1e334934fd7ae32ad705667f8c1057aa"
}

# CIS Certificate order service
resource "ibm_cis_certificate_order" "test" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  hosts     = ["example.com"]
}

# CIS Certificates data source
data "ibm_cis_certificates" "test" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
}

# CIS Certificate Upload
resource "ibm_cis_certificate_upload" "test" {
  cis_id        = data.ibm_cis.cis.id
  domain_id     = data.ibm_cis_domain.cis_domain.id
  certificate   = "xxxxx"
  private_key   = "xxxxx"
  bundle_method = "ubiquitous"
  priority      = 20
}

# CIS Certificate Upload data source
data "ibm_cis_custom_certificates" "test" {
  cis_id    = ibm_cis_certificate_upload.test.cis_id
  domain_id = ibm_cis_certificate_upload.test.domain_id
}

# CIS DNS Records import service
resource "ibm_cis_dns_records_import" "test" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  file      = "dns_records.txt"
}

# CIS Webhooks
resource "ibm_cis_webhook" "test" {
    cis_id = data.ibm_cis.cis.id
    name    = "test-Webhooks"
    url     = "https://hooks.slack.com/services/Ds3fdBFbV/1234568"
    secret = "ZaHkAf0iNXNWn8ySUJjTJHkzlanchfnR4TISjOPC_I1U"
}

# CIS Webhooks data source
data "ibm_cis_webhooks" "test1" {
  cis_id = data.ibm_cis.cis.id
}

# CIS Alert Policy
resource "ibm_cis_alert" "test" {
  depends_on  = [ibm_cis_webhook.test]
  cis_id      = data.ibm_cis.cis.id
  name        = "test-alert-police"
  description = "alert policy description"
  enabled     = true
  alert_type = "clickhouse_alert_fw_anomaly"
  mechanisms {
    email    = ["mynotifications@email.com"]
    webhooks = [ibm_cis_webhook.test.webhook_id]
  }
 filters =<<FILTERS
  		{}
  		FILTERS
 conditions =<<CONDITIONS
  		{}
  		CONDITIONS

} 
# CIS Alert Policy Data source
data "ibm_cis_alerts" "test1" {
  cis_id = data.ibm_cis.cis.id
}
# CIS Logpush Job
resource "ibm_cis_logpush_job" "test" {
    cis_id          = data.ibm_cis.cis.id
    domain_id       = data.ibm_cis_domain.cis_domain.domain_id
    name            = "MylogpushjobUpdate"
    enabled         = true
    logpull_options = "timestamps=rfc3339&timestamps=rfc3339"
    dataset         = "http_requests"
    frequency       = "high"
    logdna =<<LOG
        {
                        "hostname": "cistest-load.com",
            "ingress_key": "e2f7xxxxx73a251caxxxxxxxxxxxx",
            "region": "in-che"
        }
        LOG
}
# CIS Logpush Job Data source
data "ibm_cis_logpush_jobs" "test" {
    cis_id          = data.ibm_cis.cis.id
    domain_id       = data.ibm_cis_domain.cis_domain.domain_id
}