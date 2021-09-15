provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_trusted_profiles resource instance
resource "ibm_iam_trusted_profiles" "iam_trusted_profiles_instance" {
  name = "name"
  description = "description"
  account_id = "account_id"
}

// Create iam_trusted_profiles data source
data "ibm_iam_trusted_profiles" "iam_trusted_profiles_instance" {
  profile_id = var.iam_trusted_profiles_profile_id
}