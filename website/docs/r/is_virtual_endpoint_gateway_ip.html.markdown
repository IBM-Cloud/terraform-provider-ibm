---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway_ip"
description: |-
  Manages IBM Virtual endpoint gateway IP.
---

# ibm_is_virtual_endpoint_gateway_ip
Create, update, or delete a VPC endpoint gateway IP by using virtual endpoint gateway resource. For more information, about the VPC endpoint gateway, see [about VPC gateways](https://cloud.ibm.com/docs/vpc?topic=vpc-about-vpe).

## Example usage
The following example creates a VPN gateway IP.

```terraform
resource "ibm_is_virtual_endpoint_gateway_ip" "virtual_endpoint_gateway_ip" {
	gateway     = ibm_is_virtual_endpoint_gateway.endpoint_gateway.id
	reserved_ip = "0737-5ab3c18e-6f6c-4a69-8f48-20e3456647b5"
}

```


## Argument reference
Review the argument references that you can specify for your resource. 

- `gateway` - (Required, Forces new resource, String) The endpoint gateway ID.
- `reserver_ip` - (Required, Forces new resource, String) The endpoint gateway IP ID.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `address` - (String) The endpoint gateway IP address.
- `auto_delete` - (String) The endpoint gateway IP auto delete.
- `created_at` - (Timestamp) The created date and time of the endpoint gateway IP.
- `id` - (String) The unique identifier of the VPN gateway connection. The ID is composed of `<gateway_id>/<gateway_ip_id>`.
- `name` - (String) The endpoint gateway IP name.
- `resource_type` - (String) The endpoint gateway IP resource type.
- `target` - (List) The endpoint gateway target details.

  Nested scheme for `target`:
  - `id` - (String) The IPs target ID.
  - `name` - (String) The IPs target name.
  - `resource_type` - (String) The endpoint gateway resource type.


## Import
The `ibm_is_virtual_endpoint_gateway_ip` resource can be imported by using virtual endpoint gateway ID and gateway IP ID.

**Example**

```
$ terraform import ibm_is_virtual_endpoint_gateway_ip.example d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19d35f

```
