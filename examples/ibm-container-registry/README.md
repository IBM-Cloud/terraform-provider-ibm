# Example for ContainerRegistryV1

This example illustrates how to use the ContainerRegistryV1

These types of resources are supported:

* cr_namespace
* cr_retention_policy

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ContainerRegistryV1 resources

cr_namespace resource:

```hcl
resource "cr_namespace" "cr_namespace_instance" {
  name = var.cr_namespace_name
  resource_group_id = data.ibm_resource_group.default_group.id
  tags = var.cr_namespace_tags
}
```
cr_retention_policy resource:

```hcl
resource "cr_retention_policy" "cr_retention_policy_instance" {
  namespace = var.cr_retention_policy_namespace
  images_per_repo = var.cr_retention_policy_images_per_repo
  retain_untagged = var.cr_retention_policy_retain_untagged
}
```

## ContainerRegistryV1 Data sources


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
| name | The name of the namespace. | `string` | true |
| resource_group_id | The ID of the resource group that the namespace will be created within. | `string` | false |
| namespace | The namespace to which the retention policy is attached. | `string` | true |
| images_per_repo | Determines how many images will be retained for each repository when the retention policy is executed. The value -1 denotes 'Unlimited' (all images are retained). | `number` | true |
| retain_untagged | Determines if untagged images are retained when executing the retention policy. This is false by default meaning untagged images will be deleted when the policy is executed. | `bool` | false |

## Outputs

| Name | Description |
|------|-------------|
| cr_namespace | cr_namespace object |
| cr_retention_policy | cr_retention_policy object |
