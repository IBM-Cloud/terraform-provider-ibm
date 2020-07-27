provider "ibm" {
  generation = 1
}
data "ibm_resource_group" "group" {
  name = var.name
}

resource "ibm_resource_instance" "resource_instance" {
  name              = var.service_name
  service           = var.service_type
  plan              = var.plan
  location          = var.location
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["tag1", "tag2"]

  parameters = {
    "HMAC" = true
  }
  //User can increase timeouts 
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

