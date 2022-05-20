# IBM Cloud Cloudant - Terraform Module

This is a collection of modules that make it easier to provision cloudant instance and assign service prolicy, create a database:
* [instance](modules/instance)
* [service-policy](modules/service-policy)
* [database](modules/database)

## Compatibility

This module is meant for use with Terraform 0.13 (and higher).

## Usage

Full examples are in the [examples](./examples/) folder, but basic usage is as follows for creation of a Cloudant instance & key:

```hcl
provider "ibm" {
}

data "ibm_resource_group" "cloudant" {
  name = var.resource_group
}

module "cloudant-instance" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/instance"

  source                  = "../../modules/instance"
  provision               = var.provision
  provision_resource_key  = var.provision_resource_key

  instance_name           = var.instance_name
  resource_group_id       = data.ibm_resource_group.cloudant.id
  plan                    = var.plan
  region                  = var.region
  service_endpoints       = var.service_endpoints
  legacy_credentials      = var.legacy_credentials
  tags                    = var.tags
  create_timeout          = var.create_timeout
  update_timeout          = var.update_timeout
  delete_timeout          = var.delete_timeout
  resource_key_name       = var.resource_key_name
  role                    = var.role
  resource_key_tags       = var.resource_key_tags

  ###################
  # Service Policy
  ###################
  service_policy_provision = var.service_policy_provision
  service_name             = var.service_name
  description              = var.description
  roles                    = var.roles
}

```

## Requirements

### Terraform plugins

- [Terraform](https://www.terraform.io/downloads.html) 0.13 (or later)
- [terraform-provider-ibm](https://github.com/IBM-Cloud/terraform-provider-ibm)

## Install

### Terraform

Be sure you have the correct Terraform version (0.13), you can choose the binary here:
- https://releases.hashicorp.com/terraform/

### Terraform plugins

Be sure you have the compiled plugins on $HOME/.terraform.d/plugins/

- [terraform-provider-ibm](https://github.com/IBM-Cloud/terraform-provider-ibm)

### Pre-commit hooks

Run the following command to execute the pre-commit hooks defined in .pre-commit-config.yaml file
```
pre-commit run -a
```
You can install pre-coomit tool using

```
pip install pre-commit
```
or
```
pip3 install pre-commit
```
## How to input varaible values through a file

To review the plan for the configuration defined (no resources actually provisioned)
```
terraform plan -var-file=./input.tfvars
```
To execute and start building the configuration defined in the plan (provisions resources)
```
terraform apply -var-file=./input.tfvars
```

To destroy the VPC and all related resources
```
terraform destroy -var-file=./input.tfvars
```

## Note

All optional parameters, by default, will be set to `null` in respective example's varaible.tf file. You can also override these optional parameters.

