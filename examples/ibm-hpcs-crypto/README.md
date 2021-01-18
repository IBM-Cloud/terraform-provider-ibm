# IBM Cloud Hyper Protect Crypto Keys

# Managing HPCS Service Instances using Terraform Resources

This is a collection of resources that make it easier to provision and manage HPCS Instance IBM Cloud Platform:

* Provisioning HPCS Instances - `ibm_resource_instance`
* [Initialising HPCS Instance](https://github.com/terraform-ibm-modules/terraform-ibm-hpcs) - Initialises and Configured zeroised crypto units.
* Managing Keys on HPCS Instance - [ Key Management Service Resource](https://cloud.ibm.com/docs/terraform?topic=terraform-kp-resources#kms-key)


## Terraform versions

Terraform 0.12.

## Usage

Full example is in [main.tf](main.tf)

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

### Provision HPCS Instance

Note: `provision_instance` will determine if the instance has to be provisioned or not. If `provision_instance` is true, count will be 1 and the instance will be provisioned..
```hcl
resource "ibm_resource_instance" "hpcs_instance" {
  count    = (var.provision_instance == true ? 1 : 0)
  name     = var.hpcs_instance_name
  service  = "hs-crypto"
  plan     = var.plan
  location = var.location
  parameters = {
    units = var.units
  }
}
```



### Manage HPCS Keys
`Note:` To Manage Keys, Instance should be Initialized..

```hcl
resource "ibm_kms_key" "key" {
  instance_id  = data.ibm_resource_instance.hpcs_instance.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
}
```

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
| provision_instance | Determines if the instance has to be created or not. | `bool` | Yes |
| hpcs_instance_name | Name of HPCS Instance. | `string` | Yes |
| location | Location of HPCS Instance. | `string` | Yes |
| plan | Plan of HPCS Instance.Default: `standard` | `string` | No |
| units | No of crypto units that has to be attached to the instance. | `number` | Yes |
| key\_name | Name of the key. | `string` | Yes |

Note: COS Credententials are required when `download_from_cos` and `upload_to_cos` null resources are used

 Name | Description |
|------|-------------|
| keyID | The ID of the key.|
| InstanceGUID | The GUID of the HPCS Instance.|

