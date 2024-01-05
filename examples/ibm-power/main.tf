# Create a workspace
resource "ibm_resource_instance" "location" {
  name              = var.workspace_name
  resource_group_id = var.resource_group_id
  location          = var.datacenter
  service           = "power-iaas"
  plan              = "power-virtual-server-group"
}

# Create an image
resource "ibm_pi_image" "image" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_image_name        = var.image_name
  pi_image_id          = var.image_id
}
data "ibm_pi_image" "data_source_image" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_image_name        = resource.ibm_pi_image.image.pi_image_name
}

# Create a network
resource "ibm_pi_network" "private_network" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_network_name      = var.network_name
  pi_network_type      = var.network_type
  pi_cidr              = var.network_cidr
  pi_dns               = [var.network_dns]
  pi_network_mtu       = 2000
}
data "ibm_pi_network" "data_source_private_network" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_network_name      = resource.ibm_pi_network.private_network.pi_network_name
}

# Create a volume
resource "ibm_pi_volume" "volume" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_volume_name       = var.volume_name
  pi_volume_type       = var.volume_type
  pi_volume_size       = var.volume_size
  pi_volume_shareable  = var.volume_shareable
}
data "ibm_pi_volume" "data_source_volume" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_volume_name       = resource.ibm_pi_volume.volume.pi_volume_name
}

# Create an ssh key
resource "ibm_pi_key" "key" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_key_name          = var.ssh_key_name
  pi_ssh_key           = var.ssh_key_rsa
}
data "ibm_pi_key" "data_source_key" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_key_name          = resource.ibm_pi_key.key.pi_key_name
}

# Create an instance
resource "ibm_pi_instance" "instance" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_instance_name     = var.instance_name
  pi_memory            = var.memory
  pi_processors        = var.processors
  pi_proc_type         = var.proc_type
  pi_storage_type      = var.storage_type
  pi_sys_type          = var.sys_type
  pi_image_id          = data.ibm_pi_image.data_source_image.id
  pi_key_pair_name     = data.ibm_pi_key.data_source_key.id
  pi_network {
    network_id = data.ibm_pi_network.data_source_private_network.id
  }
  pi_volume_ids = [data.ibm_pi_volume.data_source_volume.id]
}
data "ibm_pi_instance" "data_source_instance" {
  pi_cloud_instance_id = ibm_resource_instance.location.guid
  pi_instance_name     = resource.ibm_pi_instance.instance.pi_instance_name
}
