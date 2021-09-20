
module "satellite-location" {
  source = "./modules/location"

  is_location_exist = var.is_location_exist
  location          = var.location
  managed_from      = var.managed_from
  location_zones    = var.location_zones
  location_bucket   = var.location_bucket
  ibm_region        = var.ibm_region
  resource_group    = var.resource_group
  host_labels       = var.host_labels
}