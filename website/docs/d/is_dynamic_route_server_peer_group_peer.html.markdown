---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_peer_group_peer"
description: |-
  Get information about an IBM dynamic route server peer group peer.
---

# ibm_is_dynamic_route_server_peer_group_peer

Provides a read-only data source to retrieve information about a dynamic route server peer group peer. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc).

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
data "ibm_is_dynamic_route_server_peer_group_peer" "by_id" {
  dynamic_route_server_id                 = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id      = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
  dynamic_route_server_peer_group_peer_id = ibm_is_dynamic_route_server_peer_group_peer.example.dynamic_route_server_peer_group_peer_id
}

data "ibm_is_dynamic_route_server_peer_group_peer" "by_name" {
  dynamic_route_server_id            = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
  name                               = "example-peer"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `dynamic_route_server_id` - (Required, String) The dynamic route server identifier.
- `dynamic_route_server_peer_group_id` - (Required, String) The dynamic route server peer group identifier.
- `dynamic_route_server_peer_group_peer_id` - (Optional, String) The dynamic route server peer group peer identifier. Exactly one of `dynamic_route_server_peer_group_peer_id` and `name` must be specified.
- `name` - (Optional, String) The name of the dynamic route server peer group peer. Exactly one of `dynamic_route_server_peer_group_peer_id` and `name` must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the data source in the format `<dynamic_route_server_id>/<dynamic_route_server_peer_group_id>/<dynamic_route_server_peer_group_peer_id>`.
- `asn` - (Integer) The autonomous system number (ASN) for this dynamic route server peer group peer.
- `bidirectional_forwarding_detection` - (List) The bidirectional forwarding detection (BFD) configuration for this dynamic route server peer group peer. Same nested schema as the `peers.bidirectional_forwarding_detection` attribute of the `ibm_is_dynamic_route_server_peer_group_peers` data source.
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
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  Nested schema for **lifecycle_reasons**:
  - `code` - (String) A reason code for this state.
  - `message` - (String) An explanation of the reason for this state.
  - `more_info` - (String) A link to documentation about the reason for this state.
- `lifecycle_state` - (String) The lifecycle state of the dynamic route server peer group peer.
- `priority` - (Integer) The priority of this dynamic route server peer group peer.
- `resource_type` - (String) The resource type.
- `sessions` - (List) The BGP sessions for this dynamic route server peer group peer. Same nested schema as the `peers.sessions` attribute of the `ibm_is_dynamic_route_server_peer_group_peers` data source.
- `state` - (String) The state of this dynamic route server peer group peer: `enabled` or `disabled`.
- `status` - (String) The status of this dynamic route server peer group peer connection: `down` or `up`.
- `status_reasons` - (List) The status reasons for the dynamic route server peer group peer.
  Nested schema for **status_reasons**:
  - `code` - (String) A reason code for this state.
  - `message` - (String) An explanation of the reason for this state.
  - `more_info` - (String) A link to documentation about the reason for this state.
