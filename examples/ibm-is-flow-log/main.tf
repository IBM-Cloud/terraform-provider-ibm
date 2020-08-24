
data "ibm_resource_group" "cos_group" {
  name = var.resource_group
}

data "ibm_is_instance" "ds_instance" {
  name        = "vpc1-instance"
}

resource "ibm_resource_instance" "instance1" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "bucket1" {
   bucket_name          = "us-south-bucket-vpc1"
   resource_instance_id = ibm_resource_instance.instance1.id
   region_location = var.region
   storage_class = "standard"
}

resource ibm_is_flow_log test_flowlog {
  depends_on = [ibm_cos_bucket.bucket1]
  name = "test-instance-flow-log"
  target = data.ibm_is_instance.ds_instance.id
  active = true
  storage_bucket = ibm_cos_bucket.bucket1.bucket_name
}

