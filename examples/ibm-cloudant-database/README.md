# Example for Cloudant Database

This example illustrates how to use the Cloudant database resources

These types of resources are supported:

* ibm_cloudant_database
* ibm_cloudant_replication

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## Cloudant replication resources

```hcl
module "cloudant-database" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/config-database"

  source                        = "./modules/config-database"
  cloudant_guid                 = var.cloudant_guid
  cloudant_database_partitioned = var.is_partitioned
  db_name                       = var.db_name
  cloudant_database_q           = var.cloudant_database_q

  depends_on = [module.cloudant-instance-dr]
}

module "cloudant-replication" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/config-replication"

  source = "./modules/config-replication"
  ######################
  # Replication Database
  ######################
  cloudant_guid                 = var.cloudant_guid
  cloudant_database_partitioned = var.is_partitioned
  db_name                       = var.db_name
  cloudant_database_q           = var.cloudant_database_q

  #######################
  # Replication Document
  #######################
  cloudant_replication_doc_id = var.cloudant_replication_doc_id
  source_api_key              = var.source_api_key
  target_api_key              = var.target_api_key
  source_host                 = var.source_host
  target_host                 = var.target_host
  create_target               = var.create_target
  continuous                  = var.continuous

  depends_on = [module.cloudant-database-pr]
}
```

## CloudantV1 Data sources


## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.30+ |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| provision | Enable this to bind key to cloudant instance (true/false) | `bool` | true |
| is_dr_provision | Would you like to provision a DR cloudant instance (true/false) | `bool` | true |
| pri_rg_name | Enter resource group name for primary instance | `string` | true |
| dr_rg_name | Enter resource group name for disaster recovery | `string` | true |
| pri_region | Provisioning Region for primary instance | `string` | true |
| dr_region | Provisioning Region for DR instance | `string` | true |
| pri_instance_name | Name of the cloudant instance for primary | `string` | true |
| dr_instance_name | Name of the cloudant instance for DR | `string` | true |
| pri_resource_key | Name of the resource key of the primary instance | `string` | true |
| dr_resource_key | Name of the resource key of the DR | `string` | true |
| legacy_credentials | Legacy authentication method for cloudant | `bool` | false |
| plan | plan type (standard and lite) | `string` | false |
| service_endpoints | Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private' | `string` | false |
| tags | Tags that should be applied to the service | `list` | false |
| service_policy_provision | Enable this to provision the service policy (true/false) | `bool` | true |
| service_name | Name of the service ID | `string` | true |
| description | Description to service ID | `string` | false |
| roles | service policy roles | `list` | false |
| db_name | Database name | `string` | true |
| is_partitioned | To set whether the database is partitioned | `string` | false |
| cloudant_database_q | The number of shards in the database. Each shard is a partition of the hash value range. Default is 8, unless overridden in the `cluster config` | `number` | false |
| cloudant_replication_doc_id | Path parameter to specify the document ID | `string` | true |
| create_target | Creates the target database. Requires administrator privileges on target server | `bool` | false |
| continuous | Configure the replication to be continuous | `bool` | false |