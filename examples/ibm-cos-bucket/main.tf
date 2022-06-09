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
resource "ibm_resource_instance" "activity_tracker" {
  name              = "activity_tracker"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "logdnaat"
  plan              = "lite"
  location          =  var.regional_loc
}
resource "ibm_resource_instance" "metrics_monitor" {
  name              = "metrics_monitor"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "sysdig-monitor"
  plan              = "graduated-tier"
  location          = var.regional_loc
  parameters        = {
    default_receiver = true
  }
}
resource "ibm_cos_bucket" "standard-ams03" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  single_site_location  = var.single_site_loc
  storage_class         = var.storage
  hard_quota            = var.quota
  activity_tracking {
    read_data_events     = true
    write_data_events    = true
    activity_tracker_crn = ibm_resource_instance.activity_tracker.id
  }
  metrics_monitoring {
    usage_metrics_enabled  = true
    request_metrics_enabled = true
    metrics_monitoring_crn = ibm_resource_instance.metrics_monitor.id
  }
  allowed_ip = ["223.196.168.27", "223.196.161.38", "192.168.0.1"]
}

resource "ibm_cos_bucket" "lifecycle_rule_cos" {
  bucket_name          = var.bucket_name
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = var.regional_loc
  storage_class        = var.storage
  hard_quota           = var.quota
  archive_rule {
    rule_id = var.archive_ruleid
    enable  = true
    days    = var.archive_days
    type    = var.archive_types
  }
  expire_rule {
    rule_id = var.expire_ruleid
    enable  = true
    days    = var.expire_days
    prefix  = var.expire_prefix
  }
  retention_rule {
    default = var.default_retention
    maximum = var.maximum_retention
    minimum = var.minimum_retention
    permanent = false
  }
}

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.storage
  hard_quota            = var.quota
  object_versioning {
    enable  = true
  }
  abort_incomplete_multipart_upload_days {
    rule_id = var.abort_mpu_ruleid
    enable  = true
    prefix  = var.abort_mpu_prefix
    days_after_initiation = var.abort_mpu_days_init
  }
  expire_rule {
    rule_id = var.expire_ruleid
    enable  = true
    date    = var.expire_date
    prefix  = var.expire_prefix
  }
  noncurrent_version_expiration {
    rule_id = var.nc_exp_ruleid
    enable  = true
    prefix  = var.nc_exp_prefix
    noncurrent_days = var.nc_exp_days
  }
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

//Satellite Location
resource "ibm_cos_bucket" "cos_bucket_sat" {
  bucket_name           = var.bucket_name
  resource_instance_id  = "crn:v1:bluemix:public:cloud-object-storage:satloc_wdc_c8jh7hfw0ppoapdqrmpg:a/d0c259a490e4488c83b62707ad3f5182:756ad6b6-72a6-4e55-8c94-b02e51e708b3::"
  satellite_location_id  = var.satellite_location_id
  object_versioning {
    enable  = true
  }
  expire_rule {
    rule_id = "bucket-tf-rule1"
    enable  = false
    days    = 20
    prefix  = "logs/"
  }
}
