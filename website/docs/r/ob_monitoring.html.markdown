---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_ob_monitoring"
description: |-
  Create a monitoring configuration for your cluster to automatically collect pod logs and send them to IBM Log Analysis.
---

# ibm\_ob_monitoring


Create, update, or delete a monitoring instance. This resource creates a monitoring configuration for your IBM Cloud Kubernetes Service cluster to automatically collect cluster and pod metrics, and send these metrics to your IBM Cloud Monitoring with Sysdig service instance.

## Example Usage

In the following example, you can create a monitoring configuration:

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
  name     = "TestMonitoring"
  service  = "sysdig-monitor"
  plan     = "graduated-tier"
  location = "us-south"
}

resource "ibm_resource_key" "resourceKey" {
  name                 = "TestKey"
  resource_instance_id = ibm_resource_instance.instance.id
  role                 = "Manager"
}

resource "ibm_ob_monitoring" "test2" {
  depends_on  = [ibm_resource_key.resourceKey]
  cluster     = ibm_container_cluster.testacc_cluster.id
  instance_id = ibm_resource_instance.instance.guid
}

```

## Timeouts

ibm_ob_monitoring provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating Instance.
* `update` - (Default 10 minutes) Used for updating Instance.
* `delete` - (Default 10 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `cluster` - (Required, string) The name or id of the cluster. 
* `instance_id` - (Required, string) The guid of the montoing instance.
* `sysdig_access_key` - (Optional, string) The sysdig monitoring ingestion key that you want to use for your configuration
* `private_endpoint` - (Optional, string) Add this option to connect to your Sysdig service instance through the private service endpoint.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the sysdig monitoring instance attach to cluster. The id is composed of \<cluster_name_id\>/\< monitoring_instance_id\>
* `instance_name` - Name of the monitoring instance
* `agent_key` - Sysydig agent key type.
* `agent_namespace` - Sysdig agent namespace type
* `daemonset_name` - Name of the deamonset
* `discovered_agent` - 
* `namespace` - Name of namespace
* `crn` - CRN of the instance attach

## Import

ibm_ob_monitoring can be imported using cluster_name_id, monitoring_instance_id, eg

```
$ terraform import ibm_ob_monitoring.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
