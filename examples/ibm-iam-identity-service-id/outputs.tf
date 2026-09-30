// This allows iam_service_id data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed

// for ServiceID resource
output "ibm_iam_service_id" {
  value       = ibm_iam_service_id.iam_service_id_instance
  description = "iam_service_id resource instance"
}

// for ServiceID data source
output "ibm_iam_service_ids" {
  value       = data.ibm_iam_service_id.iam_service_id_data
  description = "iam_service_id data"
}
