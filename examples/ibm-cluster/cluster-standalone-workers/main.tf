provider "ibm" {}

data "ibm_space" "space" {
  org   = "${var.org}"
  space = "${var.space}"
}

resource "ibm_container_cluster" "cluster" {
  name         = "${var.cluster_name}${random_id.name.hex}"
  datacenter   = "${var.datacenter}"
  no_subnet    = true
  subnet_id    = ["${var.subnet_id}"]
  default_pool_size = 2
  
  machine_type    = "${var.machine_type}"
  hardware = "${var.hardware}"
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
  cluster_name_id     = "${ibm_container_cluster.cluster.id}"
  service_instance_id = "${ibm_service_instance.service.id}"
  namespace_id        = "default"
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = "${ibm_container_cluster.cluster.id}"
}

resource "random_id" "name" {
  byte_length = 4
}
