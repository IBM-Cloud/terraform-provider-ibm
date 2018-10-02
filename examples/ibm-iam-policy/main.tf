provider "ibm" {}

data "ibm_org" "org" {
  org = "${var.org}"
}

data "ibm_space" "space" {
  org   = "${var.org}"
  space = "${var.space}"
}

data "ibm_account" "account" {
  org_guid = "${data.ibm_org.org.id}"
}

resource "ibm_container_cluster" "cluster" {
  name         = "${var.cluster_name}${random_id.name.hex}"
  datacenter   = "${var.datacenter}"
  org_guid     = "${data.ibm_org.org.id}"
  space_guid   = "${data.ibm_space.space.id}"
  account_guid = "${data.ibm_account.account.id}"

  workers = [{
    name   = "worker1"
    action = "add"
  },
    {
      name   = "worker2"
      action = "add"
    },
    {
      name   = "worker3"
      action = "add"
    }]

  machine_type    = "${var.machine_type}"
  isolation       = "${var.isolation}"
  public_vlan_id  = "${var.public_vlan_id}"
  private_vlan_id = "${var.private_vlan_id}"
}

resource "ibm_iam_user_policy" "iam_policy" {
  provider = "ibm.iam"
  account_guid = "${data.ibm_account.account.id}"
  ibm_id       = "${var.ibm_id1}"
  roles        = ["viewer", "editor"]

  resources = [
    {
      service_name     = "${var.service_name}"
      service_instance = ["${ibm_container_cluster.cluster.id}"]
      space_guid        = "${data.ibm_space.space.id}"
      organization_guid = "${data.ibm_org.org.id}"
    },
  ]

}

resource "random_id" "name" {
  byte_length = 4
}
