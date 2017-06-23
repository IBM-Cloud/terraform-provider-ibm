---
layout: "ibm"
page_title: "IBM: ibm_container_cluster_config"
sidebar_current: "docs-ibm-datasource-container-cluster-config"
description: |-
  Get the cluster configuration for Kubernetes on IBM Bluemix.
---

# ibm\_container_cluster_config


Download a configuration for Kubernetes clusters on IBM Bluemix.


## Example Usage

```hcl
data "ibm_container_cluster_config" "cluster_foo" {
  org_guid     = "test"
  space_guid   = "test_space"
  account_guid = "test_acc"
  name         = "FOO"
  config_dir   = "/home/foo_config"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name_id` - (Required) Name or ID of the cluster.
* `config_dir` - (Required) The directory where you want the cluster configuration to download.
* `admin` - (Optional) Set to `true` to download config for the admin. Default value: `false`.
* `download` - (Optional) Set to `false` to skip downloading the config for the admin. Default value: `true`. Since it is part of a data source, the config is downloaded for every `terraform` call by default. For a particular cluster name/ID, the config is guaranteed to be downloaded to the same path for a given `config_dir`.
* `org_guid` - (Required) The GUID for the Bluemix organization that the cluster is associated with. The value can be retrieved from the `ibm_org` data source, or by running the `bx iam orgs --guid` command in the [Bluemix CLI](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html#getting-started).
* `space_guid` - (Required) The GUID for the Bluemix space that the cluster is associated with. The value can be retrieved from the `ibm_space` data source, or by running the `bx iam space <space-name> --guid` command in the Bluemix CLI.
* `account_guid` - (Required) The GUID for the Bluemix account that the cluster is associated with. The value can be retrieved from the `ibm_account` data source, or by running the `bx iam accounts` command in the Bluemix CLI.


## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the Cluster config 
* `config_file_path` - The path to the cluster config file. Typically the Kubernetes YAML config file.
