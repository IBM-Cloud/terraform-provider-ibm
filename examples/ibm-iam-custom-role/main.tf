provider "ibm" {
}

data "ibm_iam_role_actions" "test" {
  service = var.servicename
}

resource "ibm_iam_custom_role" "customrole" {
  name         = var.name
  display_name = var.displayname
  description  = var.description
  service = var.servicename
  actions      = [data.ibm_iam_role_actions.test.manager.18]
}
