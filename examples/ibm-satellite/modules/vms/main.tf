// provider "ibm" {
//   iaas_classic_api_key = var.iaas_classic_api_key
//   iaas_classic_username = var.iaas_classic_username
// }

// resource "ibm_compute_ssh_key" "test_ssh_key" {
//     label = var.label
//     notes = var.notes
//     public_key = var.public_key
// }

resource "ibm_compute_vm_instance" "vm1" {
  depends_on = [var.module_depends_on]
  count = var.vmcount
  hostname             =   "${var.vmname}${count.index}"
  domain               = var.domain
  ssh_key_ids = [var.ssh_key_id]
  os_reference_code    = var.os
  # image_id = "4030426"
  datacenter           = var.datacenter
  network_speed        = var.network_speed
  hourly_billing       = true
#   private_network_only = true
  cores                = var.cores
  memory               = var.memory
  disks                = [var.disks]
  local_disk           = false
  # post_install_script_uri = "./addHost.sh"
}

output "vm_ip" {
  value = ibm_compute_vm_instance.vm1.*.ipv4_address
}

output "host_vm_names" {
  value = ibm_compute_vm_instance.vm1.*.hostname
}