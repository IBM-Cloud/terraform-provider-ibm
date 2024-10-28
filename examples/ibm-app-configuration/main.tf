provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}
resource "ibm_app_config_feature" "string_flag" {
    disabled_value = "string1"
    enabled_value  = "string2"
    environment_id = "test"
    feature_id     = "string_flag"
    guid           = "43f65e31-fd51-4fe1-8801-7a40b226df7f"
    name           = "test_string_flag"
    type           = "STRING"
    format         = "TEXT"
}