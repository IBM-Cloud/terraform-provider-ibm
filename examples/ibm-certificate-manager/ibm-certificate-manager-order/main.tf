provider "ibm" {
}
//creating CMS resource
resource "ibm_resource_instance" "cm" {
  name     = var.cms_name
  location = var.region
  service  = "cloudcerts"
  plan     = "free"
}
//creating CIS resource
resource "ibm_cis" "instance" {
  name = var.cis_name
  plan = var.cis_plan
}
//Creating CIS domain
resource "ibm_cis_domain" "example" {
  domain = var.domain
  cis_id = ibm_cis.instance.id
}
//ordering certificate on CMS using CIS
resource "ibm_certificate_manager_order" "cert" {
  certificate_manager_instance_id = ibm_resource_instance.cm.id
  name                            = var.order_name
  description                     = var.order_description
  domains                         = [ibm_cis_domain.example.domain]
  rotate_keys                     = var.rotate_key
  domain_validation_method        = var.dvm
  dns_provider_instance_crn       = ibm_cis.instance.id
}

