# IBM Cloud Object Storage example

The following example creates an instance of IBM Cloud Object Storage, IBM Cloud Activity Tracker, and IBM Cloud Monitoring with Sysdig. Then, multiple buckets are created and configured to send audit events and metrics to your service instances.

Following types of resources are supported:

* [Cloud Object Storage resource](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-index-of-terraform-on-ibm-cloud-resources-and-data-sources#ibm-object-storage_rd)

## Usage

To run this example you need to execute:

```sh
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Create an IBM Cloud Object Storage bucket. The bucket is used to store your data:

```terraform

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
  bucket_name          = var.bucket_name
  resource_instance_id = ibm_resource_instance.cos_instance.id
  single_site_location = "sjc04"
  #cross_region_location = var.region
  storage_class        = var.storage
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
  allowed_ip =  ["223.196.168.27","223.196.161.38","192.168.0.1"]
}

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

resource "ibm_cos_bucket" "retention_cos" {
  bucket_name          = "a-bucket-retention"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "jp-tok"
  storage_class        = standard
  force_delete        = true
  hard_quota          = 11
  retention_rule {
    default = 1
    maximum = 1
    minimum = 1
    permanent = false
  }
}

resource "ibm_cos_bucket" "objectversioning" {
  bucket_name           = "a-bucket-versioning"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = "us-east"
  storage_class         = var.storage
  hard_quota            = 1024
  object_versioning {
    enable  = true
  }
}

```

```terraform
data "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
}

data "ibm_cos_bucket" "standard-ams03" {
  bucket_name = ibm_cos_bucket.standard-ams03-firewall.bucket_name
  resource_instance_id = data.ibm_resource_instance.cos_instance.id
  bucket_type = "region_location"
  bucket_region = "us-south"
}
```

## Examples

* [Cloud Object Storage](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-cos-bucket)

<!-- COS SATELLITE PROJECT -->

## COS Satellite

The following example creates a bucket and adds object versioning and expiration features on COS Satellite location. As of now we are using existing COS instance to create bucket, so no need to create any COS instance via terraform. We do not have any resource group in Satellite. We can not use storage_class with Satellite location id.

* [IBM Satellite](https://cloud.ibm.com/docs/satellite?topic=satellite-getting-started)
* [IBM COS Satellite](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-about-cos-satellite)

## Example Usage

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

## COS Replication

Replication allows users to define rules for automatic, asynchronous copying of objects from a source bucket to a destination bucket in the same or different location.

**Note:**

You must have writer or manager platform roles on source bucket and sufficient platform roles to create new [IAM policies](https://cloud.ibm.com/docs/account?topic=account-iamoverview#iamoverview) that allow the source bucket to write to the destination bucket.

Add depends_on on ibm_iam_authorization_policy.policy in template to make sure replication is only enabled once iam authorization policy is set.

## Example usage
The following example creates an instance of IBM Cloud Object Storage. Then, multiple buckets are created and configured with replication policy.

```terraform
data "ibm_resource_group" "cos_group" {
  name = "cos-resource-group"
}

resource "ibm_resource_instance" "cos_instance_source" {
  name              = "cos-instance-src"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_resource_instance" "cos_instance_destination" {
  name              = "cos-instance-dest"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "cos_bucket_source" {
  bucket_name           = "a-bucket-source"
  resource_instance_id = ibm_resource_instance.cos_instance_source.id
  region_location      = "us-south"
  storage_class         = "standard"
  object_versioning {
    enable  = true
  }
}

resource "ibm_cos_bucket" "cos_bucket_destination" {
  bucket_name           = "a-bucket-destination"
  resource_instance_id = ibm_resource_instance.cos_instance_destination.id
  region_location      = "us-south"
  storage_class         = "standard"
  object_versioning {
    enable  = true
  }
}


### Configure IAM authorization policy

resource "ibm_iam_authorization_policy" "policy" {
  roles                  = [
      "Writer",
  ]
  subject_attributes {
    name  = "accountId"
    value = "an-account-id"
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
    value    = "an-account-id"
  }
  resource_attributes {
    name     = "serviceName"
    value    = "cloud-object-storage"
  }
  resource_attributes { 
    name  =  "serviceInstance"
    value =  ibm_resource_instance.cos_instance_destination.guid
  }
  resource_attributes { 
    name  =  "resource"
    value =   ibm_cos_bucket.cos_bucket_destination.bucket_name
  }
  resource_attributes { 
    name  =  "resourceType"
    value =  "bucket" 
  }
}

### Configure replication policy

resource "ibm_cos_bucket_replication_rule" "cos_bucket_repl" {
  depends_on = [
      ibm_iam_authorization_policy.policy
  ]
  bucket_crn	    = ibm_cos_bucket.cos_bucket_source.crn
  bucket_location = ibm_cos_bucket.cos_bucket_source.region_location
  replication_rule {
    rule_id = "a-rule-id"
    enable = "true"
    prefix = "a-prefix"
    priority = "a-priority-associated-with-the-rule"
    deletemarker_replication_status = "Enabled/Suspened"
    destination_bucket_crn = ibm_cos_bucket.cos_bucket_destination.crn
  }
}

```
## COS Object Lock

Object Lock preserves electronic records and maintains data integrity by ensuring that individual object versions are stored in a WORM (Write-Once-Read-Many), non-erasable and non-rewritable manner. This policy is enforced until a specified date or the removal of any legal holds.

## Example usage
The following example creates an instance of IBM Cloud Object Storage, creates a bucket with Object Lock enabled, and then sets Object Lock configuration on the bucket.

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

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = "a-bucket"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class         = "standard"
  object_versioning {
    enable  = true
  }
  object_lock = true
}

resource ibm_cos_bucket_object_lock_configuration "objectlock" {
 bucket_crn      = ibm_cos_bucket.cos_bucket.crn
 bucket_location = ibm_cos_bucket.cos_bucket.region_location
 object_lock_configuration{
   object_lock_enabled = "Enabled"
   object_lock_rule{
     default_retention{
        mode = "COMPLIANCE"
        days = 4
      }
    }
  }
}
```


## COS Static Webhosting

Provides an  Static web hosting configuration resource. This resource is used to  configure the website to use your documents as an index for the site and to potentially display errors.It can also be used to configure more advanced options including routing rules and request redirect for your domain.

## Example usage
The following example creates an instance of IBM Cloud Object Storage, creates a bucket and adds a website configuration on the bucket.Along with the basic bucket configuration , example of redirect all requests and adding routing rules have been given below.

```terraform

# Create a bucket
resource "ibm_cos_bucket" "cos_bucket_website_configuration" {
  bucket_name           = var.bucket_name
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = var.regional_loc
  storage_class         = var.standard_storage_class

}
# Give public access to above mentioned bucket
resource "ibm_iam_access_group_policy" "policy" { 
  depends_on = [ibm_cos_bucket.cos_bucket_website_configuration] 
  access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
  roles = ["Object Reader"] 

  resources { 
    service = "cloud-object-storage" 
    resource_type = "bucket" 
    resource_instance_id = "COS instance guid" 
    resource = data.ibm_cos_bucket.cos_bucket_website_configuration.bucket_name 
  } 
} 

# Add basic website configuration on a COS bucket
resource ibm_cos_bucket_website_configuration "website_configuration" {
  bucket_crn = "bucket_crn"
  bucket_location = data.ibm_cos_bucket.cos_bucket_website_configuration.regional_location
  website_configuration {
    error_document{
      key = "error.html"
    }
    index_document{
      suffix = "index.html"
    }
  }
}

# Add a request redirect website configuration on a COS bucket
resource ibm_cos_bucket_website_configuration "website_configuration" {
  bucket_crn = "bucket_crn"
  bucket_location = data.ibm_cos_bucket.cos_bucket_website_configuration.regional_location
  website_configuration {
    redirect_all_requests_to{
			host_name = "exampleBucketName"
			protocol = "https"
		}
  }
}


# Add a website configuration on a COS bucket with routing rule

resource ibm_cos_bucket_website_configuration "website_configuration" {
  bucket_crn = "bucket_crn"
  bucket_location = data.ibm_cos_bucket.cos_bucket_website_configuration.regional_location
  website_configuration {
    error_document{
      key = "error.html"
    }
    index_document{
      suffix = "index.html"
    }
    routing_rule {
      condition {
        key_prefix_equals = "pages/"
      }
      redirect {
        replace_key_prefix_with = "web_pages/"
      }
    }
  }
}

# Add a website configuration on a COS bucket with JSON routing rule
resource ibm_cos_bucket_website_configuration "website_configuration" {
  bucket_crn = "bucket_crn"
  bucket_location = data.ibm_cos_bucket.cos_bucket_website_configuration.regional_location
  website_configuration {
    error_document{
      key = "error.html"
      }
    index_document{
      suffix = "index.html"
    }
    routing_rules = <<EOF
			[{
			    "Condition": {
			        "KeyPrefixEquals": "pages/"
			     },
			     "Redirect": {
			        "ReplaceKeyPrefixWith": "webpages/"
			     }
			 }]
			 EOF
  }
}
```
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

| Name | Version |
|------|---------|
| terraform | >=1.0.0, <2.0 |

## Providers

| Name | Version |
|------|---------|
| ibm | Latest |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| bucket_name | Name of the bucket. | `string` | yes |
| resource_group_name | Name of the resource group. | `string` | yes |
| satellite_location_id | satellite location. | `string` | no |
| storage | The storage class that you want to use for the bucket. Supported values are **standard, vault, cold, flex, and smart**.| `string` | no |
| region | The location for a cross-regional bucket. Supported values are **us, eu, and ap**. | `string` | no |
| read_data_events | Enables sending log data to Activity Tracker and LogDNA to provide visibility into object read and write events. | `array` | no
| write_data_events | All object write events (i.e. uploads) will be sent to Activity Tracker. | `bool` | no
| activity_tracker_crn | Required the first time activity_tracking is configured. | `string` | yes
| usage_metrics_enabled | Specify **true or false** to set usage metrics (i.e. bytes_used). | `bool` | no
| request_metrics_enabled | Specify true or false to set cos request metrics (i.e. get, put, or post request). | `bool` | no
| metrics_monitoring_crn | Required the first time metrics_monitoring is configured. The instance of IBM Cloud Monitoring that will receive the bucket metrics. | `string` | yes
| regional_loc | The location for a regional bucket. Supported values are **au-syd, eu-de, eu-gb, jp-tok, us-east, or us-south**. | `string` | no
| type | Specifies the archive type to which you want the object to transition. Supported values are  **Glacier or Accelerated**. | `string` |yes
| rule_id | Unique identifier for the rule. | `string` | no
| days | Specifies the number of days when the specific expire rule action takes effect. | `int` | no
| date | After the specifies date , the current version of objects in your bucket expires. | `string` | no
| expired_object_delete_marker | Expired object delete markers can be automatically cleaned up to improve performance in bucket. This cannot be used alongside version expiration. | `bool` | no
| prefix | Specifies a prefix filter to apply to only a subset of objects with names that match the prefix. | `string` | no
| noncurrent_days | Configuration parameter in your policy that says how long to retain a non-current version before deleting it. | `int` | no
| days_after_initiation | Specifies the number of days that govern the automatic cancellation of part upload. Clean up incomplete multi-part uploads after a period of time. | `int` | no
| default | Specifies a default retention period to apply in all objects in the bucket. | `int` | yes
| maximum | Specifies maximum duration of time an object can be kept unmodified in the bucket. | `int` | yes
| minimum | Specifies minimum duration of time an object must be kept unmodified in the bucket. | `int` | yes
| permanent | Specifies a permanent retention status either enable or disable for a bucket. | `bool` | no
| enable | Specifies Versioning status either **enable or suspended** for an objects in the bucket. | `bool` | no
| hard_quota | sets a maximum amount of storage (in bytes) available for a bucket. | `int` | no
| object_lock | enables Object Lock on a bucket. | `bool` | no
| bucket\_crn | The CRN of the source COS bucket. | `string` | yes |
| bucket\_location | The location of the source COS bucket. | `string` | yes |
| destination_bucket_crn | The CRN of your destination bucket that you want to replicate to. | `String` | yes
| deletemarker_replication_status | Specifies whether Object storage replicates delete markers.  Specify true for Enabling it or false for Disabling it. | `String` | no
| status | Specifies whether the rule is enabled. Specify true for Enabling it or false for Disabling it. | `String` | yes
| rule_id | The rule id. | `String` | no
| priority | A priority is associated with each rule. The rule will be applied in a higher priority if there are multiple rules configured. The higher the number, the higher the priority | `String` | no
| prefix | An object key name prefix that identifies the subset of objects to which the rule applies. | `String` | no
| bucket_crn | The CRN of the COS bucket on which Object Lock is enabled or should be enabled. | `String` | yes
| bucket_location | Location of the COS bucket. | `String` | yes
| endpoint_type | Endpoint types of the COS bucket. | `String` | no
| object_lock_enabled | Enable Object Lock on an existing COS bucket. | `String` | yes
| mode | Retention mode for the Object Lock configuration. | `String` | yes
| years | Retention period in terms of years after which the object can be deleted. | `int` | no
| days | Retention period in terms of days after which the object can be deleted. | `int` | no
| key | Object key name to use when a 4XX class error occurs given as error document. | `String` | no
| suffix | The home or default page of the website when static web hosting configuration is added. | `String` | Yes
| hostname | Name of the host where requests are redirected. | `String` | Yes
| protocol | Protocol to use when redirecting requests. The default is the protocol that is used in the original request. | `String` | No
| http_error_code_returned_equals | HTTP error code when the redirect is applied. | `String` | No
| key_prefix_equals | Object key name prefix when the redirect is applied. | `String` | No
| host_name | Host name to use in the redirect request. | `String` | Yes
| protocol | Protocol to use when redirecting requests. | `String` | No
| http_redirect_code | HTTP redirect code to use on the response. | `String` | No
| replace_key_with | Specific object key to use in the redirect request. | `String` | No
| replace_key_prefix_with | Object key prefix to use in the redirect request. | `String` | No
{: caption="inputs"}
