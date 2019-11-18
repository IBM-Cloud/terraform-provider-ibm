variable "org" {}

variable "space" {}

variable "datacenter" {}

variable "machine_type" {}

variable "hardware" {}

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
  default = "mycluster"
}
