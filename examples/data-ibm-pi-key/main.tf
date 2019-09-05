data "ibm_pi_key" "sshkey"
{
name="${var.sshkeyname}"
powerinstanceid="${var.powerinstanceid}"

}
