#############################################################################
#This file will deploy a number (defined in the vm_count variable) of VSIs to the data 
#center along with a single common file storage. Each VSI will be connected to a common 
#private VLAN, SSH key, security group and file storage. Each VSI will have a unique host 
#name (incremntally created) and SL allocated IP address.
#############################################################################

########################################################
#Create single SSH key for all VSIs
########################################################

resource "ibm_compute_ssh_key" "ssh_key" {
  label      = "${var.ssh_label}"
  notes      = "${var.ssh_notes}"
  public_key = "${var.ssh_key}"
}

########################################################
#Create one or more VSIs based on "count" value DC2
########################################################

resource "ibm_compute_vm_instance" "app1" {
  count             = "${var.vm_count_app}"
  os_reference_code = "${var.osrefcode}"

  # incrementally created hostnames
  hostname   = "${format("app1%02d", count.index + 1)}"
  domain     = "${var.domain}"
  datacenter = "${var.datacenter1}"

  #one or more file storage based on array
  network_speed = 100
  cores         = 1
  memory        = 1024
  disks         = [25, 10]
  ssh_key_ids   = ["${ibm_compute_ssh_key.ssh_key.id}"]
  local_disk    = false

  private_security_group_ids = ["${ibm_security_group.sg_private_lamp.id}"]
  public_security_group_ids  = ["${ibm_security_group.sg_public_lamp.id}"]
  private_network_only       = false
  user_metadata              = "${data.template_cloudinit_config.app_userdata.rendered}"
  tags                       = ["group:webserver"]
}

resource "ibm_compute_vm_instance" "db1" {
  count             = "${var.vm_count_db}"
  os_reference_code = "${var.osrefcode}"

  # incrementally created hostnames
  hostname   = "${format("db1%02d", count.index + 1)}"
  domain     = "${var.domain}"
  datacenter = "${var.datacenter1}"

  #one or more file storage based on array
  network_speed = 100
  cores         = 1
  memory        = 1024
  disks         = [25, 10]
  ssh_key_ids   = ["${ibm_compute_ssh_key.ssh_key.id}"]
  local_disk    = false

  private_security_group_ids = ["${ibm_security_group.sg_private_lampdb.id}"]
  public_security_group_ids  = ["${ibm_security_group.sg_public_lamp.id}"]
  private_network_only       = false
  user_metadata              = "${data.template_cloudinit_config.db_userdata.rendered}"
  tags                       = ["group:database"]
}

########################################################
#Create one or more VSIs based on "count" value DC2
########################################################

resource "ibm_compute_vm_instance" "app2" {
  count             = "${var.vm_count_app}"
  os_reference_code = "${var.osrefcode}"

  # incrementally created hostnames
  hostname   = "${format("app2%02d", count.index + 1)}"
  domain     = "${var.domain}"
  datacenter = "${var.datacenter2}"

  #one or more file storage based on array
  network_speed = 100
  cores         = 1
  memory        = 1024
  disks         = [25, 10]
  ssh_key_ids   = ["${ibm_compute_ssh_key.ssh_key.id}"]
  local_disk    = false

  private_security_group_ids = ["${ibm_security_group.sg_private_lamp.id}"]
  public_security_group_ids  = ["${ibm_security_group.sg_public_lamp.id}"]
  private_network_only       = false
  user_metadata              = "${data.template_cloudinit_config.app_userdata.rendered}"
  tags                       = ["group:webserver"]
}

resource "ibm_compute_vm_instance" "db2" {
  count             = "${var.vm_count_db}"
  os_reference_code = "${var.osrefcode}"

  # incrementally created hostnames
  hostname   = "${format("db2%02d", count.index + 1)}"
  domain     = "${var.domain}"
  datacenter = "${var.datacenter2}"

  #one or more file storage based on array
  network_speed = 100
  cores         = 1
  memory        = 1024
  disks         = [25, 10]
  ssh_key_ids   = ["${ibm_compute_ssh_key.ssh_key.id}"]
  local_disk    = false

  # security group created in a block
  # private_security_group_ids = ["${ibm_security_group.sg1.id}"]
  private_security_group_ids = ["${ibm_security_group.sg_private_lampdb.id}"]

  public_security_group_ids = ["${ibm_security_group.sg_public_lamp.id}"]
  private_network_only      = false
  user_metadata             = "${data.template_cloudinit_config.db_userdata.rendered}"
  tags                      = ["group:database"]
}
