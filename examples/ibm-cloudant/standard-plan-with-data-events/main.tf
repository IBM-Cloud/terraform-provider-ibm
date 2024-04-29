provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.service_region
}

// Provision cloudant resource instance with Standard plan including data events
resource "ibm_cloudant" "cloudant" {
  // Required arguments:
  name     = "test_standard_plan_with_data_events_cloudant"
  location = var.service_region
  plan     = "standard"
  // Optional arguments:
  include_data_events = true
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}

// Create resource group data source
data "ibm_resource_group" "group" {
  is_default = "true"
}

// Provision activity tracker event routing using the resource group data source
resource "ibm_resource_instance" "at_instance" {
  name              = "test_standard_plan_with_data_events_at"
  service           = "logdnaat"
  plan              = "7-day"
  location          = var.service_region
  resource_group_id = data.ibm_resource_group.group.id
}
