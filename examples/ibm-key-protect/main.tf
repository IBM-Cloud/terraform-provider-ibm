resource "ibm_resource_instance" "cos_instance" {
  name     = var.cos_name
  service  = "cloud-object-storage"
  plan     = var.plan
  location = var.location
}
resource "ibm_resource_instance" "kp_instance" {
  name     = var.kp_name
  service  = "kms"
  plan     =  var.kp_plan
  location = var.kp_location
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

data "ibm_kp_key" "test" {
		key_protect_id = "${ibm_kp_key.test.key_protect_id}" 
}

resource "ibm_cos_bucket" "smart-us-south" {
  depends_on           = [ibm_iam_authorization_policy.policy]
  bucket_name          = var.bucket_name
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "smart"
  kms_key_crn          = ibm_kp_key.test.id
}