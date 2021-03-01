provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.service_region
}

// Provision cloudant resource instance with Lite plan
resource "ibm_cloudant" "cloudant" {
  // Required arguments:
  name     = "test_lite_plan_iam_cloudant"
  location = var.service_region
  plan     = "lite"
}

// Create cloudant data source
data "ibm_cloudant" "cloudant" {
  name     = ibm_cloudant.cloudant.name
}

// Provision Viewer IAM credential for our cloudant resource instance
// See also https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_key
resource "ibm_resource_key" "resourceKey" {
  name                 = "cloudantViewerCredentials"
  // Valid roles are Writer, Reader, Manager, Administrator, Operator, Viewer, and Editor:
  role                 = "Viewer"
  resource_instance_id = data.ibm_cloudant.cloudant.id
}
