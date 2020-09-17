# satellite MAIN.tf
# This file runs each of the modules

resource "ibm_compute_ssh_key" "test_ssh_key" {
  label      = var.ssh_label
  notes      = var.notes
  public_key = var.ssh_public_key
}

module "create_location" {
  source            = "./modules/create_location"
  module_depends_on = "aa"
  zone              = var.zone
  location          = var.location
  # cos_key           = var.cos_key
  # cos_key_id        = var.cos_key_id
  label             = var.label
}

module "vms" {
  source            = "./modules/vms"
  module_depends_on = module.create_location.trigger
  vmcount           = var.vmcount
  vmname            = "${var.vmname}"
  domain            = var.domain
  ssh_key_id        = ibm_compute_ssh_key.test_ssh_key.id
  os                = var.os
  datacenter        = var.datacenter
  network_speed     = var.network_speed
  cores             = var.cores
  memory            = var.memory
  disks             = var.disks
}

module "run_ssh" {
  source = "./modules/run_ssh"
  module_depends_on    = module.vms.vm_ip.0
  ipcount         = var.vmcount
  hostip          = module.vms.vm_ip
  private_ssh_key = var.private_ssh_key
  path            = "./scripts/addhost.sh"
}

#####################################################
# assign host
#####################################################
module "assign" {
  source            = "./modules/assign_host"
  module_depends_on = module.run_ssh.trigger
  ip_count          = 3
  host_vm           = module.vms.host_vm_names
  location          = var.location
}

module "register_dns" {
  source      = "./modules/register_dns"
  module_depends_on    = module.assign.trigger
  host_ip     = module.vms.vm_ip
  location = var.location
}

module "create_cluster" {
  source      = "./modules/create_cluster"
  module_depends_on    = module.register_dns.trigger
  cluster_name     = var.cluster_name
  location = var.location
}
module "assign_cluster" {
  source      = "./modules/assign_cluster"
  module_depends_on    = module.create_cluster.trigger
  cluster_name     = var.cluster_name
  location = var.location
  host_vm=module.vms.host_vm_names.3
  zone=var.host_zone
}