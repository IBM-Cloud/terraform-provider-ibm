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
variable "cbr_zone_account_id" {
  description = "The id of the account owning this zone."
  type        = string
  default     = "12ab34cd56ef78ab90cd12ef34ab56cd"
}
variable "cbr_zone_description" {
  description = "The description of the zone."
  type        = string
  default     = "this is an example of zone"
}
variable "cbr_zone_transaction_id" {
  description = "The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a transaction ID in the response.**Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each request."
  type        = string
  default     = "transaction_id"
}

// Resource arguments for cbr_rule
variable "cbr_rule_description" {
  description = "The description of the rule."
  type        = string
  default     = "this is an example of rule"
}
variable "cbr_rule_transaction_id" {
  description = "The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a transaction ID in the response.**Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each request."
  type        = string
  default     = "transaction_id"
}

// Data source arguments for cbr_zone
variable "cbr_zone_zone_id" {
  description = "The ID of a zone."
  type        = string
  default     = "zone_id"
}

// Data source arguments for cbr_rule
variable "cbr_rule_rule_id" {
  description = "The ID of a rule."
  type        = string
  default     = "rule_id"
}
