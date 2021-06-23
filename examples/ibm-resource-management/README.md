# Example for Resource Management

This example illustrates how to use the Resource Management resources and datasources

These types of resources are supported:

* resource_alias
* resource_binding

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## Resource Management resources

resource_alias resource:

```hcl
resource "ibm_resource_alias" "resource_alias_instance" {
  name = var.resource_alias_name
  source = var.resource_alias_source
  target = var.resource_alias_target
}
```
resource_binding resource:

```hcl
resource "ibm_resource_binding" "resource_binding_instance" {
  source = var.resource_binding_source
  target = var.resource_binding_target
  name = var.resource_binding_name
  parameters = var.resource_binding_parameters
  role = var.resource_binding_role
}
```

## Resource Management Data sources

resource_aliases data source:

```hcl
data "ibm_resource_aliases" "resource_aliases_instance" {
  name = var.resource_aliases_name
}
```
resource_bindings data source:

```hcl
data "ibm_resource_bindings" "resource_bindings_instance" {
  name = var.resource_bindings_name
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name of the alias. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`. | `string` | true |
| source | The short or long ID of resource instance. | `string` | true |
| target | The CRN of target name(space) in a specific environment, for example, space in Dallas YP, CFEE instance etc. | `string` | true |
| source | The short or long ID of resource alias. | `string` | true |
| target | The CRN of application to bind to in a specific environment, for example, Dallas YP, CFEE instance. | `string` | true |
| name | The name of the binding. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`. | `string` | false |
| parameters | Configuration options represented as key-value pairs. Service defined options are passed through to the target resource brokers, whereas platform defined options are not. | `` | false |
| role | The role name or it's CRN. | `string` | false |
| name | The human-readable name of the alias. | `string` | false |
| name | The human-readable name of the binding. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| resource_alias | resource_alias object |
| resource_binding | resource_binding object |
| resource_aliases | resource_aliases object |
| resource_bindings | resource_bindings object |
