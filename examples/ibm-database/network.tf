# network.tf

# Setup of network related parameters for virtual machines

#########################################################
# This file creates the security group used by the VSIs. The security group rules 
# allow internet access for install of open source packages
#########################################################

#########################################################
# Create Public Security Group 
#########################################################

resource "ibm_security_group" "sg_public_lamp" {
  name        = "sg_public_lamp"
  description = "Public access for LAMP stack to repos"
}

#########################################################
# Create policies for security group
# 1. allow tcp on 80 for HTTP access to repo's
# 2. allow tcp on 443 for HTTPS access to repo's
# Inbound http access is via private network through load balancer
#########################################################

resource "ibm_security_group_rule" "https-pub" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 443
  port_range_max    = 443
  protocol          = "tcp"
  security_group_id = "${ibm_security_group.sg_public_lamp.id}"
}

resource "ibm_security_group_rule" "http-pub" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 80
  port_range_max    = 80
  protocol          = "tcp"
  security_group_id = "${ibm_security_group.sg_public_lamp.id}"
}

# Open ICD dynamically defined port for DB access 
resource "ibm_security_group_rule" "icd-pub" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = "${ibm_database.test_acc.connectionstrings.0.hosts.0.port}"
  port_range_max    = "${ibm_database.test_acc.connectionstrings.0.hosts.0.port}"
  protocol          = "tcp"
  security_group_id = "${ibm_security_group.sg_public_lamp.id}"
}

#########################################################
# Create Private Security Group 
#########################################################

resource "ibm_security_group" "sg_private_lamp" {
  name        = "sg_private_lamp"
  description = "Private access for LAMP stack"
}

######################################################################################
# Create policies for security group
# 1. allow tcp on 80 for HTTP access to IBM Cloud hosted repo's
# 2. allow tcp on 443 for HTTPS access to repo's
# 3. allow ssh on 22 for remote administration
# 4. allow dns on 53 address lookup
# 5. allow icmp for trouble shooting
# Inbound access from internet is via private network through load balancer
# 
# source/destination ip set to generic 10.0.0.0/8, as deployment subnet not known
# at time SG is created. Similarly IPs for target data center and DAL01, WDC04, and AMS01
# not known. Future use of remote security group.
######################################################################################

resource "ibm_security_group_rule" "ssh" {
  direction         = "ingress"
  ether_type        = "IPv4"
  port_range_min    = 22
  port_range_max    = 22
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lamp.id}"
}

resource "ibm_security_group_rule" "http-in" {
  direction         = "ingress"
  ether_type        = "IPv4"
  port_range_min    = 80
  port_range_max    = 80
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lamp.id}"
}

#icmp_type specified via port_range_min 
#icmp_code specified via port_range_max 
resource "ibm_security_group_rule" "icmp" {
  direction         = "ingress"
  ether_type        = "IPv4"
  protocol          = "icmp"
  port_range_min    = 8
  port_range_max    = 0
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lamp.id}"
}

resource "ibm_security_group_rule" "http" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 80
  port_range_max    = 80
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lamp.id}"
}

resource "ibm_security_group_rule" "https" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 443
  port_range_max    = 443
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lamp.id}"
}

# Allow access to IBM DNS name servers
resource "ibm_security_group_rule" "dns" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 53
  port_range_max    = 53
  protocol          = "udp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lamp.id}"
}
