---
layout: "ibm"
page_title: "IBM: ibm_container_cluster"
sidebar_current: "docs-ibm-datasource-container-cluster"
description: |-
  Get information about a Kubernetes cluster on IBM Cloud.
---

# ibm\_container_cluster


Import the details of a Kubernetes cluster on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


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

* `cluster_name_id` - (Required, string) The name or ID of the cluster.
* `org_guid` - (Optional, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from the `ibm_org` data source or by running the `bx iam orgs --guid` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `space_guid` - (Optional, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from the `ibm_space` data source or by running the `bx iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Required, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from the `ibm_account` data source or by running the `bx iam accounts` command in the IBM Cloud CLI.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster.
* `worker_count` - The number of workers that are attached to the cluster.
* `workers` - The IDs of the workers that are attached to the cluster.
* `bounded_services` - The services that are bounded to the cluster.
* `vlans` - The VLAN'S that are attached to the cluster. Nested `vlans` blocks have the following structure:
	* `id` - The VLAN id.
	* `subnets` - The list of subnets attached to VLAN belonging to the cluster. Nested `subnets` blocks have the following structure:
		* `id` - The Subnet Id.
		* `cidr` - The cidr range.
		* `ips` - The list of ip's in the subnet.
		* `is_byoip` - `true` if user provides a ip range else `false`.
		* `is_public` - `true` if VLAN is public else `false`.
