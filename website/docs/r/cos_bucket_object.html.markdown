---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM: ibm_cos_bucket_object"
description: |-
  Manages an object in an IBM Cloud Object Storage bucket.
---

# ibm\_cos_bucket_object

Create, update, or delete an object in an IBM Cloud Object Storage bucket.

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

resource "ibm_cos_bucket" "cos_bucket" {
  bucket_name           = "my-bucket"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = "us-east"
  storage_class         = "standard"
}

resource "ibm_cos_bucket_object" "plaintext" {
  bucket_crn      = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  content         = "Hello World"
  key             = "plaintext.txt"
}

resource "ibm_cos_bucket_object" "base64" {
  bucket_crn      = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  content_base64  = "RW5jb2RlZCBpbiBiYXNlNjQ="
  key             = "base64.txt"
}

resource "ibm_cos_bucket_object" "file" {
  bucket_crn      = ibm_cos_bucket.cos_bucket.crn
  bucket_location = ibm_cos_bucket.cos_bucket.region_location
  content_file    = "${path.module}/object.json"
  key             = "file.json"
  etag            = filemd5("${path.module}/object.json")
}
```

## Argument Reference

The following arguments are supported:

* `bucket_crn` - (Required, Forces new resource, string) The CRN of the COS bucket.
* `bucket_location` - (Required, Forces new resource, string) The location of the COS bucket.
* `content` - (Optional, string) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text. Conflicts with `content_base64` and `content_file`.
* `content_base64` - (Optional, string) Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content. Conflicts with `content` and `content_file`.
* `content_file` - (Optional, string) The path to a file that will be read and uploaded as raw bytes for the object content. Conflicts with `content` and `content_base64`.
* `endpoint_type` - (Optional, string) The type of endpoint used to access COS. Accepted values: `public`, `private`, or `direct`. Default value is `public`.
* `etag` - (Optional, string) MD5 hexdigest used to trigger updates. The only meaningful value is `filemd5("path/to/file")`.
* `key` - (Required, Forces new resource, string) The name of the object in the COS bucket.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the object.
* `body` - Literal string value of the object content. Only supported for `text/*` and `application/json` content types.
* `content_length` - A standard MIME type describing the format of the object data.
* `content_type` - A standard MIME type describing the format of the object data.
* `etag` - Computed MD5 hexdigest of the object content.
* `last_modified` - Last modified date of the object. A GMT formatted date.

## Import

The `ibm_cos_bucket_object` resource can be imported using the `id`. The ID is formed from the COS bucket CRN, the object key name, and the bucket location.

id = ${bucketCRN}:object:${objectKey}:location:${bucketLocation}

```
$ terraform import ibm_cos_bucket_object.my_object <id>

$ terraform import ibm_cos_bucket.my_object crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:myBucketName:object:myObject.key:location:us-east
```
