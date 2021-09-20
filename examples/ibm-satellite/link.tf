data "ibm_satellite_location" "ds_location" {
  location = var.location

  depends_on = [module.satellite-dns]
}

module "satellite-link" {
  source = "./modules/link"

  location   = module.satellite-location.location_id
  crn        = data.ibm_satellite_location.ds_location.crn
  depends_on = [module.satellite-cluster]
}
