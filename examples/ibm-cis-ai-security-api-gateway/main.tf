##############################################################################
# IBM CIS – AI Security for Apps & API Gateway  –  End-to-End Validation
#
# Prerequisites (set as env vars or in terraform.tfvars):
#   TF_VAR_ibmcloud_api_key   = "<your IBM Cloud API key>"
#   TF_VAR_cis_crn            = "<CRN of your CIS instance>"
#   TF_VAR_domain_id          = "<zone ID of your CIS domain>"
#   TF_VAR_domain_name        = "<domain name, e.g. example.com>"
##############################################################################

terraform {
  required_providers {
    ibm = {
      source = "IBM-Cloud/ibm"
    }
  }
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

variable "ibmcloud_api_key" {}
variable "cis_crn"          {}
variable "domain_id"        {}
variable "domain_name"      {}

##############################################################################
# 1. AI Security zone-level toggle
##############################################################################

resource "ibm_cis_ai_security_settings" "ai_sec" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id
  enabled   = true
}

data "ibm_cis_ai_security_settings" "ai_sec_read" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id

  depends_on = [ibm_cis_ai_security_settings.ai_sec]
}

##############################################################################
# 2a. Single API Gateway operation (resource)
##############################################################################

resource "ibm_cis_api_gateway_operation" "single_op" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id
  method    = "POST"
  host      = var.domain_name
  endpoint  = "/api/v1/chat/completions"
}

##############################################################################
# 2b. Bulk API Gateway operations (resource)  ← new resource
##############################################################################

resource "ibm_cis_api_gateway_operations" "bulk_ops" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id

  operations {
    method   = "GET"
    host     = var.domain_name
    endpoint = "/api/v1/models"
  }

  operations {
    method   = "DELETE"
    host     = var.domain_name
    endpoint = "/api/v1/sessions"
  }

  operations {
    method   = "POST"
    host     = var.domain_name
    endpoint = "/api/v1/embeddings"
  }
}

##############################################################################
# 3. Operation labels – apply the cf-llm managed label to enable AI scanning
#    on both the single op and the bulk ops
##############################################################################

resource "ibm_cis_api_gateway_operation_labels" "llm_labels" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id

  operation_ids  = concat(
    [ibm_cis_api_gateway_operation.single_op.operation_id],
    [for op in ibm_cis_api_gateway_operations.bulk_ops.operations : op.operation_id]
  )
  managed_labels = ["cf-llm"]
}

##############################################################################
# 4. Discovery operations – mark any already-discovered ops as saved
#    (empty map is valid; populate once you know discovered operation UUIDs)
##############################################################################

resource "ibm_cis_api_gateway_discovery_operations" "promote" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id

  # Map of discovered-operation-UUID → desired state
  # Populate with actual UUIDs from the discovery data source below.
  # Left empty here so the plan/apply does not fail when discovery is empty.
  operation_states = {}
}

##############################################################################
# 5. Data sources – read-only observation
##############################################################################

# List all saved operations (Endpoint Management)
data "ibm_cis_api_gateway_operations" "all_ops" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id

  depends_on = [
    ibm_cis_api_gateway_operation.single_op,
    ibm_cis_api_gateway_operations.bulk_ops,
  ]
}

# Read a single operation by UUID
data "ibm_cis_api_gateway_operation" "single_op_read" {
  cis_id       = var.cis_crn
  domain_id    = var.domain_id
  operation_id = ibm_cis_api_gateway_operation.single_op.operation_id
}

# Discovery – operations seen in traffic but not yet in Endpoint Management
data "ibm_cis_api_gateway_discovery_operations" "discovered" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id
  state     = "review"
}

# Discovery – full OpenAPI schema document
data "ibm_cis_api_gateway_discovery" "openapi_doc" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id
}

# Schemas for saved operations
data "ibm_cis_api_gateway_schemas" "schemas" {
  cis_id    = var.cis_crn
  domain_id = var.domain_id

  depends_on = [ibm_cis_api_gateway_operations.bulk_ops]
}

##############################################################################
# Outputs – useful for spot-checking results
##############################################################################

output "ai_security_enabled" {
  description = "Current AI Security enabled state"
  value       = data.ibm_cis_ai_security_settings.ai_sec_read.enabled
}

output "single_op_id" {
  description = "operation_id of the single-op resource"
  value       = ibm_cis_api_gateway_operation.single_op.operation_id
}

output "bulk_op_ids" {
  description = "operation_ids returned by the bulk resource"
  value       = [for op in ibm_cis_api_gateway_operations.bulk_ops.operations : op.operation_id]
}

output "total_saved_ops" {
  description = "Total saved operations in Endpoint Management"
  value       = data.ibm_cis_api_gateway_operations.all_ops.total_count
}

output "discovered_ops_count" {
  description = "Number of operations currently in 'review' state"
  value       = length(data.ibm_cis_api_gateway_discovery_operations.discovered.operations)
}

output "label_result" {
  description = "Per-operation label assignment results"
  value       = ibm_cis_api_gateway_operation_labels.llm_labels.result
}
