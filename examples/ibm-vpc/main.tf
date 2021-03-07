provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision is_dedicated_host_group resource instance
resource "ibm_is_dedicated_host_group" "is_dedicated_host_group_instance" {
  class = var.is_dedicated_host_group_class
  family = var.is_dedicated_host_group_family
  name = var.is_dedicated_host_group_name
  resource_group = var.is_dedicated_host_group_resource_group
  zone = var.is_dedicated_host_group_zone
}
