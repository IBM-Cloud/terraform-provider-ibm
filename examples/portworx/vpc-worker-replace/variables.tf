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


#####################################################
# If the worker list is being provided as inputs, 
# the list should be user generated and 
# should not be passed from the `ibm_container_cluster` data source.
#
# The order of the list should not be changed until
# all the workers in the list are replaced. 
#
# This is required to avoid diffs of order changes.
#####################################################
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
