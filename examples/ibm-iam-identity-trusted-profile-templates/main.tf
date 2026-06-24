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

// Create trusted_profile_template data source
data "ibm_iam_trusted_profile_template" "trusted_profile_template_data" {
  template_id = split("/", ibm_iam_trusted_profile_template.trusted_profile_template_instance.id)[0]
  version = ibm_iam_trusted_profile_template.trusted_profile_template_instance.version
  include_history = false
}
