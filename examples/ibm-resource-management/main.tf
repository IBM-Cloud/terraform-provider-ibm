provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision resource_alias resource instance
resource "ibm_resource_alias" "resource_alias_instance" {
  name = var.resource_alias_name
  source = var.resource_alias_source
  target = var.resource_alias_target
}

// Provision resource_binding resource instance
resource "ibm_resource_binding" "resource_binding_instance" {
  source = var.resource_binding_source
  target = var.resource_binding_target
  name = var.resource_binding_name
  parameters = var.resource_binding_parameters
  role = var.resource_binding_role
}

// Create resource_aliases data source
data "ibm_resource_aliases" "resource_aliases_instance" {
  name = var.resource_aliases_name
}

// Create resource_bindings data source
data "ibm_resource_bindings" "resource_bindings_instance" {
  name = var.resource_bindings_name
}
