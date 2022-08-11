variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for collectors
variable "collectors_name" {
  description = "A unique name for your collector."
  type        = string
  default     = "IBM-collector-sample"
}
variable "collectors_is_public" {
  description = "Determines whether the collector endpoint is accessible on a public network. If set to `true`, the collector connects to resources in your account over a public network. If set to `false`, the collector connects to resources by using a private IP that is accessible only through the IBM Cloud private network."
  type        = bool
  default     = true
}
variable "collectors_managed_by" {
  description = "Determines whether the collector is an IBM or customer-managed virtual machine. Use `ibm` to allow Security and Compliance Center to create, install, and manage the collector on your behalf. The collector is installed in an OpenShift cluster and approved automatically for use. Use `customer` if you would like to install the collector by using your own virtual machine. For more information, check out the [docs](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-collector)."
  type        = string
  default     = "customer"
}
variable "collectors_description" {
  description = "A detailed description of the collector."
  type        = string
  default     = "sample collector"
}
variable "collectors_passphrase" {
  description = "To protect the credentials that you add to the service, a passphrase is used to generate a data encryption key. The key is used to securely store your credentials and prevent anyone from accessing them."
  type        = string
  default     = "secret"
}
variable "collectors_is_ubi_image" {
  description = "Determines whether the collector has a Ubi image."
  type        = bool
  default     = true
}

// Resource arguments for scopes
variable "scopes_name" {
  description = "A unique name for your scope."
  type        = string
  default     = "IBMSchema-new-048-test1"
}
variable "scopes_description" {
  description = "A detailed description of the scope."
  type        = string
  default     = "IBMSchema1"
}
variable "scopes_collector_ids" {
  description = "The unique IDs of the collectors that are attached to the scope."
  type        = list(string)
  default     = ["3"]
}
variable "scopes_credential_id" {
  description = "The unique identifier of the credential."
  type        = string
  default     = "4"
}
variable "scopes_credential_type" {
  description = "The environment that the scope is targeted to."
  type        = string
  default     = "ibm"
}

// Resource arguments for credentials
variable "credentials_enabled" {
  description = "Credentials status enabled/disbaled."
  type        = bool
  default     = true
}
variable "credentials_type" {
  description = "Credentials type."
  type        = string
  default     = "ibm_cloud"
}
variable "credentials_name" {
  description = "Credentials name."
  type        = string
  default     = "test_create1"
}
variable "credentials_description" {
  description = "Credentials description."
  type        = string
  default     = "This credential is used for testing"
}
variable "credentials_purpose" {
  description = "Purpose for which the credential is created."
  type        = string
  default     = "discovery_fact_collection_remediation"
}

// Data source arguments for list_scopes

// Data source arguments for profileDetails
variable "profileDetails_id" {
  description = "The id for the given API."
  type        = string
}
variable "profileDetails_profile_type" {
  description = "The profile type ID. This will be 4 for profiles and 6 for group profiles."
  type        = string
  default     = "4"
}

// Data source arguments for list_profiles

// Data source arguments for list_latest_scans
variable "list_latest_scans_scan_id" {
}

// Data source arguments for scans_summary
variable "scans_summary_scan_id" {
  description = "Your Scan ID."
  type        = string
}
variable "scans_summary_profile_id" {
  description = "The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID."
  type        = string
}

// Data source arguments for scan_summaries
variable "scan_summaries_report_setting_id" {
  description = "The report setting ID. This can be obtained from the /validations/latest_scans API call."
  type        = string
}

// Data source arguments for group_profile_details
variable "group_profile_details_profile_id" {
  description = "The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID."
  type        = string
}

// Data source arguments for scope_correlation
variable "scope_correlation_correlation_id" {
  description = "A correlation_Id is created when a scope is created and discovery task is triggered or when a validation is triggered on a Scope. This is used to get the status of the task(discovery or validation)."
  type        = string
}

// Data source arguments for scope
variable "scope_id" {
	description = "The scope ID. This can be obtained from the Security and Compliance Center UI by clicking on the scope name. The URL contains the ID."
  	type        = string
}

// Data source arguments for credential
variable "credential_id" {
	description = "The collector ID. This can be obtained from the Security and Compliance Center UI by clicking on the credential name. The network tab contains the ID."
  	type        = string
}

// Data source arguments for collector
variable "collector_id" {
	description = "The collector ID. This can be obtained from the Security and Compliance Center UI by clicking on the collector name. The network tab contains the ID."
  	type        = string
}