---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Object Lock Configuration"
description: 
  "Manages IBM Cloud Object Storage Object Lock Configuration"
---

# ibm_cos_bucket_object_lock_configuration
Provides an  Object Lock configuration resource. This resource is used to configure a default retention period for objects placed in the specified bucket. To enable Object Lock for a new bucket see [ibm_cos_bucket](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/cos_bucket). To configure legal hold and retention period on an object please refer [ibm_cos_bucket_object](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/cm_object).



**Note:**
To configure Object Lock on a bucket, you must  first enable object versioning on bucket by using the [Versioning objects](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-versioning).

---

## Example usage
The following example demonstrates creating a bucket with object lock enabled with default retention.

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
  bucket_name          = "a-standard-bucket"
  resource_instance_id = data.ibm_resource_instance.cos_instance.id
  bucket_region        = "us-south"
  storage_class        = "Standard"
  object_versioning {
    enable  = true
  }
  object_lock = true
}

resource ibm_cos_bucket_object_lock_configuration "objectlock" {
 bucket_crn      = ibm_cos_bucket.cos_bucket.crn
 bucket_location = ibm_cos_bucket.cos_bucket.bucket_region
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
# Enabling Object Lock configuration on an existing bucket
To enable  Object Lock configuration on an existing bucket, create a COS bucket with object versioning enabled and pass the crn of the COS bucket and location of the bucket to `ibm_cos_bucket_object_lock_configuration.bucket_crn` and `ibm_cos_bucket_object_lock_configuration.bucket_location` as shown in the example.

## Example usage

```terraform
// To only enable Object Lock configuration on an existing bucket

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name          = "a-standard-bucket"
  resource_instance_id = data.ibm_resource_instance.cos_instance.id
  bucket_region        = "us-south"
  storage_class        = "Standard"
  object_versioning {
    enable  = true
  }
}

resource ibm_cos_bucket_object_lock_configuration "objectlock" {
 bucket_crn      = ibm_cos_bucket.cos_bucket.crn
 bucket_location = ibm_cos_bucket.cos_bucket.bucket_region
 object_lock_configuration{
   object_lock_enabled = "Enabled"
  }
}

// To enable object lock configuration and set default retention on a bucket

resource ibm_cos_bucket_object_lock_configuration "objectlock" {
 bucket_crn      = ibm_cos_bucket.cos_bucket.crn
 bucket_location = ibm_cos_bucket.cos_bucket.bucket_region
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

## Argument reference
Review the argument references that you can specify for your resource. 
- `bucket_crn` - (Required, Forces new resource, String) The CRN of the COS bucket.
- `bucket_location` - (Required, Forces new resource, String) The location of the COS bucket.
- `endpoint_type`- (Optional, String) The type of the endpoint either `public` or `private` or `direct` to be used for buckets. Default value is `public`.
- `object_lock_configuration`- (Required, List) Nested block have the following structure:
  
  Nested scheme for `object_lock_configuration`:
  - `object_lock_enabled`- (String) Indicates whether this bucket has an Object Lock configuration enabled. Defaults to Enabled. Valid values: Enabled.
  - `object_lock_rule`- (List) Object Lock rule has following argument:
  
  Nested scheme for `object_lock_rule`:
  - `default_retention`- (Required) Configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket
  Nested scheme for `default_retention`:
  - `mode`- (String)  Default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Supported values: COMPLIANCE.
  - `days`- (Int) Specifies number of days after which the object can be deleted from the COS bucket.
  - `years`- (Int) Specifies number of years after which the object can be deleted from the COS bucket.

**Note:**
  The parameter `days` and `years` are mutually exclusive please provide only one of them.
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the bucket.
- `id` - (String) The ID of the bucket.

## Import IBM COS Bucket
The `ibm_cos_bucket_object_lock_configuration` resource can be imported by using the `id`. The ID is formed from the `CRN` (Cloud Resource Name). The `CRN` and bucket location can be found on the portal.

id = `$CRN:meta:$bucketlocation:$endpointtype`

**Syntax**

```
$ terraform import ibm_cos_bucket_object_lock_configuration.objectlock `$CRN:meta:$bucketlocation:public`

```

**Example**

```

$ terraform import ibm_cos_bucket_object_lock_configuration.objectlock crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:mybucketname:meta:us-south:public

```