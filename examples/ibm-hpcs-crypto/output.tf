output "InstanceGUID" {
  value = ibm_hpcs.hpcs.guid
}
output "keyID" {
  value = ibm_kms_key.key.id
}