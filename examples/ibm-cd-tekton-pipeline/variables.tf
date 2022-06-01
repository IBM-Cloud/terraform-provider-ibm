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
variable "tekton_pipeline_trigger_property_name" {
  description = "Property name."
  type        = string
  default     = "key1"
}
variable "tekton_pipeline_trigger_property_value" {
  description = "String format property value."
  type        = string
  default     = "https://github.com/IBM/tekton-tutorial.git"
}
variable "tekton_pipeline_trigger_property_enum" {
  description = "Options for SINGLE_SELECT property type."
  type        = list(string)
  default     = [ "enum" ]
}
variable "tekton_pipeline_trigger_property_default" {
  description = "Default option for SINGLE_SELECT property type."
  type        = string
  default     = "default"
}
variable "tekton_pipeline_trigger_property_type" {
  description = "Property type."
  type        = string
  default     = "TEXT"
}
variable "tekton_pipeline_trigger_property_path" {
  description = "property path for INTEGRATION type properties."
  type        = string
  default     = "path"
}

// Resource arguments for tekton_pipeline_property
variable "tekton_pipeline_property_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "tekton_pipeline_property_name" {
  description = "Property name."
  type        = string
  default     = "key1"
}
variable "tekton_pipeline_property_value" {
  description = "String format property value."
  type        = string
  default     = "https://github.com/IBM/tekton-tutorial.git"
}
variable "tekton_pipeline_property_enum" {
  description = "Options for SINGLE_SELECT property type."
  type        = list(string)
  default     = [ "enum" ]
}
variable "tekton_pipeline_property_default" {
  description = "Default option for SINGLE_SELECT property type."
  type        = string
  default     = "default"
}
variable "tekton_pipeline_property_type" {
  description = "Property type."
  type        = string
  default     = "TEXT"
}
variable "tekton_pipeline_property_path" {
  description = "property path for INTEGRATION type properties."
  type        = string
  default     = "path"
}

// Resource arguments for tekton_pipeline_trigger
variable "tekton_pipeline_trigger_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}

// Resource arguments for tekton_pipeline

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

// Data source arguments for tekton_pipeline_trigger
variable "tekton_pipeline_trigger_pipeline_id" {
  description = "The tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "tekton_pipeline_trigger_trigger_id" {
  description = "The trigger ID."
  type        = string
  default     = "1bb892a1-2e04-4768-a369-b1159eace147"
}

// Data source arguments for tekton_pipeline
variable "tekton_pipeline_id" {
  description = "ID of current instance."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
