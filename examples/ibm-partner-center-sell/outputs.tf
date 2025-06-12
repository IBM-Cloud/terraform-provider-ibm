// This output allows onboarding_resource_broker data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_onboarding_resource_broker" {
  value       = ibm_onboarding_resource_broker.onboarding_resource_broker_instance
  description = "onboarding_resource_broker resource instance"
}
// This output allows onboarding_catalog_deployment data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_onboarding_catalog_deployment" {
  value       = ibm_onboarding_catalog_deployment.onboarding_catalog_deployment_instance
  description = "onboarding_catalog_deployment resource instance"
}
// This output allows onboarding_catalog_plan data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_onboarding_catalog_plan" {
  value       = ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance
  description = "onboarding_catalog_plan resource instance"
}
// This output allows onboarding_catalog_product data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_onboarding_catalog_product" {
  value       = ibm_onboarding_catalog_product.onboarding_catalog_product_instance
  description = "onboarding_catalog_product resource instance"
}
// This output allows onboarding_iam_registration data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_onboarding_iam_registration" {
  value       = ibm_onboarding_iam_registration.onboarding_iam_registration_instance
  description = "onboarding_iam_registration resource instance"
}
// This output allows onboarding_product data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_onboarding_product" {
  value       = ibm_onboarding_product.onboarding_product_instance
  description = "onboarding_product resource instance"
}
// This output allows onboarding_registration data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_onboarding_registration" {
  value       = ibm_onboarding_registration.onboarding_registration_instance
  description = "onboarding_registration resource instance"
}
