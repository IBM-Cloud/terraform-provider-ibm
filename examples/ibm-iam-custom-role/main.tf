provider "ibm" {
}


resource "ibm_iam_custom_role" "customrole" {
  name         = var.name
  display_name = var.displayname
  description  = var.description
  service = var.servicename
  actions      = [var.action]
}
