# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

These types of resources are supported:

* iam_trusted_profiles_link

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

iam_trusted_profiles_link resource:

```hcl
resource "iam_trusted_profiles_link" "iam_trusted_profiles_link_instance" {
  profile_id = var.iam_trusted_profiles_link_profile_id
  cr_type = var.iam_trusted_profiles_link_cr_type
  link = var.iam_trusted_profiles_link_link
  name = var.iam_trusted_profiles_link_name
}
```

## IamIdentityV1 Data sources

iam_trusted_profiles_link data source:

```hcl
data "iam_trusted_profiles_link" "iam_trusted_profiles_link_instance" {
  profile_id = var.iam_trusted_profiles_link_profile_id
  link_id = var.iam_trusted_profiles_link_link_id
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
| profile_id | ID of the trusted profile. | `string` | true |
| cr_type | The compute resource type. Valid values are VSI, IKS_SA, ROKS_SA. | `string` | true |
| link | Link details. | `` | true |
| name | Optional name of the Link. | `string` | false |
| profile_id | ID of the trusted profile. | `string` | true |
| link_id | ID of the link. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| iam_trusted_profiles_link | iam_trusted_profiles_link object |
| iam_trusted_profiles_link | iam_trusted_profiles_link object |
