---
subcategory: "Global Tagging"
layout: "ibm"
page_title: "IBM : resource_tag"
description: |-
  Retrieve available tags on the account.
---

# ibm_resource_tag

Retreive information about an existing resource or access tags as a read-only data source. For more information, about resource tags, see [controlling access to resources by using tags](https://cloud.ibm.com/docs/account?topic=account-access-tags-tutorial).

## Example usage

###  Sample to attach resource tags

```terraform

data "ibm_satellite_location" "location" {
  location  = var.location
}

data "ibm_resource_tag" "read_tag" {
	resource_id = data.ibm_satellite_location.location.crn
}
```
###  Retrieve access tags

```terraform
data "ibm_resource_tag" "access_tags" {
  tag_type ="access"
}
```
###  Retrieve user tags

```terraform
data "ibm_resource_tag" "user_tags" {
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `resource_id` - (Optional, String) The CRN of the resource on which the tags should be attached.
- `resource_type` - (Optional, String) The resource type on which the tags to be attached.
- `tag_type` - (Optional, String) Type of the tag. Supported values are: `user`, `service`, or `access`. Default: `user`
## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the resource tag.
- `tags` - (String) List of tags associated with resource instance.
