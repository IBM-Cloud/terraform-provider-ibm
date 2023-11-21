// This output allows ibm_app_config_evaluate_feature_flag data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed

output "ibm_app_config_feature_flag_evaluated_value" {
  value       = data.ibm_app_config_evaluate_feature_flag.evaluate_feature_flag.result_numeric
  description = "Feature flag evaluated value."
}
output "ibm_app_config_feature_flag_evaluated_values" {
  value       = data.ibm_app_config_evaluate_feature_flag.evaluate_feature_flags["f2"].result_string
  description = "Feature flag evaluated values."
}
output "ibm_app_config_property_evaluated_value" {
  value       = data.ibm_app_config_evaluate_property.evaluate_property.result_numeric
  description = "Property evaluated value."
}
output "ibm_app_config_property_evaluated_values" {
  value       = data.ibm_app_config_evaluate_property.evaluate_properties["p2"].result_boolean
  description = "Property evaluated values."
}
