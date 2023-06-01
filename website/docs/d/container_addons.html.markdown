---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_addons"
description: |-
  Reads all enabled IBM  container add-ons.
---

# ibm_container_addons
Retrieve information about all the add-ons that are enables on a cluster. For more information, see [Cluster addons](https://cloud.ibm.com/docs/containers?topic=containers-api-at-iam#ks-cluster).

## Example usage
The following example retrieves information of an add-ons.

```terraform
data "ibm_container_addons" "addons" {
  cluster= ibm_container_addons.addons.cluster
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cluster` - (Required, String) The name or ID of the cluster.

## Attribute reference
In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `addons` - (String) The details of an enabled add-ons.

  Nested scheme for `addons`:
	- `allowed_upgrade_versions` - (String) The versions that the add-on is upgraded to.
	- `deprecated` - (String) Determines if the add-on version is deprecated.
	- `health_state` - (String) The health state of an add-on, such as critical or pending.
	- `health_status` - (String) The health status of an add-on, provides a description of the state in the form of error message.
	- `min_kube_version` - (String) The minimum Kubernetes version of the add-on.
	- `min_ocp_version` - (String) The minimum OpenShift version of the add-on.
	- `name` - (String) The add-on name such as `istio`.
	- `supported_kube_range` - (String) The supported Kubernetes version range of the add-on.
	- `target_version`-  (String) The add-on target version.
	- `version` - (String) The add-on version. Omit the version, if you need to use the default version.
	- `vlan_spanning_required`-  (String) The VLAN spanning required for multi-zone clusters.
- `id` - (String) The ID of an add-ons.
- `resource_group_id` - (String) The ID of the cluster resource group in which the `addons` is installed.