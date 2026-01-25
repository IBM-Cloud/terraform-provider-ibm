provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision logs_router_target resource instance
resource "ibm_logs_router_target" "logs_router_target_instance" {
  name = var.logs_router_target_name
  destination_crn = var.logs_router_target_destination_crn
  region = var.logs_router_target_region
  managed_by = var.logs_router_target_managed_by
}

// Provision logs_router_route resource instance
resource "ibm_logs_router_route" "logs_router_route_instance" {
  name = var.logs_router_route_name
  rules {
    action = "send"
    targets {
      id = ibm_logs_router_target.logs_router_target_instance.id
      crn = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
      name = "a-lr-target-us-south"
      target_type = "cloud_logs"
    }
    inclusion_filters {
      operand = "location"
      operator = "is"
      values = [ "us-south" ]
    }
  }
  managed_by = var.logs_router_route_managed_by
}

// Provision logs_router_settings resource instance
// Default target cannot be created with managed_by="enterprise"
resource "ibm_logs_router_settings" "logs_router_settings_instance" {
  default_targets {
    id = ibm_logs_router_target.logs_router_target_instance.id
    crn = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
    name = "a-lr-target-us-south"
    target_type = "cloud_logs"
  }
  permitted_target_regions = var.logs_router_settings_permitted_target_regions
  primary_metadata_region = var.logs_router_settings_primary_metadata_region
  backup_metadata_region = var.logs_router_settings_backup_metadata_region
  private_api_endpoint_only = var.logs_router_settings_private_api_endpoint_only
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_router_targets data source
data "ibm_logs_router_targets" "logs_router_targets_instance" {
  name = var.logs_router_targets_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_router_routes data source
data "ibm_logs_router_routes" "logs_router_routes_instance" {
  name = var.logs_router_routes_name
}
*/
