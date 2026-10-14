---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : dynamic_route_servers"
description: |-
  List IBM dynamic route servers.
---

# ibm_is_dynamic_route_servers

Retrieve the list of dynamic route servers in the region, optionally filtered by name or resource group. All pages of the collection are read. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc?topic=vpc-about-drs).

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
data "ibm_is_dynamic_route_servers" "example" {
}
```

```terraform
data "ibm_is_dynamic_route_servers" "example" {
  name              = "example-dynamic-route-server"
  resource_group_id = "11caaa983d9c4beb82690daab18717e9"
  sort              = "name"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `name` - (Optional, String) Filters the collection to resources with a `name` property matching the exact specified name.
- `resource_group_id` - (Optional, String) Filters the collection to resources with a `resource_group.id` property matching the specified identifier.
- `sort` - (Optional, String) Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. Default is `-created_at`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - (String) The unique identifier of the data source.
- `dynamic_route_servers` - (List) The dynamic route servers.

  Nested schema for **dynamic_route_servers**:
  - `id` - (String) The unique identifier for this dynamic route server.
  - `name` - (String) The name for this dynamic route server.
  - `created_at` - (String) The date and time that the dynamic route server was created.
  - `crn` - (String) The CRN for this dynamic route server.
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
  - `local_asn` - (Integer) The local autonomous system number (ASN) used for this dynamic route server and its connections.
  - `members` - (List) The members for the dynamic route server.

    Nested schema for **members**:
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

      Nested schema for **deleted**:
      - `more_info` - (String) A link to documentation about deleted resources.
    - `href` - (String) The URL for this dynamic route server member.
    - `id` - (String) The unique identifier for this dynamic route server member.
    - `name` - (String) The name for this dynamic route server member.
    - `resource_type` - (String) The resource type.
    - `virtual_network_interfaces` - (List) The virtual network interfaces for this member.

      Nested schema for **virtual_network_interfaces**:
      - `crn` - (String) The CRN for this virtual network interface.
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

        Nested schema for **deleted**:
        - `more_info` - (String) A link to documentation about deleted resources.
      - `href` - (String) The URL for this virtual network interface.
      - `id` - (String) The unique identifier for this virtual network interface.
      - `name` - (String) The name for this virtual network interface.
      - `primary_ip` - (List) The primary IP for this virtual network interface.

        Nested schema for **primary_ip**:
        - `address` - (String) The IP address.
        - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

          Nested schema for **deleted**:
          - `more_info` - (String) A link to documentation about deleted resources.
        - `href` - (String) The URL for this reserved IP.
        - `id` - (String) The unique identifier for this reserved IP.
        - `name` - (String) The name for this reserved IP.
        - `resource_type` - (String) The resource type.
      - `resource_type` - (String) The resource type.
      - `subnet` - (List) The associated subnet.

        Nested schema for **subnet**:
        - `crn` - (String) The CRN for this subnet.
        - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

          Nested schema for **deleted**:
          - `more_info` - (String) A link to documentation about deleted resources.
        - `href` - (String) The URL for this subnet.
        - `id` - (String) The unique identifier for this subnet.
        - `name` - (String) The name for this subnet.
        - `resource_type` - (String) The resource type.
  - `resource_group` - (List) The resource group for this dynamic route server.

    Nested schema for **resource_group**:
    - `href` - (String) The URL for this resource group.
    - `id` - (String) The unique identifier for this resource group.
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
