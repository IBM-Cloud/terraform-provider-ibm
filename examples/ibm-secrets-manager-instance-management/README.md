# Examples for IBM Cloud Secrets Manager Instance Management API

These examples illustrate how to use the resources and data sources associated with IBM Cloud Secrets Manager Instance Management API.

The following data sources are supported:
* ibm_sm_instance

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Cloud Secrets Manager Instance Management API data sources

### Data source: ibm_sm_instance

```hcl
data "ibm_sm_instance" "sm_instance_instance" {
  instance_id = var.sm_instance_instance_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | The service instance ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| instance_crn | The instance CRN identifier. |
| plan | Instance plan name. |
| vault_cluster | Vault cluster information for Vault Dedicated instances. |
| endpoints | Instance endpoints for Vault Dedicated instances. |
| encryption | Vault encryption configuration for Vault Dedicated instances. |

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
