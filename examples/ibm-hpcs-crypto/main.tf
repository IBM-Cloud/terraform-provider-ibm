# Provison the service using reosource_instance

/*resource "ibm_resource_instance" "hpcs"{
    name = "hpcsservice"
    service = "hs-crypto"
    plan = "standard"
    location = "us-south"
    parameters = {
        units = 2
    }
}*/

#Intialization of the hpcs crypto service may be use some scripts and null_resource to intialize the service

resource "ibm_kms_key" "key" {
  instance_id  = var.hpcs_instance_guid
  key_name     = var.key_name
  standard_key = true
  force_delete = true
}