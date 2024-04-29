---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket"
description: |-
  Get information about IBM Cloud Object Storage bucket.
---


# ibm_cos_bucket

Retrieves an IBM Cloud Object Storage bucket. It also allows object storage buckets to be updated and deleted. The ibmcloud_api_key used by Terraform must have been granted sufficient IAM rights to create and modify IBM Cloud Object Storage buckets and have access to the Resource Group the Cloud Object Storage bucket will be associated with. See https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-iam for more details on setting IAM and Access Group rights to manage COS buckets.

## Example usage

```terraform
data "ibm_resource_group" "cos_group" {
  name = "cos-resource-group"
}

data "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
}

data "ibm_cos_bucket" "standard-ams03" {
  bucket_name          = "a-standard-bucket-at-ams"
  resource_instance_id = data.ibm_resource_instance.cos_instance.id
  bucket_type          = "single_site_location"
  bucket_region        = "ams03"
}

output "bucket_private_endpoint" {
  value = data.ibm_cos_bucket.standard-ams03.s3_endpoint_private
}
```

# COS Satellite bucket

Retrieves a COS Satellite bucket. See https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-about-cos-satellite for more details on Object Storage for Satellite. 
We are using existing COS instance to create bucket so there is no need to create any COS instance via terraform

**Note:**
Object Storage for Satellite does not support all features, please refer to the documentation section [What features are currently supported?](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-about-cos-satellite#about-cos-satellite-supported) for a full list of supported features.
`Object Versioning`, `Object Expiration`, `Object Tagging` are supported, `Firewall` is not yet supported.

## Example usage

```terraform
data "ibm_satellite_location" "location" {
  location  = var.location
}

data "ibm_cos_bucket" "cos-bucket-sat" {
  bucket_name           = "cos-sat-terraform"
  resource_instance_id  = data.ibm_resource_instance.cos_instance.id
  satellite_location_id  = data.ibm_satellite_location.location.id
}
```

# ibm_cos_bucket_replication_rule

Retrieves information of replication configuration on an existing bucket. For more information about configuration options, see [Replicating objects](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-replication-overview). 

To configure a replication policy on a bucket, you must enable object versioning on both source and destination buckets by using the [Versioning objects](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-versioning).

# Key Protect Enabled COS bucket

Retrieves a COS bucket enabled with Key protect root key for data encryption.
https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/kms_key



# Hyper Protect Crypto Services (HPCS) Enabled COS bucket
Retrieves a COS bucket enabled with data encryption using root key that is  created and managed by Hyper Protect Crypto Services.
```
data "ibm_kms_key" "test" {
  instance_id = "guid-of-hs-crypto-instance"
  key_name = "name-of-key"
}
OR
data "ibm_kms_key" "test" {
  instance_id = "guid-of-hs-crypto-instance"
  alias = "alias_name"
}
OR
data "ibm_kms_key" "test" {
  instance_id = "guid-of-hs-crypto-instance"
  limit = 100
  key_name = "name-of-key"
}
resource "ibm_cos_bucket" "smart-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "smart"
  kms_key_crn          = data.ibm_kms_key.test.key.0.crn
}

```

# ibm_cos_object_lock_configuration

Retrieves an IBM Cloud Object Storage bucket Object Lock configuration set on the bucket. Allows Object Lock configuration to be updated or deleted. Object Lock cannot be disabled once enabled on a bucket..

## Example usage

```terraform
data "ibm_resource_group" "cos_group" {
  name = "cos-resource-group"
}

data "ibm_cos_bucket" "object_lock_bucket" {
  bucket_name          = "bucket-name"
  resource_instance_id = data.ibm_resource_instance.cos_instance.id
  bucket_type          = "region_location"
  bucket_region        = "bucket-region"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `bucket_name` - (Required, String) The name of the bucket.
- `bucket_region` - (Optional, String) The region of the bucket.
- `bucket_type` - (Optional, String) The type of the bucket. Supported values are `single_site_location`, `region_location`, and `cross_region_location`.
- `endpoint_type` - (Optional, String) The type of the endpoint either `public` or `private` or `direct` to be used for the buckets. Default value is `public`.
- `resource_instance_id` - (Required, String) The ID of the IBM Cloud Object Storage service instance for which you want to create a bucket.
- `storage_class`- (Optional, String)  Storage class of the bucket. Supported values are `standard`, `vault`, `cold`, `smart` for `standard` and `lite` COS plans, `onerate_active` for `cos-one-rate-plan` COS instance.
- `satellite_location_id` - (Optional, String) satellite location id. Provided by end users.
- `object_lock` - (Optional, String) Specifies Object Lock status.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 
- `allowed_ip`-  (String) List of `IPv4` or `IPv6` addresses in CIDR notation to be affected by firewall.
- `activity_tracking` (List) Nested block with the following structure.

  Nested scheme for `activity_tracking`:
  - `activity_tracker_crn` - (String) If `activity_tracker_crn` is populated, then enabled events are sent to the Activity Tracker instance specified and bucket management events are always enabled.
  **Note:**
  When the `activity_tracker_crn` is not populated, then enabled events are sent to the Activity Tracker instance associated to the container's location unless otherwise specified in the ATracker Router service configuration.
  - `read_data_events` - (Bool)  Enables sending log data to Activity Tracker to provide visibility into an object read and write events.
  - `write_data_events`- (Bool) If set to **true**, all object write events (that is `uploads`) is sent to Activity Tracker.
  - `management_events`- (Bool) If set to **true**, all the bucket management events are sent to Activity Tracker.
  **Note:**
  `management_events` field only applies if `activity_tracker_crn` is not populated.
- `archive_rule` (List) Nested block with the following structure.

  Nested scheme for `archive_rule`:
  - `days` - (String)  Specifies the number of days when the specific rule action takes effect.
  - `enable`- (Bool) Specifies archive rule status either `enable` or `disable` for a bucket.
  - `rule_id` - (String)  Unique identifier for the rule. Archive rules allow you to set a specific time frame after which objects transition to archive.
  - `type` - (String)  Specifies the storage class or archive type to which you want the object to transition. Supported values are `Glacier` or `Accelerated`.
- `abort_incomplete_multipart_upload_days` (List) Nested block with the following structure.
  
  Nested scheme for `abort_incomplete_multipart_upload_days`:
  - `days_after_initiation` - (Integer) Specifies the number of days that govern the automatic cancellation of part upload. Clean up incomplete multi-part uploads after a period of time. Must be a value greater than 0.
  - `enable` - (Bool) A rule can either be `enabled` or `disabled`. A rule is active only when enabled.
  - `prefix` - (String)  A rule with a prefix will only apply to the objects that match. You can use multiple rules for different actions for different prefixes within the same bucket.
  - `rule_id` - (String) Unique identifier for the rule. Rules allow you to set a specific time frame after which objects are deleted. Set Rule ID for cos bucket.
- `crn` - (String) The CRN of the bucket.
- `cross_region_location` - (String) The location to create a cross-regional bucket.
- `expire_rule` (List) Nested block with the following structure.

  Nested scheme for `expire_rule`:
  - `days` - (String)  Specifies the number of days when the specific rule action takes effect.
  - `date` - (String)  After the specifies date , the current version of objects in your bucket expires.
  - `enable`- (Bool) Specifies expire rule status either `enable` or `disable` for a bucket.
  - `expired_object_delete_marker` - (Bool) Expired object delete markers can be automatically cleaned up to improve performance in your bucket. This cannot be used alongside version expiration.
  - `prefix` - (String)  Specifies a prefix filter to apply to only a subset of objects with names that match the prefix.
  - `rule_id` - (String)  Unique identifier for the rule. Expire rules allow you to set a specific time frame after which objects are deleted.
- `hard_quota` - (String) Maximum bytes for the bucket.
- `id` - (String) The ID of the bucket.
- `kms_key_crn` - (String) The CRN of the IBM Key Protect instance where a root key is already provisioned. 
  **Note:**

 `key_protect` attribute has been renamed as `kms_key_crn` , hence it is recommended to all the new users to use `kms_key_crn`.Although the support for older attribute name `key_protect` will be continued for existing customers.

- `metrics_monitoring`- (List) Nested block with the following structure.
   
  Nested scheme for `metrics_monitoring`:
  - `metrics_monitoring_crn` - (String) If `metrics_monitoring_crn` is populated, then enabled events are sent to the Metrics Monitoring instance specified.
  **Note:**
 When the `metrics_monitoring_crn` is not populated, then enabled metrics are sent to the monitoring instance associated to the container's location unless otherwise specified in the Metrics Router service configuration.
  -	`request_metrics_enabled` - (Bool) If set to `true`, all request metrics `ibm_cos_bucket_all_request` is sent to the monitoring service at 1 minute (`@1mins`) granularity.
  - `usage_metrics_enabled`- (Bool) If set to **true**, all usage metrics (that is `bytes_used`) is sent to the monitoring service.
- `noncurrent_version_expiration` (List) Nested block with the following structure.
  
  Nested scheme for `noncurrent_version_expiration`:
  - `enable` - (Bool) A rule can either be `enabled` or `disabled`. A rule is active only when enabled.
  - `noncurrent_days` - (Int) Configuration parameter in your policy that says how long to retain a non-current version before deleting it. Must be greater than 0.
  - `prefix` - (String) The rule applies to any objects with keys that match this prefix. You can use multiple rules for different actions for different prefixes within the same bucket.
  - `rule_id` - (String) Unique identifier for the rule. Rules allow you to remove versions from objects. Set Rule ID for cos bucket.
- `object_versioning` - (List) Nestedblock have the following structure:

  Nested scheme for `object_verionining`:
  - `enable` - (String) Specifies versioning status either enable or suspended for the objects in the bucket.
- `region_location` - (String) The location to create a regional bucket.
- `resource_instance_id` - (String) The ID of {site.data.keyword.cos_full_notm}} instance. 
- `retention_rule` - (List) Nested block have the following structure:

  Nested scheme for `retention rule`:
  - `default` - (String) default retention period are defined by this policy and apply to all objects in the bucket.
  - `maximum` - (String) Specifies maximum duration of time an object can be kept unmodified in the bucket.
  - `minimum` - (String) Specifies minimum duration of time an object must be kept unmodified in the bucket.
  - `permanent` - (String) Specifies a permanent retention status either enable or disable for a bucket.
- `replication_rule`- (List) Nested block have the following structure:

  Nested scheme for `replication_rule`:
  - `rule_id`- (String) The rule id.
  - `enable`-  (Bool) Specifies whether the rule is enabled. Specify true for Enabling it  or false for Disabling it.
  - `prefix`- (String) An object key name prefix that identifies the subset of objects to which the rule applies.
  - `priority`- (Int) A priority is associated with each rule. The rule will be applied in a higher priority if there are multiple rules configured. The higher the number, the higher the priority
  - `deletemarker_replication_status`-  (Bool) Specifies whether Object storage replicates delete markers. Specify true for Enabling it  or false for Disabling it.
  - `destination_bucket_crn`-  (String) The CRN of your destination bucket that you want to replicate to.

- `object_lock_configuration`- (Required, List) Nested block have the following structure:
  
  Nested scheme for `object_lock_configuration`:
  - `object_lock_enabled`- (String) Indicates whether this bucket has an Object Lock configuration enabled. Defaults to Enabled. Valid values: Enabled.
  - `object_lock_rule`- (List) Object Lock rule has following arguement:
  
  Nested scheme for `object_lock_rule`:
  - `default_retention`- (Required) Configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket
  Nested scheme for `default_retention`:
  - `mode`- (String)  Default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Supported values: COMPLIANCE.
  - `days`- (Int) Specifies number of days after which the object can be deleted from the COS bucket.
  - `years`- (Int) Specifies number of years after which the object can be deleted from the COS bucket.
**Note:**

 Either days or years should be provided for default retention, both cannot be used simultaneoulsy.

 - `website_endpoint` - (String) Website endpoint, if the bucket is configured with a website. If not, this will be an empty string.

- `single_site_location` - (String) The location to create a single site bucket.
- `storage_class` - (String) The storage class of the bucket.
- `s3_endpoint_public` - (String) Public endpoint for cos bucket.
- `s3_endpoint_private` - (String) Private endpoint for cos bucket.
- `s3_endpoint_direct` - (String) Direct endpoint for cos bucket.
