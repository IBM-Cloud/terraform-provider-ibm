resource "ibm_network_vlan" "test_vlan_public" {
  name            = "${var.vlan_name_public}"
  datacenter      = "${var.datacenter}"
  type            = "PUBLIC"
}

# Create a private vlan
resource "ibm_network_vlan" "test_vlan_private" {
  name        = "${var.vlan_name_private}"
  datacenter  = "${var.datacenter}"
  type        = "PRIVATE"
}

# Create a new ssh key
resource "ibm_compute_ssh_key" "ssh_key" {
  label = "${var.ssh_label}"
  notes = "for public vlan test"
  public_key = "${var.ssh_public_key}"
}

resource "ibm_compute_vm_instance" "vm1" {
  hostname = "${var.vm_hostname}"
  os_reference_code = "${var.vm_os_reference_code}"
  domain = "${var.vm_domain}"
  datacenter = "${var.datacenter}"
  network_speed = "${var.vm_network_speed}"
  hourly_billing = true
  private_network_only = false
  cores = "${var.vm_cores}"
  memory = "${var.vm_memory}"
  disks = "${var.vm_disks}"
  user_metadata = "{\"value\":\"newvalue\"}"
  dedicated_acct_host_only = true
  local_disk = false
  ssh_key_ids = ["${ibm_compute_ssh_key.ssh_key.id}"]
  public_vlan_id  = "${ibm_network_vlan.test_vlan_public.id}"
  private_vlan_id = "${ibm_network_vlan.test_vlan_private.id}"

}

/*resource "ibm_compute_vm_instance" "webapp1" {
  hostname             = "webapp1"
  count                = 1
  domain               = "wcpclouduk.com"
  datacenter           = "lon02"
  os_reference_code    = "CENTOS_LATEST_64"
  network_speed        = 100
  flavor_key_name      = "C1_1X1X25"
  local_disk           = false
  private_network_only = true
  user_metadata        = data.template_cloudinit_config.app_userdata.rendered
}
*/
