provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision sds_volume resource instance
resource "ibm_sds_volume" "sds_volume_instance_1" {
  sds_endpoint = var.sds_endpoint

  hostnqnstring = var.sds_volume_hostnqnstring
  capacity = var.sds_volume_capacity
  name = var.sds_volume_name_1
}

// Provision sds_volume resource instance
resource "ibm_sds_volume" "sds_volume_instance_2" {
  sds_endpoint = var.sds_endpoint

  hostnqnstring = var.sds_volume_hostnqnstring
  capacity = var.sds_volume_capacity
  name = var.sds_volume_name_2
}

// Provision sds_host resource instance
resource "ibm_sds_host" "sds_host_instance" {
  sds_endpoint = var.sds_endpoint

  name = var.sds_host_name
  nqn = var.sds_host_nqn
  volumes {
    volume_id = ibm_sds_volume.sds_volume_instance_1.id
    volume_name = ibm_sds_volume.sds_volume_instance_1.name
  }
  volumes {
    volume_id = ibm_sds_volume.sds_volume_instance_2.id
    volume_name = ibm_sds_volume.sds_volume_instance_2.name
  }
}
