# Examples for IAM Identity Services

These examples illustrate how to use the resources and data sources associated with IAM Identity Services.

The following resources are supported:
* ibm_iam_trusted_profile_identities

The following data sources are supported:
* ibm_iam_trusted_profile_identities

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IAM Identity Services resources

### Resource: ibm_iam_trusted_profile_identities

```hcl
resource "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities_instance" {
  profile_id = var.iam_trusted_profile_identities_profile_id
  if_match = var.iam_trusted_profile_identities_if_match
  identities = var.iam_trusted_profile_identities_identities
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| profile_id | ID of the trusted profile. | `string` | true |
| if_match | Entity tag of the Identities to be updated. Specify the tag that you retrieved when reading the Profile Identities. This value helps identify parallel usage of this API. Pass * to indicate updating any available version, which may result in stale updates. | `string` | true |
| identities | List of identities. | `list()` | false |

## IAM Identity Services data sources

### Data source: ibm_iam_trusted_profile_identities

```hcl
data "ibm_iam_trusted_profile_identities" "iam_trusted_profile_identities_instance" {
  profile_id = var.data_iam_trusted_profile_identities_profile_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| profile_id | ID of the trusted profile. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| entity_tag | Entity tag of the profile identities response. |
| identities | List of identities. |

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
