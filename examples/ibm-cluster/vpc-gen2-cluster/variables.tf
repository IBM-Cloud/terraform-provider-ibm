variable "flavor" {
  default = "bx1.2x8"
}

variable "kube_version" {
  default = "1.17.4"
}

variable "worker_count" {
  default = "1"
}

variable "region" {
  default = "us-south"
}

variable "resource_group" {
  default = "Default"
}

variable "cluster_name" {
  default = "cluster"
}

variable "worker_pool_name" {
  default = "workerpool"
}

variable "service_instance_name" {
  default = "myservice"
}

variable "service_offering" {
  default = "speech_to_text"
}

variable "plan" {
  default = "lite"
}

