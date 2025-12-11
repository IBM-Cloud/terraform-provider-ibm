---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateways"
description: |-
  Manages IBM Cloud infrastructure virtual endpoint gateways.
---

# ibm_is_virtual_endpoint_gateways
Retrieve information of an existing IBM Cloud infrastructure virtual endpoint gateways as a read-only data source. For more information, about the VPC endpoint gateways, see [creating an endpoint gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-endpoint-gateway).

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
data "ibm_is_virtual_endpoint_gateways" "example" {
}
```

## Argument reference

Review the argument references that you can specify for your data source. 

- `dns_resolution_binding_mode` - (Optional, List) Filters the collection to endpoint gateways with a `dns_resolution_binding_mode` property matching one of the specified comma-separated values.
  - Constraints: Allowable list items are: `disabled`, `per_resource_binding`, `primary`.
- `resource_group` - (String) The ID of the Resource group this endpoint gateway belongs to
- `name` - (String) The name of the endpoint gateway

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `virtual_endpoint_gateways` - (List) List of Endpoint Gateways in the IBM Cloud infrastructure region.
  
  Nested scheme for `virtual_endpoint_gateways`:
  - `access_tags`  - (List) Access management tags associated for the virtual endpoint gateway.

  - `allow_dns_resolution_binding` - (**Deprecated**, Optional, bool) **This field has been deprecated in favor of `dns_resolution_binding_mode` and will be removed in a future version.** 
    
    Previously indicated whether to allow this endpoint gateway to participate in DNS resolution bindings with a VPC that has dns.enable_hub set to true.
    
    **Migration Guide:**
    - `false` -> use `dns_resolution_binding_mode = "disabled"`
    - `true` -> use `dns_resolution_binding_mode = "primary"`
    
    **Note:** The new `dns_resolution_binding_mode` field also supports `"per_resource_binding"` for advanced DNS sharing scenarios not available with this boolean field.
    
    ~> **Important:** Do not use both `allow_dns_resolution_binding` and `dns_resolution_binding_mode` in the same configuration. Use only `dns_resolution_binding_mode`.
  - `created_at` - (Timestamp) The created date and time of the endpoint gateway.
  - `crn` - (String) The CRN for this endpoint gateway.
  - `dns_resolution_binding_mode` - (String) The DNS resolution binding mode used for this endpoint gateway:- `disabled`: The endpoint gateway is not participating in [DNS sharing for VPE   gateways](/docs/vpc?topic=vpc-vpe-dns-sharing).- `primary`: The endpoint gateway is participating in [DNS sharing for VPE gateways]   (/docs/vpc?topic=vpc-vpe-dns-sharing) if the VPC this endpoint gateway resides in   has a DNS resolution binding to another VPC.- `per_resource_binding`: The endpoint gateway is participating in [DNS sharing for VPE   gateways](/docs/vpc?topic=vpc-vpe-dns-sharing) if the VPC this endpoint gateway   resides in has a DNS resolution binding to another VPC, and resource binding is   enabled for the `target` service.
    - Constraints: Allowable values are: `disabled`, `per_resource_binding`, `primary`.
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
  - `security_groups` (List) - The security groups to use for this endpoint gateway.
  - `service_endpoints`- (Array of Strings) The fully qualified domain names for the target service. A fully qualified domain name for the target service