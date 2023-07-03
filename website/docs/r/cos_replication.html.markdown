---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket Replication"
description: 
  "Manages IBM Cloud Object Storage Bucket Replication."
---

# ibm_cos_bucket_replication_rule
Create/replaces or delete replication configuration on an existing bucket. For more information, about configuration options, see [Replicating objects](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-replication-overview). 

To configure a replication policy on a bucket, you must enable object versioning on both source and destination buckets by using the [Versioning objects](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-versioning).

**Note:**

 you must have `writer` or `manager` platform roles on source bucket and sufficient platform roles to create new [IAM policies](https://cloud.ibm.com/docs/account?topic=account-iamoverview#iamoverview) that allow the source bucket to write to the destination bucket.
 Add depends_on on ibm_iam_authorization_policy.policy in template to make sure replication only enabled once iam  authorization policy set. 
 We are addressing Create functionality for replication now. Update functionality will be in-progress.


---

## Example usage
The following example creates an instance of IBM Cloud Object Storage. Then, multiple buckets are created and configured replication policy.

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

## Argument reference
Review the argument references that you can specify for your resource. 
- `bucket_crn` - (Required, Forces new resource, String) The CRN of the COS bucket.
- `bucket_location` - (Required, Forces new resource, String) The location of the COS bucket.
- `endpoint_type`- (Optional, String) The type of the endpoint either `public` or `private` or `direct` to be used for buckets. Default value is `public`.
- `replication_rule`- (Required, List) Nested block have the following structure:

  Nested scheme for `replication_rule`:
  - `rule_id`- (Optional, String) The rule id.
  - `enable`-  (Required, Bool) Specifies whether the rule is enabled. Specify true for Enabling it  or false for Disabling it.
  - `prefix`- (Optional, String) An object key name prefix that identifies the subset of objects to which the rule applies.
  - `priority`- (Optional, Int) A priority is associated with each rule. The rule will be applied in a higher priority if there are multiple rules configured. The higher the number, the higher the priority
  - `deletemarker_replication_status`-  (Optional, Bool) Specifies whether Object storage replicates delete markers.Specify true for Enabling it  or false for Disabling it.
  - `destination_bucket_crn`-  (Required, String) The CRN of your destination bucket that you want to replicate to.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the bucket.
- `id` - (String) The ID of the bucket.

## Import IBM COS Bucket
The `ibm_cos_bucket_replication_rule` resource can be imported by using the `id`. The ID is formed from the `CRN` (Cloud Resource Name). The `CRN` and bucket location can be found on the portal.

id = `$CRN:meta:$bucketlocation:$endpointtype`

**Syntax**

```
$ terraform import ibm_cos_bucket_replication_rule.mybucket `$CRN:meta:$bucketlocation:public`

```

**Example**

```

$ terraform import ibm_cos_bucket_replication_rule.mybucket crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:mybucketname:meta:us-south:public

```