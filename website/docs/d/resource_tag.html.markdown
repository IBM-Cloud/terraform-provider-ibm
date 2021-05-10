---
subcategory: "Global Tagging"
layout: "ibm"
page_title: "IBM : resource_tag"
description: |-
  Manages resource tags.
---

# ibm\resource_tag

Import the details of an existing resource tags as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.


## Example Usage

###  Attach resource tags

```hcl

data "ibm_satellite_location" "location" {
  location  = var.location
}

data "ibm_resource_tag" "read_tag" {
	resource_id = data.ibm_satellite_location.location.crn
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Required, string) CRN of the resource on which the tags should be attached.
* `resource_type` - (Optional, string) Resource type on which the tags should be attached.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the resource tag.
* `tags` - List of tags associated with resource instance.
