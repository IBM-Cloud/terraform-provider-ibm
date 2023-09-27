---
layout: "ibm"
page_title: "IBM : ibm_scc_provider_type_instance"
description: |-
  Get information about scc_provider_type_instance
subcategory: "Security and Compliance Center"
---

# ibm_scc_provider_type_instance

Retrieve information about a provider type instance from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_provider_type_instance" "scc_provider_type_instance" {
	provider_type_id = ibm_scc_provider_type_instance.scc_provider_type_instance.provider_type_id
	provider_type_instance_id = ibm_scc_provider_type_instance.scc_provider_type_instance_instance.providerTypeInstanceItem_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `provider_type_id` - (Required, Forces new resource, String) The provider type ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-zA-Z0-9 ,\\-_]+$/`.
* `provider_type_instance_id` - (Required, Forces new resource, String) The provider type instance ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-zA-Z0-9 ,\\-_]+$/`

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_provider_type_instance.
* `attributes` - (List) The attributes for connecting to the provider type instance.
Nested schema for **attributes**:

* `created_at` - (String) The time when the resource was created.

* `name` - (String) The name of the provider type instance.

* `provider_type_instance_item_id` - (String) The unique identifier of the provider type instance.

* `type` - (String) The type of the provider type.

* `updated_at` - (String) The time when the resource was updated.

