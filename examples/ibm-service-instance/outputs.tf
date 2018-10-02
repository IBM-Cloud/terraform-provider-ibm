# service instance guid
output "guid" {
  value = "${ibm_service_instance.service-instance.id}"
}
