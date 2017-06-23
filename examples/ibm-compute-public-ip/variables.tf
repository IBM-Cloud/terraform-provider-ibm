variable "datacenter" {
  description = "The datacenter"
  default = "dal01"
}

variable "ssh_label" {
  default = "Personal"
}

variable "ssh_key_path" {
  default = "~/.ssh/id2_rsa.pub"
}
