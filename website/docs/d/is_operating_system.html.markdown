---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : operating_system"
description: |-
  Manages IBM Operating System.
---

# ibm_is_operating_system
Retrieve information of an existing Operating System as a read only data source. For more information, about supported Operating System, see [Images](https://cloud.ibm.com/docs/vpc?topic=vpc-about-images).

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
data "ibm_is_operating_system" "example" {
  name = "centos-7-amd64"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The global unique name of an Operating System.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `allow_user_image_creation` - (String) Users may create new images with this operating system.
- `architecture` - (String) The Operating System architecture.
- `dedicated_host_only` - (String) Images with the Operating System can be used on dedicated hosts or dedicated host groups.
- `display_name` - (String) A unique, user friendly name for the Operating System.
- `family` - (String) The name of the software family the Operating System belongs to.
- `href` - (String) The URL of an Operating System.
- `id` - (String) The globally unique name of an Operating System.
- `name` - (String) The global unique name of an Operating System.
- `user_data_format` - (String) The user data format for this image.
  
  ~> **Note:** </br> Supported values are : </br>
  **&#x2022;** `cloud_init`: user_data will be interpreted according to the cloud-init standard.</br>
  **&#x2022;** `esxi_kickstart`: user_data will be interpreted as a VMware ESXi installation script.</br>
  **&#x2022;**  `ipxe`: user_data will be interpreted as a single URL to an iPXE script or as the text of an iPXE script.</br>
- `vendor` - (String) The vendor of the Operating System.
- `version` - (String) The major release version of an Operating System.
