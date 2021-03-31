---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_ob_logging"
description: |-
  Create a logging configuration for your cluster to automatically collect pod logs and send them to IBM Log Analysis.
---

# ibm\_ob_logging


Create, update, or delete a logging instance. This resource creates a logging configuration for your cluster to automatically collect pod logs and send them to IBM Log Analysis.

## Example Usage

In the following example, we can create a logging configuration:

```hcl

data "ibm_resource_group" "testacc_ds_resource_group" {
  name = "Default"
}

resource "ibm_container_cluster" "testacc_cluster" {
  name              = "TestCluster"
  datacenter        = "dal10"
  resource_group_id = data.ibm_resource_group.testacc_ds_resource_group.id
  default_pool_size = 1
  wait_till         = "MasterNodeReady"
  hardware          = "shared"
  machine_type      = "%s"
  timeouts {
    create = "720m"
    update = "720m"
  }
}

resource "ibm_resource_instance" "instance" {
  name     = "TestLogging"
  service  = "logdna"
  plan     = "7-day"
  location = "us-south"
}

resource "ibm_resource_key" "resourceKey" {
  name                 = "TestKey"
  resource_instance_id = ibm_resource_instance.instance.id
  role                 = "Manager"
}

resource "ibm_ob_logging" "test2" {
  depends_on  = [ibm_resource_key.resourceKey]
  cluster     = ibm_container_cluster.testacc_cluster.id
  instance_id = ibm_resource_instance.instance.guid
}

```

## Timeouts

ibm_ob_logging provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating Instance.
* `update` - (Default 10 minutes) Used for updating Instance.
* `delete` - (Default 10 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `cluster` - (Required, string) The name or id of the cluster. 
* `instance_id` - (Required, string) The guid of the montoing instance.
* `logdna_ingestion_key` - (Optional, string) The LogDNA ingestion key that you want to use for your configuration
* `private_endpoint` - (Optional, string) Add this option to connect to your logging service instance through the private service endpoint.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the logging instance attach to cluster. The id is composed of \<cluster_name_id\>/\< logging_instance_id\>
* `instance_name` - Name of the logging instance
* `agent_key` - LogDNA agent key
* `agent_namespace` - LogDNA agent namespace
* `daemonset_name` - Name of the deamonset
* `namespace` - Name of namespace
* `crn` - CRN of the LogDNA instance attach

## Import

ibm_ob_logging can be imported using cluster_name_id, logging_instance_id, eg:

```
$ terraform import ibm_ob_logging.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
