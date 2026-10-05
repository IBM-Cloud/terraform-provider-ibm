data "ibm_resource_group" "group" {
  name = var.resource_group
}

// Create a Red Hat AI Inference project
resource "ibm_rhaii_project" "project" {
  name              = var.project_name
  location          = var.location
  resource_group_id = data.ibm_resource_group.group.id
  tags              = var.tags
}

// Read the project back by name
data "ibm_rhaii_project" "project" {
  name              = ibm_rhaii_project.project.name
  resource_group_id = ibm_rhaii_project.project.resource_group_id
}
