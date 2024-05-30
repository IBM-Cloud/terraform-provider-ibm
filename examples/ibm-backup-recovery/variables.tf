variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for common_source_registration_request
variable "common_source_registration_request_environment" {
  description = "Specifies the environment type of the Protection Source."
  type        = string
  default     = "kPhysical"
}
variable "common_source_registration_request_name" {
  description = "A user specified name for this source."
  type        = string
  default     = "name"
}
variable "common_source_registration_request_is_internal_encrypted" {
  description = "Specifies if credentials are encrypted by internal key."
  type        = bool
  default     = true
}
variable "common_source_registration_request_encryption_key" {
  description = "Specifies the key that user has encrypted the credential with."
  type        = string
  default     = "encryption_key"
}
variable "common_source_registration_request_connection_id" {
  description = "Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user."
  type        = number
  default     = 1
}
variable "common_source_registration_request_connector_group_id" {
  description = "Specifies the connector group id of connector groups."
  type        = number
  default     = 1
}

// Data source arguments for protection_sources
variable "protection_sources_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "protection_sources_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which Sources are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_sources_include_tenants" {
  description = "If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned."
  type        = bool
  default     = false
}
variable "protection_sources_include_source_credentials" {
  description = "If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key."
  type        = bool
  default     = false
}
variable "protection_sources_encryption_key" {
  description = "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified."
  type        = string
  default     = "placeholder"
}

// Data source arguments for source_registration
variable "source_registration_ids" {
  description = "Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default."
  type        = list(number)
  default     = [ 0 ]
}
variable "source_registration_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which objects are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "source_registration_include_tenants" {
  description = "If true, the response will include Registrations which were created by all tenants which the current user has permission to see. If false, then only Registrations created by the current user will be returned."
  type        = bool
  default     = false
}
variable "source_registration_include_source_credentials" {
  description = "If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key."
  type        = bool
  default     = false
}
variable "source_registration_encryption_key" {
  description = "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified."
  type        = string
  default     = "placeholder"
}
variable "source_registration_use_cached_data" {
  description = "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source."
  type        = bool
  default     = false
}
