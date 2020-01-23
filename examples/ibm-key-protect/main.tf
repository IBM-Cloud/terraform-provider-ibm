resource "ibm_resource_instance" "cos_instance" {
  name              = "${var.cos_name}"
  service           = "cloud-object-storage"
  plan              = "${var.pln}"
  location          = "${var.location}"
}
resource "ibm_resource_instance" "kp_instance" {
  name              = "${var.kp_name}"
  service           = "kms"
  plan              = "tiered-pricing"
  location          = "us-south"
}
resource "ibm_kp_key" "test" {
  key_protect_id = "${ibm_resource_instance.kp_instance.guid}"
  key_name = "${var.key_name}"
  standard_key =  false
}
data "ibm_kp_key" "test" {
  depends_on  = ["ibm_kp_key.test"]
  key_protect_id = "${ibm_resource_instance.kp_instance.guid}"
}
resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name = "abuck4"
  resource_instance_id = "${ibm_resource_instance.cos_instance.id}"
  region_location = "us-south"
  storage_class = "flex"
  key_protect = "${data.ibm_kp_key.test.keys.0.crn}"
}