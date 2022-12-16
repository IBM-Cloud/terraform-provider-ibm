provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision metrics_router_target resource instance
resource "ibm_metrics_router_target" "metrics_router_target_instance" {
  name = var.metrics_router_target_name
  destination_crn = var.metrics_router_target_destination_crn
  region = var.metrics_router_target_region
}

// Provision metrics_router_route resource instance
resource "ibm_metrics_router_route" "metrics_router_route_instance" {
  name = var.metrics_router_route_name
  rules {
    target_ids = [ ibm_metrics_router_target.metrics_router_target_instance.id ]
    inclusion_filters {
      operand = "location"
      operator = "is"
      value = [ "value" ]
    }
  }
}

// Provision metrics_router_settings resource instance
resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
  metadata_region_primary = var.metrics_router_settings_metadata_region_primary
  private_api_endpoint_only = var.metrics_router_settings_private_api_endpoint_only
  default_targets = var.metrics_router_settings_default_targets
  permitted_target_regions = var.metrics_router_settings_permitted_target_regions
}

// Create metrics_router_targets data source
data "ibm_metrics_router_targets" "metrics_router_targets_instance" {
  name = var.metrics_router_targets_name
}

// Create metrics_router_routes data source
data "ibm_metrics_router_routes" "metrics_router_routes_instance" {
  name = var.metrics_router_routes_name
}
