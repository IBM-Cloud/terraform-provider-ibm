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
  no_subnet    = true
  subnet_id    = ["${var.subnet_id}"]
  worker_num   = 2

  machine_type    = "${var.machine_type}"
  isolation       = "${var.isolation}"
  public_vlan_id  = "${var.public_vlan_id}"
  private_vlan_id = "${var.private_vlan_id}"
}

resource "ibm_service_instance" "service" {
  name       = "${var.service_instance_name}${random_id.name.hex}"
  space_guid = "${data.ibm_space.space.id}"
  service    = "${var.service_offering}"
  plan       = "${var.plan}"
  tags       = ["my-service"]
}

resource "ibm_service_key" "key" {
  name                  = "${var.service_key}"
  service_instance_guid = "${ibm_service_instance.service.id}"
}

resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id     = "${ibm_container_cluster.cluster.name}"
  service_instance_id = "${ibm_service_instance.service.id}"
  namespace_id        = "default"
  org_guid            = "${data.ibm_org.org.id}"
  space_guid          = "${data.ibm_space.space.id}"
  account_guid        = "${data.ibm_account.account.id}"
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = "${ibm_container_cluster.cluster.name}"
  org_guid        = "${data.ibm_org.org.id}"
  space_guid      = "${data.ibm_space.space.id}"
  account_guid    = "${data.ibm_account.account.id}"
}

resource "random_id" "name" {
  byte_length = 4
}
