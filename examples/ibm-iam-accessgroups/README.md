# Example for IBM IAM Access Group Account Settings

This example illustrates how to use the IAM Access Group Account  Settings to configure public access on groups.

These types of resources are supported:

* ibm_iam_access_group_account_settings

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IAM Access Group Account Setting Resource

Access Group Account Setting Resource:

```hcl
resource "ibm_iam_access_group_account_settings" "iam_access_groups_account_settings_instance" {
  public_access_enabled = true
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name                  | Description                                 | Type      | Required |
|-----------------------|---------------------------------------------|-----------|----------|
| public_access_enabled | Defines if public access groups are enabled | `boolean` | true     |                                                                                                                  | `string` | false |

## Outputs

| Name | Description   |
|------|---------------|
| public_access_enabled | boolean |
