variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "region" {
  description = "Secrets Manager Instance region"
  default     = null
}

variable "secrets_manager_instance_id" {
  description = "Secrets Manager Instance GUID"
  type        = string
}

// Data source arguments for secrets_manager_secrets
variable "secrets_manager_secrets_secret_type" {
  description = "The secret type."
  type        = string
  default     = null
}

// Data source arguments for secrets_manager_secret
variable "secrets_manager_secret_secret_type" {
  description = "The secret type. Supported options include: arbitrary, iam_credentials, username_password."
  type        = string
  default     = "arbitrary"
}
variable "secrets_manager_secret_id" {
  description = "The v4 UUID that uniquely identifies the secret."
  type        = string
}
