provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision is_dedicated_host resource instance
resource "ibm_is_dedicated_host" "is_dedicated_host_instance" {
  dedicated_host_prototype = var.is_dedicated_host_dedicated_host_prototype
}
