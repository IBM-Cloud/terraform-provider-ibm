provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create account data source
data "ibm_account_info" "account_instance" {
  account_id = var.account_account_id
}
*/
