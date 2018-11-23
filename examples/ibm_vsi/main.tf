# Create single VSI in dal09. Hourly billed with private network connection only. 

resource "ibm_compute_vm_instance" "vm1" {
  hostname             = "vm1"
  domain               = "example.com"
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

# Create single VSI in placementgroup
resource "ibm_compute_placement_group" "group" {
  name       = "terraform_group"
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

# Create a VSI using datacenter_option
resource "ibm_compute_vm_instance" "terraform-retry" {
  hostname          = "vm3"
  domain            = "example.com"
  network_speed     = 100
  hourly_billing    = true
  cores             = 1
  memory            = 1024
  local_disk        = false
  os_reference_code = "DEBIAN_7_64"
  disks             = [25]

  datacenter_choice = [
    {
      datacenter      = "dal09"
      public_vlan_id  = 123245
      private_vlan_id = 123255
    },
    {
      datacenter = "wdc54"
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 153345
      private_vlan_id = 123255
    },
    {
      datacenter = "dal06"
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 123245
      private_vlan_id = 123255
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 1232454
      private_vlan_id = 1234567
    },
  ]
}
