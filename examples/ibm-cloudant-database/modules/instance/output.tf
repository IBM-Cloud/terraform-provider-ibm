#####################################################
# Cloudant
# Copyright 2021 IBM
#####################################################

output "cloudant_id" {
  description = "The ID of the cloudant instance"
  value       = var.provision ? concat(ibm_cloudant.cloudant_instance.*.id, [""])[0] : concat(data.ibm_cloudant.cloudant.*.id, [""])[0]
}

output "cloudant_instance_name" {
  description = "The name of the cloudant instance"
  value       = var.provision ? concat(ibm_cloudant.cloudant_instance.*.name, [""])[0] : concat(data.ibm_cloudant.cloudant.*.name, [""])[0]
}

output "cloudant_instance_crn" {
  description = "The CRN of the cloudant instance"
  value       = var.provision ? concat(ibm_cloudant.cloudant_instance.*.crn, [""])[0] : concat(data.ibm_cloudant.cloudant.*.crn, [""])[0]
}

output "cloudant_key_id" {
  description = "The ID of the cloudant key"
  value       = var.provision_resource_key ? concat(ibm_resource_key.resource_key.*.id, [""])[0] : concat(data.ibm_resource_key.cloudant_resource_key.*.id, [""])[0]
}

output "cloudant_key_host" {
  description = "API key"
  value       = var.provision_resource_key ? concat(ibm_resource_key.resource_key.*.credentials.host, [""])[0] : concat(data.ibm_resource_key.cloudant_resource_key.*.credentials.host, [""])[0]
}

output "cloudant_key_username" {
  description = "username"
  value       = var.provision_resource_key ? concat(ibm_resource_key.resource_key.*.credentials.username, [""])[0] : concat(data.ibm_resource_key.cloudant_resource_key.*.credentials.username, [""])[0]
}

output "cloudant_key_password" {
  description = "password"
  value       = var.provision_resource_key ? concat(ibm_resource_key.resource_key.*.credentials.password, [""])[0] : concat(data.ibm_resource_key.cloudant_resource_key.*.credentials.password, [""])[0]
}

output "cloudant_key_apikey" {
  description = "password"
  value       = var.provision_resource_key ? concat(ibm_resource_key.resource_key.*.credentials.apikey, [""])[0] : concat(data.ibm_resource_key.cloudant_resource_key.*.credentials.apikey, [""])[0]
}


#####################################################
# Service Policy
#####################################################

output "cloudant_service_uuid" {
  description = "The UUID of the service ID"
  value       = ibm_iam_service_policy.policy.iam_service_id
}

output "cloudant_service_iam_id" {
  description = "IAM ID of the service ID"
  value       = ibm_iam_service_policy.policy.iam_id
}

output "cloudant_service_policy_id" {
  description = "The ID of the service policy"
  value       = ibm_iam_service_policy.policy.id
}