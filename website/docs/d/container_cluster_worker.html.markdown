---
layout: "ibm"
page_title: "IBM: ibm_container_cluster_worker"
sidebar_current: "docs-ibm-datasource-container-cluster-worker"
description: |-
  Get information about a worker node that is attached to a Kubernetes cluster on IBM Cloud.
---

# ibm\_container_cluster_worker


Import details of a worker node of a Kubernetes cluster as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_container_cluster_worker" "cluster_foo" {
  worker_id    = "dev-mex10-pa70c4414695c041518603bfd0cd6e333a-w1"
  org_guid     = "test"
  space_guid   = "test_space"
  account_guid = "test_acc"
}
```

## Argument Reference

The following arguments are supported:

* `worker_id` - (Required, string) The ID of the worker node attached to the cluster.
* `org_guid` - (Optional, string) The GUID for the IBM Cloud organization that the cluster is associated with. You can retrieve the value from the `ibm_org` data source or by running the `bx iam orgs --guid` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `space_guid` - (Optional, string) The GUID for the IBM Cloud space that the cluster is associated with. You can retrieve the value from the `ibm_space` data source or by running the `bx iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Optional, string) The GUID for the IBM Cloud account that the cluster is associated with. You can retrieve the value from the `ibm_account` data source or by running the `bx iam accounts` command in the IBM Cloud CLI.
* `region` - (Optional, string) The region where the worker is provisioned. If the region is not specified it will be defaulted to provider region(BM_REGION/BLUEMIX_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.

## Attribute Reference

The following attributes are exported:

* `state` - The unique identifier of the cluster.
* `status` - The number of workers nodes attached to the cluster.
* `private_vlan` - The private VLAN of the worker node.
* `public_vlan` -  The public VLAN of the worker node.
* `private_ip` - The private IP of the worker node.
* `public_ip` -  The public IP of the worker node.
