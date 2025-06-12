---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Images"
description: |-
  Manages IBM Cloud infrastructure images.
---

# ibm_is_images
Retrieve information of an existing IBM Cloud Infrastructure images as a read-only data source. For more information, about IBM Cloud infrastructure images, see [Images](https://cloud.ibm.com/docs/vpc?topic=vpc-about-images).

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

data "ibm_is_images" "ds_images" {
}

data "ibm_is_images" "ds_images" {
  visibility = "public"
}

```
## Argument reference

Review the argument references that you can specify for your data source. 

- `catalog_managed` - (Optional, bool) Lists only those images which are managed as part of a catalog offering.
- `resource_group` - (Optional, string) The id of the resource group.
- `name` - (Optional, string) The name of the image.
- `visibility` - (Optional, string) Visibility of the image. Accepted values : **private**, **public**
- `status` - (Optional, string) Status of the image. Accepted value : **available**, **deleting**, **deprecated**, **failed**, **obsolete**, **pending**, **unusable**
- `user_data_format` - (String) The user data format for this image.  
   
    ~> **Note:** </br> Allowed values are : </br>
    **&#x2022;** `cloud_init`: user_data will be interpreted according to the cloud-init standard.</br>
    **&#x2022;** `esxi_kickstart`: user_data will be interpreted as a VMware ESXi installation script.</br>
    **&#x2022;**  `ipxe`: user_data will be interpreted as a single URL to an iPXE script or as the text of an iPXE script.</br>

## Attribute reference
You can access the following attribute references after your data source is created. 

- `images` - (List) List of all images in the IBM Cloud Infrastructure.

  Nested scheme for `images`:
  - `access_tags`  - (List) Access management tags associated for image.
  - `architecture` - (String) The architecture for this image.
  - `crn` - (String) The CRN for this image.
  - `catalog_offering` - (List) The catalog offering for this image.

      Nested scheme for **catalog_offering**:
      - `managed` - (Bool) Indicates whether this image is managed as part of a catalog offering. If an image is managed, accounts in the same enterprise with access to that catalog can specify the image's catalog offering version CRN to provision virtual server instances using the image.
      - `version` - (List) The catalog offering version associated with this image. If absent, this image is not associated with a cloud catalog offering.
      
          Nested scheme for **version**:
            - `crn` - (String) The CRN for this version of a catalog offering
  - `checksum` - (String) TThe SHA256 checksum for this image.
  - `encryption` - (String) The type of encryption used on the image.
  - `encryption_key` - (String) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource.
  - `id` - (String) The unique identifier for this image.
  - `name` - (String) The name for this image.
  - `os` - (String) The name of the Operating System.
  - `operating_system` - (List) The operating system details. 
    
      Nested scheme for `operating_system`:
      - `allow_user_image_creation` - (String) Users may create new images with this operating system.
      - `architecture` - (String) The operating system architecture.
      - `dedicated_host_only` - (Bool) Images with this operating system can only be used on dedicated hosts or dedicated host groups.
      - `display_name` - (String) A unique, display-friendly name for the operating system.
      - `family` - (String) The software family for this operating system.
      - `href` - (String) The URL for this operating system.
      - `name` - (String) The globally unique name for this operating system.
      - `user_data_format` - (String) The user data format for this image.
  
      ~> **Note:** </br> Supported values are : </br>
        **&#x2022;** `cloud_init`: user_data will be interpreted according to the cloud-init standard.</br>
        **&#x2022;** `esxi_kickstart`: user_data will be interpreted as a VMware ESXi installation script.</br>
        **&#x2022;**  `ipxe`: user_data will be interpreted as a single URL to an iPXE script or as the text of an iPXE script.</br>
      - `vendor` - (String) The vendor of the operating system.
      - `version` - (String) The major release version of this operating system.
  - `resource_group` - (List) The resource group object, for this image.
      Nested scheme for `resource_group`:
      - `href` - (String) The URL for this resource group.
      - `id` - (String) The unique identifier for this resource group.
      - `name` - (String) The user-defined name for this resource group.
  - `status` - (String) The status of this image.
  - `status_reasons` - (List) The reasons for the current status (if any).

        Nested scheme for `status_reasons`:
      - `code` - (String) The status reason code
      - `message` - (String) An explanation of the status reason
      - `more_info` - (String) Link to documentation about this status reason
  - `visibility` - (String) The visibility of the image public or private.
  - `source_volume` - The source volume id of the image.

