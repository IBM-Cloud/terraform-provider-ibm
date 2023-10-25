// This output allows billing_report_snapshot data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_billing_report_snapshot" {
  value       = ibm_billing_report_snapshot.billing_report_snapshot_instance
  description = "billing_report_snapshot resource instance"
}
