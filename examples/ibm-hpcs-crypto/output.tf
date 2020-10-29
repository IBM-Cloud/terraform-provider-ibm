output "InstanceGUID" {
  value = data.ibm_resource_instance.hpcs_instance.guid
}
output "keyID" {
  value = ibm_kms_key.key.id
}