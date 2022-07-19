---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc-routing-tables-route"
description: |-
  Manages IBM IS VPC routing tables.
---

# ibm_is_vpc_routing_table_route
Create, update, or delete of an VPC routing tables. For more information, about VPC routes, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}
resource "ibm_is_vpc_routing_table" "example" {
  vpc                           = ibm_is_vpc.example.id
  name                          = "example-routing-table"
  route_direct_link_ingress     = true
  route_transit_gateway_ingress = false
  route_vpc_zone_ingress        = false
}
resource "ibm_is_vpc_routing_table_route" "example" {
  vpc           = ibm_is_vpc.example.id
  routing_table = ibm_is_vpc_routing_table.example.routing_table
  zone          = "us-south-1"
  name          = "custom-route-2"
  destination   = "192.168.4.0/24"
  action        = "deliver"
  next_hop      = ibm_is_vpn_gateway_connection.example.gateway_connection // Example value "10.0.0.4"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `action` - (Optional, String) The action to perform with a packet matching the route `delegate`, `delegate_vpc`, `deliver`, `drop`.
- `destination` - (Required, Forces new resource, String) The destination of the route. 
- `name` - (Optional, String) The user-defined name of the route. If unspecified, the name will be a hyphenated list of randomly selected words. You need to provide unique name within the VPC routing table the route resides in.
- `next_hop` - (Required, Forces new resource, String) The next hop of the route. It accepts IP address or a VPN connection ID. For `action` other than `deliver`, you must specify `0.0.0.0`. 
- `routing_table` - (Required, String) The routing table ID.
- `vpc` - (Required, Forces new resource, String) The VPC ID.
- `zone` - (Required, Forces new resource, String)  Name of the zone. 


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `creator` - (Optional, List) If present, the resource that created the route. Routes with this property present cannot bedirectly deleted. All routes with an `origin` of `learned` or `service` will have thisproperty set, and future `origin` values may also have this property set.
  Nested scheme for **creator**:
    - `crn` - (Optional, String) The VPN gateway's CRN.
      - Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
    - `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
    Nested scheme for **deleted**:
        - `more_info` - (Required, String) Link to documentation about deleted resources.
          - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
    - `href` - (Optional, String) The VPN gateway's canonical URL.
      - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
    - `id` - (Optional, String) The unique identifier for this VPN gateway.
      - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
    - `name` - (Optional, String) The user-defined name for this VPN gateway.
      - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
    - `resource_type` - (Optional, String) The resource type.
      - Constraints: Allowable values are: `vpn_gateway`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
- `href` - (String) The routing table URL.
- `id` - (String) The routing table ID. The ID is composed of `<vpc_route_table_id>/<vpc_route_table_route_id>`.
- `is_default` - (String) Indicates the default routing table for this VPC.
- `lifecycle_state` - (String) The lifecycle state of the route.
- `origin` - (Optional, String) The origin of this route:- `service`: route was directly created by a service- `user`: route was directly created by a userThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the route on which the unexpected property value was encountered.
  - Constraints: Allowable values are: `learned`, `service`, `user`.
- `resource_type` - (String) The resource type.

## Import
The `ibm_is_vpc_routing_table_route` resource  can be imported by using VPC ID, VPC Route table ID, and VPC Route table Route ID.

**Example**

```
$ terraform import ibm_is_vpc_routing_table_route.example 56738c92-4631-4eb5-8938-8af90000006ea4/4993-a0fd-cabab477c4d1-8af911111a4/fc2667e0-9e6f-4993-a0fd-cabab55557c4d1
```

