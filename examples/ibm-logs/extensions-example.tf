# Extensions Examples - Testing all extension-related APIs
# This file can be used standalone or alongside other logs resources

# Data Source: List all available extensions
data "ibm_logs_extensions" "all_extensions" {
  instance_id = var.logs_instance_id
  region      = var.region
}

# Data Source: Get details of a specific extension (IBMCloudant)
data "ibm_logs_extension" "cloudant_extension" {
  instance_id        = var.logs_instance_id
  region             = var.region
  logs_extension_id  = "IBMCloudant"
}

# Resource: Create an extension deployment
resource "ibm_logs_extension_deployment" "cloudant_deployment" {
  instance_id  = var.logs_instance_id
  region       = var.region
  logs_extension_id = "IBMCloudant"
  
  # Use the latest version from the extension data source
  version = data.ibm_logs_extension.cloudant_extension.revisions[0].version
  
  # Deploy all items from the extension
  item_ids = [for item in data.ibm_logs_extension.cloudant_extension.revisions[0].items : item.id]
  
  # Optional: Filter by applications
  applications = ["test-app"]
  
  # Optional: Filter by subsystems
  subsystems = ["test-subsystem"]
}

# Data Source: Read the created extension deployment
data "ibm_logs_extension_deployment" "read_deployment" {
  instance_id                   = var.logs_instance_id
  region                        = var.region
  logs_extension_id  = ibm_logs_extension_deployment.cloudant_deployment.extension_deployment_id
  
  depends_on = [ibm_logs_extension_deployment.cloudant_deployment]
}

# Outputs
output "all_extensions" {
  value       = data.ibm_logs_extensions.all_extensions.extensions
  description = "List of all available extensions"
}

output "cloudant_extension_details" {
  value = {
    id          = data.ibm_logs_extension.cloudant_extension.id
    name        = data.ibm_logs_extension.cloudant_extension.name
    revisions   = data.ibm_logs_extension.cloudant_extension.revisions
  }
  description = "Details of IBMCloudant extension"
}

output "deployment_id" {
  value       = ibm_logs_extension_deployment.cloudant_deployment.extension_deployment_id
  description = "ID of the created extension deployment"
}

output "deployment_details" {
  value = {
    id           = ibm_logs_extension_deployment.cloudant_deployment.id
    extension_id = ibm_logs_extension_deployment.cloudant_deployment.logs_extension_id
    version      = ibm_logs_extension_deployment.cloudant_deployment.version
    item_ids     = ibm_logs_extension_deployment.cloudant_deployment.item_ids
    applications = ibm_logs_extension_deployment.cloudant_deployment.applications
    subsystems   = ibm_logs_extension_deployment.cloudant_deployment.subsystems
  }
  description = "Full details of the extension deployment"
}

output "read_deployment_details" {
  value = {
    id           = data.ibm_logs_extension_deployment.read_deployment.id
    version      = data.ibm_logs_extension_deployment.read_deployment.version
    item_ids     = data.ibm_logs_extension_deployment.read_deployment.item_ids
    applications = data.ibm_logs_extension_deployment.read_deployment.applications
    subsystems   = data.ibm_logs_extension_deployment.read_deployment.subsystems
  }
  description = "Details read from the extension deployment data source"
}