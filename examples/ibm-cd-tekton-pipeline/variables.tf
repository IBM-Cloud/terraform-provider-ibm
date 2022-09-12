variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for cd_tekton_pipeline_definition
variable "cd_tekton_pipeline_definition_pipeline_id" {
  description = "The Tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}

// Resource arguments for cd_tekton_pipeline_trigger_property
variable "cd_tekton_pipeline_trigger_property_pipeline_id" {
  description = "The Tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "cd_tekton_pipeline_trigger_property_trigger_id" {
  description = "The trigger ID."
  type        = string
  default     = "1bb892a1-2e04-4768-a369-b1159eace147"
}
variable "cd_tekton_pipeline_trigger_property_name" {
  description = "Property name."
  type        = string
  default     = "key1"
}
variable "cd_tekton_pipeline_trigger_property_value" {
  description = "Property value."
  type        = string
  default     = "https://github.com/IBM/tekton-tutorial.git"
}
variable "cd_tekton_pipeline_trigger_property_enum" {
  description = "Options for `single_select` property type. Only needed for `single_select` property type."
  type        = list(string)
  default     = [ "enum" ]
}
variable "cd_tekton_pipeline_trigger_property_type" {
  description = "Property type."
  type        = string
  default     = "text"
}
variable "cd_tekton_pipeline_trigger_property_path" {
  description = "A dot notation path for `integration` type properties to select a value from the tool integration. If left blank the full tool integration JSON will be selected."
  type        = string
  default     = "path"
}

// Resource arguments for cd_tekton_pipeline_property
variable "cd_tekton_pipeline_property_pipeline_id" {
  description = "The Tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "cd_tekton_pipeline_property_name" {
  description = "Property name."
  type        = string
  default     = "key1"
}
variable "cd_tekton_pipeline_property_value" {
  description = "Property value."
  type        = string
  default     = "https://github.com/IBM/tekton-tutorial.git"
}
variable "cd_tekton_pipeline_property_enum" {
  description = "Options for `single_select` property type. Only needed when using `single_select` property type."
  type        = list(string)
  default     = [ "enum" ]
}
variable "cd_tekton_pipeline_property_type" {
  description = "Property type."
  type        = string
  default     = "text"
}
variable "cd_tekton_pipeline_property_path" {
  description = "A dot notation path for `integration` type properties to select a value from the tool integration."
  type        = string
  default     = "path"
}

// Resource arguments for cd_tekton_pipeline_trigger
variable "cd_tekton_pipeline_trigger_pipeline_id" {
  description = "The Tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "cd_tekton_pipeline_trigger_type" {
  description = "Trigger type."
  type        = string
  default     = "manual"
}
variable "cd_tekton_pipeline_trigger_name" {
  description = "Trigger name."
  type        = string
  default     = "Manual Trigger"
}
variable "cd_tekton_pipeline_trigger_event_listener" {
  description = "Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline."
  type        = string
  default     = "pr-listener"
}
variable "cd_tekton_pipeline_trigger_tags" {
  description = "Trigger tags array."
  type        = list(string)
  default     = [ "tags" ]
}
variable "cd_tekton_pipeline_trigger_max_concurrent_runs" {
  description = "Defines the maximum number of concurrent runs for this trigger. Omit this property to disable the concurrency limit."
  type        = number
  default     = 3
}
variable "cd_tekton_pipeline_trigger_disabled" {
  description = "Flag whether the trigger is disabled. If omitted the trigger is enabled by default."
  type        = bool
  default     = false
}
variable "cd_tekton_pipeline_trigger_cron" {
  description = "Only needed for timer triggers. Cron expression for timer trigger."
  type        = string
  default     = "cron"
}
variable "cd_tekton_pipeline_trigger_timezone" {
  description = "Only needed for timer triggers. Timezone for timer trigger."
  type        = string
  default     = "timezone"
}

// Resource arguments for cd_tekton_pipeline
variable "cd_tekton_pipeline_enable_slack_notifications" {
  description = "Flag whether to enable slack notifications for this pipeline. When enabled, pipeline run events will be published on all slack integration specified channels in the enclosing toolchain."
  type        = bool
  default     = true
}
variable "cd_tekton_pipeline_enable_partial_cloning" {
  description = "Flag whether to enable partial cloning for this pipeline. When partial clone is enabled, only the files contained within the paths specified in definition repositories will be read and cloned. This means symbolic links may not work."
  type        = bool
  default     = true
}

// Data source arguments for cd_tekton_pipeline_definition
variable "cd_tekton_pipeline_definition_pipeline_id" {
  description = "The Tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "cd_tekton_pipeline_definition_definition_id" {
  description = "The definition ID."
  type        = string
  default     = "94299034-d45f-4e9a-8ed5-6bd5c7bb7ada"
}

// Data source arguments for cd_tekton_pipeline_trigger_property
variable "cd_tekton_pipeline_trigger_property_pipeline_id" {
  description = "The Tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "cd_tekton_pipeline_trigger_property_trigger_id" {
  description = "The trigger ID."
  type        = string
  default     = "1bb892a1-2e04-4768-a369-b1159eace147"
}
variable "cd_tekton_pipeline_trigger_property_property_name" {
  description = "The property name."
  type        = string
  default     = "debug-pipeline"
}

// Data source arguments for cd_tekton_pipeline_property
variable "cd_tekton_pipeline_property_pipeline_id" {
  description = "The Tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "cd_tekton_pipeline_property_property_name" {
  description = "The property name."
  type        = string
  default     = "debug-pipeline"
}

// Data source arguments for cd_tekton_pipeline_trigger
variable "cd_tekton_pipeline_trigger_pipeline_id" {
  description = "The Tekton pipeline ID."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
variable "cd_tekton_pipeline_trigger_trigger_id" {
  description = "The trigger ID."
  type        = string
  default     = "1bb892a1-2e04-4768-a369-b1159eace147"
}

// Data source arguments for cd_tekton_pipeline
variable "cd_tekton_pipeline_id" {
  description = "ID of current instance."
  type        = string
  default     = "94619026-912b-4d92-8f51-6c74f0692d90"
}
