---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateways"
sidebar_current: "docs-ibm-datasource-is-virtual_endpoint_gateways"
description: |-
  Manages IBM Cloud Infrastructure virtual endpoint gateways .
---

# ibm_is_virtual_endpoint_gateways

Import the details of an existing IBM Cloud Infrastructure virtual endpoint gateways as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_virtual_endpoint_gateways" "data_test" {}
```

## Argument Reference

The following arguments are supported:

## Attribute Reference

The following attributes are exported:

- `id` - Endpoint gateway id
- `name` - Endpoint gateway name
- `resource_group` - Endpoint gateway resource type
- `created_at` - Endpoint gateway created date and time
- `health_state` - Endpoint gateway health state
- `lifecycle_state` - Endpoint gateway lifecycle state
- `ips` - Collection of reserved IPs bound to an endpoint gateway
  - `id` - The unique identifier for this reserved IP
  - `name` - The user-defined or system-provided name for this reserved IP
  - `resource_type` - The resource type(subnet_reserved_ip)
- `target` - Endpoint gateway target services
  - `name` - Endpoint gateway target name
  - `resource_type` - Endpoint gateway target resource type
- `vpc` - The VPC id
