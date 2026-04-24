// This output allows iam_serviceid_group data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_serviceid_group" {
  value       = ibm_iam_serviceid_group.iam_serviceid_group_instance
  description = "iam_serviceid_group resource instance"
}
