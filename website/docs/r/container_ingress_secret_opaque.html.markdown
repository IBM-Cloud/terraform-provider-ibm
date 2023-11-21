---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_ingress_secret_opaque"
description: |-
  Registers an IBM Cloud Secrets Manager certificate secret with your cluster
---

# ibm_container_ingress_secret_opaque
Registers an IBM Cloud Secrets Manager secret type certificate with your IBM Cloud Kubernetes Service or Red Hat OpenShift on IBM Cloud cluster. For more information about how opaque secrets can be used see [about Secrets Manager secrets](https://cloud.ibm.com/docs/containers?topic=containers-secrets#non-tls)

## Example usage

```terraform
resource "ibm_container_ingress_secret_opaque" "secret" {
  cluster="exampleClusterName"
  secret_name="mySecretName"
  secret_namespace="mySecretNamespace"
  persistence = true
  fields {
    crn = "cert:region:crn"
  }
  fields {
    field_name = "myFieldName"
    crn = "cert:region:crn"
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster` - (Required, String) The cluster ID.
- `secret_name` - (Required, String) The name of the kubernetes secret.
- `secret_namespace` - (Required, String) The namespace of the kubernetes secret.
- `persistence`  - (Bool) Persist the secret data in your cluster. If the secret is later deleted from the command line or OpenShift web console, the secret is automatically re-created in your cluster.
- `update_secret` - (Optional, Integer) This argument is used to force update from upstream secrets manager instance that stores secret. Increment the value to force an update to your Ingress secret for changes made to the upstream secrets manager secret. 
- `fields` - (Required, List) List of fields of the opaque secret.
  
  Nested scheme for `fields`:
  - `crn` - (Required, String) Secrets manager secret crn
  - `field_name` - (String) Field name
  - `prefix` - (String) Prefix field name with Secrets Manager secret name

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `secret_type` - (String) The type of Kubernetes secret (Opaque).
- `status` - (String) The Status of the secret.
- `user_managed` - (Bool) Indicates whether the secret was created by a user.
- `persistence`  - (Bool) Persist the secret data in your cluster. If the secret is later deleted from the command line or OpenShift web console, the secret is automatically re-created in your cluster.
- `fields` - (List) List of fields of the opaque secret.
  
  Nested scheme for `fields`:
  - `crn` - (String) Secrets manager secret crn
  - `field_name` - (String) Requested field name
  - `name` - (String) computed field name
  - `expires_on` - (String) Expiration date of the secret
  - `type` - (String) Secrets manager secret type
  - `last_updated_timestamp` - (String) The most recent time the kubernetes secret was updated
