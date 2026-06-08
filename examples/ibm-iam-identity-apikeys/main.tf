provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_api_key resource instance
resource "ibm_iam_api_key" "iam_api_key_instance" {
  name = var.iam_api_key_name
  description = "apikey description"
  entity_lock = var.iam_api_key_entity_lock
  store_value = var.iam_api_key_store_value
  file = var.iam_api_key_file
}

// Read iam_api_key data source
data "ibm_iam_api_key" "iam_api_key_data" {
  apikey_id= ibm_iam_api_key.iam_api_key_instance.id
}
