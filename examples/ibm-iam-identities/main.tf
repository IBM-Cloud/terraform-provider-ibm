provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_trusted_profile_identities resource instance
resource "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities_instance" {
  profile_id = var.iam_trusted_profile_identities_profile_id

  dynamic "identities" {
    for_each = var.iam_trusted_profile_identities
    content {
      iam_id      = identities.value.iam_id
      type        = identities.value.type
      identifier  = identities.value.identifier
      accounts    = identities.value.accounts
      description = identities.value.description
    }
  }
}

// Create iam_trusted_profile_identities data source
data "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities_instance" {
  profile_id = var.iam_trusted_profile_identities_profile_id
}

