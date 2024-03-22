resource "ibm_resource_instance" "push_instance" {
  name     = var.name
  service  = "imfpush"
  plan     = var.plan
  location = var.location
}
