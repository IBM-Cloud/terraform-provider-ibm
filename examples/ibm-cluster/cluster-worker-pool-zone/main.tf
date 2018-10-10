provider "ibm" {}

data "ibm_org" "org" {
  org = "${var.org}"
}

data "ibm_space" "space" {
  org   = "${var.org}"
  space = "${var.space}"
}

data "ibm_resource_group" "testacc_ds_resource_group" {
  name = "default"
}

resource "ibm_container_cluster" "cluster" {
  name              = "${var.cluster_name}${random_id.name.hex}"
  datacenter        = "${var.datacenter}"
  no_subnet         = true
  subnet_id         = ["${var.subnet_id}"]
  default_pool_size = 2

  resource_group_id = "${data.ibm_resource_group.testacc_ds_resource_group.id}"
  machine_type      = "${var.machine_type}"
  isolation         = "${var.isolation}"
  public_vlan_id    = "${var.public_vlan_id}"
  private_vlan_id   = "${var.private_vlan_id}"
}

resource ibm_container_worker_pool_zone_attachment default_zone {
  cluster         = "${ibm_container_cluster.cluster.id}"
  worker_pool     = "default"
  zone            = "${var.zone}"
  public_vlan_id  = "${var.zone_public_vlan_id}"
  private_vlan_id = "${var.zone_private_vlan_id}"
}

resource ibm_container_worker_pool test_pool {
  worker_pool_name = "${var.worker_pool_name}${random_id.name.hex}"
  machine_type     = "${var.machine_type}"
  cluster          = "${ibm_container_cluster.cluster.id}"
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = "true"

  labels = {
    "test" = "test-pool"
  }
}

resource ibm_container_worker_pool_zone_attachment test_zone {
  cluster         = "${ibm_container_cluster.cluster.id}"
  worker_pool     = "${element(split("/",ibm_container_worker_pool.test_pool.id),1)}"
  zone            = "${var.zone}"
  public_vlan_id  = "${var.zone_public_vlan_id}"
  private_vlan_id = "${var.zone_private_vlan_id}"
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
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = "${ibm_container_cluster.cluster.name}"
}

resource "random_id" "name" {
  byte_length = 4
}
