# Create single VSI in dal09. Hourly billed with private network connection only. 
resource "ibm_compute_vm_instance" "vm" {
    hostname = "vm1"
    domain = "example.com"
  bulk_vms {
    hostname = "vm2"

    domain = "example.com"
  }
  os_reference_code    = "CENTOS_7_64"
  datacenter           = "dal09"
  network_speed        = 100
  hourly_billing       = true
  private_network_only = true
  cores                = 1
  memory               = 1024
  disks                = [25]
  local_disk           = false
}

data "ibm_compute_vm_instance" "tf-vg-ds-acc-test" {
  depends_on  = [ibm_compute_vm_instance.vm]
  count       = 2
  hostname    = ibm_compute_vm_instance.vm.hostname
  domain      = ibm_compute_vm_instance.vm.domain
  most_recent = true
}

