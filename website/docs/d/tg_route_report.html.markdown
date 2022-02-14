---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_route_report"
description: |-
  Manages IBM Cloud Infrastructure Transit Gateway Route Report.
---

# ibm_tg_route_report
Retrieve information of an existing IBM Cloud infrastructure transit gateway route report as a read only data source. For more information about Transit Gateway Route Reports, see [generating and viewing a route report](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-route-reports&interface=ui#generate-route-report-ui).

## Example usage

```terraform
data "ibm_tg_route_report" "tg_route_report" {
    gateway = ibm_tg_gateway.new_tg_gw.
    route_report = ibm_tg_route_report_test_tg_route_report.route_report_id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `gateway` - (Required, String) The unique identifier of the gateway.
- `route_report` - (Required, String) The unique identifier of the gateway route report


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `created_at` - (Timestamp) The date and time resource is created.
- `id` - (String) The unique identifier of this route report.
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
