provider "ibm" {}

# Create a new ssh key 
resource "ibm_compute_ssh_key" "ssh_key_performance" {
  label      = "${var.ssh_key_label}"
  notes      = "for scale group"
  public_key = "${var.ssh_public_key}"
}

resource "ibm_compute_vm_instance" "virtualguest" {
  count                   = "${var.vm_count}"
  hostname                = "ng-vm${count.index+1}"
  domain                  = "terraform.ibm.com"
  os_reference_code       = "DEBIAN_8_64"
  datacenter              = "${var.datacenter}"
  network_speed           = 10
  hourly_billing          = true
  private_network_only    = false
  cores                   = 1
  memory                  = 1024
  disks                   = [25]
  local_disk              = false
  post_install_script_uri = "https://raw.githubusercontent.com/hkantare/test/master/nginx.sh"
}

resource "ibm_lb_vpx" "citrix_vpx" {
  datacenter = "${var.datacenter}"
  speed      = 10
  version    = "10.1"
  plan       = "Standard"
  ip_count   = 2
}

resource "ibm_lb_vpx_vip" "citrix_vpx_vip" {
  name                  = "test_load_balancer_vip"
  nad_controller_id     = "${ibm_lb_vpx.citrix_vpx.id}"
  load_balancing_method = "lc"
  source_port           = "${var.port}"
  type                  = "HTTP"
  virtual_ip_address    = "${ibm_lb_vpx.citrix_vpx.vip_pool[0]}"
}

resource "ibm_lb_vpx_service" "citrix_vpx_service" {
  name                   = "ng-service${count.index+1}"
  vip_id                 = "${ibm_lb_vpx_vip.citrix_vpx_vip.id}"
  destination_ip_address = "${element(ibm_compute_vm_instance.virtualguest.*.ipv4_address, count.index)}"
  destination_port       = "${var.port}"
  weight                 = 55
  connection_limit       = 5000
  health_check           = "HTTP"
}
