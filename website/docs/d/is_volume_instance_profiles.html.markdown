---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_volume_instance_profiles"
description: |-
  Get information about VolumeInstanceProfiles
---

# ibm_is_volume_instance_profiles

Provides a read-only data source to retrieve information about an VolumeInstanceProfileCollection.For more information, about the volume concepts, see [expandable volume concepts for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-expanding-block-storage-volumes#expandable-volume-concepts)..

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
data "ibm_is_volume_instance_profiles" "example" {
  identifier = "ibm_is_instance.example.boot_volume.0.volume_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `identifier` - (Required, String) The volume identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the VolumeInstanceProfileCollection.
- `instance_profiles` - (List) A page of instance profiles compatible with the volume.
    
    Nested schema for **instance_profiles**:
	- `href` - (String) The URL for this virtual server instance profile.
	- `name` - (String) The globally unique name for this virtual server instance profile.
	- `resource_type` - (String) The resource type.
