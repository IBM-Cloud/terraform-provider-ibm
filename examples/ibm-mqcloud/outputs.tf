// This output allows mqcloud_queue_manager data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_mqcloud_queue_manager" {
  value       = ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance
  description = "mqcloud_queue_manager resource instance"
}
// This output allows mqcloud_application data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_mqcloud_application" {
  value       = ibm_mqcloud_application.mqcloud_application_instance
  description = "mqcloud_application resource instance"
}
// This output allows mqcloud_user data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_mqcloud_user" {
  value       = ibm_mqcloud_user.mqcloud_user_instance
  description = "mqcloud_user resource instance"
}
// This output allows mqcloud_keystore_certificate data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_mqcloud_keystore_certificate" {
  value       = ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate_instance
  description = "mqcloud_keystore_certificate resource instance"
}
// This output allows mqcloud_truststore_certificate data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_mqcloud_truststore_certificate" {
  value       = ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance
  description = "mqcloud_truststore_certificate resource instance"
}
// This output allows mqcloud_virtual_private_endpoint_gateway data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_mqcloud_virtual_private_endpoint_gateway" {
  value       = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance
  description = "mqcloud_virtual_private_endpoint_gateway resource instance"
}
