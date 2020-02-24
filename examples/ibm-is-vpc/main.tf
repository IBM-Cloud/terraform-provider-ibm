resource "ibm_is_vpc" "vpc1" {
  name = "vpc1"
}

resource "ibm_is_vpc_route" "route" {
  name        = "route1"
  vpc         = ibm_is_vpc.vpc1.id
  zone        = var.zone1
  destination = "192.168.4.0/24"
  next_hop    = "10.240.0.4"
  depends_on  = ["ibm_is_subnet.subnet1"]
}

resource "ibm_is_subnet" "subnet1" {
  name            = "subnet1"
  vpc             = ibm_is_vpc.vpc1.id
  zone            = var.zone1
  ipv4_cidr_block = "10.240.0.0/28"
}

resource "ibm_is_vpn_gateway" "VPNGateway1" {
  name   = "vpn1"
  subnet = ibm_is_subnet.subnet1.id
}

resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection1" {
  name          = "vpnconn1"
  vpn_gateway   = ibm_is_vpn_gateway.VPNGateway1.id
  peer_address  = ibm_is_vpn_gateway.VPNGateway1.public_ip_address
  preshared_key = "VPNDemoPassword"
  local_cidrs   = [ibm_is_subnet.subnet1.ipv4_cidr_block]
  peer_cidrs    = [ibm_is_subnet.subnet2.ipv4_cidr_block]
  ipsec_policy  = ibm_is_ipsec_policy.example.id
}

resource "ibm_is_ssh_key" "sshkey" {
  name       = "ssh1"
  public_key = file(var.ssh_public_key)
}

resource "ibm_is_instance" "instance1" {
  name    = "instance1"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet1.id
  }

  vpc       = ibm_is_vpc.vpc1.id
  zone      = var.zone1
  keys      = [ibm_is_ssh_key.sshkey.id]
  user_data = file("nginx.sh")
}

resource "ibm_is_floating_ip" "floatingip1" {
  name   = "fip1"
  target = ibm_is_instance.instance1.primary_network_interface[0].id
}

resource "ibm_is_security_group_rule" "sg1_tcp_rule" {
  depends_on = [ibm_is_floating_ip.floatingip1]
  group      = ibm_is_vpc.vpc1.default_security_group
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  tcp {
    port_min = 22
    port_max = 22
  }
}

resource "ibm_is_security_group_rule" "sg1_icmp_rule" {
  depends_on = [ibm_is_floating_ip.floatingip1]
  group      = ibm_is_vpc.vpc1.default_security_group
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  icmp {
    code = 0
    type = 8
  }
}

resource "ibm_is_security_group_rule" "sg1_app_tcp_rule" {
  depends_on = [ibm_is_floating_ip.floatingip1]
  group      = ibm_is_vpc.vpc1.default_security_group
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  tcp {
    port_min = 80
    port_max = 80
  }
}

resource "ibm_is_vpc" "vpc2" {
  name = "vpc2"
}

resource "ibm_is_subnet" "subnet2" {
  name            = "subnet2"
  vpc             = ibm_is_vpc.vpc2.id
  zone            = var.zone2
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ipsec_policy" "example" {
  name                     = "test-ipsec"
  authentication_algorithm = "md5"
  encryption_algorithm     = "3des"
  pfs                      = "disabled"
}

resource "ibm_is_ike_policy" "example" {
  name                     = "test-ike"
  authentication_algorithm = "md5"
  encryption_algorithm     = "3des"
  dh_group                 = 2
  ike_version              = 1
}

resource "ibm_is_vpn_gateway" "VPNGateway2" {
  name   = "vpn2"
  subnet = ibm_is_subnet.subnet2.id
}

resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection2" {
  name           = "vpnconn2"
  vpn_gateway    = ibm_is_vpn_gateway.VPNGateway2.id
  peer_address   = ibm_is_vpn_gateway.VPNGateway2.public_ip_address
  preshared_key  = "VPNDemoPassword"
  local_cidrs    = [ibm_is_subnet.subnet2.ipv4_cidr_block]
  peer_cidrs     = [ibm_is_subnet.subnet1.ipv4_cidr_block]
  admin_state_up = true
  ike_policy     = ibm_is_ike_policy.example.id
}

resource "ibm_is_instance" "instance2" {
  name    = "instance2"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  vpc       = ibm_is_vpc.vpc2.id
  zone      = var.zone2
  keys      = [ibm_is_ssh_key.sshkey.id]
  user_data = file("nginx.sh")
}

resource "ibm_is_floating_ip" "floatingip2" {
  name   = "fip2"
  target = ibm_is_instance.instance2.primary_network_interface[0].id
}

resource "ibm_is_security_group_rule" "sg2_tcp_rule" {
  depends_on = [ibm_is_floating_ip.floatingip2]
  group      = ibm_is_vpc.vpc2.default_security_group
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  tcp {
    port_min = 22
    port_max = 22
  }
}

resource "ibm_is_security_group_rule" "sg2_icmp_rule" {
  depends_on = [ibm_is_floating_ip.floatingip2]
  group      = ibm_is_vpc.vpc2.default_security_group
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  icmp {
    code = 0
    type = 8
  }
}

resource "ibm_is_security_group_rule" "sg2_app_tcp_rule" {
  depends_on = [ibm_is_floating_ip.floatingip2]
  group      = ibm_is_vpc.vpc2.default_security_group
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  tcp {
    port_min = 80
    port_max = 80
  }
}

resource "ibm_is_volume" "vol1" {
  name    = "vol1"
  profile = "10iops-tier"
  zone    = var.zone1
}

resource "ibm_is_volume" "vol2" {
  name     = "vol2"
  profile  = "custom"
  zone     = var.zone1
  iops     = 1000
  capacity = 200
}

resource "ibm_is_network_acl" "isExampleACL" {
  name = "is-example-acl"
  rules {
    name        = "outbound"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "outbound"
    tcp {
      port_max        = 65535
      port_min        = 1
      source_port_max = 60000
      source_port_min = 22
    }
  }
  rules {
    name        = "inbound"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    tcp {
      port_max        = 65535
      port_min        = 1
      source_port_max = 60000
      source_port_min = 22
    }
  }
}

resource "ibm_is_public_gateway" "publicgateway1" {
  name = "gateway1"
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = var.zone1
}