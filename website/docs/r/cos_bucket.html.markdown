---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket"
description: |-
  Manages IBM CloudObject Storage Bucket.
---

# ibm\_cos_bucket

Creates an IBM Cloud Object Storage bucket. It also allows object storage buckets to be updated and deleted. The ibmcloud_api_key used by Terraform must have been granted sufficient IAM rights to create and modify IBM Cloud Object Storage buckets and have access to the Resource Group the Cloud Object Storage bucket will be associated with. See https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-iam for more details on setting IAM and Access Group rights to manage COS buckets.

## Example Usage

```hcl
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
  plan              = "graduated-tier "
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

resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = "a-flex-bucket-at-us-south"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "flex"
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

resource "ibm_cos_bucket" "flex-us-south-firewall" {
  bucket_name           = "a-flex-bucket-at-us-south"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  single_site_location  = "sjc04"
  storage_class         = "flex"
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

### Configure retention rule on COS bucket
resource "ibm_cos_bucket" "retention_cos" {
  bucket_name          = "a-bucket-retention"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "jp-tok"
  storage_class        = standard
  force_delete        = true
  retention_rule {
    default = 1
    maximum = 1
    minimum = 1
    permanent = false
  }
}

```

## Argument Reference

The following arguments are supported:

* `bucket_name` - (Required, string) The name of the bucket.
* `resource_instance_id` - (Required, string) The id of Cloud Object Storage instance.
* `key_protect` - (Optional, bool) CRN of the Key Protect instance where there a root key is already provisioned. Authorization required: [Docs](https://cloud.ibm.com/docs/services/cloud-object-storage?topic=cloud-object-storage-encryption#grant-service-authorization)
* `single_site_location` - (Optional,string) Location if single site bucket is desired. Accepted values: 'ams03', 'che01', 'hkg02', 'mel01', 'mex01', 'mil01', 'mon01', 'osl01', 'par01', 'sjc04', 'sao01', 'seo01', 'sng01', 'tor01' Conflicts with: `region_location`, `cross_region_location`
* `region_location` - (Optional,string) Location if regional bucket is desired. Accepted values: 'au-syd', "ca-tor", 'eu-de', 'eu-gb', 'jp-tok', 'us-east', 'us-south' Conflicts with: `single_site_location`, `cross_region_location`
* `cross_region_location` - (Optional,string) Location if cross regional bucket is desired. Accepted values: 'us', 'eu', 'ap' Conflicts with: `single_site_location`, `region_location`
* `allowed_ip` - (Optional, list of strings) List of IPv4 or IPv6 addresses in CIDR notation to be affected by firewall in CIDR notation is supported. 
* Nested `activity_tracking` block have the following structure:
	*	`read_data_events` : (Optional, array) Enables sending log data to Activity Tracker and LogDNA to provide visibility into object read and write events.
	*	`write_data_events` : (Optional,bool) If set to true, all object write events (i.e. uploads) will be sent to Activity Tracker.
	*	`activity_tracker_crn` : (Required, string) Required the first time activity_tracking is configured.
* Nested `metrics_monitoring` block have the following structure:
	*	`usage_metrics_enabled` : (Optional,bool) If set to true, all usage metrics (i.e. bytes_used) will be sent to the monitoring service.
	*	`request_metrics_enabled` : (Optional,bool) If set to true, all request metrics (i.e. ibm_cos_bucket_all_request) will be sent to the monitoring service @1mins granulatiy.
  *	`metrics_monitoring_crn` : (Required, string) Required the first time metrics_monitoring is configured. The instance of IBM Cloud Monitoring that will receive the bucket metrics.
* **Note** - For now request metrics support is enabled only in SJC04 (single_site_location) (only through API/SDK - NO UI support yet). Later we will add the support for other regions(Notified soon).The bucket-level request metrics monitoring service is disabled in the container vault for each region except sjc04.

* **Note** - One of the location option must be present.
* `storage_class` - (Required, string) Storage class of the bucket. Accepted values: 'standard', 'vault', 'cold', 'flex', 'smart'.
* `endpoint_type` - (Optional, string) The type of the endpoint (public or private) to be used for buckets. Default value is `public`.
* `force_delete` - (Optional, bool) Since Default value set to True, it will delete all the objects in the COS Bucket and then delete the bucket.  Default value is `true`.
    * **Note** - `force_delete` will timeout on buckets with a large amount of objects.  24 hours before you delete the bucket you can set an expire rule to remove all files over a day old.  

* **Note** - Both archive_rule and expire_rule must be managed by terraform as they use the same lifecycle configuration. If user creates any of the rule outside of terraform using CLI/UI, you may see unexpected diff such as  removal of any of the rule or one rule overrides another , the policy may not match as expected due to API limitation because the LifeCycle is a single API request for both Archive and Expire.
* Nested `archive_rule` block have the following structure:
	*	`rule_id` : (Optional, Computed, string) Unique identifier for the rule. Archive rules allow you to set a specific time frame after which objects transition to the archive. 
	*	`enable` : (Required, bool) Specifies archive rule status either enable or disable for a bucket.
  *	`days` : (Required, string) Specifies the number of days when the specific rule action takes effect.
  *	`type` : (Required, string) Specifies the storage class/archive type to which you want the object to transition. It can be Glacier or Accelerated.
    * **Note** - Archive is available in certain regions only. See Integrated Services for more details-https://cloud.ibm.com/docs/cloud-object-storage/basics?topic=cloud-object-storage-service-availability

* Nested `expire_rule` block have the following structure:
	*	`rule_id` : (Optional, Computed, string) Unique identifier for the rule. Expire rules allow you to set a specific time frame after which objects are deleted. 
	*	`enable` : (Required, bool) Specifies expire rule status either enable or disable for a bucket.
  *	`days`   : (Required, string) Specifies the number of days when the specific rule action takes effect.
  *	`prefix` : (Optional, string) Specifies a prefix filter to apply to only a subset of objects with names that match the prefix.

* Nested `retention_rule` block have the following structure:
  * `default` : (Required, int) default retention period are defined by this policy and apply to   all objects in the bucket.
  * `maximum` : (Required, int) Specifies maximum duration of time an object can be kept unmodified in the bucket.
  * `minimum` : (Required, int) Specifies minimum duration of time an object must be kept unmodified in the bucket.
  * `permanent` : (Optional, bool) Specifies a permanent retention status either enable or disable for a bucket.

    * **Note**
     - Retention policies cannot be removed. For a new bucket, ensure that you are creating the bucket in a supported region.See Integrated Services for more details-https://cloud.ibm.com/docs/cloud-object-storage/basics?topic=cloud-object-storage-service-availability
     - The minimum retention period must be less than or equal to the default retention period, which in turn must be less than or equal to the maximum retention period.
     - Permanent retention can only be enabled at a IBM Cloud Object Storage bucket level with retention policy enabled and users are able to select the permanent retention period option during object uploads. Once enabled, this process can't be reversed and objects uploaded that use a permanent retention period cannot be deleted. It's the responsibility of the users to validate at their end if there's a legitimate need to permanently store objects by using Object Storage buckets with a retention policy.
     - force deleting the bucket will not work if any object is still under retention. As Objects cannot be deleted or overwritten until the retention period has expired and all the legal holds have been removed.

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

## Import

The `ibm_cos_bucket` resource can be imported using the `id`. The ID is formed from the `CRN` (Cloud Resource Name), the `bucket type` which must be `ssl` for single_site_location, `rl` for region_location or `crl` for cross_region_location, the bucket location and the endpoint type (public or private). The `CRN` and bucket location can be found on the portal.

id = $CRN:meta:$buckettype:$bucketlocation

```
$ terraform import ibm_cos_bucket.mybucket <crn>

$ terraform import ibm_cos_bucket.mybucket crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:mybucketname:meta:crl:eu:public
```
