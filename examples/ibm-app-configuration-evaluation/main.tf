provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.region
}

// Single feature flag
data "ibm_app_config_evaluate_feature_flag" "evaluate_feature_flag" {
  guid              = var.app_config_guid
  environment_id    = var.app_config_environment_id
  collection_id     = var.app_config_collection_id
  feature_id        = var.app_config_feature_id
  entity_id         = var.app_config_entity_id
  entity_attributes = var.app_config_entity_attributes
}

// Multiple feature flags
data "ibm_app_config_evaluate_feature_flag" "evaluate_feature_flags" {
  for_each          = toset(var.app_config_feature_flag_ids)
  guid              = var.app_config_guid
  environment_id    = var.app_config_environment_id
  collection_id     = var.app_config_collection_id
  feature_id        = each.value
  entity_id         = var.app_config_entity_id
  entity_attributes = var.app_config_entity_attributes
}

// Single property
data "ibm_app_config_evaluate_property" "evaluate_property" {
  guid              = var.app_config_guid
  environment_id    = var.app_config_environment_id
  collection_id     = var.app_config_collection_id
  property_id       = var.app_config_property_id
  entity_id         = var.app_config_entity_id
  entity_attributes = var.app_config_entity_attributes
}

// Multiple properties
data "ibm_app_config_evaluate_property" "evaluate_properties" {
  for_each          = toset(var.app_config_property_ids)
  guid              = var.app_config_guid
  environment_id    = var.app_config_environment_id
  collection_id     = var.app_config_collection_id
  property_id       = each.value
  entity_id         = var.app_config_entity_id
  entity_attributes = var.app_config_entity_attributes
}
