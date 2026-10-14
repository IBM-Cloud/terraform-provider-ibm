---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : dynamic_route_server"
description: |-
  Manages IBM dynamic route server.
---

# ibm_is_dynamic_route_server

Create, update, and delete a dynamic route server. A dynamic route server exchanges routes with BGP peers through its members, which are virtual network interfaces in the VPC. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc?topic=vpc-about-drs).

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

The following example creates a dynamic route server with one member that uses an existing virtual network interface.

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_virtual_network_interface" "example" {
  name   = "example-vni"
  subnet = ibm_is_subnet.example.id
}

resource "ibm_is_dynamic_route_server" "example" {
  name      = "example-dynamic-route-server"
  local_asn = 64512
  members {
    name = "example-member"
    virtual_network_interfaces {
      id = ibm_is_virtual_network_interface.example.id
    }
  }
}
```

The following example creates a member whose virtual network interface is created together with the dynamic route server.

```terraform
resource "ibm_is_dynamic_route_server" "example" {
  name      = "example-dynamic-route-server"
  local_asn = 64512
  resource_group {
    id = "11caaa983d9c4beb82690daab18717e9"
  }
  members {
    name = "example-member"
    virtual_network_interfaces {
      name = "example-member-vni"
      subnet {
        id = ibm_is_subnet.example.id
      }
      primary_ip {
        address = "10.240.0.10"
      }
    }
  }
}
```

**Note:** The dynamic route server API has no operation to change the members of an existing dynamic route server. The `members` argument therefore forces a new resource. To add a member to an existing dynamic route server without replacing it, use the `ibm_is_dynamic_route_server_member` resource.

## Argument Reference

You can specify the following arguments for this resource.

- `local_asn` - (Required, Integer) The local autonomous system number (ASN) used for this dynamic route server and its connections.
- `members` - (Required, Forces new resource, List) The initial members for the dynamic route server. At least one member is required.

  Nested schema for **members**:
  - `name` - (Required, Forces new resource, String) The name for this dynamic route server member. The name is unique across all members in the dynamic route server.
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
- `name` - (Optional, String) The name for this dynamic route server. The name is unique across all dynamic route servers in the region. If unspecified, the name is randomly generated.
  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
- `resource_group` - (Optional, Forces new resource, List) The resource group for this dynamic route server. If unspecified, the account's default resource group is used.

  Nested schema for **resource_group**:
  - `id` - (Required, Forces new resource, String) The unique identifier for this resource group.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - (String) The unique identifier of the dynamic route server.
- `created_at` - (String) The date and time that the dynamic route server was created.
- `crn` - (String) The CRN for this dynamic route server.
- `etag` - (String) The ETag of the dynamic route server, sent as `If-Match` on update and delete.
- `health_reasons` - (List) The reasons for the current `health_state` (if any).

  Nested schema for **health_reasons**:
  - `code` - (String) A reason code for this health state.
  - `message` - (String) An explanation of the reason for this health state.
  - `more_info` - (String) A link to documentation about the reason for this health state.
- `health_state` - (String) The health of this resource: `ok`, `degraded`, `faulted` or `inapplicable`.
- `href` - (String) The URL for this dynamic route server.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).

  Nested schema for **lifecycle_reasons**:
  - `code` - (String) A reason code for this lifecycle state.
  - `message` - (String) An explanation of the reason for this lifecycle state.
  - `more_info` - (String) A link to documentation about the reason for this lifecycle state.
- `lifecycle_state` - (String) The lifecycle state of the dynamic route server.
- `members` - (List) The members, with the following attributes in addition to the arguments.

  Nested schema for **members**:
  - `id` - (String) The unique identifier for this dynamic route server member.
  - `href` - (String) The URL for this dynamic route server member.
  - `resource_type` - (String) The resource type.
  - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

    Nested schema for **deleted**:
    - `more_info` - (String) A link to documentation about deleted resources.
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
- `resource_group` - (List) The resource group, with the following attributes in addition to `id`.

  Nested schema for **resource_group**:
  - `href` - (String) The URL for this resource group.
  - `name` - (String) The name for this resource group.
- `resource_type` - (String) The resource type.
- `vpc` - (List) The VPC this dynamic route server resides in.

  Nested schema for **vpc**:
  - `crn` - (String) The CRN for this VPC.
  - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

    Nested schema for **deleted**:
    - `more_info` - (String) A link to documentation about deleted resources.
  - `href` - (String) The URL for this VPC.
  - `id` - (String) The unique identifier for this VPC.
  - `name` - (String) The name for this VPC.
  - `resource_type` - (String) The resource type.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts) for certain actions:

- `create` - (Default 10 minutes) Used for creating the dynamic route server and waiting for it to become `stable`.
- `update` - (Default 10 minutes) Used for updating the dynamic route server and waiting for it to become `stable`.
- `delete` - (Default 10 minutes) Used for deleting the dynamic route server and waiting for it to be removed.

## Import

You can import the `ibm_is_dynamic_route_server` resource by using `id`, the unique identifier of the dynamic route server.

**Syntax**

```
$ terraform import ibm_is_dynamic_route_server.example <id>
```

**Example**

```
$ terraform import ibm_is_dynamic_route_server.example r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5
```

**Note:** After import, the arguments that are only used when a virtual network interface is created (`allow_ip_spoofing`, `auto_delete`, `enable_infrastructure_nat`, `protocol_state_filtering_mode`, `ips`, `security_groups`, `resource_group` and `primary_ip.auto_delete`) are set to their defaults because the API does not return them. Make sure the configuration matches these defaults, otherwise the next plan replaces the dynamic route server.
