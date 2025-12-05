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
    action = "send"
    targets {
      id = ibm_metrics_router_target.metrics_router_target_instance.id
    }
    inclusion_filters {
      operand = "location"
      operator = "is"
      values = [ "us-south" ]
    }
  }
}

// Provision metrics_router_target resource instance
resource "ibm_metrics_router_target" "metrics_router_target_instance_enterprise" {
  name = var.metrics_router_target_name
  destination_crn = var.metrics_router_target_destination_crn
  region = var.metrics_router_target_region
  managed_by = "enterprise"
}

// Provision metrics_router_route resource instance
resource "ibm_metrics_router_route" "metrics_router_route_instanc_enterprise" {
  name = var.metrics_router_route_name
  rules {
    action = "send"
    targets {
      id = ibm_metrics_router_target.metrics_router_target_instance.id
    }
    inclusion_filters {
      operand = "location"
      operator = "is"
      values = [ "us-south" ]
    }
  }
  managed_by = "enterprise"
}

// Provision metrics_router_settings resource instance
resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
  default_targets {
    id = ibm_metrics_router_target.metrics_router_target_instance.id    
  }
  permitted_target_regions = var.metrics_router_settings_permitted_target_regions
  primary_metadata_region = var.metrics_router_settings_primary_metadata_region
  backup_metadata_region = var.metrics_router_settings_backup_metadata_region
  private_api_endpoint_only = var.metrics_router_settings_private_api_endpoint_only
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create metrics_router_targets data source
data "ibm_metrics_router_targets" "metrics_router_targets_instance" {
  name = var.metrics_router_targets_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create metrics_router_routes data source
data "ibm_metrics_router_routes" "metrics_router_routes_instance" {
  name = var.metrics_router_routes_name
}
*/
