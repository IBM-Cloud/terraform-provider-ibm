---
layout: "ibm"
page_title: "IBM : ibm_pha_get_supported_location"
description: |-
  Get information about pha_get_supported_location
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_get_supported_location

Provides a read-only data source to retrieve information about a pha_get_supported_location. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pha_get_supported_location" "pha_get_supported_location" {
	if_none_match = "abcdef"
	pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `pha_instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_get_supported_location.
* `dr_locations` - (List) Array of supported DR locations.
  * Constraints: The maximum length is `16` items. The minimum length is `0` items.
Nested schema for **dr_locations**:
	* `id` - (String) Unique identifier for the location.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `name` - (String) Human-readable name of the location.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:()\\- ]+$/`.

