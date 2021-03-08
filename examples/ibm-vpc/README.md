# Example for VpcV1

This example illustrates how to use the VpcV1

These types of resources are supported:


## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## VpcV1 resources


## VpcV1 Data sources

is_dedicated_host data source:

```hcl
data "is_dedicated_host" "is_dedicated_host_instance" {
  id = var.is_dedicated_host_id
}
```
is_dedicated_hosts data source:

```hcl
data "is_dedicated_hosts" "is_dedicated_hosts_instance" {
  id = var.is_dedicated_hosts_id
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
| id | The unique identifier for this virtual server instance. | `string` | false |
| id | The unique identifier for this dedicated host. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| is_dedicated_host | is_dedicated_host object |
| is_dedicated_hosts | is_dedicated_hosts object |
