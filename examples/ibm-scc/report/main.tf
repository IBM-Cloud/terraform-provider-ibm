data "ibm_scc_latest_reports" "scc_latest_reports_instance" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	sort = "profile_name"
}

data "ibm_scc_report_rule" "scc_report_rule_instance" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	report_id = var.scc_report_id
	rule_id   = var.scc_rule_id
}

data "ibm_scc_report_tags" "scc_report_tags_instance" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	report_id = var.scc_report_id
}

data "ibm_scc_report_evaluations" "scc_report_evaluations_instance" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	report_id = var.scc_report_id
}

data "ibm_scc_report_controls" "scc_report_controls_instance" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	report_id = var.scc_report_id
}

data "ibm_scc_report_summary" "scc_report_summary_instance" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	report_id = var.scc_report_id
}

data "ibm_scc_report_violation_drift" "scc_report_violation_drift_instance" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	report_id = var.scc_report_id
}

data "ibm_scc_report" "scc_report_instance" {
	instance_id = "00000000-1111-2222-3333-444444444444"
	report_id = var.scc_report_id
}
