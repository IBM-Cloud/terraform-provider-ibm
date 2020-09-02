output "cert_order" {
  value = ibm_certificate_manager_order.cert
}
output "cis_domain" {
  value = data.ibm_cis_domain.domain
}

