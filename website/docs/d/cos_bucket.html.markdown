---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket"
description: |-
  Get information about IBM Cloud Object Storage bucket.
---

# ibm_cos_bucket

Creates an IBM Cloud Object Storage bucket. It also allows object storage buckets to be updated and deleted. The ibmcloud_api_key used by Terraform must have been granted sufficient IAM rights to create and modify IBM Cloud Object Storage buckets and have access to the Resource Group the Cloud Object Storage bucket will be associated with. See https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-iam for more details on setting IAM and Access Group rights to manage COS buckets.

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

## Argument reference
Review the argument references that you can specify for your data source. 

- `bucket_name` - (Required, String) The name of the bucket.
- `bucket_region` - (Required, String) The region of the bucket.
- `bucket_type` - (Required, String) The type of the bucket. Supported values are `single_site_location`, `region_location`, and `cross_region_location`.
- `endpoint_type` - (Optional, String) The type of the endpoint either public or private to be used for the buckets. Default value is `public`.
- `resource_instance_id` - (Required, String) The ID of the IBM Cloud Object Storage service instance for which you want to create a bucket.
- `storage_class`- (Required, String)  Storage class of the bucket. Supported values are `standard`, `vault`, `cold`, `flex`, `smart`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `allowed_ip`-  (String) List of `IPv4` or `IPv6` addresses in CIDR notation to be affected by firewall.
- `activity_tracking` (List) Nested block with the following structure.

  Nested scheme for `activity_tracking`:
  - `activity_tracker_crn` - (String) The first time activity_tracking is configured.
  - `read_data_events` - (Array)  Enables sending log data to Activity Tracker to provide visibility into an object read and write events.
  - `write_data_events`- (Bool) If set to **true**, all object write events (that is `uploads`) is sent to Activity Tracker.
- `archive_rule` (List) Nested block with the following structure.

  Nested scheme for `archive_rule`:
  - `days` - (String)  Specifies the number of days when the specific rule action takes effect.
  - `enable`- (Bool) Specifies archive rule status either `enable` or `disable` for a bucket.
  - `type` - (String)  Specifies the storage class or archive type to which you want the object to transition. Supported values are `Glacier` or `Accelerated`.
  - `rule_id` - (String)  Unique identifier for the rule. Archive rules allow you to set a specific time frame after which objects transition to archive.
- `crn` - (String) The CRN of the bucket.
- `cross_region_location` - (String) The location to create a cross-regional bucket.
- `expire_rule` (List) Nested block with the following structure.

  Nested scheme for `expire_rule`:
  - `days` - (String)  Specifies the number of days when the specific rule action takes effect.
  - `enable`- (Bool) Specifies expire rule status either `enable` or `disable` for a bucket.
  - `prefix` - (String)  Specifies a prefix filter to apply to only a subset of objects with names that match the prefix.
  - `rule_id` - (String)  Unique identifier for the rule. Expire rules allow you to set a specific time frame after which objects are deleted.
- `id` - (String) The ID of the bucket. 
- `key_protect` - (String) The CRN of the IBM Key Protect instance where a root key is already provisioned.
- `metrics_monitoring`- (List) Nested block with the following structure.
   
   Nested scheme for `metrics_monitoring`:
    - `metrics_monitoring_crn` - (String) The first time `metrics_monitoring` is configured. The instance of IBM Cloud monitoring that will receive the bucket metrics.
    -	`request_metrics_enabled` - (Bool) If set to `true`, all request metrics `ibm_cos_bucket_all_request` is sent to the monitoring service at 1 minute (`@1mins`) granularity.
    - `usage_metrics_enabled`- (Bool) If set to **true**, all usage metrics (that is `bytes_used`) is sent to the monitoring service.
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
- `hard_quota` - (String) Maximum bytes for the bucket.
- `single_site_location` - (String) The location to create a single site bucket.
- `storage_class` - (String) The storage class of the bucket.
