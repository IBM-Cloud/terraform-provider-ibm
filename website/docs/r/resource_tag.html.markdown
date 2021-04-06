---
subcategory: "Global Tagging"
layout: "ibm"
page_title: "IBM : resource_tag"
description: |-
  Manages resource tags.
---

# ibm\resource_tag

Create, update, or delete [IBM Cloud resource tags](https://cloud.ibm.com/apidocs/tagging).
Add tags to a resource.


## Example Usage

###  Attach resource tags

```hcl
data "ibm_satellite_location" "location" {
location  = var.location
}

resource "ibm_resource_tag" "tag" {
	resource_id = ibm_satellite_location.location.crn
	tags        = var.tag_names
}

```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Required, string) CRN of the resource on which the tags should be attached.
* `resource_type` - (Optional, string) Resource type on which the tags should be attached.
* `tag_type` - (Optional, string) Type of the tag. Only allowed values are: user, or service or access (default value : user).
* `tags` - (Required, array of strings) List of tags associated with resource instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the resource tag.
* `acccount_id` - The ID of the account that owns the resources to be tagged (required if tag-type is set to service).


## Import

`ibm_resource_tag` can be imported using the resource crn.

Example:

```
$ terraform import ibm_resource_tag.tag resource_id

$ terraform import ibm_resource_tag.tag  crn:v1:bluemix:public:satellite:us-east:a/ab3ed67929c2a81285fbb5f9eb22800a:c1ga7h9w0angomd44654::

```

Example for importing classic infrastructure tags:

```
$ terraform import ibm_resource_tag.tag resource_id/resource_type

$ terraform import ibm_resource_tag.tag  118398132/SoftLayer_Virtual_Guest

```
