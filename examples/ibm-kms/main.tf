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

resource "ibm_kms_kmip_adapter" "myadapter" {
    instance_id = "${ibm_kms_key.test.instance_id}" 
    profile = "native_1.0"
    profile_data = {
      "crk_id" = ibm_kms_key.test.key_id
    }
    description = "adding a description"
    name = var.kmip_adapter_name
}

resource "ibm_kms_kmip_client_cert" "mycert" {
  instance_id = "${ibm_kms_key.test.instance_id}" 
  adapter_id = "${ibm_kms_kmip_adapter.myadapter.id}"
  certificate = file("${path.module}/localhost.crt")
  name = var.kmip_cert_name
}

data "ibm_kms_kmip_adapter" "adapter_data" {
  instance_id = "${ibm_kms_key.test.instance_id}" 
  name = "${ibm_kms_kmip_adapter.myadapter.name}"
  # adapter_id = "${ibm_kms_kmip_adapter.myadapter.adapter_id}"
}

data "ibm_kms_kmip_client_cert" "cert1" {
  instance_id = "${ibm_kms_key.test.instance_id}" 
  adapter_name = "${ibm_kms_kmip_adapter.myadapter.name}"
  cert_id = "${ibm_kms_kmip_client_cert.mycert.id}"
}

data "ibm_kms_kmip_adapters" "adapters" {
  instance_id = "${ibm_kms_key.test.instance_id}" 
}

data "ibm_kms_kmip_client_certs" "cert_list" {
  instance_id = "${ibm_kms_key.test.instance_id}" 
  adapter_name = "${ibm_kms_kmip_adapter.myadapter.name}"
}

data "ibm_kms_kmip_objects" "objects_list" {
  instance_id = "${ibm_kms_key.test.instance_id}" 
  adapter_id = "${ibm_kms_kmip_adapter.myadapter.id}"
  object_state_filter = [1,2,3,4]
}

data "ibm_kms_kmip_object" "object1" {
  count = length(data.ibm_kms_kmip_objects.objects_list.objects) > 0 ? 1 : 0
  instance_id = "${ibm_kms_key.test.instance_id}" 
  adapter_id = "${ibm_kms_kmip_adapter.myadapter.id}"
  object_id = "${data.ibm_kms_kmip_objects.objects_list.objects.0.object_id}"
}