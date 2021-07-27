data "ibm_resource_group" "cos_group" {
  name = var.resource_group_name
}

resource "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.region_location
  storage_class         = var.storage_class
}

resource "ibm_cos_bucket_object" "plaintext" {
  bucket_crn      = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  content         = "Hello World"
  key             = "plaintext.txt"
}

resource "ibm_cos_bucket_object" "base64" {
  bucket_crn      = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  content_base64  = "RW5jb2RlZCBpbiBiYXNlNjQ="
  key             = "base64.txt"
}

resource "ibm_cos_bucket_object" "file" {
  bucket_crn      = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  content_file    = "${path.module}/helper/object.json"
  key             = "file.json"
  etag            = filemd5("${path.module}/helper/object.json")
}
