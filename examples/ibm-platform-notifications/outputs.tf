// This output allows notification_distribution_list_destination data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_notification_distribution_list_destination" {
  value       = ibm_notification_distribution_list_destination.notification_distribution_list_destination_instance
  description = "notification_distribution_list_destination resource instance"
}
