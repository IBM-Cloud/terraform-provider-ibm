---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket"
description: |-
  Get information about IBM CloudObject Storage Bucket.
---

# ibm\_cos_bucket

Creates an IBM Cloud Object Storage bucket. It also allows object storage buckets to be updated and deleted. The ibmcloud_api_key used by Terraform must have been granted sufficient IAM rights to create and modify IBM Cloud Object Storage buckets and have access to the Resource Group the Cloud Object Storage bucket will be associated with. See https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-iam for more details on setting IAM and Access Group rights to manage COS buckets.

## Example Usage

```hcl
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

## Argument Reference

The following arguments are supported:

* `bucket_name` - (Required, string) The name of the bucket.
* `bucket_type` - (Required, string) The type of the bucket. Accepted values: single_site_location region_location cross_region_location
* `resource_instance_id` - (Required, string) The id of Cloud Object Storage instance.
* `bucket_region` - (Required, string) The region of the bucket.
* `storage_class` - (Required, string) Storage class of the bucket. Accepted values: 'standard', 'vault', 'cold', 'flex', 'smart'.
* `endpoint_type` - (Optional, string) The type of the endpoint (public or private) to be used for buckets. Default value is `public`.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the bucket.
* `crn` - The CRN of the bucket.
* `resource_instance_id` - The id of Cloud Object Storage instance.
* `key_protect` - CRN of the Key Protect instance where there a root key is already provisioned.
* `single_site_location` - Location if single site bucket was created.
* `region_location` - Location if regional bucket was created.
* `cross_region_location` - Location if cross regional bucket was created.
* `storage_class` - Storage class of the bucket.
* `allowed_ip` - List of IPv4 or IPv6 addresses in CIDR notation to be affected by firewall.
* Nested `activity_tracking` block have the following structure:
	*	`read_data_events` : (Optional, array) Enables sending log data to Activity Tracker and LogDNA to provide visibility into object read and write events.
	*	`write_data_events` : (Optional,bool) If set to true, all object write events (i.e. uploads) will be sent to Activity Tracker.
	*	`activity_tracker_crn` : (Required, string) Required the first time activity_tracking is configured.
* Nested `metrics_monitoring` block have the following structure:
	*	`usage_metrics_enabled` : (Optional,bool) If set to true, all usage metrics (i.e. bytes_used) will be sent to the monitoring service.
	*	`request_metrics_enabled` : (Optional,bool) If set to true, all request metrics (i.e. ibm_cos_bucket_all_request) will be sent to the monitoring service @1mins granulatiy.
	*	`metrics_monitoring_crn` : (Required, string) Required the first time metrics_monitoring is configured. The instance of IBM Cloud Monitoring that will receive the bucket metrics.
* Nested `archive_rule` block have the following structure:
	*	`rule_id` : (Optional, Computed, string) Unique identifier for the rule. Archive rules allow you to set a specific time frame after which objects transition to the archive.
	*	`enable` :* (Required, bool) (Required) Specifies archive rule status either enable or disable for a bucket.
	*	`days` :* (Required, string)  Specifies the number of days when the specific rule action takes effect.
	*	`type` :* (Required, string) Specifies the storage class/archive type to which you want the object to transition. It can be Glacier or Accelerated.
* Nested `expire_rule` block have the following structure:
	*	`rule_id` : (Optional, Computed, string) Unique identifier for the rule. Expire rules allow you to set a specific time frame after which objects are deleted.
	*	`enable` :* (Required, bool) (Required) Specifies expire rule status either enable or disable for a bucket.
	*	`days`   :* (Required, string)  Specifies the number of days when the specific rule action takes effect.
	*	`prefix` :* (Optional, string) Specifies a prefix filter to apply to only a subset of objects with names that match the prefix.

* Nested `retention_rule` block have the following structure:
	*	`default` : (Required, int) default retention period are defined by this policy and apply to all objects in the bucket.
	*	`maximum` : (Required, int) Specifies maximum duration of time an object can be kept unmodified in the bucket.
	*	`minimum` : (Required, int) Specifies minimum duration of time an object must be kept unmodified in the bucket.
	*	`permanent` : (Optional, bool) Specifies a permanent retention status either enable or disable for a bucket.
