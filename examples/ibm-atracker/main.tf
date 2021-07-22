provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision atracker_target resource instance
resource "ibm_atracker_target" "atracker_target_instance" {
}

// Provision atracker_route resource instance
resource "ibm_atracker_route" "atracker_route_instance" {
  name = var.atracker_route_name
  receive_global_events = var.atracker_route_receive_global_events
  rules = var.atracker_route_rules
}

// Create atracker_targets data source
data "ibm_atracker_targets" "atracker_targets_instance" {
  name = var.atracker_targets_name
}

// Create atracker_routes data source
data "ibm_atracker_routes" "atracker_routes_instance" {
  name = var.atracker_routes_name
}

// Create atracker_endpoints data source
data "ibm_atracker_endpoints" "atracker_endpoints_instance" {
}
