resource "ibm_is_vpc" "vpc1" {
  name = "vpc1"
}

resource "ibm_is_vpc_address_prefix" "vpc-ap1" {
  name = "vpc-ap1"
  zone = "${var.zone1}"
  vpc  = "${ibm_is_vpc.vpc1.id}"
  cidr = "10.241.0.0/24"
}

resource "ibm_is_security_group" "sg1" {
  name = "sg1"
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

resource "ibm_is_security_group_network_interface_attachment" "sgnic1" {
  security_group    = "${ibm_is_security_group.sg1.id}"
  network_interface = "${ibm_is_instance.ins1.primary_network_interface.0.id}"
}

resource "ibm_is_volume" "vol1" {
  name    = "vol1"
  profile = "10iops-tier"
  zone    = "${var.zone1}"
}

resource "ibm_is_subnet" "subnet1" {
  name            = "subnet1"
  vpc             = "${ibm_is_vpc.vpc1.id}"
  zone            = "${var.zone1}"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_vpc" "vpc2" {
  name = "vpc2"
}

resource "ibm_is_subnet" "subnet2" {
  name            = "subnet2"
  vpc             = "${ibm_is_vpc.vpc2.id}"
  zone            = "${var.zone2}"
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ipsec_policy" "example" {
  name                     = "test-ipsec"
  authentication_algorithm = "md5"
  encryption_algorithm     = "3des"
  pfs                      = "disabled"
}

resource "ibm_is_ssh_key" "ssh1" {
  name       = "ssh1"
  public_key = "${var.ssh_public_key}"
}

resource "ibm_is_instance" "ins1" {
  name    = "ins1"
  image   = "${var.image}"
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

resource ibm_is_floating_ip "fip1" {
  name   = "ip1"
  target = "${ibm_is_instance.ins1.primary_network_interface.0.id}"
}

resource "ibm_is_public_gateway" "gateway1" {
  name = "gateway1"
  vpc  = "${ibm_is_vpc.vpc1.id}"
  zone = "${var.zone1}"

  //User can configure timeouts
  timeouts {
    create = "90m"
  }
}

resource "ibm_is_lb" "lb" {
    name = "loadbalancer1"
    subnets = ["${ibm_is_subnet.subnet1.id}"]
}
resource "ibm_is_lb_listener" "testacc_lb_listener" {
  lb       = "${ibm_is_lb.lb.id}"
  port     = "9080"
  protocol = "http"
}
resource "ibm_is_lb_pool" "webapptier-lb-pool" {
  lb                 = "${ibm_is_lb.lb.id}"
  name               = "a-webapptier-lb-pool"
  protocol           = "http"
  algorithm          = "round_robin"
  health_delay       = "5"
  health_retries     = "2"
  health_timeout     = "2"
  health_type        = "http"
  health_monitor_url = "/"
  depends_on = ["ibm_is_lb_listener.testacc_lb_listener"]
}
resource "ibm_is_lb_pool_member" "webapptier-lb-pool-member-zone1" {
  count = "2"
  lb    = "${ibm_is_lb.lb.id}"
  pool  = "${element(split("/",ibm_is_lb_pool.webapptier-lb-pool.id),1)}"
  port  = "86"
  target_address = "192.168.0.1"
  depends_on = ["ibm_is_lb_listener.testacc_lb_listener"]
}

resource "ibm_is_vpn_gateway" "VPNGateway1" {
  name   = "vpn1"
  subnet = "${ibm_is_subnet.subnet1.id}"
}

resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection1" {
  name          = "vpnconn1"
  vpn_gateway   = "${ibm_is_vpn_gateway.VPNGateway1.id}"
  peer_address  = "${ibm_is_vpn_gateway.VPNGateway1.public_ip_address}"
  preshared_key = "VPNDemoPassword"
  local_cidrs   = ["${ibm_is_subnet.subnet1.ipv4_cidr_block}"]
  peer_cidrs    = ["${ibm_is_subnet.subnet2.ipv4_cidr_block}"]
  ipsec_policy  = "${ibm_is_ipsec_policy.example.id}"
}
