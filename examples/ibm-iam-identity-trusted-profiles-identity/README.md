# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

These types of resources are supported:

* iam_trusted_profile_identity

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

iam_trusted_profile_identity resource:

```hcl
resource "iam_trusted_profile_identity" "iam_trusted_profile_identity_instance" {
  profile_id = var.iam_trusted_profile_identity_profile_id
  identity_type = var.iam_trusted_profile_identity_identity_type
  identifier = var.iam_trusted_profile_identity_identifier
  type = var.iam_trusted_profile_identity_type
  accounts = var.iam_trusted_profile_identity_accounts
  description = var.iam_trusted_profile_identity_description
}
```

## IamIdentityV1 Data sources


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
| identity_type | Type of the identity. | `string` | true |
| identifier | Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN. | `string` | true |
| type | Type of the identity. | `string` | true |
| accounts | Only valid for the type user. Accounts from which a user can assume the trusted profile. | `list(string)` | false |
| description | Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| iam_trusted_profile_identity | iam_trusted_profile_identity object |
