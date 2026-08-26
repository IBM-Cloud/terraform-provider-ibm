// This allows iam_api_key data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_api_key" {
  value       = ibm_iam_api_key.iam_api_key_instance
  description = "iam_api_key resource instance"
  sensitive = true
}

output "ibm_iam_api_key_data" {
  value = {
    account_id  = data.ibm_iam_api_key.iam_api_key_data.account_id
    apikey_id   = data.ibm_iam_api_key.iam_api_key_data.apikey_id
    created_at  = data.ibm_iam_api_key.iam_api_key_data.created_at
    created_by  = data.ibm_iam_api_key.iam_api_key_data.created_by
    crn         = data.ibm_iam_api_key.iam_api_key_data.crn
    description = data.ibm_iam_api_key.iam_api_key_data.description
    entity_tag  = data.ibm_iam_api_key.iam_api_key_data.entity_tag
    expires_at  = data.ibm_iam_api_key.iam_api_key_data.expires_at
    iam_id      = data.ibm_iam_api_key.iam_api_key_data.iam_id
    locked      = data.ibm_iam_api_key.iam_api_key_data.locked
    modified_at = data.ibm_iam_api_key.iam_api_key_data.modified_at
    name        = data.ibm_iam_api_key.iam_api_key_data.name
  }
  description = "iam_api_key data"
}
