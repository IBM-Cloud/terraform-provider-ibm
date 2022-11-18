variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "resource_group" {
  type        = string
  description = "Resource group within which toolchain will be created"
  default     = "Default"
}

variable "clone_repo" {
  type        = string
  description = "URL of the tekton repo to clone"
  default     = "https://github.com/open-toolchain/hello-tekton"
}

variable "repo_name" {
  type        = string
  description = "Name of the new repo that will be created in the toolchain"
  default     = "simple-tekton"
}

variable "region" {
  type        = string
  description = "IBM Cloud region where your toolchain will be created"
  default     = "us-south"
}

variable "toolchain_name" {
  type        = string
  description = "Name of the Toolchain"
  default     = "Simple Terraform Toolchain"
}

variable "toolchain_description" {
  type        = string
  description = "Description for the Toolchain"
  default     = "Toolchain created using IBM Cloud Continuous Delivery Service"
}
