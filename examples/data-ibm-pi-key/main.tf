data "ibm_pi_key" "sshkey"
{
pi_key_name="${var.sshkeyname}"
pi_cloud_instance_id="${var.powerinstanceid}"

}
