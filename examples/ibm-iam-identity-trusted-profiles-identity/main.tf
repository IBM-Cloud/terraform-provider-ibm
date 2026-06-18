provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_trusted_profile_identity resource instance
resource "ibm_iam_trusted_profile_identity" "iam_trusted_profile_identity_instance" {
  profile_id = var.iam_trusted_profile_identity_profile_id
  identity_type = var.iam_trusted_profile_identity_identity_type
  identifier = var.iam_trusted_profile_identity_identifier
  type = var.iam_trusted_profile_identity_type
  accounts = var.iam_trusted_profile_identity_accounts
  description = "description"
}

// Create iam_trusted_profile_identity data source
data "ibm_iam_trusted_profile_identity" "iam_trusted_profile_identity_data" {
  profile_id = ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance.profile_id
  identity_type = ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance.identity_type
  identifier_id = ibm_iam_trusted_profile_identity.iam_trusted_profile_identity_instance.identifier
}

// Create iam_trusted_profile_identities data source
data "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities_data" {
  profile_id = var.iam_trusted_profile_identity_profile_id
}
