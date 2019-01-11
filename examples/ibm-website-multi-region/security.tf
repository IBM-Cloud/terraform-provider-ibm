# security.tf

# Setup of network security related parameters for virtual machines

#########################################################
# This file creates the security group used by the VSIs. The security group rules 
# allow internet access for install of open source packages
#########################################################

#########################################################
# Create Public Security Group for all VSIs
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

####################################################################################
# Create Private Security Group for web and app servers
####################################################################################

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
# 6. allow mariadb on 3306 outbound to remote security group sg_private_lamp
# Inbound access from internet is via private network through load balancer
# 
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

resource "ibm_security_group_rule" "mysql-out" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 3306
  port_range_max    = 3306
  protocol          = "tcp"
  remote_group_id   = "${ibm_security_group.sg_private_lampdb.id}"
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

# Access to repos
resource "ibm_security_group_rule" "http" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 80
  port_range_max    = 80
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lamp.id}"
}

# Access to repos
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

# for testing
# resource "ibm_security_group_rule" "icmp-out" {
#   direction         = "egress"
#   ether_type        = "IPv4"
#   protocol          = "icmp"
#   port_range_min    = 8
#   port_range_max    = 0
#   remote_group_id   = "${ibm_security_group.sg_private_lampdb.id}"
#   security_group_id = "${ibm_security_group.sg_private_lamp.id}"
# }

# for testing
# resource "ibm_security_group_rule" "ssh-out" {
#   direction         = "egress"
#   ether_type        = "IPv4"
#   port_range_min    = 22
#   port_range_max    = 22
#   protocol          = "tcp"
#   remote_ip         = "10.0.0.0/8"
#   security_group_id = "${ibm_security_group.sg_private_lamp.id}"
# }

#########################################################
# Create Private Security Group for mariadb servers
#########################################################

resource "ibm_security_group" "sg_private_lampdb" {
  name        = "sg_private_lampdb"
  description = "Private access for lampdb stack"
}

######################################################################################
# Create policies for security group
# 1. allow tcp on 80 for HTTP access to IBM Cloud hosted repo's
# 2. allow tcp on 443 for HTTPS access to repo's
# 3. allow ssh on 22 for remote administration
# 4. allow dns on 53 address lookup
# 5. allow icmp for trouble shooting
# 6. allow mariadb on 3306 outbound to remote security group sg_private_lampdb for replication
# 6. allow mariadb on 3306 inbound to remote security group sg_private_lampdb for replication
# 7. allow mariadb on 3306 inbound from remote security group sg_private_lamp
# Inbound access from internet is via private network through load balancer
# 
######################################################################################

resource "ibm_security_group_rule" "ssh-db" {
  direction         = "ingress"
  ether_type        = "IPv4"
  port_range_min    = 22
  port_range_max    = 22
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

resource "ibm_security_group_rule" "http-db-in" {
  direction         = "ingress"
  ether_type        = "IPv4"
  port_range_min    = 80
  port_range_max    = 80
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

# In-bound for DB access from web/app servers
resource "ibm_security_group_rule" "mysql-db-in" {
  direction         = "ingress"
  ether_type        = "IPv4"
  port_range_min    = 3306
  port_range_max    = 3306
  protocol          = "tcp"
  remote_group_id   = "${ibm_security_group.sg_private_lamp.id}"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

# In-bound for DB replication between Marisdb VSIs
resource "ibm_security_group_rule" "mysql-db-repl-in" {
  direction         = "ingress"
  ether_type        = "IPv4"
  port_range_min    = 3306
  port_range_max    = 3306
  protocol          = "tcp"
  remote_group_id   = "${ibm_security_group.sg_private_lampdb.id}"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

# Out-bound for DB replication between Marisdb VSIs
resource "ibm_security_group_rule" "mysql-db-repl-out" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 3306
  port_range_max    = 3306
  protocol          = "tcp"
  remote_group_id   = "${ibm_security_group.sg_private_lampdb.id}"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

#icmp_type specified via port_range_min 
#icmp_code specified via port_range_max 
resource "ibm_security_group_rule" "icmp-db" {
  direction         = "ingress"
  ether_type        = "IPv4"
  protocol          = "icmp"
  port_range_min    = 8
  port_range_max    = 0
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

resource "ibm_security_group_rule" "http-db" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 80
  port_range_max    = 80
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

resource "ibm_security_group_rule" "https-db" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 443
  port_range_max    = 443
  protocol          = "tcp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

# Allow access to IBM DNS name servers
resource "ibm_security_group_rule" "dns-db" {
  direction         = "egress"
  ether_type        = "IPv4"
  port_range_min    = 53
  port_range_max    = 53
  protocol          = "udp"
  remote_ip         = "10.0.0.0/8"
  security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
}

# for testing
# resource "ibm_security_group_rule" "icmp-db-out" {
#   direction         = "egress"
#   ether_type        = "IPv4"
#   protocol          = "icmp"
#   port_range_min    = 8
#   port_range_max    = 0
#   remote_ip         = "10.0.0.0/8"
#   security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
# }

# for testing
# resource "ibm_security_group_rule" "ssh-db-out" {
#   direction         = "egress"
#   ether_type        = "IPv4"
#   port_range_min    = 22
#   port_range_max    = 22
#   protocol          = "tcp"
#   remote_ip         = "10.0.0.0/8"
#   security_group_id = "${ibm_security_group.sg_private_lampdb.id}"
# }
