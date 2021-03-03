provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Create is_dedicated_host_profile data source
data "ibm_is_dedicated_host_profile" "is_dedicated_host_profile_instance" {
  name = var.is_dedicated_host_profile_name
}

// Create is_dedicated_host_profiles data source
data "ibm_is_dedicated_host_profiles" "is_dedicated_host_profiles_instance" {
  name = var.is_dedicated_host_profiles_name
}
