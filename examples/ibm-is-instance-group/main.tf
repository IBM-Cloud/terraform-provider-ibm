resource "ibm_is_vpc" "vpc2" {
  name = var.vpc_name
}

resource "ibm_is_subnet" "subnet2" {
  name            = var.subnet_name
  vpc             = ibm_is_vpc.vpc2.id
  zone            = var.zone
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ssh_key" "sshkey" {
  name       = var.ssh_key_name
  public_key = var.ssh_key
}

resource "ibm_is_instance_template" "instancetemplate1" {
  name    = var.template_name
  image   = var.image_id
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = var.zone
  keys = [ibm_is_ssh_key.sshkey.id]
}

resource "ibm_is_instance_group" "instance_group" {
  name              = var.instance_group_name
  instance_template = ibm_is_instance_template.instancetemplate1.id
  instance_count    = var.instance_count
  subnets           = [ibm_is_subnet.subnet2.id]
}

resource "ibm_is_instance_group_manager" "instance_group_manager" {
  name                 = var.instance_group_manager_name
  aggregation_window   = var.aggregation_window
  instance_group       = ibm_is_instance_group.instance_group.id
  cooldown             = var.cooldown
  manager_type         = var.manager_type
  enable_manager       = var.enable_manager
  max_membership_count = var.max_membership_count
  min_membership_count = var.min_membership_count
}

resource "ibm_is_instance_group_manager" "instance_group_manager_scheduled" {
	name            = var.instance_group_manager_name_scheduled
	instance_group  = ibm_is_instance_group.instance_group.id
	manager_type    = var.manager_type.scheduled
	enable_manager  = var.enable_manager
}

resource "ibm_is_instance_group_manager_policy" "cpuPolicy" {
  instance_group         = ibm_is_instance_group.instance_group.id
  instance_group_manager = ibm_is_instance_group_manager.instance_group_manager.manager_id
  metric_type            = "cpu"
  metric_value           = var.metric_value
  policy_type            = "target"
  name                   = var.policy_name
}

  resource "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
		name                              = var.instance_group_manager_action_name
    instance_group                    = ibm_is_instance_group.instance_group.id
    instance_group_manager_scheduled  = ibm_is_instance_group_manager.instance_group_manager_scheduled.manager_id
    cron_spec                         = var.cron_spec
    instance_group_manager_autoscale  = ibm_is_instance_group_manager.instance_group_manager.manager_id
    min_membership_count              = var.max_membership_count
    max_membership_count              = var.min_membership_count
  }

data "ibm_is_instance_group" "instance_group_data" {
  name = ibm_is_instance_group.instance_group.name
}

data "ibm_is_instance_group_manager" "instance_group_manager" {
  instance_group = ibm_is_instance_group_manager.instance_group_manager.instance_group
  name           = ibm_is_instance_group_manager.instance_group_manager.name
}

data "ibm_is_instance_group_manager_policy" "instance_group_manager_policy" {
  instance_group         = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group
  instance_group_manager = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group_manager
  name                   = ibm_is_instance_group_manager_policy.cpuPolicy.name
}

data "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
		instance_group                    = ibm_is_instance_group_manager_action.instance_group_manager_action_autoscale.instance_group
		instance_group_manager_scheduled  = ibm_is_instance_group_manager_action.instance_group_manager_action_autoscale.instance_group_manager_scheduled
		name                              = ibm_is_instance_group_manager_action.instance_group_manager_action.name
}