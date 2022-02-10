variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for tekton_pipeline_definition
variable "tekton_pipeline_definition_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}

// Resource arguments for tekton_pipeline_trigger_property
variable "tekton_pipeline_trigger_property_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "tekton_pipeline_trigger_property_trigger_id" {
  description = "The trigger ID."
  type        = string
  default     = "1bb892a1-2e04-4768-a369-b1159eace147"
}

// Resource arguments for tekton_pipeline_property
variable "tekton_pipeline_property_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}

// Resource arguments for tekton_pipeline_trigger
variable "tekton_pipeline_trigger_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}

// Resource arguments for tekton_pipeline
variable "tekton_pipeline_integration_instance_id" {
  description = "UUID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}

// Data source arguments for tekton_pipeline_definition
variable "tekton_pipeline_definition_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "tekton_pipeline_definition_definition_id" {
  description = "The definition ID."
  type        = string
  default     = "94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
}

// Data source arguments for tekton_pipeline_trigger_property
variable "tekton_pipeline_trigger_property_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "tekton_pipeline_trigger_property_trigger_id" {
  description = "The trigger ID."
  type        = string
  default     = "1bb892a1-2e04-4768-a369-b1159eace147"
}
variable "tekton_pipeline_trigger_property_property_name" {
  description = "The property's name."
  type        = string
  default     = "debug-pipeline"
}

// Data source arguments for tekton_pipeline_property
variable "tekton_pipeline_property_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "tekton_pipeline_property_property_name" {
  description = "The property's name."
  type        = string
  default     = "debug-pipeline"
}

// Data source arguments for tekton_pipeline_workers
variable "tekton_pipeline_workers_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
