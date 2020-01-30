provider "ibm"{
}
resource "ibm_resource_instance" "cm" {
  name     = "testname"
  location = "us-south"
  service  = "cloudcerts"
  plan     = "free"
}
# created cis instance should be provided with a valid domain
resource "ibm_cis_domain" "example" {
  domain = "example.com"
  cis_id = "${ibm_cis.instance.id}"
}
resource "ibm_cis" "instance" {
  name = "test-domain"
  plan = "standard"
}

resource "ibm_certificate_manager_order" "cert" {
  certificate_manager_instance_id = "${ibm_resource_instance.cm.id}"
  name                            = "test"
  description="test description"
  domains = ["example.com"] 
  rotate_keys=false
  domain_validation_method= "dns-01" 
  dns_provider_instance_crn = "${ibm_cis.instance.id}"
}
