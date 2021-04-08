resource "ibm_satellite_location" "create_location" {
  count         = var.is_location_exist == false ? 1 : 0

  location      = var.location
  managed_from  = var.managed_from
  zones         = var.location_zones

  cos_config {
    bucket  = var.location_bucket
    region  = var.ibm_region
  }
}

data "ibm_satellite_location" "location" {
  location       = var.location
  depends_on    = [ ibm_satellite_location.create_location ]
}

output "location_id" {
  value  = data.ibm_satellite_location.location.id
}