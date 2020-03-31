resource "ibm_resource_instance" "appid_instance" {
  name     = "${var.name}"
  service  = "appid"
  plan     = "${var.plan}"
  location = "${var.location}"
}
