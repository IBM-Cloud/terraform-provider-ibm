---
subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_route_report"
description: |-
  Manages IBM Transit Gateway Route Report.
---

# ibm_tg_route_report
Create and delete a transit gateway's route report resource. For more information about Transit Gateway Route Reports, see [generating and viewing a route report](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-route-reports&interface=ui#generate-route-report-ui).

## Example usage

```terraform
resource ibm_tg_route_report" "test_tg_route_report" {
    gateway = ibm_tg_gateway.new_tg_gw.id
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `gateway` - (Required, String) The unique identifier of the gateway.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `route_report_id` - (String) The unique identifier of this route report.
- `created_at` - (Timestamp) The date and time resource is created.
- `status` - (String) The route report status.
- `updated_at` - (Timestamp) The date and time resource is last updated.
- `connections` - (String) A list of connections in the gateway

    Nested scheme for `connections`:
    - `bgps` (String) A list of the connection's bgps
        Nested scheme for `bgps`:
        - `as_path` - (String) The bgp AS path
        - `is_used` - (Bool) Indicates whether the current route is used or not
        - `local_preference` (String) The local preference
        - `prefix` - (String) The bgp prefix
    - `id` - (String) The unique identifier for the transit gateway connection
    - `name` - (String) The user-defined name for the transit gateway connection.
    - `routes` - (String) A list of the connection's routes

        Nested scheme for `routes`:
        - `prefix` - (String) The prefix used in the route
    - `type` - (String) The connection type
- `overlapping_routes` - (String) A list of overlapping routes in the gateway

    Nested scheme for `overlapping_routes`:
    - `connection_id` - (String) The unique identifier for the transit gateway connection
    - `prefix` - (String) The overlapping prefix