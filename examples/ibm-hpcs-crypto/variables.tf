# HPCS Instance Inputs

# Enable `provision_instance` to true to create hpcs instance
variable "provision_instance" {
  type        = bool
  default     = false
  description = "Determines if the instance has to be created or not"
}

variable "hpcs_instance_name" {
  type        = string
  description = "Name of HPCS Instance"
}
variable "location" {
  default     = "us-south"
  type        = string
  description = "Location of HPCS Instance"
}
variable "plan" {
  default     = "standard"
  type        = string
  description = "Plan of HPCS Instance"
}
variable "units" {
  type        = number
  description = "No of crypto units that has to be attached to the instance."
}


# Key name that has to be created on the HPCS Instance
variable "key_name" {
  type        = string
  description = "HPCS Key Name"
}
