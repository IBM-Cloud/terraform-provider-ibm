---
layout: "ibm"
page_title: "IBM: container_addons"
sidebar_current: "docs-ibm-resource-container-addons"
description: |-
  Manages IBM container addons.
---

# ibm\_container_addons

Enable, update or Disable a single AddOn or a Set of AddOns. 

## Example Usage

In the following example, you can configure a Addons:

```hcl
resource "ibm_container_addons" "addons" {
  cluster = ibm_container_vpc_cluster.cluster.name
  addons {
    name    = "istio"
    version = "1.6"
  }
}

```

## Timeouts

ibm_container_addons provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 20 minutes) Used for creating Instance.
* `update` - (Default 20 minutes) Used for updating Instance.

## Argument Reference

The following arguments are supported:

* `cluster` - (Required, string) Cluster Name or ID
* `resource_group_id` - (Optional, Forces new resource, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `addons` - (Required, set) Set of AddOns that has to be enabled
    * `name` - (Optional, string) The addon name such as 'istio'.
    * `version` - (Optional,string) The addon version, omit the version if you wish to use the default version. It is required when one wants to update the addon to specified version.

## Attribute Reference

The following attributes are exported:

* `id` - The AddOns ID.
* `addons` - Details of Enabled AddOns
    * `allowed_upgrade_versions` - The versions that the addon can be upgraded to 
    * `deprecated` - Determines if this addon version is deprecated
    * `health_state` - The health state for this addon, a short indication (e.g. critical, pending)
    * `health_status` - The health status for this addon, provides a description of the state (e.g. error message)
    * `min_kube_version` - The minimum kubernetes version for this addon.
    * `min_ocp_version` - The minimum OpenShift version for this addon.
    * `supported_kube_range` - The supported kubernetes version range for this addon.
    * `target_version` - The addon target version.
    * `vlan_spanning_required` - VLAN spanning required for multi-zone clusters.