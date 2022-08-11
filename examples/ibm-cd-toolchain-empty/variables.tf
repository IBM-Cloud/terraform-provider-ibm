variable "resource_group" {
  type        = string
  description = "Resource group within which toolchain will be created"
  default     = "Default"
}

variable "region" {
  type        = string
  description = "IBM Cloud region where your toolchain will be created"
  default     = "us-south"
}

variable "ibmcloud_api_key" {
  type        = string
  description = "IBM Cloud API KEY to interact with IBM Cloud"
}

variable "toolchain_name" {
  type        = string
  description = "Name of the Toolchain."
  default     = "Sample Empty Toolchain"
}

variable "toolchain_description" {
  type        = string
  description = "Description for the Toolchain."
  default     = "This toolchain has no preconfigured tools. If you are already familiar with toolchains, you can set up your own toolchain."
}

