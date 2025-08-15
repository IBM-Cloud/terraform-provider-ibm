provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.service_region
}

data "ibm_resource_instance" "cloudant-dedicated-cluster" {
  name     = var.cloudant_dedicated_hardware_name
}

// Provision cloudant resource instance with Standard plan and dedicated hardware Cloud Resource Name (CRN)
resource "ibm_cloudant" "cloudant" {
  // Required arguments:
  name     = "test_standard_plan_on_dedicated_hw_cloudant"
  location = var.service_region
  plan     = "standard"
  // Optional arguments (existing dedicated cluster):
  environment_crn = data.ibm_resource_instance.cloudant-dedicated-cluster.id
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}
