variable "ibmcloud_api_key" {
  description = "holds the user api key"
}

data "ibm_resource_group" "rg" {
  name = "default"
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  generation       = 2
  region           = "us-south"
}

resource "ibm_is_vpc" "test_pdns_vpc" {
  name           = "test-pdns-vpc"
  resource_group = data.ibm_resource_group.rg.id
}

resource "ibm_resource_instance" "test-pdns-instance" {
  name              = "test-pdns"
  resource_group_id = data.ibm_resource_group.rg.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "free-plan"
}

resource "ibm_dns_zone" "test-pdns-zone" {
  name        = "test.com"
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  description = "testdescription"
  label       = "testlabel"
}

resource "ibm_dns_permitted_network" "test-pdns-permitted-network" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = element(split("/", ibm_dns_zone.test-pdns-zone.id), 1)
  vpc_crn     = ibm_is_vpc.test_pdns_vpc.resource_crn
}

resource "ibm_private_dns_record" "pdns-1-records" {
  instance_id  = ibm_resource_instance.test-pdns-instance.guid
  zone_id      = element(split("/", ibm_dns_zone.test-pdns-zone.id), 1)
  type         = "A"
  ttl          = 900
  ipv4_address = "1.2.3.4"
  name         = "example"
}
