# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

The following types of resources are supported:

* trusted_profile_template_assignment

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

trusted_profile_template_assignment resource:

```hcl
resource "trusted_profile_template_assignment" "trusted_profile_template_assignment_instance" {
  template_id = var.trusted_profile_template_assignment_instance_template_id
  template_version = var.trusted_profile_template_assignment_instance_template_version
  target_type = var.trusted_profile_template_assignment_instance_target_type
  target = var.trusted_profile_template_assignment_instance_target
}
```

## IamIdentityV1 data sources

trusted_profile_template_assignment data source:

```hcl
data "trusted_profile_template_assignment" "trusted_profile_template_assignment_instance" {
  assignment_id = var.trusted_profile_template_assignment_instance_assignment_id
  include_history = var.trusted_profile_template_assignment_instance_include_history
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
| template_id | Template Id. | `string` | true |
| template_version | Template version. | `number` | true |
| target_type | Assignment target type. | `string` | true |
| target | Assignment target. | `string` | true |
| assignment_id | ID of the Assignment Record. | `string` | true |
| include_history | Defines if the entity history is included in the response. | `bool` | false |

## Outputs

| Name | Description |
|------|-------------|
| trusted_profile_template_assignment | trusted_profile_template_assignment object |
| trusted_profile_template_assignment | trusted_profile_template_assignment object |
