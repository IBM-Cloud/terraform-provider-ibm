---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : floating_ip"
description: |-
  Manages IBM floating IP.
---

# ibm_is_floating_ip
Create a floating IP address that you can associate with a Virtual Servers for VPC instance. You can use the floating IP address to access your instance from the public network, independent of whether the subnet is attached to a public gateway. For more information, see [about floating IP](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-using-the-rest-apis#create-floating-ip-api-tutorial).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example shows how to create a Virtual Servers for VPC instance and associate a floating IP address to the primary network interface of the virtual server instance.

```terraform

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "b-2x8"

  primary_network_interface {
    port_speed = "1000"
    subnet     = "70be8eae-134c-436e-a86e-04849f84cb34"
  }

  vpc  = "01eda778-b822-43a2-816d-d30713df5e13"
  zone = "us-south-1"
  keys = ["eac87f33-0c00-4da7-aa66-dc2d972148bd"]
}

resource "ibm_is_floating_ip" "testacc_floatingip" {
  name   = "testfip1"
  target = ibm_is_instance.testacc_instance.primary_network_interface[0].id
}

```

## Timeouts
The `ibm_is_instance` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the floating IP address is considered `failed` if no response is received for 10 minutes. 
- **delete**: The deletion of the floating IP address is considered `failed` if no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, String) Enter a name for the floating IP address. 
- `resource_group` - (Optional, String) The resource group ID where you want to create the floating IP.
- `target` - (Optional, String) Enter the ID of the network interface that you want to use to allocate the IP address. If you specify this option, do not specify `zone` at the same time. **Note** conflicts with `zone`. A change in `target` which is in a different `zone` will show a change to replace current floating ip with a new one.
- `tags` (Optional, Array of Strings) Enter any tags that you want to associate with your VPC. Tags might help you find your VPC more easily after it is created. Separate multiple tags with a comma (`,`).
- `zone` - (Optional, Force New Resource, String) Enter the name of the zone where you want to create the floating IP address. To list available zones, run `ibmcloud is zones`. If you specify this option, do not specify `target` at the same time. **Note** Conflicts with `target` and one of `target`, or `zone` is mandatory.

~> **Note**
  - `target` cannot be used in conjunction with the `floating_ip` argument of `ibm_is_instance_network_interface` resource and might cause cyclic dependency/unexpected issues if used used both ways.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `address` - (String) The floating IP address that was created. 
- `crn` - (String) The CRN for this floating IP. 
- `id` - (String) The unique identifier of the floating IP address. 
- `status` - (String) The provisioning status of the floating IP address.


## Import
The `ibm_is_floating_ip` resource can be imported by using floating IP ID.

**Example**

```
$ terraform import ibm_is_floating_ip.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
