#virtual ip (DNS name) of dynamically allocated load balancer

output "web_dns_name" {
  value = "http://${ibm_lbaas.lbaas1.vip}"
}
