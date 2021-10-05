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
  transaction_id = var.cbr_zone_transaction_id
}
```
cbr_rule resource:

```hcl
resource "cbr_rule" "cbr_rule_instance" {
  description = var.cbr_rule_description
  contexts = var.cbr_rule_contexts
  resources = var.cbr_rule_resources
  transaction_id = var.cbr_rule_transaction_id
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
| name | The name of the zone. | `string` | false |
| account_id | The id of the account owning this zone. | `string` | false |
| description | The description of the zone. | `string` | false |
| addresses | The list of addresses in the zone. | `list()` | false |
| excluded | The list of excluded addresses in the zone. | `list()` | false |
| transaction_id | The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a transaction ID in the response.**Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each request. | `string` | false |
| description | The description of the rule. | `string` | false |
| contexts | The contexts this rule applies to. | `list()` | false |
| resources | The resources this rule apply to. | `list()` | false |
| transaction_id | The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a transaction ID in the response.**Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each request. | `string` | false |
| zone_id | The ID of a zone. | `string` | true |
| rule_id | The ID of a rule. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| cbr_zone | cbr_zone object |
| cbr_rule | cbr_rule object |
| cbr_zone | cbr_zone object |
| cbr_rule | cbr_rule object |
