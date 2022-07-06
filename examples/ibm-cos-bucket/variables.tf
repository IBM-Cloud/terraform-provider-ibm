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

variable "single_site_loc" {
  default = "sjc04"
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

variable "expire_date" {
  default = "2022-06-09"
}

variable "expire_prefix" {
  default = ""
}

variable "nc_exp_ruleid" {
  default = "test-obj-ver-exp-3"
}

variable "nc_exp_days" {
  default = 1
}

variable "nc_exp_prefix" {
  default = ""
}

variable "abort_mpu_ruleid" {
  default = "test-abort_mpu-5"
}

variable "abort_mpu_days_init" {
  default = 1
}

variable "abort_mpu_prefix" {
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

variable "quota" {
  default = "1"
}

variable "satellite_location_id" {
  default = ""  
}
