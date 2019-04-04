---
layout: "ibm"
page_title: "IBM: container_cluster_feature"
sidebar_current: "docs-ibm-resource-container-cluster-feature"
description: |-
  Manages IBM container cluster feature.
---

# ibm\_container_cluster_feature

Enables or disables a container cluster feature. 

## Example Usage

In the following example, you can enable a private service endpoint:

```hcl
resource ibm_container_cluster_feature feature {
  cluster                 = "test1"
  private_service_endpoint = "true"
}

```

## Timeouts

ibm_container_cluster_feature provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 90 minutes) Used for creating Instance.
* `update` - (Default 90 minutes) Used for updating Instance.

## Argument Reference

The following arguments are supported:

* `cluster` - (Required, string) The cluster name or id.
* `public_service_endpoint` - (Optional, bool)  Enable or disable the public service endpoint.
* `private_service_endpoint` - (Optional, bool) Enable the private service endpoint to make the master privately accessible. Once enabled this feature cannot be disabled later.
  **NOTE**: As a prerequisite for using Service Endpoints, Account must be enabled for Virtual Routing and Forwarding (VRF). Learn more about VRF on IBM Cloud [here](https://console.bluemix.net/docs/infrastructure/direct-link/vrf-on-ibm-cloud.html#overview-of-virtual-routing-and-forwarding-vrf-on-ibm-cloud). Account must be enabled for connectivity to Service Endpoints.
* `refresh_api_servers` - (Optional, bool) To apply these changes, refresh the cluster's API server. Default value is true.
* `reload_workers` - (Optional, bool) To apply these changes, reload workers. Default value is true.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `region` - (Optional, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(BM_REGION/BLUEMIX_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.

## Attribute Reference

The following attributes are exported:

* `id` - The cluster feature ID.
* `public_service_endpoint_url` - Public service endpoint url.
* `private_service_endpoint_url` - Private service endpoint url.