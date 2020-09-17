variable "vmname" {
}

variable "ssh_key_id" {}

variable "vmcount" {
  default = "3"
}

variable "domain" {
  default = "examplehost.com"
}

variable "os" {
  default = "REDHAT_7_64"
}

variable "datacenter" {
  default = "dal10"
}

variable "network_speed" {
  default = "10"
}

variable "cores" {
  default = "4"
}

variable "memory" {
  default = "16384"
}

variable "disks" {
  default = "25"
}

variable "iaas_classic_api_key" {

}

variable "iaas_classic_username" {
}

variable "ssh_label" {
  default = "test_ssh"
}

variable "notes" {
  default = "test_ssh_key_notes"
}
variable "public_key" {

}
variable "module_depends_on" {
}