provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
  name = var.account_settings_template_name
  account_settings {

  }
}

resource "ibm_iam_account_settings_template" "account_settings_template_new_version" {
  template_id = ibm_iam_account_settings_template.account_settings_template_instance.id
  name = var.account_settings_template_name
  description = "Description for version 2"
#  committed = true
  account_settings {
    mfa = "LEVEL3"
  }
}

// data source for a pre-existing template
#data "ibm_iam_account_settings_template" "account_settings_template_instance" {
#  template_id = var.account_settings_template_template_id
#  version = var.account_settings_template_version
#}