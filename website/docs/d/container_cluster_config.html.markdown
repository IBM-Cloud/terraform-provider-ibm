---
layout: "ibm"
page_title: "IBM: ibm_container_cluster_config"
sidebar_current: "docs-ibm-datasource-container-cluster-config"
description: |-
  Get the cluster configuration for Kubernetes on IBM Cloud.
---

# ibm\_container_cluster_config


Download a configuration for Kubernetes clusters on IBM Cloud. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_container_cluster_config" "cluster_foo" {
  org_guid        = "test"
  space_guid      = "test_space"
  account_guid    = "test_acc"
  cluster_name_id = "FOO"
  config_dir      = "/home/foo_config"
  region          = "eu-de"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name_id` - (Required, string) The name or ID of the cluster.
* `config_dir` - (Required, string) The directory where you want the cluster configuration to download.
* `admin` - (Optional, boolean) Set the value to `true` to download the configuration for the administrator. The default value is `false`.
* `download` - (Optional, boolean) Set the value to `false` to skip downloading the configuration for the administrator. The default value is `true`. Because it is part of a data source, by default the configuration is downloaded for every Terraform call. For a particular cluster name or ID, the configuration is guaranteed to be downloaded to the same path for a given `config_dir`.
* `org_guid` - (Optional, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from the `ibm_org` data source or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `space_guid` - (Optional, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Optional, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from the `ibm_account` data source or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
* `region` - (Optional, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(BM_REGION/BLUEMIX_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `network` - (Optional, boolean) Set the value to `true` to download the configuration for the Calico network config with the Admin config. The default value is `false`.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster configuration.
* `config_file_path` - The path to the cluster configuration file. This is typically the Kubernetes YAML configuration file.
* `calico_config_file_path` - The path to the cluster calico configuration file.