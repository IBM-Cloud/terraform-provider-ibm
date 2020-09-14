output "WorkSpaceValues" {
  value = data.ibm_schematics_workspace.test.template_id.0
}

output "StateStoreValues" {
	value = data.ibm_schematics_state.test.state_store
}

output "OutputValues" {
	value = data.ibm_schematics_output.test.output_values
} 
