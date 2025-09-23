---
layout: "ibm"
page_title: "IBM : ibm_pi_network_peer_route_filter"
description: |-
  Manages network peer route filter in Power Virtual Server cloud.
subcategory: "Power Systems"
---

# ibm_pi_network_peer_route_filter

Create and delete a network peer route filter with this resource for on-prem locations.

## Example Usage

The following example creates a network peer route filter for your Power Systems Virtual Server.

```terraform
resource "ibm_pi_network_peer_route_filter" "pi_network_peer_route_filter_instance" {
  pi_cloud_instance_id = "6f8e2a9d-3b4c-4e4f-8e8d-f7e7e1e23456"
  pi_direction = "import"
  pi_ge = 25
  pi_index = 10
  pi_le = 30
  pi_network_peer_id = "7e1c3b2a-9f0d-4e5f-a1bc-def012345678"
  pi_prefix = "192.168.91.0/24"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Timeouts

The `ibm_pi_network_peer_route_filter` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 2 minutes) Used for creating a network peer route filter.
- **delete** - (Default 2 minutes) Used for deleting a network peer route filter.
  
## Argument Reference

Review the argument references that you can specify for this resource.

- `pi_action` - (Optional, String) Action of the filter. Allowable values are: `allow`, `deny`.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_direction` - (Required, String) Direction of the filter. Allowable values are: `import`, `export`.
- `pi_ge` - (Optional, Integer) The minimum matching length of the prefix-set(1 ≤ value ≤ 32 & value ≤ LE).
- `pi_index` - (Required, Integer) Priority or order of the filter.
- `pi_le` - (Optional, Integer) The maximum matching length of the prefix-set( 1 ≤ value ≤ 32 & value >= GE).
- `pi_network_peer_id` - (Required, String) Network peer ID.
- `pi_prefix` - (Required, String) IP prefix representing an address and mask length of the prefix-set.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `creation_date` - (String) Time stamp for create route filter.
- `error` - (String) Error description.
- `id` - (String) The unique identifier of the network peer route filter resource. Composed of `<pi_cloud_instance_id>/<pi_network_peer_id>/<route_filter_id>`
- `route_filter_id` - (String) The unique identifier of the network peer route filter.
- `state` - (String) Status of the route filter.

## Import

The `ibm_pi_network_peer_route_filter` resource can be imported by using `pi_cloud_instance_id`, `pi_network_peer_id` and `route_filter_id`.

### Example

```bash
terraform import ibm_pi_network_peer_route_filter.pi_network_peer_route_filter 6f8e2a9d-3b4c-4e4f-8e8d-f7e7e1e23456/7e1c3b2a-9f0d-4e5f-a1bc-def012345678/8a9b1c2d-3e4f-5g6h-7i8j-9k0l1m2n3o4p
```
