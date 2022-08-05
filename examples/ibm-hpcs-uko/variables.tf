variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for managed_key
variable "managed_key_uko_vault" {
  description = "The UUID of the Vault in which the update is to take place."
  type        = string
  default     = "uko_vault"
}
variable "managed_key_template_name" {
  description = "Name of the key template to use when creating a key."
  type        = string
  default     = "template_name"
}
variable "managed_key_label" {
  description = "The label of the key."
  type        = string
  default     = "IBM CLOUD KEY"
}
variable "managed_key_description" {
  description = "Description of the managed key."
  type        = string
  default     = "description"
}

// Resource arguments for key_template
variable "key_template_uko_vault" {
  description = "The UUID of the Vault in which the update is to take place."
  type        = string
  default     = "uko_vault"
}
variable "key_template_name" {
  description = "Name of the template, it will be referenced when creating managed keys."
  type        = string
  default     = "EXAMPLE-TEMPLATE"
}
variable "key_template_description" {
  description = "Description of the key template."
  type        = string
  default     = "description"
}

// Resource arguments for keystore
variable "keystore_uko_vault" {
  description = "The UUID of the Vault in which the update is to take place."
  type        = string
  default     = "uko_vault"
}

// Resource arguments for vault
variable "vault_name" {
  description = "A human-readable name to assign to your vault. To protect your privacy, do not use personal data, such as your name or location."
  type        = string
  default     = "Example Vault"
}
variable "vault_description" {
  description = "Description of the vault."
  type        = string
  default     = "The description of the creating vault"
}

// Data source arguments for managed_key
variable "managed_key_id" {
  description = "UUID of the key."
  type        = string
  default     = "id"
}

// Data source arguments for key_template
variable "key_template_id" {
  description = "UUID of the template."
  type        = string
  default     = "id"
}

// Data source arguments for keystore
variable "keystore_id" {
  description = "UUID of the keystore."
  type        = string
  default     = "id"
}

// Data source arguments for vault
variable "vault_id" {
  description = "UUID of the vault."
  type        = string
  default     = "id"
}
