---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM : ibm_container_nlb_dns"
description: |-
  Manages IBM Container Nlb
---

# ibm_container_nlb_dns

Provides a resource for container_nlb_dns. This allows to add an NLB IP's to an existing host name that you created with 'ibmcloud ks nlb-dns create'.

## Example usage

```terraform
data "ibm_container_nlb_dns" "dns" {
  cluster = var.cluster
}

resource "ibm_container_nlb_dns" "container_nlb_dns" {
  cluster 	= var.cluster
  nlb_host	= data.ibm_container_nlb_dns.dns.nlb_config.0.nlb_sub_domain
  nlb_ips 	= var.cluster_ips
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `cluster` - (Required, Forces new resource, String) The name or ID of the cluster. To list the clusters that you have access to, use the `GET /v1/clusters` API or run `ibmcloud ks cluster ls`.
* `nlb_host` - (Required, Forces new resource, String) Host Name of load Balancer.
* `nlb_ips` - (Required, Set)  NLB IPs.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the container_nlb_dns.
* `nlb_dns_type` - Type of DNS.
* `nlb_monitor_state` -  Nlb monitor state.
* `nlb_ssl_secret_name` - Name of SSL Secret.
* `nlb_ssl_secret_status` - Status of SSL Secret.
* `nlb_type` - Nlb Type.
* `secret_namespace` - Namespace of Secret.
* `resource_group_id` - The ID of the resource group that the cluster is in. To check the resource group ID of the cluster, use the GET /v1/clusters/idOrName API. To list available resource group IDs, run ibmcloud resource groups.

## Import

The import functionality is not supported for this resource.
