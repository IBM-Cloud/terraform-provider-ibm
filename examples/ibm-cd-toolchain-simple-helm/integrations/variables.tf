variable "toolchain_id" {
}

variable "resource_group" {
}

variable "region" {
}

variable "key_protect_integration_name" {
  type        = string
  description = "Name of the Key Protect Toolchain Integration"
  default     = "Key Protect Integration Instance"
}

variable "key_protect_instance_name" {
  type        = string
  description = "Name of the Key Protect Toolchain Service Instance in IBM Cloud"
}

variable "key_protect_instance_region" {
  type        = string
  description = "Region of the Key Protect Toolchain Service Instance in IBM Cloud"
}

variable "key_protect_instance_guid" {
  type        = string
  description = "GUID of the Key Protect Toolchain Service Instance in IBM Cloud"
}

variable "key_protect_service_auth" {
  type        = string
  description = "Authorization Permission for the Key Protect Toolchain Service Instance in IBM Cloud"
  default     = "[\"Viewer\", \"ReaderPlus\"]"
}