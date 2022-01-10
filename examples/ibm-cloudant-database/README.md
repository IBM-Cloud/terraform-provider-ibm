# Example for Cloudant Database

This example illustrates how to use the Cloudant database resources

These types of resources are supported:

* ibm_cloudant
* ibm_cloudant_database

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## Cloudant instance and database

```hcl
module "cloudant-instance" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/instance"

  source        = "./modules/instance"
  instance_name = var.instance_name
  plan          = var.plan
  region        = var.region
}

module "cloudant-database" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/config-database"

  source                        = "./modules/config-database"
  instance_crn                  = module.cloudant-instance.cloudant_instance_crn
  cloudant_database_partitioned = var.is_partitioned
  db_name                       = var.db_name
  cloudant_database_shards      = var.cloudant_database_shards

  depends_on = [module.cloudant-instance]
}

```

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.38+ |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| provision | Enable this to bind key to cloudant instance (true/false) | `bool` | true |
| rg_name | Enter resource group name for the instance | `string` | true |
| region | Provisioning Region for the instance | `string` | true |
| instance_name | Name of the cloudant instance | `string` | true |
| resource_key | Name of the resource key of the instance | `string` | true |
| legacy_credentials | Legacy authentication method for cloudant | `bool` | false |
| plan | plan type (standard and lite) | `string` | false |
| service_endpoints | Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private' | `string` | false |
| tags | Tags that should be applied to the service | `list` | false |
| service_policy_provision | Enable this to provision the service policy (true/false) | `bool` | true |
| service_name | Name of the service ID | `string` | true |
| description | Description to service ID | `string` | false |
| roles | service policy roles | `list` | false |
| db_name | Database name | `string` | true |
| is_partitioned | To set whether the database is partitioned | `bool` | false |
| cloudant_database_shards | The number of shards in the database. Each shard is a partition of the hash value range. Default set by server. | `number` | false |
