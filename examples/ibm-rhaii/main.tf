data "ibm_resource_group" "group" {
  name = var.resource_group
}

// Create a Red Hat AI Inference project
resource "ibm_rhaii_project" "project" {
  name              = var.project_name
  location          = var.location
  resource_group_id = data.ibm_resource_group.group.id
  tags              = var.tags
  access_tags       = var.access_tags
}

// Read the project back by name
data "ibm_rhaii_project" "project" {
  name              = ibm_rhaii_project.project.name
  resource_group_id = ibm_rhaii_project.project.resource_group_id
}

// List the inference models available in the project
data "ibm_rhaii_inference_models" "models" {
  project_id = ibm_rhaii_project.project.project_id
  location   = ibm_rhaii_project.project.location
}

// Read the details of one model
data "ibm_rhaii_inference_model" "model" {
  project_id = ibm_rhaii_project.project.project_id
  location   = ibm_rhaii_project.project.location
  model      = var.model
}
