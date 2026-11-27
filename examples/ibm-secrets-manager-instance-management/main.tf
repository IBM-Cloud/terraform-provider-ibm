provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision sm_destination resource instance
resource "ibm_sm_destination" "sm_destination_instance" {
  instance_id = var.sm_destination_instance_id
  name = var.sm_destination_name
  type = var.sm_destination_type
  description = var.sm_destination_description
  crn = var.sm_destination_crn
}

// Provision sm_hsm_destination resource instance
resource "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
  instance_id = var.sm_hsm_destination_instance_id
  name = var.sm_hsm_destination_name
  description = var.sm_hsm_destination_description
  vendor = var.sm_hsm_destination_vendor
}

// Provision sm_hsm_ha_group resource instance
resource "ibm_sm_hsm_ha_group" "sm_hsm_ha_group_instance" {
  instance_id = var.sm_hsm_ha_group_instance_id
  destination_id = var.sm_hsm_ha_group_destination_id
  name = var.sm_hsm_ha_group_name
  description = var.sm_hsm_ha_group_description
  group_label = var.sm_hsm_ha_group_group_label
  group_serial = var.sm_hsm_ha_group_group_serial
  member_serials = var.sm_hsm_ha_group_member_serials
}

// Provision sm_hsm_path resource instance
resource "ibm_sm_hsm_path" "sm_hsm_path_instance" {
  instance_id = var.sm_hsm_path_instance_id
  destination_id = var.sm_hsm_path_destination_id
  name = var.sm_hsm_path_name
  description = var.sm_hsm_path_description
  crn = var.sm_hsm_path_crn
}

// Provision sm_hsm_path_build resource instance
resource "ibm_sm_hsm_path_build" "sm_hsm_path_build_instance" {
  instance_id = var.sm_hsm_path_build_instance_id
  destination_id = var.sm_hsm_path_build_destination_id
  path_id = var.sm_hsm_path_build_path_id
}

// Provision sm_hsm_server_ca resource instance
resource "ibm_sm_hsm_server_ca" "sm_hsm_server_ca_instance" {
  instance_id = var.sm_hsm_server_ca_instance_id
  destination_id = var.sm_hsm_server_ca_destination_id
  path_id = var.sm_hsm_server_ca_path_id
  name = var.sm_hsm_server_ca_name
  description = var.sm_hsm_server_ca_description
  certificate = var.sm_hsm_server_ca_certificate
}

// Provision sm_hsm_client_certificate resource instance
resource "ibm_sm_hsm_client_certificate" "sm_hsm_client_certificate_instance" {
  instance_id = var.sm_hsm_client_certificate_instance_id
  destination_id = var.sm_hsm_client_certificate_destination_id
  name = var.sm_hsm_client_certificate_name
  description = var.sm_hsm_client_certificate_description
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_instance data source
data "ibm_sm_instance" "sm_instance_instance" {
  instance_id = var.sm_instance_instance_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_destinations data source
data "ibm_sm_destinations" "sm_destinations_instance" {
  instance_id = var.sm_destinations_instance_id
  state = var.sm_destinations_state
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_destination data source
data "ibm_sm_destination" "sm_destination_instance" {
  instance_id = var.data_sm_destination_instance_id
  destination_id = var.data_sm_destination_destination_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_destination data source
data "ibm_sm_hsm_destination" "sm_hsm_destination_instance" {
  instance_id = var.data_sm_hsm_destination_instance_id
  destination_id = var.data_sm_hsm_destination_destination_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_destinations data source
data "ibm_sm_hsm_destinations" "sm_hsm_destinations_instance" {
  instance_id = var.sm_hsm_destinations_instance_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_destination_health data source
data "ibm_sm_hsm_destination_health" "sm_hsm_destination_health_instance" {
  instance_id = var.sm_hsm_destination_health_instance_id
  destination_id = var.sm_hsm_destination_health_destination_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_ha_group data source
data "ibm_sm_hsm_ha_group" "sm_hsm_ha_group_instance" {
  instance_id = var.data_sm_hsm_ha_group_instance_id
  destination_id = var.data_sm_hsm_ha_group_destination_id
  ha_group_id = var.data_sm_hsm_ha_group_ha_group_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_ha_groups data source
data "ibm_sm_hsm_ha_groups" "sm_hsm_ha_groups_instance" {
  instance_id = var.sm_hsm_ha_groups_instance_id
  destination_id = var.sm_hsm_ha_groups_destination_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_path data source
data "ibm_sm_hsm_path" "sm_hsm_path_instance" {
  instance_id = var.data_sm_hsm_path_instance_id
  destination_id = var.data_sm_hsm_path_destination_id
  path_id = var.data_sm_hsm_path_path_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_paths data source
data "ibm_sm_hsm_paths" "sm_hsm_paths_instance" {
  instance_id = var.sm_hsm_paths_instance_id
  destination_id = var.sm_hsm_paths_destination_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_server_ca data source
data "ibm_sm_hsm_server_ca" "sm_hsm_server_ca_instance" {
  instance_id = var.data_sm_hsm_server_ca_instance_id
  destination_id = var.data_sm_hsm_server_ca_destination_id
  path_id = var.data_sm_hsm_server_ca_path_id
  server_ca_id = var.data_sm_hsm_server_ca_server_ca_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_server_cas data source
data "ibm_sm_hsm_server_cas" "sm_hsm_server_cas_instance" {
  instance_id = var.sm_hsm_server_cas_instance_id
  destination_id = var.sm_hsm_server_cas_destination_id
  path_id = var.sm_hsm_server_cas_path_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_client_certificate data source
data "ibm_sm_hsm_client_certificate" "sm_hsm_client_certificate_instance" {
  instance_id = var.data_sm_hsm_client_certificate_instance_id
  destination_id = var.data_sm_hsm_client_certificate_destination_id
  cert_id = var.data_sm_hsm_client_certificate_cert_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_hsm_client_certificates data source
data "ibm_sm_hsm_client_certificates" "sm_hsm_client_certificates_instance" {
  instance_id = var.sm_hsm_client_certificates_instance_id
  destination_id = var.sm_hsm_client_certificates_destination_id
}
*/
