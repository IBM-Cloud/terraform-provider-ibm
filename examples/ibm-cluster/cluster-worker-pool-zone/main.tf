provider "ibm" {
}

data "ibm_resource_group" "testacc_ds_resource_group" {
  name = var.resource_group
}

resource "ibm_container_cluster" "cluster" {
  name              = "${var.cluster_name}${random_id.name.hex}"
  datacenter        = var.datacenter
  no_subnet         = true
  subnet_id         = [var.subnet_id]
  default_pool_size = 2
  hardware          = "shared"
  resource_group_id = data.ibm_resource_group.testacc_ds_resource_group.id
  machine_type      = var.machine_type
  public_vlan_id    = var.public_vlan_id
  private_vlan_id   = var.private_vlan_id
}

resource "ibm_container_worker_pool_zone_attachment" "default_zone" {
  cluster         = ibm_container_cluster.cluster.id
  worker_pool     = "default"
  zone            = var.zone
  public_vlan_id  = var.zone_public_vlan_id
  private_vlan_id = var.zone_private_vlan_id
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "${var.worker_pool_name}${random_id.name.hex}"
  machine_type     = var.machine_type
  cluster          = ibm_container_cluster.cluster.id
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = "true"

  labels = {
    "test" = "test-pool"
  }
}

resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster         = ibm_container_cluster.cluster.id
  worker_pool     = element(split("/", ibm_container_worker_pool.test_pool.id), 1)
  zone            = var.zone
  public_vlan_id  = var.zone_public_vlan_id
  private_vlan_id = var.zone_private_vlan_id
}

resource "ibm_resource_instance" "cos_instance" {
  name     = var.service_instance_name
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id     = ibm_container_vpc_cluster.cluster.id
  service_instance_id = element(split(":", ibm_resource_instance.cos_instance.id), 7)
  namespace_id        = "default"
  role                = "Writer"
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = ibm_container_cluster.cluster.id
}

resource "random_id" "name" {
  byte_length = 4
}

