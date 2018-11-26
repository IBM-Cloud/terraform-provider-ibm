---
layout: "ibm"
page_title: "IBM : object_storage_account"
sidebar_current: "docs-ibm-resource-object-storage-account"
description: |-
  Manages IBM Object Storage Account.
---

# ibm\_object_storage_account

Retrieve the account name for an existing Object Storage instance within your IBM account. If no Object Storage instance exists, you can use this resource to order an Object Storage instance and to store the account name.

Do not use this resource for managing the lifecycle of an Object Storage instance in IBM. For lifecycle management, see the [Swift API](https://developer.openstack.org/api-ref/object-store/) or [Swift resources](https://github.com/TheWeatherCompany/terraform-provider-swift).

## Example Usage

```hcl
resource "ibm_object_storage_account" "foo" {
}
```

## Argument Reference

* `tags` - (Optional, array of strings) Tags associated with the object storage account instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Computed Fields

The following attributes are exported:

* `id` - The Object Storage account name, which you can use with Swift resources.
