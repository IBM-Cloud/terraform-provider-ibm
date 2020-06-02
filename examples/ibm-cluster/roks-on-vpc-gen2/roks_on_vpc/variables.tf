variable "flavor" {
  default = "bx2.16x64"
}

variable "kube_version" {
  default = "4.3_openshift"
}

variable "worker_count" {
  default = "2"
}

variable "region" {
  default = "us-south"
}

variable "resource_group" {
  default = "Default"
}

variable "cluster_name" {
  default = "cluster-roks-on-vpc"
}

variable "worker_pool_name" {
  default = "workerpool"
}
