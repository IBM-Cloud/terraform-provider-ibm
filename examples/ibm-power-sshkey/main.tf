##Template to create the sshkey to be attached to a powervm 
##

resource "ibm_power_sshkey" "sshkey"
{
name="${var.sshkeyname}"
sshkey="${var.keypair}"
}

