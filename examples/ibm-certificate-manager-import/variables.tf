variable "certfile_path" {}
variable "region" {}

variable "host" {}

variable "key" {
  default = "private_key.key"
}
variable "cert" {
  default = "certificate.pem"
}
