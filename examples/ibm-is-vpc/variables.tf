variable "zone1" {
  default = "us-south-1"
}

variable "zone2" {
  default = "us-south-2"
}

variable "ssh_public_key" {
  default = "~/.ssh/id_rsa.pub"
}

variable "image" {
  default = "fc538f61-7dd6-4408-978c-c6b85b69fe76"
}

variable "profile" {
  default = "bc1-2x8"
}