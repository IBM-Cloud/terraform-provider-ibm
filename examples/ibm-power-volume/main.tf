## Version 1.0

variable "disksize"{}
variable "diskname"{}
variable "diskshareable"{}
variable "disktype"{}




resource "ibm_power_volume" "powervolumes"{


  size = "${var.disksize}"
  name = "${var.diskname}"
  type = "${var.disktype}"
  shareable="${var.diskshareable}"

}

output "id"
{
        value = "${ibm_power_volume.powervolumes.id}"
}
