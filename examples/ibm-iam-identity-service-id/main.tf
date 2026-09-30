provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_service_id resource instance
resource "ibm_iam_service_id" "iam_service_id_instance" {
  name = var.iam_service_id_name
  description = "serviceId description"
}

// iam_service_id data source
data "ibm_iam_service_id" "iam_service_id_data" {
  name = ibm_iam_service_id.iam_service_id_instance.name

  depends_on = [ibm_iam_service_id.iam_service_id_instance]
}
