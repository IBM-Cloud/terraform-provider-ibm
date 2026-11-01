variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for sm_destination
variable "sm_destination_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "bfc50c2e-d66d-4f37-9ccf-9713f8325b39"
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

// Data source arguments for sm_instance
variable "sm_instance_sm_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "bfc50c2e-d66d-4f37-9ccf-9713f8325b39"
}

// Data source arguments for sm_destinations
variable "sm_destinations_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "bfc50c2e-d66d-4f37-9ccf-9713f8325b39"
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
  default     = "bfc50c2e-d66d-4f37-9ccf-9713f8325b39"
}
variable "data_sm_destination_destination_id" {
  description = "Destination ID."
  type        = string
  default     = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
}
