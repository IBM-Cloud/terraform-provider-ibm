provider "ibm" {
}

data "ibm_space" "spacedata" {
  space = var.space
  org   = var.org
}

resource "ibm_service_instance" "service-instance" {
  name       = var.service_instance_name
  space_guid = data.ibm_space.spacedata.id
  service    = var.service
  plan       = var.plan
  tags       = ["cluster-service", "cluster-bind"]
}

resource "ibm_service_key" "serviceKey" {
  name                  = var.service_key_name
  service_instance_guid = ibm_service_instance.service-instance.id
}

