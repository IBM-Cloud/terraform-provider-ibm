# --------------------------------------------------
# Provison the service using reosource_instance
# --------------------------------------------------
// Supported Regions are us-south and us-east
resource ibm_hpcs hpcs {
  location             = var.location
  name                 = var.hpcs_instance_name
  plan                 = var.plan
  units                = var.units
  signature_threshold  = var.signature_threshold
  revocation_threshold = var.revocation_threshold
  dynamic admins {
    for_each = var.admins
    content {
      name  = admins.value.name
      key   = admins.value.key
      token = admins.value.token
    }
  }
}
# --------------------------------
# Creating Keys for HPCS Instance
# --------------------------------

resource "ibm_kms_key" "key" {
  instance_id  = ibm_hpcs.hpcs.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
}
