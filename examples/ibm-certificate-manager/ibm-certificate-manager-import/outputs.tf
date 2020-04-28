output "cert_import_id" {
  value = ibm_certificate_manager_import.cert.id
}
output "cert_import_content" {
  value = ibm_certificate_manager_import.cert.data.content
}
