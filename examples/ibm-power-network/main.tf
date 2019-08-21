## Version 1.0

variable "name"{}
variable "cidr"{}
variable "dns"{}
variable "type"{}




resource "ibm_power_network" "powernetwork"{


  name = "${var.name}"
  cidr = "${var.cidr}"
  dns=["${var.dns}"]
  networktype="${var.type}"

}

output "id"
{
        value = "${ibm_power_network.powernetwork.id}"
}
