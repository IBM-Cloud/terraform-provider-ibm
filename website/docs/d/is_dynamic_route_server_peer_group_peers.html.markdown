---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_peer_group_peers"
description: |-
  Get information about the peers of an IBM dynamic route server peer group.
---

# ibm_is_dynamic_route_server_peer_group_peers

Provides a read-only data source to list the peers of a dynamic route server peer group. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc).

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
data "ibm_is_dynamic_route_server_peer_group_peers" "example" {
  dynamic_route_server_id            = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `dynamic_route_server_id` - (Required, String) The dynamic route server identifier.
- `dynamic_route_server_peer_group_id` - (Required, String) The dynamic route server peer group identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the data source.
- `peers` - (List) Collection of dynamic route server peer group peers.
  Nested schema for **peers**:
  - `asn` - (Integer) The autonomous system number (ASN) for this dynamic route server peer group peer.
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
  - `endpoint` - (List) The endpoint for this dynamic route server peer group peer.
    Nested schema for **endpoint**:
    - `address` - (String) The IP address of the endpoint.
    - `crn` - (String) The CRN for the transit gateway endpoint, if any.
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
      Nested schema for **deleted**:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `gateway` - (List) The gateway used to reach the endpoint address.
      Nested schema for **gateway**:
      - `address` - (String) The IP address of the gateway.
    - `href` - (String) The URL for the reserved IP used as the endpoint.
    - `id` - (String) The unique identifier for the endpoint resource.
    - `name` - (String) The name for the reserved IP used as the endpoint.
    - `resource_type` - (String) The resource type.
  - `href` - (String) The URL for this dynamic route server peer group peer.
  - `id` - (String) The unique identifier for this dynamic route server peer group peer.
  - `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
    Nested schema for **lifecycle_reasons**:
    - `code` - (String) A reason code for this state.
    - `message` - (String) An explanation of the reason for this state.
    - `more_info` - (String) A link to documentation about the reason for this state.
  - `lifecycle_state` - (String) The lifecycle state of the dynamic route server peer group peer.
  - `name` - (String) The name for this dynamic route server peer group peer.
  - `priority` - (Integer) The priority of this dynamic route server peer group peer.
  - `resource_type` - (String) The resource type.
  - `sessions` - (List) The BGP sessions for this dynamic route server peer group peer.
    Nested schema for **sessions**:
    - `established_at` - (String) The date and time that this BGP session was established.
    - `local` - (List) The dynamic route server member for the local side of this session. Same nested schema as `bidirectional_forwarding_detection.sessions.local`.
    - `protocol_state` - (String) The BGP protocol state of this session.
    - `remote` - (List) The remote side of this session.
      Nested schema for **remote**:
      - `name` - (String) The name of the remote peer.
  - `state` - (String) The state of this dynamic route server peer group peer: `enabled` or `disabled`.
  - `status` - (String) The status of this dynamic route server peer group peer connection: `down` or `up`.
  - `status_reasons` - (List) The status reasons for the dynamic route server peer group peer.
    Nested schema for **status_reasons**:
    - `code` - (String) A reason code for this state.
    - `message` - (String) An explanation of the reason for this state.
    - `more_info` - (String) A link to documentation about the reason for this state.
