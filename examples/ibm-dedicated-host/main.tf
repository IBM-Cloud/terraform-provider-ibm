data "ibm_container_dedicated_host_flavor" "dhostflavor" {
  id   = var.dhostflavorid
  zone = var.zone
}

resource "ibm_container_dedicated_host_pool" "dhostpool" {
  name         = var.dhostpoolname
  flavor_class = ibm_container_dedicated_host_flavor.dhostflavor.flavor_class
  metro        = var.metro
}

resource "ibm_container_vpc_cluster" "vpc_dhost_cluster" {
  name   = var.clustername
  vpc_id = var.vpcid
  flavor = var.dhostflavorid
  zones {
    subnet_id = var.subnetid
    name      = var.zone
  }
  resource_group_id = var.resourcegroupid
  host_pool_id      = var.dhostpoolid
}

data "ibm_container_vpc_cluster" "vpc_dhost_cluster" {
  name = var.clustername
}