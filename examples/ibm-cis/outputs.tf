output "web_dns_name" {
  value = "http://${var.dns_name}${var.domain}"
}
