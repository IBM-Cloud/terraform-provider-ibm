variable "volume_name" {
  default = "testvolume"
}

variable "volume_profile" {
  default = "custom"
}

variable "flavor" {
  default = "bx2.2x8"
}

variable "worker_count" {
  default = "1"
}

variable "resource_group" {
  default = "Default"
}

variable "name" {
  default = "cluster"
}

variable "region" {
  default = "us-south"
}

variable "kube_version" {
  default = "1.17.7"
}