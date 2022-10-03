---
layout: "ibm"
page_title: "IBM : ibm_cm_offering_instance"
description: |-
  Get information about cm_offering_instance
subcategory: "Catalog Management API"
---

# ibm_cm_offering_instance

Provides a read-only data source for cm_offering_instance. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cm_offering_instance" "cm_offering_instance" {
	instance_identifier = "instance_identifier"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_identifier` - (Required, Forces new resource, String) Version Instance identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cm_offering_instance.
* `account` - (String) The account this instance is owned by.

* `catalog_id` - (String) Catalog ID this instance was created from.

* `channel` - (String) Channel to pin the operator subscription to.

* `cluster_all_namespaces` - (Boolean) designate to install into all namespaces.

* `cluster_id` - (String) Cluster ID.

* `cluster_namespaces` - (List) List of target namespaces to install into.

* `cluster_region` - (String) Cluster region (e.g., us-south).

* `created` - (String) date and time create.

* `crn` - (String) platform CRN for this instance.

* `disabled` - (Boolean) Indicates if Resource Controller has disabled this instance.

* `id` - (String) provisioned instance ID (part of the CRN).

* `install_plan` - (String) Type of install plan (also known as approval strategy) for operator subscriptions. Can be either automatic, which automatically upgrades operators to the latest in a channel, or manual, which requires approval on the cluster.

* `kind_format` - (String) the format this instance has (helm, operator, ova...).

* `kind_target` - (String) The target kind for the installed software version.

* `label` - (String) the label for this instance.

* `last_operation` - (List) the last operation performed and status.
Nested scheme for **last_operation**:
	* `code` - (String) Error code from the last operation, if applicable.
	* `message` - (String) additional information about the last operation.
	* `operation` - (String) last operation performed.
	* `state` - (String) state after the last operation performed.
	* `transaction_id` - (String) transaction id from the last operation.
	* `updated` - (String) Date and time last updated.

* `location` - (String) String location of OfferingInstance deployment.

* `metadata` - (Map) Map of metadata values for this offering instance.

* `offering_id` - (String) Offering ID this instance was created from.

* `resource_group_id` - (String) Id of the resource group to provision the offering instance into.

* `rev` - (String) Cloudant revision.

* `schematics_workspace_id` - (String) Id of the schematics workspace, for offering instances provisioned through schematics.

* `sha` - (String) The digest value of the installed software version.

* `updated` - (String) date and time updated.

* `url` - (String) url reference to this object.

* `version` - (String) The version this instance was installed from (semver - not version id).

* `version_id` - (String) The version id this instance was installed from (version id - not semver).

