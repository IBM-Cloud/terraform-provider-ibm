# HPCS Instance Inputs
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
  default     = 2
}
variable "signature_threshold" {
  type        = number
  default     = 1
  description = "The number of administrator signatures "
}
variable "revocation_threshold" {
  type        = number
  description = "The number of administrator signatures that is required to remove an administrator after you leave imprint mode."
  default     = 1
}
variable "admins" {
  type = list(object({
    name  = string
    key   = string
    token = string
  }))
  description = "The list of administrators for the instance crypto units. "
}

# Key name that has to be created on the HPCS Instance
variable "key_name" {
  type        = string
  description = "HPCS Key Name"
}
