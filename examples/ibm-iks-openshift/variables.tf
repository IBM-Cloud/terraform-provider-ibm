
variable "datacenter" {
  default = "wdc04"
}

variable "machine_type" {
  default = "b3c.4x16"
}

variable "hardware" {
  default = "shared"
}

variable "private_vlan_id" {
  default = "2640745"
}

variable "public_vlan_id" {
  default = "2640747"
}

variable "cluster_name" {
  default = "terraform_iks_openshift"
}

variable "kube_version" {
  default = "3.11_openshift"
}

