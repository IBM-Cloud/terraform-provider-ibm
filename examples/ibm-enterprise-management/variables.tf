variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for enterprise
variable "enterprise_source_account_id" {
  description = "The ID of the account that is used to create the enterprise."
  type        = string
  default     = "source_account_id"
}
variable "enterprise_name" {
  description = "The name of the enterprise. This field must have 3 - 60 characters."
  type        = string
  default     = "name"
}
variable "enterprise_primary_contact_iam_id" {
  description = "The IAM ID of the enterprise primary contact, such as `IBMid-0123ABC`. The IAM ID must already exist."
  type        = string
  default     = "primary_contact_iam_id"
}
variable "enterprise_domain" {
  description = "A domain or subdomain for the enterprise, such as `example.com` or `my.example.com`."
  type        = string
  default     = "placeholder"
}

// Resource arguments for enterprise_account_group
variable "enterprise_account_group_parent" {
  description = "The CRN of the parent under which the account group will be created. The parent can be an existing account group or the enterprise itself."
  type        = string
  default     = "parent"
}
variable "enterprise_account_group_name" {
  description = "The name of the account group. This field must have 3 - 60 characters."
  type        = string
  default     = "name"
}
variable "enterprise_account_group_primary_contact_iam_id" {
  description = "The IAM ID of the primary contact for this account group, such as `IBMid-0123ABC`. The IAM ID must already exist."
  type        = string
  default     = "primary_contact_iam_id"
}

// Resource arguments for enterprise_account
variable "enterprise_account_parent" {
  description = "The CRN of the parent under which the account will be created. The parent can be an existing account group or the enterprise itself."
  type        = string
  default     = "parent"
}
variable "enterprise_account_name" {
  description = "The name of the account. This field must have 3 - 60 characters."
  type        = string
  default     = "name"
}
variable "enterprise_account_owner_iam_id" {
  description = "The IAM ID of the account owner, such as `IBMid-0123ABC`. The IAM ID must already exist."
  type        = string
  default     = "owner_iam_id"
}
variable "enterprise_account_traits" {
  description = "The traits object can be used to opt-out of Multi-Factor Authenticatin or for setting enterprise IAM settings setting when creating a child account in the enterprise."
  type        = set()
  default     = { enterprise_iam_managed = false }
}

variable "enterprise_account_options" {
  description = "The options object can be used to set properties on child accounts of an enterprise. You can pass a field to to create IAM service id with IAM api key when creating a child account in the enterprise."
  type        = set()
  default     = { create_iam_service_id_with_apikey_and_owner_policies : false }
}

// Data source arguments for enterprises
variable "enterprises_name" {
  description = "The name of the enterprise."
  type        = string
  default     = "placeholder"
}

// Data source arguments for account_groups
variable "account_groups_name" {
  description = "The name of the account group."
  type        = string
  default     = "placeholder"
}

variable "account_email" {
  description = "The email of enterprise."
  type        = string
  default     = "placeholder"
}

// Data source arguments for accounts
variable "accounts_name" {
  description = "The name of the account."
  type        = string
  default     = "placeholder"
}

variable "enterprise_id" {
  description = "The unique enterpriseId of enterprise account"
  type        = string
  default     = "placeholder"
}

variable "account_id" {
  description = "The unique accountId of the standalone account which needs to be imported to enterprise"
  type        = string
  default     = "placeholder"
}