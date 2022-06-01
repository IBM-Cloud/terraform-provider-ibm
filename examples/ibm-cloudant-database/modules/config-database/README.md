# Module Cloudant Database

This module is used to create cloudant database.

## Example Usage
```
module "cloudant-database" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/database"

  source                         = "../../modules/database"
  db_name                        = var.db_name
  cloudant_instance_crn          = var.instance_crn
  cloudant_database_partitioned  = var.is_partitioned
}

```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs


| Name                          | Description                                                                     | Type         | Default | Required |
|-------------------------------|---------------------------------------------------------------------------------|:-------------|:------- |:---------|
| db_name                       | Database name                                                                   | string       | n/a     | yes      |
| cloudant_instance_crn         | CRN of the cloudant instance.                                                   | string       | n/a     | yes      |
| cloudant_database_partitioned | Enable partition database                                                       | bool         | false   | no       |
| cloudant_database_shards      | The number of shards in the database. Default set by server.                    | number       | null    | no       |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->