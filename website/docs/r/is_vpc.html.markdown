---
layout: "ibm"
page_title: "IBM : vpc"
sidebar_current: "docs-ibm-resource-is-vpc"
description: |-
  Manages IBM virtual private cloud.
---

# ibm\_is_vpc

Provides a vpc resource. This allows VPC to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a VPC:

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
    name = "test"
}

```

## Argument Reference

The following arguments are supported:

* `default_network_acl` - (Optional, string) ID of the default network ACL.
* `is_default` - (Removed, bool) This field is removed.
* `classic_access` -(Optional, bool) Indicates whether this VPC should be connected to Classic Infrastructure. If true, This VPC's resources will have private network connectivity to the account's Classic Infrastructure resources. Only one VPC on an account may be connected in this way. 
* `name` - (Required, string) The name of the VPC.
* `resource_group` - (Optional, string) The resource group where the VPC to be created
* `tags` - (Optional, array of strings) Tags associated with the instance.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VPC.
* `default_security_group` - The unique identifier of the VPC default security group.
* `status` - The status of VPC.

## Import

ibm_is_vpc can be imported using ID, eg

```
$ terraform import ibm_is_vpc.example d7bec597-4726-451f-8a63-e62e6f19c32c
```