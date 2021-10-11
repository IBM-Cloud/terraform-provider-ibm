# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

These types of resources are supported:

* iam_trusted_profiles

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

iam_trusted_profiles resource:

```hcl
resource "iam_trusted_profiles" "iam_trusted_profiles_instance" {
  name = var.iam_trusted_profiles_name
  account_id = var.iam_trusted_profiles_account_id
  description = var.iam_trusted_profiles_description
}
```

## IamIdentityV1 Data sources

iam_trusted_profiles data source:

```hcl
data "iam_trusted_profiles" "iam_trusted_profiles_instance" {
  profile_id = var.iam_trusted_profiles_profile_id
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
| name | Name of the trusted profile. The name is checked for uniqueness. Therefore trusted profiles with the same names can not exist in the same account. | `string` | true |
| account_id | The account ID of the trusted profile. | `string` | true |
| description | The optional description of the trusted profile. The 'description' property is only available if a description was provided during creation of trusted profile. | `string` | false |
| profile_id | ID of the trusted profile to get. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| iam_trusted_profiles | iam_trusted_profiles object |
| iam_trusted_profiles | iam_trusted_profiles object |
