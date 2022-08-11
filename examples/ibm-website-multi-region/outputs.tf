#virtual ip (DNS name) of dynamically allocated load balancer

output "web_dns_name" {
  value = "http://${var.dns_name}${var.domain}"
}
