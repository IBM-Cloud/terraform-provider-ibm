---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : dynamic_route_server_members"
description: |-
  List IBM dynamic route server members.
---

# ibm_is_dynamic_route_server_members

Retrieve the list of members of a dynamic route server. All pages of the collection are read. For more information, see [dynamic route servers](https://cloud.ibm.com/docs/vpc?topic=vpc-about-drs).

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
data "ibm_is_dynamic_route_server_members" "example" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `dynamic_route_server_id` - (Required, String) The dynamic route server identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - (String) The unique identifier of the data source.
- `members` - (List) The members of the dynamic route server.

  Nested schema for **members**:
  - `id` - (String) The unique identifier for this dynamic route server member.
  - `name` - (String) The name for this dynamic route server member.
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
