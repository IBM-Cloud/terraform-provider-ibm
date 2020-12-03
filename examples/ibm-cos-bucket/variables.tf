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