data "ibm_pi_network" "networkdata"
{
pi_network_name="${var.networkname}"
pi_cloud_instance_id="${var.powerinstanceid}"
}



output "gateway_ip"
{
value="${data.ibm_pi_network.networkdata.gateway}"
}

output "vlan_type"
{
value="${data.ibm_pi_network.networkdata.type}"
}

output "available_ip_count"
{
value="${data.ibm_pi_network.networkdata.available_ip_count}"
}

output "used_ip_count"
{
value="${data.ibm_pi_network.networkdata.used_ip_count}"
}

output "used_ip_percent"
{
value="${data.ibm_pi_network.networkdata.used_ip_percent}"
}
