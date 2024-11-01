# Examples for Context Based Restrictions

These examples illustrate how to use the resources and data sources associated with Context Based Restrictions.

The following resources are supported:
* ibm_cbr_zone
* ibm_cbr_rule

The following data sources are supported:
* ibm_cbr_zone
* ibm_cbr_rule

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Context Based Restrictions resources

## ContextBasedRestrictionsV1 resources

cbr_zone resource:

```hcl
resource "ibm_cbr_zone" "cbr_zone_instance" {
  name = var.cbr_zone_name
  account_id = var.cbr_zone_account_id
  description = var.cbr_zone_description
  addresses = var.cbr_zone_addresses
  excluded = var.cbr_zone_excluded
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|----------|
| name | The name of the zone. | `string` | true     |
| account_id | The id of the account owning this zone. | `string` | true     |
| description | The description of the zone. | `string` | false    |
| addresses | The list of addresses in the zone. | `list()` | true     |
| excluded | The list of excluded addresses in the zone. Only addresses of type `ipAddress`, `ipRange`, and `subnet` can be excluded. | `list()` | false    |

#### Outputs

| Name | Description |
|------|-------------|
| crn | The zone CRN. |
| address_count | The number of addresses in the zone. |
| excluded_count | The number of excluded addresses in the zone. |
| href | The href link to the resource. |
| created_at | The time the resource was created. |
| created_by_id | IAM ID of the user or service which created the resource. |
| last_modified_at | The last time the resource was modified. |
| last_modified_by_id | IAM ID of the user or service which modified the resource. |

### Resource: ibm_cbr_rule

```hcl
resource "ibm_cbr_rule" "cbr_rule_instance" {
  description = var.cbr_rule_description
  contexts = var.cbr_rule_contexts
  resources = var.cbr_rule_resources
  operations = var.cbr_rule_operations
  enforcement_mode = var.cbr_rule_enforcement_mode
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|----------|
| description | The description of the rule. | `string` | false    |
| contexts | The contexts this rule applies to. | `list()` | true     |
| resources | The resources this rule apply to. | `list()` | true     |
| operations | The operations this rule applies to. | `` | false    |
| enforcement_mode | The rule enforcement mode: * `enabled` - The restrictions are enforced and reported. This is the default. * `disabled` - The restrictions are disabled. Nothing is enforced or reported. * `report` - The restrictions are evaluated and reported, but not enforced. | `string` | false    |

#### Outputs

| Name | Description |
|------|-------------|
| crn | The rule CRN. |
| href | The href link to the resource. |
| created_at | The time the resource was created. |
| created_by_id | IAM ID of the user or service which created the resource. |
| last_modified_at | The last time the resource was modified. |
| last_modified_by_id | IAM ID of the user or service which modified the resource. |

## Context Based Restrictions data sources

### Data source: ibm_cbr_zone

```hcl
data "ibm_cbr_zone" "cbr_zone_instance" {
  zone_id = var.data_cbr_zone_zone_id
}
```
### Data source: ibm_cbr_rule

```hcl
data "cbr_rule" "cbr_rule_instance" {
  rule_id = var.data_cbr_rule_rule_id
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
