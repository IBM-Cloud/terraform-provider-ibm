---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_ingress_secret_tls"
description: |-
  Registers an IBM Cloud Secrets Manager certificate secret with your cluster
---

# ibm_container_ingress_instance
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
 `fields` - (Required, string) The fields of the opaque secret.
  - `crn` - (String) Secrets manager secret crn
  - `name` - (String) Field name
  - `prefix` - (String) Prefix field name with Secrets Manager secret name

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `secret_type` - (String) The type of Kubernetes secret (TLS).
- `status` - (String) The Status of the secret.
- `user_managed` - (Bool) Indicates whether the secret was created by a user.
- `persistence`  - (Bool) Persist the secret data in your cluster. If the secret is later deleted from the command line or OpenShift web console, the secret is automatically re-created in your cluster.
 `fields` - (Required, string) The fields of the opaque secret.
  - `crn` - (String) Secrets manager secret crn
  - `name` - (String) Field name
  - `expires_on` - (String) Expiration date of the secret
  - `type` - (String) Secrets manager secret type
  - `last_updated_timestamp` - (String) The most recent time the kubernetes secret was updated
