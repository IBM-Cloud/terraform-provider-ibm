---
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket"
sidebar_current: "docs-ibm-datasource-cos-bucket"
description: |-
  Get information about IBM CloudObject Storage Bucket.
---

# ibm\_cos_bucket

Get information about already existing buckets.

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

The following attributes are exported:

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
	*	`activity_tracking.read_data_events` : (Optional, array) Enables sending log data to Activity Tracker and LogDNA to provide visibility into object read and write events.
	*	`activity_tracking.write_data_events` : (Optional,bool) If set to true, all object write events (i.e. uploads) will be sent to Activity Tracker.
	*	`activity_tracking.activity_tracker_crn` : (Required, string) Required the first time activity_tracking is configured.
* Nested `metrics_monitoring` block have the following structure:
	*	`metrics_monitoring.usage_metrics_enabled` : (Optional,bool) If set to true, all usage metrics (i.e. bytes_used) will be sent to the monitoring service.
	*	`metrics_monitoring.metrics_monitoring_crn` : (Required, string) Required the first time metrics_monitoring is configured. The instance of IBM Cloud Monitoring that will receive the bucket metrics.
