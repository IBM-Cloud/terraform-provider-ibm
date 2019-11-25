provider "ibm" {
  generation = 1
}
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_resource_instance" "resource_instance" {
  name              = var.service_instance_name
  service           = "cloud-object-storage"
  plan              = "lite"
  location          = "global"
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

/*resource "ibm_resource_instance" "cos_instance" {
  name     = var.service_instance_name
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}*/

