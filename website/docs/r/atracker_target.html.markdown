---
layout: "ibm"
page_title: "IBM : ibm_atracker_target"
description: |-
  Manages Activity Tracker Target.
subcategory: "Activity Tracker"
---

# ibm_atracker_target

Provides a resource for Activity Tracker Target. This allows Activity Tracker Target to be created, updated and deleted.

## Example usage

```terraform
resource "ibm_atracker_target" "atracker_target" {
  cos_endpoint { 
     endpoint = "endpoint" 
     target_crn = "target_crn" 
     bucket = "bucket" 
     api_key = "api_key" 
  }
  name = "my-cos-target"
  target_type = "cloud_object_storage"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `cos_endpoint` - (Required, List) Property values for a Cloud Object Storage Endpoint.
Nested scheme for **cos_endpoint**:
	* `endpoint` - (Required, String) The host name of the Cloud Object Storage endpoint.
	* `target_crn` - (Required, String) The CRN of the Cloud Object Storage instance.
	* `bucket` - (Required, String) The bucket name under the Cloud Object Storage instance.
	* `api_key` - (Required, String) The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response.
* `name` - (Required, String) The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`
* `target_type` - (Required, Forces new resource, String) The type of the target.
  * Constraints: Allowable values are: cloud_object_storage

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the Activity Tracker Target.
* `cos_write_status` - (Optional, List) The status of the write attempt with the provided cos_endpoint parameters.
Nested scheme for **cos_write_status**:
	* `status` - (Optional, String) The status such as failed or success.
	* `last_failure` - (Optional, String) The timestamp of the failure.
	* `reason_for_last_failure` - (Optional, String) Detailed description of the cause of the failure.
* `created` - (Optional, String) The timestamp of the target creation time.
* `crn` - (Required, String) The crn of the target resource.
* `encrypt_key` - (Optional, String) The encryption key that is used to encrypt events before Activity Tracker services buffer them on storage. This credential is masked in the response.
* `updated` - (Optional, String) The timestamp of the target last updated time.

## Import

You can import the `ibm_atracker_target` resource by using `id`. The uuid of the target resource.

# Syntax
```
$ terraform import ibm_atracker_target.atracker_target <id>
```

# Example
```
$ terraform import ibm_atracker_target.atracker_target f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6
```
