---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_dr_locations"
description: |-
  Get information about pdr_get_dr_locations
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_dr_locations

Retrieves the list of disaster recovery (DR) locations available for the specified service instance.
## Example Usage

```hcl
data "ibm_pdr_get_dr_locations" "pdr_get_dr_locations" {
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.(ex., en,it,fr,es,de,ja,ko,pt-BR,zh-HANS,zh-HANT)
* `instance_id` - (Required, Forces new resource, String) ID of the service instance.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_dr_locations.
* `dr_locations` - (List) List of disaster recovery locations available for the service.
Nested schema for **dr_locations**:
	* `id` - (String) Unique identifier of the DR location.
	* `name` - (String) The name of the Power virtual server DR location .
