// This allows schematics_workspace data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_workspace" {
  value       = ibm_schematics_workspace.schematics_workspace_instance
  description = "schematics_workspace resource instance"
}
// This allows schematics_action data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_action" {
  value       = ibm_schematics_action.schematics_action_instance
  description = "schematics_action resource instance"
}
// This allows schematics_job data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_job" {
  value       = ibm_schematics_job.schematics_job_instance
  description = "schematics_job resource instance"
}

// This allows schematics_policy data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_policy" {
  value       = ibm_schematics_policy.schematics_policy_instance
  description = "schematics_policy resource instance"
}
// This allows schematics_agent data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_agent" {
  value       = ibm_schematics_agent.schematics_agent_instance
  description = "schematics_agent resource instance"
}
// This allows schematics_agent_prs data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_agent_prs" {
  value       = ibm_schematics_agent_prs.schematics_agent_prs_instance
  description = "schematics_agent_prs resource instance"
}
// This allows schematics_agent_deploy data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_agent_deploy" {
  value       = ibm_schematics_agent_deploy.schematics_agent_deploy_instance
  description = "schematics_agent_deploy resource instance"
}
// This allows schematics_agent_health data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_agent_health" {
  value       = ibm_schematics_agent_health.schematics_agent_health_instance
  description = "schematics_agent_health resource instance"
}

