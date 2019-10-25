variable "org" {}

variable "space"{}

variable "subnet_id" {}

variable "vpc_id" {}

variable "flavor" {}

variable "worker_count" {}

variable "zone_name" {}

variable "cluster_name" {
  default = "cluster"
}

variable "service_instance_name" {
  default = "myservice"
}

variable "service_key" {
  default = "myservicekey"
}

variable "service_offering" {
  default = "speech_to_text"
}

variable "plan" {
  default = "lite"
}
