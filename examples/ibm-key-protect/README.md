# Key Protect Key example

This example shows how to Create a Key protect instance, generate a key and integrate that key with cos-bucket

This sample configuration will create the key protect instance, cos-bucket instance, root key, and integrate the key with a cos bucket after creating the bucket.

To run, configure your IBM Cloud provider

These types of resources and datasources are supported:

* [ Import keys ](https://cloud.ibm.com/docs/terraform?topic=terraform-kp-data-sources)
* [ Manage keys ](https://cloud.ibm.com/docs/terraform?topic=terraform-kp-resources)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.4.0`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.25.0`. Branch - `terraform_v0.11.x`.

## Deprecation Notice

  The resource `ibm_kp_key` is deprecated and replaced with `ibm_kms_key`. 

  Please refer to [https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-kms](examples/ibm-kms) for examples.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

`Create key-protect root key using kp resource and Import keys into cos`:
```hcl

resource "ibm_resource_instance" "cos_instance" {
  name     = var.cos_name
  service  = "cloud-object-storage"
  plan     = var.pln
  location = var.location
}
resource "ibm_resource_instance" "kp_instance" {
  name     = var.kp_name
  service  = "kms"
  plan     = "tiered-pricing"
  location = var.kp_location
}
resource "ibm_kp_key" "test" {
  key_protect_id = ibm_resource_instance.kp_instance.guid
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
  kms_key_crn          = ibm_kp_key.test.id
}

```

`Import existing keys`:

```hcl

data "ibm_kp_key" "test" {
    key_protect_id = var.kpkeyid
}

resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = var.bucket_name
  resource_instance_id = var.cosinstance
  region_location      = var.location
  storage_class        = "flex"
  kms_key_crn         = data.ibm_kp_key.test.keys.0.crn
}

```

## Assumptions

1. It's assumed that user has valid authorizations set for integrating kp keys with other services. This can be done using `ibm_iam_authorization_policy` resource
2. The kms service instance must be initialized before managing kp keys with kp resources. 
3. Before doing terraform destroy if force_delete flag is introduced after provisioning keys, a terraform apply must be done before terraform destroy for force_delete flag to take effect.
4. [ API Documentation for kms ](https://cloud.ibm.com/apidocs/key-protect).
5. There can be multiple key protect keys with same name in a kp instance, all keys can be imported as a list of keys.

## Notes

1. Before doing terraform destroy if force_delete flag is introduced after provisioning keys, a terraform apply must be done before terraform destroy for force_delete flag to take effect.

## Examples

* [ Key Protect Examples ](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-key-protect)



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
| kp\_plan | The key protect plan to provision| `string` | yes |
| kp\_name_ | The name of the keyprotect instance| `string` | yes |
| key\_name | The name of the kp key. | `string` | yes |
| standard\_key | Set to true to create a standard key, to create a root key set this flag to false. Default: `false` . | `bool` | no |
| plan | The cos instance plan to provision| `string` | yes |
| kp\_location | The location where key protect instance will be created| `string` | yes |
| location | The location where cos instance will be created| `string` | yes |
| cos\_name | The name of the cos instance to be provisioned| `string` | yes |
| cos\_bucket_name | The name of the cos ibucket| `string` | yes |

## Outputs

| Name | Description |
|------|-------------|
| key\_value | The list of all keys present in the instance with their properties  |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
