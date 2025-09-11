---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_link"
description: |-
  Get information about iam_trusted_profile_link
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_link

Provides a read-only data source to retrieve information about an iam_trusted_profile_link. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
	link_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link_instance.link_id
	profile_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link_instance.profile_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `link_id` - (Required, Forces new resource, String) ID of the link.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_trusted_profile_link.
* `cr_type` - (String) The compute resource type. Valid values are VSI, BMS, IKS_SA, ROKS_SA, CE.
* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
* `entity_tag` - (String) version of the link.
* `link` - (List) 
Nested schema for **link**:
	* `component_name` - (String) Component name of the compute resource, only required if cr_type is CE.
	* `component_type` - (String) Component type of the compute resource, only required if cr_type is CE.
	* `crn` - (String) The CRN of the compute resource.
	* `name` - (String) Name of the compute resource, only required if cr_type is IKS_SA or ROKS_SA.
	* `namespace` - (String) The compute resource namespace, only required if cr_type is IKS_SA or ROKS_SA.
* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.
* `name` - (String) Optional name of the Link.

