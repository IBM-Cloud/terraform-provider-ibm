---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_peer_group"
description: |-
  Get information about an IBM dynamic route server peer group.
---

# ibm_is_dynamic_route_server_peer_group

Provides a read-only data source to retrieve information about a dynamic route server peer group. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc).

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
data "ibm_is_dynamic_route_server_peer_group" "by_id" {
  dynamic_route_server_id            = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
}

data "ibm_is_dynamic_route_server_peer_group" "by_name" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
  name                    = "example-peer-group"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `dynamic_route_server_id` - (Required, String) The dynamic route server identifier.
- `dynamic_route_server_peer_group_id` - (Optional, String) The dynamic route server peer group identifier. Exactly one of `dynamic_route_server_peer_group_id` and `name` must be specified.
- `name` - (Optional, String) The name of the dynamic route server peer group. Exactly one of `dynamic_route_server_peer_group_id` and `name` must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the data source in the format `<dynamic_route_server_id>/<dynamic_route_server_peer_group_id>`.
- `created_at` - (String) The date and time that this dynamic route server peer group was created.
- `creator` - (List of String) The types of resource that created peers in this peer group.
- `health_monitor` - (List) The health monitoring configuration used for this dynamic route server peer group.
  Nested schema for **health_monitor**:
  - `mode` - (String) The mode used for this health monitor: `bfd`, `bgp` or `none`.
  - `roles` - (List of String) The roles used for this dynamic route server peer group health monitor configuration.
- `href` - (String) The URL for this dynamic route server peer group.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  Nested schema for **lifecycle_reasons**:
  - `code` - (String) A reason code for this state.
  - `message` - (String) An explanation of the reason for this state.
  - `more_info` - (String) A link to documentation about the reason for this state.
- `lifecycle_state` - (String) The lifecycle state of the dynamic route server peer group.
- `peers` - (List) The peers in this dynamic route server peer group. Each peer has the same nested schema as the `peers` attribute of the `ibm_is_dynamic_route_server_peer_group_peers` data source.
- `prefix_watcher` - (List) The prefix watcher configuration for this dynamic route server peer group.
  Nested schema for **prefix_watcher**:
  - `excluded_prefixes` - (List) The prefixes excluded from monitoring.
    Nested schema for **excluded_prefixes**:
    - `ge` - (Integer) The minimum prefix length to match.
    - `le` - (Integer) The maximum prefix length to match.
    - `prefix` - (String) The IP prefix in CIDR format.
  - `monitored_prefixes` - (List) The prefixes monitored by this peer group.
    Nested schema for **monitored_prefixes**:
    - `ge` - (Integer) The minimum prefix length to match.
    - `le` - (Integer) The maximum prefix length to match.
    - `prefix` - (String) The IP prefix in CIDR format.
- `resource_type` - (String) The resource type.
- `roles` - (List of String) The roles for this dynamic route server peer group.
- `state` - (String) The state of this dynamic route server peer group: `enabled` or `disabled`.
- `status` - (String) The status of this dynamic route server peer group.
- `status_reasons` - (List) The status reasons for the dynamic route server peer group.
  Nested schema for **status_reasons**:
  - `code` - (String) A reason code for this state.
  - `message` - (String) An explanation of the reason for this state.
  - `more_info` - (String) A link to documentation about the reason for this state.
