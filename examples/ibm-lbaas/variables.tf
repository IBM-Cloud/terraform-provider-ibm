variable "public_key" {
  description = "public SSH key to use in keypair"
}

variable "ssh_label" {
  default = "ssh_lbaas"
}

variable "name" {
  default = "terraformLbaasExample"
}

variable "notes" {
  default = "for lbaas test"
}

variable "osref" {
  default = "UBUNTU_16_64"
}

variable "domain" {
  default = "ibm.com"
}

variable "lb_method" {
  default = "round_robin"
}

variable "subnet_id" {
  default = "1851345"
}

variable "datacenter" {
  default = "dal10"
}

variable "vm-post-install-script-uri" {
  default = "https://raw.githubusercontent.com/hkantare/test/master/nginx.sh"
}

variable "hostname" {
  default = "lbaas-example"
}

