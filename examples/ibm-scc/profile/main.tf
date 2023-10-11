data "ibm_scc_control_library" "scc_control_library" {
	control_library_id = var.ibm_scc_control_library_id
}

resource "ibm_scc_profile" "scc_demo_profile" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	profile_type = "custom"
	profile_description = var.ibm_scc_profile_description
	profile_name = var.ibm_scc_profile_name
	default_parameters {
	}
	controls {
		control_library_id = var.ibm_scc_control_library_id
		control_id = "032a81ca-6ef7-4ac2-81ac-20ee4a780e3b"
	}
}

resource "ibm_scc_profile_attachment" "scc_demo_profile_attachment" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	profile_id = resource.ibm_scc_profile.scc_demo_profile.id
	name = var.ibm_scc_profile_attachment_name
	description = var.ibm_scc_profile_attachment_description
	scope {
		environment = "ibm-cloud"	
		properties {
			name = "scope_id"
			value = "62ecf99b240144dea9125666249edfcb"
		}
		properties {
			name = "scope_type"
			value = "account"
		}
	}
	schedule = "every_30_days"
	status = "enabled"
	notifications {
		enabled = false
		controls {
			failed_control_ids = []
			threshold_limit = 14
		}
	}
}
