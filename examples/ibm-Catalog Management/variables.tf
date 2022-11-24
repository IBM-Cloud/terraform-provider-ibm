variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for cm_catalog
variable "cm_catalog_label" {
  description = "Display Name in the requested language."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_short_description" {
  description = "Description in the requested language."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_catalog_icon_url" {
  description = "URL for an icon associated with this catalog."
  type        = string
  default     = "placeholder"
}
variable "cm_catalog_tags" {
  description = "List of tags associated with this catalog."
  type        = list(string)
  default     = [ "placeholder" ]
}
// Resource arguments for cm_offering
variable "cm_offering_catalog_id" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_id"
}
variable "cm_offering_label" {
  description = "Display Name in the requested language."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_offering_icon_url" {
  description = "URL for an icon associated with this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_offering_docs_url" {
  description = "URL for an additional docs with this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_offering_support_url" {
  description = "URL to be displayed in the Consumption UI for getting support on this offering."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_tags" {
  description = "List of tags associated with this catalog."
  type        = list(string)
  default     = [ "placeholder" ]
}

// Resource arguments for cm_version
variable "cm_version_catalog_id" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_id"
}
variable "cm_version_offering_id" {
  description = "Offering identification."
  type        = string
  default     = "offering_id"
}
variable "cm_version_tags" {
  description = "Tags array."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "cm_version_target_kinds" {
  description = "Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "cm_version_zipurl" {
  description = "URL path to zip location.  If not specified, must provide content in the body of this call."
  type        = string
  default     = "placeholder"
}
variable "cm_version_target_version" {
  description = "The semver value for this new version, if not found in the zip url package content."
  type        = string
  default     = "placeholder"
}

// Resource arguments for cm_offering_instance
variable "cm_offering_instance_label" {
  description = "the label for this instance."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_catalog_id" {
  description = "Catalog ID this instance was created from."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_offering_id" {
  description = "Offering ID this instance was created from."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_kind_format" {
  description = "the format this instance has (helm, operator, ova...)."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_version" {
  description = "The version this instance was installed from (not version id)."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_cluster_id" {
  description = "Cluster ID."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_cluster_region" {
  description = "Cluster region (e.g., us-south)."
  type        = string
  default     = "placeholder"
}
variable "cm_offering_instance_cluster_namespaces" {
  description = "List of target namespaces to install into."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "cm_offering_instance_cluster_all_namespaces" {
  description = "designate to install into all namespaces."
  type        = bool
  default     = false
}

// Data source arguments for cm_catalog
variable "cm_catalog_catalog_id" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_id"
}

// Data source arguments for cm_offering
variable "cm_offering_catalog_id" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_id"
}
variable "cm_offering_offering_id" {
  description = "Offering identification."
  type        = string
  default     = "offering_id"
}

// Data source arguments for cm_version
variable "cm_version_version_loc_id" {
  description = "A dotted value of `catalogID`.`versionID`."
  type        = string
  default     = "version_loc_id"
}

// Data source arguments for cm_offering_instance
variable "cm_offering_instance_instance_identifier" {
  description = "Version Instance identifier."
  type        = string
  default     = "instance_identifier"
}
