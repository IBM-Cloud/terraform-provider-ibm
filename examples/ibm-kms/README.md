# Kms Key example

This example shows how to Create a Key protect instance, generate a key and integrate that key with cos-bucket

This sample configuration will create the key protect instance, cos-bucket instance, root key, and integrate the key with a cos bucket after creating the bucket.

  **Note:**
  
 `key_protect` attribute to associate a kms_key with a COS bucket has been renamed as `kms_key_crn` , hence it is recommended to all the new users to use `kms_key_crn`.Although the support for older attribute name `key_protect` will be continued for existing custom

To run, configure your IBM Cloud provider

These types of resources and datasources are supported:

* [ Import keys ](https://cloud.ibm.com/docs/terraform?topic=terraform-kms-data-sources)
* [ Manage keys ](https://cloud.ibm.com/docs/terraform?topic=terraform-kms-resources)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.10.0`. Branch - `master`.


## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

`Create key-protect root key using kms resource and Import keys into cos`:
```hcl

resource "ibm_resource_instance" "cos_instance" {
  name     = var.cos_name
  service  = "cloud-object-storage"
  plan     = var.pln
  location = var.location
}
resource "ibm_resource_instance" "kms_instance" {
  name     = var.kms_name
  service  = "kms"
  plan     = "tiered-pricing"
  location = var.kms_location
}
resource "ibm_kms_key" "test" {
  instance_id = ibm_resource_instance.kms_instance.guid
  key_name       = var.key_name
  standard_key   = false
}

resource "ibm_iam_authorization_policy" "policy" {
  source_service_name = "cloud-object-storage"
  target_service_name = "kms"
  roles               = ["Reader"]
}

resource "ibm_cos_bucket" "flex-us-south" {
  depends_on           = [ibm_iam_authorization_policy.policy]
  bucket_name          = var.bucket_name
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = var.location
  storage_class        = "flex"
  kms_key_crn          = ibm_kms_key.test.id
}

```

`Import existing keys`:

```hcl

data "ibm_kms_keys" "test" {
    instance_id = var.kmskeyid
}

resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = var.bucket_name
  resource_instance_id = var.cosinstance
  region_location      = var.location
  storage_class        = "flex"
  kms_key_crn         = data.ibm_kms_keys.test.keys.0.crn
}

```
## Assumptions

1. It's assumed that user has valid authorizations set for integrating kms keys with other services. This can be done using `ibm_iam_authorization_policy` resource
2. The kms service instance must be initialized before managing kms keys with kms resources. 
3. Before doing terraform destroy if force_delete flag is introduced after provisioning keys, a terraform apply must be done before terraform destroy for force_delete flag to take effect.
4. [ API Documentation for kms ](https://cloud.ibm.com/apidocs/key-protect).
5. There can be multiple key protect keys with same name in a kms instance, all keys can be imported as a list of keys.

## Notes

1. Before doing terraform destroy if force_delete flag is introduced after provisioning keys, a terraform apply must be done before terraform destroy for force_delete flag to take effect.
2. KMIP adapters with active KMIP objects cannot be deleted by Terraform.

## Examples

* [ KMS Examples ](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-kms)



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
| kms\_plan | The key protect plan to provision| `string` | yes |
| kms\_name_ | The name of the keyprotect instance| `string` | yes |
| key\_name | The name of the kms key. | `string` | yes |
| standard\_key | Set to true to create a standard key, to create a root key set this flag to false. Default: `false` . | `bool` | no |
| plan | The cos instance plan to provision| `string` | yes |
| kms\_location | The location where key protect instance will be created| `string` | yes |
| location | The location where cos instance will be created| `string` | yes |
| cos\_name | The name of the cos instance to be provisioned| `string` | yes |
| cos\_bucket_name | The name of the cos ibucket| `string` | yes |

## Outputs

| Name | Description |
|------|-------------|
| key\_value | The list of all keys present in the instance with their properties  |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
