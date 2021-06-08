---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket"
description: |-
  Manages IBM Cloud Object Storage bucket.
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

### Configure object versioning on COS bucket

resource "ibm_cos_bucket" "objectversioning" {
  bucket_name           = "a-bucket-versioning"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = "us-east"
  storage_class         = var.storage
  object_versioning {
    enable  = true
  }
}

```


## Argument reference
Review the argument references that you can specify for your resource. 

- `activity_tracking`- (List of objects) Object to enable auditing with IBM Cloud Activity Tracker - Optional - Configure your IBM Cloud Activity Tracker service instance and the type of events that you want to send to your service to audit activity against your bucket. For a list of supported actions, see [Bucket actions](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-at-events#at-actions-mngt-2).

   Nested scheme for `activity_tracking`:
   - `activity_tracker_crn`-  (Required, String) The CRN of your IBM Cloud Activity Tracker service instance that you want to send your events to. This value is required only when you configure your instance for the first time.
   - `read_data_events`-  (Required, Bool)  If set to **true**, all read events against a bucket are sent to your IBM Cloud Activity Tracker service instance.
   - `request_metrics_enabled` : (Optional, Bool) If set to **true**, all request metrics `ibm_cos_bucket_all_request` is sent to the monitoring service `@1mins` granulatiy.
   - `write_data_events`-  (Required, Bool) If set to **true**, all write events against a bucket are sent to your IBM Cloud Activity Tracker service instance.
   - `usage_metrics_enabled` : (Optional, Bool) If set to **true**, all usage metrics that is `bytes_used` is sent to the monitoring service.
- `allowed_ip` - (Optional, Array of string)  A list of IPv4 or IPv6 addresses in CIDR notation that you want to allow access to your IBM Cloud Object Storage bucket.
- `archive_rule` - (Required, List) Nested archive_rule block has following structure.
    **Note** Archive is available in certain regions only. For more informaton, see [Integrated Services](https://cloud.ibm.com/docs/cloud-object-storage/basics?topic=cloud-object-storage-service-availability).
  
  Nested scheme for `archive_rule`:
  - `days` - (Required, String) Specifies the number of days when the specific rule action takes effect.
  - `enable` - (Required, Bool) Specifies archive rule status either `enable` or `disable` for a bucket.
  - `rule_id` -  (Optional, Computed, String) The unique ID for the rule. Archive rules allow you to set a specific time frame after the objects transition to the archive.
  - `type` - (Required, String) Specifies the storage class or archive type to which you want the object to transition. Allowed values are `Glacier` or `Accelerated`. **Note** Archive is available in certain regions only. For more information, see [Integrated Services](https://cloud.ibm.com/docs/cloud-object-storage/basics?topic=cloud-object-storage-service-availability).
- `bucket_name` - (Required, String) The name of the bucket.
- `cross_region_location` - (Optional, String) Specify the cross-regional bucket location. Supported values are `us`, `eu`, and `ap`. If you use this parameter, do not set `single_site_location` or `region_location` at the same time.
- `endpoint_type`- (Optional, String) The type of the endpoint either public or private to be used for buckets. Default value is `public`.
- `expire_rule` - (Required, List) Nested expire_rule block has following structure.

  Nested scheme for `expire_rule`:
  - `rule_id` -  (Optional, Computed, String) Unique ID for the rule. Expire rules allow you to set a specific time frame after which objects are deleted.
  - `enable` - (Required, Bool) Specifies expire rule status either `enable` or `disable` for a bucket.
  - `days` - (Required, String) Specifies the number of days when the specific rule action takes effect.
  - `prefix` - (Optional, String) Specifies a prefix filter to apply to only a subset of objects with names that match the prefix.

Both `archive_rule` and `expire_rule` must be managed by  Terraform as they use the same lifecycle configuration. If user creates any of the rule outside of  Terraform by using command line or console, you can see unexpected difference like removal of any of the rule or one rule overrides another. The policy cannot match as expected due to API limitations, as the lifecycle is a single API request for both archive and expire.
- `force_delete`- (Optional, Bool) As the default value set to **true**, it will delete all the objects in the COS Bucket and then delete the bucket. **Note:** `force_delete` will timeout on buckets with a large amount of objects. 24 hours before you delete the bucket you can set an expire rule to remove all the files over a day old. * **Note** Both `archive_rule` and `expire_rule` must be managed by Terraform as they use the same lifecycle configuration. If user creates any of the rule outside of Terraform by using command line, or console, you can see unexpected difference such as removal of any of the rule, or one rule overrides another, the policy may not match as expected due to API limitation because the lifecycle is a single API request for both archive and expire.
- `key_protect` - (Optional, String) The CRN of the IBM Key Protect root key that you want to use to encrypt data that is sent and stored in IBM Cloud Object Storage. Before you can enable IBM Key Protect encryption, you must provision an instance of IBM Key Protect and authorize the service to access IBM Cloud Object Storage. For more information, see [Server-Side Encryption with IBM Key Protect or Hyper Protect Crypto Services (SSE-KP)](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-encryption).
- `metrics_monitoring_crn` - (Required, string) Required the first time `metrics_monitoring` is configured. The instance of IBM Cloud Monitoring receives the bucket metrics. **Note** Request metrics are supported in all regions and console has the support. For more details check the [cloud documentiona](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-mm-cos-integration) **Note** One of the location option must be present.
- `metrics_monitoring`- (Object) to enable metrics tracking with IBM Cloud Monitoring - Optional- Set up your IBM Cloud Monitoring service instance to receive metrics for your IBM Cloud Object Storage bucket.

  Nested scheme for `metrics_monitoring`:
  - `usage_metrics_enabled` - (Optional, Bool) If set to **true**, all metrics are sent to your IBM Cloud Monitoring service instance.
  - `request_metrics_enabled` : (Optional, Bool) If set to **true**, all request metrics `ibm_cos_bucket_all_request` is sent to the monitoring service `@1mins` granulatiy.
- `object_versioning` - (List) Nested block have the following structure:

  Nested scheme for `object_versioning`:
  - `enable` : (Optional, Bool) Specifies Versioning status either enable or Suspended for the objects in the bucket.Default value set to false.

    **Note**
    - Versioning allows multiple revisions of a single object to exist in the same bucket. Each version of an object can be queried, read, restored from an archived state, or deleted.
    - If cos bucket has versioning enabled and set to false, versioning will be suspended.
    - Versioning can only be suspended, we cannot disabled once after it is enabled.
    - To permanently delete individual versions of an object, a delete request must specify a version ID.
    - Containers with object expiry cannot have versioning enabled or suspended, and containers with versioning enabled or suspended cannot have expiry lifecycle actions enabled to them.
    - COS Object versioning and COS Bucket Protection `(WORM)` cannot be used together.
    - Containers with proxy configuration cannot use versioning and vice versa.
    - SoftLayer accounts cannot use versioning.
    - Currently, you cannot support `MFA_Delete`, that is a feature to add additional security to version delete.
- `resource_instance_id` - (Required, String) The ID of the IBM Cloud Object Storage service instance for which you want to create a bucket.
- `region_location` - (Optional, String) The location of a regional bucket. Supported values are `au-syd`, `eu-de`, `eu-gb`, `jp-tok`, `us-east`, `us-south`. If you set this parameter, do not set `single_site_location` or `cross_region_location` at the same time.
- `retention_rule` - (List) Nested block have the following structure:
  
  Nested scheme for `retention rule`:
  - `default` - (Required, Integer) default retention period are defined by this policy and apply to all objects in the bucket.
  - `maximum` - (Required, Integer) Specifies maximum duration of time an object that can be kept unmodified in the bucket.
  - `minimum` - (Required, Integer) Specifies minimum duration of time an object must be kept unmodified in the bucket.
  - `permanent` : (Optional, Bool) Specifies a permanent retention status either enable or disable for a bucket.

    **Note**
     - Retention policies cannot be removed. For a new bucket, ensure that you are creating the bucket in a supported region. For more information, see [Integrated Services](https://cloud.ibm.com/docs/cloud-object-storage/basics?topic=cloud-object-storage-service-availability).
     - The minimum retention period must be less than or equal to the default retention period, that in turn must be less than or equal to the maximum retention period.
     - Permanent retention can only be enabled at a IBM Cloud Object Storage bucket level with retention policy enabled and users are able to select the permanent retention period option during object uploads. Once enabled, this process can't be reversed and objects uploaded that use a permanent retention period cannot be deleted. It's the responsibility of the users to validate at their end if there's a legitimate need to permanently store objects by using Object Storage buckets with a retention policy.
     - force deleting the bucket will not work if any object is still under retention. As objects cannot be deleted or overwritten until the retention period has expired and all the legal holds have been removed.
- `single_site_location` - (Optional, String) The location for a single site bucket. Supported values are: `ams03`, `che01`, `hkg02`, `mel01`, `mex01`, `mil01`, `mon01`, `osl01`, `par01`, `sjc04`, `sao01`, `seo01`, `sng01`, and `tor01`. If you set this parameter, do not set `region_location` or `cross_region_location` at the same time.
- `storage_class` - (Required, String) The storage class that you want to use for the bucket. Supported values are `standard`, `vault`, `cold`, `flex`, and `smart`. For more information, about storage classes, see [Use storage classes](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-classes).


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

## Import
The `ibm_cos_bucket` resource can be imported by using the `id`. The ID is formed from the `CRN` (Cloud Resource Name), the `bucket type` which must be `ssl` for single_site_location, `rl` for region_location or `crl` for cross_region_location, and the bucket location. The `CRN` and bucket location can be found on the portal.

id = `$CRN:meta:$buckettype:$bucketlocation`

**Syntax**

```
$ terraform import ibm_cos_bucket.mybucket <crn>

```

**Example**

```

$ terraform import ibm_cos_bucket.mybucket crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:mybucketname:meta:crl:eu:public

```
