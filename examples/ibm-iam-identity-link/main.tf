provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_trusted_profile_link resource instance
resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link_instance" {
  profile_id = var.iam_trusted_profile_link_profile_id
  cr_type = var.iam_trusted_profile_link_cr_type
  link {
    crn =  var.iam_trusted_profile_link_crn
  }
  name = var.iam_trusted_profile_link_name
}

// Create iam_trusted_profile_link data source
data "ibm_iam_trusted_profile_link" "iam_trusted_profile_link_instance_data" {
  profile_id = var.iam_trusted_profile_link_profile_id
  link_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link_instance.link_id
}

// Create iam_trusted_profile_links data source
data "ibm_iam_trusted_profile_links" "iam_trusted_profile_links_data" {
  profile_id = var.iam_trusted_profile_link_profile_id
}
