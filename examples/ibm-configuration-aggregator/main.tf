provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision config_aggregator_settings resource instance
resource "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
  resource_collection_enabled = var.config_aggregator_settings_resource_collection_enabled
  trusted_profile_id          = var.config_aggregator_settings_trusted_profile_id
  regions                     = var.config_aggregator_settings_regions
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create config_aggregator_configurations data source


*/

data "ibm_config_aggregator_configurations" "config_aggregator_configurations_instance" {
  config_type                = var.config_type
  location                   = var.location
  resource_crn               = var.resource_crn
  resource_group_id          = var.resource_group_id
  service_name               = var.service_name
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create config_aggregator_settings data source
data "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
    resource_collection_enabled = var.config_aggregator_settings_resource_collection_enabled
    trusted_profile_id          = var.config_aggregator_settings_trusted_profile_id
    regions                     = var.config_aggregator_settings_regions
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create config_aggregator_resource_collection_status data source
data "ibm_config_aggregator_resource_collection_status" "config_aggregator_resource_collection_status_instance" {
    status = "complete"
}
*/