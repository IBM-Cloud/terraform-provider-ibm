variable "ssh_public_key" {
}

variable "ssh_key_label" {
  default = "ssh_key_cluster"
}

variable vm_count {
  default = 2
}

variable port {
  default = 80
}

variable datacenter {
  default = "dal09"
}
