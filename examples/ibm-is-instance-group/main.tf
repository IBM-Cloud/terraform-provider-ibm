# ==========================================================================
# VPC Instance Group Configuration
# ==========================================================================

# ==========================================================================
# Base Infrastructure Resources
# ==========================================================================

# Create a VPC for the instance group
resource "ibm_is_vpc" "vpc2" {
  name = var.vpc_name # Name for the VPC
}

# Create a subnet within the VPC
resource "ibm_is_subnet" "subnet2" {
  name            = var.subnet_name    # Name for the subnet
  vpc             = ibm_is_vpc.vpc2.id # VPC ID
  zone            = var.zone           # Zone for the subnet
  ipv4_cidr_block = "10.240.64.0/28"   # IPv4 CIDR block for the subnet (16 addresses)
}

# Create an SSH key for instance access
resource "ibm_is_ssh_key" "sshkey" {
  name       = var.ssh_key_name # Name for the SSH key
  public_key = var.ssh_key      # Public key material
}

# ==========================================================================
# Instance Template Resources
# ==========================================================================

# Create an instance template for the instance group
resource "ibm_is_instance_template" "instancetemplate1" {
  name    = var.template_name # Name for the template
  image   = var.image_id      # OS image ID
  profile = var.profile       # Instance profile

  # Define the primary network interface
  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id # Subnet for instances
  }

  vpc  = ibm_is_vpc.vpc2.id         # VPC for instances
  zone = var.zone                   # Zone for instances
  keys = [ibm_is_ssh_key.sshkey.id] # SSH key for access
}

# ==========================================================================
# Instance Group Resources
# ==========================================================================

# Create an instance group using the template
resource "ibm_is_instance_group" "instance_group" {
  name              = var.instance_group_name                       # Name for the instance group
  instance_template = ibm_is_instance_template.instancetemplate1.id # Template to use
  instance_count    = var.instance_count                            # Initial number of instances
  subnets           = [ibm_is_subnet.subnet2.id]                    # Subnets for instances
}

# Get the list of memberships in the instance group
data "ibm_is_instance_group_memberships" "is_instance_group_memberships" {
  instance_group = ibm_is_instance_group.instance_group.id
}

# Manage a specific membership in the instance group
resource "ibm_is_instance_group_membership" "is_instance_group_membership" {
  instance_group = ibm_is_instance_group.instance_group.id
  # Reference the first membership from the data source
  instance_group_membership = data.ibm_is_instance_group_memberships.is_instance_group_memberships.memberships.0.instance_group_membership
  name                      = var.instance_group_membership # Name for the membership
}

# ==========================================================================
# Instance Group Manager Resources - Autoscale
# ==========================================================================

# Create an autoscale manager for the instance group
resource "ibm_is_instance_group_manager" "instance_group_manager" {
  name                 = var.instance_group_manager_name         # Name for the manager
  aggregation_window   = var.aggregation_window                  # Time window in seconds to aggregate metrics
  instance_group       = ibm_is_instance_group.instance_group.id # Instance group to manage
  cooldown             = var.cooldown                            # Cooldown period in seconds
  manager_type         = var.manager_type                        # Type of manager (autoscale)
  enable_manager       = var.enable_manager                      # Whether manager is enabled
  max_membership_count = var.max_membership_count                # Maximum instances
  min_membership_count = var.min_membership_count                # Minimum instances
}

# ==========================================================================
# Instance Group Manager Resources - Scheduled
# ==========================================================================

# Create a scheduled manager for the instance group
resource "ibm_is_instance_group_manager" "instance_group_manager_scheduled" {
  name           = var.instance_group_manager_name_scheduled # Name for the manager
  instance_group = ibm_is_instance_group.instance_group.id   # Instance group to manage
  manager_type   = var.manager_type_scheduled                # Type of manager (scheduled)
  enable_manager = var.enable_manager                        # Whether manager is enabled
}

# ==========================================================================
# Instance Group Manager Policy Resources
# ==========================================================================

# Create a CPU-based scaling policy
resource "ibm_is_instance_group_manager_policy" "cpuPolicy" {
  instance_group         = ibm_is_instance_group.instance_group.id                         # Instance group
  instance_group_manager = ibm_is_instance_group_manager.instance_group_manager.manager_id # Manager
  metric_type            = "cpu"                                                           # Metric to monitor
  metric_value           = var.metric_value                                                # Target metric value
  policy_type            = "target"                                                        # Policy type
  name                   = var.policy_name                                                 # Policy name
}

# ==========================================================================
# Instance Group Manager Action Resources
# ==========================================================================

# Create a scheduled action for the instance group
resource "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
  name                   = var.instance_group_manager_action_name                                    # Action name
  instance_group         = ibm_is_instance_group.instance_group.id                                   # Instance group
  instance_group_manager = ibm_is_instance_group_manager.instance_group_manager_scheduled.manager_id # Manager
  cron_spec              = var.cron_spec                                                             # Cron schedule
  target_manager         = ibm_is_instance_group_manager.instance_group_manager.manager_id           # Target manager
  min_membership_count   = var.min_membership_count
  max_membership_count   = var.max_membership_count
}

# ==========================================================================
# Instance Group Data Sources
# ==========================================================================

# Get information about the instance group
data "ibm_is_instance_group" "instance_group_data" {
  name = ibm_is_instance_group.instance_group.name
}

# Get information about a specific membership
data "ibm_is_instance_group_membership" "is_instance_group_membership" {
  instance_group = ibm_is_instance_group.instance_group.id
  name           = ibm_is_instance_group_membership.is_instance_group_membership.name
}

# ==========================================================================
# Instance Group Manager Data Sources
# ==========================================================================

# Get information about the instance group manager
data "ibm_is_instance_group_manager" "instance_group_manager" {
  instance_group = ibm_is_instance_group_manager.instance_group_manager.instance_group
  name           = ibm_is_instance_group_manager.instance_group_manager.name
}

# Get information about the CPU policy
data "ibm_is_instance_group_manager_policy" "instance_group_manager_policy" {
  instance_group         = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group
  instance_group_manager = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group_manager
  name                   = ibm_is_instance_group_manager_policy.cpuPolicy.name
}

# Get information about the scheduled action
data "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
  instance_group         = ibm_is_instance_group_manager_action.instance_group_manager_action.instance_group
  instance_group_manager = ibm_is_instance_group_manager_action.instance_group_manager_action.instance_group_manager
  name                   = ibm_is_instance_group_manager_action.instance_group_manager_action.name
}