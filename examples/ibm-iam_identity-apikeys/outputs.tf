// This allows iam_api_key data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_api_key" {
  value       = ibm_iam_api_key.iam_api_key_instance
  description = "iam_api_key resource instance"
}
