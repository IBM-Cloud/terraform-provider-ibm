# Example for IBM IAM User Policy

This example illustrates how to use the IAM User policy resource to assign roles and give access to certain content on cloud using policies for a particular IBM Cloud User. 

These types of resources are supported:

* [IBM IAM User Policy](https://cloud.ibm.com/docs/terraform?topic=terraform-iam-resources#iam-user-policy)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.9.0`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.31.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IAM User Policy Resource

API Gateway Endpoint resource with single OpenAPI document:

```hcl
data "ibm_resource_group" "group" {
  name = var.resource_group
}

resource "ibm_iam_user_policy" "iam_policy" {
  ibm_id = var.ibm_id1
  roles  = ["Viewer", "Editor"]

  resources {
    resource_type = "resource-group"
    resource      = data.ibm_resource_group.group.id
  }
}
```
##  IAM User Policy Data Source
Lists all Policies of a particular IBM ID user.

```hcl
data "ibm_iam_user_policy" "testacc_ds_user_policy" {
  ibm_id = ibm_iam_user_policy.policy.ibm_id
}
```

## Examples

* [IAM User Policy resource](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-iam-policy)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

Single OpenAPI document or directory of documents.

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibm_id1 | The IBM ID of the user to which policies are to be assigned. | `string` | yes |
| resource_group | The resource group name in which user policies are to be assigned. | `string` | yes |


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
