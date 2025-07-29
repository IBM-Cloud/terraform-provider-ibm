---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_image_bare_metal_server_profiles"
description: |-
  Get information about ImageBareMetalServerProfileCollection
---

# ibm_is_image_bare_metal_server_profiles

Provides a read-only data source to retrieve information about an ImageBareMetalServerProfileCollection.For more information, about infrastructure image bare metal server profiles, see [IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

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
data "ibm_is_image_bare_metal_server_profiles" "example" {
	identifier = "ibm_is_image.isExampleImage.id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `identifier` - (Required, String) The image identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the ImageBareMetalServerProfileCollection.
- `bare_metal_server_profiles` - (List) A page of bare metal server profiles compatible with the image.
    
    Nested schema for **bare_metal_server_profiles**:
	- `href` - (String) The URL for this bare metal server profile.
	- `name` - (String) The name for this bare metal server profile.
	- `resource_type` - (String) The resource type.
