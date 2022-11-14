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
  default     = "Simple Helm Toolchain"
}

variable "toolchain_description" {
  type        = string
  description = "Description for the Toolchain"
  default     = "Toolchain created using IBM Cloud Continuous Delivery Service"
}

variable "cluster" {
  type        = string
  description = "The name of your IKS cluster where you will be deploying the sample app"
}

variable "cluster_namespace" {
  type        = string
  description = "The namespace in your cluster where the app will be deployed"
  default     = "prod"
}

variable "registry_namespace" {
  type        = string
  description = "The IBM Cloud Container Registry namespace where the app image will be built and stored."
}
