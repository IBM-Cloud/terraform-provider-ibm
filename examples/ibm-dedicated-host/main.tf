data "ibm_container_dedicated_host_flavor" "dhostflavor" {
  host_flavor_id = var.dhostflavorid
  zone           = var.zone
}

resource "ibm_container_dedicated_host_pool" "dhostpool" {
  name         = var.dhostpoolname
  flavor_class = ibm_container_dedicated_host_flavor.dhostflavor.flavor_class
  metro        = var.metro
}

resource "ibm_container_dedicated_host" "dhost" {
  flavor       = ibm_container_dedicated_host_flavor.dhostflavor.host_flavor_id
  host_pool_id = ibm_container_dedicated_host_pool.dhostpool.host_id
  zone         = var.zone
}
