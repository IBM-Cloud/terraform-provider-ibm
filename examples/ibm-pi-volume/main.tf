## Version 1.0


resource "ibm_pi_volume" "powervolumes"{

  size = "${var.volumesize}"
  name = "${var.volumename}"
  type = "${var.volumetype}"
  shareable="${var.volumeshareable}"
  powerinstanceid="${var.powerinstanceid}"

}

output "id"
{
        value = "${ibm_pi_volume.powervolumes.id}"
}
