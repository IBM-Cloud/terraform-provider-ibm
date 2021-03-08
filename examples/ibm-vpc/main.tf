provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Create is_dedicated_hosts data source
data "ibm_is_dedicated_hosts" "is_dedicated_hosts_instance" {
  name = var.is_dedicated_hosts_name
}
