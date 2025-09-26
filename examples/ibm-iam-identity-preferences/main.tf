provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_identity_preference resource instance
resource "ibm_iam_identity_preference" "iam_identity_preference_instance" {
  account_id = var.iam_identity_preference_account_id
  iam_id = var.iam_identity_preference_iam_id
  service = var.iam_identity_preference_service
  preference_id = var.iam_identity_preference_preference_id
  value_string = var.iam_identity_preference_value_string
  value_list_of_strings = var.iam_identity_preference_value_list_of_strings
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create iam_identity_preference data source
data "ibm_iam_identity_preference" "iam_identity_preference_instance" {
  account_id = var.data_iam_identity_preference_account_id
  iam_id = var.data_iam_identity_preference_iam_id
  service = var.data_iam_identity_preference_service
  preference_id = var.data_iam_identity_preference_preference_id
}
*/
