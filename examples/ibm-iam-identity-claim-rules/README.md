# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

These types of resources are supported:

* iam_trusted_profiles_claim_rule

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

iam_trusted_profiles_claim_rule resource:

```hcl
resource "iam_trusted_profiles_claim_rule" "iam_trusted_profiles_claim_rule_instance" {
  profile_id = var.iam_trusted_profiles_claim_rule_profile_id
  type = var.iam_trusted_profiles_claim_rule_type
  conditions = var.iam_trusted_profiles_claim_rule_conditions
  context = var.iam_trusted_profiles_claim_rule_context
  name = var.iam_trusted_profiles_claim_rule_name
  realm_name = var.iam_trusted_profiles_claim_rule_realm_name
  cr_type = var.iam_trusted_profiles_claim_rule_cr_type
  expiration = var.iam_trusted_profiles_claim_rule_expiration
}
```

## IamIdentityV1 Data sources

iam_trusted_profiles_claim_rule data source:

```hcl
data "iam_trusted_profiles_claim_rule" "iam_trusted_profiles_claim_rule_instance" {
  profile_id = var.iam_trusted_profiles_claim_rule_profile_id
  rule_id = var.iam_trusted_profiles_claim_rule_rule_id
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
| profile_id | ID of the trusted profile to create a claim rule. | `string` | true |
| type | Type of the calim rule, either 'Profile-SAML' or 'Profile-CR'. | `string` | true |
| conditions | Conditions of this claim rule. | `list()` | true |
| context | Context with key properties for problem determination. | `` | false |
| name | Name of the claim rule to be created or updated. | `string` | false |
| realm_name | The realm name of the Idp this claim rule applies to. This field is required only if the type is specified as 'Profile-SAML'. | `string` | false |
| cr_type | The compute resource type the rule applies to, required only if type is specified as 'Profile-CR'. Valid values are VSI, IKS_SA, ROKS_SA. | `string` | false |
| expiration | Session expiration in seconds, only required if type is 'Profile-SAML'. | `number` | false |
| profile_id | ID of the trusted profile. | `string` | true |
| rule_id | ID of the claim rule to get. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| iam_trusted_profiles_claim_rule | iam_trusted_profiles_claim_rule object |
| iam_trusted_profiles_claim_rule | iam_trusted_profiles_claim_rule object |
