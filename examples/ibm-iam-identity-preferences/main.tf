provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Import an existing Identity preference, creation is not supported
import {
  to = ibm_iam_identity_preference.iam_identity_preference_instance
  id = "${var.iam_identity_preference_account_id}/${var.iam_identity_preference_iam_id}/${var.iam_identity_preference_service}/${var.iam_identity_preference_preference_id}"
}
resource "ibm_iam_identity_preference" "iam_identity_preference_instance" {
  account_id = var.iam_identity_preference_account_id
  service = var.iam_identity_preference_service
  preference_id = var.iam_identity_preference_preference_id
  value_string = "/iam"
}

// Create iam_identity_preference data source
data "ibm_iam_identity_preference" "iam_identity_preference_instance_data" {
  account_id = var.iam_identity_preference_account_id
  iam_id = var.iam_identity_preference_iam_id
  service = var.iam_identity_preference_service
  preference_id = var.iam_identity_preference_preference_id
  
  depends_on = [ibm_iam_identity_preference.iam_identity_preference_instance]
}

// Create iam_identity_preferences data source
data "ibm_iam_identity_preferences" "iam_identity_preferences_instance_list" {
  account_id = var.iam_identity_preference_account_id
  iam_id = var.iam_identity_preference_iam_id
  
  depends_on = [ibm_iam_identity_preference.iam_identity_preference_instance]
}
