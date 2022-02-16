variable "cluster" {
  description = "Satellite Location Name"
  type        = string
}

variable "location" {
  description = "Satellite Location Name"
  type        = string
}

variable "kube_version" {
  description = "Satellite Kube Version"
}

variable "resource_group" {
  description = "Resource Group Name that has to be targeted"
  type        = string
}

variable "zones" {
  type    = list(string)
  default = ["us-east-1", "us-east-2", "us-east-3"]
}

variable "worker_pool_name" {
  description = "Worker Pool Name"
  type        = string
}

variable "worker_count" {
  description = "Worker Count for default pool"
  type        = number
  default     = 1
}

variable "default_wp_labels" {
  description = "Label to add default worker pool"
  type        = map(any)

  default = {
    "poolname" = "default-worker-pool"
  }
}

variable "workerpool_labels" {
  description = "Label to add to workerpool"
  type        = map(any)

  default = {
    "poolname" = "worker-pool"
  }
}

variable "host_labels" {
  description = "Label to add to attach host script"
  type        = list(string)
}

variable "cluster_tags" {
  description = "List of tags associated with this resource."
  type        = list(string)
  default     = ["env:cluster"]
}

variable "zone_name" {
  description = "zone name"
  type        = string
}