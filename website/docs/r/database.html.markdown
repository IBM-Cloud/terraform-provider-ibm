---
subcategory: "Cloud Databases"
layout: "ibm"
page_title: "IBM : Cloud Database instance"
description: |-
  Manages IBM Cloud database instance.
---

# ibm_database

Create, update, or delete a IBM Cloud Database (ICD) instance. The `ibmcloud_api_key` that are used by  Terraform should grant IAM rights to create and modify IBM Cloud Databases and have access to the resource group the ICD instance is associated with. For more information, see [documentation](https://cloud.ibm.com/docs/services/databases-for-postgresql/reference-access-management.html#identity-and-access-management) to manage ICD instances.

If `resource_group_id` is not specified, the ICD instance is created in the default resource group. The `API_KEY` must be assigned permissions for this group.

Configuration of an ICD resource requires that the `region` parameter is set for the IBM provider in the `provider.tf` to be the same as the target ICD `location/region`. If not specified it default to `us-south`. A `terraform apply`  fails if the ICD `location` is set differently. If the Terraform configuration needs to deploy resources into multiple regions, provider alias can be used. For more information, see [Terraform provider configuration](https://www.terraform.io/docs/configuration/providers.html#multiple-provider-instances).


## Example usage
To find an example for configuring a virtual server instance that connects to a PostgreSQL database, see [here](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-database).

```terraform
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

### Sample database instance by using `node_` attributes
An example to configure and deploy database by using `node_` attributes instead of `memory_`.

```terraform
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
  node_count                = 3
  node_memory_allocation_mb = 1024
  node_disk_allocation_mb   = 20480
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

### Sample database instance by using `group` attributes
An example to configure and deploy database by using `group` attributes.

```terraform
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

  group {
    group_id = "member"

    memory {
      allocation_mb = 1024
    }

    disk {
      allocation_mb = 5120
    }

    cpu {
      allocation_count = 3
    }
  }

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

### Sample database instance by using `point_in_time_recovery`
An example for configuring `point_in_time_recovery` time by using `ibm_database` resource.


```terraform
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


### Sample database instance by using auto_scaling

```terraform
resource "ibm_database" "autoscale" {
    resource_group_id            = data.ibm_resource_group.group.id
    name                         = "redis"
    service                      = "databases-for-redis"
    plan                         = "standard"
    location                     = "us-south"
    service_endpoints            = "private"
    auto_scaling {
        cpu {
            rate_increase_percent       = 20
            rate_limit_count_per_member = 20
            rate_period_seconds         = 900
            rate_units                  = "count"
        }
        disk {
            capacity_enabled             = true
            free_space_less_than_percent = 15
            io_above_percent             = 85
            io_enabled                   = true
            io_over_period               = "15m"
            rate_increase_percent        = 15
            rate_limit_mb_per_member     = 3670016
            rate_period_seconds          = 900
            rate_units                   = "mb"
        }
         memory {
            io_above_percent         = 90
            io_enabled               = true
            io_over_period           = "15m"
            rate_increase_percent    = 10
            rate_limit_mb_per_member = 114688
            rate_period_seconds      = 900
            rate_units               = "mb"
        }
    }
}
```
### Sample Cassandra database instance
* Cassandra provisioning may require more time than the default timeout. A longer timeout value can be set with using the `timeouts` attribute.

```terraform
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "cassandra" {
  resource_group_id            = data.ibm_resource_group.test_acc.id
  name                         = "test"
  service                      = "databases-for-cassandra"
  plan                         = "enterprise"
  location                     = "us-south"
  adminpassword                = "password12"
  members_memory_allocation_mb = 36864
  members_disk_allocation_mb   = 61440
  users {
    name     = "user123"
    password = "password12"
  }
  whitelist {
    address     = "172.168.1.2/32"
    description = "desc1"
  }

  timeouts {
    create = "120m"
    update = "120m"
    delete = "15m"
  }
}
```
### Sample MongoDB Enterprise database instance
* MongoDB Enterprise provisioning may require more time than the default timeout. A longer timeout value can be set with using the `timeouts` attribute.
* Please make sure your resources meet minimum requirements of scaling. Please refer [docs](https://cloud.ibm.com/docs/databases-for-mongodb?topic=databases-for-mongodb-pricing#scaling-per-member) for more info.
* `service_endpoints` cannot be updated on this instance.

```terraform
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "mongodb" {
  resource_group_id            = data.ibm_resource_group.test_acc.id
  name                         = "test"
  service                      = "databases-for-mongodb"
  plan                         = "enterprise"
  location                     = "us-south"
  adminpassword                = "password12"
  members_disk_allocation_mb   = 61440
  members_memory_allocation_mb = 43008
  tags                         = ["one:two"]
  users {
    name     = "user123"
    password = "password12"
  }
  whitelist {
    address     = "172.168.1.2/32"
    description = "desc1"
  }
  timeouts {
    create = "120m"
    update = "120m"
    delete = "15m"
  }
}
```

### Sample MongoDB Enterprise database instance with BI Connector and Analytics
* To enable Analytics and/or BI Connector for MongoDB Enterprise, a `group` attribute must be defined for the `analytics` and `bi_connector` group types with `members` scaled to at exactly `1`.
* MongoDB Enterprise provisioning may require more time than the default timeout. A longer timeout value can be set with using the `timeouts` attribute.

```terraform
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "mongodb_enterprise" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "test"
  service           = "databases-for-mongodb"
  plan              = "enterprise"
  location          = "us-south"
  adminpassword     = "password12"
  tags              = ["one:two"]

  group {
    group_id = "member"
    
    memory { 
      allocation_mb = 14336
    }
    
    disk { 
      allocation_mb = 20480
    }
  }
  
  group {
    group_id = "analytics"
    
    members { 
      allocation_count = 1
    }
  }
  
  group {
    group_id = "bi_connector"
    
    members { 
      allocation_count = 1
    }
  }
    
  timeouts {
    create = "120m"
    update = "120m"
    delete = "15m"
  }
}

data "ibm_database_connection" "mongodb_conn" {
  deployment_id = ibm_database.mongodb_enterprise.id
  user_type     = "database"
  user_id       = "admin"
  endpoint_type = "public"
}

output "bi_connector_connection" {
  description = "BI Connector connection string"
  value       = data.ibm_database_connection.mongodb_conn.bi_connector.0.composed.0
}

output "analytics_connection" {
  description = "Analytics Node connection string"
  value       = data.ibm_database_connection.mongodb_conn.analytics.0.composed.0
}

```

### Sample EDB instance
EDB takes more time than expected. It is always advisible to extend timeouts using timeouts block

```terraform
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "edb" {
  resource_group_id            = data.ibm_resource_group.test_acc.id
  name                         = "test"
  service                      = "databases-for-enterprisedb"
  plan                         = "standard"
  location                     = "us-south"
  adminpassword                = "password12"
  members_memory_allocation_mb = 3072
  members_disk_allocation_mb   = 61440
  tags                         = ["one:two"]
  users {
    name     = "user123"
    password = "password12"
  }
  whitelist {
    address     = "172.168.1.2/32"
    description = "desc1"
  }
  timeouts {
    create = "120m"
    update = "120m"
    delete = "15m"
  }
}
```
### Updating configuration for postgres database

```terraform
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "db" {
  location                     = "us-east"
  members_cpu_allocation_count = 0
  members_disk_allocation_mb   = 10240
  members_memory_allocation_mb = 2048
  name                         = "telus-database"
  service                      = "databases-for-postgresql"
  plan                         = "standard"
  configuration           		= <<CONFIGURATION
  {
    "max_connections": 400
  }
  CONFIGURATION
}

```

**provider.tf**
Please make sure to target right region in the provider block, If database is created in region other than `us-south`

```terraform
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = "eu-gb"
}
```


For more information, about an example that are related to a VSI configuration to connect to a PostgreSQL database, refer to [VSI configured connection](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-database).


## Timeouts
The following timeouts are defined for this resource.

* `Create` The creation of an instance is considered failed when no response is received for 60 minutes.
* `Update` The update of an instance is considered failed when no response is received for 20 minutes.
* `Delete` The deletion of an instance is considered failed when no response is received for 10 minutes.

ICD create instance typically takes between 30 minutes to 45 minutes. Delete and update takes a minute. Provisioning time are unpredictable, if the apply fails due to a timeout, import the database resource once the create is completed.


## Argument reference
Review the argument reference that you can specify for your resource.

- `adminpassword` - (Optional, String)  The password for the database administrator. If not specified, an empty string is provided for the password and the user ID cannot be used. In this case, more users must be specified in a `user` block.
- `auto_scaling` (List , Optional) Configure rules to allow your database to automatically increase its resources. Single block of autoscaling is allowed at once.

  Nested scheme for `auto_scaling`:
  - `cpu` (List , Optional) Single block of CPU is allowed at once by CPU autoscaling.

    Nested scheme for `cpu`:
    - `rate_increase_percent` - (Optional, Integer) Auto scaling rate in increase percent.
    - `rate_limit_count_per_member` - (Optional, Integer) Auto scaling rate limit in count per number.
    - `rate_period_seconds` - (Optional, Integer) Period seconds of the auto scaling rate.
    - `rate_units` - (Optional, String) Auto scaling rate in units.
  - `disk` (List , Optional) Single block of disk is allowed at once in disk auto scaling.

    Nested scheme for `disk`:
    - `capacity_enabled` - (Optional, Bool) Auto scaling scalar enables or disables the scalar capacity.
    - `free_space_less_than_percent` - (Optional, Integer) Auto scaling scalar capacity free space less than percent.
    -  `io_above_percent` - (Optional, Integer) Auto scaling scalar I/O utilization above percent.
    - `io_enabled` - (Optional, Bool) Auto scaling scalar I/O utilization enabled.`
    - `rate_increase_percent` - (Optional, Integer) Auto scaling rate increase percent.
    - `rate_limit_mb_per_member` - (Optional, Integer) Auto scaling rate limit in megabytes per member.
    - `rate_period_seconds` - (Optional, Integer) Auto scaling rate period in seconds.
    - `rate_units` - (Optional, String) Auto scaling rate in units.
  - `memory` (List , Optional) Memory Auto Scaling in single block of memory is allowed at once.

    Nested scheme for `memory`:
    - `io_above_percent` - (Optional, Integer) Auto scaling scalar I/O utilization above percent.
    - `io_enabled`-Bool-Optional-Auto scaling scalar I/O utilization enabled.
    - `io_over_period` - (Optional, String) Auto scaling scalar I/O utilization over period.
    - `rate_increase_percent` - (Optional, Integer) Auto scaling rate in increase percent.
    - `rate_limit_mb_per_member` - (Optional, Integer) Auto scaling rate limit in megabytes per member.
    - `rate_period_seconds` - (Optional, Integer) Auto scaling rate period in seconds.
    - `rate_units` - (Optional, String) Auto scaling rate in units.
- `backup_id` - (Optional, String) The CRN of a backup resource to restore from. The backup is created by a database deployment with the same service ID. The backup is loaded after provisioning and the new deployment starts up that uses that data. A backup CRN is in the format `crn:v1:<…>:backup:`. If omitted, the database is provisioned empty.
- `backup_encryption_key_crn`- (Optional, Forces new resource, String) The CRN of a key protect key, that you want to use for encrypting disk that holds deployment backups. A key protect CRN is in the format `crn:v1:<...>:key:`. Backup_encryption_key_crn can be added only at the time of creation and no update support  are available.
- `configuration` - (Optional, Json String) Database Configuration in JSON format. Supported services `databases-for-postgresql`, `databases-for-redis` and `databases-for-enterprisedb`. For valid values please refer [API docs](https://cloud.ibm.com/apidocs/cloud-databases-api/cloud-databases-api-v4#setdatabaseconfiguration-request).
- `guid` - (Optional, String) The unique identifier of the database instance.
- `key_protect_key` - (Optional, Forces new resource, String) The root key CRN of a Key Management Services like Key Protect or Hyper Protect Crypto Service (HPCS)  that you want to use for disk encryption. A key CRN is in the format `crn:v1:<…>:key:`. You can specify the root key during the database creation only. After the database is created, you cannot update the root key. For more information, refer [Disk encryption](https://cloud.ibm.com/docs/cloud-databases?topic=cloud-databases-key-protect#using-the-key-protect-key) documentation.
- `key_protect_instance` - (Optional, Forces new resource, String) The instance CRN of a Key Management Services like Key Protect or Hyper Protect Crypto Service (HPCS) that you want to use for disk encryption. An instance CRN is in the format `crn:v1:<…>::`.
- `location` - (Required, String) The location where you want to deploy your instance. The location must match the `region` parameter that you specify in the `provider` block of your  Terraform configuration file. The default value is `us-south`. Currently, supported regions are `us-south`, `us-east`, `eu-gb`, `eu-de`, `au-syd`, `jp-tok`, `oslo01`.
- `group` - (Optional, Set) A set of group scaling values for the database. Multiple blocks are allowed. Can only be performed on is_adjustable=true groups. Values set are per-member. Values must be greater than or equal to the minimum size and must be a multiple of the step size.

  Nested scheme for `group`:
  - `group_id` - (Optional, String) The ID of the scaling group.

  - `members` (Set, Optional)

    Nested scheme for `members`:
    - `allocation_count` - (Optional, Integer) Allocated number of members.

  - `memory` (Set, Optional) Memory Auto Scaling in single block of memory is allowed at once.

    Nested scheme for `memory`:
    - `allocation_mb` - (Optional, Integer) Allocated memory per-member.

  - `disk` (Set, Optional)

    Nested scheme for `disk`:
    - `allocation_mb` - (Optional, Integer) Allocated disk per-member.

  - `cpu` (Set, Optional)

    Nested scheme for `cpu`:
    - `allocation_count` - (Optional, Integer) Allocated dedicated CPU per-member.

- `members_memory_allocation_mb` - (Optional, Integer) The amount of memory in megabytes for the database, split across all members. If not specified, the default setting of the database service is used, which can vary by database type.
- `members_disk_allocation_mb` - (Optional, Integer) The amount of disk space for the database, split across all members. If not specified, the default setting of the database service is used, which can vary by database type.
- `members_cpu_allocation_count` - (Optional, Integer) Enables and allocates the number of specified dedicated cores to your deployment.
- `node_count` - (Optional, Integer) The total number of nodes in the cluster. If not specified defaults to the database minimum node count. These vary by database type. See the documentation related to each database for the defaults. https://cloud.ibm.com/docs/services/databases-for-postgresql/howto-provisioning.html#list-of-additional-parameters
- `node_cpu_allocation_count` - (Optional, Integer) Enables and allocates the number of specified dedicated cores to your deployment per node.
- `node_disk_allocation_mb`  - (Optional, Integer) The disk size of the database per node. As above.
- `node_memory_allocation_mb` - (Optional,Integer) The memory size for the database per node. If not specified defaults to the database default. These vary by database type. See the documentation related to each database for the defaults. https://cloud.ibm.com/docs/services/databases-for-postgresql/howto-provisioning.html#list-of-additional-parameters

  ~> **Note:** `members_memory_allocation_mb`, `members_disk_allocation_mb`, `members_cpu_allocation_count` conflicts with `node_count`,`node_cpu_allocation_count`, `node_disk_allocation_mb`, `node_memory_allocation_mb`. `group` conflicts with `node_` and `members_` arguments. Either members, node, or group arguments have to be provided.
- `name` - (Required, String) A descriptive name that is used to identify the database instance. The name must not include spaces.
- `plan` - (Required, Forces new resource, String) The name of the service plan that you choose for your instance. All databases use `standard`. `enterprise` is supported only for cassandra (`databases-for-cassandra`) and mongodb(`databases-for-mongodb`)
* `plan_validation` - (Optional, bool) Enable or disable validating the database parameters for elasticsearch and postgres (more coming soon) during the plan phase. If not specified defaults to true.
- `point_in_time_recovery_deployment_id` - (Optional, String) The ID of the source deployment that you want to recover back to.
- `point_in_time_recovery_time` - (Optional, String) The timestamp in UTC format that you want to restore to. To retrieve the timestamp, run the `ibmcloud cdb postgresql earliest-pitr-timestamp <deployment name or CRN>` command. For more information, see [Point-in-time Recovery](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-pitr).
- `remote_leader_id` - (Optional, String) A CRN of the leader database to make the replica(read-only) deployment. The leader database is created by a database deployment with the same service ID. A read-only replica is set up to replicate all of your data from the leader deployment to the replica deployment by using asynchronous replication. For more information, see [Configuring Read-only Replicas](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-read-only-replicas).
- `resource_group_id` - (Optional, Forces new resource, String)  The ID of the resource group where you want to create the instance. To retrieve this value, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If no value is provided, the `default` resource group is used.
- `service` - (Required, Forces new resource, String) The type of Cloud Databases that you want to create. Only the following services are currently accepted: `databases-for-etcd`, `databases-for-postgresql`, `databases-for-redis`, `databases-for-elasticsearch`, `messages-for-rabbitmq`,`databases-for-mongodb`,`databases-for-mysql`, `databases-for-cassandra` and `databases-for-enterprisedb`.
- `service_endpoints` - (Optional, String) Specify whether you want to enable the public, private, or both service endpoints. Supported values are `public`, `private`, or `public-and-private`. The default is `public`.
- `tags` (Optional, Array of Strings) A list of tags that you want to add to your instance.
- `version` - (Optional, Forces new resource, String) The version of the database to be provisioned. If omitted, the database is created with the most recent major and minor version.
- `users` - (Optional, List of Objects) A list of users that you want to create on the database. Multiple blocks are allowed.

  Nested scheme for `users`:
  - `name` - (Optional, String) The user ID to add to the database instance. The user ID must be in the range 5 - 32 characters.
  - `password` - (Optional, String) The password for the user ID. The password must be in the range 10 - 32 characters.
- `whitelist` - (Optional, List of Objects) A list of allowed IP addresses for the database. Multiple blocks are allowed.

  Nested scheme for `whitelist`:
  - `address` - (Optional, String) The IP address or range of database client addresses to be whitelisted in CIDR format. Example, `172.168.1.2/32`.
  - `description` - (Optional, String) A description for the allowed IP addresses range.


## Attribute reference
In addition to all argument references list, you can access the following attribute references after your resource is created.

- `adminuser` - (String) The user ID of the database administrator. Example, `admin` or `root`.
- `configuration_schema` (String) Database Configuration Schema in JSON format.
- `connectionstrings` - (Array) A list of connection strings for the database for each user ID. For more information, about how to use connection strings, see the [documentation](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-connection-strings). The results are returned in pairs of the userid and string: `connectionstrings.1.name = admin connectionstrings.1.string = postgres://admin:$PASSWORD@79226bd4-4076-4873-b5ce-b1dba48ff8c4.b8a5e798d2d04f2e860e54e5d042c915.databases.appdomain.cloud:32554/ibmclouddb?sslmode=verify-full` Individual string parameters can be retrieved by using  Terraform variables and outputs `connectionstrings.x.hosts.x.port` and `connectionstrings.x.hosts.x.host`.
- `id` - (String) The CRN of the database instance.
- `status` - (String) The status of the instance.
- `version` - (String) The database version.

## Import
The database instance can be imported by using the ID, that is formed from the CRN. To import the resource, you must specify the `region` parameter in the `provider` block of your  Terraform configuration file. If the region is not specified, `us-south` is used by default. An  Terraform refresh or apply fails, if the database instance is not in the same region as configured in the provider or its alias.

CRN is a 120 digit character string of the form -  `crn:v1:bluemix:public:databases-for-postgresql:us-south:a/4ea1882a2d3401ed1e459979941966ea:79226bd4-4076-4873-b5ce-b1dba48ff8c4::`

**Syntax**

```
$ terraform import ibm_database.my_db <crn>
```

**Example**

```
$ terraform import ibm_database.my_db crn:v1:bluemix:public:databases-for-postgresql:us-south:a/4ea1882a2d3401ed1e459979941966ea:79226bd4-4076-4873-b5ce-b1dba48ff8c4::
```

Import requires a minimal Terraform config file to allow importing.

```terraform
resource "ibm_database" "<your_database>" {
  name              = "<your_database_name>"
```

Run `terraform state show ibm_database.<your_database>` after import to retrieve the more values to be included in the resource config file. Observe the ICD exports the admin userid. It does not export any more user IDs and passwords that are configured on the instance. These values must be retrieved from an alternative source. If new passwords need to be configured or the connection string that is retrieved to use the service, a new users block must be defined to create new users. This limitation is due to a lack of ICD functionality.
