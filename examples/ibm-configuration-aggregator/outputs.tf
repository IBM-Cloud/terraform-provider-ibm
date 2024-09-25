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