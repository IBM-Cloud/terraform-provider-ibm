resource "ibm_resource_instance" "cos_instance" {
  name     = var.cos_name
  service  = "cloud-object-storage"
  plan     = var.plan
  location = var.location
}
resource "ibm_resource_instance" "kms_instance" {
  name     = var.kms_name
  service  = "kms"
  plan     =  var.kms_plan
  location = var.kms_location
}
resource "ibm_kms_key" "test" {
  instance_id = ibm_resource_instance.kms_instance.guid
  key_name       = var.key_name
  endpoint_type = "public"
  standard_key   = false
  force_delete = true
}

resource "ibm_iam_authorization_policy" "policy" {
  source_service_name = "cloud-object-storage"
  target_service_name = "kms"
  roles               = ["Reader"]
}

data "ibm_kms_keys" "test" {
		instance_id = "${ibm_kms_key.test.instance_id}" 
}

resource "ibm_cos_bucket" "flex-us-south" {
  depends_on           = [ibm_iam_authorization_policy.policy]
  bucket_name          = var.bucket_name
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "flex"
  kms_key_crn          = ibm_kms_key.test.id
}