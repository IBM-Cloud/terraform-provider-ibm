---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_addons"
description: |-
  Manages IBM container addons.
---

# ibm_container_addons
Enable, update or disable a single add-on or a set of add-ons. For more information, see [Cluster addons](https://cloud.ibm.com/docs/containers?topic=containers-api-at-iam#ks-cluster).

## Example usage

In the following example, you can configure a add-ons:

```terraform
resource "ibm_container_addons" "addons" {
  cluster = ibm_container_vpc_cluster.cluster.name
  addons {
    name    = "istio"
    version = "1.7"
  }
  addons {
    name    = "kube-terminal"
    version = "1.0.0"
  }
  addons {
    name    = "alb-oauth-proxy"
    version = "1.0.0"
  }
  addons {
    name    = "debug-tool"
    version = "2.0.0"
  }
  addons {
    name    = "knative"
    version = "0.17.0"
  }
  addons {
    name    = "static-route"
    version = "1.0.0"
  }
  addons {
    name    = "vpc-block-csi-driver"
    version = "2.0.3"
  }
  addons {
    name    = "cluster-autoscaler"
    version = "1.0.1"
  }
}

```

## Timeouts

The `ibm_container_addons` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create** The enablement of the add-ons is considered `failed` if no response is received for 20 minutes.
- **Update** The enablement of the add-ons is considered `failed` if no response is received for 20 minutes.


## Argument reference
Review the argument references that you can specify for your resource. 

- `addons` - (Required, Set) Set of add-ons that needs to be enabled.

  Nested scheme for `addons`:
	- `name` - (Optional, String) The add-on name such as `istio`. Supported add-ons are `kube-terminal`, `alb-oauth-proxy`, `debug-tool`, `istio`, `knative`, `static-route`,`vpc-block-csi-driver`.
      * [Kubernetes Cluster](https://cloud.ibm.com/docs/containers?topic=containers-managed-addons#adding-managed-add-ons)
      * [Openshift Cluster](https://cloud.ibm.com/docs/openshift?topic=openshift-managed-addons#adding-managed-add-ons)
      * [Satellite Cluster]( https://cloud.ibm.com/docs/openshift?topic=openshift-managed-addons#addons-satellite)
  - `version`- (Optional, String) The add-on version. Omit the version that you want to use as the default version.This is required when you want to update the add-on to specified version.
- `cluster` - (Required, String) The name or ID of the cluster.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group. You can retrieve the value from data source ibm_resource_group. If not provided defaults to default resource group.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `addons` - (String) The details of an enabled add-ons.

  Nested scheme for `addons`:
	- `allowed_upgrade_versions` - (String) The versions that the add-on is upgraded to.
	- `deprecated` - (String) Determines if the add-on version is deprecated.
	- `health_state` - (String) The health state of an add-on, such as critical or pending.
	- `health_status` - (String) The health status of an add-on, provides a description of the state in the form of error message.
	- `min_kube_version` - (String) The minimum Kubernetes version of the add-on.
	- `min_ocp_version` - (String) The minimum OpenShift version of the add-on.
	- `supported_kube_range` - (String) The supported Kubernetes version range of the add-on.
  - `target_version`-  (String) The add-on target version.
  - `vlan_spanning_required`-  (String) The VLAN spanning required for multi-zone clusters.
- `id` - (String) The ID of the add-ons.
