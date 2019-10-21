##Template to create the sshkey to be attached to a powervm 
##

resource "ibm_pi_key" "sshkey"
{
pi_key_name="${var.keyname}"
pi_ssh_key="${var.keyvalue}"
pi_cloud_instance_id="${var.powerinstanceid}"
}

