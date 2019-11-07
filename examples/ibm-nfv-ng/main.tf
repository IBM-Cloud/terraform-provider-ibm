resource "ibm_is_vpc" "vpc1" {
  name = "f5-bigip-1nic-demo"
}

resource "ibm_is_vpc_address_prefix" "vpc-ap1" {
  name = "f5-bigip-1nic-demo-ap1"
  zone = "${var.zone1}"
  vpc  = "${ibm_is_vpc.vpc1.id}"
  cidr = "10.240.0.0/18"
}

resource "ibm_is_security_group" "sg1" {
  name = "f5-bigip-1nic-demo-sg1"
  vpc  = "${ibm_is_vpc.vpc1.id}"
}

resource "ibm_is_security_group_rule" "egress_all" {
  depends_on = ["ibm_is_floating_ip.fip1"]
  group      = "${ibm_is_vpc.vpc1.default_security_group}"
  direction  = "outbound"
  remote     = "0.0.0.0/0"

  tcp = {
    port_min = 80
    port_max = 80
  }
}

resource "ibm_is_security_group_rule" "sg1_f5_tcp_rule1" {
  depends_on = ["ibm_is_floating_ip.fip1"]
  group      = "${ibm_is_vpc.vpc1.default_security_group}"
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  tcp = {
    port_min = 8443
    port_max = 8443
  }
}

resource "ibm_is_security_group_rule" "sg1_app_tcp_rule1" {
  depends_on = ["ibm_is_floating_ip.fip1"]
  group      = "${ibm_is_vpc.vpc1.default_security_group}"
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  tcp = {
    port_min = 80
    port_max = 80
  }
}

resource "ibm_is_security_group_rule" "sg1_icmp_rule" {
  depends_on = ["ibm_is_floating_ip.fip1"]
  group      = "${ibm_is_vpc.vpc1.default_security_group}"
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  icmp = {
    code = 0
    type = 8
  }
}

resource "ibm_is_security_group_network_interface_attachment" "sgnic1" {
  security_group    = "${ibm_is_security_group.sg1.id}"
  network_interface = "${ibm_is_instance.ins1.primary_network_interface.0.id}"
}

resource "ibm_is_subnet" "subnet1" {
  name            = "f5-bigip-1nic-demo-subnet1"
  vpc             = "${ibm_is_vpc.vpc1.id}"
  zone            = "${var.zone1}"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "ssh1" {
  name       = "f5-bigip-1nic-demo-sshkey"
  public_key = "${var.ssh_public_key}"
}

resource "ibm_is_instance" "ins1" {
  name    = "f5-bigip-1nic-demo-appliance"
  image   = "${var.f5_image}"
  profile = "${var.profile}"

  primary_network_interface = {
    subnet = "${ibm_is_subnet.subnet1.id}"
  }

  vpc  = "${ibm_is_vpc.vpc1.id}"
  zone = "${var.zone1}"
  keys = ["${ibm_is_ssh_key.ssh1.id}"]

  //User can configure timeouts
  timeouts {
    create = "90m"
    delete = "30m"
  }
}

resource "ibm_is_instance" "backendapp1" {
  name    = "f5-bigip-1nic-demo-app01"
  image   = "${var.backend_image}"
  profile = "${var.profile}"

  primary_network_interface = {
    subnet = "${ibm_is_subnet.subnet1.id}"
  }

  vpc  = "${ibm_is_vpc.vpc1.id}"
  zone = "${var.zone1}"
  keys = ["${ibm_is_ssh_key.ssh1.id}"]
  user_data = "${file("nginx-userdata.sh")}"

  //User can configure timeouts
  timeouts {
    create = "90m"
    delete = "30m"
  }
}

resource "ibm_is_instance" "backendapp2" {
  name    = "f5-bigip-1nic-demo-app02"
  image   = "${var.backend_image}"
  profile = "${var.profile}"

  primary_network_interface = {
    subnet = "${ibm_is_subnet.subnet1.id}"
  }

  vpc  = "${ibm_is_vpc.vpc1.id}"
  zone = "${var.zone1}"
  keys = ["${ibm_is_ssh_key.ssh1.id}"]
  user_data = "${file("nginx-userdata.sh")}"

  //User can configure timeouts
  timeouts {
    create = "90m"
    delete = "30m"
  }
}

resource ibm_is_floating_ip "fip1" {
  name   = "f5-bigip-1nic-demo-ip1"
  target = "${ibm_is_instance.ins1.primary_network_interface.0.id}"
}

resource "ibm_is_public_gateway" "gateway1" {
  name = "f5-bigip-1nic-demo-gateway1"
  vpc  = "${ibm_is_vpc.vpc1.id}"
  zone = "${var.zone1}"

  //User can configure timeouts
  timeouts {
    create = "90m"
  }
}
