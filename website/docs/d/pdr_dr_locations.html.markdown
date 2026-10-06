---
layout: "ibm"
page_title: "IBM : ibm_pdr_dr_locations"
description: |-
  Get information about pdr_dr_locations
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pdr_dr_locations

Retrieves the list of disaster recovery (DR) locations available for the specified service instance.

## Example Usage

```hcl
data "ibm_pdr_dr_locations" "pdr_dr_locations" {
	accept_language = "en-US"
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```
### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) Unique identifier of the API key.
  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `dr_locations` - (List) List of disaster recovery locations available for the service.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.

Nested schema for **dr_locations**:
	* `id` - (String) Unique identifier of the DR location.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `name` - (String) The name of the Power virtual server DR location .
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
	* `power_edge_router` - (Boolean) Indicates whether the region is power edge router enabled or not.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

