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

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create sm_instance data source
data "ibm_sm_instance" "sm_instance_instance" {
  sm_instance_id = var.sm_instance_sm_instance_id
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
