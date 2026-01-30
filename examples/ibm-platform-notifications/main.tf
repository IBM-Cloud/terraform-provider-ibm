provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision notification_distribution_list_destination resource instance
resource "ibm_notification_distribution_list_destination" "notification_distribution_list_destination_instance" {
  account_id = var.notification_distribution_list_destination_account_id
  destination_id = var.notification_distribution_list_destination_destination_id
  destination_type = var.notification_distribution_list_destination_destination_type
}
