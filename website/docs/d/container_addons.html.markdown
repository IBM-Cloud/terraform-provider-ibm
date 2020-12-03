---
layout: "ibm"
page_title: "IBM: container_addons"
sidebar_current: "docs-ibm-data-source-container-addons"
description: |-
  Reads all enabled IBM  container addons.
---

# ibm\_container_addons

Reads all the addOns that are enabled on a cluster

## Example Usage

In the following example, you can get details of Addons:

```hcl
data "ibm_container_addons" "addons" {
  cluster= ibm_container_addons.addons.cluster
}

```

## Argument Reference

The following arguments are supported:

* `cluster` - (Required, string) Cluster Name or ID

## Attribute Reference

The following attributes are exported:

* `id` - The AddOns ID.
* `resource_group_id` - The ID of the cluster resource group in which the addons have to be installed.
* `addons` - Details of Enabled AddOns
  * `name` - The addon name such as 'istio'.
  * `version` - The addon version, omit the version if you wish to use the default version.
  * `allowed_upgrade_versions` - The versions that the addon can be upgraded to 
  * `deprecated` - Determines if this addon version is deprecated
  * `health_state` - The health state for this addon, a short indication (e.g. critical, pending)
  * `health_status` - The health status for this addon, provides a description of the state (e.g. error message)
  * `min_kube_version` - The minimum kubernetes version for this addon.
  * `min_ocp_version` - The minimum OpenShift version for this addon.
  * `supported_kube_range` - The supported kubernetes version range for this addon.
  * `target_version` - The addon target version.
  * `vlan_spanning_required` - VLAN spanning required for multi-zone clusters.