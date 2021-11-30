
data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_satellite_location" "create_location" {
  count = var.is_location_exist == false ? 1 : 0

  location          = var.location
  managed_from      = var.managed_from
  zones             = var.location_zones
  resource_group_id = data.ibm_resource_group.resource_group.id

  cos_config {
    bucket = var.location_bucket != null ? var.location_bucket : null
  }

  tags = var.tags
}

data "ibm_satellite_location" "location" {
  location   = var.is_location_exist == false ? ibm_satellite_location.create_location.0.id : var.location
  depends_on = [ibm_satellite_location.create_location]
}

data "ibm_satellite_attach_host_script" "script" {
  location      = data.ibm_satellite_location.location.id
  labels        = var.host_labels
  host_provider = var.host_provider
}

output "location_id" {
  value = data.ibm_satellite_location.location.id
}

output "location_crn" {
  value = data.ibm_satellite_location.location.crn
}

output "host_script" {
  value = data.ibm_satellite_attach_host_script.script.host_script
}