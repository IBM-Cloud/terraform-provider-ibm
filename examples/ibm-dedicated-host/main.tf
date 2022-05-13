data "ibm_container_dedicated_host_flavor" "dhostflavor" {
  id   = var.dhostflavorid
  zone = var.zone
}

resource "ibm_container_dedicated_host_pool" "dhostpool" {
  name          = var.dhostpoolname
  flavor_class  = "${ibm_container_dedicated_host_flavor.dhostflavor.flavor_class}"
  metro         = var.metro
}
