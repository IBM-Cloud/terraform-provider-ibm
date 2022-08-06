variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for secret_group
variable "secret_group_name" {
  description = "The name of your secret group."
  type        = string
  default     = "my-secret-group"
}
variable "secret_group_description" {
  description = "An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this group."
}
