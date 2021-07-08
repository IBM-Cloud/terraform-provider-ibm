module "satellite-link" {
  source = "./modules/link"

  location   = module.satellite-location.location_id
  crn        = module.satellite-location.location_crn
  depends_on = [module.satellite-cluster]
}