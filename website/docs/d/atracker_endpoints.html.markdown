---
layout: "ibm"
page_title: "IBM : ibm_atracker_endpoints"
description: |-
  Get information about atracker_endpoints
subcategory: "Activity Tracker"
---

# ibm_atracker_endpoints

Provides a read-only data source for atracker_endpoints. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_atracker_endpoints" "atracker_endpoints" {
}
```


## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the atracker_endpoints.
* `api_endpoint` - (Required, List) Activity Tracker API endpoint.
Nested scheme for **api_endpoint**:
	* `public_url` - (Required, String) The public URL of Activity Tracker in a region.
	* `public_enabled` - (Required, Boolean) Indicates whether or not the public endpoint is enabled in the account.
	* `private_url` - (Required, String) The private URL of Activity Tracker. This URL cannot be disabled.
	* `private_enabled` - (Optional, Boolean) The private endpoint is always enabled.

