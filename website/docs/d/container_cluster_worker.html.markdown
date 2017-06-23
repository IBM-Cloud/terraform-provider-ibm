---
layout: "ibm"
page_title: "IBM: ibm_container_cluster_worker"
sidebar_current: "docs-ibm-datasource-container-cluster-worker"
description: |-
  Get information about a worker node that is attached to a Kubernetes cluster on IBM Bluemix.
---

# ibm\_container_cluster_worker


Import details about the worker node of a Kubernetes cluster as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax. 


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

* `worker_id` - (Required) ID of the worker node attached to the cluster.
* `org_guid` - (Required) The GUID for the Bluemix organization that the cluster is associated with. The value can be retrieved from the `ibm_org` data source, or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `space_guid` - (Required) The GUID for the Bluemix space that the cluster is associated with. The value can be retrieved from the `ibm_space` data source, or by running the `bx iam space <space-name> --guid` command in the Bluemix CLI.
* `account_guid` - (Required) The GUID for the Bluemix account that the cluster is associated with. The value can be retrieved from the `ibm_account` data source, or by running the `bx iam accounts` command in the Bluemix CLI.


## Attributes Reference

The following attributes are exported:

* `state` - The unique identifier of the cluster.
* `status` - Number of workers nodes attached to the cluster.
* `private_vlan` - The private VLAN of the worker node.
* `public_vlan` -  The public VLAN of the worker node.
* `private_ip` - The private IP of the worker node.
* `public_ip` -  The public IP of the worker node.
