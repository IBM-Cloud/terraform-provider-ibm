# Examples for IAM Identity Services

These examples illustrate how to use the resources and data sources associated with IAM Identity Services.

The following resources are supported:
* ibm_iam_identity_preference

The following data sources are supported:
* ibm_iam_identity_preference

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IAM Identity Services resources

### Resource: ibm_iam_identity_preference

```hcl
resource "ibm_iam_identity_preference" "iam_identity_preference_instance" {
  account_id = var.iam_identity_preference_account_id
  iam_id = var.iam_identity_preference_iam_id
  service = var.iam_identity_preference_service
  preference_id = var.iam_identity_preference_preference_id
  value_string = var.iam_identity_preference_value_string
  value_list_of_strings = var.iam_identity_preference_value_list_of_strings
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| account_id | Account id to update preference for. | `string` | true |
| iam_id | IAM id to update the preference for. | `string` | true |
| service | Service of the preference to be updated. | `string` | true |
| preference_id | Identifier of preference to be updated. | `string` | true |
| value_string | String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present. | `string` | true |
| value_list_of_strings | List of value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| scope | Scope of the preference, 'global' or 'account'. |
| preference_id | Unique ID of the preference. |

## IAM Identity Services data sources

### Data source: ibm_iam_identity_preference

```hcl
data "ibm_iam_identity_preference" "iam_identity_preference_instance" {
  account_id = var.data_iam_identity_preference_account_id
  iam_id = var.data_iam_identity_preference_iam_id
  service = var.data_iam_identity_preference_service
  preference_id = var.data_iam_identity_preference_preference_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| account_id | Account id to get preference for. | `string` | true |
| iam_id | IAM id to get the preference for. | `string` | true |
| service | Service of the preference to be fetched. | `string` | true |
| preference_id | Identifier of preference to be fetched. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| scope | Scope of the preference, 'global' or 'account'. |
| value_string | String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present. |
| value_list_of_strings | List of value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present. |

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
