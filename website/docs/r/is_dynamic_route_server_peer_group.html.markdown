---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_peer_group"
description: |-
  Manages IBM dynamic route server peer group.
---

# ibm_is_dynamic_route_server_peer_group

Create, update, and delete a peer group of a dynamic route server. A peer group collects BGP peers of the same type so that identical policies and health monitoring apply to all of them. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc).

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

The following example creates a dynamic route server and a peer group with one BGP peer.

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
  name      = "example-drs"
  local_asn = 64512
  members {
    name = "example-member"
    virtual_network_interfaces {
      id = ibm_is_virtual_network_interface.example.id
    }
  }
}

resource "ibm_is_dynamic_route_server_peer_group" "example" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
  name                    = "example-peer-group"
  roles                   = ["next_hop"]
  state                   = "enabled"

  health_monitor {
    mode  = "bgp"
    roles = ["next_hop"]
  }

  peers {
    asn      = 64513
    name     = "example-peer"
    priority = 1
    state    = "enabled"
    endpoint {
      address = "10.240.0.10"
      gateway {
        address = "10.240.0.1"
      }
    }
  }

  prefix_watcher {
    monitored_prefixes {
      prefix = "192.168.0.0/16"
      ge     = 16
      le     = 24
    }
    excluded_prefixes {
      prefix = "192.168.100.0/24"
    }
  }
}
```

## Timeouts

The `ibm_is_dynamic_route_server_peer_group` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating the peer group and waiting for it to become `stable`.
- **update** - (Default 10 minutes) Used for updating the peer group and waiting for it to become `stable`.
- **delete** - (Default 10 minutes) Used for deleting the peer group and waiting for it to be removed.

## Argument Reference

You can specify the following arguments for this resource.

- `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
- `health_monitor` - (Required, List) The health monitoring configuration used for this dynamic route server peer group. Exactly one block.
  Nested schema for **health_monitor**:
  - `mode` - (Optional, String) The mode used for this health monitor. Allowed values are `bfd`, `bgp` and `none`. Defaults to the value chosen by the service.
  - `roles` - (Required, Forces new resource, List of String) The roles used for this health monitor configuration. Allowed values are `next_hop` and `router`.
- `name` - (Optional, String) The name for this dynamic route server peer group. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `peers` - (Required, Forces new resource, List) The peers created with this peer group. At least one block. Additional peers can be managed with the `ibm_is_dynamic_route_server_peer_group_peer` resource. Peers added that way are also returned in this list; use `lifecycle { ignore_changes = [peers] }` on this resource when mixing both approaches.
  Nested schema for **peers**:
  - `asn` - (Required, Integer) The autonomous system number (ASN) for this peer.
  - `endpoint` - (Required, List) The endpoint for this peer. Exactly one block. Specify either the reserved IP (`id` or `href`) or an `address` with an optional `gateway`.
    Nested schema for **endpoint**:
    - `address` - (Optional, String) The IP address of the endpoint.
    - `gateway` - (Optional, List) The gateway used to reach the endpoint address. At most one block.
      Nested schema for **gateway**:
      - `address` - (Required, String) The IP address of the gateway.
    - `href` - (Optional, String) The URL for the reserved IP used as the endpoint.
    - `id` - (Optional, String) The unique identifier for the reserved IP used as the endpoint.
  - `name` - (Optional, String) The name for this peer. The name must not be used by another peer in the peer group.
  - `priority` - (Optional, Integer) The priority of this peer. Lower values have higher priority.
  - `state` - (Optional, String) The state of this peer. Allowed values are `enabled` and `disabled`.
- `prefix_watcher` - (Optional, List) The prefix watcher configuration for this dynamic route server peer group. At most one block.
  Nested schema for **prefix_watcher**:
  - `excluded_prefixes` - (Optional, List) The prefixes excluded from monitoring.
    Nested schema for **excluded_prefixes**:
    - `ge` - (Optional, Integer) The minimum prefix length to match.
    - `le` - (Optional, Integer) The maximum prefix length to match.
    - `prefix` - (Required, String) The IP prefix in CIDR format.
  - `monitored_prefixes` - (Optional, List) The prefixes monitored by this peer group.
    Nested schema for **monitored_prefixes**:
    - `ge` - (Optional, Integer) The minimum prefix length to match.
    - `le` - (Optional, Integer) The maximum prefix length to match.
    - `prefix` - (Required, String) The IP prefix in CIDR format.
- `roles` - (Required, Forces new resource, List of String) The roles for this dynamic route server peer group. Allowed values are `next_hop` and `router`.
- `state` - (Optional, String) The state of this dynamic route server peer group. Allowed values are `enabled` and `disabled`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the resource in the format `<dynamic_route_server_id>/<dynamic_route_server_peer_group_id>`.
- `created_at` - (String) The date and time that this dynamic route server peer group was created.
- `creator` - (List of String) The types of resource that created peers in this peer group.
- `dynamic_route_server_peer_group_id` - (String) The unique identifier for this dynamic route server peer group.
- `etag` - (String) ETag identifier for the dynamic route server peer group.
- `href` - (String) The URL for this dynamic route server peer group.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  Nested schema for **lifecycle_reasons**:
  - `code` - (String) A reason code for this state.
  - `message` - (String) An explanation of the reason for this state.
  - `more_info` - (String) A link to documentation about the reason for this state.
- `lifecycle_state` - (String) The lifecycle state of the dynamic route server peer group.
- `peers` - (List) In addition to the arguments above, each peer exposes:
  - `id` - (String) The unique identifier for this peer.
- `resource_type` - (String) The resource type.
- `status` - (String) The status of this dynamic route server peer group.
- `status_reasons` - (List) The status reasons for the dynamic route server peer group.
  Nested schema for **status_reasons**:
  - `code` - (String) A reason code for this state.
  - `message` - (String) An explanation of the reason for this state.
  - `more_info` - (String) A link to documentation about the reason for this state.

## Import

You can import the `ibm_is_dynamic_route_server_peer_group` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, and `dynamic_route_server_peer_group_id` in the following format:

<pre>
&lt;dynamic_route_server_id&gt;/&lt;dynamic_route_server_peer_group_id&gt;
</pre>

- `dynamic_route_server_id`: A string. The dynamic route server identifier.
- `dynamic_route_server_peer_group_id`: A string. The dynamic route server peer group identifier.

**Syntax**

```
$ terraform import ibm_is_dynamic_route_server_peer_group.example <dynamic_route_server_id>/<dynamic_route_server_peer_group_id>
```

**Example**

```
$ terraform import ibm_is_dynamic_route_server_peer_group.example r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/r006-8228af60-189d-11ed-861d-0242ac120002
```
