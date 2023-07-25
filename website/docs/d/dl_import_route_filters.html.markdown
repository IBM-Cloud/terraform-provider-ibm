---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_import_route_filters"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Gateway.
---

# dl_import_route_filters

Import the details of an existing IBM Cloud Infrastructure Direct Link Gateway and its virtual connections. For more information, about IBM Cloud Direct Link, see [getting started with IBM Cloud Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-get-started-with-ibm-cloud-dl).


## Example usage

---
```terraform
data "dl_import_route_filters" "test_dl_import_route_filters" {
    gateway = ibm_dl_gateway.test_dl_gateway.id
}
```
---
## Argument reference
Review the argument reference that you can specify for your resource. 

- `gateway`- (Required, String) Direct Link Gateway ID.


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `import_route_filters` - List of all import route filters for a given direct link gateway.

   Nested scheme for `import_route_filters`:
  - `created_at` - (String) The date and time resource is created.
  - `im_filter_id` - (String) The unique identifier of Import Route Filter.
  - `action` - (String) Whether to permit or deny the prefix filter.
  - `before` - (String) Identifier of prefix filter that handles the ordering and follow semantics. When a filter reference another filter in it's before field, then the filter making the reference is applied before the referenced filter. For example: if filter A references filter B in its before field, A is applied before B.
  - `ge` - (Int) The minimum matching length of the prefix-set.
  - `le` - (Int) The maximum matching length of the prefix-set.
  - `prefix` - (String) IP prefix representing an address and mask length of the prefix-set.
  - `updated_at` - (String) The date and time resource is last updated.

