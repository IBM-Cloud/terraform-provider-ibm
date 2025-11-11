---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_grs_location_pairs"
description: |-
  Get information about pdr_get_grs_location_pairs
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_grs_location_pairs

Provides a read-only data source to retrieve information about pdr_get_grs_location_pairs. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_get_grs_location_pairs" "pdr_get_grs_location_pairs" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_grs_location_pairs.
* `location_pairs` - (Map) A map of GRS location pairs where each key is a primary location and the value is its paired location.

