variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for sm_destination
variable "sm_destination_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_destination_name" {
  description = "Destination name."
  type        = string
  default     = "name"
}
variable "sm_destination_type" {
  description = "Destination type."
  type        = string
  default     = "ibm_cloud_database"
}
variable "sm_destination_description" {
  description = "Optional description."
  type        = string
  default     = "Production PostgreSQL database"
}
variable "sm_destination_crn" {
  description = "IBM Cloud Database service instance CRN."
  type        = string
  default     = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/e91c8f42b3d74e1a9c2f05d8b67a3e10:3f8b1c7a-9d42-4e6f-b8a5-2c1d9e7f4b83::"
}

// Resource arguments for sm_hsm_destination
variable "sm_hsm_destination_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_destination_name" {
  description = "Human-readable display name."
  type        = string
  default     = "prod-hsm-destination"
}
variable "sm_hsm_destination_description" {
  description = "Optional description. Omitted if not set."
  type        = string
  default     = "Primary Thales Luna destination"
}
variable "sm_hsm_destination_vendor" {
  description = "HSM vendor. Immutable after creation."
  type        = string
  default     = "thales_luna"
}

// Resource arguments for sm_hsm_ha_group
variable "sm_hsm_ha_group_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_ha_group_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "sm_hsm_ha_group_name" {
  description = "Human-readable display name."
  type        = string
  default     = "production-ha-group"
}
variable "sm_hsm_ha_group_description" {
  description = "Optional description. Omitted if not set."
  type        = string
  default     = "Primary HA Group for prod"
}
variable "sm_hsm_ha_group_group_label" {
  description = "HA Group virtual token label. Immutable after creation."
  type        = string
  default     = "prod_ha_group"
}
variable "sm_hsm_ha_group_group_serial" {
  description = "HA Group serial number. Immutable after creation."
  type        = string
  default     = "11323465009866"
}
variable "sm_hsm_ha_group_member_serials" {
  description = "Ordered list of HSM node member serial numbers for this HA Group. Max 100 unique members."
  type        = list(string)
  default     = ["123456","123457"]
}

// Resource arguments for sm_hsm_path
variable "sm_hsm_path_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_path_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "sm_hsm_path_name" {
  description = "Human-readable display name."
  type        = string
  default     = "hsm-node-1"
}
variable "sm_hsm_path_description" {
  description = "Optional description. Omitted if not set."
  type        = string
  default     = "description"
}
variable "sm_hsm_path_crn" {
  description = "CRN of the gateway through which the HSM is reached. Immutable."
  type        = string
  default     = "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1"
}

// Resource arguments for sm_hsm_path_build
variable "sm_hsm_path_build_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_path_build_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "sm_hsm_path_build_path_id" {
  description = "Path ID."
  type        = string
  default     = "fa3d1234-0000-0000-0000-000000000001"
}

// Resource arguments for sm_hsm_server_ca
variable "sm_hsm_server_ca_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_server_ca_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "sm_hsm_server_ca_path_id" {
  description = "Path ID."
  type        = string
  default     = "fa3d1234-0000-0000-0000-000000000001"
}
variable "sm_hsm_server_ca_name" {
  description = "Human-readable display name."
  type        = string
  default     = "HSM1_Server_CA"
}
variable "sm_hsm_server_ca_description" {
  description = "Optional description. Omitted if not set."
  type        = string
  default     = "description"
}
variable "sm_hsm_server_ca_certificate" {
  description = "PEM-encoded Server CA certificate. Immutable."
  type        = string
  default     = "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----"
}

// Resource arguments for sm_hsm_client_certificate
variable "sm_hsm_client_certificate_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_client_certificate_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "sm_hsm_client_certificate_name" {
  description = "Human-readable display name."
  type        = string
  default     = "hsm-client-cert-prod"
}
variable "sm_hsm_client_certificate_description" {
  description = "Optional description. Omitted if not set."
  type        = string
  default     = "description"
}

// Data source arguments for sm_instance
variable "sm_instance_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}

// Data source arguments for sm_destinations
variable "sm_destinations_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_destinations_state" {
  description = "Filter by destination state."
  type        = string
  default     = "placeholder"
}

// Data source arguments for sm_destination
variable "data_sm_destination_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "data_sm_destination_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}

// Data source arguments for sm_hsm_destination
variable "data_sm_hsm_destination_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "data_sm_hsm_destination_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}

// Data source arguments for sm_hsm_destinations
variable "sm_hsm_destinations_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}

// Data source arguments for sm_hsm_destination_health
variable "sm_hsm_destination_health_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_destination_health_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}

// Data source arguments for sm_hsm_ha_group
variable "data_sm_hsm_ha_group_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "data_sm_hsm_ha_group_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "data_sm_hsm_ha_group_ha_group_id" {
  description = "The identifier of the HA Group to operate on. Matches the `id` field in the `HSMHAGroup` response schema."
  type        = string
  default     = "ee52ebb6-1728-4580-8ede-13f6504e3ae0"
}

// Data source arguments for sm_hsm_ha_groups
variable "sm_hsm_ha_groups_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_ha_groups_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}

// Data source arguments for sm_hsm_path
variable "data_sm_hsm_path_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "data_sm_hsm_path_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "data_sm_hsm_path_path_id" {
  description = "Path ID."
  type        = string
  default     = "fa3d1234-0000-0000-0000-000000000001"
}

// Data source arguments for sm_hsm_paths
variable "sm_hsm_paths_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_paths_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}

// Data source arguments for sm_hsm_server_ca
variable "data_sm_hsm_server_ca_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "data_sm_hsm_server_ca_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "data_sm_hsm_server_ca_path_id" {
  description = "Path ID."
  type        = string
  default     = "fa3d1234-0000-0000-0000-000000000001"
}
variable "data_sm_hsm_server_ca_server_ca_id" {
  description = "The identifier of the Server CA to operate on. Matches the `id` field in the `ServerCa` response schema."
  type        = string
  default     = "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
}

// Data source arguments for sm_hsm_server_cas
variable "sm_hsm_server_cas_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_server_cas_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "sm_hsm_server_cas_path_id" {
  description = "Path ID."
  type        = string
  default     = "fa3d1234-0000-0000-0000-000000000001"
}

// Data source arguments for sm_hsm_client_certificate
variable "data_sm_hsm_client_certificate_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "data_sm_hsm_client_certificate_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
variable "data_sm_hsm_client_certificate_cert_id" {
  description = "The identifier of the Client Certificate to operate on. Matches the `id` field in the `ClientCertificate` response schema."
  type        = string
  default     = "7c9e6679-7425-40de-944b-e07fc1f90ae7"
}

// Data source arguments for sm_hsm_client_certificates
variable "sm_hsm_client_certificates_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "fe5c5c27-3e44-474a-9484-55892e12d6d1"
}
variable "sm_hsm_client_certificates_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
