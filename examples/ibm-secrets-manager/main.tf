provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision secret_group resource instance
resource "ibm_secret_group" "secret_group_instance" {
  name = var.secret_group_name
  description = var.secret_group_description
}
