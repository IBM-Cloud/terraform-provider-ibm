---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_link"
description: |-
  Manages iam_trusted_profile_link.
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_link

Create, update, and delete iam_trusted_profile_links with this resource.

## Example Usage

```hcl
resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link_instance" {
  cr_type = "cr_type"
  link {
		crn = "crn"
		namespace = "namespace"
		name = "name"
  }
  profile_id = "profile_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `cr_type` - (Required, Forces new resource, String) The compute resource type. Valid values are VSI, IKS_SA, ROKS_SA.
* `link` - (Required, Forces new resource, List) Link details.
Nested schema for **link**:
	* `crn` - (Optional, String) The CRN of the compute resource.
	* `name` - (Optional, String) Name of the compute resource, only required if cr_type is IKS_SA or ROKS_SA.
	* `namespace` - (Optional, String) The compute resource namespace, only required if cr_type is IKS_SA or ROKS_SA.
* `name` - (Optional, Forces new resource, String) Optional name of the Link.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_trusted_profile_link.
* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
* `entity_tag` - (String) version of the link.
* `link_id` - (String) the unique identifier of the link.
* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.


## Import

You can import the `ibm_iam_trusted_profile_link` resource by using `id`.
The `id` property can be formed from `profile_id`, and `link_id` in the following format:

<pre>
&lt;profile_id&gt;/&lt;link_id&gt;
</pre>
* `profile_id`: A string. ID of the trusted profile.
* `link_id`: A string. the unique identifier of the link.

# Syntax
<pre>
$ terraform import ibm_iam_trusted_profile_link.iam_trusted_profile_link &lt;profile_id&gt;/&lt;link_id&gt;
</pre>
