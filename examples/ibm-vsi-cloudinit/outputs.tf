output "vsi private ip address" {
  value = "http://${ibm_compute_vm_instance.webapp1.ipv4_address_private}"
}
