// This output allows config_aggregator_settings data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
locals {
  entries = [
    for config in data.ibm_config_aggregator_configurations.example.configs : {
      about = {
        account_id               = jsondecode(config.about).account_id
        config_type              = jsondecode(config.about).config_type
        last_config_refresh_time = jsondecode(config.about).last_config_refresh_time
        location                 = jsondecode(config.about).location
        resource_crn             = jsondecode(config.about).resource_crn
        resource_group_id        = jsondecode(config.about).resource_group_id
        resource_name            = jsondecode(config.about).resource_name
        service_name             = jsondecode(config.about).service_name
        tags                     = {}
      }
      config = jsondecode(config.config)
    }
  ]
}
output "ibm_config_aggregator_configurations" {
  value = {
    configs=local.entries
  }
}
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

output "config_aggregator_resource_collection_status"{
    value={
      status=data.ibm_config_aggregator_resource_collection_status.config_aggregator_resource_collection_status_instance.status
    }
}