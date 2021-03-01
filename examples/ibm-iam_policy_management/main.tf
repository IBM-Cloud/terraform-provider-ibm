provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_policy resource instance
resource "ibm_iam_policy" "iam_policy_instance" {
  type = var.iam_policy_type
  subjects = var.iam_policy_subjects
  roles = var.iam_policy_roles
  resources = var.iam_policy_resources
  description = var.iam_policy_description
  accept_language = var.iam_policy_accept_language
}

// Provision iam_custom_role resource instance
resource "ibm_iam_custom_role" "iam_custom_role_instance" {
  display_name = var.iam_custom_role_display_name
  actions = var.iam_custom_role_actions
  name = var.iam_custom_role_name
  account_id = var.iam_custom_role_account_id
  service_name = var.iam_custom_role_service_name
  description = var.iam_custom_role_description
  accept_language = var.iam_custom_role_accept_language
}

// Create iam_policy data source
data "ibm_iam_policy" "iam_policy_instance" {
  policy_id = var.iam_policy_policy_id
}

// Create iam_custom_role data source
data "ibm_iam_custom_role" "iam_custom_role_instance" {
  role_id = var.iam_custom_role_role_id
}
