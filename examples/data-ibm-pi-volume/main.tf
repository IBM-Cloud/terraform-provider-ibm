data "ibm_pi_volume" "volumedata"
{
pi_volume_name="${var.volumename}"
pi_cloud_instance_id="${var.powerinstanceid}"
}



output "id"
{
value="${data.ibm_pi_volume.volumedata.volumeid}"
}

