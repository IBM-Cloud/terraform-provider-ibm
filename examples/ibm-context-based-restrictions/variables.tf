variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for cbr_zone
variable "cbr_zone_name" {
  description = "The name of the zone."
  type        = string
  default     = "an example of zone"
}
variable "cbr_zone_description" {
  description = "The description of the zone."
  type        = string
  default     = "A terraform example of network zone"
}
variable "cbr_zone_vpc" {
  description = "A vpc of the zone."
  type        = string
}

// Resource arguments for cbr_rule
variable "cbr_rule_description" {
  description = "The description of rule."
  type        = string
  default     = "A terraform example of rule"
}

// Data source arguments for cbr_zone
variable "cbr_zone_zone_id" {
  description = "The ID of a zone."
  type        = string
  default     = "559052eb8f43302824e7ae490c0281eb"
}

// Data source arguments for cbr_rule
variable "cbr_rule_rule_id" {
  description = "The ID of a rule."
  type        = string
  default     = "07bca38c06db1a6e125d9738c701f2c1"
}


// IBM cloud account ID
variable "ibmcloud_account_id" {
  description = "Account ID for rule / zone"
  type        = string
  default     = "12ab34cd56ef78ab90cd12ef34ab56cd"
}
