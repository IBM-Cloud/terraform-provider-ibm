variable "pipeline_id" {
}

variable "resource_group" {
}

variable "app_name" {
}

variable "app_image_name" {
}

variable "cluster_name" {
}

variable "cluster_namespace" {
}

variable "cluster_region" {
}

variable "registry_namespace" {
}

variable "registry_region" {
}

variable "region" {
}

variable "ibm_cloud_api_key" {
}

variable "kp_integration_name" {
}

variable "app_repo" {
}

variable "pipeline_repo" {
}

variable "tekton_tasks_catalog_repo" {
}

variable "ci_pipeline_manual_trigger_name" {
  type        = string
  description = "The name of Manual Trigger for CI Pipeline as defined in tekton definition."
  default     = "Manual Trigger"
}

variable "ci_pipeline_manual_trigger_type" {
  type        = string
  description = "The type of Manual Trigger for CI Pipeline as defined in tekton definition."
  default     = "manual"
}

variable "ci_pipeline_manual_trigger_disabled" {
  type        = bool
  description = "Flag to disable manual CI Trigger"
  default     = false
}

variable "ci_pipeline_manual_trigger_listener_name" {
  type        = string
  description = "The name of EventListener for the CI Pipeline SCM Trigger as defined in tekton definition."
  default     = "manual-run"
}

variable "ci_pipeline_scm_trigger_name" {
  type        = string
  description = "The name of SCM Trigger for CI Pipeline as defined in tekton definition."
  default     = "SCM Trigger"
}

variable "ci_pipeline_scm_trigger_type" {
  type        = string
  description = "The type of SCM Trigger for CI Pipeline as defined in tekton definition."
  default     = "scm"
}

variable "ci_pipeline_scm_trigger_disabled" {
  type        = bool
  description = "Flag to disable SCM CI Trigger"
  default     = false
}

variable "ci_pipeline_scm_trigger_listener_name" {
  type        = string
  description = "The name of EventListener for the CI Pipeline SCM Trigger as defined in tekton definition."
  default     = "grit-or-gitlab-commit"
}

variable "ci_pipeline_max_concurrent_runs" {
  type        = number
  description = "The number of maximum concurrent runs to be supported by CI Pipeline"
  default     = 1
}

