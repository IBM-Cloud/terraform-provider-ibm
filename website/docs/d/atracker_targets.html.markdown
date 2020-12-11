---
layout: "ibm"
page_title: "IBM : atracker_targets"
sidebar_current: "docs-ibm-datasource-atracker-targets"
description: |-
  Get information about A list of target resources.
---

# ibm\_atracker_targets

Provides a read-only data source for A list of target resources.. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "atracker_targets" "atracker_targets" {
	name = "a-cos-target-us-south"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of this target resource.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the A list of target resources..
* `targets` - A list of target resources. Nested `targets` blocks have the following structure:
	* `id` - The uuid of this target resource.
	* `name` - The name of this target resource.
	* `instance_id` - The uuid of ATracker services in this region.
	* `crn` - The crn of this target type resource.
	* `target_type` - The type of this target.
	* `encrypt_key` - The encryption key used to encrypt events before ATracker services buffer them on storage. This credential will be masked in the response.
	* `cos_endpoint` - Property values for a Cloud Object Storage Endpoint. Nested `cos_endpoint` blocks have the following structure:
		* `endpoint` - The host name of this COS endpoint.
		* `target_crn` - The CRN of this COS instance.
		* `bucket` - The bucket name under this COS instance.
		* `api_key` - The IAM Api key that have writer access to this cos instance. This credential will be masked in the response.

