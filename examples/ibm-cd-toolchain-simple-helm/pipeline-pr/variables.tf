variable "pipeline_id" {
}

variable "resource_group" {
}

variable "app_name" {
}

variable "region" {
}

variable "ibmcloud_api_key" {
}

variable "kp_integration_name" {
}

variable "app_repo" {
}

variable "app_repo_branch" {
}

variable "ibmcloud_api" {
}

variable "pipeline_repo" {
    type        = string
    description = "The repository url containing pipeline definitions for Simple Helm Toolchain."
}

variable "pipeline_repo_branch" {
}

variable "pipeline_path" {
  type        = string
  description = "The relative folder path within pipeline definitions repository containing tekton definitions for pipelines."
  default     = ".pr-pipeline"
}

variable "tekton_tasks_catalog_repo" {
}

variable "definitions_branch" {
}

variable "pr_pipeline_scm_trigger_type" {
  type        = string
  description = "The type of SCM Trigger for PR Pipeline as defined in tekton definition."
  default     = "scm"
}

variable "pr_pipeline_scm_trigger_name" {
  type        = string
  description = "The name of SCM Trigger for PR Pipeline as defined in tekton definition."
  default     = "Git PR Trigger"
}

variable "pr_pipeline_scm_trigger_listener_name" {
  type        = string
  description = "The name of EventListener for the PR Pipeline SCM Trigger as defined in tekton definition."
  default     = "gitlab-pr-listener"
}

variable "pr_pipeline_scm_trigger_enabled" {
  type        = bool
  description = "Flag to enable or disable SCM CI Trigger"
  default     = true
}

variable "pr_pipeline_max_concurrent_runs" {
  type        = number
  description = "The number of maximum concurrent runs to be supported by PR Pipeline"
  default     = 1
}