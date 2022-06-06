provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.service_region
}

// Provision cloudant resource instance with Standard plan and 2 capacity throughput blocks
resource "ibm_cloudant" "cloudant" {
  // Required arguments:
  name     = "test_standard_plan_cloudant"
  location = var.service_region
  plan     = "standard"
  // Optional arguments:
  capacity = 2
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}

// Create database with
resource "ibm_cloudant_database" "database" {
  db           = var.db_name
  instance_crn = ibm_cloudant.cloudant.crn
}

data "ibm_cloudant_database" "database" {
  db           = ibm_cloudant_database.database.db
  instance_crn = ibm_cloudant_database.database.instance_crn
}