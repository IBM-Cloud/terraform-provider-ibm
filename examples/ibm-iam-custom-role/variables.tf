
variable "name" {
  default = "Role1"
}

variable "displayname" {
  default = "Role1Display"
}

variable "description" {
  default = "Description for role"
}

variable "servicename" {
  default = "kms"
}

variable "action" {
  default = "kms.secrets.rotate"
}

