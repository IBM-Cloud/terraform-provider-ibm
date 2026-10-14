---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_peer_group_peer"
description: |-
  Manages IBM dynamic route server peer group peer.
---

# ibm_is_dynamic_route_server_peer_group_peer

Create, update, and delete a BGP peer in a dynamic route server peer group. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc).

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

The following example adds a peer to an existing peer group.

```terraform
resource "ibm_is_dynamic_route_server_peer_group_peer" "example" {
  dynamic_route_server_id            = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
  name                               = "example-peer"
  asn                                = 64514
  priority                           = 2
  state                              = "enabled"
  endpoint {
    address = "10.240.0.20"
    gateway {
      address = "10.240.0.1"
    }
  }
}
```

## Timeouts

The `ibm_is_dynamic_route_server_peer_group_peer` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating the peer and waiting for it to become `stable`.
- **update** - (Default 10 minutes) Used for updating the peer and waiting for it to become `stable`.
- **delete** - (Default 10 minutes) Used for deleting the peer and waiting for it to be removed.

## Argument Reference

You can specify the following arguments for this resource.

- `asn` - (Required, Integer) The autonomous system number (ASN) for this dynamic route server peer group peer.
- `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
- `dynamic_route_server_peer_group_id` - (Required, Forces new resource, String) The dynamic route server peer group identifier.
- `endpoint` - (Required, Forces new resource, List) The endpoint for this dynamic route server peer group peer. Exactly one block. Specify either the reserved IP (`id` or `href`) or an `address` with an optional `gateway`.
  Nested schema for **endpoint**:
  - `address` - (Optional, String) The IP address of the endpoint.
  - `gateway` - (Optional, List) The gateway used to reach the endpoint address. At most one block.
    Nested schema for **gateway**:
    - `address` - (Required, String) The IP address of the gateway.
  - `href` - (Optional, String) The URL for the reserved IP used as the endpoint.
  - `id` - (Optional, String) The unique identifier for the reserved IP used as the endpoint.
- `name` - (Optional, String) The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.
- `priority` - (Optional, Integer) The priority of this dynamic route server peer group peer. Lower values have higher priority. Allowed values are `0` to `4`.
- `state` - (Optional, String) The state of this dynamic route server peer group peer. Allowed values are `enabled` and `disabled`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the resource in the format `<dynamic_route_server_id>/<dynamic_route_server_peer_group_id>/<dynamic_route_server_peer_group_peer_id>`.
- `bidirectional_forwarding_detection` - (List) The bidirectional forwarding detection (BFD) configuration for this dynamic route server peer group peer.
  Nested schema for **bidirectional_forwarding_detection**:
  - `detect_multiplier` - (Integer) The detection multiplier used for BFD.
  - `enabled` - (Boolean) Indicates whether BFD is enabled for this peer.
  - `mode` - (String) The BFD mode.
  - `receive_interval` - (Integer) The minimum receive interval in milliseconds.
  - `role` - (String) The BFD role of this peer.
  - `sessions` - (List) The BFD sessions for this peer.
    Nested schema for **sessions**:
    - `local` - (List) The dynamic route server member for the local side of this session.
      Nested schema for **local**:
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
        Nested schema for **deleted**:
        - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for this resource.
      - `id` - (String) The unique identifier for this resource.
      - `name` - (String) The name for this resource.
      - `resource_type` - (String) The resource type.
      - `virtual_network_interfaces` - (List) The virtual network interfaces for this dynamic route server member.
        Nested schema for **virtual_network_interfaces**:
        - `crn` - (String) The CRN for this resource, if any.
        - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
          Nested schema for **deleted**:
          - `more_info` - (String) Link to documentation about deleted resources.
        - `href` - (String) The URL for this resource.
        - `id` - (String) The unique identifier for this resource.
        - `name` - (String) The name for this resource.
        - `resource_type` - (String) The resource type.
    - `remote` - (List) The peer for the remote side of this session.
      Nested schema for **remote**:
      - `crn` - (String) The CRN for this resource, if any.
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
        Nested schema for **deleted**:
        - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for this resource.
      - `id` - (String) The unique identifier for this resource.
      - `name` - (String) The name for this resource.
      - `resource_type` - (String) The resource type.
    - `state` - (String) The state of this BFD session.
  - `transmit_interval` - (Integer) The minimum transmit interval in milliseconds.
- `created_at` - (String) The date and time that this dynamic route server peer group peer was created.
- `creator` - (String) The type of resource that created this peer: `dynamic_route_server` or `transit_gateway`.
- `dynamic_route_server_peer_group_peer_id` - (String) The unique identifier for this dynamic route server peer group peer.
- `etag` - (String) ETag identifier for the dynamic route server peer group peer.
- `href` - (String) The URL for this dynamic route server peer group peer.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  Nested schema for **lifecycle_reasons**:
  - `code` - (String) A reason code for this state.
  - `message` - (String) An explanation of the reason for this state.
  - `more_info` - (String) A link to documentation about the reason for this state.
- `lifecycle_state` - (String) The lifecycle state of the dynamic route server peer group peer.
- `resource_type` - (String) The resource type.
- `sessions` - (List) The BGP sessions for this dynamic route server peer group peer.
  Nested schema for **sessions**:
  - `established_at` - (String) The date and time that this BGP session was established.
  - `local` - (List) The dynamic route server member for the local side of this session. Same nested schema as `bidirectional_forwarding_detection.sessions.local`.
  - `protocol_state` - (String) The BGP protocol state of this session.
  - `remote` - (List) The remote side of this session.
    Nested schema for **remote**:
    - `name` - (String) The name of the remote peer.
- `status` - (String) The status of this dynamic route server peer group peer connection: `down` or `up`.
- `status_reasons` - (List) The status reasons for the dynamic route server peer group peer.
  Nested schema for **status_reasons**:
  - `code` - (String) A reason code for this state.
  - `message` - (String) An explanation of the reason for this state.
  - `more_info` - (String) A link to documentation about the reason for this state.

## Import

You can import the `ibm_is_dynamic_route_server_peer_group_peer` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, `dynamic_route_server_peer_group_id`, and `dynamic_route_server_peer_group_peer_id` in the following format:

<pre>
&lt;dynamic_route_server_id&gt;/&lt;dynamic_route_server_peer_group_id&gt;/&lt;dynamic_route_server_peer_group_peer_id&gt;
</pre>

- `dynamic_route_server_id`: A string. The dynamic route server identifier.
- `dynamic_route_server_peer_group_id`: A string. The dynamic route server peer group identifier.
- `dynamic_route_server_peer_group_peer_id`: A string. The dynamic route server peer group peer identifier.

**Syntax**

```
$ terraform import ibm_is_dynamic_route_server_peer_group_peer.example <dynamic_route_server_id>/<dynamic_route_server_peer_group_id>/<dynamic_route_server_peer_group_peer_id>
```

**Example**

```
$ terraform import ibm_is_dynamic_route_server_peer_group_peer.example r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/r006-8228af60-189d-11ed-861d-0242ac120002/r006-8228af60-189d-11ed-861d-0242ac120004
```
