---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_volume"
description: |-
  Manages IBM Cloud VPC volume.
---

# ibm_is_volume
Retrieve information of an existing IBM Cloud VSI volume. For more information, about the volume concepts, see [expandable volume concepts for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-expanding-block-storage-volumes#expandable-volume-concepts).

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
resource "ibm_is_volume" "example" {
  name    = "example-volume"
  profile = "10iops-tier"
  zone    = "us-south-1"
}
data "ibm_is_volume" "example" {
  name = ibm_is_volume.example.name
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `identifier` - (Optional, String) The id of the volume. (one of `identifier`, `name` is required)
- `name` - (Optional, String) The name of the volume. (one of `identifier`, `name` is required)
- `zone` - (Optional, String) The zone name of the volume.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `access_tags`  - (List) Access management tags associated for the volume.
- `active` - (Boolean) Indicates whether a running virtual server instance has an attachment to this volume.
- `attachment_state` - (Boolean) The attachment state of the volume
- `adjustable_capacity_states` - (List) The attachment states that support adjustable capacity for this volume. Allowable list items are: `attached`, `unattached`, `unusable`. 
- `adjustable_iops_states` - (List) The attachment states that support adjustable IOPS for this volume. Allowable list items are: `attached`, `unattached`, `unusable`.
- `bandwidth` - The maximum bandwidth (in megabits per second) for the volume
- `busy` - (Boolean) Indicates whether this volume is performing an operation that must be serialized. This must be `false` to perform an operation that is specified to require serialization.
- `capacity` - (String) The capacity of the volume in gigabytes.
- `catalog_offering` - (List) The catalog offering this volume was created from. If a virtual server instance is provisioned with a boot_volume_attachment specifying this volume, the virtual server instance will use this volume's catalog offering, including its pricing plan.If absent, this volume was not created from a catalog offering.

  Nested scheme for `catalog_offering`:
    - `version_crn` - (String) The CRN for this version of a catalog offering
    - `plan_crn` - (String) The CRN for this catalog offering version's billing plan
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
    
      Nested schema for `deleted`:
        - `more_info`  - (String) Link to documentation about deleted resources.
- `created_at` - (String) The date and time that the volume was created.
- `crn` - (String) The crn of this volume.
- `encryption_key` - (String) The key to use for encrypting this volume.
- `encryption_type` - (String) The type of ecryption used in the volume [**provider_managed**, **user_managed**].
- `health_reasons` - (List) The reasons for the current health_state (if any).

  Nested scheme for `health_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this health state.
  - `message` - (String) An explanation of the reason for this health state.
  - `more_info` - (String) Link to documentation about the reason for this health state.
- `health_state` - (String) The health of this resource.
- `iops` - (String) The bandwidth for the volume.
- `operating_system` - (List) The operating system associated with this volume. If absent, this volume was not created from an image, or the image did not include an operating system.
  Nested scheme for **operating_system**:
  - `architecture` - (String) The operating system architecture
  - `dedicated_host_only` - (Boolean) Images with this operating system can only be used on dedicated hosts or dedicated host groups
  - `display_name` - (String) A unique, display-friendly name for the operating system
  - `family` - (String) The software family for this operating system
  - `href` - (String) The URL for this operating system.
  - `name` - (String) The globally unique name for this operating system.
  - `vendor` - (String) The vendor of the operating system
  - `version` - (String) The major release version of this operating system
- `profile` - (String) The profile to use for this volume.
- `resource_group` - (String) The resource group ID for this volume.
- `source_snapshot` - ID of the snapshot, if volume was created from it.
- `status` - (String) The status of the volume. Supported values are **available**, **failed**, **pending**, **unusable**, **pending_deletion**.
- `status_reasons` - (List) Array of reasons for the current status.
  
  Nested scheme for `status_reasons`:
  - `code` - (String)  A snake case string identifying the status reason.
  - `message` - (String)  An explanation of the status reason
  - `more_info` - (String) Link to documentation about this status reason
- `tags` - (String) User Tags associated with the volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
- `unattached_capacity_update_supported` - (Boolean) Indicates whether the capacity for the volume can be changed when not attached to a running virtual server instance.
- `unattached_iops_update_supported` - (Boolean) Indicates whether the IOPS for the volume can be changed when not attached to a running virtual server instance.
