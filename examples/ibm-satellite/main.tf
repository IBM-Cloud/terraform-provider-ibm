
module "satellite-location" {
  source            = "./modules/location"

  is_location_exist = var.is_location_exist
  location          = var.location
  managed_from      = var.managed_from
  location_zones    = var.location_zones
  ibmcloud_api_key  = var.ibmcloud_api_key
  ibm_region        = var.ibm_region
  resource_group    = var.resource_group
  tags              = var.tags
}

data "ibm_satellite_attach_host_script" "script" {
  location          = module.satellite-location.location_id
  labels            = var.host_labels
  host_provider     = "ibm"
}