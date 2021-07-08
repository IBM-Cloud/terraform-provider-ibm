/*

//By default satellite link exist in location. If satellite link doesn't exist, Use 'ibm_satellite_link' resource to provision
// Provision satellite_link resource instance
resource "ibm_satellite_link" "satellite_link" {
  location = var.location
  crn      = var.crn
}

*/

data "ibm_satellite_link" "satellite_link" {
  location = var.location
}