## Template to be used by the IBM Provider for Power
##
variable "servername"
{
  description="Name of the server"
}

variable "memory"{
  description="Memory Of the Power VM Instance"
}

variable "processors"{
  description="Processor Count on the server"
}

variable "proctype"{
  description="Processor Type for the LPAR - shared/dedicated"
}


variable "sshkeyname"
{
description="Key Name to be passed"
}

variable "volumename"
{
description="Volume Name to be created"
}

variable "volumesize"
{
description="Volume Size to be created"
}

variable "volumetype"
{
description="Type of volume to be created - ssd/shared"
}

variable "shareable"{
description="Should the volume be shared or not true/false"
}

resource "ibm_power_volume" "powervolumes"{


  size = "${var.volumesize}"
  name = "${var.volumename}"
  type = "${var.volumetype}"
  shareable="${var.shareable}"

}

output "volumeid"
{
  value = "${ibm_power_volume.powervolumes.id}"
}



variable "networks" {
  #default=["APP","DB","e46ed0b9-f7ef-4f9a-aed1-65a8d5623d34"]
  default=["APP","DB","e46ed0b9-f7ef-4f9a-aed1-65a8d5623d34"]
}


#variable "volumes" {
#default=["vg9"]
#}


variable "systemtype"{
  description = "Systemtype of the server"
  default="s922"
}

variable "migratable"{
  description = "Server can be migrated"
}

variable "imagename"{
  description = "Name of the image"
  default="7200-03-03"
}

#data "ibm_power_volume" "powervolume"
#{
#count = "${length(var.volumes)}"
#volumename="${var.volumes[count.index]}"
#}

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
