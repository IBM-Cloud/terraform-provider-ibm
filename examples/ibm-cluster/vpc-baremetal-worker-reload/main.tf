resource "random_id" "name" {
  byte_length = 2
}

locals {
  ZONE = "${var.region}-1"
}

resource "ibm_is_vpc" "vpc" {
  name = "vpc-${random_id.name.hex}"
}

resource "ibm_is_subnet" "subnet" {
  name                     = "subnet-${random_id.name.hex}"
  vpc                      = ibm_is_vpc.vpc.id
  zone                     = local.ZONE
  total_ipv4_address_count = 256
}

data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}


resource "ibm_container_vpc_cluster" "cluster" {
  name              = "${var.cluster_name}-${random_id.name.hex}"
  vpc_id            = ibm_is_vpc.vpc.id
  flavor            = var.flavor
  worker_count      = 1
  resource_group_id = data.ibm_resource_group.resource_group.id
  wait_till         = "OneWorkerNodeReady"

  zones {
    subnet_id = ibm_is_subnet.subnet.id
    name      = local.ZONE
  }
}

data "ibm_container_vpc_cluster" "cluster_data" {
  name              = ibm_container_vpc_cluster.cluster.name
  resource_group_id = data.ibm_resource_group.resource_group.id
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = ibm_container_vpc_cluster.cluster.id
}

// Action: Reload bare metal worker and wait for completion
// Invoke with: terraform apply -invoke action.ibm_container_vpc_bare_metal_worker_reload.reload
action "ibm_container_vpc_bare_metal_worker_reload" "reload" {
  config {
    cluster_name_id      = ibm_container_vpc_cluster.cluster.id
    bare_metal_server_id = data.ibm_container_vpc_cluster.cluster_data.workers[0]
  }
}

// Action: Reload bare metal worker without waiting (fire-and-forget)
// Invoke with: terraform apply -invoke action.ibm_container_vpc_bare_metal_worker_reload.reload_no_wait
action "ibm_container_vpc_bare_metal_worker_reload" "reload_no_wait" {
  config {
    cluster_name_id      = ibm_container_vpc_cluster.cluster.id
    bare_metal_server_id = data.ibm_container_vpc_cluster.cluster_data.workers[0]
    no_wait              = true
  }
}