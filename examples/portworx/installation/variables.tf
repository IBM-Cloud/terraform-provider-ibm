
variable "cluster_name" {
  default = "mycluster"
}

variable "secret_cert" {
  description = "the etcd certificate in base 64 format"
}

variable "secret_username" {
  description = "the etcd username in base 64 format"
}
variable "secret_password" {
  description = "the etcd password in base 64 format"
}

variable "etcd_endpoint" {
  description = "the etcd endpoint url"
}

