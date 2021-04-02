# Example for Fetching SecretsManager Secrets

This example illustrates how to use the SecretsManagerV1 to fetch SecretsManager Secrets

These types of resources are supported:


## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## SecretsManagerV1 resources


## SecretsManagerV1 Data sources

secrets_manager_secrets data source:

```hcl
data "secrets_manager_secrets" "secrets_manager_secrets_instance" {
  instance_id = var.secrets_manager_instance_id
  secret_type = var.secrets_manager_secrets_secret_type
}
```
secrets_manager_secret data source:

```hcl
data "secrets_manager_secret" "secrets_manager_secret_instance" {
  instance_id = var.secrets_manager_instance_id
  secret_type = var.secrets_manager_secret_secret_type
  secret_id = var.secrets_manager_secret_id
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
| ibm | 1.22.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` |  | true |
| region | Secrets Manager Instance region | `string` | us-south | false |
| secrets\_manager\_instance\_id | Secrets Manager Instance GUID | `string` |  | true |
| secrets\_manager\_secrets\_secret\_type | The secret type. Supported options include: arbitrary, iam_credentials, username_password. | `string` | null | false |
| secrets\_manager\_secret\_secret\_type | The secret type. Supported options include: arbitrary, iam_credentials, username_password. | `string` |  | true |
| secret\_id | The v4 UUID that uniquely identifies the secret. | `string` |  | true |

## Outputs

| Name | Description |
|------|-------------|
| secrets\_manager\_secrets | secrets\_manager\_secrets object |
| secrets\_manager\_secret | secrets\_manager\_secret object |
