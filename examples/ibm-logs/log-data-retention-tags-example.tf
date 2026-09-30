# Example configuration for PR 188 changes:
# 1. New resource: ibm_logs_log_data_retention_tags
# 2. Updated field: archive_retention_tag in ibm_logs_policy

# Prerequisites:
# - IBM Cloud Logs instance must be provisioned
# - Archive bucket must be configured in the Logs instance
# - Set the following environment variables or update terraform.tfvars:
#   - logs_instance_id: GUID of your IBM Cloud Logs instance
#   - region: Region where your Logs instance is located (e.g., us-south, eu-gb)

# ============================================================================
# Provider Configuration
# ============================================================================
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

# ============================================================================
# 1. Configure Log Data Retention Tags
# ============================================================================
# This resource manages the 3 editable retention tags (IDs 1, 2, 3)
# Note: Tags 0 (default) and 4-9 (reserved) are read-only and managed by the system
# The tags list must contain exactly 3 tag names in order (for tag IDs 1, 2, 3)

resource "ibm_logs_log_data_retention_tags" "example" {
  instance_id = var.logs_instance_id
  region      = var.region

  # List of 3 tag names for retention tags 1, 2, and 3
  # Order matters: [tag_1_name, tag_2_name, tag_3_name]
  tags = [
    "short-term",   # Tag ID 1: Short-term retention (e.g., 7 days)
    "medium-term",  # Tag ID 2: Medium-term retention (e.g., 30 days)
    "long-term"     # Tag ID 3: Long-term retention (e.g., 90 days)
  ]
}

# ============================================================================
# 2. Create a Policy with Archive Retention Tag
# ============================================================================
# This policy uses the new archive_retention_tag field instead of the old
# archive_retention object. The tag name must match one of the tags configured
# in ibm_logs_log_data_retention_tags resource.

resource "ibm_logs_policy" "example_with_retention_tag" {
  instance_id = var.logs_instance_id
  region      = var.region

  name        = "Example Policy with Retention Tag"
  description = "Policy demonstrating the new archive_retention_tag field"
  priority    = "type_high"
  enabled     = true

  # Application rule: match logs from specific applications
  application_rule {
    rule_type_id = "start_with"
    name         = "production"
  }

  # Subsystem rule: match logs from specific subsystems
  subsystem_rule {
    rule_type_id = "is"
    name         = "api-gateway"
  }

  # NEW FIELD: archive_retention_tag
  # References a retention tag name from ibm_logs_log_data_retention_tags
  # This replaces the old archive_retention { id = "..." } syntax
  # Using index 1 to reference the second tag (medium-term)
  archive_retention_tag = ibm_logs_log_data_retention_tags.example.tags[1]

  # Log rules: define which log severities to apply the policy to
  log_rules {
    severities = ["info", "warning", "error", "critical"]
  }

  # Ensure retention tags are configured before creating the policy
  depends_on = [ibm_logs_log_data_retention_tags.example]
}

# ============================================================================
# 3. Data Source: Read Policy with Retention Tag
# ============================================================================
# Read a specific policy to verify the archive_retention_tag field

data "ibm_logs_policy" "example" {
  instance_id    = var.logs_instance_id
  region         = var.region
  logs_policy_id = ibm_logs_policy.example_with_retention_tag.policy_id

  depends_on = [ibm_logs_policy.example_with_retention_tag]
}

# ============================================================================
# 4. Data Source: Read Retention Tags Configuration
# ============================================================================
# Read the retention tags configuration using the new data source

data "ibm_logs_log_data_retention_tags" "example" {
  instance_id = var.logs_instance_id
  region      = var.region

  depends_on = [ibm_logs_log_data_retention_tags.example]
}

# ============================================================================
# 5. Data Source: List All Policies (includes deprecated fields)
# ============================================================================
# Read all policies to verify backward compatibility with deprecated archive_retention field

data "ibm_logs_policies" "example" {
  instance_id = var.logs_instance_id
  region      = var.region

  depends_on = [ibm_logs_policy.example_with_retention_tag]
}

# ============================================================================
# Outputs
# ============================================================================

output "retention_tags_configuration" {
  description = "Configured retention tags"
  value = {
    tag_1 = ibm_logs_log_data_retention_tags.example.tags[0]
    tag_2 = ibm_logs_log_data_retention_tags.example.tags[1]
    tag_3 = ibm_logs_log_data_retention_tags.example.tags[2]
  }
}

output "policy_archive_retention_tag" {
  description = "Archive retention tag used by the policy"
  value       = data.ibm_logs_policy.example.archive_retention_tag
}

output "policy_details" {
  description = "Policy details including the new archive_retention_tag field"
  value = {
    id                    = ibm_logs_policy.example_with_retention_tag.policy_id
    name                  = data.ibm_logs_policy.example.name
    archive_retention_tag = data.ibm_logs_policy.example.archive_retention_tag
    enabled               = data.ibm_logs_policy.example.enabled
    priority              = data.ibm_logs_policy.example.priority
  }
}

output "retention_tags_from_datasource" {
  description = "Retention tags read from data source"
  value = {
    tag_1 = data.ibm_logs_log_data_retention_tags.example.tags[0]
    tag_2 = data.ibm_logs_log_data_retention_tags.example.tags[1]
    tag_3 = data.ibm_logs_log_data_retention_tags.example.tags[2]
  }
}

output "all_policies_count" {
  description = "Number of policies in the instance"
  value       = length(data.ibm_logs_policies.example.policies)
}