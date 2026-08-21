provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_api_key resource instance
resource "ibm_iam_api_key" "iam_api_key_instance" {
  name = var.iam_api_key_name
  description = "apikey desc"
  file = var.iam_api_key_file_path
}

// Read iam_api_key data source
data "ibm_iam_api_key" "iam_api_key_data" {
  apikey_id = ibm_iam_api_key.iam_api_key_instance.apikey_id
}
