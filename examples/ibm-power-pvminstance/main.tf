## Template to be used by the IBM Provider for Power
##

resource "ibm_power_volume" "powervolumes"{


  size = "${var.volumesize}"
  name = "${var.volumename}"
  type = "${var.volumetype}"
  shareable="${var.shareable}"

}

output "volumeid"
{
value="${ibm_power_volume.powervolumes.id}"
}


data "ibm_power_network" "powernetworks"
{
  count = "${length(var.networks)}"
  networkname = "${var.networks[count.index]}"
}


data "ibm_power_image" "powerimages"{
  name="${var.imagename}"
}


resource "ibm_power_pvminstance" "pvminstance"
{
	memory = "${var.memory}"
	processors = "${var.processors}"
	servername = "${var.servername}"
    proctype="${var.proctype}"
    migratable="${var.migratable}"
    imageid = "${data.ibm_power_image.powerimages.imageid}"
    volumeids = ["${ibm_power_volume.powervolumes.id}"]
	networkids = ["${data.ibm_power_network.powernetworks.*.networkid}"]
	keypairname="${var.sshkeyname}"
	systype="${var.systemtype}"
	replicationpolicy="${var.replicationpolicy}"
	replicants="${var.replicants}"
}

output "status"
{
value="${ibm_power_pvminstance.pvminstance.status}"
}


output "minproc"
{
value="${ibm_power_pvminstance.pvminstance.minproc}"
}

output "healthstatus"
{
value="${ibm_power_pvminstance.pvminstance.healthstatus}"
}

output "ipaddress"
{
value="${ibm_power_pvminstance.pvminstance.addresses}"
}

output "progress"
{
value="${ibm_power_pvminstance.pvminstance.progress}"
}
