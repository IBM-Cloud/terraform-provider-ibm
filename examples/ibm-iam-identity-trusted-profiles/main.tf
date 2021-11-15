provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_trusted_profile resource instance
resource "ibm_iam_trusted_profile" "iam_trusted_profile_instance" {
  name = "name"
  description = "description"
}

// Create iam_trusted_profile data source
data "ibm_iam_trusted_profile" "iam_trusted_profile_instance" {
  profile_id = var.iam_trusted_profile_profile_id
}