provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_iam_trusted_profile_template" "trusted_profile_template_instance" {
  name = var.trusted_profile_template_name
  description = "description"
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
  policy_template_references {
    id      = "policyTemplate-1c748620-8a08-4fd2-a1f8-86992ce12c1c"
    version = "1"
  }
}

// Create trusted_profile_template data source
data "ibm_iam_trusted_profile_template" "trusted_profile_template_instance_data" {
  template_id = ibm_iam_trusted_profile_template.trusted_profile_template_instance.id
  version = ibm_iam_trusted_profile_template.trusted_profile_template_instance.version
  include_history = true

  depends_on = [ibm_iam_trusted_profile_template.trusted_profile_template_instance]
}

data "ibm_iam_trusted_profile_template" "trusted_profile_template_version_data" {
  template_id = ibm_iam_trusted_profile_template.trusted_profile_template_version.id
  version = ibm_iam_trusted_profile_template.trusted_profile_template_version.version
  include_history = true

  depends_on = [ibm_iam_trusted_profile_template.trusted_profile_template_version]
}
