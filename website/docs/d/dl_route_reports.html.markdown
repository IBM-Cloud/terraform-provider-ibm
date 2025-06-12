---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_route_reports"
description: |-
  Retrieve all route reports for the specified Direct Link gateway.
---

# ibm_dl_router_reports

Import the details of an existing infrastructure Direct Link Route Reports as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about Direct Link route reports, see [Direct Link Route Report](https://cloud.ibm.com/docs/dl?topic=dl-generate-route-reports&interface=ui).


## Example usage

```terraform
data "ibm_dl_route_reports" "test_dl_reports" {
	gateway = ibm_dl_gateway.test_dl_gateway.id
}
```

## Argument reference
The argument reference that you need to specify for the data source. 

- `gateway`- (Required, String) Direct Link Gateway ID.

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your data source is created.
- `route_reports` - (String) List of all route reports for the transit gateway
    Nested scheme for `route_reports`
    - `created_at` - (String) The date and time resource created.
    - `advertised_routes` - (List) List of connection prefixes advertised to the on-prem network.
        Nested scheme for `advertised_routes`:
        - `as_path` - (String) The BGP AS path of the route.
        - `prefix` - (String) The prefix used in the route.
    - `gateway_routes` - (List) List of local/direct routes.
        Nested scheme for `gateway_routes`:
        - `prefix` - (String) The prefix used in the route.
    - `id` - (String) Route report identifier.
    - `on_prem_routes` - (List) List of on premises routes
        Nested scheme for `on_prem_routes`:
        - `as_path` - (String) The BGP AS path of the route.
        - `next_hop` - (String) Next hop address.
        - `prefix` - (String) The prefix used in the route.
    - `overlapping_routes` - (List) List of overlapping routes.
        Nested scheme for `overlapping_routes`:
        - `routes` - (List) List of overlapping connection/prefix pairs.
            Nested scheme for `routes`:
            - `prefix` - (String) The overlapping prefix.
            - `type` - (String) The type of route.
            - `virtual_connection_id` - (String) Virtual Connection ID. This is set only when type of route is virtual_connection.
    - `status` - (String) The route report status.
    - `updated_at` - (String) The date and time resource was updated.
    - `virtual_connection_routes` - (List) List of routes on virtual connections.
        Nested scheme for `virtual_connection_routes`
        - `routes` - (List) List of virtual connection routes.
            Nested scheme for `routes`:
            - `active` - (Bool) Indicates whether the route is the preferred path of the prefix.
            - `local_preference` - (String) The local preference of the route. This attribute can manipulate the chosen path on routes.
            - `prefix` - (String) The prefix used in the route.
        - `virtual_connection_id` - (String) Virtual Connection ID
        - `virtual_connection_name` - (String) Virtual Connection name
        - `virtual_connection_type` - (String) Virtual Connection type
