# Example for ConfigurationGovernanceV1

This example illustrates how to use the ConfigurationGovernanceV1

These types of resources are supported:

* scc_template
* scc_template_attachment
* scc_rule
* scc_rule_attachment

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ConfigurationGovernanceV1 resources

scc_template resource:

```hcl
resource "scc_template" "scc_template_instance" {
  template = var.scc_template_template
}
```
scc_template_attachment resource:

```hcl
resource "scc_template_attachment" "scc_template_attachment_instance" {
  template_id = var.scc_template_attachment_template_id
  attachment = var.scc_template_attachment_attachment
}
```
scc_rule resource:

```hcl
resource "scc_rule" "scc_rule_instance" {
  rule = var.scc_rule_rule
}
```
scc_rule_attachment resource:

```hcl
resource "scc_rule_attachment" "scc_rule_attachment_instance" {
  rule_id = var.scc_rule_attachment_rule_id
  attachment = var.scc_rule_attachment_attachment
}
```

## ConfigurationGovernanceV1 Data sources


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
| template | A list of templates to be created. | `list()` | true |
| template_id | The UUID that uniquely identifies the template. | `string` | true |
| attachment |  | `list()` | true |
| rule | A list of rules to be created. | `list()` | true |
| rule_id | The UUID that uniquely identifies the rule. | `string` | true |
| attachment |  | `list()` | true |

## Outputs

| Name | Description |
|------|-------------|
| scc_template | scc_template object |
| scc_template_attachment | scc_template_attachment object |
| scc_rule | scc_rule object |
| scc_rule_attachment | scc_rule_attachment object |
