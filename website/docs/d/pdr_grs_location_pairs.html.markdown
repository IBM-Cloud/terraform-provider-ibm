---
layout: "ibm"
page_title: "IBM : ibm_pdr_grs_location_pairs"
description: |-
  Get information about pdr_grs_location_pairs
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pdr_grs_location_pairs

Retrieves the (GRS) location pairs associated with the specified service instance based on managed VMs.

## Example Usage

```hcl
data "ibm_pdr_grs_location_pairs" "pdr_grs_location_pairs" {
	accept_language = "en-US"
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```
### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) Unique identifier of the API key.
  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `location_pairs` - (Map) A map of GRS location pairs where each key is a primary location and the value is its paired location.
  * Constraints: Each map value has a maximum length of `32` characters. Each map value has a minimum length of `2` characters. Each map value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.


