resource "ibm_is_vpc" "vpc1" {
  name = "vpc1"
}

resource "ibm_is_vpc_address_prefix" "testacc_vpc_address_prefix" {
  name       = "vpcaddressprefix"
  zone       = var.zone1
  vpc        = ibm_is_vpc.vpc1.id
  cidr       = var.cidr1
  is_default = true
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
    subnet            = ibm_is_subnet.subnet2.id
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
      iops     = 3000
      profile  = "general-purpose"
      capacity = 200
    }
  }
}

resource "ibm_is_instance_template" "instancetemplate2" {
  name    = "testtemplate1"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet            = ibm_is_subnet.subnet2.id
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

// Load balancer with private DNS
resource "ibm_is_lb" "example" {
  name    = "example-load-balancer"
  subnets = [ibm_is_subnet.subnet1.id]
  profile = "network-fixed"
  dns {
    instance_crn = "crn:v1:staging:public:dns-svcs:global:a/exxxxxxxxxxxxx-xxxxxxxxxxxxxxxxx:5xxxxxxx-xxxxx-xxxxxxxxxxxxxxx-xxxxxxxxxxxxxxx::"
    zone_id      = "bxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx"
  }
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
  name                                = "test_pool"
  lb                                  = ibm_is_lb.lb2.id
  algorithm                           = "round_robin"
  protocol                            = "https"
  health_delay                        = 60
  health_retries                      = 5
  health_timeout                      = 30
  health_type                         = "https"
  proxy_protocol                      = "v1"
  session_persistence_type            = "app_cookie"
  session_persistence_app_cookie_name = "cookie1"
}

resource "ibm_is_lb_pool" "testacc_pool" {
  name                     = "test_pool"
  lb                       = ibm_is_lb.lb2.id
  algorithm                = "round_robin"
  protocol                 = "https"
  health_delay             = 60
  health_retries           = 5
  health_timeout           = 30
  health_type              = "https"
  proxy_protocol           = "v1"
  session_persistence_type = "http_cookie"
}

resource "ibm_is_lb_pool" "testacc_pool" {
  name                     = "test_pool"
  lb                       = ibm_is_lb.lb2.id
  algorithm                = "round_robin"
  protocol                 = "https"
  health_delay             = 60
  health_retries           = 5
  health_timeout           = 30
  health_type              = "https"
  proxy_protocol           = "v1"
  session_persistence_type = "source_ip"
}

data "ibm_is_lb_listener" "is_lb_listener" {
  lb          = ibm_is_lb.lb2.id
  listener_id = ibm_is_lb_listener.lb_listener2.listener_id
}
data "ibm_is_lb_listeners" "is_lb_listeners" {
  lb = ibm_is_lb.lb2.id
}

data "ibm_is_lb_listener_policy" "is_lb_listener_policy" {
  lb        = ibm_is_lb.lb2.id
  listener  = ibm_is_lb_listener.lb_listener2.listener_id
  policy_id = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
}
data "ibm_is_lb_listener_policies" "is_lb_listener_policies" {
  lb       = ibm_is_lb.lb2.id
  listener = ibm_is_lb_listener.lb_listener2.listener_id
}

data "ibm_is_lb_listener_policy_rule" "is_lb_listener_policy_rule" {
  lb       = ibm_is_lb.lb2.id
  listener = ibm_is_lb_listener.lb_listener2.listener_id
  policy   = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
  rule     = ibm_is_lb_listener_policy_rule.lb_listener_policy_rule.rule
}

data "ibm_is_lb_listener_policy_rules" "is_lb_listener_policy_rules" {
  lb       = ibm_is_lb.lb2.id
  listener = ibm_is_lb_listener.lb_listener2.listener_id
  policy   = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
}

resource "ibm_is_vpn_gateway" "VPNGateway1" {
  name   = "vpn1"
  subnet = ibm_is_subnet.subnet1.id
}

// Deprecated: peer_address, local_cidrs, peer_cidrs
resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection1_deprecated" {
  name          = "vpnconn1-deprecated"
  vpn_gateway   = ibm_is_vpn_gateway.VPNGateway1.id
  peer_address  = ibm_is_vpn_gateway.VPNGateway1.public_ip_address
  preshared_key = "VPNDemoPassword"
  local_cidrs   = [ibm_is_subnet.subnet1.ipv4_cidr_block]
  peer_cidrs    = [ibm_is_subnet.subnet2.ipv4_cidr_block]
  ipsec_policy  = ibm_is_ipsec_policy.example.id
}

resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection1" {
  name          = "vpnconn1"
  vpn_gateway   = ibm_is_vpn_gateway.VPNGateway1.id
  peer_address  = ibm_is_vpn_gateway.VPNGateway1.public_ip_address
  preshared_key = "VPNDemoPassword"
  peer {
    address    = ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address != "0.0.0.0" ? ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address : ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address2
    peer_cidrs = [ibm_is_subnet.subnet2.ipv4_cidr_block]
  }
  local {
    cidrs = [ibm_is_subnet.subnet1.ipv4_cidr_block]
  }
  ipsec_policy = ibm_is_ipsec_policy.example.id
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

  vpc  = ibm_is_vpc.vpc1.id
  zone = var.zone1
  keys = [ibm_is_ssh_key.sshkey.id]
}

data "ibm_is_instance" "ds_instance" {
  name        = ibm_is_instance.instance1.name
  private_key = file("~/.ssh/id_rsa")
  passphrase  = ""
}


resource "ibm_is_instance_network_interface" "is_instance_network_interface" {
  instance             = ibm_is_instance.instance1.id
  subnet               = ibm_is_subnet.subnet1.id
  allow_ip_spoofing    = true
  name                 = "my-network-interface"
  primary_ipv4_address = "10.0.0.5"
}

data "ibm_is_instance_network_interface" "is_instance_network_interface" {
  instance_name          = ibm_is_instance.instance1.name
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
  authentication_algorithm = "sha256"
  encryption_algorithm     = "aes128"
  pfs                      = "disabled"
}

resource "ibm_is_ike_policy" "example" {
  name                     = "test-ike"
  authentication_algorithm = "sha256"
  encryption_algorithm     = "aes128"
  dh_group                 = 14
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
  vpc            = ibm_is_vpc.vpc2.id
  zone           = var.zone2
  keys           = [ibm_is_ssh_key.sshkey.id]
}

resource "ibm_is_instance" "instance3" {
  name    = "instance3"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }
  dedicated_host_group = ibm_is_dedicated_host_group.dh_group01.id
  vpc                  = ibm_is_vpc.vpc2.id
  zone                 = var.zone2
  keys                 = [ibm_is_ssh_key.sshkey.id]
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
  vpc  = ibm_is_vpc.vpc1.id
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
  name        = "isexample-rule"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"
  icmp {
    code = 1
    type = 1
  }
}

data "ibm_is_network_acl_rule" "testacc_dsnaclrule" {
  network_acl = ibm_is_network_acl.isExampleACL.id
  name        = ibm_is_network_acl_rule.isExampleACL.name
}

data "ibm_is_network_acl_rules" "testacc_dsnaclrules" {
  network_acl = ibm_is_network_acl.isExampleACL.id
}

data "ibm_is_network_acl" "is_network_acl" {
  network_acl = ibm_is_network_acl.isExampleACL.id
}

data "ibm_is_network_acl" "is_network_acl1" {
  name     = ibm_is_network_acl.isExampleACL.name
  vpc_name = ibm_is_vpc.vpc1.name
}

data "ibm_is_network_acls" "is_network_acls" {
}

resource "ibm_is_public_gateway" "publicgateway1" {
  name = "gateway1"
  vpc  = ibm_is_vpc.vpc1.id
  zone = var.zone1
}

// subnet public gateway attachment
resource "ibm_is_subnet_public_gateway_attachment" "example" {
  subnet         = ibm_is_subnet.subnet1.id
  public_gateway = ibm_is_public_gateway.publicgateway1.id
}

data "ibm_is_public_gateway" "testacc_dspgw" {
  name = ibm_is_public_gateway.publicgateway1.name
}

data "ibm_is_public_gateways" "publicgateways" {
}

data "ibm_is_vpc" "vpc1" {
  name = ibm_is_vpc.vpc1.name
}

// added for vpcs datasource
data "ibm_is_vpc" "vpcs" {
}

data "ibm_is_volume_profile" "volprofile" {
  name = "general-purpose"
}

data "ibm_is_volume_profiles" "volprofiles" {
}

data "ibm_resource_group" "default" {
  name = "Default" ///give your resource grp
}

resource "ibm_is_dedicated_host_group" "dh_group01" {
  family         = "balanced"
  class          = "bx2d"
  zone           = "us-south-1"
  name           = "my-dh-group-01"
  resource_group = data.ibm_resource_group.default.id
}
data "ibm_is_dedicated_host_group" "dgroup" {
  name = ibm_is_dedicated_host_group.dh_group01.name
}
resource "ibm_is_dedicated_host" "is_dedicated_host" {
  profile        = "bx2d-host-152x608"
  name           = "my-dedicated-host-01"
  host_group     = ibm_is_dedicated_host_group.dh_group01.id
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
  name       = ibm_is_dedicated_host.is_dedicated_host.name
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

  volumes = [ibm_is_volume.vol3.id]

  primary_network_interface {
    subnet = ibm_is_subnet.subnet1.id
  }

  vpc  = ibm_is_vpc.vpc1.id
  zone = var.zone1
  keys = [ibm_is_ssh_key.sshkey.id]
}

// creating a snapshot from boot volume with clone
resource "ibm_is_snapshot" "b_snapshot" {
  name          = "my-snapshot-boot"
  source_volume = ibm_is_instance.instance4.volume_attachments[0].volume_id
  clones        = [var.zone1]
  tags          = ["tags1"]
}

// creating a snapshot from data volume
resource "ibm_is_snapshot" "d_snapshot" {
  name          = "my-snapshot-data"
  source_volume = ibm_is_instance.instance4.volume_attachments[1].volume_id
  tags          = ["tags1"]
}

// data source for snapshot by name
data "ibm_is_snapshot" "ds_snapshot" {
  name = "my-snapshot-boot"
}

// data source for snapshots
data "ibm_is_snapshots" "ds_snapshots" {
}

// data source for snapshot clones
data "ibm_is_snapshot_clones" "ds_snapshot_clones" {
  snapshot = ibm_is_snapshot.b_snapshot.id
}

// data source for snapshot clones
data "ibm_is_snapshot_clones" "ds_snapshot_clone" {
  snapshot = ibm_is_snapshot.b_snapshot.id
  zone     = var.zone1
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
  tags    = ["tag1"]
}

// creating a volume attachment on an existing instance using an existing volume
resource "ibm_is_instance_volume_attachment" "att1" {
  instance                           = ibm_is_instance.instance5.id
  volume                             = ibm_is_volume.vol5.id
  name                               = "vol-att-1"
  delete_volume_on_attachment_delete = false
  delete_volume_on_instance_delete   = false
}

// creating a volume attachment on an existing instance using a new volume
resource "ibm_is_instance_volume_attachment" "att2" {
  instance                           = ibm_is_instance.instance5.id
  name                               = "vol-att-2"
  profile                            = "general-purpose"
  snapshot                           = ibm_is_snapshot.d_snapshot.id
  delete_volume_on_instance_delete   = true
  delete_volume_on_attachment_delete = true
  volume_name                        = "vol4-restore"
}

// data source for volume attachment
data "ibm_is_instance_volume_attachment" "ds_vol_att" {
  instance = ibm_is_instance.instance5.id
  name     = ibm_is_instance_volume_attachment.att2.name
}

// data source for volume attachments
data "ibm_is_instance_volume_attachment" "ds_vol_atts" {
  instance = ibm_is_instance.instance5.id
}

// creating an instance using an existing instance template
resource "ibm_is_instance" "instance6" {
  name              = "instance4"
  instance_template = ibm_is_instance_template.instancetemplate1.id
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

resource "ibm_is_instance_disk_management" "disks" {
  instance = ibm_is_instance.instance1.id
  disks {
    name = "mydisk01"
    id   = ibm_is_instance.instance1.disks.0.id
  }
}

data "ibm_is_instance_disks" "disk1" {
  instance = ibm_is_instance.instance1.id
}

// reserved ips

resource "ibm_is_instance" "instance7" {
  name    = "instance5"
  profile = var.profile
  boot_volume {
    name     = "boot-restore"
    snapshot = ibm_is_snapshot.b_snapshot.id
  }
  auto_delete_volume = true
  primary_network_interface {
    primary_ip {
      address     = "10.0.0.5"
      auto_delete = true
    }
    name   = "test-reserved-ip"
    subnet = ibm_is_subnet.subnet2.id
  }
  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}

// catalog images 

data "ibm_is_images" "imageslist" {
  catalog_managed = true
}
resource "ibm_is_instance" "instance8" {
  name               = "instance8"
  profile            = var.profile
  auto_delete_volume = true
  primary_network_interface {
    primary_ip {
      name        = "example-reserved-ip"
      auto_delete = true
    }
    name   = "test-reserved-ip"
    subnet = ibm_is_subnet.subnet2.id
  }
  catalog_offering {
    version_crn = data.ibm_is_images.imageslist.images.0.catalog_offering.0.version.0.crn
    plan_crn    = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
  }
  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}

resource "ibm_is_instance_template" "instancetemplate3" {
  name = "instancetemplate-3"
  catalog_offering {
    version_crn = data.ibm_is_images.imageslist.images.0.catalog_offering.0.version.0.crn
    plan_crn    = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
  }
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}


data "ibm_is_instance_network_interface_reserved_ip" "data_reserved_ip" {
  instance          = ibm_is_instance.test_instance.id
  network_interface = ibm_is_instance.test_instance.network_interfaces.0.id
  reserved_ip       = ibm_is_instance.test_instance.network_interfaces.0.ips.0.id
}

data "ibm_is_instance_network_interface_reserved_ips" "data_reserved_ips" {
  instance          = ibm_is_instance.test_instance.id
  network_interface = ibm_is_instance.test_instance.network_interfaces.0.id
}

data "ibm_is_instance_disk" "disk1" {
  instance = ibm_is_instance.instance1.id
  disk     = data.ibm_is_instance_disks.disk1.disks.0.id
}

data "ibm_is_dedicated_host_disks" "dhdisks" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
}

data "ibm_is_dedicated_host_disk" "dhdisk" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
  disk           = ibm_is_dedicated_host_disk_management.disks.disks.0.id
}

resource "ibm_is_dedicated_host_disk_management" "disks" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
  disks {
    name = "newdisk01"
    id   = data.ibm_is_dedicated_host.dhost.disks.0.id

  }
  disks {
    name = "newdisk02"
    id   = data.ibm_is_dedicated_host.dhost.disks.1.id

  }
}

data "ibm_is_operating_system" "os" {
  name = "red-8-amd64"
}

data "ibm_is_operating_systems" "oslist" {
}

#### BARE METAL SERVER


resource "ibm_is_bare_metal_server" "bms" {
  profile = "bx2-metal-192x768"
  name    = "my-bms"
  image   = "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e"
  zone    = "us-south-3"
  keys    = [ibm_is_ssh_key.sshkey.id]
  primary_network_interface {
    subnet = ibm_is_subnet.subnet1.id
  }
  vpc = ibm_is_vpc.vpc1.id
}

resource "ibm_is_bare_metal_server_disk" "this" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  disk              = ibm_is_bare_metal_server.bms.disks.0.id
  name              = "bms-disk-update"
}

resource "ibm_is_bare_metal_server_action" "this" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  action            = "stop"
  stop_type         = "hard"
}

data "ibm_is_bare_metal_server_profiles" "this" {
}

data "ibm_is_bare_metal_server_profile" "this" {
  name = data.ibm_is_bare_metal_server_profiles.this.profiles.0.name
}

data "ibm_is_bare_metal_server_disk" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
  disk              = ibm_is_bare_metal_server.this.disks.0.id
}

data "ibm_is_bare_metal_server_disks" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
}

resource "ibm_is_bare_metal_server_network_interface" "bms_nic" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id

  subnet            = ibm_is_subnet.subnet1.id
  name              = "eth2"
  allow_ip_spoofing = true
  allowed_vlans     = [101, 102]
}

resource "ibm_is_bare_metal_server_network_interface_allow_float" "bms_vlan_nic" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id

  subnet = ibm_is_subnet.subnet1.id
  name   = "eth2"
  vlan   = 102
}

resource "ibm_is_bare_metal_server_network_interface" "bms_nic2" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id

  subnet            = ibm_is_subnet.subnet1.id
  name              = "eth2"
  allow_ip_spoofing = true
  vlan              = 101
}

resource "ibm_is_floating_ip" "testacc_fip" {
  name = "testaccfip"
  zone = ibm_is_subnet.subnet1.zone
}

resource "ibm_is_bare_metal_server_network_interface_floating_ip" "bms_nic_fip" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  network_interface = ibm_is_bare_metal_server_network_interface.bms_nic2.id
  floating_ip       = ibm_is_floating_ip.testacc_fip.id
}

data "ibm_is_bare_metal_server_network_interface" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
  network_interface = ibm_is_bare_metal_server.this.primary_network_interface.id
}

data "ibm_is_bare_metal_server_network_interfaces" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
}

data "ibm_is_bare_metal_server" "this" {
  identifier = ibm_is_bare_metal_server.this.id
}

data "ibm_is_bare_metal_servers" "this" {
}

data "ibm_is_bare_metal_server_initialization" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
}

resource "ibm_is_floating_ip" "floatingipbms" {
  name = "fip1"
  zone = ibm_is_subnet.subnet1.zone
}

data "ibm_is_bare_metal_server_network_interface_floating_ip" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
  network_interface = ibm_is_bare_metal_server.this.primary_network_interface[0].id
  floating_ip       = ibm_is_floating_ip.floatingipbms.id
}

data "ibm_is_bare_metal_server_network_interface_floating_ips" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
  network_interface = ibm_is_bare_metal_server.this.primary_network_interface[0].id
}

resource "ibm_is_placement_group" "is_placement_group" {
  strategy       = "%s"
  name           = "%s"
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

data "ibm_is_vpc_address_prefix" "example" {
  vpc            = ibm_is_vpc.vpc1.id
  address_prefix = ibm_is_vpc_address_prefix.testacc_vpc_address_prefix.address_prefix
}

data "ibm_is_vpc_address_prefix" "example-1" {
  vpc_name       = ibm_is_vpc.vpc1.name
  address_prefix = ibm_is_vpc_address_prefix.testacc_vpc_address_prefix.address_prefix
}

data "ibm_is_vpc_address_prefix" "example-2" {
  vpc                 = ibm_is_vpc.vpc1.id
  address_prefix_name = ibm_is_vpc_address_prefix.testacc_vpc_address_prefix.name
}

data "ibm_is_vpc_address_prefix" "example-3" {
  vpc_name            = ibm_is_vpc.vpc1.name
  address_prefix_name = ibm_is_vpc_address_prefix.testacc_vpc_address_prefix.name
}

## Security Groups/Rules/Rule
// Create is_security_groups data source
data "ibm_is_security_groups" "example" {
}

// Create is_security_group data source
resource "ibm_is_security_group" "example" {
  name = "example-security-group"
  vpc  = ibm_is_vpc.vpc1.id
}

resource "ibm_is_security_group_rule" "exampleudp" {
  depends_on = [
    ibm_is_security_group.example,
  ]
  group     = ibm_is_security_group.example.id
  direction = "inbound"
  remote    = "127.0.0.1"
  udp {
    port_min = 805
    port_max = 807
  }
}

data "ibm_is_security_group_rule" "example" {
  depends_on = [
    ibm_is_security_group_rule.exampleudp,
  ]
  security_group_rule = ibm_is_security_group_rule.exampleudp.rule_id
  security_group      = ibm_is_security_group.example.id
}

// Create is_security_group_rules data source
resource "ibm_is_security_group_rule" "exampletcp" {
  group     = ibm_is_security_group.example.id
  direction = "outbound"
  remote    = "127.0.0.1"
  tcp {
    port_min = 8080
    port_max = 8080
  }
  depends_on = [
    ibm_is_security_group.example,
  ]
}

data "ibm_is_security_group_rules" "example" {
  depends_on = [
    ibm_is_security_group_rule.exampletcp,
  ]
}

data "ibm_is_vpn_gateway" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
}
data "ibm_is_vpn_gateway" "example-1" {
  vpn_gateway_name = ibm_is_vpn_gateway.example.name
}
data "ibm_is_vpn_gateway_connection" "example" {
  vpn_gateway            = ibm_is_vpn_gateway.example.id
  vpn_gateway_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
}
data "ibm_is_vpn_gateway_connection" "example-1" {
  vpn_gateway                 = ibm_is_vpn_gateway.example-1.id
  vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.example.name
}
data "ibm_is_vpn_gateway_connection" "example-2" {
  vpn_gateway_name       = ibm_is_vpn_gateway.example.name
  vpn_gateway_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
}
data "ibm_is_vpn_gateway_connection" "example-3" {
  vpn_gateway_name            = ibm_is_vpn_gateway.example.name
  vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.example.name
}
data "ibm_is_ike_policies" "example" {
}

data "ibm_is_ipsec_policies" "example" {
}

data "ibm_is_ike_policy" "example" {
  ike_policy = ibm_is_ike_policy.example.id
}

data "ibm_is_ipsec_policy" "example1" {
  ipsec_policy = ibm_is_ipsec_policy.example.id
}

data "ibm_is_ike_policy" "example2" {
  name = "my-ike-policy"
}

data "ibm_is_ipsec_policy" "example3" {
  name = "my-ipsec-policy"
}

# List ssh keys 
data "ibm_is_ssh_keys" "example" {
}

# List ssh keys by Resource group id
data "ibm_is_ssh_keys" "example" {
  resource_group = data.ibm_resource_group.default.id
}

# List volumes
data "ibm_is_volumes" "example-volumes" {
}

# List Volumes by Name 
data "ibm_is_volumes" "example" {
  volume_name = "worrier-mailable-timpani-scowling"
}

# List Volumes by Zone name
data "ibm_is_volumes" "example" {
  zone_name = "us-south-1"
}

## Backup Policy
resource "ibm_is_backup_policy" "is_backup_policy" {
  match_user_tags     = ["tag1"]
  name                = "my-backup-policy"
  match_resource_type = "volume"
}

resource "ibm_is_backup_policy" "is_backup_policy" {
  match_user_tags     = ["tag1"]
  name                = "my-backup-policy-instance"
  match_resource_type = "instance"
  included_content    = ["boot_volume", "data_volumes"]
}

resource "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
  backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
  cron_spec        = "30 09 * * *"
  active           = false
  attach_user_tags = ["tag2"]
  copy_user_tags   = true
  deletion_trigger {
    delete_after      = 20
    delete_over_count = "20"
  }
  name = "my-backup-policy-plan-1"
}
resource "ibm_is_backup_policy_plan" "is_backup_policy_plan_clone" {
  backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
  cron_spec        = "30 09 * * *"
  active           = false
  attach_user_tags = ["tag2"]
  copy_user_tags   = true
  deletion_trigger {
    delete_after      = 20
    delete_over_count = "20"
  }
  name = "my-backup-policy-plan-1"
  clone_policy {
    zones         = ["us-south-1", "us-south-2"]
    max_snapshots = 3
  }
}

data "ibm_is_backup_policies" "is_backup_policies" {
}

data "ibm_is_backup_policy" "is_backup_policy" {
  name = "my-backup-policy"
}

data "ibm_is_backup_policy_plans" "is_backup_policy_plans" {
  backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
}

data "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
  backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
  name             = "my-backup-policy-plan"
}

//backup policies for enterprise

resource "ibm_is_backup_policy" "ent-baas-example" {
  match_user_tags = ["tag1"]
  name            = "example-enterprise-backup-policy"
  scope {
    crn = "crn:v1:bluemix:public:is:us-south:a/123456::reservation:7187-ba49df72-37b8-43ac-98da-f8e029de0e63"
  }
}

data "ibm_is_backup_policy" "enterprise_backup" {
  name = ibm_is_backup_policy.ent-baas-example.name
}

// Vpn Server
resource "ibm_is_vpn_server" "is_vpn_server" {
  certificate_crn = var.is_certificate_crn
  client_authentication {
    method    = "certificate"
    client_ca = var.is_client_ca
  }
  client_ip_pool         = "10.5.0.0/21"
  subnets                = [ibm_is_subnet.subnet1.id]
  client_dns_server_ips  = ["192.168.3.4"]
  client_idle_timeout    = 2800
  enable_split_tunneling = false
  name                   = "example-vpn-server"
  port                   = 443
  protocol               = "udp"
}

resource "ibm_is_vpn_server_route" "is_vpn_server_route" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
  destination   = "172.16.0.0/16"
  action        = "translate"
  name          = "example-vpn-server-route"
}

data "ibm_is_backup_policy_job" "is_backup_policy_job" {
  backup_policy_id = ibm_is_backup_policy.is_backup_policy.id
  identifier       = ""
}

data "ibm_is_backup_policy_jobs" "is_backup_policy_jobs" {
  backup_policy_plan_id = ibm_is_backup_policy.is_backup_policy.backup_policy_plan_id
  backup_policy_id      = ibm_is_backup_policy.is_backup_policy.id
}

data "ibm_is_vpn_server" "is_vpn_server" {
  identifier = ibm_is_vpn_server.is_vpn_server.vpn_server
}
data "ibm_is_vpn_servers" "is_vpn_servers" {
}

data "ibm_is_vpn_server_routes" "is_vpn_server_routes" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
}

data "ibm_is_vpn_server_route" "is_vpn_server_route" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
  identifier    = ibm_is_vpn_server_route.is_vpn_server_route.vpn_route
}
data "ibm_is_vpn_server_clients" "is_vpn_server_clients" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
}
data "ibm_is_vpn_server_client" "is_vpn_server_client" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
  identifier    = "0726-61b2f53f-1e95-42a7-94ab-55de8f8cbdd5"
}
resource "ibm_is_image_export_job" "example" {
  image = ibm_is_image.image1.id
  name  = "my-image-export"
  storage_bucket {
    name = "bucket-27200-lwx4cfvcue"
  }
}

data "ibm_is_image_export_jobs" "example" {
  image = ibm_is_image_export_job.example.image
}

data "ibm_is_image_export_job" "example" {
  image            = ibm_is_image_export_job.example.image
  image_export_job = ibm_is_image_export_job.example.image_export_job
}
resource "ibm_is_vpc" "vpc" {
  name = "my-vpc"
}
resource "ibm_is_share" "share" {
  zone        = "us-south-1"
  size        = 30000
  name        = "my-share"
  profile     = "dp2"
  tags        = ["share1", "share3"]
  access_tags = ["access:dev"]
}

resource "ibm_is_share" "sharereplica" {
  zone                  = "us-south-2"
  name                  = "my-share-replica"
  profile               = "dp2"
  replication_cron_spec = "0 */5 * * *"
  source_share          = ibm_is_share.share.id
  tags                  = ["share1", "share3"]
  access_tags           = ["access:dev"]
}

resource "ibm_is_share_mount_target" "is_share_mount_target" {
  share = ibm_is_share.is_share.id
  vpc   = ibm_is_vpc.vpc1.id
  name  = "my-share-target-1"
}

data "ibm_is_share_mount_target" "is_share_mount_target" {
  share        = ibm_is_share.is_share.id
  mount_target = ibm_is_share_mount_target.is_share_target.mount_target
}

data "ibm_is_share_mount_targets" "is_share_mount_targets" {
  share = ibm_is_share.is_share.id
}

data "ibm_is_share" "is_share" {
  share = ibm_is_share.is_share.id
}

data "ibm_is_shares" "is_shares" {
}

resource "ibm_is_share_snapshot" "example" {
  name  = "my-example-share-snapshot"
  share = ibm_is_share.share.id
  tags  = ["my-example-share-snapshot-tag"]
}
data "ibm_is_share_snapshots" "example" {
  share = ibm_is_share.share.id
}

// Retrieve all the snapshots from all the shares
data "ibm_is_share_snapshots" "example1" {
}
data "ibm_is_share_snapshot" "example1" {
  share          = ibm_is_share.share.id
  share_snapshot = ibm_is_share_snapshot.example.share_snapshot
}
// vpc dns resolution bindings

// list all dns resolution bindings on a vpc
data "ibm_is_vpc_dns_resolution_bindings" "is_vpc_dns_resolution_bindings" {
  vpc_id = ibm_is_vpc.vpc1.id
}
// get a dns resolution bindings on a vpc
data "ibm_is_vpc_dns_resolution_binding" "is_vpc_dns_resolution_binding" {
  vpc_id = ibm_is_vpc.vpc1.id
  id     = ibm_is_vpc.vpc2.id
}
data "ibm_resource_group" "rg" {
  is_default = true
}
// creating a hub enabled vpc, hub disabled vpc, creating custom resolvers for both then
// delegating the vpc by uncommenting the configuration in hub_false_delegated vpc
resource "ibm_is_vpc" "hub_true" {
  name = "${var.name}-vpc-hub-true"
  dns {
    enable_hub = true
  }
}

resource "ibm_is_vpc" "hub_false_delegated" {
  name = "${var.name}-vpc-hub-false-del"
  dns {
    enable_hub = false
    # resolver {
    # 	type = "delegated"
    # 	vpc_id = ibm_is_vpc.hub_true.id
    # }
  }
}

resource "ibm_is_subnet" "hub_true_sub1" {
  name                     = "hub-true-subnet1"
  vpc                      = ibm_is_vpc.hub_true.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}
resource "ibm_is_subnet" "hub_true_sub2" {
  name                     = "hub-true-subnet2"
  vpc                      = ibm_is_vpc.hub_true.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}
resource "ibm_is_subnet" "hub_false_delegated_sub1" {
  name                     = "hub-false-delegated-subnet1"
  vpc                      = ibm_is_vpc.hub_false_delegated.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}
resource "ibm_is_subnet" "hub_false_delegated_sub2" {
  name                     = "hub-false-delegated-subnet2"
  vpc                      = ibm_is_vpc.hub_false_delegated.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}
resource "ibm_resource_instance" "dns-cr-instance" {
  name              = "dns-cr-instance"
  resource_group_id = data.ibm_resource_group.rg.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "standard-dns"
}
resource "ibm_dns_custom_resolver" "test_hub_true" {
  name              = "test-hub-true-customresolver"
  instance_id       = ibm_resource_instance.dns-cr-instance.guid
  description       = "new test CR - TF"
  high_availability = true
  enabled           = true
  locations {
    subnet_crn = ibm_is_subnet.hub_true_sub1.crn
    enabled    = true
  }
  locations {
    subnet_crn = ibm_is_subnet.hub_true_sub2.crn
    enabled    = true
  }
}
resource "ibm_dns_custom_resolver" "test_hub_false_delegated" {
  name              = "test-hub-false-customresolver"
  instance_id       = ibm_resource_instance.dns-cr-instance.guid
  description       = "new test CR - TF"
  high_availability = true
  enabled           = true
  locations {
    subnet_crn = ibm_is_subnet.hub_false_delegated_sub1.crn
    enabled    = true
  }
  locations {
    subnet_crn = ibm_is_subnet.hub_false_delegated_sub2.crn
    enabled    = true
  }
}

resource "ibm_is_vpc_dns_resolution_binding" "dnstrue" {
  name   = "hub-spoke-binding"
  vpc_id = ibm_is_vpc.hub_false_delegated.id
  vpc {
    id = ibm_is_vpc.hub_true.id
  }
}


// snapshot cross region

provider "ibm" {
  alias  = "eu-de"
  region = "eu-de"
}

resource "ibm_is_snapshot" "b_snapshot_copy" {
  provider            = ibm.eu-de
  name                = "my-snapshot-boot-copy"
  source_snapshot_crn = ibm_is_snapshot.b_snapshot.crn
}

// image deprecate and obsolete

resource "ibm_is_image_deprecate" "example" {
  image = ibm_is_image.image1.id
}

resource "ibm_is_image_obsolete" "example" {
  image = ibm_is_image.image1.id
}


// vni

resource "ibm_is_vpc" "testacc_vpc" {
  name = "${var.name}-vpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name                     = "${var.name}-subnet"
  vpc                      = ibm_is_vpc.testacc_vpc.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16

}

resource "ibm_is_virtual_network_interface" "testacc_vni" {
  name                      = var.name
  subnet                    = ibm_is_subnet.testacc_subnet.id
  enable_infrastructure_nat = true
  allow_ip_spoofing         = true
}

resource "ibm_is_floating_ip" "testacc_floatingip" {
  name = "${var.name}-floating"
  zone = ibm_is_subnet.testacc_subnet.zone
}
resource "ibm_is_virtual_network_interface_floating_ip" "testacc_vni_floatingip" {
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
  floating_ip               = ibm_is_floating_ip.testacc_floatingip.id
}
data "ibm_is_virtual_network_interface_floating_ip" "is_vni_floating_ip" {
  depends_on                = [ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip]
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
  floating_ip               = ibm_is_floating_ip.testacc_floatingip.id
}
data "ibm_is_virtual_network_interface_floating_ips" "is_vni_floating_ips" {
  depends_on                = [ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip]
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
}

data "ibm_is_virtual_network_interface_ips" "is_vni_reservedips" {
  depends_on                = [ibm_is_virtual_network_interface_ip.testacc_vni_reservedip]
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
}
data "ibm_is_virtual_network_interface_ip" "is_vni_reservedip" {
  depends_on                = [ibm_is_virtual_network_interface_ip.testacc_vni_reservedip]
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
  reserved_ip               = ibm_is_subnet_reserved_ip.testacc_reservedip.reserved_ip
}

resource "ibm_is_subnet_reserved_ip" "testacc_reservedip" {
  subnet = ibm_is_subnet.testacc_subnet.id
  name   = "${var.name}-reserved-ip"
}
resource "ibm_is_virtual_network_interface_ip" "testacc_vni_reservedip" {
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
  reserved_ip               = ibm_is_subnet_reserved_ip.testacc_reservedip.reserved_ip
}

resource "ibm_is_virtual_network_interface" "testacc_vni2" {
  name                      = "${var.name}-2"
  subnet                    = ibm_is_subnet.testacc_subnet.id
  enable_infrastructure_nat = true
  allow_ip_spoofing         = true
}
resource "ibm_is_virtual_network_interface" "testacc_vni3" {
  name                      = "${var.name}-3"
  subnet                    = ibm_is_subnet.testacc_subnet.id
  enable_infrastructure_nat = true
  allow_ip_spoofing         = true
}
resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "${var.name}-ssh"
  public_key = file("~/.ssh/id_rsa.pub")
}


resource "ibm_is_bare_metal_server" "testacc_bms" {
  profile = "cx2-metal-96x192"
  name    = "${var.name}-bms"
  image   = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"
  zone    = "${var.region}-2"
  keys    = [ibm_is_ssh_key.testacc_sshkey.id]
  primary_network_attachment {
    name = "vni-221"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.testacc_vni.id
    }
    allowed_vlans = [100, 102]
  }
  vpc = ibm_is_vpc.testacc_vpc.id
}

resource "ibm_is_bare_metal_server_network_attachment" "na" {
  bare_metal_server = ibm_is_bare_metal_server.testacc_bms.id
  # interface_type = "vlan"
  vlan = 100
}
resource "ibm_is_bare_metal_server_network_attachment" "na2" {
  bare_metal_server = ibm_is_bare_metal_server.testacc_bms.id
  # interface_type = "pci"
  allowed_vlans = [200, 202]
}

resource "ibm_is_instance_network_attachment" "ina" {
  instance = ibm_is_instance.ins.id
  name     = "viability-undecided-jalapeno-unbuilt"
  virtual_network_interface {
    id = ibm_is_virtual_network_interface.testacc_vni2.id
  }
}
resource "ibm_is_instance" "ins" {
  name    = "${var.name}-vsi2"
  profile = "bx2-2x8"
  image   = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"
  primary_network_attachment {
    name = "vni-test"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.testacc_vni3.id
    }
  }
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "${var.region}-2"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
}
resource "ibm_is_instance_template" "ins_temp" {
  name    = "${var.name}-vsi2"
  profile = "bx2-2x8"
  image   = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"
  primary_network_attachment {
    name = "vni-test"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.testacc_vni3.id
    }
  }
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "${var.region}-2"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
}

resource "ibm_is_share" "share" {
  zone                  = "us-east-1"
  source_share_crn      = "crn:v1:staging:public:is:us-south-1:a/efe5afc483594adaa8325e2b4d1290df::share:r134-d8c8821c-a227-451d-a9ed-0c0cd2358829"
  encryption_key        = "crn:v1:staging:public:kms:us-south:a/efe5afc483594adaa8325e2b4d1290df:1be45161-6dae-44ca-b248-837f98004057:key:3dd21cc5-cc20-4f7c-bc62-8ec9a8a3d1bd"
  replication_cron_spec = "5 * * * *"
  name                  = "tfp-temp-crr"
  profile               = "dp2"
}
//snapshot consistency group

resource "ibm_is_snapshot_consistency_group" "is_snapshot_consistency_group_instance" {
  delete_snapshots_on_delete = true
  snapshots {
    name          = "exmaple-snapshot"
    source_volume = ibm_is_instance.instance.volume_attachments[0].volume_id
  }
  snapshots {
    name          = "example-snapshot-1"
    source_volume = ibm_is_instance.instance.volume_attachments[1].volume_id
  }
  name = "example-snapshot-consistency-group"
}

data "ibm_is_snapshot_consistency_group" "is_snapshot_consistency_group_instance" {
  identifier = ibm_is_snapshot_consistency_group.is_snapshot_consistency_group_instance.id
}
data "ibm_is_snapshot_consistency_group" "is_snapshot_consistency_group_instance" {
  name = "example-snapshot-consistency-group"
}
data "ibm_is_snapshot_consistency_groups" "is_snapshot_consistency_group_instance" {
  depends_on = [ibm_is_snapshot_consistency_group.is_snapshot_consistency_group_instance]
  name       = "example-snapshot-consistency-group"
}

//reservation

resource "ibm_is_reservation" "example" {
  capacity {
    total = 5
  }
  committed_use {
    term = "one_year"
  }
  profile {
    name          = "ba2-2x8"
    resource_type = "instance_profile"
  }
  zone = "us-east-3"
}

resource "ibm_is_reservation_activate" "example" {
  reservation = ibm_is_reservation.example.id
}

data "ibm_is_reservations" "example" {
}

data "ibm_is_reservation" "example" {
  identifier = ibm_is_reservation.example.id
}

// cluster examples
# =============================================================================================================
variable "prefix" {
  default = "test-cluster"
}
variable "is_instances_resource_group_id" {
  default = "efhiorho4388yf348y83yvchrc083h0r30c"
}
variable "region" {
  default = "us-east"
}
variable "is_instances_name" {
  default = "test-vsi"
}
data "ibm_is_cluster_network_profile" "is_cluster_network_profile_instance" {
  name = "h100"
}
data "ibm_is_cluster_network_profiles" "is_cluster_network_profiles_instance" {
}
# Create VPC
resource "ibm_is_vpc" "is_vpc" {
  name = "${var.prefix}-vpc"
}
resource "ibm_is_vpc" "is_vpc2" {
  name = "${var.prefix}-vpc2"
}

# # # Create Subnet
resource "ibm_is_subnet" "is_subnet" {
  name                     = "${var.prefix}-subnet"
  vpc                      = ibm_is_vpc.is_vpc.id
  total_ipv4_address_count = 64
  zone                     = "${var.region}-3"
}

data "ibm_is_instance_profile" "is_instance_profile_instance" {
  name = "gx3d-160x1792x8h100"
}
data "ibm_is_instance_profiles" "is_instance_profiles_instance" {
}
data "ibm_is_image" "is_image" {
  name = "ibm-ubuntu-20-04-6-minimal-amd64-6"
}
resource "ibm_is_cluster_network" "is_cluster_network_instance" {
  name           = "${var.prefix}-cluster"
  profile        = "h100"
  resource_group = var.is_instances_resource_group_id
  subnet_prefixes {
    cidr = "10.1.0.0/24"
  }
  vpc {
    id = ibm_is_vpc.is_vpc.id
  }
  zone = "${var.region}-3"
}
resource "ibm_is_cluster_network" "is_cluster_network_instance" {
  name    = "${var.prefix}-cluster-updated"
  profile = "h100"
  subnet_prefixes {
    cidr = "10.0.0.0/24"
  }
  vpc {
    id = ibm_is_vpc.is_vpc.id
  }
  zone = ibm_is_subnet.is_subnet.zone
}
resource "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
  cluster_network_id       = ibm_is_cluster_network.is_cluster_network_instance.id
  name                     = "${var.prefix}-cluster-subnet"
  total_ipv4_address_count = 64
}

resource "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
  cluster_network_id        = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
  address                   = "10.1.0.4"
  name                      = "${var.prefix}-cluster-subnet-r-ip"
}

resource "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
  name               = "${var.prefix}-cluster-ni"
  primary_ip {
    id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_subnet_reserved_ip_id
  }
  subnet {
    id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
  }
}

resource "ibm_is_instance_template" "is_instance_template" {
  name    = "${var.prefix}-cluster-it"
  image   = data.ibm_is_image.is_image.id
  profile = "gx3d-160x1792x8h100"
  primary_network_attachment {
    name = "my-pna-it"
    virtual_network_interface {
      auto_delete = true
      subnet      = ibm_is_subnet.is_subnet.id
    }
  }
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
    }
  }
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
    }
  }
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
    }
  }
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
    }
  }
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
    }
  }
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
    }
  }
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
    }
  }
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
    }
  }
  vpc  = ibm_is_vpc.is_vpc.id
  zone = ibm_is_subnet.is_subnet.zone
  keys = [ibm_is_ssh_key.is_key.id]
}
resource "ibm_is_ssh_key" "is_key" {
  name       = "my-key"
  public_key = file("~/.ssh/id_ed25519.pub")
  type       = "ed25519"
}

resource "ibm_is_instance" "is_instance" {
  name    = "${var.prefix}-cluster-ins"
  image   = data.ibm_is_image.is_image.id
  profile = "gx3d-160x1792x8h100"
  primary_network_interface {
    subnet = ibm_is_subnet.is_subnet.id
  }
  cluster_network_attachments {
    name = "cna-1"
    cluster_network_interface {
      auto_delete = true
      name        = "cni-1"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-2"
    cluster_network_interface {
      auto_delete = true
      name        = "cni-2"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-3"
    cluster_network_interface {
      auto_delete = true
      name        = "cni-3"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-4"
    cluster_network_interface {
      auto_delete = true
      name        = "cni-4"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-5"
    cluster_network_interface {
      auto_delete = true
      name        = "cni-5"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-6"
    cluster_network_interface {
      auto_delete = true
      name        = "cni-6"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-7"
    cluster_network_interface {
      auto_delete = true
      name        = "cni-7"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  cluster_network_attachments {
    name = "cna-8"
    cluster_network_interface {
      auto_delete = true
      name        = "cni-8"
      subnet {
        id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
      }
    }
  }
  vpc  = ibm_is_vpc.is_vpc.id
  zone = ibm_is_subnet.is_subnet.zone
  keys = [ibm_is_ssh_key.is_key.id]
}

resource "ibm_is_instance_action" "is_instance_stop_before" {
  action   = "stop"
  instance = ibm_is_instance.is_instance.id
}

resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance" {
  depends_on  = [ibm_is_instance_action.is_instance_stop_before]
  instance_id = ibm_is_instance.is_instance.id
  before {
    id = ibm_is_instance.is_instance.cluster_network_attachments.0.id
  }
  cluster_network_interface {
    name = "my-cluster-network-interface"
    subnet {
      id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
    }
  }
  name = "cna-9"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance10" {
  depends_on  = [ibm_is_instance_action.is_instance_stop_before]
  instance_id = ibm_is_instance.is_instance.id
  before {
    id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance.instance_cluster_network_attachment_id
  }
  cluster_network_interface {
    name = "my-cluster-network-interface-10"
    subnet {
      id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
    }
  }
  name = "cna-10"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance11" {
  depends_on  = [ibm_is_instance_action.is_instance_stop_before]
  instance_id = ibm_is_instance.is_instance.id
  before {
    id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance10.instance_cluster_network_attachment_id
  }
  cluster_network_interface {
    name = "my-cluster-network-interface-11"
    subnet {
      id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
    }
  }
  name = "cna-11"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance12" {
  depends_on  = [ibm_is_instance_action.is_instance_stop_before]
  instance_id = ibm_is_instance.is_instance.id
  before {
    id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance11.instance_cluster_network_attachment_id
  }
  cluster_network_interface {
    name = "my-cluster-network-interface12"
    subnet {
      id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
    }
  }
  name = "cna-12"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance13" {
  depends_on  = [ibm_is_instance_action.is_instance_stop_before]
  instance_id = ibm_is_instance.is_instance.id
  before {
    id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance12.instance_cluster_network_attachment_id
  }
  cluster_network_interface {
    name = "my-cluster-network-interface13"
    subnet {
      id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
    }
  }
  name = "cna-13"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance14" {
  depends_on  = [ibm_is_instance_action.is_instance_stop_before]
  instance_id = ibm_is_instance.is_instance.id
  before {
    id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance13.instance_cluster_network_attachment_id
  }
  cluster_network_interface {
    name = "my-cluster-network-interface14"
    subnet {
      id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
    }
  }
  name = "cna-149"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance15" {
  depends_on  = [ibm_is_instance_action.is_instance_stop_before]
  instance_id = ibm_is_instance.is_instance.id
  before {
    id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance14.instance_cluster_network_attachment_id
  }
  cluster_network_interface {
    name = "my-cluster-network-interface15"
    subnet {
      id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
    }
  }
  name = "cna-15"
}
resource "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance16" {
  depends_on  = [ibm_is_instance_action.is_instance_stop_before]
  instance_id = ibm_is_instance.is_instance.id
  before {
    id = ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance15.instance_cluster_network_attachment_id
  }
  cluster_network_interface {
    name = "my-cluster-network-interface16"
    subnet {
      id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
    }
  }
  name = "cna-16"
}
resource "ibm_is_instance_action" "is_instance_start_after" {
  # depends_on = [ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance16]
  action   = "start"
  instance = ibm_is_instance.is_instance.id
}
resource "ibm_is_instance_action" "is_instance_stop_update" {
  # depends_on = [ibm_is_instance_cluster_network_attachment.is_instance_cluster_network_attachment_instance16]
  action   = "stop"
  instance = ibm_is_instance.is_instance.id
}

data "ibm_is_cluster_network" "is_cluster_network_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
}
data "ibm_is_cluster_networks" "is_cluster_networks_instance" {
}

data "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
  cluster_network_id           = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_interface_id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
}
data "ibm_is_cluster_network_interfaces" "is_cluster_network_interfaces_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
}

data "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
  cluster_network_id        = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
}
data "ibm_is_cluster_network_subnets" "is_cluster_network_subnets_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
}
data "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
  cluster_network_id                    = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id             = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
  cluster_network_subnet_reserved_ip_id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_subnet_reserved_ip_id

}
data "ibm_is_cluster_network_subnet_reserved_ips" "is_cluster_network_subnet_reserved_ips_instance" {
  cluster_network_id        = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
}

data "ibm_is_instance_template" "is_instance_template_instance" {
  name = ibm_is_instance_template.is_instance_template.name
}
data "ibm_is_instance" "is_instance_instance" {
  name = ibm_is_instance.is_instance.name
}
data "ibm_is_instances" "is_instances_instance" {
}
data "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance" {
  instance_id                            = ibm_is_instance.is_instance.id
  instance_cluster_network_attachment_id = ibm_is_instance.is_instance.cluster_network_attachments.0.id
}
data "ibm_is_instance_cluster_network_attachments" "is_instance_cluster_network_attachments_instance" {
  instance_id = ibm_is_instance.is_instance.id
}
data "ibm_is_instances" "is_instances_instance" {
  # resource_group_id = var.is_instances_resource_group_id
  name = var.is_instances_name
  # cluster_network_id = var.is_instances_cluster_network_id
  # cluster_network_crn = var.is_instances_cluster_network_crn
  # cluster_network_name = var.is_instances_cluster_network_name
  # dedicated_host_id = var.is_instances_dedicated_host_id
  # dedicated_host_crn = var.is_instances_dedicated_host_crn
  # dedicated_host_name = var.is_instances_dedicated_host_name
  # placement_group_id = var.is_instances_placement_group_id
  # placement_group_crn = var.is_instances_placement_group_crn
  # placement_group_name = var.is_instances_placement_group_name
  # reservation_id = var.is_instances_reservation_id
  # reservation_crn = var.is_instances_reservation_crn
  # reservation_name = var.is_instances_reservation_name
  # vpc_id = var.is_instances_vpc_id
  # vpc_crn = var.is_instances_vpc_crn
  # vpc_name = var.is_instances_vpc_name
}