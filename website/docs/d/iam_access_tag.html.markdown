---
subcategory: "Global Tagging"
layout: "ibm"
page_title: "IBM : iam_access_tag"
description: |-
  Retrieve iam access tag.
---

# ibm_iam_access_tag

Retrieve an existing IBM Cloud access management tag as a read-only data source. For more information, about access tags, see [IBM Cloud access management tags](https://cloud.ibm.com/apidocs/tagging#create-tag).

## Example usage

###  Sample to retrieve an access tag

```terraform

data "ibm_iam_access_tag" "example_access_tag" {
  name  = var.access_tag_name
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the access management tag.

## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the access tag. Same as `name`.
- `tag_type` - (String) Type of the tag(`access`)
