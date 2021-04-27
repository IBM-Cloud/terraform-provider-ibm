---
subcategory: "Cloud Databases"
layout: "ibm"
page_title: "IBM : Cloud Databases instance"
description: |-
  Get information on an IBM Cloud Database Instance.
---

# ibm\_database

Creates a read only copy of an existing IBM Cloud Databases resource.  

Configuration of an ICD data_source requires that the `region` parameter is set for the IBM provider in the `provider.tf` to be the same as the ICD service `location/region` that the service will be deployed in. If not specified it will default to `us-south`. A `terraform refresh` of the data_source will fail if the ICD `location` is  different to that specified on the provider.  

## Example Usage

```hcl
data "ibm_database" "<your_database>" {
  name = "<your_database>"
  location = "<your-db-location>"
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required, string) The name used to identify the IBM Cloud Database instance in the IBM Cloud UI. IBM Cloud does not enforce that service names are unique and it is possible that duplicate service names exist. The first located service instance is used by Terraform. The name must not include spaces.

* `resource_group_id` - (Optional, string) The id of the resource group where the resource instance exists. If not provided it takes the default resource group.

* `location` - (Optional, string) The location or the environment in which instance exists.

* `service` - (Optional, string) The service type of the instance. You can retrieve the value by running the `ibmcloud catalog service-marketplace` or `ibmcloud catalog search` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of this ICD instance.
* `plan` - The service plan for this ICD instance
* `location` - The location or region for this ICD instance. This will be the same as defined for the provider or its alias.
* `status` - Status of the ICD instance.
* `id` - The unique identifier of the new database instance (CRN).
* `status` - Status of resource instance.
* `adminuser` - userid of the default administration user for the database, usually `admin` or `root`.
* `version` - Database version.
* `platform_options` - Platform-specific options for this deployment.
  * `key_protect_key_id` - The CRN of Key protect key.
  * `disk_encryption_key_crn` - The CRN of Disk encryption Key.
  * `backup_encryption_key_crn` - The CRN of Backup encryption Key.
* `cert_file_path` - The absolute path to certificate PEM file.
* `connectionstrings` - List of connection strings by userid for the database. See the IBM Cloud documentation for more details of how to use connection strings in ICD for database access: https://cloud.ibm.com/docs/services/databases-for-postgresql/howto-getting-connection-strings.html#getting-your-connection-strings. The results are returned in pairs of the userid and string:
  `connectionstrings.1.name = admin`
  `connectionstrings.1.string = postgres://admin:$PASSWORD@79226bd4-4076-4873-b5ce-b1dba48ff8c4.b8a5e798d2d04f2e860e54e5d042c915.databases.appdomain.cloud:32554/ibmclouddb?sslmode=verify-full`
* `whitelist` - List of whitelisted IP addresses or ranges.
* `guid` - Unique identifier of resource instance.
* `auto_scaling` -  Configure rules to allow your database to automatically increase its resources.
  * `cpu` - CPU AutoScaling
    * `rate_increase_percent` - Auto Scaling Rate: Increase Percent
    * `rate_limit_count_per_member` - Auto Scaling Rate: Limit count per number
    * `rate_period_seconds` - Auto Scaling Rate: Period Seconds
    * `rate_units` - Auto Scaling Rate: Units
  * `disk` - Disk AutoScaling
    * `capacity_enabled` - Auto Scaling Scalar: Enables or disable the capacity scalar
    * `free_space_less_than_percent` - Auto Scaling Scalar: Capacity Free Space Less Than Percent
    * `io_above_percent` - Auto Scaling Scalar: IO Utilization Above Percent
    * `io_enabled` - Auto Scaling Scalar: IO Utilization Enabled
    * `io_over_period` - Auto Scaling Scalar: IO Utilization Over Period
    * `rate_increase_percent` - Auto Scaling Rate: Increase Percent
    * `rate_limit_mb_per_member` - Auto Scaling Rate: Limit mb per member
    * `rate_period_seconds` - Auto Scaling Rate: Period Seconds
    * `rate_units` - Auto Scaling Rate: Units
  * `memory` - Memory AutoScaling
    * `io_above_percent` - Auto Scaling Scalar: IO Utilization Above Percent
    * `io_enabled` - Auto Scaling Scalar: IO Utilization Enabled
    * `io_over_period` - Auto Scaling Scalar: IO Utilization Over Period
    * `rate_increase_percent` - Auto Scaling Rate: Increase Percent
    * `rate_limit_mb_per_member` - Auto Scaling Rate: Limit mb per member
    * `rate_period_seconds` - Auto Scaling Rate: Period Seconds
    * `rate_units` - Auto Scaling Rate: Units

Note that the provider only exports the admin userid and associated connectionstring. It does not export any userids additionally configured for the instance. This is due to a lack of ICD function.

