variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for mqcloud_queue_manager
variable "mqcloud_queue_manager_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_queue_manager_name" {
  description = "A queue manager name conforming to MQ restrictions."
  type        = string
  default     = "testqm"
}
variable "mqcloud_queue_manager_display_name" {
  description = "A displayable name for the queue manager - limited only in length."
  type        = string
  default     = "A test queue manager"
}
variable "mqcloud_queue_manager_location" {
  description = "The locations in which the queue manager could be deployed."
  type        = string
  default     = "reserved-eu-fr-cluster-f884"
}
variable "mqcloud_queue_manager_size" {
  description = "The queue manager sizes of deployment available. Deployment of lite queue managers for aws_us_east_1 and aws_eu_west_1 locations is not available."
  type        = string
  default     = "lite"
}
variable "mqcloud_queue_manager_version" {
  description = "The MQ version of the queue manager."
  type        = string
  default     = "9.3.2_2"
}

// Resource arguments for mqcloud_application
variable "mqcloud_application_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_application_name" {
  description = "The name of the application - conforming to MQ rules."
  type        = string
  default     = "test-app"
}

// Resource arguments for mqcloud_user
variable "mqcloud_user_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}

// Data source arguments for mqcloud_application
variable "mqcloud_application_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_user_name" {
  description = "The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance."
  type        = string
  default     = "name"
}

variable "mqcloud_application_name" {
  description = "The name of the application - conforming to MQ rules."
  type        = string
  default     = "name"
}
variable "mqcloud_user_name" {
  description = "The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance."
  type        = string
  default     = "t0scie98o57a"
}
variable "mqcloud_user_email" {
  description = "The email of the user."
  type        = string
  default     = "user@example.com"
}

// Resource arguments for mqcloud_keystore_certificate
variable "mqcloud_keystore_certificate_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_keystore_certificate_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}
variable "mqcloud_keystore_certificate_label" {
  description = "Certificate label in queue manager store."
  type        = string
  default     = "label"
}

// Resource arguments for mqcloud_truststore_certificate
variable "mqcloud_truststore_certificate_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_truststore_certificate_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}

variable "mqcloud_truststore_certificate_label" {
  description = "Certificate label in queue manager store."
  type        = string
  default     = "label"
}

// Data source arguments for mqcloud_queue_manager
variable "mqcloud_queue_manager_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}

variable "mqcloud_queue_manager_name" {
  description = "A queue manager name conforming to MQ restrictions."
  type        = string
  default     = "name"
}

// Data source arguments for mqcloud_queue_manager_status
variable "mqcloud_queue_manager_status_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_queue_manager_status_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}

// Data source arguments for mqcloud_application
variable "mqcloud_application_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}

variable "mqcloud_application_name" {
  description = "The name of the application - conforming to MQ rules."
  type        = string
  default     = "name"
}

// Data source arguments for mqcloud_user
variable "mqcloud_user_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}

variable "mqcloud_user_name" {
  description = "The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance."
  type        = string
  default     = "name"
}

// Data source arguments for mqcloud_truststore_certificate
variable "mqcloud_truststore_certificate_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_truststore_certificate_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}
variable "mqcloud_truststore_certificate_label" {
  description = "Certificate label in queue manager store."
  type        = string
  default     = "label"
}

// Data source arguments for mqcloud_keystore_certificate
variable "mqcloud_keystore_certificate_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQ on Cloud service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_keystore_certificate_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}
variable "mqcloud_keystore_certificate_label" {
  description = "Certificate label in queue manager store."
  type        = string
  default     = "label"
}
