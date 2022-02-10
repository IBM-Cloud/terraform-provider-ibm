provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision tekton_pipeline_property resource instance
resource "ibm_tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  name = var.tekton_pipeline_property_name
  value = var.tekton_pipeline_property_value
  options = var.tekton_pipeline_property_options
  type = var.tekton_pipeline_property_type
  path = var.tekton_pipeline_property_path
}

// Create tekton_pipeline_property data source
data "ibm_tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  property_name = var.tekton_pipeline_property_property_name
}

// Create tekton_pipeline_workers data source
data "ibm_tekton_pipeline_workers" "tekton_pipeline_workers_instance" {
  pipeline_id = var.tekton_pipeline_workers_pipeline_id
}
