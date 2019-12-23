variable "ssh_label" {
  default = "ssh_key_vlan"
}

variable "ssh_public_key" {
}

variable "vm_hostname" {
  default = "vlan-provider"
}
variable "vm_os_reference_code" {
  default = "CENTOS_7_64"
}
variable "datacenter" {
  default = "dal06"
}
variable "vm_domain" {
  default = "terraformuat.ibm.com"
}
variable "vm_network_speed" {
  default = 10
}
variable "vm_cores" {
  default = 1
}
variable "vm_memory" {
  default = 1024
}
variable "vm_disks" {
  default = [25, 10, 20]
}

variable "vlan_name_public" {
  default = "test_vlan_public"
}

variable "vlan_name_private" {
  default = "test_vlan_private"
}
