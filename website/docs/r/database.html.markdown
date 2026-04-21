---
subcategory: "Cloud Databases"
layout: "ibm"
page_title: "IBM : Cloud Database instance"
description: |-
  Manages IBM Cloud database instance.
---

# ibm_database

Create, update, or delete an IBM Cloud Database (ICD) instance. The `ibmcloud_api_key` that is used by Terraform should grant IAM rights to create and modify IBM Cloud Databases and have access to the resource group the ICD instance is associated with. For more information, see [documentation](https://cloud.ibm.com/docs/services/databases-for-postgresql/reference-access-management.html#identity-and-access-management) to manage ICD instances.

If `resource_group_id` is not specified, the ICD instance is created in the default resource group. The `API_KEY` must be assigned permissions for this group.

Configuration of an ICD resource requires that the `region` parameter is set for the IBM provider in the `provider.tf` to be the same as the target ICD `location/region`. If not specified it defaults to `us-south`. A `terraform apply` fails if the ICD `location` is set differently. If the Terraform configuration needs to deploy resources into multiple regions, provider alias can be used. For more information, see [Terraform provider configuration](https://www.terraform.io/docs/configuration/providers.html#multiple-provider-instances).

## Plan Selection: Classic vs Gen2

IBM Cloud Databases offers two infrastructure generations:

- **Classic plans**: Standard plans like `standard`, `enterprise`, and `platinum`
- **Gen2 plans**: Next-generation plans with `-gen2` suffix (e.g., `standard-gen2`, `enterprise-gen2`)

The plan you select determines which infrastructure your database uses. Both generations support the core database functionality, but have some differences in available features and management approaches. See the [Argument Reference](#argument-reference) section below for details on feature availability by plan type.

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

  adminpassword                = "password12345678"

  group {
    group_id = "member"

    memory {
      allocation_mb = 14336
    }

    disk {
      allocation_mb = 20480
    }

    cpu {
      allocation_count = 3
    }
  }

  users {
    name     = "user123"
    password = "password12345678"
    type     = "database"
  }

  allowlist {
    address     = "172.168.1.1/32"
    description = "desc"
  }
}

output "ICD Etcd database connection string" {
  value  = "http://${ibm_database.test_acc.ibm_database_connection.icd_conn}"
}

```

### Gen2 database instance example

An example to configure and deploy a Gen2 database instance. Gen2 plans use the `-gen2` suffix and have specific scaling requirements.

```terraform
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_database" "<your_database>" {
  name              = "<your_database_name>"
  plan              = "standard-gen2"
  location          = "us-south"
  service           = "databases-for-postgresql"
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["tag1", "tag2"]

  group {
    group_id = "member"

    members {
      allocation_count = 3
    }

    disk {
      allocation_mb = 20480
    }

    host_flavor {
      id = "bx3d.4x20"
    }
  }
}

output "ICD Gen2 database connection" {
  description = "Use ibm_resource_key to obtain connection credentials"
  value       = "Connection details available via ibm_resource_key resource"
}
```

**Note:** Gen2 instances use `ibm_resource_key` for credential management instead of the `adminpassword` and `users` attributes used in Classic plans.

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

  adminpassword                = "password12345678"

  group {
    group_id = "member"

    memory {
      allocation_mb = 10240
    }

    disk {
      allocation_mb = 256000
    }

    cpu {
      allocation_count = 3
    }
  }

  users {
    name     = "user123"
    password = "password12345678"
  }

  allowlist {
    address     = "172.168.1.1/32"
    description = "desc"
  }
}

output "ICD Etcd database connection string" {
  value = "http://${ibm_database.test_acc.ibm_database_connection.icd_conn}"
}

```

### Sample database instance by using `host_flavor` attribute
An example to configure and deploy database by using `host_flavor` attribute.

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

    host_flavor {
      id = "b3c.8x32.encrypted"
    }

    disk {
      allocation_mb = 256000
    }
  }

  users {
    name     = "user123"
    password = "password12"
  }

  allowlist {
    address     = "172.168.1.1/32"
    description = "desc"
  }
}

output "ICD Etcd database connection string" {
  value = "http://${ibm_database.test_acc.ibm_database_connection.icd_conn}"
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
  adminpassword                = "password12345678"

  group {
    group_id = "member"

    memory {
      allocation_mb = 24576
    }

    disk {
      allocation_mb = 122880
    }

    cpu {
      allocation_count = 6
    }
  }

  tags                         = ["one:two"]

  users {
    name      = "dbuser"
    password  = "password12345678"
    type      = "database"
  }

  users {
    name     = "opsmanageruser"
    password = "$ecurepa$$word12"
    type     = "ops_manager"
    role     = "group_read_only"
  }

  allowlist {
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
* To enable Analytics and/or BI Connector for MongoDB Enterprise, a `group` attribute must be defined for the `analytics` and `bi_connector` group types with `members` scaled to at exactly `1`. Read more about Analytics and BI Connector [here](https://cloud.ibm.com/docs/databases-for-mongodb?topic=databases-for-mongodb-mongodbee-analytics)

    > 🛑 **Deprectaed:** Analytics Add-On for Databases for MongoDB Enterprise Edition is deprecated after March 31,2025.

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
  adminpassword     = "password12345678"
  tags              = ["one:two"]

  group {
    group_id = "member"

    memory {
      allocation_mb = 24576
    }

    disk {
      allocation_mb = 122880
    }

    cpu {
      allocation_count = 6
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
  adminpassword                = "password12345678"

  group {
    group_id = "member"

    memory {
      allocation_mb = 12288
    }

    disk {
      allocation_mb = 131072
    }

    cpu {
      allocation_count = 3
    }
  }

  tags                         = ["one:two"]

  users {
    name      = "user123"
    password  = "password12345678"
    type      = "database"
  }

  allowlist {
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

### Sample Elasticsearch Enterprise instance

```terraform
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "es" {
  resource_group_id            = data.ibm_resource_group.test_acc.id
  name                         = "es-enterprise"
  service                      = "databases-for-elasticsearch"
  plan                         = "enterprise"
  location                     = "eu-gb"
  adminpassword                = "password12345678"
  version                      = "7.17"
  group {
    group_id = "member"
    members {
      allocation_count = 3
    }
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
    password = "password12345678"
  }
  allowlist {
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
### Sample Elasticsearch Platinum instance

```terraform
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "es" {
  resource_group_id            = data.ibm_resource_group.test_acc.id
  name                         = "es-platinum"
  service                      = "databases-for-elasticsearch"
  plan                         = "platinum"
  location                     = "eu-gb"
  adminpassword                = "password12345678"
  group {
    group_id = "member"
    members {
      allocation_count = 3
    }
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
    password = "password12345678"
  }
  allowlist {
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
  group {
    group_id = "member"

    memory {
      allocation_mb = 12288
    }

    disk {
      allocation_mb = 131072
    }

    cpu {
      allocation_count = 3
    }
  }
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

### Creating logical replication slot for postgres database

```terraform
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "db" {
  name                         = "example-database"
  service                      = "databases-for-postgresql"
  plan                         = "standard"
  location                     = "us-east"

  users {
    name     = "repl"
    password = "repl12345password"
  }

  configuration                = <<CONFIGURATION
  {
    "wal_level": "logical",
    "max_replication_slots": 21,
    "max_wal_senders": 21
  }
  CONFIGURATION

  logical_replication_slot {
    name = "wj123"
    database_name = "ibmclouddb"
    plugin_type = "wal2json"
  }
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

ICD create instance typically takes between 30 minutes to 45 minutes. Delete and update takes a minute with the exception of an in place version upgrade. Provisioning time are unpredictable, if the apply fails due to a timeout, import the database resource once the create is completed.


## Argument reference
Review the argument reference that you can specify for your resource.

- `adminpassword` - (Optional, String) The password for the database administrator. Password must be between 15 and 32 characters in length and contain a letter and a number. The only special characters allowed are `-_`.

  **Gen2:** Accepted but ignored. Gen2 instances do not have a default admin user. Use the `ibm_resource_key` resource to create service credentials for database access.

- `auto_scaling` (List, Optional) Configure rules to allow your database to automatically increase its resources. Single block of autoscaling is allowed at once.

  **Gen2:** Accepted but ignored. Auto-scaling policies are not available in Gen2. Monitor your database and manually adjust scaling as needed.

   - Nested scheme for `auto_scaling`:
     - `disk` (List , Optional) Single block of disk is allowed at once in disk auto scaling.
        - Nested scheme for `disk`:
          - `capacity_enabled` - (Optional, Bool) Auto scaling scalar enables or disables the scalar capacity.
          - `free_space_less_than_percent` - (Optional, Integer) Auto scaling scalar capacity free space less than percent.
          - `io_above_percent` - (Optional, Integer) Auto scaling scalar I/O utilization above percent.
          - `io_enabled` - (Optional, Bool) Auto scaling scalar I/O utilization enabled.`
          - `io_over_period` - (Optional, String) Auto scaling scalar I/O utilization over period.
          - `rate_increase_percent` - (Optional, Integer) Auto scaling rate increase percent.
          - `rate_limit_mb_per_member` - (Optional, Integer) Auto scaling rate limit in megabytes per member.
          - `rate_period_seconds` - (Optional, Integer) Auto scaling rate period in seconds.
          - `rate_units` - (Optional, String) Auto scaling rate in units.

     - `memory` (List , Optional) Memory Auto Scaling in single block of memory is allowed at once.
       - Nested scheme for `memory`:
         - `io_above_percent` - (Optional, Integer) Auto scaling scalar I/O utilization above percent.
         - `io_enabled`-Bool-Optional-Auto scaling scalar I/O utilization enabled.
         - `io_over_period` - (Optional, String) Auto scaling scalar I/O utilization over period.
         - `rate_increase_percent` - (Optional, Integer) Auto scaling rate in increase percent.
         - `rate_limit_mb_per_member` - (Optional, Integer) Auto scaling rate limit in megabytes per member.
         - `rate_period_seconds` - (Optional, Integer) Auto scaling rate period in seconds.
         - `rate_units` - (Optional, String) Auto scaling rate in units.

- `backup_id` - (Optional, String) The CRN of a backup resource to restore from. The backup is created by a database deployment with the same service ID. The backup is loaded after provisioning and the new deployment starts up that uses that data. A backup CRN is in the format `crn:v1:<…>:backup:`. If omitted, the database is provisioned empty.
- `backup_encryption_key_crn`- (Optional, Forces new resource, String) The CRN of a key protect key, that you want to use for encrypting disk that holds deployment backups. A key protect CRN is in the format `crn:v1:<...>:key:`. Backup_encryption_key_crn can be added only at the time of creation and no update support  are available.
- `configuration` - (Optional, Json String) Database Configuration in JSON format. Supported services: `databases-for-postgresql`, `databases-for-redis`, `databases-for-mysql`,`messages-for-rabbitmq` and `databases-for-enterprisedb`. For valid values please refer [API docs](https://cloud.ibm.com/apidocs/cloud-databases-api/cloud-databases-api-v5#updatedatabaseconfiguration).

  **Gen2:** Accepted but ignored. Database configuration management is not yet implemented for Gen2 instances.

- `logical_replication_slot` - (Optional, List of Objects) A list of logical replication slots that you want to create on the database. Multiple blocks are allowed. This is only available for `databases-for-postgresql`.

  **Gen2:** Accepted but ignored. Logical replication slots are not available for Gen2 instances.

  Nested scheme for `logical_replication_slot`:
  - `name` - (Required, String) The name of the `logical_replication_slot`.
  - `database_name` - (Required, String) The name of the database on which you want to create the `logical_replication_slot`.
  - `plugin_type` - (Required, String) The plugin type that is used to create the `logical_replication_slot`. Only `wal2json` is supported.

  Prereqs to creating a logical replication slot:
  - Make sure the replication user's (`repl`) password has been changed.
  - Make sure that your database is configured such that logical replication can be enabled. This means thats the `wal_level` needs to be set to `logical`. Also, `max_replication_slots` and `max_wal_senders` must be greater than 20.
  - For more information on enabling logical replication slots please see [Configuring Wal2json](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-wal2json)
- `guid` - (Optional, String) The unique identifier of the database instance.
- `key_protect_key` - (Optional, Forces new resource, String) The root key CRN of a Key Management Services like Key Protect or Hyper Protect Crypto Service (HPCS) that you want to use for disk encryption. A key CRN is in the format `crn:v1:<…>:key:`. You can specify the root key during the database creation only. After the database is created, you cannot update the root key. For more information, refer [Disk encryption](https://cloud.ibm.com/docs/cloud-databases?topic=cloud-databases-key-protect#using-the-key-protect-key) documentation.

- `key_protect_instance` - (Optional, Forces new resource, String) The instance CRN of a Key Management Services like Key Protect or Hyper Protect Crypto Service (HPCS) that you want to use for disk encryption. An instance CRN is in the format `crn:v1:<…>::`.

  **Gen2:** Accepted but not used. Use `key_protect_key` for disk encryption and `backup_encryption_key_crn` for backup encryption.
- `location` - (Required, String) The location where you want to deploy your instance. The location must match the `region` parameter that you specify in the `provider` block of your  Terraform configuration file. The default value is `us-south`. Currently, supported regions are `us-south`, `us-east`, `eu-gb`, `eu-de`, `au-syd`, `jp-tok`, `oslo01`.
- `group` - (Optional, Set) A set of group scaling values for the database. Multiple blocks are allowed. Can only be performed on is_adjustable=true groups. Values set are per-member. Values must be greater than or equal to the minimum size and must be a multiple of the step size.

  **Gen2:** Fully supported for `members`, `disk`, and `host_flavor`. Note: Multitenant hosts are not supported in Gen2. Plan fails if `memory` or `cpu` allocations are set, as memory and CPU are determined by the dedicated `host_flavor` and cannot be set independently.

  - Nested scheme for `group`:
    - `group_id` - (Optional, String) The ID of the scaling group. Scaling group ID allowed values:  `member`, `analytics`, or `bi_connector`. Read more about `analytics` and `bi_connector` [here](https://cloud.ibm.com/docs/databases-for-mongodb?topic=databases-for-mongodb-mongodbee-analytics).


    - `members` (Set, Optional)
      - Nested scheme for `members`:
        - `allocation_count` - (Optional, Integer) Allocated number of members. **Gen2: Supported.**

    - `memory` (Set, Optional) Memory Auto Scaling in single block of memory is allowed at once.
      - Nested scheme for `memory`:
        - `allocation_mb` - (Optional, Integer) Allocated memory per-member. **Gen2: Plan fails if set** - memory is determined by `host_flavor`.

    - `disk` (Set, Optional)
      - Nested scheme for `disk`:
        - `allocation_mb` - (Optional, Integer) Allocated disk per-member. **Gen2: Supported** (specified as storage_gb in the API).

    - `cpu` (Set, Optional)
      - Nested scheme for `cpu`:
        - `allocation_count` - (Optional, Integer) Allocated dedicated CPU per-member. **Gen2: Plan fails if set** - CPU is determined by `host_flavor`.

    - `host_flavor` (Set, Optional)
      - Nested scheme for `host_flavor`:
        - `id` - (Optional, String) The hosting infrastructure identifier. **Gen2: Supported.**

          **Classic:** Selecting `multitenant` places your database on a logically separated, multi-tenant machine. With this identifier, minimum resource configurations apply. Alternatively, setting the identifier to any of the following host sizes places your database on the specified host size with no other tenants:
          - `b3c.4x16.encrypted`
          - `b3c.8x32.encrypted`
          - `m3c.8x64.encrypted`
          - `b3c.16x64.encrypted`
          - `b3c.32x128.encrypted`
          - `m3c.30x240.encrypted`

          **Gen2:** Use dedicated host flavors like `bx3d.4x20`, `bx3d.8x40`, `mx3d.8x64`, etc. The `multitenant` option is not supported for Gen2 plans.

- `name` - (Required, String) A descriptive name that is used to identify the database instance. The name must not include spaces.
- `offline_restore` - (Optional, Boolean) Enable or disable the Offline Restore option while performing a Point-in-time Recovery for MongoDB EE in a disaster recovery scenario when the source region is unavailable, see [Point-in-time Recovery](https://cloud.ibm.com/docs/databases-for-mongodb?topic=databases-for-mongodb-pitr&interface=api#pitr-offline-restore)
- `plan` - (Required, Forces new resource, String) The name of the service plan that you choose for your instance. The plan determines whether your instance uses Classic or Gen2 infrastructure:
  - **Classic plans**: `standard`, `enterprise`, `platinum`
  - **Gen2 plans**: `standard-gen2`, `enterprise-gen2`, `platinum-gen2`

  Plans ending with `-gen2` use Gen2 infrastructure. `enterprise` is supported only for elasticsearch (`databases-for-elasticsearch`) and mongodb (`databases-for-mongodb`). `platinum` is supported for elasticsearch (`databases-for-elasticsearch`).
- `point_in_time_recovery_deployment_id` - (Optional, String) The ID of the source deployment that you want to recover back to.
- `point_in_time_recovery_time` - (Optional, String) The timestamp in UTC format that you want to restore to. To retrieve the timestamp, run the `ibmcloud cdb postgresql earliest-pitr-timestamp <deployment name or CRN>` command. To restore to the latest available time, use a blank string `""` as the timestamp. For more information, see [Point-in-time Recovery](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-pitr).
- `remote_leader_id` - (Optional, String) A CRN of the leader database to make the replica(read-only) deployment. The leader database is created by a database deployment with the same service ID. A read-only replica is set up to replicate all of your data from the leader deployment to the replica deployment by using asynchronous replication. Removing the `remote_leader_id` attribute from an existing read-only replica will promote the deployment to a standalone deployment. The deployment will restart and break its connection with the leader. This will disable all database users associated with this deployment. For more information, see [Configuring Read-only Replicas](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-read-only-replicas).

  **Gen2:** Plan fails if set. Read-only replica creation and promotion are not supported for Gen2 instances.

- `skip_initial_backup` - (Optional, Boolean) Should only be set when promoting a read-only replica. By setting this value to `true`, you skip the initial backup that would normally be taken upon promotion. Skipping the initial backup means that your replica becomes available more quickly, but there is no immediate backup available. The default is `false`. For more information, see [Configuring Read-only Replicas]

  **Gen2:** Accepted but ignored (Classic-only feature for read replica promotion).
- `async_restore` - (Optional, Boolean) Should only be set for asynchronous restore. By setting this value to `true`, the restore is initiated as an asynchronous operation, which helps to reduce end-to-end restore time. Only applicable when restoring a PostgreSQL instance.
- `resource_group_id` - (Optional, Forces new resource, String)  The ID of the resource group where you want to create the instance. To retrieve this value, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If no value is provided, the `default` resource group is used.
- `service` - (Required, Forces new resource, String) The type of Cloud Databases that you want to create. Only the following services are currently accepted: `databases-for-etcd`, `databases-for-postgresql`, `databases-for-redis`, `databases-for-elasticsearch`, `messages-for-rabbitmq`,`databases-for-mongodb`,`databases-for-mysql`, and `databases-for-enterprisedb`.
- `service_endpoints` - (Optional, String) Specify whether you want to enable the public, private, or both service endpoints. Supported values are `public`, `private`, or `public-and-private`.

  **Gen2:** Optional; must be `private` if set. Gen2 instances only support private endpoints and default to `private`. Plan fails if set to `public` or `public-and-private`.
- `tags` (Optional, Array of Strings) A list of tags that you want to add to your instance.
- `version` - (Optional, String) The version of the database to be provisioned or upgraded to. If omitted, the database is created with the latest supported major and minor version.

  **Classic:** This field can be updated to perform an in-place upgrade without forcing the creation of a new resource. The database will be put into READ-ONLY mode during upgrade.

  **Gen2:** Can be set at creation only. Updates fail with error. In-place version upgrades are not supported for Gen2 plans. To upgrade a Gen2 instance, create a new instance with the desired version.

    > ⚠️ **Warning:** For Classic plans, upgrading may require more time than the default timeout.
    > A longer timeout value can be set using the timeouts attribute.

- `version_upgrade_skip_backup` - (Optional, Boolean) Whether to skip taking a backup before upgrading the database version. This is only applicable to databases that do not support point-in-time restore (PITR). To learn more, refer to the version upgrade documentation.

  **Gen2:** Accepted but ignored (Classic-only feature for version upgrades).

    > ⚠️ **Warning:** For Classic plans, skipping a backup is **not recommended**.
    > Skipping a backup before a version upgrade is dangerous and may result in **data loss** if the upgrade fails at any stage — there will be **no immediate backup** to restore from.

- `deletion_protection` - (Optional, Boolean) If the DB instance should have deletion protection within terraform enabled. This is not a property of the resource and does not prevent deletion outside of terraform. The database can't be deleted by terraform when this value is set to `true`. The default is `false`.
- `users` - (Optional, List of Objects) A list of users that you want to create on the database. Multiple blocks are allowed.

  **Gen2:** Plan fails if set. Use the `ibm_resource_key` resource to create service credentials for Gen2 instances.

  Nested scheme for `users`:
  - `name` - (Required, String) The user name to add to the database instance. The user name must be in the range 5 - 32 characters.
  - `password` - (Required, String) The password for the user. Passwords must be between 15 and 32 characters in length and contain a letter and a number. Users with an `ops_manager` user type must have a password containing a special character `~!@#$%^&*()=+[]{}|;:,.<>/?_-` as well as a letter and a number. Other user types may only use special characters `-_`.
  - `type` - (Optional, String) The type for the user. Examples: `database`, `ops_manager`, `read_only_replica`. The default value is `database`.
  - `role` - (Optional, String) The role for the user. Only available for `ops_manager` user type or Redis 6.0 and above. Example roles for `ops_manager`: `group_read_only`, `group_data_access_admin`. For, Redis 6.0 and above, `role` must be in Redis ACL syntax for adding and removing command categories i.e. `+@category` or  `-@category`. Allowed command categories are `all`, `admin`, `read`, `write`. Example Redis `role`: `-@all +@read`

- `allowlist` - (Optional, List of Objects) A list of allowed IP addresses for the database. Multiple blocks are allowed.

  **Gen2:** Plan fails if set. IP allowlist configuration is not available for Gen2 instances.

  Nested scheme for `allowlist`:
  - `address` - (Optional, String) The IP address or range of database client addresses to be allowlisted in CIDR format. Example, `172.168.1.2/32`.
  - `description` - (Optional, String) A description for the allowed IP addresses range.

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your resource is created.

- `adminuser` - (String) The user ID of the database administrator. Example, `admin` or `root`.

  **Gen2:** Always empty. Gen2 instances do not have a default admin user. Use `ibm_resource_key` for credentials.

- `configuration_schema` (String) Database Configuration Schema in JSON format.

  **Gen2:** Always empty (not available).

- `id` - (String) The CRN of the database instance.
- `status` - (String) The status of the instance.
- `version` - (String) The database version.
- `groups` - (List) A list of scaling groups for the database. This is a computed attribute that shows the current resource allocations.

## Gen2 Feature Summary

The following table summarizes feature availability for Classic and Gen2 plans:

| Feature | Classic Plans | Gen2 Plans |
|---------|--------------|------------|
| Basic provisioning (name, location, service, plan) | ✅ Supported | ✅ Supported |
| Resource group assignment | ✅ Supported | ✅ Supported |
| Tags | ✅ Supported | ✅ Supported |
| Encryption (key_protect_key, backup_encryption_key_crn) | ✅ Supported | ✅ Supported |
| Restore from backup (backup_id) | ✅ Supported | ❌ Plan fails if set |
| Point-in-time recovery (point_in_time_recovery_deployment_id, point_in_time_recovery_time) | ✅ Supported | ❌ Plan fails if set |
| Offline restore (MongoDB) | ✅ Supported | ❌ Accepted but ignored |
| Async restore (PostgreSQL) | ✅ Supported | ❌ Accepted but ignored |
| Scaling (members, disk, host_flavor) | ✅ Supported | ✅ Supported |
| Scaling (memory, cpu) | ✅ Supported | ❌ Plan fails if set (controlled by host_flavor) |
| Service endpoints | ✅ public, private, public-and-private | ⚠️ private only (plan fails if public) |
| Admin password | ✅ Supported | ❌ Accepted but ignored (use ibm_resource_key) |
| User management | ✅ Supported | ❌ Plan fails if set (use ibm_resource_key) |
| IP allowlist | ✅ Supported | ❌ Plan fails if set (use ibm_resource_key) |
| Database configuration | ✅ Supported | ❌ Accepted but ignored |
| Auto-scaling | ✅ Supported | ❌ Accepted but ignored |
| Logical replication slots | ✅ Supported | ❌ Accepted but ignored |
| Read-only replicas | ✅ Supported | ❌ Plan fails if set |
| In-place version upgrades | ✅ Supported | ❌ Updates fail with error |
| Deletion protection | ✅ Supported | ✅ Supported |

### Gen2 Validation Behavior

Gen2 plans handle unsupported features in two ways:

- **Plan fails if set**: Terraform plan will fail with a validation error if these attributes are configured. You must remove them from your configuration to use Gen2 plans.
  - Examples: `backup_id`, `point_in_time_recovery_deployment_id`, `point_in_time_recovery_time`, `users`, `allowlist`, `remote_leader_id`, memory/cpu in `group`

- **Accepted but ignored**: These attributes can remain in your configuration for easier migration, but they have no effect on Gen2 instances. They are silently ignored during apply and cleared during read operations.
  - Examples: `auto_scaling`, `configuration`, `logical_replication_slot`, `offline_restore`, `async_restore`, `adminpassword`

**Note:** For Gen2 instances, use the `ibm_resource_key` resource to create service credentials and obtain connection information.

## Import
The database instance can be imported by using the ID, that is formed from the CRN. To import the resource, you must specify the `region` parameter in the `provider` block of your Terraform configuration file. If the region is not specified, `us-south` is used by default. A Terraform refresh or apply fails if the database instance is not in the same region as configured in the provider or its alias.

**Both Classic and Gen2 instances can be imported using the same method.** The plan type (Classic vs Gen2) is automatically detected from the instance's service plan.

CRN is a 120 digit character string of the form - `crn:v1:bluemix:public:databases-for-postgresql:us-south:a/4ea1882a2d3401ed1e459979941966ea:79226bd4-4076-4873-b5ce-b1dba48ff8c4::`

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
}
```

Run `terraform state show ibm_database.<your_database>` after import to retrieve the more values to be included in the resource config file.

**For Classic instances:** The import exports the admin userid but does not export any other user IDs and passwords that are configured on the instance. These values must be retrieved from an alternative source. If new passwords need to be configured or the connection string is needed to use the service, a new users block must be defined to create new users.

**For Gen2 instances:** User management is not available via Terraform. Use the `ibm_resource_key` resource to create service credentials and obtain connection information for Gen2 instances.
