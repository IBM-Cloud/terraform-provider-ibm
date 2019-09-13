## Version 1.0


resource "ibm_pi_volume" "powervolumes"{


  pi_volume_size = "${var.volumesize}"
  pi_volume_name = "${var.volumename}"
  pi_volume_type = "${var.volumetype}"
  pi_volume_shareable="${var.volumeshareable}"
  pi_cloud_instance_id="${var.powerinstanceid}"

}


output "id"
{
        value = "${ibm_pi_volume.powervolumes.id}"
}
