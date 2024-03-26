resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
  instance_id = "00000000-1111-2222-3333-444444444444"
  provider_type_id = var.scc_provider_type_id
  name = var.scc_provider_type_instance_instance
  attributes = var.scc_provider_type_instance_attributes
}

