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
