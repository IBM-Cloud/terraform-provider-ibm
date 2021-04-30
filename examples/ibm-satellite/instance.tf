data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_is_vpc" "satellite_vpc" {
  name = "${var.is_prefix}-vpc"
}

resource "ibm_is_subnet" "satellite_subnet" {
  count                    = 3
  name                     = "${var.is_prefix}-subnet-${count.index}"
  vpc                      = ibm_is_vpc.satellite_vpc.id
  total_ipv4_address_count = 256
  zone                     = "${var.ibm_region}-${count.index + 1}"
}

resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "ibm_is_ssh_key" "satellite_ssh" {
  depends_on = [module.satellite-location]

  name        = "${var.is_prefix}-ssh"
  public_key  = var.public_key != "" ? var.public_key : tls_private_key.example.public_key_openssh
}

resource "ibm_is_instance" "satellite_instance" {
  count          = 3

  name           = "${var.is_prefix}-instance-${count.index}"
  vpc            = ibm_is_vpc.satellite_vpc.id
  zone           = "${var.ibm_region}-${count.index + 1}"
  image          = "r014-931515d2-fcc3-11e9-896d-3baa2797200f"
  profile        = "mx2-8x64"
  keys           = [ibm_is_ssh_key.satellite_ssh.id]
  resource_group = data.ibm_resource_group.resource_group.id
  user_data      = data.ibm_satellite_attach_host_script.script.host_script

  primary_network_interface {
    subnet = ibm_is_subnet.satellite_subnet[count.index].id
  }
}

resource "ibm_is_floating_ip" "satellite_ip" {
  count  = 3
  
  name   = "${var.is_prefix}-fip-${count.index}"
  target = ibm_is_instance.satellite_instance[count.index].primary_network_interface[0].id
}