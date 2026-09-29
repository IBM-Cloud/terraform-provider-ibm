provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_instance data source
data "ibm_sm_instance" "sm_instance_instance" {
  instance_id = var.sm_instance_instance_id
}
*/

// Generate a Vault admin token for authenticating to your Vault Dedicated cluster.
resource "ibm_sm_admin_token" "sm_admin_token" {
  instance_id = "var.sm_instance_instance_id
}

