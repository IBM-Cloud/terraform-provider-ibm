---
layout: "ibm"
page_title: "IBM : Cloud Database instance"
sidebar_current: "docs-ibm-resource-database"
description: |-
  Manages IBM Cloud Database Instance.
---

# ibm\_database

Creates an IBM Cloud Database (ICD) instance resource. This resource allows database instances to be created, updated, and deleted. The ibmcloud_api_key used by Terraform must have been granted sufficient IAM rights to create and modify IBM Cloud Databases and have access to the Resource Group the ICD instance will be associated with. See https://cloud.ibm.com/docs/services/databases-for-postgresql/reference-access-management.html#identity-and-access-management for more details on setting IAM and Access Group rights to manage ICD instances.  

If no resource_group_id is specified, the ICD instance is created under the default resource group. The API_KEY must have been assigned permissions for this group.  

Configuration of an ICD resource requires that the `region` parameter is set for the IBM provider in the `provider.tf` to be the same as the target ICD `location/region`. If not specified it will default to `us-south`. A `terraform apply` will fail if the ICD `location` is set differently. If the Terraform configuration needs to deploy resources into multiple regions, provider alias' can be used. https://www.terraform.io/docs/configuration/providers.html#multiple-provider-instances

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_database" "<your_database>" {
  name              = "<your_database_name>"
  plan              = "standard"
  location          = "eu-gb"
  service           = "databases-for-etcd"
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["tag1", "tag2"]

  adminpassword                = "password12"
  members_memory_allocation_mb = 3072
  members_disk_allocation_mb   = 61440
  users {
    name     = "user123"
    password = "password12"
  }
  whitelist {
    address     = "172.168.1.1/32"
    description = "desc"
  }
}

output "ICD Etcd database connection string" {
  value = "http://${ibm_database.test_acc.connectionstrings[0].composed}"
}

```
## Example Usage using point_in_time_recovery time

```hcl
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_database" "test_acc" {
  resource_group_id                    = data.ibm_resource_group.group.id
  name                                 = "<your_database_name>"
  service                              = "databases-for-postgresql"
  plan                                 = "standard"
  location                             = "eu-gb"
  point_in_time_recovery_time          = "2020-04-20T05:27:36Z"
  point_in_time_recovery_deployment_id = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/4448261269a14562b839e0a3019ed980:0b8c37b0-0f01-421a-bb32-056c6565b461::"
}
```

provider.tf

```hcl

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = "eu-gb"
}
```

See https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-database for an example of a VSI configured to connect to a PostgreSQL DB.  


## Timeouts

ibm_database provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for Creating Instance.
* `update` - (Default 20 minutes) Used for Updating Instance.
* `delete` - (Default 10 minutes) Used for Deleting Instance.

ICD instance create typically takes between 10 to 20 minutes. Delete and update in minutes. Provisioning time can be unpredictable. If the apply fails due to a timeout, import the database resource after it has finished creation.  


## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify the database instance. The name must not include spaces. 
* `plan` - (Required, string) The name of the plan type for an IBM Cloud Database. The only currently supported value is "standard"
* `location` - (Required, string) Any of the currently supported ICD regions. The IBM provider `location` in the provider definition also needs to be set to the same region as the target ICD region. The default provider region is `us-south`. The following regions are currently supported: `us-south`, `us-east`, `eu-gb`, `eu-de`, `au-syd`, `jp-tok`, `oslo01`.  
* `resource_group_id` - (Optional, string) The ID of the resource group where you want to create the service. You can retrieve the value from data source `ibm_resource_group`. If not provided it creates the service in default resource group.
* `tags` - (Optional, array of strings) Tags associated with the instance.
* `service` - (Required, string) The ICD database type to be created. Only the following services are currently accepted: 
`databases-for-etcd`, `databases-for-postgresql`, `databases-for-redis`, `databases-for-elasticsearch`, `messages-for-rabbitmq`, `databases-for-mongodb`
* `version` - (Optiona, Forces new resource, string)  The version of the database to be provisioned. If omitted, the database is created with the most recent major and minor version.
* `adminpassword` - (Optional, string) If not specified the password is unitialised and the id unusable. In this case addditional users must be specified in a user block.   
* `members_memory_allocation_mb` - (Optional) The memory size for the database, split across all members. If not specified defaults to the database default. These vary by database type. See the documentation related to each database for the defaults. https://cloud.ibm.com/docs/services/databases-for-postgresql/howto-provisioning.html#list-of-additional-parameters
* `members_disk_allocation_mb`  - (Optional) The disk size of the database, split across all members. As above.
* `members_cpu_allocation_count` - (Optional, int) Enables and allocates the number of specified dedicated cores to your deployment. 
* `backup_id` - (Optional, string) A CRN of a backup resource to restore from. The backup must have been created by a database deployment with the same service ID. The backup is loaded after provisioning and the new deployment starts up that uses that data. A backup CRN is in the format crn:v1:<...>:backup:<uuid>. If omitted, the database is provisioned empty.
* `remote_leader_id` - (Optional, string) A CRN of the leader database to make the replica(read-only) deployment. The leader database must have been created by a database deployment with the same service ID. A read-only replica is set up to replicate all of your data from the leader deployment to the replica deployment using asynchronous replication. See the documentation related to Read-only Replicas here. https://cloud.ibm.com/docs/services/databases-for-postgresql?topic=databases-for-postgresql-read-only-replicas
* `key_protect_key` - (Optional, Force new resource, string) The CRN of a Key Protect key, which is then used for disk encryption. A key protect CRN is in the format crn:v1:<...>:key:<id>. No update support available. `key_protect_key` can be added only at the time of creation.
* `key_protect_instance` - (Optional, Force new resource, string) The CRN of a Key Protect instance, which is then used for disk encryption. A key protect CRN is in the format crn:v1:<...>::.No update support available. `key_protect_instance` can be added only at the time of creation.
* `point_in_time_recovery_deployment_id` - (Optional, string) The source deployment's ID.
* `point_in_time_recovery_time` - (Optional, string) The timestamp in UTC you want to restore to. PITR time stamp can be retrieved using [`ibmcloud cdb postgresql earliest-pitr-timestamp <deployment name or CRN>`] For more info on how to get PITR time refer [point-in-time-recovery-docs](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-pitr)
* `service_endpoints` - (Optional, string) Selects the types Service Endpoints supported on your deployment. Options are public, private, or public-and-private. The default is `public`.

* `users` - (Optional) - Multiple blocks allowed       
  * `name` - Name of the userid to add to the database instance, Minimum of 5 characters up to 32.  
  * `password` - Password for the userid, minimum of 10 characters up to 32. 
            
* `whitelist` - (Optional) - Multiple blocks allowed             
  * `address` - IP address or range of db client addresses to be whitelisted in CIDR format, `172.168.1.2/32`
  * `description` -  Unique description for white list range
* `guid` - Unique identifier of resource instance.



## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new database instance (CRN).
* `status` - Status of resource instance.
* `adminuser` - userid of the default admistration user for the database, usually `admin` or `root`.
* `version` - Database version. 
* `connectionstrings` - List of connection strings by userid for the database. See the IBM Cloud documentation for more details of how to use connection strings in ICD for database access: https://cloud.ibm.com/docs/services/databases-for-postgresql/howto-getting-connection-strings.html#getting-your-connection-strings. The results are returned in pairs of the userid and string:
  `connectionstrings.1.name = admin`
  `connectionstrings.1.string = postgres://admin:$PASSWORD@79226bd4-4076-4873-b5ce-b1dba48ff8c4.b8a5e798d2d04f2e860e54e5d042c915.databases.appdomain.cloud:32554/ibmclouddb?sslmode=verify-full`
Individual string parameters can be retrieved using TF vars and outputs  `connectionstrings.x.hosts.x.port` and `connectionstrings.x.hosts.x.host` 


## Import

The `ibm_database` resource can be imported using the `ID`. The ID is formed from the `CRN` (Cloud Resource Name) from the **Overview** page of the Cloud Database instance. It can be found under the heading **Deployment Details**
* CRN is a 120 digit character string of the form: `crn:v1:bluemix:public:databases-for-postgresql:us-south:a/4ea1882a2d3401ed1e459979941966ea:79226bd4-4076-4873-b5ce-b1dba48ff8c4::`

The `region` parameter must be set for the IBM provider in `provider.tf` to be the same as the ICD service `location(region)`. If not specified it will default to `us-south`. A `terraform refresh/apply` of the data_source will fail if the ICD instance is not in the same region as specified for the provider or its alias.  

```
$ terraform import ibm_database.my_db <crn>

$ terraform import ibm_database.my_db crn:v1:bluemix:public:databases-for-postgresql:us-south:a/4ea1882a2d3401ed1e459979941966ea:79226bd4-4076-4873-b5ce-b1dba48ff8c4::
```

Import requires a minimal Terrform config file to allow importing. 

```hcl
resource "ibm_database" "<your_database>" {
  name              = "<your_database_name>"
```

Run `terraform state show ibm_database.<your_database>` after import to retrieve the additional values to be included in the resource config file. Note that ICD only exports the admin userid. It does not export any additional userids and passwords configured on the instance. These values must be retrieved from an alternative source. If new passwords need to be configured or the connection string retrieved to use the service, a new `users` block must be defined to create new users. This limitation is due to a lack of ICD functionality.  

