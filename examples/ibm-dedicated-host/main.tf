data "ibm_container_dedicated_host_flavor" "dhostflavor" {
  host_flavor_id = var.dhostflavorid
  zone           = var.zone
}

resource "ibm_container_dedicated_host_pool" "dhostpool" {
  name              = var.dhostpoolname
  flavor_class      = data.ibm_container_dedicated_host_flavor.dhostflavor.flavor_class
  metro             = var.metro
  resource_group_id = var.resource_group_id
}

resource "ibm_container_dedicated_host" "dhost" {
  flavor       = data.ibm_container_dedicated_host_flavor.dhostflavor.host_flavor_id
  host_pool_id = ibm_container_dedicated_host_pool.dhostpool.id
  zone         = var.zone
}

data "ibm_is_vpc" "vpc" {
  name = var.vpc_name
}

data "ibm_is_subnet" "subnet" {
  name = var.subnet_name
}

resource "ibm_container_vpc_cluster" "dhost_vpc_cluster" {
  name   = var.cluster_name
  vpc_id = data.ibm_is_vpc.vpc.id
  flavor = var.flavor
  zones {
    name      = var.zone
    subnet_id = data.ibm_is_subnet.subnet.id
  }
  worker_count      = var.worker_count
  resource_group_id = var.resource_group_id
  wait_till         = "OneWorkerNodeReady"
  host_pool_id      = ibm_container_dedicated_host_pool.dhostpool.id

  depends_on        = [
    ibm_container_dedicated_host.dhost
    ]
}

resource "ibm_container_vpc_worker_pool" "dhost_vpc_worker_pool" {
  cluster          = ibm_container_vpc_cluster.dhost_vpc_cluster.name
  worker_pool_name = var.worker_pool_name
  flavor           = ibm_container_vpc_cluster.dhost_vpc_cluster.flavor
  vpc_id           = ibm_container_vpc_cluster.dhost_vpc_cluster.vpc_id
  host_pool_id     = ibm_container_vpc_cluster.dhost_vpc_cluster.host_pool_id
  worker_count     = var.worker_count
  zones {
    name      = var.zone
    subnet_id = data.ibm_is_subnet.subnet.id
  }
  depends_on = [
    ibm_container_dedicated_host.dhost
  ]
}

