// Provision satellite_storage_assignment resource instance
resource "ibm_satellite_storage_assignment" "instance" {
  assignment_name = var.assignment_name
  cluster = var.cluster
  config = var.config
  controller = var.controller
}