provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision sds_volume resource instance
resource "ibm_sds_volume" "sds_volume_instance_1" {
  sds_endpoint = var.sds_endpoint

  capacity = var.sds_volume_capacity
  name = var.sds_volume_name_1
}

// Provision sds_volume resource instance
resource "ibm_sds_volume" "sds_volume_instance_2" {
  sds_endpoint = var.sds_endpoint

  capacity = var.sds_volume_capacity
  name = var.sds_volume_name_2
}

// Provision sds_host resource instance
resource "ibm_sds_host" "sds_host_instance" {
  sds_endpoint = var.sds_endpoint

  name = var.sds_host_name
  nqn = var.sds_host_nqn
}

// Provision sds_volume_mapping resource instance
resource "ibm_sds_volume_mapping" "sds_vm_1" {

  depends_on = [time_sleep.wait_5_seconds]

  host_id = ibm_sds_host.sds_host_instance.id
  volume {
    id = ibm_sds_volume.sds_volume_instance_1.id
  }
}

// Provision sds_volume_mapping resource instance
resource "ibm_sds_volume_mapping" "sds_vm_2" {

  depends_on = [time_sleep.wait_5_seconds]

  host_id = ibm_sds_host.sds_host_instance.id
  volume {
    id = ibm_sds_volume.sds_volume_instance_2.id
  }
}

// Use this sleep to allow the volume mappings to delete first before deleting the volumes and hosts
resource "time_sleep" "wait_5_seconds" {
  depends_on = [
    ibm_sds_volume.sds_volume_instance_1,
    ibm_sds_volume.sds_volume_instance_2,
    ibm_sds_host.sds_host_instance
  ]
  destroy_duration = "5s"
}
