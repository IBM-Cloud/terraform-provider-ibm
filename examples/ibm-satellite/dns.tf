module "satellite-dns" {
  source = "./modules/dns"

  location          = module.satellite-location.location_id
  cluster           = module.satellite-cluster.cluster_id
  control_plane_ips = slice(ibm_is_floating_ip.satellite_ip[*].address, 0, var.host_count)
  cluster_ips       = slice(ibm_is_floating_ip.satellite_ip[*].address, var.host_count, (var.host_count + var.addl_host_count))
  resource_group    = var.resource_group
  depends_on        = [module.satellite-cluster]
}