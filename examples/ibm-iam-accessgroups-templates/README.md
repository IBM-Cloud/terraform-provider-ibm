# Example for IAMAccessGroupsV2

This example illustrates how to use the IAMAccessGroupsV2

The following types of resources are supported:

* ibm_iam_access_group_template
* ibm_iam_access_group_template_version
* ibm_iam_access_group_template_assignment

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IAMAccessGroupsV2 resources

ibm_iam_access_group_template resource:

```hcl
resource "iam_access_group_template" "iam_access_group_template_instance" {
  transaction_id = var.iam_access_group_template_transaction_id
  name = var.iam_access_group_template_name
  description = var.iam_access_group_template_description
  account_id = var.iam_access_group_template_account_id
  group = var.iam_access_group_template_group
  policy_template_references = var.iam_access_group_template_policy_template_references
}
```
ibm_iam_access_group_template_version resource:

```hcl
resource "iam_access_group_template_version" "iam_access_group_template_version_instance" {
  template_id = var.iam_access_group_template_version_template_id
  transaction_id = var.iam_access_group_template_version_transaction_id
  name = var.iam_access_group_template_version_name
  description = var.iam_access_group_template_version_description
  group = var.iam_access_group_template_version_group
  policy_template_references = var.iam_access_group_template_version_policy_template_references
}
```
ibm_iam_access_group_template_assignment resource:

```hcl
resource "iam_access_group_template_assignment" "iam_access_group_template_assignment_instance" {
  transaction_id = var.iam_access_group_template_assignment_transaction_id
  template_id = var.iam_access_group_template_assignment_template_id
  template_version = var.iam_access_group_template_assignment_template_version
  target_type = var.iam_access_group_template_assignment_target_type
  target = var.iam_access_group_template_assignment_target
}
```

## IamAccessGroupsV2 data sources

ibm_iam_access_group_template data source:

```hcl
data "iam_access_group_template" "iam_access_group_template_instance" {
  account_id = var.iam_access_group_template_account_id
  transaction_id = var.iam_access_group_template_transaction_id
  verbose = var.iam_access_group_template_verbose
}
```
ibm_iam_access_group_template_versions data source:

```hcl
data "ibm_iam_access_group_template_version" "ibm_iam_access_group_template_version_instance" {
  template_id = var.ibm_iam_access_group_template_version_template_id
}
```
ibm_iam_access_group_template_assignment data source:

```hcl
data "iam_access_group_template_assignment" "iam_access_group_template_assignment_instance" {
  account_id = var.iam_access_group_template_assignment_account_id
  template_id = var.iam_access_group_template_assignment_template_id
  template_version = var.iam_access_group_template_assignment_template_version
  target = var.iam_access_group_template_assignment_target
  status = var.iam_access_group_template_assignment_status
  transaction_id = var.iam_access_group_template_assignment_transaction_id
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
| transaction_id | An optional transaction id for the request. | `string` | false |
| name | The name of the access group template. | `string` | true |
| description | The description of the access group template. | `string` | false |
| account_id | The ID of the account to which the access group template is assigned. | `string` | true |
| group | Access Group Component. | `` | false |
| policy_template_references | References to policy templates assigned to the access group template. | `list()` | false |
| template_id | ID of the template that you want to create a new version of. | `string` | true |
| transaction_id | An optional transaction id for the request. | `string` | false |
| name | The name of the access group template. | `string` | false |
| description | The description of the access group template. | `string` | false |
| group | Access Group Component. | `` | false |
| policy_template_references | References to policy templates assigned to the access group template. | `list()` | false |
| transaction_id | An optional transaction id for the request. | `string` | false |
| template_id | The ID of the template that the assignment is based on. | `string` | true |
| template_version | The version of the template that the assignment is based on. | `string` | true |
| target_type | The type of the entity that the assignment applies to. | `string` | true |
| target | The ID of the entity that the assignment applies to. | `string` | true |
| account_id | Enterprise account ID. | `string` | true |
| transaction_id | An optional transaction id for the request. | `string` | false |
| verbose | If `verbose=true`, IAM resource details are returned. If performance is a concern, leave the `verbose` parameter off so that details are not retrieved. | `bool` | false |
| template_id | ID of the template that you want to list all versions of. | `string` | true |
| account_id | Enterprise account ID. | `string` | true |
| template_id | Filter results by Template Id. | `string` | false |
| template_version | Filter results by Template Version. | `string` | false |
| target | Filter results by the assignment target. | `string` | false |
| status | Filter results by the assignment status. | `string` | false |
| transaction_id | An optional transaction id for the request. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| iam_access_group_template | iam_access_group_template object |
| iam_access_group_template_version | iam_access_group_template_version object |
| iam_access_group_template_assignment | iam_access_group_template_assignment object |
| iam_access_group_template | iam_access_group_template object |
| ibm_iam_access_group_template_version | ibm_iam_access_group_template_version object |
| iam_access_group_template_assignment | iam_access_group_template_assignment object |
