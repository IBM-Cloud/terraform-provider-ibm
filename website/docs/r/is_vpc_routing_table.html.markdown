---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc-routing-tables"
description: |-
  Manages IBM IS VPC routing tables.
---

# ibm_is_vpc_routing_table
Create, update, or delete an VPC routing tables. For more information, about VPC routes, see [routing tables for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).

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
  name                          = "example-vpc-routing-table"
  route_direct_link_ingress     = true
  route_transit_gateway_ingress = false
  route_vpc_zone_ingress        = false
}

```

## Example usage: Advertising routes
```
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}
resource "ibm_is_vpc_routing_table" "is_vpc_routing_table_instance" {
  vpc                           = ibm_is_vpc.example.id
  name                          = "example-vpc-routing-table"
  route_direct_link_ingress     = true
  route_transit_gateway_ingress = false
  route_vpc_zone_ingress        = false
  advertise_routes_to           = ["direct_link", "transit_gateway"]

}
```
# Example usage for accept_routes_from_resource_type
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_routing_table" "example" {
  vpc                              = ibm_is_vpc.example.id
  name                             = "example-vpc-routing-table"
  route_direct_link_ingress        = true
  route_transit_gateway_ingress    = false
  route_vpc_zone_ingress           = false
  accept_routes_from_resource_type = ["vpn_server"]
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the routing table.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `advertise_routes_to` - (Optional, List) The ingress sources to advertise routes to. Routes in the table with `advertise` enabled will be advertised to these sources.

  ->**Options** An ingress source that routes can be advertised to:</br>
        **&#x2022;** `direct_link` (requires `route_direct_link_ingress` be set to `true`)</br>
        **&#x2022;** `transit_gateway` (requires `route_transit_gateway_ingress` be set to `true`)
- `accept_routes_from_resource_type` - (Optional, List) The resource type filter specifying the resources that may create routes in this routing table. Ex: `vpn_server`, `vpn_gateway`
- `created_at` - (Timestamp)  The date and time when the routing table was created.
- `name` - (Optional, String) The routing table name.
- `route_direct_link_ingress` - (Optional, Bool)  If set to **true**, the routing table is used to route traffic that originates from Direct Link to the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**.
- `route_internet_ingress` - (Optional, Bool) If set to **true**, this routing table will be used to route traffic that originates from the internet. For this to succeed, the VPC must not already have a routing table with this property set to **true**.
- `route_transit_gateway_ingress` - (Optional, Bool) If set to **true**, the routing table is used to route traffic that originates from Transit Gateway to the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**.
- `route_vpc_zone_ingress` - (Optional, Bool) If set to true, the routing table is used to route traffic that originates from subnets in other zones in the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**.
- `tags` - (Optional, Array of Strings) Enter any tags that you want to associate with your routing table. Tags might help you find your routing table more easily after it is created. Separate multiple tags with a comma (`,`).
- `vpc` - (Required, Forces new resource, String) The VPC ID. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn`-  (String) CRN of the default routing table.
- `href` - (String) The routing table URL.
- `id` - (String) The unique identifier of the routing table. The ID is composed of `<vpc_id>/<vpc_routing_table_id>`.
- `is_default` - (String)  Indicates the default routing table for this VPC.
- `lifecycle_state` - (String) The lifecycle state of the routing table.
- `resource_type` - (String) The resource type.
- `resource_group` - (List) The resource group for this routing table. 

  Nested scheme for `resource_group`:
  - `href` - (String) The URL for this resource group.
  - `id` - (String) The unique identifier for this resource group.
  - `name` - (String) The name for this resource group. 
- `routing_table` - (String) The unique routing table identifier.
- `routes` - (List) The routes for the routing table.

  Nested scheme for `routes`:
  - `id` - (String) The unique ID of the route.
  - `name`-  (String) The user-defined name of the route.
- `subnets` - (List) The subnets to which routing table is attached.

  Nested scheme for `subnets`:
  - `id` - (String) The unique ID of the subnet.
  - `name` - (String) The user defined name of the subnet.

## Import
The `ibm_is_vpc_routing_table` resource can be imported by using VPC ID and VPC Route table ID.

**Example**

```
$ terraform import ibm_is_vpc_routing_table.example 56738c92-4631-4eb5-8938-8af9211a6ea4/fc2667e0-9e6f-4993-a0fd-cabab477c4d1
```
