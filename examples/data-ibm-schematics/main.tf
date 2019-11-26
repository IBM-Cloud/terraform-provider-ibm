data "ibm_schematics_workspace" "test" {
  workspace_id = "${var.workspaceID}"
}

data "ibm_schematics_state" "test" {
	workspace_id = "${var.workspaceID}"
	template_id = "${var.templateID}"
}
	  
data "ibm_schematics_output" "test" {
	workspace_id = "${var.workspaceID}"
	template_id = "${var.templateID}"
}
	  
	
