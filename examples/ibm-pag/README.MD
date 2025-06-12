# Example for Privileged Access Gateway (PAG)

This example illustrates how to provision Privileged Access Gateway Service.

The following type of resource is supported:

* ibm_pag_instance

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## Example Usage

Create an IBM Privileged Access Gateway service.


```hcl
data "ibm_resource_group" "pag" {
  name = var.ibm_resource_group_name
}

data "ibm_resource_instance" "pag-cos" {
  name              = var.ibm_cos_instance_name
  resource_group_id = data.ibm_resource_group.pag.id
  service           = "cloud-object-storage"
}

data "ibm_cos_bucket" "pag-cos-bucket" {
  bucket_name          = var.ibm_cos_bucket_name
  resource_instance_id = data.ibm_resource_instance.pag-cos.id
  bucket_type          = var.ibm_cos_bucket_type
  bucket_region        = var.ibm_cos_bucket_region
}

data "ibm_is_vpc" "pag" {
  name = var.ibm_vpc_name
}

data "ibm_is_subnet" "pag_instance_1" {
  name = var.ibm_vpc_subnet_name_instance_1
}

data "ibm_is_subnet" "pag_instance_2" {
  name = var.ibm_vpc_subnet_name_instance_2
}

data "ibm_is_security_group" "pag_instance_1" {
  name     = each.value
  for_each = var.ibm_vpc_security_groups_instance_1
}

data "ibm_is_security_group" "pag_instance_2" {
  name     = each.value
  for_each = var.ibm_vpc_security_groups_instance_2
}

resource "ibm_pag_instance" "pag" {
  name              = var.ibm_pag_instance_name
  resource_group_id = data.ibm_resource_group.pag.id
  service           = "privileged-access-gateway"
  plan              = var.ibm_pag_service_plan
  location          = var.region
  parameters_json = jsonencode(
    {
      "cosinstance" : data.ibm_resource_instance.pag-cos.crn,
      "cosbucket" : var.ibm_cos_bucket_name,
      "cosendpoint" : data.ibm_cos_bucket.pag-cos-bucket.s3_endpoint_direct
      "proxies" : [
        {
          "name" : "proxy1",
          "securitygroups" : [for sg in data.ibm_is_security_group.pag_instance_1 : sg.id],
          "subnet" : {
            "crn" : data.ibm_is_subnet.pag_instance_1.crn,
            "cidr" : data.ibm_is_subnet.pag_instance_1.ipv4_cidr_block
          }
        },
        {
          "name" : "proxy2",
          "securitygroups" : [for sg in data.ibm_is_security_group.pag_instance_2 : sg.id],
          "subnet" : {
            "crn" : data.ibm_is_subnet.pag_instance_2.crn,
            "cidr" : data.ibm_is_subnet.pag_instance_2.ipv4_cidr_block
          }
        }
      ]
    }
  )
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

resource "ibm_iam_authorization_policy" "pag-cos-iam-policy" {
  source_service_name         = "privileged-access-gateway"
  source_resource_instance_id = ibm_pag_instance.pag.guid
  roles                       = ["Object Writer"]
  resource_attributes {
    name     = "serviceName"
    operator = "stringEquals"
    value    = "cloud-object-storage"
  }

  resource_attributes {
    name     = "accountId"
    operator = "stringEquals"
    value    = data.ibm_resource_group.pag.account_id
  }
  resource_attributes {
    name     = "serviceInstance"
    operator = "stringEquals"
    value    = data.ibm_resource_instance.pag-cos.guid
  }

  resource_attributes {
    name     = "resourceType"
    operator = "stringEquals"
    value    = "bucket"
  }

  resource_attributes {
    name     = "resource"
    operator = "stringEquals"
    value    = var.ibm_cos_bucket_name
  }

}


locals {
  pag_hostnames = [for i in range(var.num_instances) : join(".", ["${ibm_pag_instance.pag.guid}-${i + 1}", "${ibm_pag_instance.pag.location}", "pag", "appdomain", "cloud"])]
}
output "pag-hosts" {
  value = local.pag_hostnames
}
```

## Requirements

You need Terraform v1.0.0 installer to execute this example in your account. For more information, about Terraform installation, see [Installing the Terraform CLI](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-getting-started)

| Name | Version |
|------|---------|
| terraform | >= 1.2.0, < 2.0.0 |

## Providers

| Name | Version |
|------|---------|
| ibm | latest |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| ibmcloud_api_key | IBM Cloud API Key. | `string` | n/a | yes |
| region | The location or the region for your PAG instance. | `string` | `us-south` | yes |
| ibm_pag_instance_name| The name for your PAG instance.| `string` | n/a | yes |
| ibm_resource_group_name | The resource group for your PAG instance. | `string` | n/a | yes |
| ibm_cos_instance_name | Cloud Object Storage instance name where you would like the operator session recordings to be stored. | `string` | n/a | yes |
| ibm_cos_bucket_name | Bucket name where you would like the operator session recordings to be stored. | `string` | n/a | yes |
| ibm_cos_bucket_region | IBM COS Bucket region. | `string` | n/a | yes |
| ibm_cos_bucket_type | IBM COS Bucket type. | `string` | `single_site_location` | yes |
| ibm_vpc_name | VPC name where network interface will reside. | `string` | n/a | yes |
| ibm_pag_service_plan | IBM PAG Service Plan. | `string` | `standard` | yes |
| ibm_vpc_subnet_name_instance_1 | Subnet name for the PAG instance 1.	 | `string` | n/a |  yes |
| ibm_vpc_subnet_name_instance_2 | Subnet name for the PAG instance 2.	 | `string` | n/a |  yes |
| ibm_vpc_security_groups_instance_1 | Security group name for the PAG instance 1.	 | `["<set(string)>"]` | n/a |  yes |
| ibm_vpc_security_groups_instance_2 | Security group name for the PAG instance 2.	 | `["<set(string)>"]` | n/a |  yes |