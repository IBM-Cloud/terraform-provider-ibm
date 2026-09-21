// This allows iam_trusted_profile_link data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed

// for link CRUD operations 
output "ibm_iam_trusted_profile_link" {
  value       = ibm_iam_trusted_profile_link.iam_trusted_profile_link_instance
  description = "iam_trusted_profile_link resource instance"
}

output "ibm_iam_trusted_profile_link_data" {
  value       = data.ibm_iam_trusted_profile_link.iam_trusted_profile_link_instance_data
  description = "iam_trusted_profile_link data"
}

// for link list operation
output "ibm_iam_trusted_profile_links_data" {
  value       = data.ibm_iam_trusted_profile_links.iam_trusted_profile_links_data
  description = "iam_trusted_profile_links data"
}
