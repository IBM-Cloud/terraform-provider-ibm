# ==========================================================================
# VPC Resources
# ==========================================================================

# Create a standard VPC for primary workloads
resource "ibm_is_vpc" "vpc1" {
  name = "vpc1"
}

# Create a second VPC for separation of workloads
resource "ibm_is_vpc" "vpc2" {
  name = "vpc2"
}

# Create a VPC with DNS hub enabled for central DNS resolution
resource "ibm_is_vpc" "hub_true" {
  name = "vpc-hub-true"
  dns {
    enable_hub = true # Enable as DNS hub
  }
}

# Create a VPC that can use the DNS hub VPC for name resolution
resource "ibm_is_vpc" "hub_false_delegated" {
  name = "vpc-hub-false-del"
  dns {
    enable_hub = false
    # Uncomment to delegate DNS resolution to the hub VPC
    # resolver {
    #   type = "delegated"
    #   vpc_id = ibm_is_vpc.hub_true.id
    # }
  }
}

# Create a VPC for test resources
resource "ibm_is_vpc" "testacc_vpc" {
  name = "${var.name}-vpc"
}

# Create a VPC for cluster resources
resource "ibm_is_vpc" "is_vpc" {
  name = "${var.prefix}-vpc"
}

# Create a second VPC for cluster resources
resource "ibm_is_vpc" "is_vpc2" {
  name = "${var.prefix}-vpc2"
}

# ==========================================================================
# VPC Address Prefix Resources
# ==========================================================================

# Create an address prefix for VPC1 to define the IP address range
resource "ibm_is_vpc_address_prefix" "vpc_address_prefix" {
  name       = "vpcaddressprefix"
  zone       = var.zone1          # Zone where the prefix will be created
  vpc        = ibm_is_vpc.vpc1.id # Reference to the VPC
  cidr       = var.cidr1          # CIDR block for the address range
  is_default = true               # Set as the default address prefix
}

# ==========================================================================
# Subnet Resources
# ==========================================================================

# Create a subnet in VPC1 with a small CIDR block (16 IP addresses)
resource "ibm_is_subnet" "subnet1" {
  name            = "subnet1"
  vpc             = ibm_is_vpc.vpc1.id # VPC where the subnet will be created
  zone            = var.zone1          # Zone for the subnet
  ipv4_cidr_block = "10.240.0.0/28"    # IPv4 CIDR block for the subnet (16 IP addresses)
}

# Create a subnet in VPC2
resource "ibm_is_subnet" "subnet2" {
  name            = "subnet2"
  vpc             = ibm_is_vpc.vpc2.id
  zone            = var.zone2
  ipv4_cidr_block = "10.240.64.0/28" # Different CIDR range than subnet1
}

# Create a subnet for the test VPC
resource "ibm_is_subnet" "testacc_subnet" {
  name                     = "${var.name}-subnet"
  vpc                      = ibm_is_vpc.testacc_vpc.id
  zone                     = "${var.region}-2"
  total_ipv4_address_count = 16 # Alternative way to specify size
}

# Create a subnet for the cluster VPC
resource "ibm_is_subnet" "is_subnet" {
  name                     = "${var.prefix}-subnet"
  vpc                      = ibm_is_vpc.is_vpc.id
  total_ipv4_address_count = 64 # 64 IP addresses for cluster resources
  zone                     = "${var.region}-3"
}

# Create subnets for the DNS hub VPC
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

# Create subnets for the non-hub VPC
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

# ==========================================================================
# Subnet Reserved IP Resources
# ==========================================================================

# Reserve a specific IP address in a subnet
resource "ibm_is_subnet_reserved_ip" "testacc_reservedip" {
  subnet = ibm_is_subnet.testacc_subnet.id
  name   = "${var.name}-reserved-ip"
}

# ==========================================================================
# Public Gateway Resources
# ==========================================================================

# Create a public gateway to allow outbound internet access
resource "ibm_is_public_gateway" "publicgateway1" {
  name = "gateway1"
  vpc  = ibm_is_vpc.vpc1.id # VPC for the gateway
  zone = var.zone1          # Zone for the gateway
}

# Attach a subnet to a public gateway
resource "ibm_is_subnet_public_gateway_attachment" "example" {
  subnet         = ibm_is_subnet.subnet1.id                # Subnet to attach
  public_gateway = ibm_is_public_gateway.publicgateway1.id # Public gateway to attach to
}

# ==========================================================================
# VPC DNS Resolution Resources
# ==========================================================================

# Create a DNS service instance
resource "ibm_resource_instance" "dns-cr-instance" {
  name              = "dns-cr-instance"
  resource_group_id = data.ibm_resource_group.default.id
  location          = "global"       # Global service
  service           = "dns-svcs"     # DNS Services
  plan              = "standard-dns" # Standard DNS plan
}

# Create a custom resolver for the hub VPC
resource "ibm_dns_custom_resolver" "test_hub_true" {
  name              = "test-hub-true-customresolver"
  instance_id       = ibm_resource_instance.dns-cr-instance.guid
  description       = "new test CR - TF"
  high_availability = true # Enable high availability
  enabled           = true # Enable the resolver

  # First location for the resolver
  locations {
    subnet_crn = ibm_is_subnet.hub_true_sub1.crn
    enabled    = true
  }

  # Second location for high availability
  locations {
    subnet_crn = ibm_is_subnet.hub_true_sub2.crn
    enabled    = true
  }
}

# Create a custom resolver for the non-hub VPC
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

# Create a DNS resolution binding between VPCs
resource "ibm_is_vpc_dns_resolution_binding" "dnstrue" {
  name   = "hub-spoke-binding"
  vpc_id = ibm_is_vpc.hub_false_delegated.id # The binding is created in this VPC
  vpc {
    id = ibm_is_vpc.hub_true.id # The target VPC (hub) for resolution
  }
}

# ==========================================================================
# VPC Data Sources
# ==========================================================================

# Get VPC by name
data "ibm_is_vpc" "vpc1" {
  name = ibm_is_vpc.vpc1.id
}

# List all VPCs
data "ibm_is_vpc" "vpcs" {
}

# ==========================================================================
# VPC Address Prefix Data Sources
# ==========================================================================

# Get address prefix by VPC ID and prefix
data "ibm_is_vpc_address_prefix" "example" {
  vpc            = ibm_is_vpc.vpc1.id
  address_prefix = ibm_is_vpc_address_prefix.vpc_address_prefix.address_prefix
}

# Get address prefix by VPC name and prefix
data "ibm_is_vpc_address_prefix" "example-1" {
  vpc_name       = ibm_is_vpc.vpc1.name
  address_prefix = ibm_is_vpc_address_prefix.vpc_address_prefix.address_prefix
}

# Get address prefix by VPC ID and prefix name
data "ibm_is_vpc_address_prefix" "example-2" {
  vpc                 = ibm_is_vpc.vpc1.id
  address_prefix_name = ibm_is_vpc_address_prefix.vpc_address_prefix.name
}

# Get address prefix by VPC name and prefix name
data "ibm_is_vpc_address_prefix" "example-3" {
  vpc_name            = ibm_is_vpc.vpc1.name
  address_prefix_name = ibm_is_vpc_address_prefix.vpc_address_prefix.name
}

# ==========================================================================
# Public Gateway Data Sources
# ==========================================================================

# Get public gateway by name
data "ibm_is_public_gateway" "testacc_dspgw" {
  name = ibm_is_public_gateway.publicgateway1.name
}

# List all public gateways
data "ibm_is_public_gateways" "publicgateways" {
}

# ==========================================================================
# DNS Resolution Data Sources
# ==========================================================================

# List all DNS resolution bindings for a VPC
data "ibm_is_vpc_dns_resolution_bindings" "is_vpc_dns_resolution_bindings" {
  vpc_id = ibm_is_vpc.vpc1.id
}

# Get a specific DNS resolution binding
data "ibm_is_vpc_dns_resolution_binding" "is_vpc_dns_resolution_binding" {
  vpc_id = ibm_is_vpc.vpc1.id
  id     = ibm_is_vpc.vpc2.id
}

# ==========================================================================
# Security Group Resources
# ==========================================================================

# Create a custom security group
resource "ibm_is_security_group" "example" {
  name = "example-security-group"
  vpc  = ibm_is_vpc.vpc1.id
}

# ==========================================================================
# Security Group Rules Resources
# ==========================================================================

# Allow SSH access (TCP port 22) to VPC1 default security group
resource "ibm_is_security_group_rule" "sg1_tcp_rule" {
  depends_on = [ibm_is_floating_ip.floatingip1] # Create after floating IP is created
  group      = ibm_is_vpc.vpc1.default_security_group
  direction  = "inbound"
  remote     = "0.0.0.0/0" # Allow from any IP address

  tcp {
    port_min = 22
    port_max = 22
  }
}

# Allow ICMP (ping) to VPC1 default security group
resource "ibm_is_security_group_rule" "sg1_icmp_rule" {
  depends_on = [ibm_is_floating_ip.floatingip1]
  group      = ibm_is_vpc.vpc1.default_security_group
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  icmp {
    code = 0 # Echo Reply
    type = 8 # Echo Request
  }
}

# Allow HTTP access (TCP port 80) to VPC1 default security group
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

# Allow SSH access (TCP port 22) to VPC2 default security group
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

# Allow ICMP (ping) to VPC2 default security group
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

# Allow HTTP access (TCP port 80) to VPC2 default security group
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

# Add UDP rule to custom security group
resource "ibm_is_security_group_rule" "exampleudp" {
  depends_on = [ibm_is_security_group.example]
  group      = ibm_is_security_group.example.id
  direction  = "inbound"
  remote     = "127.0.0.1" # Allow from localhost only

  udp {
    port_min = 805
    port_max = 807
  }
}

# Add TCP rule to custom security group
resource "ibm_is_security_group_rule" "exampletcp" {
  group     = ibm_is_security_group.example.id
  direction = "outbound"
  remote    = "127.0.0.1"

  tcp {
    port_min = 8080
    port_max = 8080
  }

  depends_on = [ibm_is_security_group.example]
}

# ==========================================================================
# Network ACL Resources
# ==========================================================================

# Create a network ACL with outbound and inbound rules
resource "ibm_is_network_acl" "isExampleACL" {
  name = "is-example-acl"
  vpc  = ibm_is_vpc.vpc1.id

  # Outbound rule - allow all TCP traffic
  rules {
    name        = "outbound"
    action      = "allow"
    source      = "0.0.0.0/0" # From any source
    destination = "0.0.0.0/0" # To any destination
    direction   = "outbound"
    tcp {
      port_max        = 65535 # Maximum TCP port
      port_min        = 1     # Minimum TCP port
      source_port_max = 60000
      source_port_min = 22
    }
  }

  # Inbound rule - allow all TCP traffic
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

# Add an additional rule to the network ACL
resource "ibm_is_network_acl_rule" "isExampleACLRule" {
  network_acl = ibm_is_network_acl.isExampleACL.id
  name        = "isexample-rule"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"

  # ICMP rule
  icmp {
    code = 1
    type = 1
  }
}

# ==========================================================================
# Floating IP Resources
# ==========================================================================

# Create a floating IP for instance1
resource "ibm_is_floating_ip" "floatingip1" {
  name   = "fip1"
  target = ibm_is_instance.instance1.primary_network_interface[0].id # Attach to primary network interface
}

# Create a floating IP for instance2
resource "ibm_is_floating_ip" "floatingip2" {
  name   = "fip2"
  target = ibm_is_instance.instance2.primary_network_interface[0].id
}

# Create a floating IP in a specific zone (not attached to any resource)
resource "ibm_is_floating_ip" "testacc_floatingip" {
  name = "testaccfip"
  zone = ibm_is_subnet.subnet1.zone # Create in the same zone as subnet1
}

# Create a floating IP for bare metal server
resource "ibm_is_floating_ip" "floatingipbms" {
  name = "fip1"
  zone = ibm_is_subnet.subnet1.zone
}

# ==========================================================================
# Security Group Data Sources
# ==========================================================================

# Get information about a specific security group rule
data "ibm_is_security_group_rule" "example" {
  depends_on          = [ibm_is_security_group_rule.exampleudp]
  security_group_rule = ibm_is_security_group_rule.exampleudp.rule_id
  security_group      = ibm_is_security_group.example.id
}

# List all security group rules in a security group
data "ibm_is_security_group_rules" "example" {
  depends_on = [ibm_is_security_group_rule.exampletcp]
}

# List all security groups
data "ibm_is_security_groups" "example" {
}

# ==========================================================================
# Network ACL Data Sources
# ==========================================================================

# Get information about a specific network ACL rule
data "ibm_is_network_acl_rule" "testacc_dsnaclrule" {
  network_acl = ibm_is_network_acl.isExampleACL.id
  name        = ibm_is_network_acl_rule.isExampleACLRule.name
}

# List all rules in a network ACL
data "ibm_is_network_acl_rules" "testacc_dsnaclrules" {
  network_acl = ibm_is_network_acl.isExampleACL.id
}

# Get information about a network ACL by ID
data "ibm_is_network_acl" "is_network_acl" {
  network_acl = ibm_is_network_acl.isExampleACL.id
}

# Get information about a network ACL by name and VPC name
data "ibm_is_network_acl" "is_network_acl1" {
  name     = ibm_is_network_acl.isExampleACL.name
  vpc_name = ibm_is_vpc.vpc1.name
}

# List all network ACLs
data "ibm_is_network_acls" "is_network_acls" {
}

# ==========================================================================
# Load Balancer Resources
# ==========================================================================

# Create a load balancer with private DNS configuration
resource "ibm_is_lb" "example" {
  name    = "example-load-balancer"
  subnets = [ibm_is_subnet.subnet1.id] # Subnets where the load balancer will be deployed
  profile = "network-fixed"            # Type of load balancer profile

  # Private DNS configuration for the load balancer
  dns {
    instance_crn = "crn:v1:staging:public:dns-svcs:global:a/exxxxxxxxxxxxx-xxxxxxxxxxxxxxxxx:5xxxxxxx-xxxxx-xxxxxxxxxxxxxxx-xxxxxxxxxxxxxxx::"
    zone_id      = "bxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx"
  }
}

# Create a simple load balancer
resource "ibm_is_lb" "lb2" {
  name    = "mylb"
  subnets = [ibm_is_subnet.subnet1.id] # Subnets where the load balancer will be deployed
}

# ==========================================================================
# Load Balancer Listener Resources
# ==========================================================================

# Create a load balancer listener for HTTP traffic
resource "ibm_is_lb_listener" "lb_listener2" {
  lb       = ibm_is_lb.lb2.id # Load balancer where the listener will be created
  port     = "9086"           # Port where the listener will listen
  protocol = "http"           # Protocol for the listener
}

# ==========================================================================
# Load Balancer Listener Policy Resources
# ==========================================================================

# Create a load balancer listener policy for traffic redirection
resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
  lb                      = ibm_is_lb.lb2.id
  listener                = ibm_is_lb_listener.lb_listener2.listener_id
  action                  = "redirect" # Policy action - redirect traffic
  priority                = 2          # Priority of the policy
  name                    = "mylistenerpolicy"
  target_http_status_code = 302                      # HTTP status code for redirection
  target_url              = "https://www.google.com" # Target URL for redirection

  # Rule for the policy
  rules {
    condition = "contains" # Condition to match
    type      = "header"   # Type of the condition
    field     = "1"        # Field to check
    value     = "2"        # Value to match
  }
}

# Create a load balancer listener policy rule
resource "ibm_is_lb_listener_policy_rule" "lb_listener_policy_rule" {
  lb        = ibm_is_lb.lb2.id
  listener  = ibm_is_lb_listener.lb_listener2.listener_id
  policy    = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
  condition = "equals"        # Condition to match
  type      = "header"        # Type of the condition
  field     = "MY-APP-HEADER" # Header field to check
  value     = "UpdateVal"     # Value to match
}

# ==========================================================================
# Load Balancer Pool Resources
# ==========================================================================

# Create a load balancer pool with app cookie session persistence
resource "ibm_is_lb_pool" "app_cookie_pool" {
  name                                = "test_pool"
  lb                                  = ibm_is_lb.lb2.id
  algorithm                           = "round_robin" # Load balancing algorithm
  protocol                            = "https"       # Protocol for the pool
  health_delay                        = 60            # Delay between health checks
  health_retries                      = 5             # Number of retries for health checks
  health_timeout                      = 30            # Timeout for health checks
  health_type                         = "https"       # Protocol for health checks
  proxy_protocol                      = "v1"          # Proxy protocol version
  session_persistence_type            = "app_cookie"  # Type of session persistence
  session_persistence_app_cookie_name = "cookie1"     # Name of the application cookie
}

# Create a load balancer pool with HTTP cookie session persistence
resource "ibm_is_lb_pool" "http_cookie_pool" {
  name                     = "test_pool"
  lb                       = ibm_is_lb.lb2.id
  algorithm                = "round_robin"
  protocol                 = "https"
  health_delay             = 60
  health_retries           = 5
  health_timeout           = 30
  health_type              = "https"
  proxy_protocol           = "v1"
  session_persistence_type = "http_cookie" # HTTP cookie session persistence
}

# Create a load balancer pool with source IP session persistence
resource "ibm_is_lb_pool" "source_ip_pool" {
  name                     = "test_pool"
  lb                       = ibm_is_lb.lb2.id
  algorithm                = "round_robin"
  protocol                 = "https"
  health_delay             = 60
  health_retries           = 5
  health_timeout           = 30
  health_type              = "https"
  proxy_protocol           = "v1"
  session_persistence_type = "source_ip" # Source IP session persistence
}

# ==========================================================================
# VPN Gateway Resources
# ==========================================================================

# Create a VPN Gateway in subnet1
resource "ibm_is_vpn_gateway" "VPNGateway1" {
  name   = "vpn1"
  subnet = ibm_is_subnet.subnet1.id # Subnet where the VPN gateway will be created
}

# Create a second VPN Gateway in subnet2
resource "ibm_is_vpn_gateway" "VPNGateway2" {
  name   = "vpn2"
  subnet = ibm_is_subnet.subnet2.id
}

# Create a VPN Gateway for examples
resource "ibm_is_vpn_gateway" "example" {
  name   = "vpn-example"
  subnet = ibm_is_subnet.subnet1.id
}

# Create a VPN Gateway for testing
resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
  name   = "vpn-test"
  subnet = ibm_is_subnet.subnet1.id
}

# ==========================================================================
# VPN Gateway Connection Resources
# ==========================================================================

# Create a VPN Gateway connection (deprecated format - using peer_cidrs and local_cidrs)
# Note: This is the deprecated format and should be avoided for new deployments
resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection1_deprecated" {
  name          = "vpnconn1-deprecated"
  vpn_gateway   = ibm_is_vpn_gateway.VPNGateway1.id
  peer_address  = ibm_is_vpn_gateway.VPNGateway1.public_ip_address
  preshared_key = "VPNDemoPassword"
  local_cidrs   = [ibm_is_subnet.subnet1.ipv4_cidr_block]
  peer_cidrs    = [ibm_is_subnet.subnet2.ipv4_cidr_block]
  ipsec_policy  = ibm_is_ipsec_policy.example.id
}

# Create a VPN Gateway connection (recommended format - using peer and local blocks)
resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection1" {
  name          = "vpnconn1"
  vpn_gateway   = ibm_is_vpn_gateway.VPNGateway1.id
  peer_address  = ibm_is_vpn_gateway.VPNGateway1.public_ip_address
  preshared_key = "VPNDemoPassword"

  # Peer configuration with failover handling between primary and secondary addresses
  peer {
    address    = ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address != "0.0.0.0" ? ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address : ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address2
    peer_cidrs = [ibm_is_subnet.subnet2.ipv4_cidr_block]
  }

  # Local network configuration
  local {
    cidrs = [ibm_is_subnet.subnet1.ipv4_cidr_block]
  }

  ipsec_policy = ibm_is_ipsec_policy.example.id # IPSec policy for the connection
}

# Create a VPN Gateway connection for VPNGateway2 with IKE policy
resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection2" {
  name           = "vpnconn2"
  vpn_gateway    = ibm_is_vpn_gateway.VPNGateway2.id
  peer_address   = ibm_is_vpn_gateway.VPNGateway2.public_ip_address
  preshared_key  = "VPNDemoPassword"
  local_cidrs    = [ibm_is_subnet.subnet2.ipv4_cidr_block]
  peer_cidrs     = [ibm_is_subnet.subnet1.ipv4_cidr_block]
  admin_state_up = true                         # Enable the connection
  ike_policy     = ibm_is_ike_policy.example.id # IKE policy for the connection
}

# Create a VPN connection for examples
resource "ibm_is_vpn_gateway_connection" "example" {
  name          = "vpnconn-example"
  vpn_gateway   = ibm_is_vpn_gateway.example.id
  peer_address  = "169.21.50.5"
  preshared_key = "VPNDemoPassword"
  local_cidrs   = [ibm_is_subnet.subnet1.ipv4_cidr_block]
  peer_cidrs    = ["10.45.0.0/16"]
}

# ==========================================================================
# IPSec and IKE Policies Resources
# ==========================================================================

# Create an IPSec policy for VPN connections
resource "ibm_is_ipsec_policy" "example" {
  name                     = "test-ipsec"
  authentication_algorithm = "sha256"   # Authentication algorithm
  encryption_algorithm     = "aes128"   # Encryption algorithm
  pfs                      = "disabled" # Perfect Forward Secrecy
}

# Create an IKE policy for VPN connections
resource "ibm_is_ike_policy" "example" {
  name                     = "test-ike"
  authentication_algorithm = "sha256" # Authentication algorithm
  encryption_algorithm     = "aes128" # Encryption algorithm
  dh_group                 = 14       # Diffie-Hellman group
  ike_version              = 1        # IKE protocol version
}

# ==========================================================================
# VPN Server Resources
# ==========================================================================

# Create a VPN server
resource "ibm_is_vpn_server" "is_vpn_server" {
  certificate_crn = var.is_certificate_crn # Certificate CRN for TLS

  # Client authentication configuration
  client_authentication {
    method    = "certificate"    # Authentication method
    client_ca = var.is_client_ca # Certificate authority for client authentication
  }

  client_ip_pool         = "10.5.0.0/21"              # IP pool for VPN clients (2048 IPs)
  subnets                = [ibm_is_subnet.subnet1.id] # Subnets for the VPN server
  client_dns_server_ips  = ["192.168.3.4"]            # DNS servers for VPN clients
  client_idle_timeout    = 2800                       # Timeout in seconds for idle clients
  enable_split_tunneling = false                      # All traffic routed through VPN (no split tunneling)
  name                   = "example-vpn-server"
  port                   = 443   # Port for the VPN server
  protocol               = "udp" # Protocol for the VPN server
}

# Create a route for the VPN server
resource "ibm_is_vpn_server_route" "is_vpn_server_route" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
  destination   = "172.16.0.0/16" # Destination network
  action        = "translate"     # Action for the route (translate, deliver, drop)
  name          = "example-vpn-server-route"
}

# ==========================================================================
# Load Balancer Data Sources
# ==========================================================================

# Get information about a load balancer listener
data "ibm_is_lb_listener" "is_lb_listener" {
  lb          = ibm_is_lb.lb2.id
  listener_id = ibm_is_lb_listener.lb_listener2.listener_id
}

# List all listeners for a load balancer
data "ibm_is_lb_listeners" "is_lb_listeners" {
  lb = ibm_is_lb.lb2.id
}

# Get information about a listener policy
data "ibm_is_lb_listener_policy" "is_lb_listener_policy" {
  lb        = ibm_is_lb.lb2.id
  listener  = ibm_is_lb_listener.lb_listener2.listener_id
  policy_id = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
}

# List all policies for a listener
data "ibm_is_lb_listener_policies" "is_lb_listener_policies" {
  lb       = ibm_is_lb.lb2.id
  listener = ibm_is_lb_listener.lb_listener2.listener_id
}

# Get information about a listener policy rule
data "ibm_is_lb_listener_policy_rule" "is_lb_listener_policy_rule" {
  lb       = ibm_is_lb.lb2.id
  listener = ibm_is_lb_listener.lb_listener2.listener_id
  policy   = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
  rule     = ibm_is_lb_listener_policy_rule.lb_listener_policy_rule.rule
}

# List all rules for a listener policy
data "ibm_is_lb_listener_policy_rules" "is_lb_listener_policy_rules" {
  lb       = ibm_is_lb.lb2.id
  listener = ibm_is_lb_listener.lb_listener2.listener_id
  policy   = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
}

# ==========================================================================
# VPN Gateway Data Sources
# ==========================================================================

# Get information about a VPN gateway
data "ibm_is_vpn_gateway" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
}

# Get information about a VPN gateway by name
data "ibm_is_vpn_gateway" "example-1" {
  vpn_gateway_name = ibm_is_vpn_gateway.example.name
}

# Get information about a VPN gateway connection
data "ibm_is_vpn_gateway_connection" "example" {
  vpn_gateway            = ibm_is_vpn_gateway.example.id
  vpn_gateway_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
}

# Get information about a VPN gateway connection by name
data "ibm_is_vpn_gateway_connection" "example-1" {
  vpn_gateway                 = ibm_is_vpn_gateway.example-1.id
  vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.example.name
}

# Get information about a VPN gateway connection with gateway name
data "ibm_is_vpn_gateway_connection" "example-2" {
  vpn_gateway_name       = ibm_is_vpn_gateway.example.name
  vpn_gateway_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
}

# Get information about a VPN gateway connection with gateway name and connection name
data "ibm_is_vpn_gateway_connection" "example-3" {
  vpn_gateway_name            = ibm_is_vpn_gateway.example.name
  vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.example.name
}

# ==========================================================================
# IKE and IPSec Policy Data Sources
# ==========================================================================

# List all IKE policies
data "ibm_is_ike_policies" "example" {
}

# List all IPSec policies
data "ibm_is_ipsec_policies" "example" {
}

# Get information about an IKE policy by ID
data "ibm_is_ike_policy" "example" {
  ike_policy = ibm_is_ike_policy.example.id
}

# Get information about an IPSec policy by ID
data "ibm_is_ipsec_policy" "example1" {
  ipsec_policy = ibm_is_ipsec_policy.example.id
}

# Get information about an IKE policy by name
data "ibm_is_ike_policy" "example2" {
  name = "my-ike-policy"
}

# Get information about an IPSec policy by name
data "ibm_is_ipsec_policy" "example3" {
  name = "my-ipsec-policy"
}

# ==========================================================================
# VPN Server Data Sources
# ==========================================================================

# Get information about a VPN server
data "ibm_is_vpn_server" "is_vpn_server" {
  identifier = ibm_is_vpn_server.is_vpn_server.vpn_server
}

# List all VPN servers
data "ibm_is_vpn_servers" "is_vpn_servers" {
}

# List all routes for a VPN server
data "ibm_is_vpn_server_routes" "is_vpn_server_routes" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
}

# Get information about a VPN server route
data "ibm_is_vpn_server_route" "is_vpn_server_route" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
  identifier    = ibm_is_vpn_server_route.is_vpn_server_route.vpn_route
}

# List all clients connected to a VPN server
data "ibm_is_vpn_server_clients" "is_vpn_server_clients" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
}

# Get information about a specific VPN server client
data "ibm_is_vpn_server_client" "is_vpn_server_client" {
  vpn_server_id = ibm_is_vpn_server.is_vpn_server.vpn_server
  identifier    = "0726-61b2f53f-1e95-42a7-94ab-55de8f8cbdd5" # Client ID
}

# ==========================================================================
# SSH Key Resources
# ==========================================================================

# Create an SSH key for accessing instances
resource "ibm_is_ssh_key" "sshkey" {
  name       = "ssh1"
  public_key = file(var.ssh_public_key) # Read public key from a file
}

# Create an SSH key with RSA key type
resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "${var.name}-ssh"
  public_key = file("~/.ssh/id_rsa.pub") # Read public key from file
}

# Create an SSH key with ED25519 key type
resource "ibm_is_ssh_key" "is_key" {
  name       = "my-key"
  public_key = file("~/.ssh/id_ed25519.pub")
  type       = "ed25519" # Explicitly set key type
}

# ==========================================================================
# Instance Template Resources
# ==========================================================================

# Create an instance template with a new volume
resource "ibm_is_instance_template" "instancetemplate1" {
  name    = "testtemplate"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315" # OS image ID
  profile = "bx2-8x32"                                  # Instance profile (8 vCPU, 32 GB RAM)

  # Primary network interface configuration
  primary_network_interface {
    subnet            = ibm_is_subnet.subnet2.id # Subnet for the interface
    allow_ip_spoofing = true                     # Allow IP spoofing
  }

  vpc  = ibm_is_vpc.vpc2.id         # VPC for the instance
  zone = "us-south-2"               # Zone for the instance
  keys = [ibm_is_ssh_key.sshkey.id] # SSH key for access

  # Boot volume configuration
  boot_volume {
    name                             = "testbootvol"
    delete_volume_on_instance_delete = true # Delete boot volume when instance is deleted
  }

  # Additional volume attachment with volume prototype
  volume_attachments {
    delete_volume_on_instance_delete = true
    name                             = "volatt-01"
    volume_prototype {
      iops     = 3000              # IOPS for the volume
      profile  = "general-purpose" # Volume profile
      capacity = 200               # Volume size in GB
    }
  }
}

# Create an instance template with an existing volume
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

  # Volume attachment using an existing volume
  volume_attachments {
    delete_volume_on_instance_delete = true
    name                             = "volatt-01"
    volume                           = ibm_is_volume.vol1.id # Reference to existing volume
  }
}

# Create an instance template with a primary network attachment using VNI
resource "ibm_is_instance_template" "ins_temp" {
  name    = "${var.name}-vsi2"
  profile = "bx2-2x8"
  image   = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"

  # Use a virtual network interface for the primary network attachment
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

# Create an instance template for GPU cluster nodes
resource "ibm_is_instance_template" "is_instance_template" {
  name    = "${var.prefix}-cluster-it"
  image   = data.ibm_is_image.is_image.id
  profile = "gx3d-160x1792x8h100" # GPU instance profile

  # Primary network attachment
  primary_network_attachment {
    name = "my-pna-it"
    virtual_network_interface {
      auto_delete = true
      subnet      = ibm_is_subnet.is_subnet.id
    }
  }

  # Cluster network attachments - 8 attachments for 8 GPUs
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

# ==========================================================================
# Virtual Server Instance Resources
# ==========================================================================

# Create a basic virtual server instance in vpc1/subnet1
resource "ibm_is_instance" "instance1" {
  name    = "instance1"
  image   = var.image   # OS image ID from variable
  profile = var.profile # Instance profile from variable

  primary_network_interface {
    subnet = ibm_is_subnet.subnet1.id # Primary subnet for the instance
  }

  vpc  = ibm_is_vpc.vpc1.id         # VPC for the instance
  zone = var.zone1                  # Zone for the instance
  keys = [ibm_is_ssh_key.sshkey.id] # SSH key for access
}

# Create an instance on a dedicated host in vpc2/subnet2
resource "ibm_is_instance" "instance2" {
  name    = "instance2"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  dedicated_host = ibm_is_dedicated_host.is_dedicated_host.id # Place instance on specific dedicated host
  vpc            = ibm_is_vpc.vpc2.id
  zone           = var.zone2
  keys           = [ibm_is_ssh_key.sshkey.id]
}

# Create an instance on a dedicated host group in vpc2/subnet2
resource "ibm_is_instance" "instance3" {
  name    = "instance3"
  image   = var.image
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  dedicated_host_group = ibm_is_dedicated_host_group.dh_group01.id # Place instance in a dedicated host group
  vpc                  = ibm_is_vpc.vpc2.id
  zone                 = var.zone2
  keys                 = [ibm_is_ssh_key.sshkey.id]
}

# Create an instance with attached volumes
resource "ibm_is_instance" "instance4" {
  name    = "instance4"
  image   = var.image
  profile = var.profile

  volumes = [ibm_is_volume.vol3.id] # Attach additional volumes to the instance

  primary_network_interface {
    subnet = ibm_is_subnet.subnet1.id
  }

  vpc  = ibm_is_vpc.vpc1.id
  zone = var.zone1
  keys = [ibm_is_ssh_key.sshkey.id]
}

# Create an instance with boot volume from snapshot
resource "ibm_is_instance" "instance5" {
  name    = "instance5"
  profile = var.profile

  # Boot from a snapshot instead of an image
  boot_volume {
    name     = "boot-restore"
    snapshot = ibm_is_snapshot.b_snapshot.id # Use a snapshot for boot volume
  }

  auto_delete_volume = true # Delete volumes when instance is deleted

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}

# Create an instance using an instance template
resource "ibm_is_instance" "instance6" {
  name              = "instance4"
  instance_template = ibm_is_instance_template.instancetemplate1.id # Create from template
}

# Create an instance with reserved IP address
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
      address     = "10.0.0.5" # Specific IP address for the interface
      auto_delete = true       # Delete the reserved IP when the instance is deleted
    }
    name   = "test-reserved-ip"
    subnet = ibm_is_subnet.subnet2.id
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}

# Create an instance from catalog offering
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

  # Use a catalog offering instead of a standard image
  catalog_offering {
    version_crn = data.ibm_is_images.imageslist.images.0.catalog_offering.0.version.0.crn
    plan_crn    = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}

# Create an instance with a primary network attachment using VNI
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

# Create a cluster instance with built-in network interfaces
resource "ibm_is_instance" "is_instance" {
  name    = "${var.prefix}-cluster-ins"
  image   = data.ibm_is_image.is_image.id
  profile = "gx3d-160x1792x8h100" # GPU instance profile

  # Primary network interface in regular subnet
  primary_network_interface {
    subnet = ibm_is_subnet.is_subnet.id
  }

  # Add cluster network attachments - one for each GPU
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

# ==========================================================================
# Instance Network Interface Management
# ==========================================================================

# Create a network interface on an existing instance
resource "ibm_is_instance_network_interface" "is_instance_network_interface" {
  instance             = ibm_is_instance.instance1.id
  subnet               = ibm_is_subnet.subnet1.id
  allow_ip_spoofing    = true # Allow IP spoofing on the interface
  name                 = "my-network-interface"
  primary_ipv4_address = "10.0.0.5" # Specific IP address for the interface
}

# Create a network attachment for an instance
resource "ibm_is_instance_network_attachment" "ina" {
  instance = ibm_is_instance.ins.id
  name     = "viability-undecided-jalapeno-unbuilt"
  virtual_network_interface {
    id = ibm_is_virtual_network_interface.testacc_vni2.id
  }
}

# ==========================================================================
# Instance Volume Attachment Management
# ==========================================================================

# Attach an existing volume to an instance
resource "ibm_is_instance_volume_attachment" "att1" {
  instance                           = ibm_is_instance.instance5.id
  volume                             = ibm_is_volume.vol5.id # Existing volume
  name                               = "vol-att-1"
  delete_volume_on_attachment_delete = false # Keep volume when attachment is deleted
  delete_volume_on_instance_delete   = false # Keep volume when instance is deleted
}

# Create a new volume from snapshot and attach to an instance
resource "ibm_is_instance_volume_attachment" "att2" {
  instance                           = ibm_is_instance.instance5.id
  name                               = "vol-att-2"
  profile                            = "general-purpose"             # Volume profile
  snapshot                           = ibm_is_snapshot.d_snapshot.id # Create from snapshot
  delete_volume_on_instance_delete   = true                          # Delete volume when instance is deleted
  delete_volume_on_attachment_delete = true                          # Delete volume when attachment is deleted
  volume_name                        = "vol4-restore"                # Name for the new volume
}

# ==========================================================================
# Instance Disk Management
# ==========================================================================

# Update disk names on an instance
resource "ibm_is_instance_disk_management" "disks" {
  instance = ibm_is_instance.instance1.id
  disks {
    name = "mydisk01"
    id   = ibm_is_instance.instance1.disks.0.id # Reference to existing disk
  }
}

# ==========================================================================
# Instance Cluster Network Attachment Management
# ==========================================================================

# Stop the instance before adding more attachments
resource "ibm_is_instance_action" "is_instance_stop_before" {
  action   = "stop"
  instance = ibm_is_instance.is_instance.id
}

# Add additional cluster network attachments
# These additions are for instances that need more than 8 interfaces
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

# Start the instance after adding attachments
resource "ibm_is_instance_action" "is_instance_start_after" {
  action   = "start"
  instance = ibm_is_instance.is_instance.id
}

# ==========================================================================
# SSH Key Data Sources
# ==========================================================================

# List all SSH keys
data "ibm_is_ssh_keys" "example" {
}

# List all SSH keys in a resource group
data "ibm_is_ssh_keys" "example_by_rg" {
  resource_group = data.ibm_resource_group.default.id # Filter by resource group
}

# ==========================================================================
# Instance Template Data Sources
# ==========================================================================

# Get information about an instance template by ID
data "ibm_is_instance_template" "instancetemplates" {
  identifier = ibm_is_instance_template.instancetemplate2.id
}

# Get information about an instance template by name
data "ibm_is_instance_template" "is_instance_template_instance" {
  name = ibm_is_instance_template.is_instance_template.name
}

# ==========================================================================
# Instance Data Sources
# ==========================================================================

# Get information about an instance including connection info
data "ibm_is_instance" "ds_instance" {
  name        = ibm_is_instance.instance1.name
  private_key = file("~/.ssh/id_rsa") # Private key for SSH connection
  passphrase  = ""                    # Passphrase for the private key if needed
}

# Get information about an instance
data "ibm_is_instance" "is_instance_instance" {
  name = ibm_is_instance.is_instance.name
}

# List all instances
data "ibm_is_instances" "is_instances_instance" {
}

# Filter instances by name
data "ibm_is_instances" "is_instances_filtered" {
  name = var.is_instances_name
}

# ==========================================================================
# Instance Network Interface Data Sources
# ==========================================================================

# Get information about an instance network interface
data "ibm_is_instance_network_interface" "is_instance_network_interface" {
  instance_name          = ibm_is_instance.instance1.name
  network_interface_name = ibm_is_instance_network_interface.is_instance_network_interface.name
}

# List all network interfaces for an instance
data "ibm_is_instance_network_interfaces" "is_instance_network_interfaces" {
  instance_name = ibm_is_instance.instance1.name
}

# ==========================================================================
# Instance Volume Attachment Data Sources
# ==========================================================================

# Get information about a volume attachment
data "ibm_is_instance_volume_attachment" "ds_vol_att" {
  instance = ibm_is_instance.instance5.id
  name     = ibm_is_instance_volume_attachment.att2.name
}

# List all volume attachments for an instance
data "ibm_is_instance_volume_attachment" "ds_vol_atts" {
  instance = ibm_is_instance.instance5.id
}

# ==========================================================================
# Instance Disk Data Sources
# ==========================================================================

# List all disks for an instance
data "ibm_is_instance_disks" "disk1" {
  instance = ibm_is_instance.instance1.id
}

# Get information about a specific disk
data "ibm_is_instance_disk" "disk1" {
  instance = ibm_is_instance.instance1.id
  disk     = data.ibm_is_instance_disks.disk1.disks.0.id
}

# ==========================================================================
# Instance Cluster Network Attachment Data Sources
# ==========================================================================

# Get information about a cluster network attachment
data "ibm_is_instance_cluster_network_attachment" "is_instance_cluster_network_attachment_instance" {
  instance_id                            = ibm_is_instance.is_instance.id
  instance_cluster_network_attachment_id = ibm_is_instance.is_instance.cluster_network_attachments.0.id
}

# List all cluster network attachments for an instance
data "ibm_is_instance_cluster_network_attachments" "is_instance_cluster_network_attachments_instance" {
  instance_id = ibm_is_instance.is_instance.id
}

# ==========================================================================
# Instance Profile Data Sources
# ==========================================================================

# Get information about a specific instance profile
data "ibm_is_instance_profile" "is_instance_profile_instance" {
  name = "gx3d-160x1792x8h100" # 160 vCPU, 1792 GB RAM, 8x H100 GPUs
}

# List all instance profiles
data "ibm_is_instance_profiles" "is_instance_profiles_instance" {
}

# ==========================================================================
# Bare Metal Server Resources
# ==========================================================================

# Create a basic bare metal server
resource "ibm_is_bare_metal_server" "bms" {
  profile = "bx2-metal-192x768" # Bare metal server profile
  name    = "my-bms"
  image   = "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e" # OS image
  zone    = "us-south-3"                                # Zone for the server
  keys    = [ibm_is_ssh_key.sshkey.id]                  # SSH key for access

  primary_network_interface {
    subnet = ibm_is_subnet.subnet1.id # Primary subnet for the server
  }

  vpc = ibm_is_vpc.vpc1.id # VPC for the server
}

# Create a bare metal server with a primary network attachment using VNI
resource "ibm_is_bare_metal_server" "testacc_bms" {
  profile = "cx2-metal-96x192"
  name    = "${var.name}-bms"
  image   = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"
  zone    = "${var.region}-2"
  keys    = [ibm_is_ssh_key.testacc_sshkey.id]

  # Primary network attachment using a virtual network interface
  primary_network_attachment {
    name = "vni-221"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.testacc_vni.id
    }
    allowed_vlans = [100, 102] # VLANs allowed on this attachment
  }

  vpc = ibm_is_vpc.testacc_vpc.id
}

# ==========================================================================
# Bare Metal Server Management Resources
# ==========================================================================

# Update a disk name on a bare metal server
resource "ibm_is_bare_metal_server_disk" "this" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  disk              = ibm_is_bare_metal_server.bms.disks.0.id # Reference to existing disk
  name              = "bms-disk-update"                       # New name for the disk
}

# Perform an action on a bare metal server (stop)
resource "ibm_is_bare_metal_server_action" "this" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  action            = "stop" # Action to perform (stop, start, restart)
  stop_type         = "hard" # Hard stop (vs. soft stop)
}

# ==========================================================================
# Bare Metal Server Network Resources
# ==========================================================================

# Add a network interface to a bare metal server
resource "ibm_is_bare_metal_server_network_interface" "bms_nic" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  subnet            = ibm_is_subnet.subnet1.id
  name              = "eth2"
  allow_ip_spoofing = true       # Allow IP spoofing
  allowed_vlans     = [101, 102] # VLANs allowed on this interface
}

# Add a network interface with a specific VLAN
resource "ibm_is_bare_metal_server_network_interface" "bms_nic2" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  subnet            = ibm_is_subnet.subnet1.id
  name              = "eth2"
  allow_ip_spoofing = true
  vlan              = 101 # VLAN ID for the interface
}

# Add a floating interface (VLAN) to a bare metal server
resource "ibm_is_bare_metal_server_network_interface_allow_float" "bms_vlan_nic" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  subnet            = ibm_is_subnet.subnet1.id
  name              = "eth2"
  vlan              = 102 # VLAN ID for the interface
}

# Create a network attachment for a bare metal server
resource "ibm_is_bare_metal_server_network_attachment" "na" {
  bare_metal_server = ibm_is_bare_metal_server.testacc_bms.id
  # interface_type = "vlan"    # Interface type - commented out in original
  vlan = 100 # VLAN ID for the attachment
}

# Create a network attachment with multiple allowed VLANs
resource "ibm_is_bare_metal_server_network_attachment" "na2" {
  bare_metal_server = ibm_is_bare_metal_server.testacc_bms.id
  # interface_type = "pci"    # Interface type - commented out in original
  allowed_vlans = [200, 202] # VLANs allowed on this attachment
}

# Associate a floating IP with a bare metal server network interface
resource "ibm_is_bare_metal_server_network_interface_floating_ip" "bms_nic_fip" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  network_interface = ibm_is_bare_metal_server_network_interface.bms_nic2.id
  floating_ip       = ibm_is_floating_ip.testacc_fip.id # Floating IP to associate
}

# ==========================================================================
# Bare Metal Server Profile Data Sources
# ==========================================================================

# List all bare metal server profiles
data "ibm_is_bare_metal_server_profiles" "this" {
}

# Get information about a specific bare metal server profile
data "ibm_is_bare_metal_server_profile" "this" {
  name = data.ibm_is_bare_metal_server_profiles.this.profiles.0.name
}

# ==========================================================================
# Bare Metal Server Data Sources
# ==========================================================================

# Get information about a bare metal server
data "ibm_is_bare_metal_server" "this" {
  identifier = ibm_is_bare_metal_server.this.id
}

# List all bare metal servers
data "ibm_is_bare_metal_servers" "this" {
}

# Get bare metal server initialization data
data "ibm_is_bare_metal_server_initialization" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
}

# ==========================================================================
# Bare Metal Server Disk Data Sources
# ==========================================================================

# Get information about a specific disk
data "ibm_is_bare_metal_server_disk" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
  disk              = ibm_is_bare_metal_server.this.disks.0.id
}

# List all disks for a bare metal server
data "ibm_is_bare_metal_server_disks" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
}

# ==========================================================================
# Bare Metal Server Network Interface Data Sources
# ==========================================================================

# Get information about a network interface
data "ibm_is_bare_metal_server_network_interface" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
  network_interface = ibm_is_bare_metal_server.this.primary_network_interface.id
}

# List all network interfaces for a bare metal server
data "ibm_is_bare_metal_server_network_interfaces" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
}

# Get information about a floating IP associated with a network interface
data "ibm_is_bare_metal_server_network_interface_floating_ip" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
  network_interface = ibm_is_bare_metal_server.this.primary_network_interface[0].id
  floating_ip       = ibm_is_floating_ip.floatingipbms.id
}

# List all floating IPs associated with a network interface
data "ibm_is_bare_metal_server_network_interface_floating_ips" "this" {
  bare_metal_server = ibm_is_bare_metal_server.this.id
  network_interface = ibm_is_bare_metal_server.this.primary_network_interface[0].id
}

# ==========================================================================
# Block Storage Volume Resources
# ==========================================================================

# Create a volume with predefined IOPS tier
resource "ibm_is_volume" "vol1" {
  name    = "vol1"
  profile = "10iops-tier" # Predefined IOPS tier
  zone    = var.zone1
}

# Create a volume with custom IOPS
resource "ibm_is_volume" "vol2" {
  name     = "vol2"
  profile  = "custom" # Custom profile for specifying IOPS
  zone     = var.zone1
  iops     = 1000 # Custom IOPS value
  capacity = 200  # Volume size in GB
}

# Create a volume to attach to instance4
resource "ibm_is_volume" "vol3" {
  name    = "vol3"
  profile = "10iops-tier"
  zone    = var.zone1
}

# Create a volume for instance5
resource "ibm_is_volume" "vol5" {
  name    = "vol5"
  profile = "10iops-tier"
  zone    = "us-south-2"
  tags    = ["tag1"] # Tag the volume
}

# ==========================================================================
# Snapshot Resources
# ==========================================================================

# Create a snapshot from boot volume with clone
resource "ibm_is_snapshot" "b_snapshot" {
  name          = "my-snapshot-boot"
  source_volume = ibm_is_instance.instance4.volume_attachments[0].volume_id # Source is boot volume
  clones        = [var.zone1]                                               # Create clones in the specified zone
  tags          = ["tags1"]                                                 # Add tags
}

# Create a snapshot from data volume
resource "ibm_is_snapshot" "d_snapshot" {
  name          = "my-snapshot-data"
  source_volume = ibm_is_instance.instance4.volume_attachments[1].volume_id # Source is data volume
  tags          = ["tags1"]
}

# Create a snapshot copy in another region
resource "ibm_is_snapshot" "b_snapshot_copy" {
  provider            = ibm.eu-de # Use eu-de provider
  name                = "my-snapshot-boot-copy"
  source_snapshot_crn = ibm_is_snapshot.b_snapshot.crn # Source is another snapshot
}

# Create a snapshot consistency group for instance snapshots
resource "ibm_is_snapshot_consistency_group" "is_snapshot_consistency_group_instance" {
  delete_snapshots_on_delete = true # Delete snapshots when group is deleted

  # First snapshot in the group
  snapshots {
    name          = "exmaple-snapshot"
    source_volume = ibm_is_instance.instance.volume_attachments[0].volume_id
  }

  # Second snapshot in the group
  snapshots {
    name          = "example-snapshot-1"
    source_volume = ibm_is_instance.instance.volume_attachments[1].volume_id
  }

  name = "example-snapshot-consistency-group"
}

# ==========================================================================
# File Share Resources
# ==========================================================================

# Create a file share
resource "ibm_is_share" "share" {
  zone        = "us-south-1" # Zone for the share
  size        = 30000        # Size in GB
  name        = "my-share"
  profile     = "dp2" # Performance profile
  tags        = ["share1", "share3"]
  access_tags = ["access:dev"] # Access control tags
}

# Create a replica of a file share
resource "ibm_is_share" "sharereplica" {
  zone                  = "us-south-2" # Zone for the replica
  name                  = "my-share-replica"
  profile               = "dp2"                 # Performance profile
  replication_cron_spec = "0 */5 * * *"         # Replication schedule (every 5 hours)
  source_share          = ibm_is_share.share.id # Source share to replicate
  tags                  = ["share1", "share3"]
  access_tags           = ["access:dev"]
}

# Create a cross-regional share
resource "ibm_is_share" "cross_regional_share" {
  zone                  = "us-east-1"
  source_share_crn      = "crn:v1:staging:public:is:us-south-1:a/efe5afc483594adaa8325e2b4d1290df::share:r134-d8c8821c-a227-451d-a9ed-0c0cd2358829"
  encryption_key        = "crn:v1:staging:public:kms:us-south:a/efe5afc483594adaa8325e2b4d1290df:1be45161-6dae-44ca-b248-837f98004057:key:3dd21cc5-cc20-4f7c-bc62-8ec9a8a3d1bd"
  replication_cron_spec = "5 * * * *" # Replication schedule (every hour at 5 minutes past)
  name                  = "tfp-temp-crr"
  profile               = "dp2"
}

# Create a mount target for a file share
resource "ibm_is_share_mount_target" "is_share_mount_target" {
  share = ibm_is_share.share.id # Share to mount
  vpc   = ibm_is_vpc.vpc1.id    # VPC where the mount target will be created
  name  = "my-share-target-1"
}

# Create a snapshot of a file share
resource "ibm_is_share_snapshot" "example" {
  name  = "my-example-share-snapshot"
  share = ibm_is_share.share.id
  tags  = ["my-example-share-snapshot-tag"]
}

# ==========================================================================
# Image Resources
# ==========================================================================

# Create an image from a URL in Cloud Object Storage
resource "ibm_is_image" "image1" {
  href             = var.image_cos_url # URL to the image in Cloud Object Storage
  name             = "my-img-1"
  operating_system = var.image_operating_system # OS type
}

# Create an image from a volume
resource "ibm_is_image" "image2" {
  source_volume = data.ibm_is_instance.instance1.volume_attachments.0.volume_id # Source volume
  name          = "my-img-1"
}

# Mark an image as deprecated
resource "ibm_is_image_deprecate" "example" {
  image = ibm_is_image.image1.id # Image to deprecate
}

# Mark an image as obsolete
resource "ibm_is_image_obsolete" "example" {
  image = ibm_is_image.image1.id # Image to mark as obsolete
}

# Create an image export job to export an image to Cloud Object Storage
resource "ibm_is_image_export_job" "example" {
  image = ibm_is_image.image1.id
  name  = "my-image-export"
  storage_bucket {
    name = "bucket-27200-lwx4cfvcue" # Destination bucket
  }
}

# ==========================================================================
# Backup Policy Resources
# ==========================================================================

# Create a backup policy for volumes
resource "ibm_is_backup_policy" "volume_backup_policy" {
  match_user_tags     = ["tag1"] # Tag to match for backup
  name                = "my-backup-policy"
  match_resource_type = "volume" # Backup for volumes
}

# Create a backup policy for instances
resource "ibm_is_backup_policy" "instance_backup_policy" {
  match_user_tags     = ["tag1"]
  name                = "my-backup-policy-instance"
  match_resource_type = "instance"                      # Backup for instances
  included_content    = ["boot_volume", "data_volumes"] # Include both boot and data volumes
}

# Create a backup policy plan
resource "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
  backup_policy_id = ibm_is_backup_policy.volume_backup_policy.id
  cron_spec        = "30 09 * * *" # Daily at 9:30 AM
  active           = false         # Not active initially
  attach_user_tags = ["tag2"]      # Tags to attach to backups
  copy_user_tags   = true          # Copy tags from source

  # Deletion trigger for old backups
  deletion_trigger {
    delete_after      = 20   # Delete after 20 days
    delete_over_count = "20" # Keep only last 20 backups
  }

  name = "my-backup-policy-plan-1"
}

# Create a backup policy plan with clones
resource "ibm_is_backup_policy_plan" "is_backup_policy_plan_clone" {
  backup_policy_id = ibm_is_backup_policy.volume_backup_policy.id
  cron_spec        = "30 09 * * *"
  active           = false
  attach_user_tags = ["tag2"]
  copy_user_tags   = true

  deletion_trigger {
    delete_after      = 20
    delete_over_count = "20"
  }

  name = "my-backup-policy-plan-1"

  # Clone policy for creating backup copies in other zones
  clone_policy {
    zones         = ["us-south-1", "us-south-2"] # Zones to create clones in
    max_snapshots = 3                            # Maximum number of snapshot clones
  }
}

# Create an enterprise backup policy
resource "ibm_is_backup_policy" "ent-baas-example" {
  match_user_tags = ["tag1"]
  name            = "example-enterprise-backup-policy"

  # Scope for enterprise backup
  scope {
    crn = "crn:v1:bluemix:public:is:us-south:a/123456::reservation:7187-ba49df72-37b8-43ac-98da-f8e029de0e63"
  }
}

# ==========================================================================
# Volume Data Sources
# ==========================================================================

# Get information about a volume profile
data "ibm_is_volume_profile" "volprofile" {
  name = "general-purpose"
}

# List all volume profiles
data "ibm_is_volume_profiles" "volprofiles" {
}

# List all volumes
data "ibm_is_volumes" "example-volumes" {
}

# List all volumes with a specific name
data "ibm_is_volumes" "example-by-name" {
  volume_name = "worrier-mailable-timpani-scowling" # Filter volumes by name
}

# List all volumes in a zone
data "ibm_is_volumes" "example-by-zone" {
  zone_name = "us-south-1" # Filter volumes by zone
}

# ==========================================================================
# Snapshot Data Sources
# ==========================================================================

# Get information about a snapshot by name
data "ibm_is_snapshot" "ds_snapshot" {
  name = "my-snapshot-boot"
}

# List all snapshots
data "ibm_is_snapshots" "ds_snapshots" {
}

# List all clones of a snapshot
data "ibm_is_snapshot_clones" "ds_snapshot_clones" {
  snapshot = ibm_is_snapshot.b_snapshot.id
}

# Get information about a specific snapshot clone
data "ibm_is_snapshot_clones" "ds_snapshot_clone" {
  snapshot = ibm_is_snapshot.b_snapshot.id
  zone     = var.zone1
}

# Get information about a snapshot consistency group by ID
data "ibm_is_snapshot_consistency_group" "by_id" {
  identifier = ibm_is_snapshot_consistency_group.is_snapshot_consistency_group_instance.id
}

# Get information about a snapshot consistency group by name
data "ibm_is_snapshot_consistency_group" "by_name" {
  name = "example-snapshot-consistency-group"
}

# List snapshot consistency groups by name
data "ibm_is_snapshot_consistency_groups" "is_snapshot_consistency_group_instance" {
  depends_on = [ibm_is_snapshot_consistency_group.is_snapshot_consistency_group_instance]
  name       = "example-snapshot-consistency-group"
}

# ==========================================================================
# File Share Data Sources
# ==========================================================================

# Get information about a share mount target
data "ibm_is_share_mount_target" "is_share_mount_target" {
  share        = ibm_is_share.share.id
  mount_target = ibm_is_share_mount_target.is_share_target.mount_target
}

# List all mount targets for a share
data "ibm_is_share_mount_targets" "is_share_mount_targets" {
  share = ibm_is_share.share.id
}

# Get information about a share
data "ibm_is_share" "is_share" {
  share = ibm_is_share.share.id
}

# List all shares
data "ibm_is_shares" "is_shares" {
}

# List all snapshots for a share
data "ibm_is_share_snapshots" "example" {
  share = ibm_is_share.share.id
}

# List all snapshots across all shares
data "ibm_is_share_snapshots" "example1" {
}

# Get information about a specific share snapshot
data "ibm_is_share_snapshot" "example1" {
  share          = ibm_is_share.share.id
  share_snapshot = ibm_is_share_snapshot.example.share_snapshot
}

# ==========================================================================
# Image Data Sources
# ==========================================================================

# Get information about an image by name
data "ibm_is_image" "dsimage" {
  name = ibm_is_image.image1.name
}

# List all images
data "ibm_is_images" "dsimages" {
}

# List all catalog-managed images
data "ibm_is_images" "imageslist" {
  catalog_managed = true # Only return catalog-managed images
}

# Get Ubuntu image for examples
data "ibm_is_image" "is_image" {
  name = "ibm-ubuntu-20-04-6-minimal-amd64-6"
}

# List all image export jobs for an image
data "ibm_is_image_export_jobs" "example" {
  image = ibm_is_image_export_job.example.image
}

# Get information about a specific image export job
data "ibm_is_image_export_job" "example" {
  image            = ibm_is_image_export_job.example.image
  image_export_job = ibm_is_image_export_job.example.image_export_job
}

# ==========================================================================
# Operating System Data Sources
# ==========================================================================

# Get information about a specific operating system
data "ibm_is_operating_system" "os" {
  name = "red-8-amd64" # Get information about a specific OS
}

# List all available operating systems
data "ibm_is_operating_systems" "oslist" {
}

# ==========================================================================
# Backup Policy Data Sources
# ==========================================================================

# List all backup policies
data "ibm_is_backup_policies" "is_backup_policies" {
}

# Get information about a backup policy by name
data "ibm_is_backup_policy" "is_backup_policy" {
  name = "my-backup-policy"
}

# Get information about an enterprise backup policy
data "ibm_is_backup_policy" "enterprise_backup" {
  name = ibm_is_backup_policy.ent-baas-example.name
}

# List all plans for a backup policy
data "ibm_is_backup_policy_plans" "is_backup_policy_plans" {
  backup_policy_id = ibm_is_backup_policy.volume_backup_policy.id
}

# Get information about a backup policy plan
data "ibm_is_backup_policy_plan" "is_backup_policy_plan" {
  backup_policy_id = ibm_is_backup_policy.volume_backup_policy.id
  name             = "my-backup-policy-plan"
}

# Get information about a backup policy job
data "ibm_is_backup_policy_job" "is_backup_policy_job" {
  backup_policy_id = ibm_is_backup_policy.volume_backup_policy.id
  identifier       = ""
}

# List all jobs for a backup policy
data "ibm_is_backup_policy_jobs" "is_backup_policy_jobs" {
  backup_policy_plan_id = ibm_is_backup_policy.volume_backup_policy.backup_policy_plan_id
  backup_policy_id      = ibm_is_backup_policy.volume_backup_policy.id
}

# ==========================================================================
# Cluster Network Resources
# ==========================================================================

# Create a cluster network for high-performance computing
resource "ibm_is_cluster_network" "is_cluster_network_instance" {
  name           = "${var.prefix}-cluster"
  profile        = "h100" # High-performance cluster network profile
  resource_group = var.is_instances_resource_group_id

  # Define subnet prefix for the cluster
  subnet_prefixes {
    cidr = "10.1.0.0/24" # 256 IP addresses
  }

  # Link to VPC
  vpc {
    id = ibm_is_vpc.is_vpc.id
  }

  zone = "${var.region}-3"
}

# Create an updated cluster network
resource "ibm_is_cluster_network" "is_cluster_network_instance_updated" {
  name    = "${var.prefix}-cluster-updated"
  profile = "h100"

  # Define subnet prefix
  subnet_prefixes {
    cidr = "10.0.0.0/24"
  }

  # Link to VPC
  vpc {
    id = ibm_is_vpc.is_vpc.id
  }

  zone = ibm_is_subnet.is_subnet.zone
}

# ==========================================================================
# Cluster Network Subnet Resources
# ==========================================================================

# Create a subnet in the cluster network
resource "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
  cluster_network_id       = ibm_is_cluster_network.is_cluster_network_instance.id
  name                     = "${var.prefix}-cluster-subnet"
  total_ipv4_address_count = 64 # 64 IP addresses
}

# Reserve an IP in the cluster network subnet
resource "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
  cluster_network_id        = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
  address                   = "10.1.0.4" # Specific IP address
  name                      = "${var.prefix}-cluster-subnet-r-ip"
}

# ==========================================================================
# Cluster Network Interface Resources
# ==========================================================================

# Create a network interface in the cluster network
resource "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
  name               = "${var.prefix}-cluster-ni"

  # Use reserved IP for primary IP
  primary_ip {
    id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_subnet_reserved_ip_id
  }

  # Attach to subnet
  subnet {
    id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
  }
}

# ==========================================================================
# Virtual Network Interface Resources
# ==========================================================================

# Create a virtual network interface
resource "ibm_is_virtual_network_interface" "testacc_vni" {
  name                      = var.name
  subnet                    = ibm_is_subnet.testacc_subnet.id
  enable_infrastructure_nat = true # Enable NAT for the interface
  allow_ip_spoofing         = true # Allow IP spoofing
}

# Create another virtual network interface
resource "ibm_is_virtual_network_interface" "testacc_vni2" {
  name                      = "${var.name}-2"
  subnet                    = ibm_is_subnet.testacc_subnet.id
  enable_infrastructure_nat = true
  allow_ip_spoofing         = true
}

# Create a third virtual network interface
resource "ibm_is_virtual_network_interface" "testacc_vni3" {
  name                      = "${var.name}-3"
  subnet                    = ibm_is_subnet.testacc_subnet.id
  enable_infrastructure_nat = true
  allow_ip_spoofing         = true
}

# Associate a floating IP with a virtual network interface
resource "ibm_is_virtual_network_interface_floating_ip" "testacc_vni_floatingip" {
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
  floating_ip               = ibm_is_floating_ip.testacc_floatingip.id
}

# Associate a reserved IP with a virtual network interface
resource "ibm_is_virtual_network_interface_ip" "testacc_vni_reservedip" {
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
  reserved_ip               = ibm_is_subnet_reserved_ip.testacc_reservedip.reserved_ip
}

# ==========================================================================
# Cluster Network Profile Data Sources
# ==========================================================================

# Get information about a specific cluster network profile
data "ibm_is_cluster_network_profile" "is_cluster_network_profile_instance" {
  name = "h100" # Profile for high-performance computing with H100 GPUs
}

# List all cluster network profiles
data "ibm_is_cluster_network_profiles" "is_cluster_network_profiles_instance" {
}

# ==========================================================================
# Cluster Network Data Sources
# ==========================================================================

# Get information about a cluster network
data "ibm_is_cluster_network" "is_cluster_network_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
}

# List all cluster networks
data "ibm_is_cluster_networks" "is_cluster_networks_instance" {
}

# ==========================================================================
# Cluster Network Interface Data Sources
# ==========================================================================

# Get information about a cluster network interface
data "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
  cluster_network_id           = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_interface_id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
}

# List all interfaces in a cluster network
data "ibm_is_cluster_network_interfaces" "is_cluster_network_interfaces_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
}

# ==========================================================================
# Cluster Network Subnet Data Sources
# ==========================================================================

# Get information about a cluster network subnet
data "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
  cluster_network_id        = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
}

# List all subnets in a cluster network
data "ibm_is_cluster_network_subnets" "is_cluster_network_subnets_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
}

# ==========================================================================
# Cluster Network Subnet Reserved IP Data Sources
# ==========================================================================

# Get information about a reserved IP in a cluster network subnet
data "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
  cluster_network_id                    = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id             = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
  cluster_network_subnet_reserved_ip_id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_subnet_reserved_ip_id
}

# List all reserved IPs in a cluster network subnet
data "ibm_is_cluster_network_subnet_reserved_ips" "is_cluster_network_subnet_reserved_ips_instance" {
  cluster_network_id        = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
}

# ==========================================================================
# Virtual Network Interface Data Sources
# ==========================================================================

# Get information about a floating IP associated with a virtual network interface
data "ibm_is_virtual_network_interface_floating_ip" "is_vni_floating_ip" {
  depends_on                = [ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip]
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
  floating_ip               = ibm_is_floating_ip.testacc_floatingip.id
}

# List all floating IPs associated with a virtual network interface
data "ibm_is_virtual_network_interface_floating_ips" "is_vni_floating_ips" {
  depends_on                = [ibm_is_virtual_network_interface_floating_ip.testacc_vni_floatingip]
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
}

# List all reserved IPs associated with a virtual network interface
data "ibm_is_virtual_network_interface_ips" "is_vni_reservedips" {
  depends_on                = [ibm_is_virtual_network_interface_ip.testacc_vni_reservedip]
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
}

# Get information about a reserved IP associated with a virtual network interface
data "ibm_is_virtual_network_interface_ip" "is_vni_reservedip" {
  depends_on                = [ibm_is_virtual_network_interface_ip.testacc_vni_reservedip]
  virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
  reserved_ip               = ibm_is_subnet_reserved_ip.testacc_reservedip.reserved_ip
}

# ==========================================================================
# Resource Group Resources
# ==========================================================================

# Get the default resource group
data "ibm_resource_group" "default" {
  name = "Default" # Use the default resource group
}

# ==========================================================================
# Dedicated Host Resources
# ==========================================================================

# Create a dedicated host group (a group of dedicated hosts)
resource "ibm_is_dedicated_host_group" "dh_group01" {
  family         = "balanced"   # Host family - balanced, compute, memory
  class          = "bx2d"       # Host class
  zone           = "us-south-1" # Zone for the host group
  name           = "my-dh-group-01"
  resource_group = data.ibm_resource_group.default.id # Resource group for the host group
}

# Create a dedicated host in the host group
resource "ibm_is_dedicated_host" "is_dedicated_host" {
  profile        = "bx2d-host-152x608" # Host profile (152 vCPU, 608 GB RAM)
  name           = "my-dedicated-host-01"
  host_group     = ibm_is_dedicated_host_group.dh_group01.id # Host group for the host
  resource_group = data.ibm_resource_group.default.id        # Resource group for the host
}

# Manage disk names on a dedicated host
resource "ibm_is_dedicated_host_disk_management" "disks" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id

  # Update the name of the first disk
  disks {
    name = "newdisk01"
    id   = data.ibm_is_dedicated_host.dhost.disks.0.id
  }

  # Update the name of the second disk
  disks {
    name = "newdisk02"
    id   = data.ibm_is_dedicated_host.dhost.disks.1.id
  }
}

# ==========================================================================
# Placement Group Resources
# ==========================================================================

# Create a placement group for controlling instance placement
resource "ibm_is_placement_group" "is_placement_group" {
  strategy       = "host_spread" # Placement strategy - spread across hosts for high availability
  name           = "my-placement-group"
  resource_group = data.ibm_resource_group.default.id
}

# ==========================================================================
# Reservation Resources
# ==========================================================================

# Create a reservation for capacity
resource "ibm_is_reservation" "example" {
  capacity {
    total = 5 # Total capacity to reserve (5 instances)
  }

  committed_use {
    term = "one_year" # Commitment term - one year
  }

  profile {
    name          = "ba2-2x8"          # Profile for the reservation (2 vCPU, 8 GB RAM)
    resource_type = "instance_profile" # Type of resource to reserve
  }

  zone = "us-east-3" # Zone for the reservation
}

# Activate a reservation
resource "ibm_is_reservation_activate" "example" {
  reservation = ibm_is_reservation.example.id
}

# ==========================================================================
# Dedicated Host Data Sources
# ==========================================================================

# Get information about a dedicated host group
data "ibm_is_dedicated_host_group" "dgroup" {
  name = ibm_is_dedicated_host_group.dh_group01.name
}

# List all dedicated host groups
data "ibm_is_dedicated_host_groups" "dgroups" {
}

# Get information about a specific dedicated host profile
data "ibm_is_dedicated_host_profile" "ibm_is_dedicated_host_profile" {
  name = "bx2d-host-152x608"
}

# List all dedicated host profiles
data "ibm_is_dedicated_host_profiles" "ibm_is_dedicated_host_profiles" {
}

# List all dedicated hosts
data "ibm_is_dedicated_hosts" "dhosts" {
}

# Get information about a specific dedicated host
data "ibm_is_dedicated_host" "dhost" {
  name       = ibm_is_dedicated_host.is_dedicated_host.name
  host_group = data.ibm_is_dedicated_host_group.dgroup.id
}

# List all disks on a dedicated host
data "ibm_is_dedicated_host_disks" "dhdisks" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
}

# Get information about a specific disk on a dedicated host
data "ibm_is_dedicated_host_disk" "dhdisk" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
  disk           = ibm_is_dedicated_host_disk_management.disks.disks.0.id
}

# ==========================================================================
# Placement Group Data Sources
# ==========================================================================

# Get information about a placement group
data "ibm_is_placement_group" "is_placement_group" {
  name = ibm_is_placement_group.is_placement_group.name
}

# List all placement groups
data "ibm_is_placement_groups" "is_placement_groups" {
}

# ==========================================================================
# Reservation Data Sources
# ==========================================================================

# List all reservations
data "ibm_is_reservations" "example" {
}

# Get information about a specific reservation
data "ibm_is_reservation" "example" {
  identifier = ibm_is_reservation.example.id
}

# ==========================================================================
# Geography Data Sources
# ==========================================================================

# List all regions
data "ibm_is_regions" "regions" {
}
# Get information about a specific reservation
data "ibm_is_region" "region" {
  name = var.region
}

# List all zones
data "ibm_is_zones" "zones" {
  region = var.region
  status = ""
}
# Get information about a specific zone
data "ibm_is_zone" "zone" {
  name   = var.zone1
  region = var.region
}
