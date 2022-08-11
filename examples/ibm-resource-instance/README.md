# IBM Cloud resource instance example

The following example creates an instance of IBM Cloud resource. Instance could be of any resource, for example Cloud Object Storage, Activity Tracker, metrics monitor etc. By specifying the right value to argument `service`, we can provision respective resource instance.
Document reference http://servicedata.mybluemix.net/

Following types of resources are supported:

* [Resource Instance](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_instance)


## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Create an IBM Cloud Object Storage bucket resource instance. 

```terraform

data "ibm_resource_group" "cos_group" {
  name = var.resource_group_name
}

resource "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = var.service_type
  plan              = var.plan_type
  location          = var.location
}

data "ibm_resource_instance" "test" {
  name    = ibm_resource_instance.cos_instance.name
  service = var.service_type
}

```

## Examples

* [Resource Instance](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-resource-instance)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

You need Terraform v1.0.0 installer to execute this example in your account. For more information, about Terraform installation, see [Installing the Terraform CLI](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-getting-started)

| Name | Version |
|------|---------|
| terraform | >=1.0.0, <2.0 |

## Providers

| Name | Version |
|------|---------|
| ibm | latest |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_name | The name of the service offering. | `string` | yes |
| service_type | The type of the service offering. | `string` | yes |
| plan| The name of the plan type supported by service.| `string` | yes |
| location | Target location or environment to create the resource instance. | `string` | yes |
| resource_group | Name of your resource group. | `string` | yes |

