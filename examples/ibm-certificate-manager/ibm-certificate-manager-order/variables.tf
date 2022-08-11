variable "region" {
  description = "Region in which resource has to be provisioned."
  type        = string
  default     = "us-south"
}
variable "cms_name" {
  type        = string
  description = "CMS Service instance name"
}
variable "cis_name" {
  type        = string
  description = "CIS Service instance name"
}
variable "domain" {
  type        = string
  description = "CIS Domain name"
}
variable "order_name" {
  type        = string
  description = "Name of certificate that has to be orderd."
}
variable "order_description" {
  type        = string
  description = "Description of certificate that has to be orderd"
}
variable "rotate_key" {
  type        = bool
  description = "Rotate Keys"
  default     = false
}
variable "dvm" {
  type        = string
  description = "Domain Validation Method of the CIS Domain"
  default     = "dns-01"
}
