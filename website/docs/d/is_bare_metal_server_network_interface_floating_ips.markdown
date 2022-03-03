---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_network_interface_floating_ips"
description: |-
  Fetches all floating ip information associated with network interface of a bare metal server.
---

# is_bare_metal_server_network_interface_floating_ips
Retrieve an information of list of floating IPs on a bare metal network interface IBM Cloud as a read-only data source. For more information, about floating IP, see [about floating IP](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-using-the-rest-apis#create-floating-ip-api-tutorial).

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

    data "is_bare_metal_server_network_interface_floating_ips" "test" {
		  bare_metal_server   = "xxxx-xxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx"
		  network_interface   = "xxxx-xxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx"
    }

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `bare_metal_server` - (Required, String) The bare metal server id.
- `network_interface` - (Required, String) The identifier of the bare metal server network interface.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `floating_ips` - (List of objects) A list of all floating ips on the network interface of the bare metal server.

  Nested scheme for `floating_ips`:
    - `address` - (String) The floating IP address.
    - `crn` - (String) The CRN for this floating IP.
    - `id` - (String) The unique identifier of the floating IP.
    - `status` - (String) Provisioning status of the floating IP address.
    - `tags` - (String) The tags associated with VPC.
    - `target` - (String) The ID of the network interface used to allocate the floating IP address.
    - `zone` - (String) The zone name where to create the floating IP address.
