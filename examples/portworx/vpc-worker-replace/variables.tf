######################################################
#IBM-Cloud Authentication Credentials
######################################################

variable "ibmcloud_api_key" {
  type        = string
  description = "IBM-Cloud API Key"
}

#####################################################
# Vpc Kubernetes cluster
# Copyright 2020 IBM
#####################################################
variable "cluster_name" {
  description = "Name of the VPC cluster"
  type        = string
}

variable "worker_list" {
    description = "List of workers to process"
    type        = list(string)
}

variable "resource_group" {
  description = "Name of resource group."
  type        = string
  default     = null
}

variable "create_timeout" {
  type        = string
  description = "Timeout duration for create."
  default     = null
}

variable "delete_timeout" {
  type        = string
  description = "Timeout duration for delete."
  default     = null
}

variable "ptx_timeout" {
  type        = string
  description = "Timeout duration for checking ptx pods/status."
  default     = null
}

variable "kube_config_path" {
    description = "Path of downloaded cluster config"
    type        = string
    default     = ""
}

variable "check_ptx_status" {
    description = "Check status of portworx on replaced workers"
    type        = bool
    default     = true
}
