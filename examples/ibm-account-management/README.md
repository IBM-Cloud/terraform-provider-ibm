# Examples for Account Management

These examples illustrate how to use the resources and data sources associated with Account Management.

The following data sources are supported:
* ibm_account_info

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Account Management data sources

### Data source: ibm_account_info

```hcl
data "ibm_account_info" "account_instance" {
  account_id = var.account_account_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| account_id | The unique identifier of the account you want to retrieve. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name |  |
| owner |  |
| owner_userid |  |
| owner_iamid |  |
| type |  |
| status |  |
| linked_softlayer_account |  |
| team_directory_enabled |  |
| traits |  |

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
