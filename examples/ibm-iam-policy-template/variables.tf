// Resource arguments for policy_template
variable "policy_template_name" {
  description = "policy template name"
  type        = string
}

// Resource arguments for policy_template
variable "policy_template_accountId" {
  description = "enterprise account id"
}

variable "policy_template_description" {
  description = "Description of the policy template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the policy for enterprise users managing IAM templates."
  type        = string
}
variable "policy_template_committed" {
  description = "Committed status of the template version."
  type        = bool
  default     = false
}