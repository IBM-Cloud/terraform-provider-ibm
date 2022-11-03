variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for cm_catalog
variable "cm_catalog_label" {
  description = "Display Name in the requested language."
  type        = string
  default     = "label"
}
variable "cm_catalog_label_i18n" {
  description = "A map of translated strings, by language code."
  type        = map(string)
  default     = { "key": "inner" }
}
variable "cm_catalog_short_description" {
  description = "Description in the requested language."
  type        = string
  default     = "short_description"
}
variable "cm_catalog_short_description_i18n" {
  description = "A map of translated strings, by language code."
  type        = map(string)
  default     = { "key": "inner" }
}
variable "cm_catalog_catalog_icon_url" {
  description = "URL for an icon associated with this catalog."
  type        = string
  default     = "catalog_icon_url"
}
variable "cm_catalog_tags" {
  description = "List of tags associated with this catalog."
  type        = list(string)
  default     = [ "tags" ]
}
variable "cm_catalog_disabled" {
  description = "Denotes whether a catalog is disabled."
  type        = bool
  default     = true
}
variable "cm_catalog_resource_group_id" {
  description = "Resource group id the catalog is owned by."
  type        = string
  default     = "resource_group_id"
}
variable "cm_catalog_owning_account" {
  description = "Account that owns catalog."
  type        = string
  default     = "owning_account"
}
variable "cm_catalog_kind" {
  description = "Kind of catalog. Supported kinds are offering and vpe."
  type        = string
  default     = "kind"
}
variable "cm_catalog_metadata" {
  description = "Catalog specific metadata."
  type        = map()
  default     = { "key": null }
}

// Resource arguments for cm_offering
variable "cm_offering_catalog_id" {
  description = "Catalog identifier."
  type        = string
  default     = "catalog_id"
}
variable "cm_offering_url" {
  description = "The url for this specific offering."
  type        = string
  default     = "url"
}
variable "cm_offering_crn" {
  description = "The crn for this specific offering."
  type        = string
  default     = "crn"
}
variable "cm_offering_label" {
  description = "Display Name in the requested language."
  type        = string
  default     = "label"
}
variable "cm_offering_label_i18n" {
  description = "A map of translated strings, by language code."
  type        = map(string)
  default     = { "key": "inner" }
}
variable "cm_offering_name" {
  description = "The programmatic name of this offering."
  type        = string
  default     = "name"
}
variable "cm_offering_offering_icon_url" {
  description = "URL for an icon associated with this offering."
  type        = string
  default     = "offering_icon_url"
}
variable "cm_offering_offering_docs_url" {
  description = "URL for an additional docs with this offering."
  type        = string
  default     = "offering_docs_url"
}
variable "cm_offering_offering_support_url" {
  description = "[deprecated] - Use offering.support instead.  URL to be displayed in the Consumption UI for getting support on this offering."
  type        = string
  default     = "offering_support_url"
}
variable "cm_offering_tags" {
  description = "List of tags associated with this catalog."
  type        = list(string)
  default     = [ "tags" ]
}
variable "cm_offering_keywords" {
  description = "List of keywords associated with offering, typically used to search for it."
  type        = list(string)
  default     = [ "keywords" ]
}
variable "cm_offering_created" {
  description = "The date and time this catalog was created."
  type        = string
  default     = "2021-01-31T09:44:12Z"
}
variable "cm_offering_updated" {
  description = "The date and time this catalog was last updated."
  type        = string
  default     = "2021-01-31T09:44:12Z"
}
variable "cm_offering_short_description" {
  description = "Short description in the requested language."
  type        = string
  default     = "short_description"
}
variable "cm_offering_short_description_i18n" {
  description = "A map of translated strings, by language code."
  type        = map(string)
  default     = { "key": "inner" }
}
variable "cm_offering_long_description" {
  description = "Long description in the requested language."
  type        = string
  default     = "long_description"
}
variable "cm_offering_long_description_i18n" {
  description = "A map of translated strings, by language code."
  type        = map(string)
  default     = { "key": "inner" }
}
variable "cm_offering_pc_managed" {
  description = "Offering is managed by Partner Center."
  type        = bool
  default     = true
}
variable "cm_offering_publish_approved" {
  description = "Offering has been approved to publish to permitted to IBM or Public Catalog."
  type        = bool
  default     = true
}
variable "cm_offering_share_with_all" {
  description = "Denotes public availability of an Offering - if share_enabled is true."
  type        = bool
  default     = true
}
variable "cm_offering_share_with_ibm" {
  description = "Denotes IBM employee availability of an Offering - if share_enabled is true."
  type        = bool
  default     = true
}
variable "cm_offering_share_enabled" {
  description = "Denotes sharing including access list availability of an Offering is enabled."
  type        = bool
  default     = true
}
variable "cm_offering_permit_request_ibm_public_publish" {
  description = "Is it permitted to request publishing to IBM or Public."
  type        = bool
  default     = true
}
variable "cm_offering_ibm_publish_approved" {
  description = "Indicates if this offering has been approved for use by all IBMers."
  type        = bool
  default     = true
}
variable "cm_offering_public_publish_approved" {
  description = "Indicates if this offering has been approved for use by all IBM Cloud users."
  type        = bool
  default     = true
}
variable "cm_offering_public_original_crn" {
  description = "The original offering CRN that this publish entry came from."
  type        = string
  default     = "public_original_crn"
}
variable "cm_offering_publish_public_crn" {
  description = "The crn of the public catalog entry of this offering."
  type        = string
  default     = "publish_public_crn"
}
variable "cm_offering_portal_approval_record" {
  description = "The portal's approval record ID."
  type        = string
  default     = "portal_approval_record"
}
variable "cm_offering_portal_ui_url" {
  description = "The portal UI URL."
  type        = string
  default     = "portal_ui_url"
}
variable "cm_offering_catalog_id" {
  description = "The id of the catalog containing this offering."
  type        = string
  default     = "catalog_id"
}
variable "cm_offering_catalog_name" {
  description = "The name of the catalog."
  type        = string
  default     = "catalog_name"
}
variable "cm_offering_metadata" {
  description = "Map of metadata values for this offering."
  type        = map()
  default     = { "key": null }
}
variable "cm_offering_disclaimer" {
  description = "A disclaimer for this offering."
  type        = string
  default     = "disclaimer"
}
variable "cm_offering_hidden" {
  description = "Determine if this offering should be displayed in the Consumption UI."
  type        = bool
  default     = true
}
variable "cm_offering_provider" {
  description = "Deprecated - Provider of this offering."
  type        = string
  default     = "provider"
}
variable "cm_offering_product_kind" {
  description = "The product kind.  Valid values are module, solution, or empty string."
  type        = string
  default     = "product_kind"
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
  default     = [ "tags" ]
}
variable "cm_version_name" {
  description = "Name of version. Required for virtual server image for VPC."
  type        = string
  default     = "name"
}
variable "cm_version_label" {
  description = "Display name of version. Required for virtual server image for VPC."
  type        = string
  default     = "label"
}
variable "cm_version_install_kind" {
  description = "Install type. Example: instance. Required for virtual server image for VPC."
  type        = string
  default     = "install_kind"
}
variable "cm_version_target_kinds" {
  description = "Deployment target of the content being onboarded. Current valid values are iks, roks, vcenter, power-iaas, terraform, and vpc-x86. Required for virtual server image for VPC."
  type        = list(string)
  default     = [ "target_kinds" ]
}
variable "cm_version_format_kind" {
  description = "Format of content being onboarded. Example: vsi-image. Required for virtual server image for VPC."
  type        = string
  default     = "format_kind"
}
variable "cm_version_product_kind" {
  description = "Optional product kind for the software being onboarded.  Valid values are software, module, or solution.  Default value is software."
  type        = string
  default     = "product_kind"
}
variable "cm_version_sha" {
  description = "SHA256 fingerprint of the image file. Required for virtual server image for VPC."
  type        = string
  default     = "sha"
}
variable "cm_version_version" {
  description = "Semantic version of the software being onboarded. Required for virtual server image for VPC."
  type        = string
  default     = "version"
}
variable "cm_version_working_directory" {
  description = "Optional - The sub-folder within the specified tgz file that contains the software being onboarded."
  type        = string
  default     = "working_directory"
}
variable "cm_version_zipurl" {
  description = "URL path to zip location.  If not specified, must provide content in the body of this call."
  type        = string
  default     = "zipurl"
}
variable "cm_version_target_version" {
  description = "The semver value for this new version, if not found in the zip url package content."
  type        = string
  default     = "target_version"
}
variable "cm_version_include_config" {
  description = "Add all possible configuration values to this version when importing."
  type        = bool
  default     = true
}
variable "cm_version_is_vsi" {
  description = "Indicates that the current terraform template is used to install a virtual server image."
  type        = bool
  default     = true
}
variable "cm_version_repotype" {
  description = "The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'."
  type        = string
  default     = "repotype"
}
variable "cm_version_x_auth_token" {
  description = "Authentication token used to access the specified zip file."
  type        = string
  default     = "x_auth_token"
}

// Resource arguments for cm_offering_instance
variable "cm_offering_instance_x_auth_refresh_token" {
  description = "IAM Refresh token."
  type        = string
  default     = "x_auth_refresh_token"
}
variable "cm_offering_instance_rev" {
  description = "Cloudant revision."
  type        = string
  default     = "rev"
}
variable "cm_offering_instance_url" {
  description = "url reference to this object."
  type        = string
  default     = "url"
}
variable "cm_offering_instance_crn" {
  description = "platform CRN for this instance."
  type        = string
  default     = "crn"
}
variable "cm_offering_instance_label" {
  description = "the label for this instance."
  type        = string
  default     = "label"
}
variable "cm_offering_instance_catalog_id" {
  description = "Catalog ID this instance was created from."
  type        = string
  default     = "catalog_id"
}
variable "cm_offering_instance_offering_id" {
  description = "Offering ID this instance was created from."
  type        = string
  default     = "offering_id"
}
variable "cm_offering_instance_kind_format" {
  description = "the format this instance has (helm, operator, ova...)."
  type        = string
  default     = "kind_format"
}
variable "cm_offering_instance_version" {
  description = "The version this instance was installed from (semver - not version id)."
  type        = string
  default     = "version"
}
variable "cm_offering_instance_version_id" {
  description = "The version id this instance was installed from (version id - not semver)."
  type        = string
  default     = "version_id"
}
variable "cm_offering_instance_cluster_id" {
  description = "Cluster ID."
  type        = string
  default     = "cluster_id"
}
variable "cm_offering_instance_cluster_region" {
  description = "Cluster region (e.g., us-south)."
  type        = string
  default     = "cluster_region"
}
variable "cm_offering_instance_cluster_namespaces" {
  description = "List of target namespaces to install into."
  type        = list(string)
  default     = [ "cluster_namespaces" ]
}
variable "cm_offering_instance_cluster_all_namespaces" {
  description = "designate to install into all namespaces."
  type        = bool
  default     = true
}
variable "cm_offering_instance_schematics_workspace_id" {
  description = "Id of the schematics workspace, for offering instances provisioned through schematics."
  type        = string
  default     = "schematics_workspace_id"
}
variable "cm_offering_instance_install_plan" {
  description = "Type of install plan (also known as approval strategy) for operator subscriptions. Can be either automatic, which automatically upgrades operators to the latest in a channel, or manual, which requires approval on the cluster."
  type        = string
  default     = "install_plan"
}
variable "cm_offering_instance_channel" {
  description = "Channel to pin the operator subscription to."
  type        = string
  default     = "channel"
}
variable "cm_offering_instance_created" {
  description = "date and time create."
  type        = string
  default     = "2021-01-31T09:44:12Z"
}
variable "cm_offering_instance_updated" {
  description = "date and time updated."
  type        = string
  default     = "2021-01-31T09:44:12Z"
}
variable "cm_offering_instance_metadata" {
  description = "Map of metadata values for this offering instance."
  type        = map()
  default     = { "key": null }
}
variable "cm_offering_instance_resource_group_id" {
  description = "Id of the resource group to provision the offering instance into."
  type        = string
  default     = "resource_group_id"
}
variable "cm_offering_instance_location" {
  description = "String location of OfferingInstance deployment."
  type        = string
  default     = "location"
}
variable "cm_offering_instance_disabled" {
  description = "Indicates if Resource Controller has disabled this instance."
  type        = bool
  default     = true
}
variable "cm_offering_instance_account" {
  description = "The account this instance is owned by."
  type        = string
  default     = "account"
}
variable "cm_offering_instance_kind_target" {
  description = "The target kind for the installed software version."
  type        = string
  default     = "kind_target"
}
variable "cm_offering_instance_sha" {
  description = "The digest value of the installed software version."
  type        = string
  default     = "sha"
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
