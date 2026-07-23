provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_iam_identity_preference" "iam_identity_preference_instance_left_nav" {
  account_id = var.iam_identity_preference_account_id
  iam_id = var.iam_identity_preference_iam_id
  service = var.iam_identity_preference_service
  preference_id = "global_left_navigation"
  value_list_of_strings = var.iam_identity_preference_value_list_of_strings
}

resource "ibm_iam_identity_preference" "iam_identity_preference_instance_landing" {
  account_id = var.iam_identity_preference_account_id
  iam_id = var.iam_identity_preference_iam_id
  service = var.iam_identity_preference_service
  preference_id = "landing_page"
  value_string = var.iam_identity_preference_value_string
}

// Create iam_identity_preference data source
data "ibm_iam_identity_preference" "iam_identity_preference_instance_data" {
  account_id = var.iam_identity_preference_account_id
  iam_id = var.iam_identity_preference_iam_id
  service = var.iam_identity_preference_service
  preference_id = "global_left_navigation"

  depends_on = [ibm_iam_identity_preference.iam_identity_preference_instance_left_nav]
}

// Create iam_identity_preferences data source
data "ibm_iam_identity_preferences" "iam_identity_preferences_instance_list" {
  account_id = var.iam_identity_preference_account_id
  iam_id = var.iam_identity_preference_iam_id

  depends_on = [ibm_iam_identity_preference.iam_identity_preference_instance_left_nav, ibm_iam_identity_preference.iam_identity_preference_instance_landing]
}
