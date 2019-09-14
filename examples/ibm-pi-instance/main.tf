## Template to be used by the IBM Provider for Power
##
## Added the cloud_init_data to be passed in 

resource "ibm_pi_volume" "powervolumes"{


  pi_volume_size = "${var.volumesize}"
  pi_volume_name = "${var.volumename}"
  pi_volume_type = "${var.volumetype}"
  pi_volume_shareable="${var.shareable}"
  pi_cloud_instance_id="${var.powerinstanceid}"

}

output "volumeid"
{
value="${ibm_pi_volume.powervolumes.id}"
}


data "ibm_pi_network" "powernetworks"
{
  count = "${length(var.networks)}"
  pi_network_name = "${var.networks[count.index]}"
  pi_cloud_instance_id="${var.powerinstanceid}"
}


data "ibm_pi_image" "powerimages"{
  pi_image_name="${var.imagename}"
  pi_cloud_instance_id="${var.powerinstanceid}"
}

data "ibm_pi_key" "powerkey"{
pi_cloud_instance_id="${var.powerinstanceid}"
pi_key_name="${var.sshkeyname}"
}


resource "ibm_pi_instance" "pvminstance"
{
	pi_memory = "${var.memory}"
	pi_processors = "${var.processors}"
	pi_instance_name = "${var.servername}"
    pi_proc_type="${var.proctype}"
    pi_migratable="${var.migratable}"
    pi_image_id = "${data.ibm_pi_image.powerimages.imageid}"
   pi_volume_ids = ["${ibm_pi_volume.powervolumes.id}"]
	pi_network_ids = ["${data.ibm_pi_network.powernetworks.*.networkid}"]
	pi_key_pair_name="${var.sshkeyname}"
	pi_sys_type="${var.systemtype}"
	pi_replication_policy="${var.replicationpolicy}"
	pi_replicants="${var.replicants}"
        pi_cloud_instance_id="${var.powerinstanceid}"
	pi_user_data="${var.cloud_init_data}"
}

output "status"
{
value="${ibm_pi_instance.pvminstance.pi_instance_status}"
}


output "minproc"
{
value="${ibm_pi_instance.pvminstance.pi_minproc}"
}

output "healthstatus"
{
value="${ibm_pi_instance.pvminstance.pi_health_status}"
}

output "ipaddress"
{
value="${ibm_pi_instance.pvminstance.addresses}"
}

output "progress"
{
value="${ibm_pi_instance.pvminstance.pi_progress}"
}
