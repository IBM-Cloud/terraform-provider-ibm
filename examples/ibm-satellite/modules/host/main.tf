resource "ibm_satellite_host" "assign_host" {
  count         = var.host_count

  location      = var.location
  host_id       = element(var.host_vms, count.index)
  labels        = var.host_labels
  zone          = element(var.location_zones, count.index)
  host_provider = var.host_provider
}

