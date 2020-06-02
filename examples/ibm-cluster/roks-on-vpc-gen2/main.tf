module "cluster_and_workerpool" {
  source = "./roks_on_vpc"
}

module "cos_service_binding" {
  source =  "./cloud_object_storage"
  cluster_id = module.cluster_and_workerpool.cluster_id
}