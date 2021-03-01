# Example for IamPolicyManagementV1

This example illustrates how to use the IamPolicyManagementV1

These types of resources are supported:

* iam_policy
* iam_custom_role

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamPolicyManagementV1 resources

iam_policy resource:

```hcl
resource "iam_policy" "iam_policy_instance" {
  type = var.iam_policy_type
  subjects = var.iam_policy_subjects
  roles = var.iam_policy_roles
  resources = var.iam_policy_resources
  description = var.iam_policy_description
  accept_language = var.iam_policy_accept_language
}
```
iam_custom_role resource:

```hcl
resource "iam_custom_role" "iam_custom_role_instance" {
  display_name = var.iam_custom_role_display_name
  actions = var.iam_custom_role_actions
  name = var.iam_custom_role_name
  account_id = var.iam_custom_role_account_id
  service_name = var.iam_custom_role_service_name
  description = var.iam_custom_role_description
  accept_language = var.iam_custom_role_accept_language
}
```

## IamPolicyManagementV1 Data sources

iam_policy data source:

```hcl
data "iam_policy" "iam_policy_instance" {
  policy_id = var.iam_policy_policy_id
}
```
iam_custom_role data source:

```hcl
data "iam_custom_role" "iam_custom_role_instance" {
  role_id = var.iam_custom_role_role_id
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
| type | The policy type; either 'access' or 'authorization'. | `string` | true |
| subjects | The subjects associated with a policy. | `list()` | true |
| roles | A set of role cloud resource names (CRNs) granted by the policy. | `list()` | true |
| resources | The resources associated with a policy. | `list()` | true |
| description | Customer-defined description. | `string` | false |
| accept_language | Translation language code. | `string` | false |
| display_name | The display name of the role that is shown in the console. | `string` | true |
| actions | The actions of the role. | `list(string)` | true |
| name | The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized. | `string` | true |
| account_id | The account GUID. | `string` | true |
| service_name | The service name. | `string` | true |
| description | The description of the role. | `string` | false |
| accept_language | Translation language code. | `string` | false |
| policy_id | The policy ID. | `string` | true |
| role_id | The role ID. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| iam_policy | iam_policy object |
| iam_custom_role | iam_custom_role object |
| iam_policy | iam_policy object |
| iam_custom_role | iam_custom_role object |
