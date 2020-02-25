resource "ibm_resource_instance" "cos_instance" {
  name     = var.cos_name
  service  = "cloud-object-storage"
  plan     = var.pln
  location = var.location
}
resource "ibm_resource_instance" "kp_instance" {
  name     = var.kp_name
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
resource "ibm_kp_key" "test" {
  key_protect_id = ibm_resource_instance.kp_instance.guid
  key_name       = var.key_name
  standard_key   = false
}

resource "ibm_iam_authorization_policy" "policy" {
  source_service_name = "cloud-object-storage"
  target_service_name = "kms"
  roles               = ["Reader"]
}

resource "ibm_cos_bucket" "flex-us-south" {
  depends_on           = [ibm_iam_authorization_policy.policy]
  bucket_name          = "abuck4"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "flex"
  key_protect          = ibm_kp_key.test.id
}