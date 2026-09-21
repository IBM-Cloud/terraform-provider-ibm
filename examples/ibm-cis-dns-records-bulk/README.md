# IBM CIS DNS Records Batch Example

This example shows how to use `ibm_cis_dns_records_batch` to create, update, and delete multiple CIS DNS records in a single API call.

## Usage

```hcl
module "dns_batch" {
  source             = "../../examples/ibm-cis-dns-records-bulk"
  cis_instance_name  = "my-cis-instance"
  domain             = "example.com"
  resource_group     = "Default"
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.0 |
| ibm | >= 1.83.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| resource_group | IBM Cloud Resource Group name | string | `"Default"` | no |
| cis_instance_name | Name of the IBM Cloud Internet Services instance | string | n/a | yes |
| domain | DNS domain managed by the CIS instance | string | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| result_posts | DNS records created by the batch post operation |
| result_puts | DNS records replaced by the batch put operation |
| result_patches | DNS records updated by the batch patch operation |
| result_deletes | DNS records removed by the batch delete operation |
