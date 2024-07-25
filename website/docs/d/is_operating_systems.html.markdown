---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : operating_systems"
description: |-
  Manages IBM Operating Systems.
---

# ibm_is_operating_systems
Retrieve information of an existing IBM Cloud Infrastructure Operating Systems as a read only data source. For more information, about supported Operating System, see [Images](https://cloud.ibm.com/docs/vpc?topic=vpc-about-images).

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
data "ibm_is_operating_systems" "example"{
}
```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `operating_systems` - (List) List of all Operating Systems in the IBM Cloud Infrastructure region.

  Nested scheme for `operating_system`:
  - `allow_user_image_creation` - (String) Users may create new images with this operating system.
  - `architecture` - (String) The Operating System architecture.
  - `dedicated_host_only` - (String) Images with this Operating System can only be used on dedicated hosts or dedicated host groups.
  - `display_name` - (String) A unique, display-friendly name for the Operating System.
  - `family` - (String) The name of the software family this Operating System belongs to.
  - `href` - (String) The URL for this Operating System.
  - `name` - (String) The globally unique name for this Operating System.
  - `user_data_format` - (String) The user data format for this image.
  
    ~> **Note:** </br> Supported values are : </br>
    **&#x2022;** `cloud_init`: user_data will be interpreted according to the cloud-init standard.</br>
    **&#x2022;** `esxi_kickstart`: user_data will be interpreted as a VMware ESXi installation script.</br>
    **&#x2022;**  `ipxe`: user_data will be interpreted as a single URL to an iPXE script or as the text of an iPXE script.</br>
  - `vendor` - (String) The vendor of the Operating System.
  - `version` - (String) The major release version of this Operating System.
