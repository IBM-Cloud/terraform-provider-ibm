---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_filter"
description: |-
  Get information on an IBM Cloud Internet Services Filters.
---

# ibm_cis_filter

Imports a read only copy of an existing Internet Services Filters resource.

## Example Usage

```terraform
data "ibm_cis_filter" "test" {
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
-
- `expression` - The expression of filter.
- `paused` - Whether this filter is currently disabled.
- `description` - Some useful information about this filter to help identify the purpose of it.

