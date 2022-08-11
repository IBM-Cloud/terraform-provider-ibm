output "web_dns_name" {
  value = "http://${var.dns_name}${var.domain}"
}
output "instance_id" {
  value = ibm_cis.web_domain.id
}
output "domain_id" {
  value = ibm_cis_domain.web_domain.id
}
output "monitor" {
  value = ibm_cis_healthcheck.root.id
}
output "rate_limit_id" {
  value = ibm_cis_rate_limit.ratelimit.id
}
output "ibm_cis_edge_functions_action_output" {
  value = ibm_cis_edge_functions_action.test_action
}

output "ibm_cis_edge_function_trigger_output" {
  value = ibm_cis_edge_functions_trigger.test_trigger
}

output "domain_setting" {
  value = ibm_cis_domain_settings.test_domain_settings
}

output "ibm_cis_tls_settings_output" {
  value = ibm_cis_tls_settings.tls_settings
}

output "ibm_cis_routing_output" {
  value = ibm_cis_routing.routing
}

output "cache_settings" {
  value = ibm_cis_cache_settings.test
}

output "ibm_cis_custom_page_output" {
  value = ibm_cis_custom_page.custom_page
}

output "ibm_cis_firewall_ouput" {
  value = ibm_cis_firewall.lockdown
}

output "ibm_cis_page_rule_output" {
  value = ibm_cis_page_rule.page_rule
}

output "ibm_cis_waf_package_output" {
  value = ibm_cis_waf_package.test
}

output "ibm_cis_waf_group_output" {
  value = ibm_cis_waf_group.test
}

output "ibm_cis_range_app_output" {
  value = ibm_cis_range_app.app
}

output "ibm_cis_waf_rules_output" {
  value = ibm_cis_waf_rule.test
}

output "ibm_cis_certificate_order_output" {
  value = ibm_cis_certificate_order.test
}

output "ibm_cis_certificate_upload_output" {
  value = ibm_cis_certificate_upload.test
}
