---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_grs_location_pairs"
description: |-
  Get information about pdr_get_grs_location_pairs
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_grs_location_pairs

Retrieves the (GRS) location pairs associated with the specified service instance based on managed VMs.

## Example Usage

```hcl
data "ibm_pdr_get_grs_location_pairs" "pdr_get_grs_location_pairs" {
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.(ex., en,it,fr,es,de,ja,ko,pt-BR,zh-HANS,zh-HANT)
* `instance_id` - (Required, Forces new resource, String) ID of the service instance.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_grs_location_pairs.
* `location_pairs` - (Map) A map of GRS location pairs where each key is a primary location and the value is its paired location.
