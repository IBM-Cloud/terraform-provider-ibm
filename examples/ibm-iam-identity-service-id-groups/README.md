# Examples for IAM Identity Services

These examples illustrate how to use the resources and data sources associated with IAM Identity Services.

The following resources are supported:
* ibm_iam_serviceid_group

The following data sources are supported:
* ibm_iam_serviceid_group

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IAM Identity Services resources

### Resource: ibm_iam_serviceid_group

```hcl
resource "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
  account_id = var.iam_serviceid_group_account_id
  name = var.iam_serviceid_group_name
  description = var.iam_serviceid_group_description
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| account_id | ID of the account the service ID group belongs to. | `string` | true |
| name | Name of the service ID group. Unique in the account. | `string` | true |
| description | Description of the service ID group. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| entity_tag | Version of the service ID group details object. You need to specify this value when updating the service ID group to avoid stale updates. |
| crn | Cloud Resource Name of the item. |
| created_at | Timestamp of when the service ID group was created. |
| created_by | IAM ID of the user or service which created the Service Id group. |
| modified_at | Timestamp of when the service ID group was modified. |

## IAM Identity Services data sources

### Data source: ibm_iam_serviceid_group

```hcl
data "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
  iam_serviceid_group_id = var.data_iam_serviceid_group_iam_serviceid_group_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| iam_serviceid_group_id | Unique ID of the service ID group. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| entity_tag | Version of the service ID group details object. You need to specify this value when updating the service ID group to avoid stale updates. |
| account_id | ID of the account the service ID group belongs to. |
| crn | Cloud Resource Name of the item. |
| name | Name of the service ID group. Unique in the account. |
| description | Description of the service ID group. |
| created_at | Timestamp of when the service ID group was created. |
| created_by | IAM ID of the user or service which created the Service Id group. |
| modified_at | Timestamp of when the service ID group was modified. |

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
