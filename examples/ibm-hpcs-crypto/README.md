# IBM Cloud Hyper Protect Crypto Keys

Pre-req : Provision a IBM Hyper Protect Crypto Service and Intialize the service.

The following example creates a new key with specified key material.

Following types of resources are supported:

* [ Key Management Service Resource](https://cloud.ibm.com/docs/terraform?topic=terraform-kp-resources#kms-key)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.10.1`. Branch - `master`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Creates a new key with specified key material. 

```hcl

resource "ibm_kms_key" "hpcstest" {
	instance_id = "5af62d5d-5d90-4b84-bbcd-90d2123ae6c8"
	key_name = "mystandardkey"
	standard_key =  true
	force_delete = true
}

```


## Examples

* [ Key Management Service  ](https://cloud.ibm.com/docs/terraform?topic=terraform-kp-resources#kms-key)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| key\_name | Name of the key. | `string` | No |
| hpcs\_instance\_guid |GUID of the hpcs service instance. | `string` | yes |


 Name | Description |
|------|-------------|
| keyID | The ID of the key.|

