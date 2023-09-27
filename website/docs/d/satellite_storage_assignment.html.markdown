---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_storage_assignment"
description: |-
  Get information about an IBM Cloud Satellite Storage Assignment.
---

# ibm_satellite_storage_assignment
Retrieve information of an existing Satellite Storage Assignment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax. For more information, about IBM Cloud Satellite Storage Configurations see [Satellite Storage](https://cloud.ibm.com/docs/satellite?topic=satellite-storage-template-ov&interface=ui).


## Example usage

```terraform
data "ibm_satellite_storage_assignment" "assignment" {
  uuid  = var.uuid
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `uuid` - (String) The Universally Unique IDentifier (UUID) of the Assignment.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `assignment_name` - (String) The name of the assignment
- `config` - (String) The name of the storage configuration assigned.
- `groups` - (List(String)) A list of strings of cluster groups assigned to the defined configuration.
- `cluster` - (Required, String) The id of the cluster assigned to the defined configuration.
  * Constraints: Required with `controller` and Conflicts with `groups`
- `owner` - (String) The Owner of the Assignment.
- `svc_cluster` - (String) ID of the Service Cluster that you have assigned the configuration to.
- `sat_cluster` - (String) ID of the Satellite cluster that you have assigned the configuration to.
- `config_uuid` - (String) The Universally Unique IDentifier (UUID) of the Storage Configuration.
- `config_version` - (String) The Current Storage Configuration Version.
- `config_version_uuid` - (String) The Universally Unique IDentifier (UUID) of the Storage Configuration Version.
- `assignment_type` - (String) The Type of Assignment.
- `created` - (String) The Time of Creation of the Assignment.
- `rollout_success_count` - (String) The Rollout Success Count of the Assignment.
- `rollout_error_count` - (String) The Rollout Error Count of the Assignment.
- `is_assignment_upgrade_available` - (Bool) Whether a Configuration Revision Update is Available for the Assignment.
- `id` - (String) ID of the Storage Assignment Resource
