provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_trusted_profile resource instance
resource "ibm_iam_trusted_profile" "iam_trusted_profile_instance" {
  name = var.iam_trusted_profile_name
  description = "description"
}

// Create iam_trusted_profile data source
data "ibm_iam_trusted_profile" "iam_trusted_profile_instance_data" {
  profile_id = ibm_iam_trusted_profile.iam_trusted_profile_instance.id
  include_activity = false
}

// Create iam_trusted_profiles data source
data "ibm_iam_trusted_profiles" "iam_trusted_profiles_list_data" {
  account_id = var.iam_trusted_profiles_account_id
  include_history = false
}