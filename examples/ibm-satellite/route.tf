module "satellite-route" {
  source                = "./modules/route"
  
  is_endpoint_provision = var.is_endpoint_provision
  ibmcloud_api_key      = var.ibmcloud_api_key
  cluster_master_url    = module.satellite-cluster.master_url
  route_name            = var.route_name
}