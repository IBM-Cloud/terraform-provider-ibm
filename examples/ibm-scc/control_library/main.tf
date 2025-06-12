resource "ibm_scc_control_library" "scc_demo_control_library" {
  instance_id = "00000000-1111-2222-3333-444444444444"
  control_library_name = var.scc_control_library_name 
  control_library_description = var.scc_control_library_description
  control_library_type = "custom"
  control_library_version = var.scc_control_version
  version_group_label = "d755830f-1d83-4fab-b5d5-1dfb2b0dad1f"
  latest = true
  controls {
	control_id = "032a81ca-6ef7-4ac2-81ac-20ee4a780e3b"
	control_name = var.scc
	control_description = "Boundary Protection"
	control_category = "System and Communications Protection"
	control_requirement = true
	status = "enabled"
	control_docs {}
	control_specifications {
		control_specification_id = "5c7d6f88-a92f-4734-9b49-bd22b0900184"
		control_specification_description = "IBM Cloud"
		component_id = "iam-identity"
		component_name = "IAM Identity Service"
		environment = "ibm-cloud"
		assessments {
			assessment_type = "automated"
			assessment_method = "ibm-cloud-rule"
			assessment_id = "rule-a637949b-7e51-46c4-afd4-b96619001bf1"
			assessment_description = "All assessments related to iam_identity"
			parameters {
				parameter_name = "session_invalidation_in_seconds"
				parameter_display_name = "Sign out due to inactivity in seconds"
				parameter_type = "numeric"
			}
		}
	    responsibility = "user"
	}
  }
}  

