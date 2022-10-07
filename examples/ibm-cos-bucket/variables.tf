variable "bucket_name" {
  default = "a-standard-bucket-at-ams-firewall"
}

variable "resource_group_name" {
  default = "Default"
}

variable "standard_storage_class" {
  default = "standard"
}

variable "onerate_storage_class" {
  default = "onerate_active"
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

variable "replicate_ruleid" {
  default = ""
}

variable "replicate_prefix" {
  default = ""
}

variable "replicate_priority" {
  default = "1"
}

variable "delmarkerrep_status" {
  default = true
}

variable "dest_rep_bkt_crn" {
  default = ""
}

variable "hpcs_location" {
  default     = "us-south"
  type        = string
}
variable "hpcs_plan" {
  default     = "standard"
  type        = string
}
variable "hpcs_crypto_units" {
  type        = number
  default     = 2
}
variable "hpcs_signature_threshold" {
  type        = number
  default     = 1
}
variable "hpcs_revocation_threshold" {
  type        = number
  default     = 1
}
variable "hpcs_crypto_unit_admins" {
  type = list(object({
    name  = string
    key   = string
    token = string
  }))
}
# Key name that has to be created on the HPCS Instance
variable "hpcs_key_name" {
  type        = string
}
variable "hpcs_uko_rootkeycrn" {
  default = ""
}
