variable "flavor" {
  default =  "bx2.2x8"
}

variable "worker_count" {
  default = "1"
}

variable "kube_version" {
  default = "1.17.5"
}

variable "region" {
  default = "us-south"
}

variable "resource_group" {
  default = "default"
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
