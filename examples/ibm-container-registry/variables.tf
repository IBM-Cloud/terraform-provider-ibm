variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for cr_namespace
variable "cr_namespace_name" {
  description = "The name of the namespace."
  type        = string
  default     = "name"
}
variable "cr_namespace_resource_group_id" {
  description = "The ID of the resource group that the namespace will be created within."
  type        = string
  default     = "placeholder"
}
variable "cr_namespace_tags" {
  description = "Local tags associated with cr_namespace"
  type        = set(string)
  default     = []
}

// Resource arguments for cr_retention_policy
variable "cr_retention_policy_namespace" {
  description = "The namespace to which the retention policy is attached."
  type        = string
  default     = "birds"
}
variable "cr_retention_policy_images_per_repo" {
  description = "Determines how many images will be retained for each repository when the retention policy is executed. The value -1 denotes 'Unlimited' (all images are retained)."
  type        = number
  default     = 10
}
variable "cr_retention_policy_retain_untagged" {
  description = "Determines if untagged images are retained when executing the retention policy. This is false by default meaning untagged images will be deleted when the policy is executed."
  type        = bool
  default     = false
}
