variable "ssh_key_path" {
  default = "~/.ssh/id_rsa.pub"
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
