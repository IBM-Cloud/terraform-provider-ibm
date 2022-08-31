variable "resource_group" {
  type        = string
  description = "Resource group within which toolchain will be created"
  default     = "Default"
}

variable "ibm_cloud_api_key" {
  type        = string
  description = "IBM Cloud API KEY to fetch cloud resources"
}

variable "ibm_cloud_api" {
  type        = string
  description = "IBM Cloud API Endpoint"
  default     = "https://cloud.ibm.com"
}

variable "region" {
  type        = string
  description = "IBM Cloud region where your toolchain will be created"
  default     = "us-south"
}

variable "toolchain_name" {
  type        = string
  description = "Name of the Toolchain."
  default     = "Simple Helm Toolchain"
}

variable "toolchain_description" {
  type        = string
  description = "Description for the Toolchain."
  default     = "Toolchain created using IBM Cloud Continuous Delivery Service"
}

variable "app_name" {
  type        = string
  description = "Name of the application."
  default     = "simple-helm-app"
}

variable "app_image_name" {
  type        = string
  description = "Name of the application image."
  default     = "simple-helm-app"
}

variable "cluster_name" {
  type        = string
  description = "Name of the kubernetes cluster where the application will be deployed."
  default     = "mycluster-free"
}

variable "cluster_namespace" {
  type        = string
  description = "Name of the kubernetes cluster where the application will be deployed."
  default     = "prod"
}

variable "cluster_region" {
  type        = string
  description = "Region of the kubernetes cluster where the application will be deployed."
  default     = "ibm:yp1:us-south"
}

variable "registry_namespace" {
  type        = string
  description = "Namespace within the IBM Cloud Container Registry where application image need to be stored."
  default     = "myregistry-free"
}

variable "registry_region" {
  type        = string
  description = "IBM Cloud Region where the IBM Cloud Container Registry where registry is to be created."
  default     = "ibm:yp1:us-south"
}

variable "kp_name" {
  type        = string
  description = "Name of the Key Protect Instance to store the secrets."
  default     = "Key Protect Service"
}

variable "kp_region" {
  type        = string
  description = "IBM Cloud Region where the Key Protect Instance is created."
  default     = "ibm:yp1:us-south"
}

variable "app_repo" {
  type        = string
  description = "Repository url for the repository containing application source code."
  default     = "https://us-south.git.cloud.ibm.com/open-toolchain/hello-helm.git"
}

variable "pipeline_repo" {
  type        = string
  description = "Repository url for the repository containing pipeline source code."
  default     = "https://us-south.git.cloud.ibm.com/open-toolchain/simple-helm-toolchain.git"
}

variable "tekton_tasks_catalog_repo" {
  type        = string
  description = "Repository url for the repository containing commonly used tekton tasks."
  default     = "https://us-south.git.cloud.ibm.com/open-toolchain/tekton-catalog.git"
}