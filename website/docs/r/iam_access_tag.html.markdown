---
subcategory: "Global Tagging"
layout: "ibm"
page_title: "IBM : iam_access_tag"
description: |-
  Manages iam access tags.
---

# ibm_iam_access_tag

Create or delete IBM Cloud access management tags. For more information, about access tags, see [IBM Cloud access management tags](https://cloud.ibm.com/apidocs/tagging#create-tag).


## Example usage
The following example enables you to create access management tags

```terraform
resource "ibm_iam_access_tag" "example" {
	name        = "example:tag"
}

```

## Argument reference
Review the argument references that you can specify for your resource.

- `name` - (Required, String) The name of the access management tag.


## Attributes reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the access tag. Same as `name`.
- `tag_type` - (String) Type of the tag(`access`)


## Import

The `ibm_iam_access_tag` resource can be imported by using the name.

**Syntax**

```
$ terraform import ibm_iam_access_tag.tag tag_name
```

**Example**

Example for importing access tags.

**Syntax**

```
$ terraform import ibm_iam_access_tag.tag tag_name
```

**Example**

```
$ terraform import ibm_iam_access_tag.tag example:test
```
