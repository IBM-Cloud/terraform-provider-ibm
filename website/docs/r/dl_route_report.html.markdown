---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_route_report"
description: |-
  Create and Delete Route Report for a DirectLink Gateway.
---

# ibm_dl_route_report

Provides a resource for ibm_dl_route_report. This allows to create and delete a route report for a directlink gateway. For more information, see [about Direct Link Route Report](https://cloud.ibm.com/docs/dl?topic=dl-generate-route-reports&interface=ui).


## Example usage to create Direct Link Route Report on dedicated gateway
In the following example, you can create Direct Link Route Report for dedicated gateway:

```terraform
data "ibm_dl_routers" "test_dl_routers" {
		offering_type = "dedicated"
		location_name = "dal10"
	}

resource ibm_dl_gateway test_dl_gateway {
  bgp_asn =  64999
  global = true 
  metered = false
  name = "Gateway1"
  speed_mbps = 1000 
  type =  "dedicated" 
  cross_connect_router = data.ibm_dl_routers.test_dl_routers.cross_connect_routers[0].router_name
  location_name = data.ibm_dl_routers.test_dl_routers.location_name
  customer_name = "Customer1" 
  carrier_name = "Carrier1"

} 

resource ibm_dl_route_report dl_route_report {
    gateway = ibm_dl_gateway.test_dl_gateway.id
}
```

## Sample usage to create Direct Link Route Report on connect gateway
In the following example, you can create Direct Link Route Report on connect gateway:


```terraform
data "ibm_dl_ports" "test_ds_dl_ports" {
 
 }
resource "ibm_dl_gateway" "test_dl_connect" {
  bgp_asn =  64999
  global = true
  metered = false
  name = "dl-connect-gw-1"
  speed_mbps = 1000
  type =  "connect"
  port =  data.ibm_dl_ports.test_ds_dl_ports.ports[0].port_id
}

resource ibm_dl_route_report dl_route_report {
    gateway = ibm_dl_gateway.test_dl_gateway.id
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `gateway`- (Required, Forces New, Integer) Direct Link Gateway ID.

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your resource is created.
- `created_at` - (String) The date and time resource created.
- `advertised_routes` - (List) List of connection prefixes advertised to the on-prem network.
    Nested scheme for `advertised_routes`:
    - `as_path` - (String) The BGP AS path of the route.
    - `prefix` - (String) The prefix used in the route.
- `gateway_routes` - (List) List of local/direct routes.
    Nested scheme for `gateway_routes`:
    - `prefix` - (String) The prefix used in the route.
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
- `route_report_id` - (String) The unique identifier of this route report.
- `status` - (String) The route report status.
- `updated_at` - (String) The date and time resource was updated.
- `virtual_connection_routes` - (List) List of routes on virtual connections.
    Nested scheme for `virtual_connection_routes`
    - `routes` - (List) List of connection routes.
        Nested scheme for `routes`:
        - `active` - (Bool) Indicates whether the route is the preferred path of the prefix.
        - `local_preference` - (String) The local preference of the route. This attribute can manipulate the chosen path on routes.
        - `prefix` - (String) The prefix used in the route.
    - `virtual_connection_id` - (String) Virtual Connection ID
    - `virtual_connection_name` - (String) Virtual Connection name
    - `virtual_connection_type` - (String) Virtual Connection type

## Import

You can import the `ibm_dl_route_report` resource by using `id`.
The `id` property can be formed from `gateway` and `route_report_id` in the following format:

```
<gateway>/<route_report_id>
```
* `gateway`: A String. The unique identifier of a directlink gateway.
* `route_report_id`: A String. The unique identifier of the route report.

```
$ terraform import ibm_dl_route_report.dl_route_report <gateway>/<route_report_id>
```
