data "ibm_pi_instance_ip" "networkdata"
{
pi_network_name="${var.networkname}"
pi_instance_name="${var.instancename}"
pi_cloud_instance_id="${var.powerinstanceid}"
}


output "ipoctet"
{
value="${data.ibm_pi_instance_ip.networkdata.ipoctet}"
}


output "assignedip"
{
value="${data.ibm_pi_instance_ip.networkdata.requestip}"
}
