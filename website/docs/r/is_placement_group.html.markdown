---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_placement_group"
description: |-
  Manages PlacementGroup.
---

# ibm_is_placement_group

Create, update, or delete a placement group. For more information, about placement group, see [managing placement groups](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-placement-group&interface=ui).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_placement_group" "example" {
  strategy = "host_spread"
  name     = "example-placement-group"
}
```

## Argument reference

Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the placement group.

  ~> **Note:** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag).
- `name` - (Optional, string) The unique user-defined name for this placement group. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `resource_group` - (Optional, string, Forces new resource) The unique identifier of the resource group to use. If unspecified, the account's 
- `strategy` - (Required, string, Forces new resource) The strategy for this placement group- `host_spread`: place on different compute hosts- `power_spread`: place on compute hosts that use different power sources. The enumerated values for this property may expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the placement group on which the unexpected strategy was encountered.
[default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
- `tags`  - (Optional, List of Strings) The user tags to attach to the placement group.


## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - The unique identifier of the PlacementGroup.
- `created_at` - The date and time that the placement group was created.
- `crn` - The CRN for this placement group.
- `href` - The URL for this placement group.
- `lifecycle_state` - The lifecycle state of the placement group.
- `resource_type` - The resource type.

## Import

The `ibm_is_placement_group` can be imported by using ID.

**Example**

```
$ terraform import ibm_is_placement_group.example d7bec597-4726-451f-8a63-e62e6f19c32c
```