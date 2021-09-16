---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateways"
description: |-
  Manages IBM Cloud infrastructure virtual endpoint gateways.
---

# ibm_is_virtual_endpoint_gateways
Retrieve information of an existing IBM Cloud infrastructure virtual endpoint gateways as a read-only data source. For more information, about the VPC endpoint gateways, see [creating an endpoint gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-endpoint-gateway).

## Example usage

```terraform
data "ibm_is_virtual_endpoint_gateways" "data_test" {
}
```

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `virtual_endpoint_gateways` - (List) List of Endpoint Gateways in the IBM Cloud infrastructure region.
  
  Nested scheme for `virtual_endpoint_gateways`:
  - `created_at` - (Timestamp) The created date and time of the endpoint gateway.
  - `crn` - (String) The CRN for this endpoint gateway.
  - `health_state` - (String) The endpoint gateway health state. **ok: Healthy**, **degraded: Suffering from compromised performance, capacity, or connectivity**, **faulted: Completely unreachable, inoperative, or entirely incapacitated**, **inapplicable: The health state does not apply because of the current lifecycle state**. A resource with a lifecycle state of failed or deleting will have a health state of inapplicable. A pending resource may have this state.
  - `lifecycle_state` - (String) The endpoint gateway lifecycle state, supported values are `deleted`, `deleting`, `failed`, `pending`, `stable`, `updating`, `waiting`, `suspended`.
  - `id` - (String) The endpoint gateway ID.
  - `ips` - (List) The collection of reserved IPs bound to an endpoint gateway.
  
    Nested scheme for `ips`:
    - `ips.id` - (String) The unique identifier for the reserved IP.
    - `ips.name` - (String) The user defined or system provided name of the resource IP.
    - `ips.resource_type` - (String) The endpoint gateway IP resource type or the subnet reserved IP.
  - `name` - (String) The endpoint gateway name.
  - `resource_group` - (String) The unique identifier for the resource group.
  - `target` - (List) The endpoint gateway target services.
  
    Nested scheme for `target`:
    - `crn` - (String) The endpoint gateway target CRN.
    - `name` - (String) The endpoint gateway target name.
    - `resource_type` - (String) The endpoint gateway target resource type.
  - `vpc` - (String) The VPC ID.
