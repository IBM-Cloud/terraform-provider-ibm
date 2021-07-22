---
layout: "ibm"
page_title: "IBM : ibm_atracker_endpoints"
description: |-
  Get information about atracker_endpoints
subcategory: "Activity Tracking API"
---

# ibm_atracker_endpoints

Provides a read-only data source for atracker_endpoints. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_atracker_endpoints" "atracker_endpoints" {
}
```

## Argument Reference

The following arguments are supported:


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the atracker_endpoints.
* `api_endpoint` - Activity Tracking API endpoint. Nested `api_endpoint` blocks have the following structure:
	* `public_url` - The public URL of Activity Tracking in a region.
	* `public_enabled` - Indicates whether or not the public endpoint is enabled in the account.
	* `private_url` - The private URL of Activity Tracking. This URL cannot be disabled.
	* `private_enabled` - The private endpoint is always enabled.

