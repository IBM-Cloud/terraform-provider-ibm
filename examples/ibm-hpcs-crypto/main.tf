# --------------------------------------------------
# Provison the service using reosource_instance
# --------------------------------------------------

resource "ibm_resource_instance" "hpcs_instance" {
  count    = (var.provision_instance == true ? 1 : 0)
  name     = var.hpcs_instance_name
  service  = "hs-crypto"
  plan     = var.plan
  location = var.location
  parameters = {
    units = var.units
  }
}
# data source of the hpcs instance
data "ibm_resource_instance" "hpcs_instance" {
  name     = (var.provision_instance == true ? ibm_resource_instance.hpcs_instance.0.name : var.hpcs_instance_name)
  service  = "hs-crypto"
  location = var.location
}

# ------------------------------------------------------------------------------------------------------------
# #Intialization of the hpcs crypto service may be use some scripts and null_resource to intialize the service
# -------------------------------------------------------------------------------------------------------------

# --------------------------------
# Cresting Keys for HPCS Instance
# --------------------------------

resource "ibm_kms_key" "key" {
  depends_on   = [null_resource.hpcs_init]
  instance_id  = data.ibm_resource_instance.hpcs_instance.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
}
