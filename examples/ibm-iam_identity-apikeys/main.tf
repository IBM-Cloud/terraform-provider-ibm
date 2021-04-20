provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Read iam_api_key data source
data "ibm_iam_api_key" "iam_api_key" {
  id= "ApiKey-toRead"
}

// Provision iam_api_key resource instance
resource "ibm_iam_api_key" "iam_api_key_instance" {
  name = "apikey name"
  iam_id = "IBMid"
  description = "apikey desc"
  account_id = "unique account id"
}


