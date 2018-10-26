resource "ibm_compute_vm_instance" "webapp1" {
  hostname   = "webapp1"
  count  = 1
  domain = "wcpclouduk.com"
  datacenter = "lon02"
  os_reference_code = "CENTOS_LATEST_64"
  network_speed = 100
  flavor_key_name = "C1_1X1X25"
  local_disk           = false
  private_network_only       = true
  user_metadata              = "${data.template_cloudinit_config.app_userdata.rendered}"
}