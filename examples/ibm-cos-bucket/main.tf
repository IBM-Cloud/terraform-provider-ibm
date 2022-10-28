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
  storage_class         = var.standard_storage_class
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
  storage_class        = var.standard_storage_class
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
  storage_class         = var.standard_storage_class
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
//Replication
resource "ibm_resource_instance" "cos_instance_source" {
  name              = "cos-instance-src"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "cos_bucket_source" {
  bucket_name           = "sourcetest"
  resource_instance_id = ibm_resource_instance.cos_instance_source.id
  region_location      = var.regional_loc
  storage_class         = var.standard_storage_class
  object_versioning {
    enable  = true
  }
}

resource "ibm_resource_instance" "cos_instance_destination" {
  name              = "cos-instance-dest"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "cos_bucket_destination" {
  bucket_name           = "desttest"
  resource_instance_id = ibm_resource_instance.cos_instance_destination.id
  region_location      = var.regional_loc
  storage_class         = var.standard_storage_class
  object_versioning {
    enable  = true
  }
}

resource "ibm_cos_bucket" "cos_bucket_destination_1" {
  bucket_name           = "desttest01"
  resource_instance_id = ibm_resource_instance.cos_instance_destination.id
  region_location      = var.regional_loc
  storage_class         = var.standard_storage_class
  object_versioning {
    enable  = true
  }
}

resource "ibm_iam_authorization_policy" "policy" {
  roles                  = [
      "Writer",
  ]
  subject_attributes {
    name  = "accountId"
    value = "12345"
  }
  subject_attributes {
    name  = "serviceName"
    value = "cloud-object-storage"
  }
  subject_attributes {
    name  = "serviceInstance"
    value = ibm_resource_instance.cos_instance_source.guid
  }
  subject_attributes {
    name  = "resource"
    value = ibm_cos_bucket.cos_bucket_source.bucket_name
  }
  subject_attributes {
    name  = "resourceType"
    value = "bucket"
  }
  resource_attributes {
    name     = "accountId"
    operator = "stringEquals"
    value    = "12345"
  }
  resource_attributes {
    name     = "serviceName"
    operator = "stringEquals"
    value    = "cloud-object-storage"
  }
  resource_attributes { 
    name  =  "serviceInstance"
    operator = "stringEquals"
    value =  ibm_resource_instance.cos_instance_destination.guid
  }
  resource_attributes { 
    name  =  "resource"
    operator = "stringEquals"
    value =  ibm_cos_bucket.cos_bucket_destination.bucket_name
  }
  resource_attributes { 
    name  =  "resourceType"
    operator = "stringEquals"
    value =  "bucket" 
  }
}

resource "ibm_iam_authorization_policy" "policy1" {
  roles                  = [
      "Writer",
  ]
  subject_attributes {
    name  = "accountId"
    value = "12345"
  }
  subject_attributes {
    name  = "serviceName"
    value = "cloud-object-storage"
  }
  subject_attributes {
    name  = "serviceInstance"
    value = ibm_resource_instance.cos_instance_source.guid
  }
  subject_attributes {
    name  = "resource"
    value = ibm_cos_bucket.cos_bucket_source.bucket_name
  }
  subject_attributes {
    name  = "resourceType"
    value = "bucket"
  }
  resource_attributes {
    name     = "accountId"
    operator = "stringEquals"
    value    = "12345"
  }
  resource_attributes {
    name     = "serviceName"
    operator = "stringEquals"
    value    = "cloud-object-storage"
  }
  resource_attributes { 
    name  =  "serviceInstance"
    operator = "stringEquals"
    value =  ibm_resource_instance.cos_instance_destination.guid
  }
  resource_attributes { 
    name  =  "resource"
    operator = "stringEquals"
    value =  ibm_cos_bucket.cos_bucket_destination_1.bucket_name
  }
  resource_attributes { 
    name  =  "resourceType"
    operator = "stringEquals"
    value =  "bucket" 
  }
}


resource "ibm_cos_bucket_replication_rule" "cos_bucket_repl" {
  depends_on = [
      ibm_iam_authorization_policy.policy, ibm_iam_authorization_policy.policy1
  ]
  bucket_crn      = ibm_cos_bucket.cos_bucket_source.crn
  bucket_location = ibm_cos_bucket.cos_bucket_source.region_location
  replication_rule {
    enable = true
    prefix = var.replicate_prefix
    priority = var.replicate_priority
    deletemarker_replication_status = var.delmarkerrep_status
    destination_bucket_crn = ibm_cos_bucket.cos_bucket_destination.crn
  }
  replication_rule {
    enable = true
    priority = "2"
    deletemarker_replication_status = var.delmarkerrep_status
    destination_bucket_crn = ibm_cos_bucket.cos_bucket_destination_1.crn
  }
}

//HPCS - standard plan
resource ibm_hpcs hpcs {
  location             = var.location
  name                 = "hpcs-instance"
  plan                 = var.hpcs_plan
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
resource "ibm_iam_authorization_policy" "policy2" {
  source_service_name = "cloud-object-storage"
  target_service_name = "hs-crypto"
  roles               = ["Reader"]
}
resource "ibm_kms_key" "key" {
  instance_id  = ibm_hpcs.hpcs.guid
  key_name     = var.hpcs_key_name
  standard_key = false
  force_delete = true
}

resource "ibm_cos_bucket" "hpcs-enabled" {
  depends_on           = [ibm_iam_authorization_policy.policy2]
  bucket_name          = var.bucket_name
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class
  key_protect          = ibm_kms_key.key.id
}

//HPCS - UKO plan
resource "ibm_cos_bucket" "hpcs-uko-enabled" {
  depends_on           = [ibm_iam_authorization_policy.policy2]
  bucket_name          = var.bucket_name
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class
  key_protect           = var.hpcs_uko_rootkeycrn
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

//One Rate COS plan

resource "ibm_resource_instance" "cos_instance_onerate" {
  name              = "cos-instance-onerate"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "cos-one-rate-plan"
  location          = "global"
}
resource "ibm_cos_bucket" "cos_bucket_onerate" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance_onerate.id
  region_location       = var.regional_loc
  storage_class         = var.onerate_storage_class
  }
  