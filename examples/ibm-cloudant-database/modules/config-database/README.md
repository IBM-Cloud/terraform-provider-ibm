# Module Cloudant Database

This module is used to create cloudant instance and create primary & disaster recovery database.

## Example Usage
```
provider "ibm" {
}

data "ibm_resource_group" "pri_group" {
  name = var.pri_rg_name
}

data "ibm_resource_group" "dr_group" {
  name = var.dr_rg_name
}

module "cloudant-database" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/database"

  source                = "../../modules/database"
  bind_resource_key     = var.bind_resource_key
  provision             = var.provision
  pri_resource_group_id = data.ibm_resource_group.pri_group.id
  dr_resource_group_id  = data.ibm_resource_group.dr_group.id
  pri_instance_name     = var.pri_instance_name
  dr_instance_name      = var.dr_instance_name
  pri_region            = var.pri_region
  dr_region             = var.dr_region
  db_name               = var.db_name
  plan                  = var.plan
  pri_resource_key      = var.pri_resource_key
  dr_resource_key       = var.dr_resource_key
  role                  = var.role
  resource_key_tags     = var.resource_key_tags
  is_partitioned        = var.is_partitioned
  service_endpoints     = var.service_endpoints
  parameters            = var.parameters
  tags                  = var.tags
  create_timeout        = var.create_timeout
  update_timeout        = var.update_timeout
  delete_timeout        = var.delete_timeout
}

```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs


| Name                 | Description                                                                     | Type         | Default | Required |
|----------------------|---------------------------------------------------------------------------------|:-------------|:------- |:---------|
| bind_resource_key    | Indicating that instance key should be bind to cloudant instance                | bool         | true    | no       |
| provision            | Indicating to provision a new cloudant instance                                 | bool         | true    | no       |
| pri\_instance\_name  | A descriptive name used to identify the resource instance for primary           | string       | n/a     | yes      |
| dr\_instance\_name   | A descriptive name used to identify the resource instance for disaster recovery | string       | n/a     | yes      |
| pri\_resource\_key   | A descriptive name used to identify the resource key for primary                | string       | n/a     | yes      |
| dr\_resource\_key    | A descriptive name used to identify the resource key for disaster recovery      | string       | n/a     | yes      |
| db_name              | Database name                                                                   | string       | n/a     | yes      |
| is_partitioned       | Enable partition database                                                       | bool         | false   | no       |
| role                 | Name of the user role.                                                          | string       | n/a     | yes      |
| plan                 | The name of the plan type supported by service.                                 | string       | n/a     | yes      |
| region               | Target location or environment to create the resource instance.                 | string       | n/a     | yes      |
| pri\_resource\_group | Name of the primary db resource group                                           | string       | n/a     | yes      |
| dr\_resource\_group  | Name of the disaster recovery db  resource group                                | string       | n/a     | yes      |
| service\_endpoints   | Possible values are 'public', 'private', 'public-and-private'.                  | string       | n/a     | no       |
| tags                 | Tags that should be applied to the service                                      | list(string) | n/a     | no       |
| resource_key_tags    | Tags that should be applied to the service key                                  | list(string) | n/a     | no       |
| parameters           | Arbitrary parameters to pass                                                    | map(string)  | n/a     | no       |
| create_timeout       | Timeout duration for create                                                     | string       | n/a     | no       |
| update_timeout       | Timeout duration for update                                                     | string       | n/a     | no       |
| delete_timeout       | Timeout duration for delete                                                     | string       | n/a     | no       |

NOTE: We can set the create, update and delete timeouts as string. For e.g say we want to set 15 minutes timeout then the value should be "15m".

## Outputs

| Name                    | Description                            |
|-------------------------|----------------------------------------|
| pri_cloudant_key_id     | ID of the primary cloudant key         |
| pri_cloudant_guid       | GUID of the primary cloudant instance  |
| pri_cloudant_id         | ID of the primary cloudant instance    |
| dr_cloudant_key_id      | ID of the dr cloudant key              |
| dr_cloudant_guid        | GUID of the dr cloudant instance       |
| dr_cloudant_id          | ID of the dr cloudant instance         |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->