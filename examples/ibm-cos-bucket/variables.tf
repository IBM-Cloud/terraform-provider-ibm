variable "bucket_name" {
  default = "a-standard-bucket-at-ams-firewall"
}

variable "resource_group_name" {
  default = "Default"
}

variable "storage" {
  default = "standard"
}

variable "region" {
  default = "us"
}

variable "archive_ruleid" {
  default = ""
}

variable "regional_loc" {
  default = "us-south"
}

variable "archive_days" {
  default = 0
}

variable "archive_types" {
  default = "ACCELERATED"
}

variable "expire_ruleid" {
  default = ""
}

variable "expire_days" {
  default = 1
}

variable "expire_prefix" {
  default = ""
}

variable "default_retention" {
  default = "0"
}

variable "minimum_retention" {
  default = "0"
}

variable "maximum_retention" {
  default = "1"
}