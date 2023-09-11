---
layout: "ibm"
page_title: "IBM : ibm_scc_report_tags"
description: |-
  Get information about scc_report_tags
subcategory: "Results"
---

# ibm_scc_report_tags

Retrieve information about report tags from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_report_tags" "scc_report_tags" {
	report_id = "report_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `report_id` - (Required, Forces new resource, String) The ID of the scan that is associated with a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_report_tags.
* `tags` - (List) The collection of different types of tags.
Nested schema for **tags**:
	* `access` - (List) The collection of access tags.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	* `service` - (List) The collection of service tags.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	* `user` - (List) The collection of user tags.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.

