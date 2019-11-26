output "WorkSpace Values" {
  value = "${data.ibm_schematics_workspace.test.template_id.0}"
}

output "StateStore Values" {
	value = "${data.ibm_schematics_state.test.state_store}"
}

output "Output Values" {
	value = "${data.ibm_schematics_output.test.output_values}"
}