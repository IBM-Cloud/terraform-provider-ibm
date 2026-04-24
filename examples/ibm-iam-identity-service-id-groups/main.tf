provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_serviceid_group resource instance
resource "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
  account_id = var.iam_serviceid_group_account_id
  name = var.iam_serviceid_group_name
  description = var.iam_serviceid_group_description
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create iam_serviceid_group data source
data "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
  iam_serviceid_group_id = var.data_iam_serviceid_group_iam_serviceid_group_id
}
*/
