---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway"
sidebar_current: "docs-ibm-datasource-is-virtual-endpoint-gateway"
description: |-
  Manages IBM Cloud Infrastructure virtual endpoint gateway .
---

# ibm_is_virtual_endpoint_gateway

Import the details of an existing IBM Cloud Infrastructure virtual endpoint gateway  as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_virtual_endpoint_gateway" "data_test" {    
    name = ibm_is_virtual_endpoint_gateway.endpoint_gateway.name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Endpoint gateway name


## Attribute Reference

The following attributes are exported:

- `resource_group` - The unique identifier for this resource group
- `created_at` - Endpoint gateway created date and time
- `health_state` - Endpoint gateway health state(ok: Healthy,degraded: Suffering from compromised performance, capacity, or connectivity,faulted: Completely unreachable, inoperative, or otherwise entirely incapacitated,inapplicable: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of failed or deleting will have a health state of inapplicable. A pending resource may also have this state.)
- `lifecycle_state` - Endpoint gateway lifecycle state(deleted, deleting, failed, pending, stable, updating, waiting, suspended)
- `ips` -  Collection of reserved IPs bound to an endpoint gateway
    - `id` - The unique identifier for this reserved IP
    - `name` - The user-defined or system-provided name for this reserved IP
    - `resource_type` - Endpoint gateway IP resource type
- `target` -  Endpoint gateway target
    - `name` - TThe target name
    - `resource_type` - The resource type(subnet_reserved_ip) 
- `vpc` -  The VPC id   