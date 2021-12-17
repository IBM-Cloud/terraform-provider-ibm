resource "ibm_is_vpc" "vpc1" {
  name = "vpc1"
}

resource "ibm_is_vpc_address_prefix" "testacc_vpc_address_prefix" {
  name        = "vpcaddressprefix"
  zone        = var.zone1
  vpc         = ibm_is_vpc.vpc1.id
	cidr        = var.cidr1
	is_default  = true
}

resource "ibm_is_vpc_route" "route1" {
  name        = "route1"
  vpc         = ibm_is_vpc.vpc1.id
  zone        = var.zone1
  destination = "192.168.4.0/24"
  next_hop    = "10.240.0.4"
  depends_on  = [ibm_is_subnet.subnet1]
}

resource "ibm_is_subnet" "subnet1" {
  name            = "subnet1"
  vpc             = ibm_is_vpc.vpc1.id
  zone            = var.zone1
  ipv4_cidr_block = "10.240.0.0/28"
}

resource "ibm_is_instance_template" "instancetemplate1" {
  name    = "testtemplate"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
    allow_ip_spoofing = true
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]

  boot_volume {
    name                             = "testbootvol"
    delete_volume_on_instance_delete = true
  }
  volume_attachments {
      delete_volume_on_instance_delete = true
      name                             = "volatt-01"
      volume_prototype {
          iops = 3000
          profile = "general-purpose"
          capacity = 200
      }
  }
}

resource "ibm_is_instance_template" "instancetemplate2" {
  name    = "testtemplate1"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
    allow_ip_spoofing = true
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]

  boot_volume {
    name                             = "testbootvol"
    delete_volume_on_instance_delete = true
  }
   volume_attachments {
        delete_volume_on_instance_delete = true
        name                             = "volatt-01"
        volume                           = ibm_is_volume.vol1.id
    }
}

// datasource for instance template
data "ibm_is_instance_template" "instancetemplates" {
	identifier = ibm_is_instance_template.instancetemplate2.id
}

resource "ibm_is_lb" "lb2" {
  name    = "mylb"
  subnets = [ibm_is_subnet.subnet1.id]
}

resource "ibm_is_lb_listener" "lb_listener2" {
  lb       = ibm_is_lb.lb2.id
  port     = "9086"
  protocol = "http"
}
resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
  lb                      = ibm_is_lb.lb2.id
  listener                = ibm_is_lb_listener.lb_listener2.listener_id
  action                  = "redirect"
  priority                = 2
  name                    = "mylistenerpolicy"
  target_http_status_code = 302
  target_url              = "https://www.google.com"
  rules {
    condition = "contains"
    type      = "header"
    field     = "1"
    value     = "2"
  }
}

resource "ibm_is_lb_listener_policy_rule" "lb_listener_policy_rule" {
  lb        = ibm_is_lb.lb2.id
  listener  = ibm_is_lb_listener.lb_listener2.listener_id
  policy    = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
  condition = "equals"
  type      = "header"
  field     = "MY-APP-HEADER"
  value     = "UpdateVal"
}

resource "ibm_is_lb_pool" "testacc_pool" {
  name           = "test_pool"
  lb             = ibm_is_lb.lb2.id
  algorithm      = "round_robin"
  protocol       = "https"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "https"
  proxy_protocol = "v1"
  session_persistence_type = "app_cookie"
  session_persistence_app_cookie_name = "cookie1"
}

resource "ibm_is_lb_pool" "testacc_pool" {
  name           = "test_pool"
  lb             = ibm_is_lb.lb2.id
  algorithm      = "round_robin"
  protocol       = "https"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "https"
  proxy_protocol = "v1"
  session_persistence_type = "http_cookie"
}

resource "ibm_is_lb_pool" "testacc_pool" {
  name           = "test_pool"
  lb             = ibm_is_lb.lb2.id
  algorithm      = "round_robin"
  protocol       = "https"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "https"
  proxy_protocol = "v1"
  session_persistence_type = "source_ip"
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
}

data "ibm_is_instance" "ds_instance" {
  name = ibm_is_instance.instance1.name
  private_key = file("~/.ssh/id_rsa")
  passphrase = ""
}


resource "ibm_is_instance_network_interface" "is_instance_network_interface" {
  instance = ibm_is_instance.instance1.id
  subnet = ibm_is_subnet.subnet1.id
  allow_ip_spoofing = true
  name = "my-network-interface"
  primary_ipv4_address = "10.0.0.5"
}

data "ibm_is_instance_network_interface" "is_instance_network_interface" {
	instance_name = ibm_is_instance.instance1.name
	network_interface_name = ibm_is_instance_network_interface.is_instance_network_interface.name
}

data "ibm_is_instance_network_interfaces" "is_instance_network_interfaces" {
	instance_name = ibm_is_instance.instance1.name
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
  encryption_algorithm     = "triple_des"
  pfs                      = "disabled"
}

resource "ibm_is_ike_policy" "example" {
  name                     = "test-ike"
  authentication_algorithm = "md5"
  encryption_algorithm     = "triple_des"
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
  dedicated_host = ibm_is_dedicated_host.is_dedicated_host.id
  vpc       = ibm_is_vpc.vpc2.id
  zone      = var.zone2
  keys      = [ibm_is_ssh_key.sshkey.id]
}

resource "ibm_is_instance" "instance3" {
  name    = "instance3"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }
  dedicated_host_group = ibm_is_dedicated_host_group.dh_group01.id
  vpc       = ibm_is_vpc.vpc2.id
  zone      = var.zone2
  keys      = [ibm_is_ssh_key.sshkey.id]
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
  vpc = ibm_is_vpc.vpc1.id
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

resource "ibm_is_network_acl_rule" "isExampleACLRule" {
  network_acl = ibm_is_network_acl.isExampleACL.id
  name           = "isexample-rule"
  action         = "allow"
  source         = "0.0.0.0/0"
  destination    = "0.0.0.0/0"
  direction      = "outbound"
  icmp {
    code = 1
    type = 1
  }
}

data "ibm_is_network_acl_rule" "testacc_dsnaclrule" {
  network_acl = ibm_is_network_acl.isExampleACL.id
  name = ibm_is_network_acl_rule.isExampleACL.name
}

data "ibm_is_network_acl_rules" "testacc_dsnaclrules" {
  network_acl = ibm_is_network_acl.isExampleACL.id
}

data "ibm_is_network_acl" "is_network_acl" {
	network_acl = ibm_is_network_acl.isExampleACL.id
}

data "ibm_is_network_acl" "is_network_acl1" {
	name = ibm_is_network_acl.isExampleACL.name
	vpc_name = ibm_is_vpc.vpc1.name
}

data "ibm_is_network_acls" "is_network_acls" {
}

resource "ibm_is_public_gateway" "publicgateway1" {
  name = "gateway1"
  vpc  = ibm_is_vpc.vpc1.id
  zone = var.zone1
}

data "ibm_is_public_gateway" "testacc_dspgw"{
  name = ibm_is_public_gateway.publicgateway1.name
}

data "ibm_is_public_gateways" "publicgateways"{
}

data "ibm_is_vpc" "vpc1" {
  name = ibm_is_vpc.vpc1.name
}

// added for vpcs datasource
data "ibm_is_vpc" "vpcs"{
}

data "ibm_is_volume_profile" "volprofile"{
  name = "general-purpose"
}

data "ibm_is_volume_profiles" "volprofiles"{
}

data "ibm_resource_group" "default" {
name = "Default" ///give your resource grp
}

resource "ibm_is_dedicated_host_group" "dh_group01" {
  family = "balanced"
  class = "bx2d"
  zone = "us-south-1"
  name = "my-dh-group-01"
  resource_group = data.ibm_resource_group.default.id
}
data "ibm_is_dedicated_host_group" "dgroup" {
	name = ibm_is_dedicated_host_group.dh_group01.name
}
resource "ibm_is_dedicated_host" "is_dedicated_host" {
  profile = "bx2d-host-152x608"
  name = "my-dedicated-host-01"
	host_group = ibm_is_dedicated_host_group.dh_group01.id
  resource_group = data.ibm_resource_group.default.id
}

data "ibm_is_dedicated_host_groups" "dgroups" {
}

data "ibm_is_dedicated_host_profile" "ibm_is_dedicated_host_profile" {
	name = "bx2d-host-152x608"
} 

data "ibm_is_dedicated_host_profiles" "ibm_is_dedicated_host_profiles" {
} 


data "ibm_is_dedicated_hosts" "dhosts" {

}

data "ibm_is_dedicated_host" "dhost" {
  name = ibm_is_dedicated_host.is_dedicated_host.name
  host_group = data.ibm_is_dedicated_host_group.dgroup.id
}

resource "ibm_is_volume" "vol3" {
  name    = "vol3"
  profile = "10iops-tier"
  zone    = var.zone1
}

// creating an instance with volumes
resource "ibm_is_instance" "instance4" {
  name    = "instance4"
  image   = var.image
  profile = var.profile

  volumes = [ ibm_is_volume.vol3.id ]

  primary_network_interface {
    subnet = ibm_is_subnet.subnet1.id
  }

  vpc       = ibm_is_vpc.vpc1.id
  zone      = var.zone1
  keys      = [ibm_is_ssh_key.sshkey.id]
}

// creating a snapshot from boot volume
resource "ibm_is_snapshot" "b_snapshot" {
  name          = "my-snapshot-boot"
  source_volume = ibm_is_instance.instance4.volume_attachments[0].volume_id
}

// creating a snapshot from data volume
resource "ibm_is_snapshot" "d_snapshot" {
  name          = "my-snapshot-data"
  source_volume = ibm_is_instance.instance4.volume_attachments[1].volume_id
}

// data source for snapshot by name
data "ibm_is_snapshot" "ds_snapshot" {
	name = "my-snapshot-boot"
}

// data source for snapshots
data "ibm_is_snapshots" "ds_snapshots" {
}

// restoring a boot volume from snapshot in a new instance
resource "ibm_is_instance" "instance5" {
  name    = "instance5"
  profile = var.profile
  boot_volume {
    name     = "boot-restore"
    snapshot = ibm_is_snapshot.b_snapshot.id
  }
  auto_delete_volume = true
  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }
  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}

// creating a volume 
resource "ibm_is_volume" "vol5" {
  name    = "vol5"
  profile = "10iops-tier"
  zone    = "us-south-2"
}

// creating a volume attachment on an existing instance using an existing volume
resource "ibm_is_instance_volume_attachment" "att1" {
  instance                            = ibm_is_instance.instance5.id
  volume                              = ibm_is_volume.vol5.id
  name                                = "vol-att-1"
  delete_volume_on_attachment_delete  = false
  delete_volume_on_instance_delete    = false
}

// creating a volume attachment on an existing instance using a new volume
resource "ibm_is_instance_volume_attachment" "att2" {
  instance                            = ibm_is_instance.instance5.id
  name                                = "vol-att-2"
  profile                             = "general-purpose"
  snapshot                            = ibm_is_snapshot.d_snapshot.id
  delete_volume_on_instance_delete    = true
  delete_volume_on_attachment_delete  = true
  volume_name                         = "vol4-restore"
}

// data source for volume attachment
data "ibm_is_instance_volume_attachment" "ds_vol_att" {
  instance  = ibm_is_instance.instance5.id
  name      = ibm_is_instance_volume_attachment.att2.name
}

// data source for volume attachments
data "ibm_is_instance_volume_attachment" "ds_vol_atts" {
  instance = ibm_is_instance.instance5.id
}

// creating an instance using an existing instance template
resource "ibm_is_instance" "instance6" {
  name              = "instance4"
  instance_template   = ibm_is_instance_template.instancetemplate1.id
}

resource "ibm_is_image" "image1" {
  href             = var.image_cos_url
  name             = "my-img-1"
  operating_system = var.image_operating_system
}

resource "ibm_is_image" "image2" {
  source_volume = data.ibm_is_instance.instance1.volume_attachments.0.volume_id
  name          = "my-img-1"
}

data "ibm_is_image" "dsimage" {
  name = ibm_is_image.image1.name
}

data "ibm_is_images" "dsimages" {
}

resource "ibm_is_instance_disk_management" "disks"{
  instance = ibm_is_instance.instance1.id 
  disks {
    name = "mydisk01"
    id = ibm_is_instance.instance1.disks.0.id
  }
}

data "ibm_is_instance_disks" "disk1" {
  instance = ibm_is_instance.instance1.id
}

data "ibm_is_instance_disk" "disk1" {
  instance = ibm_is_instance.instance1.id
  disk = data.ibm_is_instance_disks.disk1.disks.0.id
}

data "ibm_is_dedicated_host_disks" "dhdisks" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
}

data "ibm_is_dedicated_host_disk" "dhdisk" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
  disk = ibm_is_dedicated_host_disk_management.disks.disks.0.id
}

resource "ibm_is_dedicated_host_disk_management" "disks" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
  disks  {
    name = "newdisk01"
    id = data.ibm_is_dedicated_host.dhost.disks.0.id
  
  }
  disks  {
    name = "newdisk02"
    id = data.ibm_is_dedicated_host.dhost.disks.1.id
  
  }
}

data "ibm_is_operating_system" "os"{
  name = "red-8-amd64"
}

data "ibm_is_operating_systems" "oslist"{
}

resource "ibm_is_placement_group" "is_placement_group" {
  strategy = "%s"
  name = "%s"
  resource_group = data.ibm_resource_group.default.id
}

data "ibm_is_placement_group" "is_placement_group" {
  name = ibm_is_placement_group.is_placement_group.name
}

data "ibm_is_placement_groups" "is_placement_groups" {
}

## List regions 
data "ibm_is_regions" "regions" {
}
