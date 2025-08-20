provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision pdr_managedr resource instance
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id = var.pdr_managedr_instance_id
  stand_by_redeploy = var.pdr_managedr_stand_by_redeploy
  accept_language = var.pdr_managedr_accept_language
  if_none_match = var.pdr_managedr_if_none_match
  accepts_incomplete = var.pdr_managedr_accepts_incomplete
}

// Provision pdr_validate_apikey resource instance
resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
  instance_id = var.pdr_validate_apikey_instance_id
  accept_language = var.pdr_validate_apikey_accept_language
  if_none_match = var.pdr_validate_apikey_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_get_deployment_status data source
data "ibm_pdr_get_deployment_status" "pdr_get_deployment_status_instance" {
  instance_id = var.pdr_get_deployment_status_instance_id
  if_none_match = var.pdr_get_deployment_status_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_get_event data source
data "ibm_pdr_get_event" "pdr_get_event_instance" {
  provision_id = var.pdr_get_event_provision_id
  event_id = var.pdr_get_event_event_id
  accept_language = var.pdr_get_event_accept_language
  if_none_match = var.pdr_get_event_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_get_events data source
data "ibm_pdr_get_events" "pdr_get_events_instance" {
  provision_id = var.pdr_get_events_provision_id
  time = var.pdr_get_events_time
  from_time = var.pdr_get_events_from_time
  to_time = var.pdr_get_events_to_time
  accept_language = var.pdr_get_events_accept_language
  if_none_match = var.pdr_get_events_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_get_machine_types data source
data "ibm_pdr_get_machine_types" "pdr_get_machine_types_instance" {
  instance_id = var.pdr_get_machine_types_instance_id
  primary_workspace_name = var.pdr_get_machine_types_primary_workspace_name
  accept_language = var.pdr_get_machine_types_accept_language
  if_none_match = var.pdr_get_machine_types_if_none_match
  standby_workspace_name = var.pdr_get_machine_types_standby_workspace_name
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_get_managed_vm_list data source
data "ibm_pdr_get_managed_vm_list" "pdr_get_managed_vm_list_instance" {
  instance_id = var.pdr_get_managed_vm_list_instance_id
  accept_language = var.pdr_get_managed_vm_list_accept_language
  if_none_match = var.pdr_get_managed_vm_list_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_last_operation data source
data "ibm_pdr_last_operation" "pdr_last_operation_instance" {
  instance_id = var.pdr_last_operation_instance_id
  accept_language = var.pdr_last_operation_accept_language
  if_none_match = var.pdr_last_operation_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_validate_clustertype data source
data "ibm_pdr_validate_clustertype" "pdr_validate_clustertype_instance" {
  instance_id = var.pdr_validate_clustertype_instance_id
  orchestrator_cluster_type = var.pdr_validate_clustertype_orchestrator_cluster_type
  accept_language = var.pdr_validate_clustertype_accept_language
  if_none_match = var.pdr_validate_clustertype_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_validate_proxyip data source
data "ibm_pdr_validate_proxyip" "pdr_validate_proxyip_instance" {
  instance_id = var.pdr_validate_proxyip_instance_id
  proxyip = var.pdr_validate_proxyip_proxyip
  vpc_location = var.pdr_validate_proxyip_vpc_location
  vpc_id = var.pdr_validate_proxyip_vpc_id
  if_none_match = var.pdr_validate_proxyip_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_validate_workspace data source
data "ibm_pdr_validate_workspace" "pdr_validate_workspace_instance" {
  instance_id = var.pdr_validate_workspace_instance_id
  workspace_id = var.pdr_validate_workspace_workspace_id
  crn = var.pdr_validate_workspace_crn
  location_url = var.pdr_validate_workspace_location_url
  if_none_match = var.pdr_validate_workspace_if_none_match
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
// Create pdr_get_dr_summary_response data source
data "ibm_pdr_get_dr_summary_response" "pdr_get_dr_summary_response_instance" {
  instance_id = var.pdr_get_dr_summary_response_instance_id
  accept_language = var.pdr_get_dr_summary_response_accept_language
  if_none_match = var.pdr_get_dr_summary_response_if_none_match
}
