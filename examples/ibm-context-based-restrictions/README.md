# Example for ContextBasedRestrictionsV1

This example illustrates how to use the ContextBasedRestrictionsV1

These types of resources are supported:

* cbr_zone
* cbr_rule

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ContextBasedRestrictionsV1 resources

cbr_zone resource:

```hcl
resource "cbr_zone" "cbr_zone_instance" {
  name = var.cbr_zone_name
  account_id = var.cbr_zone_account_id
  description = var.cbr_zone_description
  addresses = var.cbr_zone_addresses
  excluded = var.cbr_zone_excluded
}
```
cbr_rule resource:

```hcl
resource "cbr_rule" "cbr_rule_instance" {
  description = var.cbr_rule_description
  contexts = var.cbr_rule_contexts
  resources = var.cbr_rule_resources
}
```

## ContextBasedRestrictionsV1 Data sources

cbr_zone data source:

```hcl
data "cbr_zone" "cbr_zone_instance" {
  zone_id = var.cbr_zone_zone_id
}
```
cbr_rule data source:

```hcl
data "cbr_rule" "cbr_rule_instance" {
  rule_id = var.cbr_rule_rule_id
}
```

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
| name | The name of the zone. | `string` | false |
| account_id | The id of the account owning this zone. | `string` | false |
| description | The description of the zone. | `string` | false |
| addresses | The list of addresses in the zone. | `list()` | false |
| excluded | The list of excluded addresses in the zone. Only addresses of type `ipAddress`, `ipRange`, and `subnet` can be excluded. | `list()` | false |
| description | The description of the rule. | `string` | false |
| contexts | The contexts this rule applies to. | `list()` | false |
| resources | The resources this rule apply to. | `list()` | false |
| operations | The operations this rule applies to. | `` | false |
| enforcement_mode | The rule enforcement mode: * `enabled` - The restrictions are enforced and reported. This is the default. * `disabled` - The restrictions are disabled. Nothing is enforced or reported. * `report` - The restrictions are evaluated and reported, but not enforced. | `string` | false |
| zone_id | The ID of a zone. | `string` | true |
| rule_id | The ID of a rule. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| cbr_zone | cbr_zone object |
| cbr_rule | cbr_rule object |
| cbr_zone | cbr_zone object |
| cbr_rule | cbr_rule object |
