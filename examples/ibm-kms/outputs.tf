
output "key_value" {
  value = data.ibm_kms_key.test.keys.*
}