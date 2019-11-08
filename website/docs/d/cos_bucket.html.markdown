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
  resource_group_id = "${data.ibm_resource_group.cos_group.id}"
  service           = "cloud-object-storage"
}

data "ibm_cos_bucket" "standard-ams03" {
  bucket_name = "a-standard-bucket-at-ams"
  resource_instance_id = "${data.ibm_resource_instance.cos_instance.id}"
  bucket_type = "single_site_location"
  region = "ams03"
}

output "bucket_private_endpoint" {
  value = "${data.ibm_cos_bucket.standard-ams03.s3_endpoint_private}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket_name` - (Required, string) The name of the bucket.
* `bucket_type` - (Required, string) The type of the bucket. Accepted values: single_site_location region_location cross_region_location
* `resource_instance_id` - (Required, string) The id of Cloud Object Storage instance.
* `bucket_region` - (Required, string) The region of the bucket.

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
