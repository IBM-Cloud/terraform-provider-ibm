---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_group"
description: |-
  Get information about dedicated host group.
---

# ibm_is_dedicated_host_group
Retrieve the dedicated host group data sources. For more information, about dedicated host group in your IBM Cloud VPC, see [dedicated hosts groups](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-dedicated-hosts-instances).

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
data "ibm_is_dedicated_host_group" "example" {
  name = "example-host-group"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The unique user defined name of this dedicated host group.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `class` -  (String) The dedicated host profile class for hosts in this group.
- `created_at` -  (String) The date and time that the dedicated host group was created.
- `crn` -  (String) The CRN for this dedicated host group.
- `dedicated_hosts` -  (List) The dedicated hosts that are in this dedicated host group. 
  
  Nested scheme for `dedicated_hosts`:
  - `crn` -  (String) The CRN for this dedicated host.
  - `deleted` -  (List) If present, this property indicates the referenced resource has been deleted and provides supplementary information. Nested deleted blocks have the following structure.

    Nested scheme for `deleted`:
    - `more_info` -  (String) Link to documentation about deleted resources.
  - `href` -  (String) The URL for this dedicated host.
  - `id` -  (String) The unique identifier for this dedicated host.
  - `name` -  (String) The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly selected words.
  - `resource_type` -  (String) The type of resource referenced.
- `family` -  (String) The dedicated host profile family for hosts in this group.
- `href` -  (String) The URL for this dedicated host group.
- `id` -  (String) The unique identifier of the dedicated host group.
- `resource_group` -  (String) The unique identifier of the resource group for this dedicated host.
- `resource_type` -  (String) The type of resource referenced.
- `supported_instance_profiles` -  (List) Array of instance profiles that can be used by instances placed on this dedicated host group. 

  Nested scheme for `supported_instance_profiles`:
  - `href` -  (String) The URL for this virtual server instance profile.
  - `name` -  (String) The unique name for this virtual server instance profile.
- `zone` -  (String) The zone this dedicated host group resides in.


