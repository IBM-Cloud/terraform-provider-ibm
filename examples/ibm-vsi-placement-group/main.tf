# Create single VSI in placementgroup
resource "ibm_compute_placement_group" "group" {
  name       = "group"
  pod        = "pod01"
  datacenter = "dal05"
}

resource "ibm_compute_vm_instance" "vm-pgroup" {
  hostname           = "vm2"
  domain             = "example.com"
  network_speed      = 10
  hourly_billing     = true
  datacenter         = "dal05"
  cores              = 1
  memory             = 1024
  local_disk         = false
  os_reference_code  = "DEBIAN_8_64"
  disks              = [25]
  placement_group_id = "${ibm_compute_placement_group.group.id}"
}