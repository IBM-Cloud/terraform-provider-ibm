---
subcategory: ""
layout: "ibm"
page_title: "Binding a Service to a Cluster"
description: |-
    Bind an IBM Cloud service to an IBM Cloud Kubernetes Service cluster.
---

# Binding a Service to a Cluster

Bind an IBM Cloud service to an IBM Cloud Kubernetes Service cluster. Service binding is a quick way to create service credentials for an IBM Cloud service by using its public service endpoint and storing these credentials in a Kubernetes secret in your cluster. The Kubernetes secret is automatically encrypted in etcd to protect your data.

To bind a service to your cluster, you need to:

1. Create a resource key for the service
2. Use the target cluster's config as the `kubernetes` provider's configuration
3. Create a kubernetes secret, using the resource key's credentials

## Example

In the following example, we bind the `ibm_resource_instance.kms` service to `ibm_container_cluster.cluster`.

```terraform
// create resource key
resource "ibm_resource_key" "kms_key" {
    name                 = "kms_key"
    resource_instance_id = ibm_resource_instance.kms.id
}

// get cluster config by cluster ID
data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = ibm_container_cluster.cluster.id
}

// use kubernetes provider configuration from cluster
provider "kubernetes" {
  host                   = data.ibm_container_cluster_config.cluster_config.host
  token                  = data.ibm_container_cluster_config.cluster_config.token
  cluster_ca_certificate = data.ibm_container_cluster_config.cluster_config.ca_certificate
}

// create kubernetes secret from resource key's credentials
resource "kubernetes_secret_v1" "kms_secret" {
  metadata {
    name      = "kms-secret"
    namespace = "default"
  }

  data = ibm_resource_key.kms_key.credentials
}
```
