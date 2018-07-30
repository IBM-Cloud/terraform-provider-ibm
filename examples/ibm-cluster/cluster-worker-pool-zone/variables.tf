variable "org" {}

variable "space" {}

variable "region" {}

variable "datacenter" {}

variable "machine_type" {}

variable "isolation" {}

variable "private_vlan_id" {}

variable "public_vlan_id" {}

variable "subnet_id" {}

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

variable "cluster_name" {
  default = "cluster"
}

variable "worker_pool_name" {
  default = "worker_pool"
}

variable "zone" {}

variable "zone_private_vlan_id" {}

variable "zone_public_vlan_id" {}
