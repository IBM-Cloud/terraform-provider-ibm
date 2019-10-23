## Version 1.0


resource "ibm_pi_network" "powernetwork"{


  pi_network_type = "${var.network_type}"
  pi_network_name = "${var.network_name}"
  pi_dns = ["${var.network_dns}"]
  pi_cidr="${var.network_CIDR}"
  pi_cloud_instance_id="${var.cloudinstanceid}"

}


output "id"
{
        value = "${ibm_pi_network.powernetwork.id}"
}
