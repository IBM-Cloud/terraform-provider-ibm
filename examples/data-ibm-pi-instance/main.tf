data "ibm_pi_instance" "instance"
{
pi_instance_name="${var.instancename}"
pi_cloud_instance_id="${var.powerinstanceid}"

}



output "state"
{
value="${data.ibm_pi_instance.instance.status}"
}

output "address"
{
value="${data.ibm_pi_instance.instance.addresses}"
}

output "id"
{
value="${data.ibm_pi_instance.instance.id}"
}

