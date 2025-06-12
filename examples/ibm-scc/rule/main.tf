// Provision scc_rule resource instance
resource "ibm_scc_rule" "scc_rule_tf_demo" {
	instance_id = "00000000-1111-2222-3333-444444444444"
    description = var.scc_description
    target {
        service_name = "cloud-object-storage"
        resource_kind = "bucket"
    }
    labels = ["SOC2"]
    required_config {
        description = "this is a terraform update to description"
        // This is Terraform HCL, not JSON
		property = "storage_class"
		operator = "string_equals"
		value    = "smart"
    } 
}
