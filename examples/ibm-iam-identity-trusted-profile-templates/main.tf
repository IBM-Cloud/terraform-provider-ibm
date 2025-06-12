provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_iam_trusted_profile_template" "trusted_profile_template_instance" {
  name = var.trusted_profile_template_name
  description = var.trusted_profile_template_description
  profile {
    name = "name"
    description = "description"
    rules {
      name = "name"
      type = "Profile-SAML"
      realm_name = "test-realm-101"
      expiration = 1
      conditions {
        claim = "claim"
        operator = "EQUALS"
        value = "\"value\""
      }
    }
  }
}

resource "ibm_iam_trusted_profile_template" "trusted_profile_template_version" {
  template_id = split("/", ibm_iam_trusted_profile_template.trusted_profile_template_instance.id)[0]
  name = var.trusted_profile_template_name
  description = "new description"
  committed = true
  profile {
    name = "name"
    description = "description"
  }
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create trusted_profile_template data source
data "ibm_iam_trusted_profile_template" "trusted_profile_template_instance" {
  template_id = var.trusted_profile_template_template_id
  version = var.trusted_profile_template_version
  include_history = var.trusted_profile_template_include_history
}
*/
