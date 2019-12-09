provider "ibm" {}

resource "ibm_iam_access_group" "accgrp" {
  name        = "test"
  description = "New access group"
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_iam_access_group_policy" "policy" {
  access_group_id = "${ibm_iam_access_group.accgrp.id}"
  roles        = ["Operator", "Writer"]

  resources = [{
    resource_group_id = "${data.ibm_resource_group.group.id}"
  }]
}

resource "ibm_iam_user_invite" "invite_user" {
    users = ["${var.user1}", "${var.user2}"]
    access_groups = ["${ibm_iam_access_group.accgrp.id}"]

}