---
layout: "ibm"
page_title: "IBM: ibm_container_cluster"
sidebar_current: "docs-ibm-datasource-container-cluster"
description: |-
  Get information about a Kubernetes cluster on IBM Bluemix.
---

# ibm\_container_cluster


Import the details for a Kubernetes cluster on IBM Bluemix as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax. 


## Example Usage

```hcl
data "ibm_container_cluster" "cluster_foo" {
  cluster_name_id = "FOO"
  org_guid        = "test"
  space_guid      = "test_space"
  account_guid    = "test_acc"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name_id` - (Required) Name or ID of the cluster.
* `org_guid` - (Required) The GUID for the Bluemix organization that the cluster is associated with. The value can be retrieved from the `ibm_org` data source, or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `space_guid` - (Required) The GUID for the Bluemix space that the cluster is associated with. The value can be retrieved from the `ibm_space` data source, or by running the `bx iam space <space-name> --guid` command in the Bluemix CLI.
* `account_guid` - (Required) The GUID for the Bluemix account that the cluster is associated with. The value can be retrieved from the `ibm_account` data source, or by running the `bx iam accounts` command in the Bluemix CLI.


## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster.
* `worker_count` - Number of workers attached to the cluster.
* `workers` - IDs of the worker attached to the cluster.
* `bounded_services` - Services that are bounded to the cluster.
