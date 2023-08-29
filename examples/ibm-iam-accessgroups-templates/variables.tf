variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_access_group_template
variable "iam_access_group_template_transaction_id" {
  description = "An optional transaction id for the request."
  type        = string
  default     = "transaction_id"
}
variable "iam_access_group_template_name" {
  description = "The name of the access group template."
  type        = string
  default     = "IAM Admin Group template"
}
variable "iam_access_group_template_description" {
  description = "The description of the access group template."
  type        = string
  default     = "This access group template allows admin access to all IAM platform services in the account."
}

// Resource arguments for iam_access_group_template_version
variable "iam_access_group_template_version_template_id" {
  description = "ID of the template that you want to create a new version of."
  type        = string
  default     = "template_id"
}
variable "iam_access_group_template_version_transaction_id" {
  description = "An optional transaction id for the request."
  type        = string
  default     = "transaction_id"
}
variable "iam_access_group_template_version_name" {
  description = "The name of the access group template."
  type        = string
  default     = "IAM Admin Group template 2"
}
variable "iam_access_group_template_version_description" {
  description = "The description of the access group template."
  type        = string
  default     = "This access group template allows admin access to all IAM platform services in the account."
}

// Resource arguments for iam_access_group_template_assignment
variable "iam_access_group_template_assignment_transaction_id" {
  description = "An optional transaction id for the request."
  type        = string
  default     = "transaction_id"
}
variable "iam_access_group_template_assignment_template_id" {
  description = "The ID of the template that the assignment is based on."
  type        = string
  default     = "AccessGroupTemplateId-4be4"
}
variable "iam_access_group_template_assignment_template_version" {
  description = "The version of the template that the assignment is based on."
  type        = string
  default     = "1"
}
variable "iam_access_group_template_assignment_target_type" {
  description = "The type of the entity that the assignment applies to."
  type        = string
  default     = "AccountGroup"
}
variable "iam_access_group_template_assignment_target" {
  description = "The ID of the entity that the assignment applies to."
  type        = string
  default     = "0a45594d0f-123"
}

// Data source arguments for iam_access_group_template
variable "iam_access_group_template_transaction_id" {
  description = "An optional transaction id for the request."
  type        = string
  default     = "placeholder"
}
variable "iam_access_group_template_verbose" {
  description = "If `verbose=true`, IAM resource details are returned. If performance is a concern, leave the `verbose` parameter off so that details are not retrieved."
  type        = bool
  default     = true
}

// Data source arguments for ibm_iam_access_group_template_version
variable "ibm_iam_access_group_template_version_template_id" {
  description = "ID of the template that you want to list all versions of."
  type        = string
  default     = "template_id"
}

// Data source arguments for iam_access_group_template_assignment
variable "iam_access_group_template_assignment_template_id" {
  description = "Filter results by Template Id."
  type        = string
  default     = "placeholder"
}
variable "iam_access_group_template_assignment_template_version" {
  description = "Filter results by Template Version."
  type        = string
  default     = "placeholder"
}
variable "iam_access_group_template_assignment_target" {
  description = "Filter results by the assignment target."
  type        = string
  default     = "placeholder"
}
variable "iam_access_group_template_assignment_status" {
  description = "Filter results by the assignment status."
  type        = string
  default     = "placeholder"
}
variable "iam_access_group_template_assignment_transaction_id" {
  description = "An optional transaction id for the request."
  type        = string
  default     = "placeholder"
}
