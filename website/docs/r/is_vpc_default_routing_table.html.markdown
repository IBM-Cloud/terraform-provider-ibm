---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc_default_routing_table"
description: |-
  Manages IBM IS VPC default routing table.
---

# ibm_is_vpc_default_routing_table

Provides a resource to manage the default routing table of a VPC. This resource allows you to manage the default routing table that is automatically created when a VPC is created. Unlike custom routing tables, the default routing table cannot be deleted but can be configured to control traffic routing behavior.

~> **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ibm_is_vpc_default_routing_table` resource behaves differently from normal resources. Terraform does not _create_ this resource but instead attempts to "adopt" it into management.

Every VPC has a default routing table that can be managed but not destroyed. When Terraform first adopts a default routing table, it manages the configuration of the existing default routing table. This resource allows you to configure ingress routing behavior, route acceptance filters, and route advertisement settings for the default routing table.

For more information, about VPC, see [getting started with Virtual Private Cloud](https://cloud.ibm.com/docs/vpc?topic=vpc-getting-started). For more information, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes&interface=ui).

For more information, about VPC routes, see [routing tables for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Basic usage
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_routing_table" "example" {
  vpc = ibm_is_vpc.example.id
}
```

### Configure ingress routing
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_routing_table" "example" {
  vpc                           = ibm_is_vpc.example.id
  route_direct_link_ingress     = true
  route_transit_gateway_ingress = false
  route_vpc_zone_ingress        = false
}
```

### Example usage with route acceptance filters
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_routing_table" "example" {
  vpc                              = ibm_is_vpc.example.id
  route_direct_link_ingress        = true
  route_transit_gateway_ingress    = false
  route_vpc_zone_ingress           = false
  accept_routes_from_resource_type = ["vpn_server", "vpn_gateway"]
}
```

### Example usage with route advertisement
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_routing_table" "example" {
  vpc                           = ibm_is_vpc.example.id
  route_direct_link_ingress     = true
  route_transit_gateway_ingress = true
  route_vpc_zone_ingress        = false
  advertise_routes_to           = ["direct_link", "transit_gateway"]
}
```

### Example usage with tags
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_default_routing_table" "example" {
  vpc                           = ibm_is_vpc.example.id
  route_direct_link_ingress     = true
  route_transit_gateway_ingress = false
  route_vpc_zone_ingress        = false
  tags                          = ["env:production", "team:networking"]
  access_tags                   = ["project:web-app"]
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `accept_routes_from_resource_type` - (Optional, List) The resource type filter specifying the resources that may create routes in this routing table. Supported values: `vpn_server`, `vpn_gateway`.
- `access_tags` - (Optional, List of Strings) A list of access management tags to attach to the default routing table.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `advertise_routes_to` - (Optional, List) The ingress sources to advertise routes to. Routes in the table with `advertise` enabled will be advertised to these sources.

  ->**Options** An ingress source that routes can be advertised to:</br>
        **&#x2022;** `direct_link` (requires `route_direct_link_ingress` be set to `true`)</br>
        **&#x2022;** `transit_gateway` (requires `route_transit_gateway_ingress` be set to `true`)
- `route_direct_link_ingress` - (Optional, Bool) If set to **true**, the routing table is used to route traffic that originates from Direct Link to the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**. Default: `false`.
- `route_internet_ingress` - (Optional, Bool) If set to **true**, this routing table will be used to route traffic that originates from the internet. For this to succeed, the VPC must not already have a routing table with this property set to **true**. Default: `false`.
- `route_transit_gateway_ingress` - (Optional, Bool) If set to **true**, the routing table is used to route traffic that originates from Transit Gateway to the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**. Default: `false`.
- `route_vpc_zone_ingress` - (Optional, Bool) If set to true, the routing table is used to route traffic that originates from subnets in other zones in the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**. Default: `false`.
- `tags` - (Optional, Array of Strings) Enter any tags that you want to associate with your default routing table. Tags might help you find your routing table more easily after it is created. Separate multiple tags with a comma (`,`).
- `vpc` - (Required, Forces new resource, String) The VPC ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` - (String) The date and time when the routing table was created.
- `crn` - (String) The CRN of the default routing table.
- `href` - (String) The routing table URL.
- `id` - (String) The unique identifier of the routing table. The ID is composed of `<vpc_id>/<vpc_routing_table_id>`.
- `is_default` - (Bool) Indicates whether this is the default routing table for this VPC. Will always be `true` for this resource.
- `lifecycle_state` - (String) The lifecycle state of the routing table.
- `name` - (String) The user-defined name for this routing table. This is automatically assigned by IBM Cloud for default routing tables.
- `resource_group` - (List) The resource group for this routing table.

  Nested scheme for `resource_group`:
  - `href` - (String) The URL for this resource group.
  - `id` - (String) The unique identifier for this resource group.
  - `name` - (String) The name for this resource group.
- `resource_type` - (String) The resource type.
- `routing_table` - (String) The unique routing table identifier.
- `subnets` - (List) The subnets to which this routing table is attached.

  Nested scheme for `subnets`:
  - `id` - (String) The unique ID of the subnet.
  - `name` - (String) The user defined name of the subnet.

## Import
The `ibm_is_vpc_default_routing_table` resource can be imported by using VPC ID and the default VPC routing table ID.

**Example**

```
$ terraform import ibm_is_vpc_default_routing_table.example 56738c92-4631-4eb5-8938-8af9211a6ea4/fc2667e0-9e6f-4993-a0fd-cabab477c4d1
```

## Differences from custom routing tables

The default routing table resource differs from custom routing tables (`ibm_is_vpc_routing_table`) in several key ways:

- **Cannot be deleted**: The default routing table is permanent and cannot be destroyed
- **Always exists**: Created automatically when a VPC is created
- **Only one per VPC**: Each VPC has exactly one default routing table
- **No route management**: Individual routes are managed separately using `ibm_is_vpc_route` resources

## Notes

- The default routing table cannot be deleted. When you run `terraform destroy`, the resource will be removed from Terraform state, but the actual routing table will remain in IBM Cloud.
- To manage individual routes in the default routing table, use the `ibm_is_vpc_route` resource with the default routing table ID.
- Changes to ingress routing settings may affect traffic flow to your VPC. Plan changes carefully.