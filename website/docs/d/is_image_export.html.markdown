---
layout: "ibm"
page_title: "IBM : ibm_is_image_export_job"
description: |-
  Get information about ImageExportJob
subcategory: "VPC infrastructure"
---

# ibm_is_image_export_job

Provides a read-only data source for ImageExportJob. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information about VPC custom images export, see [IBM Cloud Docs: Virtual Private Cloud - Exporting a custom image to IBM Cloud Object Storage](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-custom-images&interface=ui#custom-image-export-to-cos).

## Example Usage

```hcl
resource "ibm_is_image" "example" {
  name               = "example-image"
  href               = "cos://us-south/custom-image-vpc-bucket/customImage-0.vhd"
  operating_system   = "ubuntu-16-04-amd64"
  encrypted_data_key = "eJxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0="
  encryption_key     = "crn:v1:bluemix:public:kms:us-south:a/6xxxxxxxxxxxxxxx:xxxxxxx-xxxx-xxxx-xxxxxxx:key:dxxxxxx-fxxx-4xxx-9xxx-7xxxxxxxx"
}
resource "ibm_is_image_export_job" "example" {
  image          = ibm_is_image.example.id
  name           = "my-image-export"
  storage_bucket {
    name = "bucket-27200-lwx4cfvcue"
  }
}
data "ibm_is_image_export_job" "example" {
  image_export_job = ibm_is_image_export_job.example.image_export_job
  image            = ibm_is_image_export_job.is_image_export.image
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `image_export_job` - (Required, String) The image export job identifier.
- `image` - (Required, String) The image identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the ImageExportJob.
- `completed_at` - (Optional, String) The date and time that the image export job was completed.If absent, the export job has not yet completed.
- `created_at` - (String) The date and time that the image export job was created.
- `encrypted_data_key` - (Optional, String) A base64-encoded, encrypted representation of the key that was used to encrypt the data for the exported image. This key can be unwrapped with the image's `encryption_key` root key using either Key Protect or Hyper Protect Crypto Service.If absent, the export job is for an unencrypted image.
- `format` - (String) The format of the exported image. Allowable values are: `qcow2`, `vhd`.
- `href` - (String) The URL for this image export job.
- `image_export_job` - (String) The unique identifier for this image export job.
- `name` - (String) The user-defined name for this image export job.
- `resource_type` - (String) The type of resource referenced.
- `started_at` - (Optional, String) The date and time that the image export job started running.If absent, the export job has not yet started.
- `status` - (String) The status of this image export job:- `deleting`: Export job is being deleted- `failed`: Export job could not be completed successfully- `queued`: Export job is queued- `running`: Export job is in progress- `succeeded`: Export job was completed successfullyThe exported image object is automatically deleted for `failed` jobs. Allowable values are: `deleting`, `failed`, `queued`, `running`, `succeeded`.
- `status_reasons` - (List) The reasons for the current status (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
Nested scheme for **status_reasons**:
  - `code` - (String) A snake case string succinctly identifying the status reason.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (Optional, String) Link to documentation about this status reason.
- `storage_bucket` - (List) The Cloud Object Storage bucket of the exported image object.
Nested scheme for **storage_bucket**:
  - `crn` - (String) The CRN of this Cloud Object Storage bucket.
  - `name` - (String) The globally unique name of this Cloud Object Storage bucket.
- `storage_href` - (String) The Cloud Object Storage location of the exported image object. The object at this location may not exist until the job is started, and will be incomplete while the job is running.After the job completes, the exported image object is not managed by the IBM VPC service, and may be removed or replaced with a different object by any user or service with IAM authorization to the bucket.
- `storage_object` - (List) The Cloud Object Storage object for the exported image. This object may not exist untilthe job is started, and will not be complete until the job completes.
Nested scheme for **storage_object**:
  - `name` - (String) The name of this Cloud Object Storage object. Names are unique within a Cloud Object Storage bucket.

