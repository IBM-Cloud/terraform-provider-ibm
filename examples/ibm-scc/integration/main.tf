resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
  provider_type_id = var.scc_provider_type_id
  name = var.scc_provider_type_instance_instance
  attributes = var.scc_provider_type_instance_attributes
}

