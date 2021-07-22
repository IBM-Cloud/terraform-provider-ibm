---
layout: "ibm"
page_title: "IBM : ibm_atracker_targets"
description: |-
  Get information about atracker_targets
subcategory: "Activity Tracking API"
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

The following arguments are supported:

* `name` - (Optional, string) The name of the target resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the atracker_targets.
* `targets` - A list of target resources. Nested `targets` blocks have the following structure:
	* `id` - The uuid of the target resource.
	* `name` - The name of the target resource.
	* `crn` - The crn of the target resource.
	* `target_type` - The type of the target.
	* `encrypt_key` - The encryption key that is used to encrypt events before Activity Tracking services buffer them on storage. This credential is masked in the response.
	* `cos_endpoint` - Property values for a Cloud Object Storage Endpoint. Nested `cos_endpoint` blocks have the following structure:
		* `endpoint` - The host name of the Cloud Object Storage endpoint.
		* `target_crn` - The CRN of the Cloud Object Storage instance.
		* `bucket` - The bucket name under the Cloud Object Storage instance.
		* `api_key` - The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response.
	* `cos_write_status` - The status of the write attempt with the provided cos_endpoint parameters. Nested `cos_write_status` blocks have the following structure:
		* `status` - The status such as failed or success.
		* `last_failure` - The timestamp of the failure.
		* `reason_for_last_failure` - Detailed description of the cause of the failure.
	* `created` - The timestamp of the target creation time.
	* `updated` - The timestamp of the target last updated time.

