module "satellite-cluster" {
  source = "./modules/cluster"

  cluster           = var.cluster
  location          = module.satellite-location.location_id
  kube_version      = var.kube_version
  default_wp_labels = var.default_wp_labels
  zones             = var.cluster_zones
  resource_group    = var.resource_group
  worker_pool_name  = var.worker_pool_name
  worker_count      = var.worker_count
  workerpool_labels = var.workerpool_labels
  cluster_tags      = var.cluster_tags
  host_labels       = var.host_labels
  zone_name         = var.zone_name

  depends_on = [module.satellite-host]
}