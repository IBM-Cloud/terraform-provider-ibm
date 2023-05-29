// This allows cm_catalog data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cm_catalog" {
  value       = ibm_cm_catalog.cm_catalog_instance
  description = "cm_catalog resource instance"
}
// This allows cm_offering data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cm_offering" {
  value       = ibm_cm_offering.cm_offering_instance
  description = "cm_offering resource instance"
}
// This allows cm_version data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cm_version" {
  value       = ibm_cm_version.cm_version_instance
  description = "cm_version resource instance"
}
// This allows cm_offering_instance data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cm_offering_instance" {
  value       = ibm_cm_offering_instance.cm_offering_instance_instance
  description = "cm_offering_instance resource instance"
}
