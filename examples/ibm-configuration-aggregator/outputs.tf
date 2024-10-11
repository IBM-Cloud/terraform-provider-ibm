// This output allows config_aggregator_settings data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "config_aggregator_settings" {
  value = {
    additional_scope            = []
    regions                     = ["all"]
    resource_collection_enabled  = ibm_config_aggregator_settings.config_aggregator_settings_instance.resource_collection_enabled
    trusted_profile_id           = ibm_config_aggregator_settings.config_aggregator_settings_instance.trusted_profile_id
  }
}

output "aggregator_settings" {
  value = {
    additional_scope            = []
    regions                     = ["all"]
    resource_collection_enabled  = ibm_config_aggregator_settings.config_aggregator_settings_instance.resource_collection_enabled
    trusted_profile_id           = ibm_config_aggregator_settings.config_aggregator_settings_instance.trusted_profile_id
  }
}

output "ibm_config_aggregator_configurations" {
  value = data.ibm_config_aggregator_configurations.config_aggregator_configurations_instance
}

output "config_aggregator_resource_collection_status"{
    value={
      status=data.ibm_config_aggregator_resource_collection_status.config_aggregator_resource_collection_status_instance.status
    }
}