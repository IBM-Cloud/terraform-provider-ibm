---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_filters"
description: |-
  Get information on an IBM Cloud Internet Services Filters.
---

# ibm_cis_filters

Imports a read only copy of an existing Internet Services Filters resource.

## Example Usage

```terraform
data "ibm_cis_filters" "test" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
- `id` - The Filter ID. It is a combination of <`filter_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":"
- `cis_filters_list` - 
  - `expression` - (Required,string) The expression of filter.
  - `paused` - (Optional,boolean). Whether this filter is currently disabled.
  - `description` - (Optional,string) Some useful information about this filter to help identify the purpose of it.
  - `filter_id` - (Computed, string) The Filter ID.

