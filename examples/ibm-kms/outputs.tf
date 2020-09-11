
output "key_value" {
  value = data.ibm_kms_keys.test.keys.*
}