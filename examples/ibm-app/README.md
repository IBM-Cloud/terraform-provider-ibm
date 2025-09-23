# Example for app

This example shows how to deploy an application in the IBM PaaS. In the variables.tf you would find git_repo which is git url of a Cloud Foundry application repository. You must provide valid values for the variables org and space.
When you perform terraform apply the provisioner will download the code from the git_repo and zip it at location specified by variable app_zip.

The example provisions a cloudant db service instance, create routes and assigns that route and service instance to the application.



These types of resources are supported:

* [ app ](https://cloud.ibm.com/docs/terraform?topic=terraform-cloud-foundry-resources#cf-app)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.4.0`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.25.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## Cloud Foundry ibm app

`create app`:
```hcl

data "ibm_space" "space" {
  org   = "example.com"
  space = "dev"
}

resource "ibm_app" "app" {
  name                 = var.name
  space_guid           = data.ibm_space.space.id
  app_path             = var.path
  wait_timeout_minutes = 90
  buildpack            = "sdk-for-nodejs"
}

```

## Examples

* [ app ](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-app)


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

| Name | Description                                  | Type     | Required |
|------|----------------------------------------------|----------|----------|
| name | The name of the app that you want to create. | `string` | yes      |
| path | path to the compressed file of the app.      | `string` | yes      |

## Outputs

| Name     | Description                       |
|----------|-----------------------------------|
| id       | The unique identifier of the app. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
