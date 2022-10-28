---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket"
description: 
  "Manages IBM Cloud Object Storage bucket."
---

# ibm_cos_bucket
Create or delete an IBM Cloud Object Storage bucket. The bucket is used to store your data. For more information, about configuration options, see [Create some buckets to store your data](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-getting-started-cloud-object-storage#gs-create-buckets). 

To create a bucket, you must provision an IBM Cloud Object Storage instance first by using the [`ibm_resource_instance`](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-resource-mgmt-resources#resource-instance) resource.

## Example usage
The following example creates an instance of IBM Cloud Object Storage, IBM Cloud Activity Tracker, and IBM Cloud Monitoring. Then, multiple buckets are created and configured to send audit events and metrics to your service instances.

```terraform
data "ibm_resource_group" "cos_group" {
  name = "cos-resource-group"
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
  location          = "us-south"
}
resource "ibm_resource_instance" "metrics_monitor" {
  name              = "metrics_monitor"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "sysdig-monitor"
  plan              = "graduated-tier"
  location          = "us-south"
  parameters        = {
    default_receiver = true
  }
}

resource "ibm_cos_bucket" "standard-ams03" {
  bucket_name          = "a-standard-bucket-at-ams"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  single_site_location = "ams03"
  storage_class        = "standard"
}

resource "ibm_cos_bucket" "smart-us-south" {
  bucket_name          = "a-smart-bucket-at-us-south"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "smart"
}

resource "ibm_cos_bucket" "cold-ap" {
  bucket_name           = "a-cold-bucket-at-ap"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  cross_region_location = "ap"
  storage_class         = "cold"
}

resource "ibm_cos_bucket" "standard-ams03-firewall" {
  bucket_name           = "a-standard-bucket-at-ams-firewall"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  single_site_location  = "sjc04"
  storage_class         = "standard"
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

resource "ibm_cos_bucket" "smart-us-south-firewall" {
  bucket_name           = "a-smart-bucket-at-us-south"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  single_site_location  = "sjc04"
  storage_class         = "smart"
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

resource "ibm_cos_bucket" "cold-ap-firewall" {
  bucket_name           = "a-cold-bucket-at-ap"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  single_site_location  = "sjc04"
  storage_class         = "cold"
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

### Configure archive and expire rules on COS Bucket

resource "ibm_cos_bucket" "archive_expire_rule_cos" {
  bucket_name          = "a-bucket-archive-expire"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "standard"
  force_delete         = true
  archive_rule {
    rule_id = "a-bucket-arch-rule"
    enable  = true
    days    = 0
    type    = "GLACIER"
  }
  expire_rule {
    rule_id = "a-bucket-expire-rule"
    enable  = true
    days    = 30
    prefix  = "logs/"
  }
}

### Configure expire rule to prepare a COS Bucket with a large number of objects for deletion

resource "ibm_cos_bucket" "expire_rule_cos" {
  bucket_name          = "a-bucket-expire"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "standard"
  force_delete         = true
  expire_rule {
    rule_id = "a-bucket-expire-rule"
    enable  = true
    days    = 1
  }
}

### Configure expire date/days with non current version expiration enabled on COS bucket

resource "ibm_cos_bucket" "expirebucket" {
  bucket_name          = "a-bucket-expiredat"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "standard"
  force_delete         = true
  object_versioning {
    enable  = true
  }
  expire_rule {
    rule_id = "a-bucket-expire-rule"
    enable  = true
    date    = "2021-11-18"
    prefix  = "logs/"
  }
  noncurrent_version_expiration {
    rule_id = "my-rule-id-bucket-ncversion"
    enable  = true
    prefix  = ""
    noncurrent_days = 1
  }
}

### Configure clean up expired object delete markers  on COS bucket

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = "a-bucket-expireddelemarkertest"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = "us-south"
  storage_class         = "standard"
  object_versioning {
    enable  = true
  }
  expire_rule {
    rule_id = "my-rule-id-bucket-expired"
    enable  = true
    expired_object_delete_marker = true
  }
  noncurrent_version_expiration {
    rule_id = "my-rule-id-bucket-ncversion"
    enable  = true
    prefix  = ""
    noncurrent_days = 1
  }
}

### Configure abort incomplete multipart upload on COS bucket

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = "a-bucket-multipartupload"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = "us-south"
  storage_class         = "standard"
  abort_incomplete_multipart_upload_days {
    rule_id = var.abort_mpu_ruleid
    enable  = true
    prefix  = ""
    days_after_initiation = 1
  }
}

### Configure retention rule on COS bucket

resource "ibm_cos_bucket" "retention_cos" {
  bucket_name          = "a-bucket-retention"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "jp-tok"
  storage_class        = standard
  hard_quota           = 1024
  force_delete        = true
  retention_rule {
    default = 1
    maximum = 1
    minimum = 1
    permanent = false
  }
}

### Configure object versioning on COS bucket

resource "ibm_cos_bucket" "objectversioning" {
  bucket_name           = "a-bucket-versioning"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = "us-east"
  storage_class         = var.storage
  hard_quota            = 11
  object_versioning {
    enable  = true
  }
}

```

# cos satellite bucket

Create or delete an COS satellite bucket. See the architecture of COS Satellite 
https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-about-cos-satellite for more details. We are using existing cos instance to create bucket , so no need to create any cos instance via a terraform. Cos satellite does not support all features see the section **What features are currently supported?** in https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-about-cos-satellite.
IBM Satellite documentation https://cloud.ibm.com/docs/satellite?topic=satellite-getting-started. We are supporting object versioning and expiration features as of now. Firewall is not supported yet.

## Example usage

```terraform
data "ibm_resource_group" "group" {
    name = "Default"
}

resource "ibm_satellite_location" "create_location" {
  location          = var.location
  zones             = var.location_zones
  managed_from      = var.managed_from
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = "cos-sat-terraform"
  resource_instance_id  = data.ibm_resource_instance.cos_instance.id
  satellite_location_id  = data.ibm_satellite_location.create_location.id
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
```


# Key Protect enabled COS bucket

Create or delete an COS bucket with a key protect root key.For more details about key protect see https://cloud.ibm.com/docs/key-protect?topic=key-protect-about  .We  need to create and manage root key using  **ibm_kms_key** resource. We are using existing cos instance to create bucket , so no need to create any cos instance via a terraform. https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/kms_key

## Example usage

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
resource "ibm_kms_key" "test" {
  instance_id  = ibm_resource_instance.kms_instance.guid
  key_name     = "key-name"
  standard_key = false
  force_delete =true
}
resource "ibm_iam_authorization_policy" "policy" {
	source_service_name = "cloud-object-storage"
	target_service_name = "kms"
	roles               = ["Reader"]
}
resource "ibm_cos_bucket" "smart-us-south" {
  depends_on           = [ibm_iam_authorization_policy.policy]
  bucket_name          = "atest-bucket"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "smart"
  key_protect          = ibm_kms_key.test.id
}
```


# HPCS enabled COS bucket

Create or delete a COS bucket with a Hyper Protect Crypto Services (HPCS) root key.For more details about HPCS see https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started .To enable HPCS on a COS bucket, an HPCS instance is required and needs to be initialized by loading the master key to create and manage HPCS keys. For more information on initializing the HPCS instance, see https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-initialize-hsm-recovery-crypto-unit. To create an HPCS instance using terraform, see https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/hpcs.


## Example usage

```terraform
resource ibm_hpcs hpcs {
  location             = "us-south"
  name                 = "test-hpcs"
  plan                 = "standard"
  units                = 2
  signature_threshold  = 1
  revocation_threshold = 1
  admins {
    name  = "admin1"
    key   = "/cloudTKE/1.sigkey"
    token = "<sensitive1234>"
  }
  admins {
    name  = "admin2"
    key   = "/cloudTKE/2.sigkey"
    token = "<sensitive1234>"
  }
}
resource "ibm_kms_key" "key" {
  instance_id  = ibm_hpcs.hpcs.guid
  key_name     = "key-name"
  standard_key = false
  force_delete = true
}

resource "ibm_iam_authorization_policy" "policy1" {
  source_service_name = "cloud-object-storage"
  target_service_name = "hs-crypto"
  roles               = ["Reader"]
}
resource "ibm_cos_bucket" "smart-us-south" {
  depends_on           = [ibm_iam_authorization_policy.policy]
  bucket_name          = "atest-bucket"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "smart"
  key_protect          = ibm_kms_key.key.id
}

```



# COS One-rate plan
One-rate is one of the plans for cloud object storage instance .The One Rate plan is best suited for active workloads with large amounts of outbound bandwidth relative to storage capacity.For more information, see https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-onerate&mhsrc=ibmsearch_a&mhq=One+rate

## Example usage

```terraform
resource "ibm_resource_instance" "cos_instance_onerate" {
  name              = "cos-instance-onerate"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "cos-one-rate-plan"
  location          = "global"
}
resource "ibm_cos_bucket" "cos_bucket_onerate" {
  bucket_name           = "bucket-name"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = "us-south"
  storage_class         = "onerate_active"
  }


```


## Argument reference
Review the argument references that you can specify for your resource. 

- `abort_incomplete_multipart_upload_days` (Optional,List) Nested block with the following structure.
  
  Nested scheme for `abort_incomplete_multipart_upload_days`:
  - `days_after_initiation` - (Optional, Integer) Specifies the number of days that govern the automatic cancellation of part upload. Clean up incomplete multi-part uploads after a period of time. Must be a value greater than 0 and less than 3650.
  - `enable` - (Required, Bool) A rule can either be `enabled` or `disabled`. A rule is active only when enabled.
  - `prefix` - (Optional, String)  A rule with a prefix will only apply to the objects that match. You can use multiple rules for different actions for different prefixes within the same bucket.
  - `rule_id` - (Optional, String) Unique identifier for the rule. Rules allow you to set a specific time frame after which objects are deleted. Set Rule ID for cos bucket.
- `allowed_ip` - (Optional, Array of string)  A list of IPv4 or IPv6 addresses in CIDR notation that you want to allow access to your IBM Cloud Object Storage bucket.
- `activity_tracking`- (List of objects) Object to enable auditing with IBM Cloud Activity Tracker - Optional - Configure your IBM Cloud Activity Tracker service instance and the type of events that you want to send to your service to audit activity against your bucket. For a list of supported actions, see [Bucket actions](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-at-events#at-actions-mngt-2).

  Nested scheme for `activity_tracking`:
  - `activity_tracker_crn`-  (Required, String) The CRN of your IBM Cloud Activity Tracker service instance that you want to send your events to. This value is required only when you configure your instance for the first time. 
  - `read_data_events`-  (Required, Bool)  If set to **true**, all read events against a bucket are sent to your IBM Cloud Activity Tracker service instance.
  - `write_data_events`-  (Required, Bool) If set to **true**, all write events against a bucket are sent to your IBM Cloud Activity Tracker service instance.
- `archive_rule` - (Required, List) Nested archive_rule block has following structure.
  
  Nested scheme for `archive_rule`:
  - `days` - (Required, Integer) Specifies the number of days when the specific rule action takes effect.
  - `enable` - (Required, Bool) Specifies archive rule status either `enable` or `disable` for a bucket.
  - `rule_id` -  (Optional, Computed, String) The unique ID for the rule. Archive rules allow you to set a specific time frame after the objects transition to the archive.
  - `type` - (Required, String) Specifies the storage class or archive type to which you want the object to transition. Allowed values are `Glacier` or `Accelerated`. 
  
    **Note:** 
    - Archive is available in certain regions only. For more information, see [Integrated Services](https://cloud.ibm.com/docs/cloud-object-storage/basics?topic=cloud-object-storage-service-availability).
    - Restoring object once archive is not supported yet.
- `bucket_name` - (Required, String) The name of the bucket.
- `cross_region_location` - (Optional, String) Specify the cross-regional bucket location. Supported values are `us`, `eu`, and `ap`. If you use this parameter, do not set `single_site_location` or `region_location` at the same time.
- `endpoint_type`- (Optional, String) The type of the endpoint either `public` or `private` or `direct` to be used for buckets. Default value is `public`.
- `expire_rule` - (Required, List) An expiration rule deletes objects after a defined period (from the object creation date). see [lifecycle actions](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-versioning). Nested expire_rule block has following structure.
  
  Nested scheme for `expire_rule`:
  - `days` - (Optional, Integer) Specifies the number of days when the specific rule action takes effect.
  - `date` - (Optional, String) After the specifies date , the current version of objects in your bucket expires.
  - `enable` - (Required, Bool) Specifies expire rule status either `enable` or `disable` for a bucket.
  - `expired_object_delete_marker` - (Optional, String) Expired object delete markers can be automatically cleaned up to improve performance in your bucket. This cannot be used alongside version expiration. This element for the Expiration action which will only remove delete markers that have no non-current versions at all & objects whose only version is a single delete marker.
  - `prefix` - (Optional, String) Specifies a prefix filter to apply to only a subset of objects with names that match the prefix.
  - `rule_id` -  (Optional, Computed, String) Unique ID for the rule. Expire rules allow you to set a specific time frame after which objects are deleted.

    **Note:** 
    - Both `archive_rule` and `expire_rule` must be managed by  Terraform as they use the same lifecycle configuration. If user creates any of the rule outside of  Terraform by using command line or console, you can see unexpected difference like removal of any of the rule or one rule overrides another. The policy cannot match as expected due to API limitations, as the lifecycle is a single API request for both archive and expire.
    - When versioning is enabled/suspended, regular object expiration will no longer remove objects, instead it will create a delete marker, unless the current version is already a delete marker, then nothing happens. If the only version of the object is a delete marker, then the delete marker is removed after X days, or on a specific date.
    - expired_object_delete_marker element can not be used in conjunction with other expiry action elements (Days or Date).
    - The expiry 3 action elements (Days, Date, ExpiredObjectDeleteMarker) are all mutually exclusive.Anyone parameter can apply among 3 (Days, Date, ExpiredObjectDeleteMarker) in expire_rule.
    - You cannot specify both a Days and ExpiredObjectDeleteMarker tag on the same rule. Specifying the Days tag will automatically perform ExpiredObjectDeleteMarker cleanup once delete markers are old enough to satisfy the age criteria. You can create a separate rule with only the tag ExpiredObjectDeleteMarker to clean up delete markers as soon as they become the only version.
- `force_delete`- (Optional, Bool) As the default value set to **true**, it will delete all the objects in the COS Bucket and then delete the bucket. 

    **Note:** `force_delete` will timeout on buckets with a large amount of objects. 24 hours before you delete the bucket you can set an expire rule to remove all the files over a day old.
- `hard_quota` - (Optional, Integer) Sets a maximum amount of storage (in bytes) available for a bucket. For more information, check the [cloud documention](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-quota).
- `key_protect` - (Optional, String) The CRN of the IBM Key Protect root key that you want to use to encrypt data that is sent and stored in IBM Cloud Object Storage. Before you can enable IBM Key Protect encryption, you must provision an instance of IBM Key Protect and authorize the service to access IBM Cloud Object Storage. For more information, see [Server-Side Encryption with IBM Key Protect or Hyper Protect Crypto Services (SSE-KP)](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-encryption).
- `metrics_monitoring`- (Object) to enable metrics tracking with IBM Cloud Monitoring - Optional- Set up your IBM Cloud Monitoring service instance to receive metrics for your IBM Cloud Object Storage bucket.

  Nested scheme for `metrics_monitoring`:
  - `metrics_monitoring_crn` - (Required, string) Required the first time `metrics_monitoring` is configured. The instance of IBM Cloud Monitoring receives the bucket metrics. 
  - `request_metrics_enabled` : (Optional, Bool) If set to **true**, all request metrics `ibm_cos_bucket_all_request` is sent to the monitoring service `@1mins` granulatiy.
  - `usage_metrics_enabled` : (Optional, Bool) If set to **true**, all usage metrics that is `bytes_used` is sent to the monitoring service.e.

    **Note:** 
    - Request metrics are supported in all regions and console has the support. For more details check the [cloud documention](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-mm-cos-integration).
    - One of the location option must be present. 
- `noncurrent_version_expiration` - (Required, List) lifecycle has a versioning related expiration action: non-current version expiration. This can remove old versions of objects after they've been non-current for a specified number of days which is specified with a NoncurrentDays parameter on the rule. see [lifecycle actions](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-versioning). Nested noncurrent_version_expiration block has following structure.

  Nested scheme for `noncurrent_version_expiration`:
  - `enable` - (Requried, Bool) A rule can either be `enabled` or `disabled`. A rule is active only when enabled.
  - `noncurrent_days` - (Optional, Integer) Configuration parameter in your policy that says how long to retain a non-current version before deleting it. Must be greater than 0.
  - `prefix` - (Optional, String) The rule applies to any objects with keys that match this prefix. You can use multiple rules for different actions for different prefixes within the same bucket.
  - `rule_id` - (Optional, String) Unique identifier for the rule. Rules allow you to remove versions from objects. Set Rule ID for cos bucket.
- `object_versioning` - (List) Object Versioning allows the COS user to keep multiple versions of an objet in a bucke to protect against accidental deletion or overwrites. With versioning, you can easilyrecover from both unintended user actions and application failure. Nested block have the following structure:

  Nested scheme for `object_versioning`:
  - `enable` : (Optional, Bool) Specifies Versioning status either enable or Suspended for the objects in the bucket.Default value set to false.

    **Note:**
    - Versioning allows multiple revisions of a single object to exist in the same bucket. Each version of an object can be queried, read, restored from an archived state, or deleted.
    - If cos bucket has versioning enabled and set to false, versioning will be suspended.
    - Versioning can only be suspended, we cannot disabled once after it is enabled.
    - To permanently delete individual versions of an object, a delete request must specify a version ID.
    - COS Object versioning and COS Bucket Protection `(WORM)` cannot be used together.
    - Containers with proxy configuration cannot use versioning and vice versa.
    - SoftLayer accounts cannot use versioning.
    - Currently, you cannot support `MFA_Delete`, that is a feature to add additional security to version delete.
- `region_location` - (Optional, String) The location of a regional bucket. Supported values are `au-syd`, `eu-de`, `eu-gb`, `jp-tok`, `us-east`, `us-south`, `ca-tor`, `jp-osa`, `br-sao`. If you set this parameter, do not set `single_site_location` or `cross_region_location` at the same time.
- `resource_instance_id` - (Required, String) The ID of the IBM Cloud Object Storage service instance for which you want to create a bucket.
- `retention_rule` - (List) Nested block have the following structure:
  
  Nested scheme for `retention rule`:
  - `default` - (Required, Integer) default retention period are defined by this policy and apply to all objects in the bucket.
  - `maximum` - (Required, Integer) Specifies maximum duration of time an object that can be kept unmodified in the bucket.
  - `minimum` - (Required, Integer) Specifies minimum duration of time an object must be kept unmodified in the bucket.
  - `permanent` : (Optional, Bool) Specifies a permanent retention status either enable or disable for a bucket.

    **Note:**
     - Retention policies cannot be removed. For a new bucket, ensure that you are creating the bucket in a supported region. For more information, see [Integrated Services](https://cloud.ibm.com/docs/cloud-object-storage/basics?topic=cloud-object-storage-service-availability).
     - The minimum retention period must be less than or equal to the default retention period, that in turn must be less than or equal to the maximum retention period.
     - Permanent retention can only be enabled at a IBM Cloud Object Storage bucket level with retention policy enabled and users are able to select the permanent retention period option during object uploads. Once enabled, this process can't be reversed and objects uploaded that use a permanent retention period cannot be deleted. It's the responsibility of the users to validate at their end if there's a legitimate need to permanently store objects by using Object Storage buckets with a retention policy.
     - force deleting the bucket will not work if any object is still under retention. As objects cannot be deleted or overwritten until the retention period has expired and all the legal holds have been removed.
- `single_site_location` - (Optional, String) The location for a single site bucket. Supported values are: `ams03`, `che01`, `hkg02`, `mel01`, `mex01`, `mil01`, `mon01`, `osl01`, `par01`, `sjc04`, `sao01`, `seo01`, `sng01`, and `tor01`. If you set this parameter, do not set `region_location` or `cross_region_location` at the same time.
- `storage_class` - (Optional, String) The storage class that you want to use for the bucket. Supported values are `standard`, `vault`, `cold` and `smart` for `standard` and `lite` COS plans, `onerate_active` for `cos-one-rate-plan` COS plan.For more information, about storage classes, see [Use storage classes](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-classes).`storage_class` should not be used with Satellite location id.
- `satellite_location_id` - (Optional, String) satellite location id. Provided by end users.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the bucket.
- `cross_region_location` - (String) The location if you created a cross-regional bucket.
- `id` - (String) The ID of the bucket. 
- `key_protect` - (String) The CRN of the IBM Key Protect instance that you use to encrypt your data in IBM Cloud Object Storage.
- `region_location` - (String) The location if you created a regional bucket.
- `resource_instance_id` - (String) The ID of IBM Cloud Object Storage instance. 
- `single_site_location` - (String) The location if you created a single site bucket.
- `storage_class` - (String) The storage class of the bucket.
- `s3_endpoint_public` - (String) Public endpoint for cos bucket.
- `s3_endpoint_private` - (String) Private endpoint for cos bucket.
- `s3_endpoint_direct` - (String) Direct endpoint for cos bucket.

## Import IBM COS Bucket
The `ibm_cos_bucket` resource can be imported by using the `id`. The ID is formed from the `CRN` (Cloud Resource Name), the `bucket type` which must be `ssl` for single_site_location, `rl` for region_location or `crl` for cross_region_location, and the bucket location. The `CRN` and bucket location can be found on the portal.

id = `$CRN:meta:$buckettype:$bucketlocation`

**Syntax**

```
$ terraform import ibm_cos_bucket.mybucket `$CRN:meta:$buckettype:$bucketlocation`

```

**Example**

```

$ terraform import ibm_cos_bucket.mybucket crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:mybucketname:meta:crl:eu:public

```

## Import COS Satelllite Bucket
The `cos satellite bucket` resource can be imported by using the `id`. The ID is formed from the `CRN` (Cloud Resource Name), the `satellite_location_id` which must be `sl` for satellite_location_id and the bucket location. The `CRN` and bucket location can be found on the portal.

id = `$CRN:meta:$buckettype:$bucketlocation`

**Example**

```

$ terraform import ibm_cos_bucket.cos_bucket crn:v1:staging:public:cloud-object-storage:satloc_dal_c8fctn320qtrspbisg80:a/81ee25188545f05150650a0a4ee015bb:a2deec95-0836-4720-bfc7-ca41c28a8c66:bucket:tf-listbuckettest:meta:sl:c8fctn320qtrspbisg80:public

```