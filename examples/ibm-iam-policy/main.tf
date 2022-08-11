provider "ibm" {
}

data "ibm_resource_group" "group" {
  name = var.resource_group
}

resource "ibm_iam_user_policy" "iam_policy" {
  ibm_id = var.ibm_id1
  roles  = ["Viewer", "Editor"]

  resources {
    resource_type = "resource-group"
    resource      = data.ibm_resource_group.group.id
  }
}
