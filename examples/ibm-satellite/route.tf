data "ibm_satellite_cluster" "ds_cluster" {
  name       = var.cluster
  depends_on = [module.satellite-dns]
}

module "satellite-route" {
  source = "./modules/route"

  is_endpoint_provision = var.is_endpoint_provision
  ibmcloud_api_key      = var.ibmcloud_api_key
  cluster_master_url    = data.ibm_satellite_cluster.ds_cluster.server_url
  route_name            = var.route_name
}