variable "resource_group" {
  description = "IBM Cloud Resource Group name"
  type        = string
  default     = "Default"
}

variable "cis_instance_name" {
  description = "Name of the IBM Cloud Internet Services instance"
  type        = string
}

variable "domain" {
  description = "DNS domain managed by the CIS instance"
  type        = string
}
