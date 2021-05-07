
##################################################
# IBMCLOUD Satellite Location Variables
##################################################

variable "location" {
  description = "Location Name"
  type         = string
}

variable "managed_from" {
  description = "The IBM Cloud region to manage your Satellite location from. Choose a region close to your on-prem data center for better performance."
  type         = string
  default = "wdc"
}

variable "location_zones" {
  description = "Allocate your hosts across these three zones"
  type        = list(string)
  default     = ["us-east-1", "us-east-2", "us-east-3"]
}

variable "is_location_exist" {
  description = "Location Name"
  type         = bool
  default      = false
}

variable "location_bucket" {
  description = "COS bucket name"
  default     = ""
}

#################################################################################################
# IBMCLOUD -  Authentication , Target Variables.
#################################################################################################

variable "ibmcloud_api_key" {
  description  = "IBM Cloud API Key"
  type         = string
}

variable "resource_group" {
  description = "Name of the resource group on which location has to be created"
}

variable "ibm_region" {
  description = "Region of the IBM Cloud account. Currently supported regions for satellite are us-east and eu-gb region."
  default     = "us-east"
}

variable "host_provider" {
    description  = "The cloud provider of host|vms"
    type         = string
    default      = "ibm"
}

variable "tags" {
  description = "List of tags associated with this satellite."
  type        = list(string)
  default     = [ "env:dev" ]
}