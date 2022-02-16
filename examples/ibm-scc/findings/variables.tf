variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "account_id" {
  description = "IBM Cloud Account ID"
  type        = string
}

variable "provider_id" {
  description = "Part of parent. This field contains the provider_id for example: providers/{provider_id}"
  type        = string
}

variable "note_id" {
  description = "Second part of note name: providers/{provider_id}/notes/{note_id}"
  type        = string
}

variable "occurrence_id" {
  description = "Second part of occurrence name: providers/{provider_id}/occurrences/{occurrence_id}"
  type        = string
}

