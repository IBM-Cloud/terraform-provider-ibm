---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_connection_prefix_filter"
description: |-
  Gets Information about IBM Cloud Infrastructure Transit Gateway Connection Prefix Filter.
---

# ibm_tg_connection_prefix_filter
Retrieve information of an existing IBM Cloud infrastructure transit gateway connection prefix filter as a read only data source. For more information about Transit Gateway Prefix filters, see [adding and deleting prefix filters](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-adding-prefix-filters&interface=ui).

## Example usage

```terraform
data "ibm_tg_connection_prefix_filter" "tg_prefix_filter" {
    gateway = ibm_tg_gateway.new_tg_gw.id
    connection_id = ibm_tg_connection.test_ibm_tg_connection.connection_id
	filter_id = ibm_tg_connection_prefix_filter.test_tg_prefix_filter.filter_id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `gateway` - (Required, String) The unique identifier of the gateway.
- `connection_id` - (Required, String) The unique identifier of the gateway connection
- `filter_id` - - (Required, String) The unique identifier of the gateway connection prefix filter

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `created_at` - (String) The date and time resource is created.
- `id` - (String) The unique identifier of this prefix filter.
- `action` - (String) Whether to permit or deny the prefix filter
- `before` - (String) Identifier of prefix filter that handles the ordering and follow semantics. When a filter reference another filter in it's before field, then the filter making the reference is applied before the referenced filter. For example: if filter A references filter B in its before field, A is applied before B.
- `ge` - (Int) The IP Prefix GE. The GE (greater than or equal to) value can be included to match all less-specific prefixes within a parent prefix above a certain length.
- `le` - (Int) The IP Prefix LE. The LE (less than or equal to) value can be included to match all more-specific prefixes within a parent prefix up to a certain length.
- `prefix` - (String) The IP Prefix
- `updated_at` - (String) The date and time resource is last updated.