# Configure the IBM Cloud Provider

# 09-11-2018 - Added session_stickiness = "SOURCE_IP" to lbaas to avoid Wordpress logon 
# redirect loop.



# To use HTTPS for the website
# Uncomment tls_certificate_id on resource "ibm_lbaas" "lbaas1" 
# Specify frontend_protocol = "HTTPS" and frontend_port = 443
# Uncomment statement resource "ibm_compute_ssl_certificate" "lbaas-cert"

resource "ibm_compute_ssh_key" "ssh_key" {
  label      = "${var.ssh_label}"
  notes      = "${var.ssh_notes}"
  public_key = "${var.ssh_key}"
}

# Cloud-init using cloud-config
# This configuration is designed to validate a fully working website has been created.

# commands are written within ''s. []s causes runcmd to fail on host
# manage_etc_hosts set to true to explicitly call attention to the fact that this
# is the default on IBM Cloud. /etc/hosts is refreshed at each reboot from
# /etc/cloud/templates/hosts.redhat.tmpl. 

# Install Apache web server, copy index.html to avoid 403 http code
# Record final msg in /var/log/cloud-init.log

# package_upgrade not set to true. Avoids long execution time in demo mode. 
# base64_encode and gzip set to false as not supported by IBM Cloud
data "template_cloudinit_config" "app_userdata" {
  base64_encode = false
  gzip          = false

  part {
    content = <<EOF
#cloud-config
manage_etc_hosts: true
package_upgrade: false
packages:
- httpd
runcmd:
- 'cp /usr/share/httpd/noindex/index.html /var/www/html' 
- 'systemctl start httpd'
final_message: "The system is finally up, after $UPTIME seconds"
EOF
  }
}

########################################################
# Create VMs for web tier and db tier
########################################################

resource "ibm_compute_vm_instance" "app1" {
  count             = "${var.vm_count_app}"
  os_reference_code = "${var.osrefcode}"

  # incrementally created hostnames
  hostname   = "${format("app1%02d", count.index + 1)}"
  domain     = "${var.dns_domain}"
  datacenter = "${var.datacenter1}"

  #one or more file storage based on array
  network_speed = 100
  cores         = 1
  memory        = 1024
  disks         = [25, 25]
  ssh_key_ids   = ["${ibm_compute_ssh_key.ssh_key.id}"]
  local_disk    = true

  private_security_group_ids = ["${ibm_security_group.sg_private_lamp.id}"]
  public_security_group_ids  = ["${ibm_security_group.sg_public_lamp.id}"]
  private_network_only       = false
  #user_metadata              = "${data.template_cloudinit_config.app_userdata.rendered}"
  tags                       = ["group:webserver"]
}

data "template_cloudinit_config" "db_userdata" {
  base64_encode = false
  gzip          = false

  part {
    content = <<EOF
#cloud-config
manage_etc_hosts: true
package_upgrade: false
final_message: "The system is finally up, after $UPTIME seconds"
EOF
  }
}

resource "ibm_compute_vm_instance" "db1" {
  count             = "${var.vm_count_db}"
  os_reference_code = "${var.osrefcode}"

  # incrementally created hostnames
  hostname   = "${format("db1%02d", count.index + 1)}"
  domain     = "${var.dns_domain}"
  datacenter = "${var.datacenter1}"

  #one or more file storage based on array
  network_speed = 100
  cores         = 1
  memory        = 1024
  disks         = [25, 25]
  ssh_key_ids   = ["${ibm_compute_ssh_key.ssh_key.id}"]
  local_disk    = true

  private_security_group_ids = ["${ibm_security_group.sg_private_lamp.id}"]
  public_security_group_ids  = ["${ibm_security_group.sg_public_lamp.id}"]
  private_network_only       = false
  user_metadata              = "${data.template_cloudinit_config.db_userdata.rendered}"
  tags                       = ["group:database"]
}

# tag cloudloadbalancer required for Ansible dynamic inventory
resource "ibm_lbaas" "lbaas1" {
  name        = "${var.lb_name}"
  description = "lbaas example"
  subnets     = ["${ibm_compute_vm_instance.app1.*.private_subnet_id[0]}"]

  # HTTP/80 default to avoid requiement for SSL cert when used as demo
  protocols = [{
    frontend_protocol = "HTTP"
    frontend_port     = 80

    #frontend_protocol     = "HTTPS"
    #frontend_port         = 443
    backend_protocol = "HTTP"
    # Session stickiness set to avoid Wordpress admin logon loop with HTTP
    session_stickiness = "SOURCE_IP"
    backend_port          = 80
    load_balancing_method = "${var.lb_method}"

    #tls_certificate_id    = "${ibm_compute_ssl_certificate.lbaas-cert.id}"
  }]
}

# resource "ibm_compute_ssl_certificate" "lbaas-cert" {
#   certificate = "${var.ssl_cert}"
#   private_key = "${var.ssl_private_key}"
# }

resource "ibm_lbaas_server_instance_attachment" "lbaas_member" {
  count              = "${var.vm_count_app}"
  private_ip_address = "${element(ibm_compute_vm_instance.app1.*.ipv4_address_private,count.index)}"
  weight             = 40
  lbaas_id           = "${ibm_lbaas.lbaas1.id}"
  depends_on         = ["ibm_lbaas.lbaas1"]
}

resource "ibm_lbaas_health_monitor" "lbaas_hm" {
  protocol   = "${ibm_lbaas.lbaas1.health_monitors.0.protocol}"
  port       = "${ibm_lbaas.lbaas1.health_monitors.0.port}"
  timeout    = 3
  lbaas_id   = "${ibm_lbaas.lbaas1.id}"
  monitor_id = "${ibm_lbaas.lbaas1.health_monitors.0.monitor_id}"
  depends_on = ["ibm_lbaas_server_instance_attachment.lbaas_member"]
}
