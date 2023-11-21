---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM: ibm_cos_bucket_object"
description: |-
  Manages an object in an IBM Cloud Object Storage bucket.
---

# ibm_cos_bucket_object

Create, update, or delete an object in an IBM Cloud Object Storage bucket. For more information, about an IBM Cloud Object Storage bucket, see [Create some buckets to store your data](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-getting-started-cloud-object-storage#gs-create-buckets). 

## Example usage

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
# Object Lock

Object Lock preserves electronic records and maintains data integrity by ensuring that individual object versions are stored in a WORM (Write-Once-Read-Many), non-erasable and non-rewritable manner. This policy is enforced until a specified date or the removal of any legal holds.

**Note:**
Object Lock must be enabled on the bucket to configure `object_lock_mode` , `object_lock_retain_until_date` , `object_lock_legal_hold_status` on the object.

## Example usage

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
  bucket_name           = "my-bucket"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  region_location       = "us-east"
  storage_class         = "standard"
  object_versioning {
    enable  = true
  }
  object_lock = true
}

resource "ibm_cos_bucket_object" "cos_object_objectlock" {
  bucket_crn      = data.ibm_cos_bucket.cos_bucket.crn
  bucket_location = data.ibm_cos_bucket.cos_bucket.bucket_region
  key             = "object.json"
  object_lock_mode              = "COMPLIANCE"
  object_lock_retain_until_date = "2023-02-15T18:00:00Z"
  object_lock_legal_hold_status = "ON"
}
```

# Website redirect 

Requests can be redirected for an object to another object or URL by setting the website redirect location in the metadata of the object.

**Note:**
COS bucket with website configuration and public access enabled is a pre-requisite.For adding website configuration to a bucket please follow [static_web_hosting]()

## Example usage

```terraform

resource "ibm_cos_bucket_object" "cos_object_objectlock" {
  bucket_crn      = "bucket-crn"
  bucket_location = "us-south"
  key             = "page1.html"
  website_redirect = "/page2.html"
}
```


## Argument reference
Review the argument references that you can specify for your resource.

- `bucket_crn` - (Required, Forces new resource, String) The CRN of the COS bucket.
- `bucket_location` - (Required, Forces new resource, String) The location of the COS bucket.
- `content` - (Optional, String) Literal string value to use as an object content, which will be uploaded as UTF-8 encoded text. Conflicts with `content_base64` and `content_file`.
- `content_base64` - (Optional, String) Base64-encoded data that will be decoded and uploaded as raw bytes for an object content. This safely uploads `non-UTF8` binary data, but is recommended only for small content. Conflicts with `content` and `content_file`.
- `content_file` - (Optional, String) The path to a file that will be read and uploaded as raw bytes for an object content. Conflicts with `content` and `content_base64`.
- `endpoint_type` - (Optional, String) The type of endpoint used to access COS. Supported values are `public`, `private`, or `direct`. Default value is `public`.
- `etag` - (Optional, String) MD5 hexdigest used to trigger updates. The only meaningful value is `filemd5("path/to/file")`.
- `key` - (Required, Forces new resource, String) The name of an object in the COS bucket.
- `website_redirect` - (Optional, String) Target URL for website redirect.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of an object.
- `body` - (String) Literal string value of an object content. Only supported for `text/*` and `application/json` content types.
- `content_length` - (String) A standard MIME type describing the format of an object data.
- `content_type` - (String) A standard MIME type describing the format of an object data.
- `etag` - (String) Computed MD5 hexdigest of an object content.
- `last_modified` - (Timestamp) Last modified date of an object. A GMT formatted date.
- `object_sql_url` - (String) Access the object using an SQL Query instance. The SQL URL is a reference url used inside of an SQL statement. The reference url is used to perform queries against objects storing structured data.

## Import

The `ibm_cos_bucket_object` resource can be imported by using the `id`. The ID is formed from the COS bucket CRN, an object key name, and the bucket location.

id = ${bucketCRN}:object:${objectKey}:location:${bucketLocation}

**Syntax**

```
$ terraform import ibm_cos_bucket_object.my_object <id>
```

**Example**

```
$ terraform import ibm_cos_bucket.my_object crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:myBucketName:object:myObject.key:location:us-east
```
