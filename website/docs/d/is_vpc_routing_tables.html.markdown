---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Routing Tables"
description: |-
  Get information about IBM VPC routing tables.
---

# ibm_is_vpc_routing_tables
Retrieve information of an existing IBM Cloud infrastructure VPC default routing tables. For more information, about VPC routing tables, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes)

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

data "ibm_is_vpc_routing_tables" "example" {
  vpc = ibm_is_vpc.example.id
}
```


## Argument reference
Review the argument references that you can specify for your data source. 

- `vpc` - (Required, String) The ID of the VPC.
- `is_default` - (Optional, Boolean) Indicate whether this is the default routing table for this VPC

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `routing_tables` (List) List of all the routing tables in a VPC.

  Nested scheme for `routing_tables`:
    - `access_tags`  - (List) Access management tags associated for the routing table.
	- `accept_routes_from` - (List) The filters specifying the resources that may create routes in this routing table.At present, only the `resource_type` filter is permitted, and only the `vpn_gateway` value is supported, but filter support is expected to expand in the future.
		Nested scheme for **accept_routes_from**:
		- `resource_type` - (String) The resource type.		
	- `advertise_routes_to` - (Optional, List) The ingress sources to advertise routes to. Routes in the table with `advertise` enabled will be advertised to these sources.

		->**Options** An ingress source that routes can be advertised to:</br>
				**&#x2022;** `direct_link` (requires `route_direct_link_ingress` be set to `true`)</br>
				**&#x2022;** `transit_gateway` (requires `route_transit_gateway_ingress` be set to `true`)
    - `created_at` - (Timestamp)  The date and time the routing table was created.
	- `crn` -  (String) The crn for this routing table.
	- `href` - (String) The routing table URL.
	- `is_default` - (String)  Indicates whether the default routing table.
	- `lifecycle_state` - (String) The lifecycle state of the routing table.
	- `name` - (String) The name for the default routing tables.
	- `resource_type` - (String) The type of resource referenced.
	- `resource_group` - (List) The resource group for this routing table.
		Nested scheme for `resource_group`:
		- `href` - (String) The URL for this resource group.
		- `id` - (String) The unique identifier for this resource group.
		- `name` - (String) The name for this resource group.
	- `route_table` - (String) The unique ID for the routing table.
	- `route_direct_link_ingress` - (String) Indicates if the routing table is used to route traffic that originates from Direct Link to the VPC.
	- `route_internet_ingress` - (Boolean) Indicates whether this routing table is used to route traffic that originates from the internet.
	- `route_transit_gateway_ingress` - (String) Indicates if the routing table is used to route traffic that originates from Transit Gateway to the VPC.
	- `route_vpc_zone_ingress` - (String)  Indicates if the routing table is used to route traffic that originates from subnets in other zones of the VPC.
	- `routes` - (List) The routes for the routing table.	
		Nested scheme for `routes`:
		- `id` - (String) The unique ID of the route.
		- `name`-  (String) The user-defined name of the route.
	- `subnets` - (List) The subnets to which routing table is attached.    
		Nested scheme for `subnets`:
		- `id` - (String) The unique ID of the subnet.
		- `name` - (String) The user-defined name of the subnet.
	- `tags` - (List) Tags associated with the routing table.
