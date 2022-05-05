# Example for Transit Gateway resources

This example shows how to create Transit Gateway resources.

Following types of resources are supported:

* [Transit Gateway](https://cloud.ibm.com/docs/terraform?topic=terraform-tg-resource#tg-gateway-resource)
* [Transit Gateway Connection](https://cloud.ibm.com/docs/terraform?topic=terraform-tg-resource#tg-connection)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.5.1`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.27.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Transit Gateway Resources

Create a transit gateway:

```hcl
resource "ibm_tg_gateway" "new_tg_gw"{
	name=var.name
	location=var.location
	global=true
	resource_group = data.ibm_resource_group.rg.id
}  
```

Create a transit gateway connection:

```hcl
resource "ibm_is_vpc" "test_tg_vpc" {
  name = var.vpc_name
}

resource "ibm_tg_connection" "test_ibm_tg_connection"{
	gateway = "${ibm_tg_gateway.new_tg_gw.id}"
	network_type = var.network_type
	name= vc_name
	network_id = ibm_is_vpc.test_tg_vpc.resource_crn
}  
```

Create a transit gateway connection prefix filter:
```hcl
resource "ibm_tg_connection_prefix_filter" "test_tg_prefix_filter" {
    gateway = ibm_tg_gateway.new_tg_gw.id
    connection_id = ibm_tg_connection.test_ibm_tg_connection.connection_id
    action = "permit"
    prefix = "192.168.100.0/24"
    le = 0
    ge = 32
}
```

Create a transit gateway route report:

```hcl
resource ibm_tg_route_report" "test_tg_route_report" {
	gateway = ibm_tg_gateway.new_tg_gw.id
}
```
## Transit Gateway Data Sources

Retrieves specified Transit Gateway:

```hcl

data "ibm_tg_gateway" "tg_gateway" {
	name= ibm_tg_gateway.new_tg_gw.name
}

```
List all the Transit Gateways in the account.

```hcl
data "ibm_tg_gateways" "all_tg_gws"{
}
```
List all locations that support Transit Gateways
```hcl
data "ibm_tg_locations" "tg_locations" {
}
```
Get the details of a Transit Gateway Location.
```hcl
data "ibm_tg_location" "tg_location" {
	name = var.location
} 
```
List all prefix filters for a Transit Gateway Connection
````
data "ibm_tg_connection_prefix_filters" "tg_prefix_filters" {
    gateway = ibm_tg_gateway.new_tg_gw.id
    connection_id = ibm_tg_connection.test_ibm_tg_connection.connection_id
}
```
Retrieve specified Transit Gateway Connection Prefix Filter
```
data "ibm_tg_connection_prefix_filter" "tg_prefix_filter" {
    gateway = ibm_tg_gateway.new_tg_gw.id
    connection_id = ibm_tg_connection.test_ibm_tg_connection.connection_id
	filter_id = ibm_tg_connection_prefix_filter.test_tg_prefix_filter.filter_id
}
```
List all route reports for a Transit Gateway
```
data "ibm_tg_route_reports" "tg_route_reports" {
	gateway = ibm_tg_gateway.new_tg_gw.id
}
```
Retrieve specified Transit Gateway Route Report
```
data "ibm_tg_route_report" "tg_route_report" {
	gateway = ibm_tg_gateway.new_tg_gw.
	route_report = ibm_tg_route_report_test_tg_route_report.route_report_id
}
```
## Examples

* [ Transit Gateway](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-transit-gateway)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | Name of Transit Gateway Services. | `string` | yes |
| location |  Location of Transit Gateway Services. | `string` | yes |
| global | Allow global routing for a Transit Gateway. If unspecified, the default value is false. | `boolean` | no |
| vc_name | Name of Transit Gateway Connection . | `string` | yes |
| vpc_name | Name of VPC . | `string` | yes |



<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
