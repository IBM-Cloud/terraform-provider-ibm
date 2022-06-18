data "ibm_pi_image" "data_source_image" {
  pi_cloud_instance_id = var.cloud_instance_id
  pi_image_name        = var.image_name
}
resource "ibm_pi_key" "key" {
  pi_cloud_instance_id = var.cloud_instance_id
  pi_key_name          = var.ssh_key_name
  pi_ssh_key           = var.ssh_key_rsa
}
data "ibm_pi_key" "data_source_key" {
  depends_on = [ibm_pi_key.key]

  pi_cloud_instance_id = var.cloud_instance_id
  pi_key_name          = var.ssh_key_name
}
resource "ibm_pi_network" "network" {
  pi_cloud_instance_id = var.cloud_instance_id
  pi_network_name      = var.network_name
  pi_network_type      = var.network_type
  count                = var.network_count
}
data "ibm_pi_public_network" "data_source_network" {
  depends_on = [ibm_pi_network.network]

  pi_cloud_instance_id = var.cloud_instance_id
}
resource "ibm_pi_volume" "volume" {
  pi_cloud_instance_id = var.cloud_instance_id
  pi_volume_name       = var.volume_name
  pi_volume_type       = var.volume_type
  pi_volume_size       = var.volume_size
  pi_volume_shareable  = var.volume_shareable
}
data "ibm_pi_volume" "data_source_volume" {
  depends_on = [ibm_pi_volume.volume]

  pi_cloud_instance_id = var.cloud_instance_id
  pi_volume_name       = var.volume_name
}
resource "ibm_pi_instance" "instance" {
  depends_on = [data.ibm_pi_image.data_source_image,
    data.ibm_pi_key.data_source_key,
    data.ibm_pi_volume.data_source_volume,
  data.ibm_pi_public_network.data_source_network]

  pi_cloud_instance_id = var.cloud_instance_id
  pi_instance_name     = var.instance_name
  pi_memory            = var.memory
  pi_processors        = var.processors
  pi_proc_type         = var.proc_type
  pi_storage_type      = var.storage_type
  pi_sys_type          = var.sys_type
  pi_image_id          = data.ibm_pi_image.data_source_image.id
  pi_key_pair_name     = data.ibm_pi_key.data_source_key.id
  pi_network { network_id = data.ibm_pi_public_network.data_source_network.id }
  pi_volume_ids = [data.ibm_pi_volume.data_source_volume.id]
}

data "ibm_pi_instance" "data_source_instance" {
  depends_on = [ibm_pi_instance.instance]

  pi_cloud_instance_id = var.cloud_instance_id
  pi_instance_name     = var.instance_name
}