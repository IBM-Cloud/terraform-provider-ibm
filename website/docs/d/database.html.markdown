---
subcategory: "Cloud Databases"
layout: "ibm"
page_title: "IBM : Cloud Databases instance"
description: |-
  Get information on an IBM Cloud database instance.
---

# ibm_database

Retrieve information about an existing [IBM Cloud Database instance](https://cloud.ibm.com/docs/cloud-databases).

**Note**
Configuration of an IBM Cloud Databases `data_source` requires that the `region` parameter is set for the IBM provider in the `provider.tf`. The region must be the same as the `location` that the IBM Cloud Databases instance is deployed into. If not specified, `us-south` is used by default. A `terraform refresh` of the `data_source` fails if the region and the location differ.

## Example usage
The following example retrieves information about the `mydatabase` instance in `us-east`.

```terraform
data "ibm_database" "database" {
  name = "mydatabase"
  location = "us-east"
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Required, String) The name of the IBM Cloud Databases instance. IBM Cloud does not enforce that service names are unique and it is possible that duplicate service names exist. The first located service instance is used by  Terraform. The name must not include spaces.
- `location` - (Optional, String) The location where the IBM Cloud Databases instance is deployed into.
- `resource_group_id`- (Optional, String) The ID of the resource group where the IBM Cloud Databases instance is deployed into. The default is `default`.
- `service` - (Optional, String) The service type of the instance. To retrieve this value, run `ibmcloud catalog service-marketplace` or `ibmcloud catalog search`.

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your data source is created. 

- `adminuser` - (String)  The user ID of the default administration user for the database, such as `admin` or `root`.
- `cert_file_path` - (String)  The absolute path to certificate PEM file.
- `connectionstrings`  (List) List of connection strings by userid for the database. For information about how to use connection strings, see the [documentation](https://cloud.ibm.com/docs/databases-for-postgresql?topic=databases-for-postgresql-connection-strings). The results are returned in pairs of the userid and string: `connectionstrings.1.name = admin connectionstrings.1.string = postgres://admin:$PASSWORD@12345aa1-1111-1111-a1aa-a1aaa11aa1a1.a1a1a111a1a11a1a111a111a1a111a111.databases.appdomain.cloud:32554/ibmclouddb?sslmode=verify-full`.
- `configuration_schema` (String) Database Configuration Schema in JSON format.
- `id` - (String) The CRN of the IBM Cloud Databases instance.
- `guid` - (String) The unique identifier of the IBM Cloud Databases instance.
- `plan` - (String)  The service plan of the IBM Cloud Databases instance.
- `location` - (String)  The location where the IBM Cloud Databases instance is deployed into.
- `status` - (String)  The status of the IBM Cloud Databases instance.
- `version` - (String) The database version.
- `platform_options`-  (String) The CRN of key protect key.
   
   Nested scheme for `platform_options`:
   - `key_protect_key_id`-  (String) The CRN of key protect key.
   - `disk_encryption_key_crn`-  (String) The CRN of disk encryption key.
   - `backup_encryption_key_crn`-  (String) The CRN of backup encryption key.
   
- `auto_scaling` (List)Configure rules to allow your database to automatically increase its resources. Single block of autoscaling is allowed at once.

  Nested scheme for `auto_scaling`:
  - `cpu` (List)Autoscaling CPU.
  
     Nested scheme for `cpu`:
     - `rate_increase_percent`- (Integer) Auto scaling rate in increase percent.
     - `rate_limit_count_per_member`- (Integer) Auto scaling rate limit in count per number.
     - `rate_period_seconds`- (Integer) Auto scaling rate in period seconds.
     - `rate_units` - (String) Auto scaling rate in units.
  
  - `disk` (List) Disk auto scaling.
  
    Nested scheme for `disk`:
    - `capacity_enabled`- (Boolean) Auto scaling scalar enables or disables the scalar capacity.
    - `free_space_less_than_percent`- (Integer) Auto scaling scalar capacity free space less than percent.
    - `io_above_percent`- (Integer) Auto scaling scalar I/O utilization above percent.
    - `io_enabled`- (Boolean) Auto scaling scalar I/O utilization enabled.
    - `io_over_period`- (Boolean) Auto scaling scalar I/O utilization over period.
    - `rate_increase_percent`- (Integer) Auto scaling rate increase percent.
    - `rate_limit_mb_per_member`- (Integer) Auto scaling rate limit in megabytes per member.
    - `rate_period_seconds`- (Integer) Auto scaling rate period in seconds.
    - `rate_units` - (String) Auto scaling rate in units.
	
- `memory` (List) Memory Auto Scaling.

  Nested scheme for `memory`:
  - `io_above_percent`- (Integer) Auto scaling scalar I/O utilization above percent.
  - `io_enabled`- (Boolean) Auto scaling scalar I/O utilization enabled.
  - `io_over_period` - (String) Auto scaling scalar I/O utilization over period.
  - `rate_increase_percent`- (Integer) Auto scaling rate in increase percent.
  - `rate_limit_mb_per_member`- (Integer) Auto scaling rate limit in megabytes per member.
  - `rate_period_seconds`- (Integer) Auto scaling rate period in seconds.
  - `rate_units` - (String) Auto scaling rate in units.
- `whitelist` (List) A list of allowed IP addresses or ranges.


**Note**
The provider only exports the admin user ID and associated connection string. It does not export any user IDs that are configured for the instance in addition. 
