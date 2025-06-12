# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

These types of resources are supported:


## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources


## IamIdentityV1 Data sources

iam_user_mfa_enrollments data source:

```hcl
data "iam_user_mfa_enrollments" "iam_user_mfa_enrollments_instance" {
  account_id = var.iam_user_mfa_enrollments_account_id
  iam_id = var.iam_user_mfa_enrollments_iam_id
}
```
iam_mfa_report data source:

```hcl
data "iam_mfa_report" "iam_mfa_report_instance" {
  account_id = var.iam_mfa_report_account_id
  reference = var.iam_mfa_report_reference
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
| account_id | ID of the account. | `string` | true |
| iam_id | iam_id of the user. This user must be the member of the account. | `string` | true |
| account_id | ID of the account. | `string` | true |
| reference | Reference for the report to be generated, You can use 'latest' to get the latest report for the given account. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| iam_user_mfa_enrollments | iam_user_mfa_enrollments object |
| iam_mfa_report | iam_mfa_report object |
