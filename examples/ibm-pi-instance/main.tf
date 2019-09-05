## Template to be used by the IBM Provider for Power
##

resource "ibm_pi_volume" "powervolumes"{


  size = "${var.volumesize}"
  name = "${var.volumename}"
  type = "${var.volumetype}"
  shareable="${var.shareable}"
  powerinstanceid="${var.powerinstanceid}"

}

output "volumeid"
{
value="${ibm_pi_volume.powervolumes.id}"
}


data "ibm_pi_network" "powernetworks"
{
  count = "${length(var.networks)}"
  networkname = "${var.networks[count.index]}"
  powerinstanceid="${var.powerinstanceid}"
}


data "ibm_pi_image" "powerimages"{
  name="${var.imagename}"
  powerinstanceid="${var.powerinstanceid}"
}


resource "ibm_pi_instance" "pvminstance"
{
	memory = "${var.memory}"
	processors = "${var.processors}"
	servername = "${var.servername}"
    proctype="${var.proctype}"
    migratable="${var.migratable}"
    imageid = "${data.ibm_pi_image.powerimages.imageid}"
    volumeids = ["${ibm_pi_volume.powervolumes.id}"]
	networkids = ["${data.ibm_pi_network.powernetworks.*.networkid}"]
	keypairname="${var.sshkeyname}"
	systype="${var.systemtype}"
	replicationpolicy="${var.replicationpolicy}"
	replicants="${var.replicants}"
        powerinstanceid="${var.powerinstanceid}"
}

output "status"
{
value="${ibm_pi_instance.pvminstance.status}"
}


output "minproc"
{
value="${ibm_pi_instance.pvminstance.minproc}"
}

output "healthstatus"
{
value="${ibm_pi_instance.pvminstance.healthstatus}"
}

output "ipaddress"
{
value="${ibm_pi_instance.pvminstance.addresses}"
}

output "progress"
{
value="${ibm_pi_instance.pvminstance.progress}"
}
