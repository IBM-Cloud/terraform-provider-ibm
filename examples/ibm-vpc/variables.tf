variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for is_dedicated_host_group
variable "is_dedicated_host_group_class" {
  description = "The dedicated host profile class for hosts in this group."
  type        = string
  default     = "mx2"
}
variable "is_dedicated_host_group_family" {
  description = "The dedicated host profile family for hosts in this group."
  type        = string
  default     = "balanced"
}
variable "is_dedicated_host_group_name" {
  description = "The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words."
  type        = string
  default     = "placeholder"
}
variable "is_dedicated_host_group_resource_group" {
  description = "The resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used."
  type        = object({ example=string })
  default     = [ { example: "object" } ]
}
variable "is_dedicated_host_group_zone" {
  description = "The zone this dedicated host group will reside in."
  type        = object({ example=string })
  default     = [ {"name":"us-south-1"} ]
}
