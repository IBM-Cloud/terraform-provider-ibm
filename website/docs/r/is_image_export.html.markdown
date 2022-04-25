---
layout: "ibm"
page_title: "IBM : ibm_is_image_export"
description: |-
  Manages ImageExportJob.
subcategory: "VPC infrastructure"
---

# ibm_is_image_export

Provides a resource for ImageExportJob. This allows ImageExportJob to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_is_image_export" "example" {
  image = "d7bec597-4726-451f-8a63-e62e6f121c32c"
  name = "my-image-export"
  storage_bucket_name = "bucket-27200-lwx4cfvcue"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `format` - (Optional, Forces new resource, String) The format to use for the exported image. If the image is encrypted, only `qcow2` is supported. Allowable values are: `qcow2`, `vhd`.
- `image` - (Required, Forces new resource, String) The image identifier.
- `name` - (Optional, String) The user-defined name for this image export job. Names must be unique within the image this export job resides in. If unspecified, the name will be a hyphenated list of randomly-selected words prefixed with the first 16 characters of the parent image name.The exported image object name in Cloud Object Storage (`storage_object.name` in the response) will be based on this name. The object name will be unique within the bucket.
- `storage_bucket_name` - (Optional, Forces new resource, String) The name of the Cloud Object Storage bucket to export the image to. The bucket must exist and an IAMservice authorization must grant `Image Service for VPC` of`VPC Infrastructure Services` writer access to the bucket.
- `storage_bucket_crn` - (Optional, Forces new resource, String) The CRN of the Cloud Object Storage bucket to export the image to. The bucket must exist and an IAMservice authorization must grant `Image Service for VPC` of`VPC Infrastructure Services` writer access to the bucket.
  
  -> **NOTE:**
  `storage_bucket_name` and `storage_bucket_crn` are mutually exclusive. Provide either one of them.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the ImageExportJob. Follows the format <image_id>/<image_export_job_id>.
- `completed_at` - (String) The date and time that the image export job was completed.If absent, the export job has not yet completed.
- `created_at` - (String) The date and time that the image export job was created.
- `encrypted_data_key` - (Optional, String) A base64-encoded, encrypted representation of the key that was used to encrypt the data for the exported image. This key can be unwrapped with the image's `encryption_key` root key using either Key Protect or Hyper Protect Crypto Service.If absent, the export job is for an unencrypted image.
- `href` - (Required, String) The URL for this image export job.
- `image_export_job` - (Required, String) The unique identifier for this image export job.
- `resource_type` - (Required, String) The type of resource referenced.
- `started_at` - (Optional, String) The date and time that the image export job started running.If absent, the export job has not yet started.
- `status` - (Required, String) The status of this image export job:- `deleting`: Export job is being deleted- `failed`: Export job could not be completed successfully- `queued`: Export job is queued- `running`: Export job is in progress- `succeeded`: Export job was completed successfullyThe exported image object is automatically deleted for `failed` jobs. Allowable values are: `deleting`, `failed`, `queued`, `running`, `succeeded`.
- `status_reasons` - (Required, List) The reasons for the current status (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
Nested scheme for **status_reasons**:
  - `code` - (Required, String) A snake case string succinctly identifying the status reason.
  - `message` - (Required, String) An explanation of the status reason.
  - `more_info` - (Optional, String) Link to documentation about this status reason.
- `storage_href` - (Required, String) The Cloud Object Storage location of the exported image object. The object at this location may not exist until the job is started, and will be incomplete while the job is running.After the job completes, the exported image object is not managed by the IBM VPC service, and may be removed or replaced with a different object by any user or service with IAM authorization to the bucket.
- `storage_object` - (Required, List) The Cloud Object Storage object for the exported image. This object may not exist untilthe job is started, and will not be complete until the job completes.
Nested scheme for **storage_object**:
  - `name` - (Required, String) The name of this Cloud Object Storage object. Names are unique within a Cloud Object Storage bucket.


## Import

You can import the `ibm_is_image_export` resource by using `id`.
The `id` property can be formed from `image_id`, and `id` in the following format:

```
<image_id>/<id>
```
- `image_id`: A string. The image identifier.
- `id`: A string. The image export job identifier.

# Syntax
```
$ terraform import ibm_is_image_export.is_image_export <image_id>/<id>
```
