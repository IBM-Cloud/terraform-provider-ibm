variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
  sensitive   = true
}

variable "resource_group" {
  description = "Name of the resource group for the project"
  type        = string
  default     = "Default"
}

variable "project_name" {
  description = "Name of the Red Hat AI Inference project"
  type        = string
  default     = "my-rhaii-project"
}

variable "location" {
  description = "Region of the Red Hat AI Inference project"
  type        = string
  default     = "us-east"
}

variable "tags" {
  description = "Tags for the project"
  type        = list(string)
  default     = []
}

variable "access_tags" {
  description = "Access management tags for the project, in key:value format. The tags must already exist in the account (create them with ibm_iam_access_tag)."
  type        = list(string)
  default     = []
}

variable "model" {
  description = "ID of the inference model to read"
  type        = string
  default     = "llama-3-3-70b-instruct"
}
