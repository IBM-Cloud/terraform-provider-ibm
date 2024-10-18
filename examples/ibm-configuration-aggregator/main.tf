provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
  instance_id=var.instance_id
  region =var.region
  resource_collection_enabled = var.config_aggregator_settings_resource_collection_enabled
  trusted_profile_id          = var.config_aggregator_settings_trusted_profile_id
  resource_collection_regions                    = var.config_aggregator_settings_regions
}

data "ibm_config_aggregator_configurations" "example" {
  instance_id=var.instance_id
  region =var.region

}


data "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
    instance_id=var.instance_id
    region =var.region
}

data "ibm_config_aggregator_resource_collection_status" "config_aggregator_resource_collection_status_instance" {
    instance_id=var.instance_id
    region =var.region
}