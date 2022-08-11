# IBM Cloud Hyper Protect Crypto Keys

# Managing HPCS Service Instances using Terraform Resources

This is a collection of resources that make it easier to provision and manage HPCS Instance IBM Cloud Platform:

* Provisioning HPCS Instances - [Hyper Protect Crypto Service instance](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/hpcs) `ibm_hpcs` 
* Managing Keys on HPCS Instance - [ Key Management Service Resource](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/kms_key)


## Terraform versions

Terraform 0.13 and above.

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


```terraform
resource ibm_hpcs hpcs {
  location             = var.location
  name                 = var.hpcs_instance_name
  plan                 = var.plan
  units                = var.units
  signature_threshold  = var.signature_threshold
  revocation_threshold = var.revocation_threshold
  dynamic admins {
    for_each = var.admins
    content {
      name   = admins.value.name
      key   = admins.value.key
      token   = admins.value.token
    }
  }
}
```
### Manage HPCS Keys

```terraform
resource "ibm_kms_key" "key" {
  instance_id  = ibm_hpcs.hpcs.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.13 |

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
| signature_threshold | The number of administrator signatures. | `number` | Yes |
| revocation_threshold | The number of administrator signatures that is required to remove an administrator after you leave imprint mode. | `number` | Yes |
| admins | The list of administrators for the instance crypto units.  | `list` | Yes |
| key\_name | Name of the key. | `string` | Yes |

### Note: Please refer [docs](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/hpcs#admins) for `admins` object details
## Outputs
 Name | Description |
|------|-------------|
| keyID | The ID of the key.|
| InstanceGUID | The GUID of the HPCS Instance.|

