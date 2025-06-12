// This output allows vmaas_vdc data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_vmaas_vdc" {
  value       = ibm_vmaas_vdc.vmaas_vdc_instance
  description = "vmaas_vdc resource instance"
}
