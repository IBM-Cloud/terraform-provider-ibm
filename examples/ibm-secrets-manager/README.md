# Example for SecretsManagerV1

This example illustrates how to use the SecretsManagerV1

These types of resources are supported:

* SecretGroup

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## SecretsManagerV1 resources

secret_group resource:

```hcl
resource "secret_group" "secret_group_instance" {
  name = var.secret_group_name
  description = var.secret_group_description
}
```

## SecretsManagerV1 Data sources


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
| name | The name of your secret group. | `string` | true |
| description | An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| secret_group | secret_group object |
