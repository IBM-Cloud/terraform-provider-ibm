provider "ibm" {}

data "ibm_org" "org" {
  org = "${var.org}"
}

data "ibm_space" "space" {
  org   = "${var.org}"
  space = "${var.space}"
}

data "ibm_resource_group" "resource_group" {
  name = "Default"
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "${var.cluster_name}${random_id.name.hex}" 
  vpc_id            = "${var.vpc_id}"
  flavor            = "${var.flavor}"
  worker_count      = "${var.worker_count}"
  resource_group_id = "${data.ibm_resource_group.resource_group.id}"
  zones = [
      {
         subnet_id = "${var.subnet_id}"
         name = "${var.zone_name}"
      }
  ]
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
  cluster_name_id     = "${ibm_container_vpc_cluster.cluster.id}"
  service_instance_id = "${ibm_service_instance.service.id}"
  namespace_id        = "default"
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = "${ibm_container_vpc_cluster.cluster.id}"
}

resource "random_id" "name" {
  byte_length = 4
}
