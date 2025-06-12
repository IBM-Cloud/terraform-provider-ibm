---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_ingress_secret_tls"
description: |-
  Registers an IBM Cloud Secrets Manager certificate secret with your cluster
---

# ibm_container_ingress_secret_tls
Registers an IBM Cloud Secrets Manager secret type certificate with your IBM Cloud Kubernetes Service or Red Hat OpenShift on IBM Cloud cluster. For more information about how TLS secrets can be used see [about Secrets Manager secrets](https://cloud.ibm.com/docs/containers?topic=containers-secrets#tls)

## Example usage

```terraform
resource "ibm_container_ingress_secret_tls" "secret" {
  cluster="exampleClusterName"
  secret_name="mySecretName"
  secret_namespace="mySecretNamespace"
  cert_crn="cert:region:crn"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster` - (Required, String) The cluster ID.
- `secret_name` - (Required, String) The name of the kubernetes secret.
- `secret_namespace` - (Required, string) The namespace of the kubernetes secret.
- `cert_crn` - (Required, string) The Secrets Manager crn for a secret of type certificate.
- `update_secret` - (Optional, Integer) This argument is used to force update from upstream secrets manager instance that stores secret. Increment the value to force an update to your Ingress secret for changes made to the upstream secrets manager secret. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `persistence`  - (Bool) Persist the secret data in your cluster. If the secret is later deleted from the command line or OpenShift web console, the secret is automatically re-created in your cluster.
- `domain_name` - (String) Domain name.
- `expires_on` - (String) Certificate expires on date.
- `status` - (String) The Status of the secret.
- `user_managed` - (Bool) Indicates whether the secret was created by a user.
- `secret_type` - (String) The type of Kubernetes secret (TLS).
- `type` - (String) Type of Secret Manager secret.
- `last_updated_timestamp` - (String) Timestamp secret was last updated in cluster.
