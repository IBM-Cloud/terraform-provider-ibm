data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_satellite_location_nlb_dns" "location_dns_instance" {
  location = var.location
  ips      = var.control_plane_ips
}

data "ibm_satellite_location_nlb_dns" "location_dns_instance" {
  location = ibm_satellite_location_nlb_dns.location_dns_instance.location
}

data "ibm_container_nlb_dns" "dns" {
  cluster = var.cluster
}

resource "ibm_container_nlb_dns" "container_nlb_dns" {
  cluster           = var.cluster
  nlb_host          = data.ibm_container_nlb_dns.dns.nlb_config.0.nlb_sub_domain
  nlb_ips           = var.cluster_ips
  resource_group_id = data.ibm_resource_group.resource_group.id
}