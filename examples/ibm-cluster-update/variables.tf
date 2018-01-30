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
  default = "cloudantNoSQLDB"
}

variable "plan" {
  default = "Lite"
}

variable "cluster_name" {
  default = "cluster"
}
