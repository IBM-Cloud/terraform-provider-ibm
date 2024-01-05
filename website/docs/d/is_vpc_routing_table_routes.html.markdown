---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Routing Table Routes"
description: |-
  Get information about IBM VPC routing table routes.
---

# ibm_is_vpc_routing_table_routes
Retrieve information of an existing IBM Cloud Infrastructure Virtual Private Cloud routing table routes as a read-only data source. For more information, about VPC default routing table, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).

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
  name = "example-routing-table"
  vpc  = ibm_is_vpc.example.id
}


data "ibm_is_vpc_routing_table_routes" "example" {
  vpc           = ibm_is_vpc.example.id
  routing_table = ibm_is_vpc_routing_tables.example.routing_table
}
```
## Argument reference
Review the argument references that you can specify for your data source. 

- `vpc` - (Required, String) The ID of the VPC.
- `routing_table` - (Required, String) The ID of the routing table.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `routes` (List) List of all the routing table in a VPC.

  Nested scheme for `routes`:
	- `name` - (String) The name for the default routing table.
	- `route_id` - (String) The unique ID for the route.
	- `lifecycle_state` - (String) The lifecycle state of the route.
	- `href` - (String) The routing table URL.
	- `created_at` - (Timestamp)  The date and time that the route was created.
	- `creator` - (List) If present, the resource that created the route. Routes with this property present cannot bedirectly deleted. All routes with an `origin` of `learned` or `service` will have thisproperty set, and future `origin` values may also have this property set.
		Nested scheme for **creator**:
		- `crn` - (String) The VPN gateway's CRN.
		- Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		Nested scheme for **deleted**:
			- `more_info` - (Required, String) Link to documentation about deleted resources.
			- Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `href` - (String) The VPN gateway's canonical URL.
		- Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `id` - (String) The unique identifier for this VPN gateway.
		- Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		- `name` - (String) The user-defined name for this VPN gateway.
		- Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		- `resource_type` - (String) The resource type.
		- Constraints: Allowable values are: `vpn_gateway`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	- `action` - (String) The action to perform with a packet matching the route.
	- `advertise` - (Boolean) Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.
	- `destination` - (String) The destination of the route.
	- `next_hop` - (String) The next hop address of the route.
	- `origin` - (String) The origin of this route:- `service`: route was directly created by a service - `user`: route was directly created by a userThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the route on which the unexpected property value was encountered.
  		- Constraints: Allowable values are: `learned`, `service`, `user`.
	- `priority` - (Integer) The route's priority. Smaller values have higher priority. If a routing table contains routes with the same destination, the route with the highest priority (smallest value) is selected. For Example (2), supports values from 0 to 4. Default is 2.
	- `zone` - (String) The zone name of the route.
