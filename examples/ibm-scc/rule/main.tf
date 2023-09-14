// Provision scc_rule resource instance
resource "ibm_scc_rule" "scc_rule_tf_demo" {
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
