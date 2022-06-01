provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Update the current account settings
resource "ibm_scc_account_settings" "ibm_scc_account_settings_instance" {
  // Optional input of location
  location {
    location_id = "us"
  }
  // Optional input of event_notifications
  event_notifications {
    // instance_crn = "instance_crn"
  }
}

// Read the current account location settings
data "ibm_scc_account_settings" "scc_account_location_settings_instance" {}

// Read the details of a given location
data "ibm_scc_account_location" "scc_account_location_instance" {
  location_id = var.scc_account_location_location_id
}

// Read all the available locations
data "ibm_scc_account_locations" "scc_account_locations_instance" {}
