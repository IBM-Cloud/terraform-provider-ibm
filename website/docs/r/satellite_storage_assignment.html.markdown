---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_storage_assignment"
description: |-
  Manages IBM Cloud Satellite Storage Assignment.
---

# ibm_satellite_storage_assignment

Create, update, or delete [IBM Cloud Storage Assignment](https://cloud.ibm.com/docs/satellite?topic=satellite-storage-template-ov&interface=ui). With Storage Assignments, you can assign your storage configuration to clusters, service clusters, and cluster groups in your location.

## Example usage

###  Sample to create a storage assignment to a cluster

```terraform
resource "ibm_satellite_storage_assignment" "assignment" {
    assignment_name = "assignment-name"
    cluster = "cluster-id"
    config = "storage-config-name"
    controller = "satellite-location"
}
```

###  Sample to create a storage assignment to a group
```terraform
resource "ibm_satellite_storage_assignment" "assignment" {
    assignment_name = "assignment-name"
    config = "storage-config-name"
    groups = ["cluster-group-1","cluster-group-2"]
    controller = "satellite-location"
}
```

###  Sample to update the storage configuration revision to a cluster/group
```terraform
resource "ibm_satellite_storage_assignment" "assignment" {
    assignment_name = "assignment-name"
    config = "storage-config-name"
    groups = ["cluster-group-1","cluster-group-2"]
    update_config_revision = true
}
```
## Timeouts
The `ibm_satellite_storage_assignment` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 20 minutes) Used for creating Instance.
- **update** - (Default 20 minutes) Used for creating Instance.
- **delete** - (Default 20 minutes) Used for deleting Instance.

## Argument reference
Review the argument references that you can specify for your resource. 

- `assignment_name` - (Required, String) The name of the assignment
- `controller` - (Required, String) The name of the location where the storage configuration is created.
- `config` - (Required, String) The name of the storage configuration to be assigned.
- `groups` - (Required, List(String)) A list of strings of cluster groups you want to assign the defined configuration too.
  * Constraints: Required with `controller`
- `cluster` - (Required, String) The id of the cluster you wish to assign the defined configuration too.
  * Constraints: Required with `controller`
- `update_config_revision` - (Optional, Bool) Set to true to update the assignment with the latest revision of the storage configuration.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `uuid` - (String) The Universally Unique IDentifier (UUID) of the Assignment.
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