provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
  name = "Temp name"
  committed = true
  account_settings {
    mfa = "LEVEL3"
  }
}

resource "ibm_iam_account_settings_template_assignment" "account_settings_template_assignment_instance" {
  template_id = split("/", ibm_iam_account_settings_template.account_settings_template_instance.id)[0]
  template_version = ibm_iam_account_settings_template.account_settings_template_instance.version
  target_type = var.account_settings_template_assignment_target_type
  target = var.account_settings_template_assignment_target
  depends_on = [
    ibm_iam_account_settings_template.account_settings_template_instance
  ]
}

data "ibm_iam_account_settings_template_assignment" "account_settings_template_assignment_instance" {
  assignment_id = ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance.id
}