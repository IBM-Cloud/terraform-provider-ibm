resource "ibm_pi_key" "key" {
  pi_cloud_instance_id = var.powerinstanceid
  pi_key_name          = var.sshkeyname
  pi_ssh_key           = var.sshkey
}

data "ibm_pi_key" "dskey" {
  depends_on           = [ibm_pi_key.key]
  pi_cloud_instance_id = var.powerinstanceid
  pi_key_name          = var.sshkeyname
}

resource "ibm_pi_network" "power_networks" {
  count                = 1
  pi_network_name      = var.networkname
  pi_cloud_instance_id = var.powerinstanceid
  pi_network_type      = "pub-vlan"
}

data "ibm_pi_public_network" "dsnetwork" {
  depends_on           = [ibm_pi_network.power_networks]
  pi_cloud_instance_id = var.powerinstanceid
}

resource "ibm_pi_volume" "volume" {
  pi_volume_size       = 20
  pi_volume_name       = var.volname
  pi_volume_type       = "ssd"
  pi_volume_shareable  = true
  pi_cloud_instance_id = var.powerinstanceid // Get it by running cmd "ibmcloud resource service-instances --long"
}

data "ibm_pi_volume" "dsvolume" {
  depends_on           = [ibm_pi_volume.volume]
  pi_cloud_instance_id = var.powerinstanceid
  pi_volume_name       = var.volname
}

data "ibm_pi_image" "powerimages" {
  pi_image_name        = var.imagename
  pi_cloud_instance_id = var.powerinstanceid
}

resource "ibm_pi_instance" "test-instance" {
  pi_memory            = "4"
  pi_processors        = "2"
  pi_instance_name     = var.instancename
  pi_proc_type         = "shared"
  pi_image_id          = data.ibm_pi_image.powerimages.id
  pi_key_pair_name     = data.ibm_pi_key.dskey.id
  pi_sys_type          = "s922"
  pi_cloud_instance_id = var.powerinstanceid
  pi_volume_ids        = [data.ibm_pi_volume.dsvolume.id]

  pi_network {
    network_id = data.ibm_pi_public_network.dsnetwork.id
  }
}
