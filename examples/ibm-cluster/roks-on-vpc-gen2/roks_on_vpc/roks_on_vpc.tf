resource "random_id" "name1" {
  byte_length = 2
}

resource "random_id" "name2" {
  byte_length = 2
}

locals {
  ZONE1 = "${var.region}-1"
  ZONE2 = "${var.region}-2"
}

resource "ibm_is_vpc" "vpc1" {
  name = "vpc-${random_id.name1.hex}"
}

resource "ibm_is_security_group_rule" "testacc_security_group_rule_tcp" {
    group = ibm_is_vpc.vpc1.default_security_group
    direction = "inbound"
    tcp {
        port_min = 30000
        port_max = 32767
    }
 }

resource "ibm_is_subnet" "subnet1" {
  name                     = "subnet-${random_id.name1.hex}"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = local.ZONE1
  total_ipv4_address_count = 256
}

resource "ibm_is_subnet" "subnet2" {
  name                     = "subnet-${random_id.name2.hex}"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = local.ZONE2
  total_ipv4_address_count = 256
}

data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "${var.cluster_name}${random_id.name1.hex}"
  vpc_id            = ibm_is_vpc.vpc1.id
  kube_version      = var.kube_version
  flavor            = var.flavor
  worker_count      = var.worker_count
  resource_group_id = data.ibm_resource_group.resource_group.id
  entitlement       = var.entitlement
  cos_instance_crn  = var.cos_instance_crn

  zones {
    subnet_id = ibm_is_subnet.subnet1.id
    name      = local.ZONE1
  }
}

resource "ibm_container_vpc_worker_pool" "cluster_pool" {
  cluster           = ibm_container_vpc_cluster.cluster.id
  worker_pool_name  = "${var.worker_pool_name}${random_id.name1.hex}"
  flavor            = var.flavor
  vpc_id            = ibm_is_vpc.vpc1.id
  worker_count      = var.worker_count
  resource_group_id = data.ibm_resource_group.resource_group.id
  entitlement       = var.entitlement
  zones {
    name      = local.ZONE2
    subnet_id = ibm_is_subnet.subnet2.id
  }
}
