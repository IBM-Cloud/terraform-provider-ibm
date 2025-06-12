---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM: ibm_cos_bucket_object"
description: |-
  Get information about an object in an IBM Cloud Object Storage bucket.
---

# ibm_cos_bucket_object

Retrieves information of an object in IBM Cloud Object Storage bucket. For more information, about an IBM Cloud Object Storage bucket, see [Create some buckets to store your data](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-getting-started-cloud-object-storage#gs-create-buckets). 

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

data "ibm_cos_bucket" "cos_bucket" {
  resource_instance_id = data.ibm_resource_instance.cos_instance.id
  bucket_name          = "my-bucket"
  bucket_type          = "region_location"
  bucket_region        = "us-east"
}

data "ibm_cos_bucket_object" "cos_object" {
  bucket_crn      = data.ibm_cos_bucket.cos_bucket.crn
  bucket_location = data.ibm_cos_bucket.cos_bucket.bucket_region
  key             = "object.json"
}
```
## Argument reference
Review the argument references that you can specify for your data source. 

- `bucket_crn` - (Required, String) The CRN of the COS bucket.
- `bucket_location` - (Required, String) The location of the COS bucket.
- `endpoint_type` - (Optional, String) The type of endpoint used to access COS. Accepted values: `public`, `private`, or `direct`. Default value is `public`.
- `key` - (Required, String) The name of an object in the COS bucket.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of an object.
- `body` - (String) Literal string value of an object content. Only supported for `text/*` and `application/json` content types.
- `content_length` - (String) A standard MIME type describing the format of an object data.
- `content_type` - (String) A standard MIME type describing the format of an object data.
- `etag` - (String) Computed MD5 hexdigest of an object content.
- `last_modified` - (Timestamp) Last modified date of an object in a GMT formatted date.
- `object_sql_url` - (String) Access the object using an SQL Query instance. The SQL URL is a reference URL used inside an SQL statement. The reference URL is used to perform queries against objects storing structured data.
- `object_lock_mode` - (String) This is the retention mode for an object.
- `object_lock_retain_until_date` - (String) A date after which the object can be deleted from the COS bucket.
- `object_lock_legal_hold_status` - (String) If the value of this attribute is **ON**, then the object cannot be deleted from the COS bucket.
- - `website_redirect` - (String) If this value is set then incoming request will be redirected to the set value..
