provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_trusted_profiles_link resource instance
resource "ibm_iam_trusted_profiles_link" "iam_trusted_profiles_link_instance" {
  profile_id = var.iam_trusted_profiles_link_profile_id
  cr_type = var.iam_trusted_profiles_link_cr_type
  link {
    crn = "crn"
    namespace = "namespace"
    name = "name"
  }
  name = var.iam_trusted_profiles_link_name
}

// Create iam_trusted_profiles_link data source
data "ibm_iam_trusted_profiles_link" "iam_trusted_profiles_link_instance" {
  profile_id = var.iam_trusted_profiles_link_profile_id
  link_id = var.iam_trusted_profiles_link_link_id
}
