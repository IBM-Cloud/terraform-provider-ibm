// This allows iam_api_key data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_api_key" {
  value       = ibm_iam_api_key.iam_api_key_instance
  description = "iam_api_key resource instance"
  sensitive = true
}

output "ibm_iam_api_key_data" {
  value = {
    apikey_id   = data.ibm_iam_api_key.iam_api_key_data.apikey_id
    name        = data.ibm_iam_api_key.iam_api_key_data.name
    description = data.ibm_iam_api_key.iam_api_key_data.description
    account_id  = data.ibm_iam_api_key.iam_api_key_data.account_id
    crn         = data.ibm_iam_api_key.iam_api_key_data.crn
    iam_id      = data.ibm_iam_api_key.iam_api_key_data.iam_id
    expires_at  = data.ibm_iam_api_key.iam_api_key_data.expires_at
  }
  description = "iam_api_key data"
}
