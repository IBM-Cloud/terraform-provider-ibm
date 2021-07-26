---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : object_storage_account"
description: |-
  Manages IBM Cloud Object Storage Account.
---

# ibm_object_storage_account
Retrieve the account name for an existing Object Storage instance within your IBM account. If no Object Storage instance exists, you can use this resource to order an Object Storage instance and to store the account name. For more information, about IBM Cloud Object Storage account, see [exporting an image to IBM Cloud Object Storage](https://cloud.ibm.com/docs/image-templates?topic=image-templates-exporting-an-image-to-ibm-cloud-object-storage).

Do not use this resource for managing the lifecycle of an Object Storage instance in IBM. For lifecycle management, see [Swift API](https://docs.openstack.org/api-ref/object-store/) or [Swift resources](https://github.com/TheWeatherCompany/terraform-provider-swift).

## Example usage

```terraform
resource "ibm_object_storage_account" "foo" {
}
```

## Argument reference 
Review the argument references that you can specify for your resource.

- `tags` - (Optional, Array of String) Tags associated with the Object Storage Account instance.  
  **Note** `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute reference
In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The Object Storage account name, which you can use with Swift resources.
