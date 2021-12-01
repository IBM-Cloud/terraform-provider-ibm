---
layout: "ibm"
page_title: "IBM : ibm_atracker_targets"
description: |-
  Get information about atracker_targets
subcategory: "Activity Tracker"
---

# ibm_atracker_targets

Provides a read-only data source for atracker_targets. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_atracker_targets" "atracker_targets" {
	name = "a-cos-target-us-south"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `name` - (Optional, String) The name of the target resource.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the atracker_targets.
* `targets` - (Required, List) A list of target resources.
Nested scheme for **targets**:
	* `id` - (Required, String) The uuid of the target resource.
	* `name` - (Required, String) The name of the target resource.
	* `crn` - (Required, String) The crn of the target resource.
	* `target_type` - (Required, String) The type of the target.
	  * Constraints: Allowable values are: cloud_object_storage
	* `encrypt_key` - (Optional, String) The encryption key that is used to encrypt events before Activity Tracker services buffer them on storage. This credential is masked in the response.
	* `cos_endpoint` - (Optional, List) Property values for a Cloud Object Storage Endpoint.
	Nested scheme for **cos_endpoint**:
		* `endpoint` - (Required, String) The host name of the Cloud Object Storage endpoint.
		* `target_crn` - (Required, String) The CRN of the Cloud Object Storage instance.
		* `bucket` - (Required, String) The bucket name under the Cloud Object Storage instance.
		* `api_key` - (Required, String) The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response.
	* `cos_write_status` - (Optional, List) The status of the write attempt with the provided cos_endpoint parameters.
	Nested scheme for **cos_write_status**:
		* `status` - (Optional, String) The status such as failed or success.
		* `last_failure` - (Optional, String) The timestamp of the failure.
		* `reason_for_last_failure` - (Optional, String) Detailed description of the cause of the failure.
	* `created` - (Optional, String) The timestamp of the target creation time.
	* `updated` - (Optional, String) The timestamp of the target last updated time.

