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

# COS Credentials
variable "api_key" {
  type        = string
  description = "api key of the COS bucket"
}
variable "cos_crn" {
  type        = string
  description = "COS instance CRN"
}
variable "endpoint" {
  type        = string
  description = "COS endpoint"
}
variable "bucket_name" {
  type        = string
  description = "COS bucket name"
}

# Input Json file
variable "input_file_name" {
  type        = string
  description = "Input json file name that is present in the cos-bucket or in the local"
}
# Path to which CLOUDTKEFILES has to be exported
variable "tke_files_path" {
  default     = "/Users/Kavya"
  type        = string
  description = "Path to which tke files has to be exported"
}
# Key name that has to be created on the HPCS Instance
variable "key_name" {
  type        = string
  description = "HPCS Key Name"
}
