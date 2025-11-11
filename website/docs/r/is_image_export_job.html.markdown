---
layout: "ibm"
page_title: "IBM : ibm_is_image_export_job"
description: |-
  Manages ImageExportJob.
subcategory: "VPC infrastructure"
---

# ibm_is_image_export_job

Provides a resource for ImageExportJob. This allows ImageExportJob to be created, updated and deleted. For more information about VPC custom images export, see [IBM Cloud Docs: Virtual Private Cloud - Exporting a custom image to IBM Cloud Object Storage](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-custom-images&interface=ui#custom-image-export-to-cos).

~> **Note**
  Image export jobs are asynchronous. Time taken to export the image depends on its size. Hence the resource will not wait for job status to be completed. It is recommended to check the status of the export job by refreshing this resource or the datasources `ibm_is_image_export_job` and `ibm_is_image_export_jobs` and recreate the export resource if it is failed.

## Example Usage

```hcl
resource "ibm_is_image" "example" {
  name               = "example-image"
  href               = "cos://us-south/custom-image-vpc-bucket/customImage-0.vhd"
  operating_system   = "ubuntu-16-04-amd64"
  encrypted_data_key = "eJxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0="
  encryption_key     = "crn:v1:bluemix:public:kms:us-south:a/6xxxxxxxxxxxxxxx:xxxxxxx-xxxx-xxxx-xxxxxxx:key:dxxxxxx-fxxx-4xxx-9xxx-7xxxxxxxx"
}
// Create export job with storage bucket name
resource "ibm_is_image_export_job" "example" {
  image               = ibm_is_image.example.id
  name                = "my-image-export"
  storage_bucket {
    name = "bucket-27200-lwx4cfvcue"
  }
}
// Create export job with storage bucket CRN
resource "ibm_is_image_export_job" "example" {
  image               = ibm_is_image.example.id
  name                = "my-image-export"
  storage_bucket {
    crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/XXXXeaXXXX5XXXX0f0XXXX92ff85XXXX:aaXXXXX1-07XX-42XX-b8d0-aXXXXXX243:bucket:dallas-bucket"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `format` - (Optional, Forces new resource, String) The format to use for the exported image. If the image is encrypted, only `qcow2` is supported. Allowable values are: `qcow2`, `vhd`. Default value is `qcow2`.
- `image` - (Required, Forces new resource, String) The image identifier.
- `name` - (Optional, String) The user-defined name for this image export job. Names must be unique within the image this export job resides in. If unspecified, the name will be a hyphenated list of randomly-selected words prefixed with the first 16 characters of the parent image name.The exported image object name in Cloud Object Storage (`storage_object.name` in the response) will be based on this name. The object name will be unique within the bucket.
- `storage_bucket` - (Required, Forces new resource, List) The Cloud Object Storage bucket to export the image to. The bucket must exist and an IAM service authorization must grant Image Service for VPC of VPC Infrastructure Services writer access to the bucket.

  Nested scheme for `storage_bucket`:
  - `name` - (Optional, String) Name of this Cloud Object Storage bucket.
  - `crn` - (Optional, String) The CRN of this Cloud Object Storage bucket

  -> **NOTE:**
  Within `storage_bucket`, `name` and `crn` are mutually exclusive. Provide either one of them.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `completed_at` - (String) The date and time that the image export job was completed.If absent, the export job has not yet completed.
- `created_at` - (String) The date and time that the image export job was created.
- `encrypted_data_key` - (Optional, String) A base64-encoded, encrypted representation of the key that was used to encrypt the data for the exported image. This key can be unwrapped with the image's `encryption_key` root key using either Key Protect or Hyper Protect Crypto Service.If absent, the export job is for an unencrypted image.
- `href` - (String) The URL for this image export job.
- `id` - The unique identifier of the ImageExportJob. Follows the format <image_id>/<image_export_job_id>.
- `image_export_job` - (String) The unique identifier for this image export job.
- `resource_type` - (String) The type of resource referenced.
- `started_at` - (String) The date and time that the image export job started running.If absent, the export job has not yet started.
- `status` - (String) The status of this image export job:- `deleting`: Export job is being deleted- `failed`: Export job could not be completed successfully- `queued`: Export job is queued- `running`: Export job is in progress- `succeeded`: Export job was completed successfullyThe exported image object is automatically deleted for `failed` jobs. Allowable values are: `deleting`, `failed`, `queued`, `running`, `succeeded`.
- `status_reasons` - (List) The reasons for the current status (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
Nested scheme for **status_reasons**:
  - `code` - (String) A snake case string succinctly identifying the status reason.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (Optional, String) Link to documentation about this status reason.
- `storage_href` - (String) The Cloud Object Storage location of the exported image object. The object at this location may not exist until the job is started, and will be incomplete while the job is running.After the job completes, the exported image object is not managed by the IBM VPC service, and may be removed or replaced with a different object by any user or service with IAM authorization to the bucket.
- `storage_object` - (List) The Cloud Object Storage object for the exported image. This object may not exist untilthe job is started, and will not be complete until the job completes.
Nested scheme for **storage_object**:
  - `name` - (String) The name of this Cloud Object Storage object. Names are unique within a Cloud Object Storage bucket.


## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_image_export_job` resource by using `id`.
The `id` property can be formed from `image_id`, and `id`. For example:

```terraform
import {
  to = ibm_is_image_export_job.is_image_export
  id = "<image_id>/<id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_image_export_job.is_image_export <image_id>/<id>
```