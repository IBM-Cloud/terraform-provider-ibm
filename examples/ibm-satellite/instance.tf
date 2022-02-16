data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

data "ibm_is_image" "rhel7" {
  name = "ibm-redhat-7-9-minimal-amd64-4"
}

resource "ibm_is_vpc" "vpc" {
  name = "${var.is_prefix}-vpc"
}

resource "ibm_is_subnet" "satellite_subnet" {
  count                    = 3
  name                     = "${var.is_prefix}-subnet-${count.index}"
  vpc                      = ibm_is_vpc.vpc.id
  total_ipv4_address_count = 256
  zone                     = "${var.ibm_region}-${count.index + 1}"
}

resource "ibm_is_security_group_rule" "sg-rule-inbound-ssh" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 22
    port_max = 22
  }
}

# Port 80 is needed for default routes into OpenShift. A TODO is to figure out
# how to do this securely using HTTPS instead.
resource "ibm_is_security_group_rule" "sg-rule-inbound-http" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 80
    port_max = 80
  }
}

resource "ibm_is_security_group_rule" "sg-rule-inbound-https" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 443
    port_max = 443
  }
}

resource "ibm_is_security_group_rule" "sg-rule-inbound-api" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 30000
    port_max = 32767
  }
}

resource "ibm_is_security_group_rule" "sg-rule-inbound-api2" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  udp {
    port_min = 30000
    port_max = 32767
  }
}

resource "ibm_is_security_group_rule" "sg-rule-inbound-icmp" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  icmp {
    type = 8
  }
}

resource "ibm_is_security_group_rule" "sg-rule-outbound" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "outbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 1
    port_max = 65535
  }
}

# Hosts must have TCP/UDP/ICMP Layer 3 connectivity for all ports across hosts.
# You cannot block access to certain ports that might block communication across hosts.
resource "ibm_is_security_group_rule" "sg-rule-inbound-from-the-group" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "inbound"
  remote    = ibm_is_vpc.vpc.default_security_group
}

resource "ibm_is_security_group_rule" "sg-rule-outbound-to-the-group" {
  group     = ibm_is_vpc.vpc.default_security_group
  direction = "outbound"
  remote    = ibm_is_vpc.vpc.default_security_group
}


resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "ibm_is_ssh_key" "satellite_ssh" {
  depends_on = [module.satellite-location]

  name       = "${var.is_prefix}-ssh"
  public_key = var.public_key != null ? var.public_key : tls_private_key.example.public_key_openssh
}

locals {
  zones      = ["${var.ibm_region}-1", "${var.ibm_region}-2", "${var.ibm_region}-3"]
  subnet_ids = [ibm_is_subnet.satellite_subnet[0].id, ibm_is_subnet.satellite_subnet[1].id, ibm_is_subnet.satellite_subnet[2].id]
}

resource "ibm_is_instance" "satellite_instance" {
  count = var.host_count + var.addl_host_count

  name           = "${var.is_prefix}-instance-${count.index}"
  vpc            = ibm_is_vpc.vpc.id
  zone           = element(local.zones, count.index)
  image          = data.ibm_is_image.rhel7.id
  profile        = "mx2-8x64"
  keys           = [ibm_is_ssh_key.satellite_ssh.id]
  resource_group = data.ibm_resource_group.resource_group.id
  user_data      = module.satellite-location.host_script

  primary_network_interface {
    subnet = element(local.subnet_ids, count.index)
  }
}

resource "ibm_is_floating_ip" "satellite_ip" {
  count = var.host_count + var.addl_host_count

  name   = "${var.is_prefix}-fip-${count.index}"
  target = ibm_is_instance.satellite_instance[count.index].primary_network_interface[0].id
}
