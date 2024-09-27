---
subcategory: "Global Tagging"
layout: "ibm"
page_title: "IBM : resource_tag"
description: |-
  Manages resource tags.
---

# ibm_resource_tag

Create, update, or delete IBM Cloud resource tags. For more information, about tagging, see [IBM Cloud resource tags](https://cloud.ibm.com/apidocs/tagging).


## Example usage
The following example enables you to attach resource tags

```terraform
data "ibm_satellite_location" "location" {
location  = var.location
}

resource "ibm_resource_tag" "tag" {
	resource_id = ibm_satellite_location.location.crn
	tags        = var.tag_names
}

resource "ibm_resource_tag" "tag_with_timeout" {
	resource_id = ibm_satellite_location.location.crn
	tags        = var.tag_names
	timeout		{
		create = "45s"
	}
}

```

## Timeouts
The `ibm_resource_tag` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The default timeout for tag creation is 30 seconds. 


## Argument reference
Review the argument references that you can specify for your resource.

- `resource_id` - (Required, String) The CRN of the resource on which the tags is be attached.
- `resource_type` - (Optional, String) The resource type on which the tags should be attached.
- `tag_type` - (Optional, String) Type of the tag. Supported values are: `user` or `access`. The default value is user.
- `tags` - (Required, Array of strings) List of tags associated with resource instance.
- `replace` - (Optional, Bool) If true, it indicates that the attaching operation is a replacement operation

## Attributes reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the resource tag.
- `acccount_id` - (String) The ID of an account that owns the resources to be tagged. This is required if tag-type is set to service.


## Import

The `ibm_resource_tag` resource can be imported by using the resource CRN.

**Syntax**

```
$ terraform import ibm_resource_tag.tag resource_id
```

**Example**

```
$ terraform import ibm_resource_tag.tag  crn:v1:bluemix:public:satellite:us-east:a/ab3ed67929c2a81285fbb5f9eb22800a:c1ga7h9w0angomd44654::

```

Example for importing classic infrastructure tags.

**Syntax**

```
$ terraform import ibm_resource_tag.tag resource_id/resource_type
```

**Example**

```
$ terraform import ibm_resource_tag.tag 118398132/SoftLayer_Virtual_Guest
```
