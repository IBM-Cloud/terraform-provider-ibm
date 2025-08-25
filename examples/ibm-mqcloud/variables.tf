variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for mqcloud_queue_manager
variable "mqcloud_queue_manager_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
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
  default     = "reserved-eu-de-cluster-f884"
}
variable "mqcloud_queue_manager_size" {
  description = "The queue manager sizes of deployment available."
  type        = string
  default     = "xsmall"
}
variable "mqcloud_queue_manager_version" {
  description = "The MQ version of the queue manager."
  type        = string
  default     = "9.3.2_2"
}

// Resource arguments for mqcloud_application
variable "mqcloud_application_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
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
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_user_name" {
  description = "The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance."
  type        = string
  default     = "testuser"
}
variable "mqcloud_user_email" {
  description = "The email of the user."
  type        = string
  default     = "testuser@ibm.com"
}

// Resource arguments for mqcloud_keystore_certificate
variable "mqcloud_keystore_certificate_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_keystore_certificate_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}
variable "mqcloud_keystore_certificate_label" {
  description = "The label to use for the certificate to be uploaded."
  type        = string
  default     = "certlabel"
}
variable "mqcloud_keystore_certificate_certificate_file" {
  description = "The filename and path of the certificate to be uploaded."
  type        = string
  default     = "SGVsbG8gd29ybGQ="
}
variable "mqcloud_keystore_certificate_config_ams_channel_name" {
  description = "A channel's information that is configured with this certificate."
  type        = string
  default     = "CLOUD.APP.SVRCONN"
}

variable "mqcloud_keystore_certificate_config_ams_channel_name" {
  description = "A channel's information that is configured with this certificate."
  type        = string
  default     = "CLOUD.APP.SVRCONN"
}

// Resource arguments for mqcloud_truststore_certificate
variable "mqcloud_truststore_certificate_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_truststore_certificate_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}
variable "mqcloud_truststore_certificate_label" {
  description = "The label to use for the certificate to be uploaded."
  type        = string
  default     = "certlabel"
}
variable "mqcloud_truststore_certificate_certificate_file" {
  description = "The filename and path of the certificate to be uploaded."
  type        = string
  default     = "SGVsbG8gd29ybGQ="
}

// Resource arguments for mqcloud_virtual_private_endpoint_gateway
variable "mqcloud_virtual_private_endpoint_gateway_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_virtual_private_endpoint_gateway_trusted_profile" {
  description = "The CRN of the trusted profile to assume for this request."
  type        = string
  default     = "CRN of Trusted Profile"
}
variable "mqcloud_virtual_private_endpoint_gateway_name" {
  description = "The name of the virtual private endpoint gateway, created by the user."
  type        = string
  default     = "vpe-gateway1-to-vpe-gateway2"
}
variable "mqcloud_virtual_private_endpoint_gateway_target_crn" {
  description = "The CRN of the reserved capacity service instance the user is trying to connect to."
  type        = string
  default     = "Virtual Private Endpoint CRN"
}

// Data source arguments for mqcloud_queue_manager_options
variable "mqcloud_queue_manager_options_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}

// Data source arguments for mqcloud_queue_manager
variable "data_mqcloud_queue_manager_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "data_mqcloud_queue_manager_name" {
  description = "A queue manager name conforming to MQ restrictions."
  type        = string
  default     = "name"
}

// Data source arguments for mqcloud_queue_manager_status
variable "mqcloud_queue_manager_status_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "mqcloud_queue_manager_status_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}

// Data source arguments for mqcloud_application
variable "data_mqcloud_application_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "data_mqcloud_application_name" {
  description = "The name of the application - conforming to MQ rules."
  type        = string
  default     = "name"
}

// Data source arguments for mqcloud_user
variable "data_mqcloud_user_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "data_mqcloud_user_name" {
  description = "The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance."
  type        = string
  default     = "name"
}

// Data source arguments for mqcloud_truststore_certificate
variable "data_mqcloud_truststore_certificate_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "data_mqcloud_truststore_certificate_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}
variable "data_mqcloud_truststore_certificate_label" {
  description = "Certificate label in queue manager store."
  type        = string
  default     = "label"
}

// Data source arguments for mqcloud_keystore_certificate
variable "data_mqcloud_keystore_certificate_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "Service Instance ID"
}
variable "data_mqcloud_keystore_certificate_queue_manager_id" {
  description = "The id of the queue manager to retrieve its full details."
  type        = string
  default     = "Queue Manager ID"
}
variable "data_mqcloud_keystore_certificate_label" {
  description = "Certificate label in queue manager store."
  type        = string
  default     = "label"
}

// Data source arguments for mqcloud_virtual_private_endpoint_gateways
variable "mqcloud_virtual_private_endpoint_gateways_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "a2b4d4bc-dadb-4637-bcec-9b7d1e723af8"
}
variable "mqcloud_virtual_private_endpoint_gateways_trusted_profile" {
  description = "The CRN of the trusted profile to assume for this request."
  type        = string
  default     = "crn:v1:staging:public:mq-eude-stackname:eu-de:::endpoint:qm1.private.stackname.mq2.test.appdomain.cloud"
}
variable "mqcloud_virtual_private_endpoint_gateways_name" {
  description = "The name of the virtual private endpoint gateway, created by the user."
  type        = string
  default     = "name"
}

// Data source arguments for mqcloud_virtual_private_endpoint_gateway
variable "data_mqcloud_virtual_private_endpoint_gateway_service_instance_guid" {
  description = "The GUID that uniquely identifies the MQaaS service instance."
  type        = string
  default     = "a2b4d4bc-dadb-4637-bcec-9b7d1e723af8"
}
variable "data_mqcloud_virtual_private_endpoint_gateway_virtual_private_endpoint_gateway_guid" {
  description = "The id of the virtual private endpoint gateway."
  type        = string
  default     = "r010-ebab3c08-c9a8-40c4-8869-61c09ddf7b44"
}
variable "data_mqcloud_virtual_private_endpoint_gateway_trusted_profile" {
  description = "The CRN of the trusted profile to assume for this request."
  type        = string
  default     = "crn:v1:staging:public:mq-eude-stackname:eu-de:::endpoint:qm1.private.stackname.mq2.test.appdomain.cloud"
}
