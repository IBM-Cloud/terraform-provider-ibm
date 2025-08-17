# =====================================================================================
# IBM VPC Infrastructure
# =====================================================================================

# =====================================================================================
# Resource Groups
# =====================================================================================

data "ibm_resource_group" "default" {
  name = "Default" // Update with your resource group name
}

# =====================================================================================
# VPCs
# =====================================================================================

resource "ibm_is_vpc" "vpc-primary" {
  name = "vpc-primary"
}

resource "ibm_is_vpc" "vpc-secondary-1" {
  name = "vpc-secondary-1"
}

resource "ibm_is_vpc" "vpc-secondary-2" {
  name = "vpc-secondary-2"
}

// VPC with DNS hub configuration
resource "ibm_is_vpc" "vpc-hub-enabled" {
  name = "vpc-hub-enabled"
  dns {
    enable_hub = true
  }
}

// VPC with DNS hub disabled (can be delegated later)
resource "ibm_is_vpc" "vpc-hub-disabled" {
  name = "vpc-hub-disabled"
  dns {
    enable_hub = false
    # resolver {
    # 	type = "delegated"
    # 	vpc_id = ibm_is_vpc.vpc-hub-enabled.id
    # }
  }
}

# =====================================================================================
# VPC Address Prefixes
# =====================================================================================

resource "ibm_is_vpc_address_prefix" "address-prefix-1" {
  name       = "address-prefix-1"
  zone       = var.zone1
  vpc        = ibm_is_vpc.vpc-primary.id
  cidr       = var.cidr1
  is_default = false
}

# =====================================================================================
# Subnets
# =====================================================================================

resource "ibm_is_subnet" "subnet-primary-1" {
  name            = "subnet-primary-1"
  vpc             = ibm_is_vpc.vpc-primary.id
  zone            = var.zone1
  ipv4_cidr_block = "10.240.0.0/28"
}

resource "ibm_is_subnet" "subnet-secondary-1" {
  name            = "subnet-secondary-1"
  vpc             = ibm_is_vpc.vpc-secondary-1.id
  zone            = var.zone2
  ipv4_cidr_block = "10.240.64.0/28"
}

// Subnets for DNS hub configuration
resource "ibm_is_subnet" "subnet-hub-1" {
  name                     = "subnet-hub-1"
  vpc                      = ibm_is_vpc.vpc-hub-enabled.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}

resource "ibm_is_subnet" "subnet-hub-2" {
  name                     = "subnet-hub-2"
  vpc                      = ibm_is_vpc.vpc-hub-enabled.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}

resource "ibm_is_subnet" "subnet-delegated-1" {
  name                     = "subnet-delegated-1"
  vpc                      = ibm_is_vpc.vpc-hub-disabled.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}

resource "ibm_is_subnet" "subnet-delegated-2" {
  name                     = "subnet-delegated-2"
  vpc                      = ibm_is_vpc.vpc-hub-disabled.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}

// Additional subnet for VNI examples
resource "ibm_is_subnet" "subnet-vni-example" {
  name                     = "subnet-vni-example"
  vpc                      = ibm_is_vpc.vpc-secondary-2.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16
}

# =====================================================================================
# SSH Keys
# =====================================================================================

resource "ibm_is_ssh_key" "ssh-key-1" {
  name       = "ssh-key-1"
  public_key = file(var.ssh_public_key)
}

resource "ibm_is_ssh_key" "ssh-key-2" {
  name       = "ssh-key-2"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "ibm_is_ssh_key" "ssh-key-ed25519" {
  name       = "ssh-key-ed25519"
  public_key = file("~/.ssh/id_ed25519.pub")
  type       = "ed25519"
}

# =====================================================================================
# Security Groups
# =====================================================================================

resource "ibm_is_security_group" "security-group-1" {
  name = "security-group-1"
  vpc  = ibm_is_vpc.vpc-primary.id
}

# =====================================================================================
# Security Group Rules
# =====================================================================================

// SSH access rule
resource "ibm_is_security_group_rule" "sg-ssh-rule" {
  group     = ibm_is_vpc.vpc-primary.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 22
    port_max = 22
  }
}

// ICMP rule
resource "ibm_is_security_group_rule" "sg-icmp-rule" {
  group     = ibm_is_vpc.vpc-primary.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  icmp {
    code = 0
    type = 8
  }
}

// HTTP access rule
resource "ibm_is_security_group_rule" "sg-http-rule" {
  group     = ibm_is_vpc.vpc-primary.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 80
    port_max = 80
  }
}

// UDP rule example
resource "ibm_is_security_group_rule" "sg-udp-rule" {
  group     = ibm_is_security_group.security-group-1.id
  direction = "inbound"
  remote    = "127.0.0.1"

  udp {
    port_min = 805
    port_max = 807
  }
}

// TCP outbound rule
resource "ibm_is_security_group_rule" "sg-tcp-outbound" {
  group     = ibm_is_security_group.security-group-1.id
  direction = "outbound"
  remote    = "127.0.0.1"

  tcp {
    port_min = 8080
    port_max = 8080
  }
}

// Security group rules for VPC 2
resource "ibm_is_security_group_rule" "sg2-ssh-rule" {
  group     = ibm_is_vpc.vpc-secondary-1.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 22
    port_max = 22
  }
}

resource "ibm_is_security_group_rule" "sg2-icmp-rule" {
  group     = ibm_is_vpc.vpc-secondary-1.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  icmp {
    code = 0
    type = 8
  }
}

resource "ibm_is_security_group_rule" "sg2-http-rule" {
  group     = ibm_is_vpc.vpc-secondary-1.default_security_group
  direction = "inbound"
  remote    = "0.0.0.0/0"

  tcp {
    port_min = 80
    port_max = 80
  }
}

# =====================================================================================
# Network ACLs
# =====================================================================================

resource "ibm_is_network_acl" "network-acl-1" {
  name = "network-acl-1"
  vpc  = ibm_is_vpc.vpc-primary.id

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

# =====================================================================================
# Network ACL Rules
# =====================================================================================

resource "ibm_is_network_acl_rule" "network-acl-rule-1" {
  network_acl = ibm_is_network_acl.network-acl-1.id
  name        = "network-acl-rule-1"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"

  icmp {
    code = 1
    type = 1
  }
}

# =====================================================================================
# Reserved IPs
# =====================================================================================

resource "ibm_is_subnet_reserved_ip" "reserved-ip-1" {
  subnet = ibm_is_subnet.subnet-vni-example.id
  name   = "reserved-ip-1"
}

# =====================================================================================
# Virtual Network Interfaces (VNIs)
# =====================================================================================

resource "ibm_is_virtual_network_interface" "vni-1" {
  name                      = "vni-1"
  subnet                    = ibm_is_subnet.subnet-vni-example.id
  enable_infrastructure_nat = true
  allow_ip_spoofing         = true
}

resource "ibm_is_virtual_network_interface" "vni-2" {
  name                      = "vni-2"
  subnet                    = ibm_is_subnet.subnet-vni-example.id
  enable_infrastructure_nat = true
  allow_ip_spoofing         = true
}

resource "ibm_is_virtual_network_interface" "vni-3" {
  name                      = "vni-3"
  subnet                    = ibm_is_subnet.subnet-vni-example.id
  enable_infrastructure_nat = true
  allow_ip_spoofing         = true
}

# =====================================================================================
# VNI Reserved IP Attachments
# =====================================================================================

resource "ibm_is_virtual_network_interface_ip" "vni-reserved-ip-1" {
  virtual_network_interface = ibm_is_virtual_network_interface.vni-1.id
  reserved_ip               = ibm_is_subnet_reserved_ip.reserved-ip-1.reserved_ip
}

# =====================================================================================
# Data Sources for Foundation Resources
# =====================================================================================

// VPC data sources
data "ibm_is_vpc" "vpc-primary-data" {
  name = ibm_is_vpc.vpc-primary.name
}

data "ibm_is_vpcs" "all-vpcs" {
}

// VPC address prefix data sources
data "ibm_is_vpc_address_prefix" "address-prefix-data" {
  vpc            = ibm_is_vpc.vpc-primary.id
  address_prefix = ibm_is_vpc_address_prefix.address-prefix-1.address_prefix
}

data "ibm_is_vpc_address_prefix" "address-prefix-by-name" {
  vpc_name            = ibm_is_vpc.vpc-primary.name
  address_prefix_name = ibm_is_vpc_address_prefix.address-prefix-1.name
}

// Security group data sources
data "ibm_is_security_groups" "all-security-groups" {
}

data "ibm_is_security_group_rule" "sg-rule-data" {
  security_group_rule = ibm_is_security_group_rule.sg-udp-rule.rule_id
  security_group      = ibm_is_security_group.security-group-1.id
}

data "ibm_is_security_group_rules" "sg-rules-data" {
  security_group = ibm_is_security_group.security-group-1.id
}

// Network ACL data sources
data "ibm_is_network_acl" "network-acl-data" {
  network_acl = ibm_is_network_acl.network-acl-1.id
}

data "ibm_is_network_acl" "network-acl-by-name" {
  name     = ibm_is_network_acl.network-acl-1.name
  vpc_name = ibm_is_vpc.vpc-primary.name
}

data "ibm_is_network_acls" "all-network-acls" {
}

data "ibm_is_network_acl_rule" "network-acl-rule-data" {
  network_acl = ibm_is_network_acl.network-acl-1.id
  name        = ibm_is_network_acl_rule.network-acl-rule-1.name
}

data "ibm_is_network_acl_rules" "network-acl-rules-data" {
  network_acl = ibm_is_network_acl.network-acl-1.id
}

// SSH key data sources
data "ibm_is_ssh_keys" "all-ssh-keys" {
}

data "ibm_is_ssh_keys" "ssh-keys-by-resource-group" {
}

// VNI data sources
data "ibm_is_virtual_network_interface_ips" "vni-reserved-ips" {
  virtual_network_interface = ibm_is_virtual_network_interface.vni-1.id
}

data "ibm_is_virtual_network_interface_ip" "vni-reserved-ip-data" {
  virtual_network_interface = ibm_is_virtual_network_interface.vni-1.id
  reserved_ip               = ibm_is_subnet_reserved_ip.reserved-ip-1.reserved_ip
}

# =====================================================================================
# DNS Resolution Bindings
# =====================================================================================

// Create DNS resolution binding between VPCs
resource "ibm_is_vpc_dns_resolution_binding" "dns-binding-1" {
  name   = "dns-binding-1"
  vpc_id = ibm_is_vpc.vpc-hub-disabled.id
  vpc {
    id = ibm_is_vpc.vpc-hub-enabled.id
  }
}

// DNS resolution binding data sources
data "ibm_is_vpc_dns_resolution_bindings" "dns-bindings-data" {
  vpc_id = ibm_is_vpc.vpc-primary.id
}

data "ibm_is_vpc_dns_resolution_binding" "dns-binding-data" {
  vpc_id     = ibm_is_vpc.vpc-hub-disabled.id
  identifier = split("/", ibm_is_vpc_dns_resolution_binding.dns-binding-1.id)[1]
}

# =====================================================================================
# Volumes
# =====================================================================================

resource "ibm_is_volume" "volume-1" {
  name    = "volume-1"
  profile = "10iops-tier"
  zone    = var.zone1
}

resource "ibm_is_volume" "volume-2" {
  name     = "volume-2"
  profile  = "custom"
  zone     = var.zone1
  iops     = 1000
  capacity = 200
}

resource "ibm_is_volume" "volume-3" {
  name    = "volume-3"
  profile = "10iops-tier"
  zone    = var.zone1
  tags    = ["tag1"]
}

resource "ibm_is_volume" "volume-4" {
  name    = "volume-4"
  profile = "10iops-tier"
  zone    = "us-south-2"
  tags    = ["tag1"]
}

# =====================================================================================
# Virtual Server Instances (VSIs)
# =====================================================================================

// Basic instance
resource "ibm_is_instance" "vsi-1" {
  name    = "vsi-1"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-primary-1.id
  }

  vpc  = ibm_is_vpc.vpc-primary.id
  zone = var.zone1
  keys = [ibm_is_ssh_key.ssh-key-1.id]
}


// Instance with dedicated host
resource "ibm_is_instance" "vsi-2" {
  name    = "vsi-2"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-secondary-1.id
  }

  dedicated_host = ibm_is_dedicated_host.dedicated-host-1.id
  vpc            = ibm_is_vpc.vpc-secondary-1.id
  zone           = var.zone2
  keys           = [ibm_is_ssh_key.ssh-key-1.id]
}

// Instance with dedicated host group
resource "ibm_is_instance" "vsi-3" {
  name    = "vsi-3"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-secondary-1.id
  }

  dedicated_host_group = ibm_is_dedicated_host_group.host-group-1.id
  vpc                  = ibm_is_vpc.vpc-secondary-1.id
  zone                 = var.zone2
  keys                 = [ibm_is_ssh_key.ssh-key-1.id]
}

// Instance with volumes attached
resource "ibm_is_instance" "vsi-4" {
  name    = "vsi-4"
  image   = var.image
  profile = var.profile

  volumes = [ibm_is_volume.volume-3.id]

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-primary-1.id
  }

  vpc  = ibm_is_vpc.vpc-primary.id
  zone = var.zone1
  keys = [ibm_is_ssh_key.ssh-key-1.id]
}

// Instance with boot volume from snapshot
resource "ibm_is_instance" "vsi-5" {
  name    = "vsi-5"
  profile = var.profile

  boot_volume {
    name     = "boot-volume-restored"
    snapshot = ibm_is_snapshot.snapshot-boot.id
  }

  auto_delete_volume = true

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-secondary-1.id
  }

  vpc  = ibm_is_vpc.vpc-secondary-1.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.ssh-key-1.id]
}

// Instance with reserved IP
resource "ibm_is_instance" "vsi-6" {
  name    = "vsi-6"
  profile = var.profile

  boot_volume {
    name     = "boot-volume-reserved-ip"
    snapshot = ibm_is_snapshot.snapshot-boot.id
  }

  auto_delete_volume = true

  primary_network_interface {
    primary_ip {
      address     = "10.0.0.5"
      auto_delete = true
    }
    name   = "network-interface-reserved-ip"
    subnet = ibm_is_subnet.subnet-secondary-1.id
  }

  vpc  = ibm_is_vpc.vpc-secondary-1.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.ssh-key-1.id]
}

// Instance with named reserved IP
resource "ibm_is_instance" "vsi-7" {
  name               = "vsi-7"
  profile            = var.profile
  auto_delete_volume = true

  primary_network_interface {
    primary_ip {
      name        = "reserved-ip-named"
      auto_delete = true
    }
    name   = "network-interface-named-ip"
    subnet = ibm_is_subnet.subnet-secondary-1.id
  }

  catalog_offering {
    version_crn = data.ibm_is_images.catalog-images.images.0.catalog_offering.0.version.0.crn
    plan_crn    = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
  }

  vpc  = ibm_is_vpc.vpc-secondary-1.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.ssh-key-1.id]
}

// Instance using VNI (Virtual Network Interface)
resource "ibm_is_instance" "vsi-8" {
  name    = "vsi-8"
  profile = "bx2-2x8"
  image   = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"

  primary_network_attachment {
    name = "vni-attachment-1"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.vni-3.id
    }
  }

  vpc  = ibm_is_vpc.vpc-secondary-2.id
  zone = "${var.region}-2"
  keys = [ibm_is_ssh_key.ssh-key-ed25519.id]
}

# =====================================================================================
# Instance Network Interfaces
# =====================================================================================

resource "ibm_is_instance_network_interface" "network-interface-1" {
  instance             = ibm_is_instance.vsi-1.id
  subnet               = ibm_is_subnet.subnet-primary-1.id
  allow_ip_spoofing    = true
  name                 = "network-interface-1"
  primary_ipv4_address = "10.0.0.5"
}

// Instance network attachment using VNI
resource "ibm_is_instance_network_attachment" "network-attachment-1" {
  instance = ibm_is_instance.vsi-8.id
  name     = "network-attachment-1"
  virtual_network_interface {
    id = ibm_is_virtual_network_interface.vni-2.id
  }
}

# =====================================================================================
# Floating IPs
# =====================================================================================

resource "ibm_is_floating_ip" "floating-ip-1" {
  name   = "floating-ip-1"
  target = ibm_is_instance.vsi-1.primary_network_interface[0].id
}

resource "ibm_is_floating_ip" "floating-ip-2" {
  name   = "floating-ip-2"
  target = ibm_is_instance.vsi-2.primary_network_interface[0].id
}

resource "ibm_is_floating_ip" "floating-ip-3" {
  name = "floating-ip-3"
  zone = ibm_is_subnet.subnet-vni-example.zone
}

// VNI floating IP attachment
resource "ibm_is_virtual_network_interface_floating_ip" "vni-floating-ip-1" {
  virtual_network_interface = ibm_is_virtual_network_interface.vni-1.id
  floating_ip               = ibm_is_floating_ip.floating-ip-3.id
}

# =====================================================================================
# Volume Attachments
# =====================================================================================

// Volume attachment using existing volume
resource "ibm_is_instance_volume_attachment" "volume-attachment-1" {
  instance                           = ibm_is_instance.vsi-5.id
  volume                             = ibm_is_volume.volume-4.id
  name                               = "volume-attachment-1"
  delete_volume_on_attachment_delete = false
  delete_volume_on_instance_delete   = false
}

// Volume attachment creating new volume from snapshot
resource "ibm_is_instance_volume_attachment" "volume-attachment-2" {
  instance                           = ibm_is_instance.vsi-5.id
  name                               = "volume-attachment-2"
  profile                            = "general-purpose"
  snapshot                           = ibm_is_snapshot.snapshot-data.id
  delete_volume_on_instance_delete   = true
  delete_volume_on_attachment_delete = true
  volume_name                        = "volume-from-snapshot"
}

# =====================================================================================
# Instance Actions
# =====================================================================================

resource "ibm_is_instance_action" "vsi-stop-action" {
  action   = "stop"
  instance = ibm_is_instance.vsi-8.id
}

resource "ibm_is_instance_action" "vsi-start-action" {
  action   = "start"
  instance = ibm_is_instance.vsi-8.id
}

# =====================================================================================
# Instance Disk Management
# =====================================================================================

resource "ibm_is_instance_disk_management" "disk-management-1" {
  instance = ibm_is_instance.vsi-1.id
  disks {
    name = "disk-1"
    id   = ibm_is_instance.vsi-1.disks.0.id
  }
}

# =====================================================================================
# Data Sources for Compute Resources
# =====================================================================================

// Instance data sources
data "ibm_is_instance" "vsi-1-data" {
  name        = ibm_is_instance.vsi-1.name
  private_key = file("~/.ssh/id_rsa")
  passphrase  = ""
}

data "ibm_is_instances" "all-instances" {
}

// Instance network interface data sources
data "ibm_is_instance_network_interface" "network-interface-data" {
  instance_name          = ibm_is_instance.vsi-1.name
  network_interface_name = ibm_is_instance_network_interface.network-interface-1.name
}

data "ibm_is_instance_network_interfaces" "all-network-interfaces" {
  instance_name = ibm_is_instance.vsi-1.name
}

// Instance network interface reserved IP data sources
data "ibm_is_instance_network_interface_reserved_ip" "reserved-ip-data" {
  instance          = ibm_is_instance.vsi-6.id
  network_interface = ibm_is_instance.vsi-6.network_interfaces.0.id
  reserved_ip       = ibm_is_instance.vsi-6.network_interfaces.0.primary_ip.0.reserved_ip
}

data "ibm_is_instance_network_interface_reserved_ips" "reserved-ips-data" {
  instance          = ibm_is_instance.vsi-6.id
  network_interface = ibm_is_instance.vsi-6.network_interfaces.0.id
}

// Volume data sources
data "ibm_is_volumes" "all-volumes" {
}

data "ibm_is_volumes" "volumes-by-name" {
  volume_name = "volume-1"
}

data "ibm_is_volumes" "volumes-by-zone" {
  zone_name = "us-south-1"
}

data "ibm_is_volume_profile" "volume-profile-data" {
  name = "general-purpose"
}

data "ibm_is_volume_profiles" "all-volume-profiles" {
}

// Volume attachment data sources
data "ibm_is_instance_volume_attachment" "volume-attachment-data" {
  instance = ibm_is_instance.vsi-5.id
  name     = ibm_is_instance_volume_attachment.volume-attachment-2.name
}

data "ibm_is_instance_volume_attachments" "all-volume-attachments" {
  instance = ibm_is_instance.vsi-5.id
}

// Instance disk data sources
data "ibm_is_instance_disks" "instance-disks" {
  instance = ibm_is_instance.vsi-1.id
}

data "ibm_is_instance_disk" "instance-disk-data" {
  instance = ibm_is_instance.vsi-1.id
  disk     = data.ibm_is_instance_disks.instance-disks.disks.0.id
}

// VNI floating IP data sources
data "ibm_is_virtual_network_interface_floating_ip" "vni-floating-ip-data" {
  virtual_network_interface = ibm_is_virtual_network_interface.vni-1.id
  floating_ip               = ibm_is_floating_ip.floating-ip-3.id
}

data "ibm_is_virtual_network_interface_floating_ips" "vni-floating-ips-data" {
  virtual_network_interface = ibm_is_virtual_network_interface.vni-1.id
}

// Image data sources
data "ibm_is_image" "image-data" {
  name = "my-custom-image"
}

data "ibm_is_images" "all-images" {
}

data "ibm_is_images" "catalog-images" {
  catalog_managed = true
}

// Operating system data sources
data "ibm_is_operating_system" "os-data" {
  name = "red-8-amd64"
}

data "ibm_is_operating_systems" "all-operating-systems" {
}

// Regions data source
data "ibm_is_regions" "all-regions" {
}

# =====================================================================================
# Public Gateways
# =====================================================================================

resource "ibm_is_public_gateway" "public-gateway-1" {
  name = "public-gateway-1"
  vpc  = ibm_is_vpc.vpc-primary.id
  zone = var.zone1
}

// Subnet public gateway attachment
resource "ibm_is_subnet_public_gateway_attachment" "subnet-pgw-attachment-1" {
  subnet         = ibm_is_subnet.subnet-primary-1.id
  public_gateway = ibm_is_public_gateway.public-gateway-1.id
}

# =====================================================================================
# IKE and IPSec Policies
# =====================================================================================

resource "ibm_is_ike_policy" "ike-policy-1" {
  name                     = "ike-policy-1"
  authentication_algorithm = "sha256"
  encryption_algorithm     = "aes128"
  dh_group                 = 14
  ike_version              = 1
}

resource "ibm_is_ipsec_policy" "ipsec-policy-1" {
  name                     = "ipsec-policy-1"
  authentication_algorithm = "sha256"
  encryption_algorithm     = "aes128"
  pfs                      = "disabled"
}

# =====================================================================================
# VPN Gateways
# =====================================================================================

resource "ibm_is_vpn_gateway" "vpn-gateway-1" {
  name   = "vpn-gateway-1"
  subnet = ibm_is_subnet.subnet-primary-1.id
}

resource "ibm_is_vpn_gateway" "vpn-gateway-2" {
  name   = "vpn-gateway-2"
  subnet = ibm_is_subnet.subnet-secondary-1.id
}

# =====================================================================================
# VPN Gateway Connections
# =====================================================================================

// Deprecated format for reference
resource "ibm_is_vpn_gateway_connection" "vpn-connection-1-deprecated" {
  name          = "vpn-connection-1-deprecated"
  vpn_gateway   = ibm_is_vpn_gateway.vpn-gateway-1.id
  peer_address  = ibm_is_vpn_gateway.vpn-gateway-1.public_ip_address
  preshared_key = "VPNDemoPassword"
  local_cidrs   = [ibm_is_subnet.subnet-primary-1.ipv4_cidr_block]
  peer_cidrs    = [ibm_is_subnet.subnet-secondary-1.ipv4_cidr_block]
  ipsec_policy  = ibm_is_ipsec_policy.ipsec-policy-1.id
}

// Current format
resource "ibm_is_vpn_gateway_connection" "vpn-connection-1" {
  name          = "vpn-connection-1"
  vpn_gateway   = ibm_is_vpn_gateway.vpn-gateway-1.id
  preshared_key = "VPNDemoPassword"

  peer {
    address = ibm_is_vpn_gateway.vpn-gateway-2.public_ip_address != "0.0.0.0" ? ibm_is_vpn_gateway.vpn-gateway-2.public_ip_address : ibm_is_vpn_gateway.vpn-gateway-2.public_ip_address2
    cidrs   = [ibm_is_subnet.subnet-secondary-1.ipv4_cidr_block]
  }

  local {
    cidrs = [ibm_is_subnet.subnet-primary-1.ipv4_cidr_block]
  }

  ipsec_policy = ibm_is_ipsec_policy.ipsec-policy-1.id
}

resource "ibm_is_vpn_gateway_connection" "vpn-connection-2" {
  name           = "vpn-connection-2-deprecated"
  vpn_gateway    = ibm_is_vpn_gateway.vpn-gateway-2.id
  peer_address   = ibm_is_vpn_gateway.vpn-gateway-2.public_ip_address
  preshared_key  = "VPNDemoPassword"
  local_cidrs    = [ibm_is_subnet.subnet-secondary-1.ipv4_cidr_block]
  peer_cidrs     = [ibm_is_subnet.subnet-primary-1.ipv4_cidr_block]
  admin_state_up = true
  ike_policy     = ibm_is_ike_policy.ike-policy-1.id
}

# =====================================================================================
# Load Balancers
# =====================================================================================

// Load balancer with private DNS
resource "ibm_is_lb" "load-balancer-1" {
  name    = "load-balancer-1"
  subnets = [ibm_is_subnet.subnet-primary-1.id]
  profile = "network-fixed"
  dns {
    instance_crn = "crn:v1:staging:public:dns-svcs:global:a/exxxxxxxxxxxxx-xxxxxxxxxxxxxxxxx:5xxxxxxx-xxxxx-xxxxxxxxxxxxxxx-xxxxxxxxxxxxxxx::"
    zone_id      = "bxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx"
  }
}

// Basic load balancer
resource "ibm_is_lb" "load-balancer-2" {
  name    = "load-balancer-2"
  subnets = [ibm_is_subnet.subnet-primary-1.id]
}

# =====================================================================================
# Load Balancer Listeners
# =====================================================================================

resource "ibm_is_lb_listener" "lb-listener-1" {
  lb       = ibm_is_lb.load-balancer-2.id
  port     = "9086"
  protocol = "http"
}

# =====================================================================================
# Load Balancer Listener Policies
# =====================================================================================

resource "ibm_is_lb_listener_policy" "lb-listener-policy-1" {
  lb                      = ibm_is_lb.load-balancer-2.id
  listener                = ibm_is_lb_listener.lb-listener-1.listener_id
  action                  = "redirect"
  priority                = 2
  name                    = "lb-listener-policy-1"
  target_http_status_code = 302
  target_url              = "https://www.google.com"

  rules {
    condition = "contains"
    type      = "header"
    field     = "1"
    value     = "2"
  }
}

# =====================================================================================
# Load Balancer Listener Policy Rules
# =====================================================================================

resource "ibm_is_lb_listener_policy_rule" "lb-listener-policy-rule-1" {
  lb        = ibm_is_lb.load-balancer-2.id
  listener  = ibm_is_lb_listener.lb-listener-1.listener_id
  policy    = ibm_is_lb_listener_policy.lb-listener-policy-1.policy_id
  condition = "equals"
  type      = "header"
  field     = "MY-APP-HEADER"
  value     = "UpdateVal"
}

# =====================================================================================
# Load Balancer Pools
# =====================================================================================

// Pool with app cookie session persistence
resource "ibm_is_lb_pool" "lb-pool-1" {
  name                                = "lb-pool-1"
  lb                                  = ibm_is_lb.load-balancer-2.id
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

// Pool with HTTP cookie session persistence
resource "ibm_is_lb_pool" "lb-pool-2" {
  name                     = "lb-pool-2"
  lb                       = ibm_is_lb.load-balancer-2.id
  algorithm                = "round_robin"
  protocol                 = "https"
  health_delay             = 60
  health_retries           = 5
  health_timeout           = 30
  health_type              = "https"
  proxy_protocol           = "v1"
  session_persistence_type = "http_cookie"
}

// Pool with source IP session persistence
resource "ibm_is_lb_pool" "lb-pool-3" {
  name                     = "lb-pool-3"
  lb                       = ibm_is_lb.load-balancer-2.id
  algorithm                = "round_robin"
  protocol                 = "https"
  health_delay             = 60
  health_retries           = 5
  health_timeout           = 30
  health_type              = "https"
  proxy_protocol           = "v1"
  session_persistence_type = "source_ip"
}

# =====================================================================================
# Data Sources for Advanced Networking
# =====================================================================================

// Public gateway data sources
data "ibm_is_public_gateway" "public-gateway-data" {
  name = ibm_is_public_gateway.public-gateway-1.name
}

data "ibm_is_public_gateways" "all-public-gateways" {
}

// VPN data sources
data "ibm_is_vpn_gateway" "vpn-gateway-data" {
  vpn_gateway = ibm_is_vpn_gateway.vpn-gateway-1.id
}

data "ibm_is_vpn_gateway" "vpn-gateway-by-name" {
  vpn_gateway_name = ibm_is_vpn_gateway.vpn-gateway-1.name
}

data "ibm_is_vpn_gateway_connection" "vpn-connection-data" {
  vpn_gateway            = ibm_is_vpn_gateway.vpn-gateway-2.id
  vpn_gateway_connection = ibm_is_vpn_gateway_connection.vpn-connection-1.gateway_connection
}

data "ibm_is_vpn_gateway_connection" "vpn-connection-by-name" {
  vpn_gateway                 = ibm_is_vpn_gateway.vpn-gateway-2.id
  vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.vpn-connection-1.name
}

data "ibm_is_vpn_gateway_connection" "vpn-connection-by-gateway-name" {
  vpn_gateway_name       = ibm_is_vpn_gateway.vpn-gateway-1.name
  vpn_gateway_connection = ibm_is_vpn_gateway_connection.vpn-connection-1.gateway_connection
}

data "ibm_is_vpn_gateway_connection" "vpn-connection-by-both-names" {
  vpn_gateway_name            = ibm_is_vpn_gateway.vpn-gateway-1.name
  vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.vpn-connection-1.name
}

// IKE and IPSec policy data sources
data "ibm_is_ike_policies" "all-ike-policies" {
}

data "ibm_is_ipsec_policies" "all-ipsec-policies" {
}

data "ibm_is_ike_policy" "ike-policy-data" {
  ike_policy = ibm_is_ike_policy.ike-policy-1.id
}

data "ibm_is_ipsec_policy" "ipsec-policy-data" {
  ipsec_policy = ibm_is_ipsec_policy.ipsec-policy-1.id
}

data "ibm_is_ike_policy" "ike-policy-by-name" {
  name = "my-ike-policy"
}

data "ibm_is_ipsec_policy" "ipsec-policy-by-name" {
  name = "my-ipsec-policy"
}

// Load balancer data sources
data "ibm_is_lb_listener" "lb-listener-data" {
  lb          = ibm_is_lb.load-balancer-2.id
  listener_id = ibm_is_lb_listener.lb-listener-1.listener_id
}

data "ibm_is_lb_listeners" "lb-listeners-data" {
  lb = ibm_is_lb.load-balancer-2.id
}

data "ibm_is_lb_listener_policy" "lb-listener-policy-data" {
  lb        = ibm_is_lb.load-balancer-2.id
  listener  = ibm_is_lb_listener.lb-listener-1.listener_id
  policy_id = ibm_is_lb_listener_policy.lb-listener-policy-1.policy_id
}

data "ibm_is_lb_listener_policies" "lb-listener-policies-data" {
  lb       = ibm_is_lb.load-balancer-2.id
  listener = ibm_is_lb_listener.lb-listener-1.listener_id
}

data "ibm_is_lb_listener_policy_rule" "lb-listener-policy-rule-data" {
  lb       = ibm_is_lb.load-balancer-2.id
  listener = ibm_is_lb_listener.lb-listener-1.listener_id
  policy   = ibm_is_lb_listener_policy.lb-listener-policy-1.policy_id
  rule     = ibm_is_lb_listener_policy_rule.lb-listener-policy-rule-1.rule
}

data "ibm_is_lb_listener_policy_rules" "lb-listener-policy-rules-data" {
  lb       = ibm_is_lb.load-balancer-2.id
  listener = ibm_is_lb_listener.lb-listener-1.listener_id
  policy   = ibm_is_lb_listener_policy.lb-listener-policy-1.policy_id
}

# =====================================================================================
# Snapshots
# =====================================================================================

// Creating a snapshot from boot volume with clone
resource "ibm_is_snapshot" "snapshot-boot" {
  name          = "snapshot-boot"
  source_volume = ibm_is_instance.vsi-4.volume_attachments[0].volume_id
  clones        = [var.zone1]
  tags          = ["boot-snapshot"]
}

// Creating a snapshot from data volume
resource "ibm_is_snapshot" "snapshot-data" {
  name          = "snapshot-data"
  source_volume = ibm_is_instance.vsi-4.volume_attachments[1].volume_id
  tags          = ["data-snapshot"]
}

// Additional snapshots for examples
resource "ibm_is_snapshot" "snapshot-1" {
  name          = "snapshot-1"
  source_volume = ibm_is_volume.volume-1.id
  tags          = ["volume-snapshot"]
}

resource "ibm_is_snapshot" "snapshot-2" {
  name          = "snapshot-2"
  source_volume = ibm_is_volume.volume-2.id
  tags          = ["custom-volume-snapshot"]
}

# =====================================================================================
# Cross-Region Snapshot Copy
# =====================================================================================

// Provider for different region
provider "ibm" {
  alias  = "eu-de"
  region = "eu-de"
}

// Cross-region snapshot copy
resource "ibm_is_snapshot" "snapshot-cross-region" {
  provider            = ibm.eu-de
  name                = "snapshot-cross-region"
  source_snapshot_crn = ibm_is_snapshot.snapshot-boot.crn
}

# =====================================================================================
# Snapshot Consistency Groups
# =====================================================================================

resource "ibm_is_snapshot_consistency_group" "snapshot-consistency-group-1" {
  delete_snapshots_on_delete = true

  snapshots {
    name          = "snapshot-consistency-1"
    source_volume = ibm_is_instance.vsi-4.volume_attachments[0].volume_id
  }

  snapshots {
    name          = "snapshot-consistency-2"
    source_volume = ibm_is_instance.vsi-4.volume_attachments[1].volume_id
  }

  name = "snapshot-consistency-group-1"
}

# =====================================================================================
# Custom Images
# =====================================================================================

// Image from COS URL
resource "ibm_is_image" "image-1" {
  href             = var.image_cos_url
  name             = "image-1"
  operating_system = var.image_operating_system
}

// Image from instance volume
resource "ibm_is_image" "image-2" {
  source_volume = data.ibm_is_instance.vsi-1-data.volume_attachments.0.volume_id
  name          = "image-2"
}

# =====================================================================================
# Image Management
# =====================================================================================

// Image deprecation
resource "ibm_is_image_deprecate" "image-deprecate-1" {
  image = ibm_is_image.image-1.id
}

// Image obsolescence
resource "ibm_is_image_obsolete" "image-obsolete-1" {
  image = ibm_is_image.image-1.id
}

# =====================================================================================
# Image Export Jobs
# =====================================================================================

resource "ibm_is_image_export_job" "image-export-1" {
  image = ibm_is_image.image-1.id
  name  = "image-export-1"
  storage_bucket {
    name = "bucket-27200-lwx4cfvcue"
  }
}

# =====================================================================================
# File Shares
# =====================================================================================

// Basic file share
resource "ibm_is_share" "share-1" {
  zone        = "us-south-1"
  size        = 30000
  name        = "share-1"
  profile     = "dp2"
  tags        = ["share1", "share3"]
  access_tags = ["access:dev"]
}

// Replica share with replication
resource "ibm_is_share" "share-replica-1" {
  zone                  = "us-south-2"
  name                  = "share-replica-1"
  profile               = "dp2"
  replication_cron_spec = "0 * /5 * * *"
  source_share          = ibm_is_share.share-1.id
  tags                  = ["share1", "share3"]
  access_tags           = ["access:dev"]
}

// Share from source share (cross-region replication)
resource "ibm_is_share" "share-cross-region" {
  zone                  = "us-east-1"
  source_share_crn      = "crn:v1:staging:public:is:us-south-1:a/efe5afc483594adaa8325e2b4d1290df::share:r134-d8c8821c-a227-451d-a9ed-0c0cd2358829"
  encryption_key        = "crn:v1:staging:public:kms:us-south:a/efe5afc483594adaa8325e2b4d1290df:1be45161-6dae-44ca-b248-837f98004057:key:3dd21cc5-cc20-4f7c-bc62-8ec9a8a3d1bd"
  replication_cron_spec = "5 * * * *"
  name                  = "share-cross-region"
  profile               = "dp2"
}

# =====================================================================================
# Share Mount Targets
# =====================================================================================

resource "ibm_is_share_mount_target" "share-mount-target-1" {
  share = ibm_is_share.share-1.id
  vpc   = ibm_is_vpc.vpc-primary.id
  name  = "share-mount-target-1"
}

# =====================================================================================
# Share Snapshots
# =====================================================================================

resource "ibm_is_share_snapshot" "share-snapshot-1" {
  name  = "share-snapshot-1"
  share = ibm_is_share.share-1.id
  tags  = ["share-snapshot-tag"]
}

# =====================================================================================
# Backup Policies
# =====================================================================================

// Backup policy for volumes
resource "ibm_is_backup_policy" "backup-policy-volumes" {
  match_user_tags     = ["tag1"]
  name                = "backup-policy-volumes"
  match_resource_type = "volume"
}

// Backup policy for instances
resource "ibm_is_backup_policy" "backup-policy-instances" {
  match_user_tags     = ["tag1"]
  name                = "backup-policy-instances"
  match_resource_type = "instance"
  included_content    = ["boot_volume", "data_volumes"]
}

// Enterprise backup policy with scope
resource "ibm_is_backup_policy" "backup-policy-enterprise" {
  match_user_tags = ["tag1"]
  name            = "backup-policy-enterprise"
  scope {
    crn = "crn:v1:bluemix:public:is:us-south:a/123456::reservation:7187-ba49df72-37b8-43ac-98da-f8e029de0e63"
  }
}

# =====================================================================================
# Backup Policy Plans
# =====================================================================================

// Basic backup policy plan
resource "ibm_is_backup_policy_plan" "backup-policy-plan-1" {
  backup_policy_id = ibm_is_backup_policy.backup-policy-volumes.id
  cron_spec        = "30 09 * * *"
  active           = false
  attach_user_tags = ["tag2"]
  copy_user_tags   = true

  deletion_trigger {
    delete_after      = 20
    delete_over_count = "20"
  }

  name = "backup-policy-plan-1"
}

// Backup policy plan with clone policy
resource "ibm_is_backup_policy_plan" "backup-policy-plan-clone" {
  backup_policy_id = ibm_is_backup_policy.backup-policy-volumes.id
  cron_spec        = "30 09 * * *"
  active           = false
  attach_user_tags = ["tag2"]
  copy_user_tags   = true

  deletion_trigger {
    delete_after      = 20
    delete_over_count = "20"
  }

  name = "backup-policy-plan-clone"

  clone_policy {
    zones         = ["us-south-1", "us-south-2"]
    max_snapshots = 3
  }
}

# =====================================================================================
# Data Sources for Storage & Backup
# =====================================================================================

// Snapshot data sources
data "ibm_is_snapshot" "snapshot-data" {
  name = "snapshot-boot"
}

data "ibm_is_snapshots" "all-snapshots" {
}

data "ibm_is_snapshot_clones" "snapshot-clones" {
  snapshot = ibm_is_snapshot.snapshot-boot.id
}

data "ibm_is_snapshot_clones" "snapshot-clone-by-zone" {
  snapshot = ibm_is_snapshot.snapshot-boot.id
}

// Snapshot consistency group data sources
data "ibm_is_snapshot_consistency_group" "snapshot-consistency-group-data" {
  identifier = ibm_is_snapshot_consistency_group.snapshot-consistency-group-1.id
}

data "ibm_is_snapshot_consistency_group" "snapshot-consistency-group-by-name" {
  name = "snapshot-consistency-group-1"
}

data "ibm_is_snapshot_consistency_groups" "all-snapshot-consistency-groups" {
  name = "snapshot-consistency-group-1"
}

// Image data sources
data "ibm_is_image" "custom-image-data" {
  name = ibm_is_image.image-1.name
}

data "ibm_is_images" "all-custom-images" {
}

// Image export job data sources
data "ibm_is_image_export_jobs" "image-export-jobs" {
  image = ibm_is_image_export_job.image-export-1.image
}

data "ibm_is_image_export_job" "image-export-job-data" {
  image            = ibm_is_image_export_job.image-export-1.image
  image_export_job = ibm_is_image_export_job.image-export-1.image_export_job
}

// Share data sources
data "ibm_is_share" "share-data" {
  share = ibm_is_share.share-1.id
}

data "ibm_is_shares" "all-shares" {
}

// Share mount target data sources
data "ibm_is_share_mount_target" "share-mount-target-data" {
  share        = ibm_is_share.share-1.id
  mount_target = ibm_is_share_mount_target.share-mount-target-1.mount_target
}

data "ibm_is_share_mount_targets" "share-mount-targets" {
  share = ibm_is_share.share-1.id
}

// Share snapshot data sources
data "ibm_is_share_snapshots" "share-snapshots-by-share" {
  share = ibm_is_share.share-1.id
}

data "ibm_is_share_snapshots" "all-share-snapshots" {
}

data "ibm_is_share_snapshot" "share-snapshot-data" {
  share          = ibm_is_share.share-1.id
  share_snapshot = ibm_is_share_snapshot.share-snapshot-1.share_snapshot
}

// Backup policy data sources
data "ibm_is_backup_policies" "all-backup-policies" {
}

data "ibm_is_backup_policy" "backup-policy-data" {
  name = "backup-policy-volumes"
}

data "ibm_is_backup_policy" "backup-policy-enterprise-data" {
  name = ibm_is_backup_policy.backup-policy-enterprise.name
}

// Backup policy plan data sources
data "ibm_is_backup_policy_plans" "backup-policy-plans" {
  backup_policy_id = ibm_is_backup_policy.backup-policy-volumes.id
}

data "ibm_is_backup_policy_plan" "backup-policy-plan-data" {
  backup_policy_id = ibm_is_backup_policy.backup-policy-volumes.id
  name             = "backup-policy-plan-1"
}

// Backup policy job data sources
data "ibm_is_backup_policy_job" "backup-policy-job-data" {
  backup_policy_id = ibm_is_backup_policy.backup-policy-volumes.id
  identifier       = ""
}

data "ibm_is_backup_policy_jobs" "backup-policy-jobs" {
  backup_policy_id = ibm_is_backup_policy.backup-policy-volumes.id
}

# =====================================================================================
# Dedicated Host Groups
# =====================================================================================

resource "ibm_is_dedicated_host_group" "host-group-1" {
  family         = "balanced"
  class          = "bx2d"
  zone           = "us-south-1"
  name           = "host-group-1"
  resource_group = data.ibm_resource_group.default.id
}

# =====================================================================================
# Dedicated Hosts
# =====================================================================================

resource "ibm_is_dedicated_host" "dedicated-host-1" {
  profile        = "bx2d-host-152x608"
  name           = "dedicated-host-1"
  host_group     = ibm_is_dedicated_host_group.host-group-1.id
  resource_group = data.ibm_resource_group.default.id
}

# =====================================================================================
# Dedicated Host Disk Management
# =====================================================================================

resource "ibm_is_dedicated_host_disk_management" "dedicated-host-disks-1" {
  dedicated_host = ibm_is_dedicated_host.dedicated-host-1.id

  disks {
    name = "dedicated-host-disk-1"
    id   = ibm_is_dedicated_host.dedicated-host-1.disks.0.id
  }

  disks {
    name = "dedicated-host-disk-2"
    id   = ibm_is_dedicated_host.dedicated-host-1.disks.1.id
  }
}

# =====================================================================================
# Instance Templates
# =====================================================================================

// Basic instance template
resource "ibm_is_instance_template" "instance-template-1" {
  name    = "instance-template-1"
  image   = "r006-618b224d-eb88-492f-9825-dc246bb5211a"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet            = ibm_is_subnet.subnet-secondary-1.id
    allow_ip_spoofing = true
  }

  vpc  = ibm_is_vpc.vpc-secondary-1.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.ssh-key-1.id]

  boot_volume {
    name                             = "instance-template-boot-vol"
    delete_volume_on_instance_delete = true
  }

  volume_attachments {
    delete_volume_on_instance_delete = true
    name                             = "instance-template-vol-attachment-1"
    volume_prototype {
      iops     = 3000
      profile  = "general-purpose"
      capacity = 200
    }
  }
}

// Instance template with existing volume
resource "ibm_is_instance_template" "instance-template-2" {
  name    = "instance-template-2"
  image   = "r006-618b224d-eb88-492f-9825-dc246bb5211a"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet            = ibm_is_subnet.subnet-secondary-1.id
    allow_ip_spoofing = true
  }

  vpc  = ibm_is_vpc.vpc-secondary-1.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.ssh-key-1.id]

  boot_volume {
    name                             = "instance-template-boot-vol-2"
    delete_volume_on_instance_delete = true
  }

  volume_attachments {
    delete_volume_on_instance_delete = true
    name                             = "instance-template-vol-attachment-2"
    volume                           = ibm_is_volume.volume-1.id
  }
}

// Instance template with catalog offering
resource "ibm_is_instance_template" "instance-template-3" {
  name = "instance-template-3"

  catalog_offering {
    version_crn = data.ibm_is_images.catalog-images.images.0.catalog_offering.0.version.0.crn
    plan_crn    = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
  }

  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-secondary-1.id
  }

  vpc  = ibm_is_vpc.vpc-secondary-1.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.ssh-key-1.id]
}

// Instance template with VNI
resource "ibm_is_instance_template" "instance-template-vni" {
  name    = "instance-template-vni"
  profile = "bx2-2x8"
  image   = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"

  primary_network_attachment {
    name = "template-vni-attachment"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.vni-3.id
    }
  }

  vpc  = ibm_is_vpc.vpc-secondary-2.id
  zone = "${var.region}-2"
  keys = [ibm_is_ssh_key.ssh-key-ed25519.id]
}

# =====================================================================================
# Instance from Template
# =====================================================================================

// Creating an instance using an existing instance template
resource "ibm_is_instance" "vsi-from-template" {
  name              = "vsi-from-template"
  instance_template = ibm_is_instance_template.instance-template-1.id
  zone              = "us-south-3"
  keys              = [ibm_is_ssh_key.ssh-key-1.id]

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-primary-1.id
  }

  vpc = ibm_is_vpc.vpc-primary.id
}

# =====================================================================================
# Bare Metal Servers
# =====================================================================================

// Basic bare metal server
resource "ibm_is_bare_metal_server" "bare-metal-1" {
  profile = "bx2-metal-192x768"
  name    = "bare-metal-1"
  image   = "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e"
  zone    = "us-south-3"
  keys    = [ibm_is_ssh_key.ssh-key-1.id]

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-primary-1.id
  }

  vpc = ibm_is_vpc.vpc-primary.id
}

// Bare metal server with VNI
resource "ibm_is_bare_metal_server" "bare-metal-vni" {
  profile = "cx2-metal-96x192"
  name    = "bare-metal-vni"
  image   = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"
  zone    = "${var.region}-2"
  keys    = [ibm_is_ssh_key.ssh-key-ed25519.id]

  primary_network_attachment {
    name = "bare-metal-vni-attachment"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.vni-1.id
    }
    allowed_vlans = [100, 102]
  }

  vpc = ibm_is_vpc.vpc-secondary-2.id
}

# =====================================================================================
# Bare Metal Server Actions
# =====================================================================================

resource "ibm_is_bare_metal_server_action" "bare-metal-stop" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  action            = "stop"
  stop_type         = "hard"
}

# =====================================================================================
# Bare Metal Server Disk Management
# =====================================================================================

resource "ibm_is_bare_metal_server_disk" "bare-metal-disk-1" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  disk              = ibm_is_bare_metal_server.bare-metal-1.disks.0.id
  name              = "bare-metal-disk-1"
}

# =====================================================================================
# Bare Metal Server Network Interfaces
# =====================================================================================

// Basic network interface
resource "ibm_is_bare_metal_server_network_interface" "bare-metal-nic-1" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  subnet            = ibm_is_subnet.subnet-primary-1.id
  name              = "bare-metal-nic-1"
  allow_ip_spoofing = true
  allowed_vlans     = [101, 102]
}

// VLAN network interface
resource "ibm_is_bare_metal_server_network_interface" "bare-metal-nic-vlan" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  subnet            = ibm_is_subnet.subnet-primary-1.id
  name              = "bare-metal-nic-vlan"
  allow_ip_spoofing = true
  vlan              = 101
}

// Allow float network interface
resource "ibm_is_bare_metal_server_network_interface_allow_float" "bare-metal-nic-float" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  subnet            = ibm_is_subnet.subnet-primary-1.id
  name              = "bare-metal-nic-float"
  vlan              = 102
}

// Bare metal server network attachments for VNI-based server
resource "ibm_is_bare_metal_server_network_attachment" "bare-metal-attachment-1" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-vni.id
  vlan              = 100
}

resource "ibm_is_bare_metal_server_network_attachment" "bare-metal-attachment-2" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-vni.id
  allowed_vlans     = [200, 202]
}

# =====================================================================================
# Bare Metal Server Floating IPs
# =====================================================================================

resource "ibm_is_floating_ip" "bare-metal-floating-ip" {
  name = "bare-metal-floating-ip"
  zone = ibm_is_subnet.subnet-primary-1.zone
}

resource "ibm_is_bare_metal_server_network_interface_floating_ip" "bare-metal-nic-fip" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  network_interface = ibm_is_bare_metal_server_network_interface.bare-metal-nic-vlan.id
  floating_ip       = ibm_is_floating_ip.bare-metal-floating-ip.id
}

# =====================================================================================
# Placement Groups
# =====================================================================================

resource "ibm_is_placement_group" "placement-group-1" {
  strategy       = "host_spread"
  name           = "placement-group-1"
  resource_group = data.ibm_resource_group.default.id
}

# =====================================================================================
# Reservations
# =====================================================================================

resource "ibm_is_reservation" "reservation-1" {
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

resource "ibm_is_reservation_activate" "reservation-1-activate" {
  reservation = ibm_is_reservation.reservation-1.id
}

# =====================================================================================
# Data Sources for Advanced Compute
# =====================================================================================

// Dedicated host group data sources
data "ibm_is_dedicated_host_group" "host-group-data" {
  name = ibm_is_dedicated_host_group.host-group-1.name
}

data "ibm_is_dedicated_host_groups" "all-host-groups" {
}

// Dedicated host data sources
data "ibm_is_dedicated_host" "dedicated-host-data" {
  name       = ibm_is_dedicated_host.dedicated-host-1.name
  host_group = data.ibm_is_dedicated_host_group.host-group-data.id
}

data "ibm_is_dedicated_hosts" "all-dedicated-hosts" {
}

data "ibm_is_dedicated_host_profile" "dedicated-host-profile-data" {
  name = "bx2d-host-152x608"
}

data "ibm_is_dedicated_host_profiles" "all-dedicated-host-profiles" {
}

// Dedicated host disk data sources
data "ibm_is_dedicated_host_disks" "dedicated-host-disks-data" {
  dedicated_host = data.ibm_is_dedicated_host.dedicated-host-data.id
}

data "ibm_is_dedicated_host_disk" "dedicated-host-disk-data" {
  dedicated_host = data.ibm_is_dedicated_host.dedicated-host-data.id
  disk           = ibm_is_dedicated_host_disk_management.dedicated-host-disks-1.disks.0.id
}

// Instance template data sources
data "ibm_is_instance_template" "instance-template-data" {
  identifier = ibm_is_instance_template.instance-template-2.id
}

data "ibm_is_instance_template" "instance-template-by-name" {
  name = ibm_is_instance_template.instance-template-vni.name
}

// Bare metal server data sources
data "ibm_is_bare_metal_servers" "all-bare-metal-servers" {
}

data "ibm_is_bare_metal_server" "bare-metal-server-data" {
  identifier = ibm_is_bare_metal_server.bare-metal-1.id
}

data "ibm_is_bare_metal_server_profiles" "all-bare-metal-profiles" {
}

data "ibm_is_bare_metal_server_profile" "bare-metal-profile-data" {
  name = data.ibm_is_bare_metal_server_profiles.all-bare-metal-profiles.profiles.0.name
}

// Bare metal server initialization data
data "ibm_is_bare_metal_server_initialization" "bare-metal-init-data" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
}

// Bare metal server disk data sources
data "ibm_is_bare_metal_server_disk" "bare-metal-disk-data" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  disk              = ibm_is_bare_metal_server.bare-metal-1.disks.0.id
}

data "ibm_is_bare_metal_server_disks" "bare-metal-disks-data" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
}

// Bare metal server network interface data sources
data "ibm_is_bare_metal_server_network_interface" "bare-metal-nic-data" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  network_interface = ibm_is_bare_metal_server.bare-metal-1.primary_network_interface.0.id
}

data "ibm_is_bare_metal_server_network_interfaces" "bare-metal-nics-data" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
}

// Bare metal server floating IP data sources
data "ibm_is_bare_metal_server_network_interface_floating_ip" "bare-metal-fip-data" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  network_interface = ibm_is_bare_metal_server.bare-metal-1.primary_network_interface[0].id
  floating_ip       = ibm_is_floating_ip.bare-metal-floating-ip.id
}

data "ibm_is_bare_metal_server_network_interface_floating_ips" "bare-metal-fips-data" {
  bare_metal_server = ibm_is_bare_metal_server.bare-metal-1.id
  network_interface = ibm_is_bare_metal_server.bare-metal-1.primary_network_interface[0].id
}

// Placement group data sources
data "ibm_is_placement_group" "placement-group-data" {
  name = ibm_is_placement_group.placement-group-1.name
}

data "ibm_is_placement_groups" "all-placement-groups" {
}

// Reservation data sources
data "ibm_is_reservations" "all-reservations" {
}

data "ibm_is_reservation" "reservation-data" {
  identifier = ibm_is_reservation.reservation-1.id
}

// Instance profile data sources
data "ibm_is_instance_profile" "instance-profile-data" {
  name = "gx3d-160x1792x8h100"
}

data "ibm_is_instance_profiles" "all-instance-profiles" {
}

# =====================================================================================
# DNS Services
# =====================================================================================

// DNS service instance
resource "ibm_resource_instance" "dns-service-instance" {
  name              = "dns-service-instance"
  resource_group_id = data.ibm_resource_group.default.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "standard-dns"
}

// Custom DNS resolver for hub VPC
resource "ibm_dns_custom_resolver" "dns-resolver-hub" {
  name              = "dns-resolver-hub"
  instance_id       = ibm_resource_instance.dns-service-instance.guid
  description       = "Custom DNS resolver for hub VPC"
  high_availability = true
  enabled           = true

  locations {
    subnet_crn = ibm_is_subnet.subnet-hub-1.crn
    enabled    = true
  }

  locations {
    subnet_crn = ibm_is_subnet.subnet-hub-2.crn
    enabled    = true
  }
}

// Custom DNS resolver for delegated VPC
resource "ibm_dns_custom_resolver" "dns-resolver-delegated" {
  name              = "dns-resolver-delegated"
  instance_id       = ibm_resource_instance.dns-service-instance.guid
  description       = "Custom DNS resolver for delegated VPC"
  high_availability = true
  enabled           = true

  locations {
    subnet_crn = ibm_is_subnet.subnet-delegated-1.crn
    enabled    = true
  }

  locations {
    subnet_crn = ibm_is_subnet.subnet-delegated-2.crn
    enabled    = true
  }
}

# =====================================================================================
# VPN Servers
# =====================================================================================

resource "ibm_is_vpn_server" "vpn-server-1" {
  certificate_crn = var.is_certificate_crn

  client_authentication {
    method        = "certificate"
    client_ca_crn = var.is_client_ca
  }

  client_ip_pool         = "10.5.0.0/21"
  subnets                = [ibm_is_subnet.subnet-primary-1.id]
  client_dns_server_ips  = ["192.168.3.4"]
  client_idle_timeout    = 2800
  enable_split_tunneling = false
  name                   = "vpn-server-1"
  port                   = 443
  protocol               = "udp"
}

# =====================================================================================
# VPN Server Routes
# =====================================================================================

resource "ibm_is_vpn_server_route" "vpn-server-route-1" {
  vpn_server  = ibm_is_vpn_server.vpn-server-1.vpn_server
  destination = "172.16.0.0/16"
  action      = "translate"
  name        = "vpn-server-route-1"
}

# =====================================================================================
# Cluster Networks
# =====================================================================================

// Cluster network profile data sources
data "ibm_is_cluster_network_profile" "cluster-network-profile-h100" {
  name = "h100"
}

data "ibm_is_cluster_network_profiles" "all-cluster-network-profiles" {
}

// Cluster network
resource "ibm_is_cluster_network" "cluster-network-1" {
  name           = "cluster-network-1"
  profile        = "h100"
  resource_group = data.ibm_resource_group.default.id

  subnet_prefixes {
    cidr = "10.1.0.0/24"
  }

  vpc {
    id = ibm_is_vpc.vpc-primary.id
  }

  zone = "${var.region}-3"
}

# =====================================================================================
# Cluster Network Subnets
# =====================================================================================

resource "ibm_is_cluster_network_subnet" "cluster-subnet-1" {
  cluster_network_id       = ibm_is_cluster_network.cluster-network-1.id
  name                     = "cluster-subnet-1"
  total_ipv4_address_count = 64
}

# =====================================================================================
# Cluster Network Reserved IPs
# =====================================================================================

resource "ibm_is_cluster_network_subnet_reserved_ip" "cluster-reserved-ip-1" {
  cluster_network_id        = ibm_is_cluster_network.cluster-network-1.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
  address                   = "10.1.0.4"
  name                      = "cluster-reserved-ip-1"
}

# =====================================================================================
# Cluster Network Interfaces
# =====================================================================================

resource "ibm_is_cluster_network_interface" "cluster-network-interface-1" {
  cluster_network_id = ibm_is_cluster_network.cluster-network-1.id
  name               = "cluster-network-interface-1"

  primary_ip {
    id = ibm_is_cluster_network_subnet_reserved_ip.cluster-reserved-ip-1.cluster_network_subnet_reserved_ip_id
  }

  subnet {
    id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
  }
}

# =====================================================================================
# High-Performance Computing Instances
# =====================================================================================

// HPC-optimized instance with cluster network attachments
resource "ibm_is_instance" "hpc-instance-1" {
  name    = "hpc-instance-1"
  image   = data.ibm_is_image.ubuntu-image.id
  profile = "gx3d-160x1792x8h100"

  primary_network_interface {
    subnet = ibm_is_subnet.subnet-primary-1.id
  }

  // Multiple cluster network attachments for high-performance networking
  cluster_network_attachments {
    name = "cluster-attachment-1"
    cluster_network_interface {
      auto_delete = true
      name        = "cluster-interface-1"
      subnet {
        id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
      }
    }
  }

  cluster_network_attachments {
    name = "cluster-attachment-2"
    cluster_network_interface {
      auto_delete = true
      name        = "cluster-interface-2"
      subnet {
        id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
      }
    }
  }

  cluster_network_attachments {
    name = "cluster-attachment-3"
    cluster_network_interface {
      auto_delete = true
      name        = "cluster-interface-3"
      subnet {
        id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
      }
    }
  }

  cluster_network_attachments {
    name = "cluster-attachment-4"
    cluster_network_interface {
      auto_delete = true
      name        = "cluster-interface-4"
      subnet {
        id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
      }
    }
  }

  vpc  = ibm_is_vpc.vpc-primary.id
  zone = ibm_is_subnet.subnet-primary-1.zone
  keys = [ibm_is_ssh_key.ssh-key-1.id]
}

# =====================================================================================
# Instance Cluster Network Attachments (Dynamic)
# =====================================================================================

// Stop instance before adding more cluster network attachments
resource "ibm_is_instance_action" "hpc-instance-stop" {
  action   = "stop"
  instance = ibm_is_instance.hpc-instance-1.id
}

// Add additional cluster network attachments
resource "ibm_is_instance_cluster_network_attachment" "cluster-attachment-5" {
  depends_on  = [ibm_is_instance_action.hpc-instance-stop]
  instance_id = ibm_is_instance.hpc-instance-1.id

  before {
    id = ibm_is_instance.hpc-instance-1.cluster_network_attachments.0.id
  }

  cluster_network_interface {
    name = "cluster-interface-5"
    subnet {
      id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
    }
  }

  name = "cluster-attachment-5"
}

resource "ibm_is_instance_cluster_network_attachment" "cluster-attachment-6" {
  depends_on  = [ibm_is_instance_action.hpc-instance-stop]
  instance_id = ibm_is_instance.hpc-instance-1.id

  before {
    id = ibm_is_instance_cluster_network_attachment.cluster-attachment-5.instance_cluster_network_attachment_id
  }

  cluster_network_interface {
    name = "cluster-interface-6"
    subnet {
      id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
    }
  }

  name = "cluster-attachment-6"
}

// Start instance after adding cluster network attachments
resource "ibm_is_instance_action" "hpc-instance-start" {
  depends_on = [
    ibm_is_instance_cluster_network_attachment.cluster-attachment-5,
    ibm_is_instance_cluster_network_attachment.cluster-attachment-6
  ]
  action   = "start"
  instance = ibm_is_instance.hpc-instance-1.id
}

# =====================================================================================
# HPC Instance Template with Cluster Networking
# =====================================================================================

resource "ibm_is_instance_template" "hpc-template-1" {
  name    = "hpc-template-1"
  image   = data.ibm_is_image.ubuntu-image.id
  profile = "gx3d-160x1792x8h100"

  primary_network_attachment {
    name = "hpc-template-primary-attachment"
    virtual_network_interface {
      auto_delete = true
      subnet      = ibm_is_subnet.subnet-primary-1.id
    }
  }

  // Template with multiple cluster network attachments
  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.cluster-network-interface-1.cluster_network_interface_id
    }
  }

  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.cluster-network-interface-1.cluster_network_interface_id
    }
  }

  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.cluster-network-interface-1.cluster_network_interface_id
    }
  }

  cluster_network_attachments {
    cluster_network_interface {
      id = ibm_is_cluster_network_interface.cluster-network-interface-1.cluster_network_interface_id
    }
  }

  vpc  = ibm_is_vpc.vpc-primary.id
  zone = ibm_is_subnet.subnet-primary-1.zone
  keys = [ibm_is_ssh_key.ssh-key-1.id]
}

# =====================================================================================
# Data Sources for Specialized Services
# =====================================================================================

// Ubuntu image for HPC instances
data "ibm_is_image" "ubuntu-image" {
  name = "ibm-ubuntu-20-04-6-minimal-amd64-6"
}

// VPN server data sources
data "ibm_is_vpn_server" "vpn-server-data" {
  identifier = ibm_is_vpn_server.vpn-server-1.vpn_server
}

data "ibm_is_vpn_servers" "all-vpn-servers" {
}

// VPN server route data sources
data "ibm_is_vpn_server_routes" "vpn-server-routes" {
  vpn_server = ibm_is_vpn_server.vpn-server-1.vpn_server
}

data "ibm_is_vpn_server_route" "vpn-server-route-data" {
  vpn_server = ibm_is_vpn_server.vpn-server-1.vpn_server
  identifier = ibm_is_vpn_server_route.vpn-server-route-1.vpn_route
}

// VPN server client data sources
data "ibm_is_vpn_server_clients" "vpn-server-clients" {
  vpn_server = ibm_is_vpn_server.vpn-server-1.vpn_server
}

data "ibm_is_vpn_server_client" "vpn-server-client-data" {
  vpn_server = ibm_is_vpn_server.vpn-server-1.vpn_server
  identifier = "0726-61b2f53f-1e95-42a7-94ab-55de8f8cbdd5"
}

// Cluster network data sources
data "ibm_is_cluster_network" "cluster-network-data" {
  cluster_network_id = ibm_is_cluster_network.cluster-network-1.id
}

data "ibm_is_cluster_networks" "all-cluster-networks" {
}

// Cluster network interface data sources
data "ibm_is_cluster_network_interface" "cluster-interface-data" {
  cluster_network_id           = ibm_is_cluster_network.cluster-network-1.id
  cluster_network_interface_id = ibm_is_cluster_network_interface.cluster-network-interface-1.cluster_network_interface_id
}

data "ibm_is_cluster_network_interfaces" "cluster-interfaces-data" {
  cluster_network_id = ibm_is_cluster_network.cluster-network-1.id
}

// Cluster network subnet data sources
data "ibm_is_cluster_network_subnet" "cluster-subnet-data" {
  cluster_network_id        = ibm_is_cluster_network.cluster-network-1.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
}

data "ibm_is_cluster_network_subnets" "cluster-subnets-data" {
  cluster_network_id = ibm_is_cluster_network.cluster-network-1.id
}

// Cluster network reserved IP data sources
data "ibm_is_cluster_network_subnet_reserved_ip" "cluster-reserved-ip-data" {
  cluster_network_id                    = ibm_is_cluster_network.cluster-network-1.id
  cluster_network_subnet_id             = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
  cluster_network_subnet_reserved_ip_id = ibm_is_cluster_network_subnet_reserved_ip.cluster-reserved-ip-1.cluster_network_subnet_reserved_ip_id
}

data "ibm_is_cluster_network_subnet_reserved_ips" "cluster-reserved-ips-data" {
  cluster_network_id        = ibm_is_cluster_network.cluster-network-1.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.cluster-subnet-1.cluster_network_subnet_id
}

// Instance cluster network attachment data sources
data "ibm_is_instance_cluster_network_attachment" "cluster-attachment-data" {
  instance_id                            = ibm_is_instance.hpc-instance-1.id
  instance_cluster_network_attachment_id = ibm_is_instance.hpc-instance-1.cluster_network_attachments.0.id
}

data "ibm_is_instance_cluster_network_attachments" "cluster-attachments-data" {
  instance_id = ibm_is_instance.hpc-instance-1.id
}

// HPC instance template data source
data "ibm_is_instance_template" "hpc-template-data" {
  name = ibm_is_instance_template.hpc-template-1.name
}

// HPC instance data source
data "ibm_is_instance" "hpc-instance-data" {
  name = ibm_is_instance.hpc-instance-1.name
}