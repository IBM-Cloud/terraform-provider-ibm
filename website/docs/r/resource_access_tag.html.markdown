---
subcategory: "Global Tagging"
layout: "ibm"
page_title: "IBM : resource_access_tag"
description: |-
  Manages resource access tags.
---

# ibm_resource_access_tag

  ~>**Deprecated:**
  The ability to use the ibm_resource_access_tag resource to create or delete IBM Cloud access management tags in Terraform has been removed in favor of a dedicated ibm_iam_access_tag resource. For more information, check out [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/iam_access_tag)

Create, update, or delete IBM Cloud access management tags. For more information, about tagging, see [IBM Cloud access management tags](https://cloud.ibm.com/apidocs/tagging#create-tag).


## Example usage
The following example enables you to create access management tags

```terraform
resource "ibm_resource_access_tag" "example" {
	name        = "example:tag"
}

```

## Argument reference
Review the argument references that you can specify for your resource.

- `name` - (Required, String) The name of the access management tag.


## Attributes reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the resource tag. Same as `name`.
- `tag_type` - (String) Type of the tag(`access`)


## Import

The `ibm_resource_access_tag` resource can be imported by using the resource CRN.

**Syntax**

```
$ terraform import ibm_resource_access_tag.tag tag_name
```

**Example**

```
$ terraform import ibm_resource_access_tag.tag  crn:v1:bluemix:public:satellite:us-east:a/ab3ed67929c2a81285fbb5f9eb22800a:c1ga7h9w0angomd44654::

```

Example for importing access tags.

**Syntax**

```
$ terraform import ibm_resource_access_tag.tag tag_name
```

**Example**

```
$ terraform import ibm_resource_access_tag.tag example:test
```
