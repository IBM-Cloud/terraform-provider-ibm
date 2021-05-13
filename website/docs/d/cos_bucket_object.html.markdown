---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM: ibm_cos_bucket_object"
description: |-
  Get information about an object in an IBM Cloud Object Storage bucket.
---

# ibm\_cos_bucket_object

Retrieves information of an object in IBM Cloud Object Storage bucket.

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

## Argument Reference

The following arguments are supported:

* `bucket_crn` - (Required, string) The CRN of the COS bucket.
* `bucket_location` - (Required, string) The location of the COS bucket.
* `endpoint_type` - (Optional, string) The type of endpoint used to access COS. Accepted values: `public`, `private`, or `direct`. Default value is `public`.
* `key` - (Required, string) The name of the object in the COS bucket.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the object.
* `body` - Literal string value of the object content. Only supported for `text/*` and `application/json` content types.
* `content_length` - A standard MIME type describing the format of the object data.
* `content_type` - A standard MIME type describing the format of the object data.
* `etag` - Computed MD5 hexdigest of the object content.
* `last_modified` - Last modified date of the object. A GMT formatted date.
