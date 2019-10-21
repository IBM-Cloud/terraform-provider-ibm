
data "ibm_pi_instance" "instance"
{ 
instancename="${var.pvm_instance_name}"
powerinstanceid="${var.power_instance_id}"

}


resource "ibm_pi_operations" "instance"
{
pvm_instance_name="${data.ibm_pi_instance.instance.id}"
power_instance_id="${var.power_instance_id}"
operation="${var.operation}"
}

