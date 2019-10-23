## Version 1.0


resource "ibm_pi_volume_attach" "powervolumes"{


  pi_volume_attach_name = "${var.volumename}"
  pi_instance_name = "${var.instancename}"
  pi_cloud_instance_id="${var.powerinstanceid}"

}


output "id"
{
        value = "${ibm_pi_volume_attach.powervolumes.id}"
}
