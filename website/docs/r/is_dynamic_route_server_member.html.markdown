---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : dynamic_route_server_member"
description: |-
  Manages IBM dynamic route server member.
---

# ibm_is_dynamic_route_server_member

Create and delete a member of a dynamic route server. A member is a set of virtual network interfaces in the VPC through which the dynamic route server exchanges routes with its peers. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc?topic=vpc-about-drs).

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

The following example adds a member that uses an existing virtual network interface to a dynamic route server.

```terraform
resource "ibm_is_virtual_network_interface" "example" {
  name   = "example-member-vni"
  subnet = ibm_is_subnet.example.id
}

resource "ibm_is_dynamic_route_server_member" "example" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
  name                    = "example-member"
  virtual_network_interfaces {
    id = ibm_is_virtual_network_interface.example.id
  }
}
```

The following example adds a member whose virtual network interface is created together with the member.

```terraform
resource "ibm_is_dynamic_route_server_member" "example" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
  name                    = "example-member"
  virtual_network_interfaces {
    name = "example-member-vni"
    subnet {
      id = ibm_is_subnet.example.id
    }
    primary_ip {
      address = "10.240.0.11"
    }
  }
}
```

**Note:** The dynamic route server member API has no update operation. Every argument forces a new resource, so changing an argument deletes the member and creates a new one.

## Argument Reference

You can specify the following arguments for this resource.

- `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
- `name` - (Required, Forces new resource, String) The name for this dynamic route server member. The name is unique across all members in the dynamic route server.
  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
- `virtual_network_interfaces` - (Required, Forces new resource, List) The virtual network interfaces for this member. At least one entry is required. Each entry either identifies an existing virtual network interface with exactly one of `id`, `href` or `crn`, or describes a new virtual network interface to create, in which case `subnet` is required.

  Nested schema for **virtual_network_interfaces**:
  - `id` - (Optional, Forces new resource, String) The unique identifier for an existing virtual network interface.
  - `href` - (Optional, Forces new resource, String) The URL for an existing virtual network interface.
  - `crn` - (Optional, Forces new resource, String) The CRN for an existing virtual network interface.
  - `name` - (Optional, Forces new resource, String) The name for the virtual network interface to create. The name is unique across all virtual network interfaces in the VPC.
  - `subnet` - (Optional, Forces new resource, List) The subnet for the virtual network interface to create. Required when the entry is not an identity. Exactly one of `id`, `crn` or `href` must be set.

    Nested schema for **subnet**:
    - `id` - (Optional, Forces new resource, String) The unique identifier for this subnet.
    - `crn` - (Optional, Forces new resource, String) The CRN for this subnet.
    - `href` - (Optional, Forces new resource, String) The URL for this subnet.
  - `allow_ip_spoofing` - (Optional, Forces new resource, Bool) Indicates whether source IP spoofing is allowed on the virtual network interface to create. Default is `false`.
  - `auto_delete` - (Optional, Forces new resource, Bool) Indicates whether the virtual network interface to create is automatically deleted when the member is deleted. Default is `true`.
  - `enable_infrastructure_nat` - (Optional, Forces new resource, Bool) If `true`, the VPC infrastructure performs any needed NAT operations for the virtual network interface to create. Default is `true`.
  - `protocol_state_filtering_mode` - (Optional, Forces new resource, String) The protocol state filtering mode for the virtual network interface to create. Allowed values are `auto`, `enabled` and `disabled`.
  - `primary_ip` - (Optional, Forces new resource, List) The primary IP for the virtual network interface to create. Identify an existing reserved IP with `id` or `href`, or describe a reserved IP to create with `address`, `name` and `auto_delete`.

    Nested schema for **primary_ip**:
    - `id` - (Optional, Forces new resource, String) The unique identifier for an existing reserved IP.
    - `href` - (Optional, Forces new resource, String) The URL for an existing reserved IP.
    - `address` - (Optional, Forces new resource, String) The IP address to reserve, which must not already be reserved on the subnet.
    - `auto_delete` - (Optional, Forces new resource, Bool) Indicates whether the reserved IP is automatically deleted when the virtual network interface is deleted. Default is `true`.
    - `name` - (Optional, Forces new resource, String) The name for the reserved IP to create.
  - `ips` - (Optional, Forces new resource, List) Additional reserved IPs to bind to the virtual network interface to create. Each entry identifies an existing reserved IP with `id` or `href`, or describes a reserved IP to create.

    Nested schema for **ips**:
    - `id` - (Optional, Forces new resource, String) The unique identifier for an existing reserved IP.
    - `href` - (Optional, Forces new resource, String) The URL for an existing reserved IP.
    - `address` - (Optional, Forces new resource, String) The IP address to reserve, which must not already be reserved on the subnet.
    - `auto_delete` - (Optional, Forces new resource, Bool) Indicates whether the reserved IP is automatically deleted when the virtual network interface is deleted. Default is `true`.
    - `name` - (Optional, Forces new resource, String) The name for the reserved IP to create.
  - `resource_group` - (Optional, Forces new resource, List) The resource group for the virtual network interface to create.

    Nested schema for **resource_group**:
    - `id` - (Required, Forces new resource, String) The unique identifier for this resource group.
  - `security_groups` - (Optional, Forces new resource, List) The security groups for the virtual network interface to create. Each entry must set exactly one of `id`, `crn` or `href`.

    Nested schema for **security_groups**:
    - `id` - (Optional, Forces new resource, String) The unique identifier for this security group.
    - `crn` - (Optional, Forces new resource, String) The CRN for this security group.
    - `href` - (Optional, Forces new resource, String) The URL for this security group.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - (String) The unique identifier of the resource, in the format `<dynamic_route_server_id>/<dynamic_route_server_member_id>`.
- `dynamic_route_server_member_id` - (String) The unique identifier for this dynamic route server member.
- `health_reasons` - (List) The reasons for the current `health_state` (if any).

  Nested schema for **health_reasons**:
  - `code` - (String) A reason code for this health state.
  - `message` - (String) An explanation of the reason for this health state.
  - `more_info` - (String) A link to documentation about the reason for this health state.
- `health_state` - (String) The health of this dynamic route server member: `ok`, `faulted` or `inapplicable`.
- `href` - (String) The URL for this dynamic route server member.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).

  Nested schema for **lifecycle_reasons**:
  - `code` - (String) A reason code for this lifecycle state.
  - `message` - (String) An explanation of the reason for this lifecycle state.
  - `more_info` - (String) A link to documentation about the reason for this lifecycle state.
- `lifecycle_state` - (String) The lifecycle state of the dynamic route server member.
- `resource_type` - (String) The resource type.
- `virtual_network_interfaces` - (List) The virtual network interfaces, with the following attributes in addition to the arguments.

  Nested schema for **virtual_network_interfaces**:
  - `id` - (String) The unique identifier for this virtual network interface.
  - `href` - (String) The URL for this virtual network interface.
  - `crn` - (String) The CRN for this virtual network interface.
  - `name` - (String) The name for this virtual network interface.
  - `resource_type` - (String) The resource type.
  - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

    Nested schema for **deleted**:
    - `more_info` - (String) A link to documentation about deleted resources.
  - `primary_ip` - (List) The primary IP for this virtual network interface, with `id`, `href`, `address`, `name`, `resource_type` and `deleted`.
  - `subnet` - (List) The associated subnet, with `id`, `href`, `crn`, `name`, `resource_type` and `deleted`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) for certain actions:

- `create` - (Default 10 minutes) Used for creating the member and waiting for it to become `stable`.
- `delete` - (Default 10 minutes) Used for deleting the member and waiting for it to be removed.

## Import

You can import the `ibm_is_dynamic_route_server_member` resource by using `id`. The `id` property is a combination of the `dynamic_route_server_id` and the member identifier, separated by `/`.

**Syntax**

```
$ terraform import ibm_is_dynamic_route_server_member.example <dynamic_route_server_id>/<dynamic_route_server_member_id>
```

**Example**

```
$ terraform import ibm_is_dynamic_route_server_member.example r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc
```

**Note:** After import, the arguments that are only used when a virtual network interface is created (`allow_ip_spoofing`, `auto_delete`, `enable_infrastructure_nat`, `protocol_state_filtering_mode`, `ips`, `security_groups`, `resource_group` and `primary_ip.auto_delete`) are set to their defaults because the API does not return them. Make sure the configuration matches these defaults, otherwise the next plan replaces the member.
