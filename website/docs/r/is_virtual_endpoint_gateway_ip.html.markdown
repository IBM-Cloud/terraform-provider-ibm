---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway_ip"
sidebar_current: "docs-ibm-resource-is-virtual-endpoint-gateway-ip"
description: |-
  Manages IBM Virtual endpoint gateway IP
---

# ibm_is_virtual_endpoint_gateway_ip

Provides a Virtual endpoint gateway resource. This allows Virtual endpoint gateway IP to be created, updated, and cancelled.

## Example Usage

In the following example, you can create a endpoint gateway IP:

```hcl
resource "ibm_is_virtual_endpoint_gateway_ip" "virtual_endpoint_gateway_ip" {
	gateway     = ibm_is_virtual_endpoint_gateway.endpoint_gateway.id
	reserved_ip = "0737-5ab3c18e-6f6c-4a69-8f48-20e3456647b5"
}

```

## Argument Reference

The following arguments are supported:

- `gateway` - (Required, string,ForceNew) Endpoint gateway ID
- `reserved_ip` - (Required, string,ForceNew) Endpoint gateway IP id

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the endpoint gateway connection. The id is composed of <`gateway_id`>/<`reserved_ip_id`>.
- `name` - Endpoint gateway IP name
- `created_at` - Endpoint gateway IP created date and time
- `resource_type` - Endpoint gateway IP resource type
- `auto_delete` - Endpoint gateway IP auto delete
- `address` - Endpoint gateway IP address
- `target` - Endpoint gateway detail
  - `id` - The IPs target id
  - `name` - The IPs target name
  - `resource_type` - Endpoint gateway resource type

## Import

ibm_is_virtual_endpoint_gateway_ip can be imported using virtual endpoint gateway ID and gateway ip id, eg

```
$ terraform import ibm_is_virtual_endpoint_gateway_ip.example d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19d35f

```
