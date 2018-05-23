---
layout: "ibm"
page_title: "IBM: ibm_container_cluster_versions"
sidebar_current: "docs-ibm-datasource-container-cluster-versions"
description: |-
  List supported kubernetes versions on IBM Cloud.
---

# ibm\_container_cluster_versions

Get the list of supported kubernetes versions on IBM Cloud. Please refer to https://console.bluemix.net/docs/containers/cs_versions.html#cs_versions for detail instructions.

## Example Usage

```hcl
data "ibm_container_cluster_versions" "cluster_versions" {
  org_guid        = "test"
  space_guid      = "test_space"
  account_guid    = "test_acc"
}
```

## Argument Reference

The following arguments are supported:

* `org_guid` - (Optional, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from the `ibm_org` data source or by running the `bx iam orgs --guid` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `space_guid` - (Optional, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from the `ibm_space` data source or by running the `bx iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Required, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from the `ibm_account` data source or by running the `bx iam accounts` command in the IBM Cloud CLI.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster versions.
* `valid_kube_versions` - The supported kubernetes versions on IBM Cloud.