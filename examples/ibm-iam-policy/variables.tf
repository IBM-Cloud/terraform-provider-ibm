variable "org" {}

variable "space" {}

variable "region" {}

variable "iam_region" {}

variable "ibm_id1" {}

variable "datacenter" {}

variable "machine_type" {}

variable "isolation" {}

variable "private_vlan_id" {}

variable "public_vlan_id" {}

variable "cluster_name" {
  default = "cluster"
}

variable "service_name" {
  default = "IBM Bluemix Container Service"
}
