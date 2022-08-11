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
resource "ibm_iam_access_group" "accgrp" {
	name = var.agname
}

data "ibm_iam_roles" "test" {
	service = var.servicename
  }
resource "ibm_iam_access_group_policy" "policy" {
	access_group_id = ibm_iam_access_group.accgrp.id
	roles           = [ibm_iam_custom_role.customrole.display_name,"Viewer"]
	tags            = ["tag1"]
	resources {
	  service = var.servicename
	}
}