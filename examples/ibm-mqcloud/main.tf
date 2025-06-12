provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision mqcloud_queue_manager resource instance
resource "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
  service_instance_guid = var.mqcloud_queue_manager_service_instance_guid
  name = var.mqcloud_queue_manager_name
  display_name = var.mqcloud_queue_manager_display_name
  location = var.mqcloud_queue_manager_location
  size = var.mqcloud_queue_manager_size
  version = var.mqcloud_queue_manager_version
}

// Provision mqcloud_application resource instance
resource "ibm_mqcloud_application" "mqcloud_application_instance" {
  service_instance_guid = var.mqcloud_application_service_instance_guid
  name = var.mqcloud_application_name
}

// Provision mqcloud_user resource instance
resource "ibm_mqcloud_user" "mqcloud_user_instance" {
  service_instance_guid = var.mqcloud_user_service_instance_guid
  name = var.mqcloud_user_name
  email = var.mqcloud_user_email
}

// Provision mqcloud_keystore_certificate resource instance
resource "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
  service_instance_guid = var.mqcloud_keystore_certificate_service_instance_guid
  queue_manager_id      = var.mqcloud_keystore_certificate_queue_manager_id
  label                 = var.mqcloud_keystore_certificate_label
  certificate_file      = var.mqcloud_keystore_certificate_certificate_file

  certificate_file      = var.mqcloud_keystore_certificate_certificate_file

  config {
    ams {
      channels {
        name = var.mqcloud_keystore_certificate_config_ams_channel_name
      }
    }
  }
}

// Provision mqcloud_truststore_certificate resource instance
resource "ibm_mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
  service_instance_guid = var.mqcloud_truststore_certificate_service_instance_guid
  queue_manager_id = var.mqcloud_truststore_certificate_queue_manager_id
  label = var.mqcloud_truststore_certificate_label
  certificate_file = var.mqcloud_truststore_certificate_certificate_file
}

// Provision mqcloud_virtual_private_endpoint_gateway resource instance
resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
  service_instance_guid = var.mqcloud_virtual_private_endpoint_gateway_service_instance_guid
  trusted_profile = var.mqcloud_virtual_private_endpoint_gateway_trusted_profile
  name = var.mqcloud_virtual_private_endpoint_gateway_name
  target_crn = var.mqcloud_virtual_private_endpoint_gateway_target_crn
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_queue_manager_options data source
data "ibm_mqcloud_queue_manager_options" "mqcloud_queue_manager_options_instance" {
  service_instance_guid = var.mqcloud_queue_manager_options_service_instance_guid
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_queue_manager data source
data "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
  service_instance_guid = var.data_mqcloud_queue_manager_service_instance_guid
  name = var.data_mqcloud_queue_manager_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_queue_manager_status data source
data "ibm_mqcloud_queue_manager_status" "mqcloud_queue_manager_status_instance" {
  service_instance_guid = var.mqcloud_queue_manager_status_service_instance_guid
  queue_manager_id = var.mqcloud_queue_manager_status_queue_manager_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_application data source
data "ibm_mqcloud_application" "mqcloud_application_instance" {
  service_instance_guid = var.data_mqcloud_application_service_instance_guid
  name = var.data_mqcloud_application_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_user data source
data "ibm_mqcloud_user" "mqcloud_user_instance" {
  service_instance_guid = var.data_mqcloud_user_service_instance_guid
  name = var.data_mqcloud_user_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_truststore_certificate data source
data "ibm_mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
  service_instance_guid = var.data_mqcloud_truststore_certificate_service_instance_guid
  queue_manager_id = var.data_mqcloud_truststore_certificate_queue_manager_id
  label = var.data_mqcloud_truststore_certificate_label
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_keystore_certificate data source
data "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
  service_instance_guid = var.data_mqcloud_keystore_certificate_service_instance_guid
  queue_manager_id = var.data_mqcloud_keystore_certificate_queue_manager_id
  label = var.data_mqcloud_keystore_certificate_label
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_virtual_private_endpoint_gateways data source
data "ibm_mqcloud_virtual_private_endpoint_gateways" "mqcloud_virtual_private_endpoint_gateways_instance" {
  service_instance_guid = var.mqcloud_virtual_private_endpoint_gateways_service_instance_guid
  trusted_profile = var.mqcloud_virtual_private_endpoint_gateways_trusted_profile
  name = var.mqcloud_virtual_private_endpoint_gateways_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create mqcloud_virtual_private_endpoint_gateway data source
data "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
  service_instance_guid = var.data_mqcloud_virtual_private_endpoint_gateway_service_instance_guid
  virtual_private_endpoint_gateway_guid = var.data_mqcloud_virtual_private_endpoint_gateway_virtual_private_endpoint_gateway_guid
  trusted_profile = var.data_mqcloud_virtual_private_endpoint_gateway_trusted_profile
}
*/
