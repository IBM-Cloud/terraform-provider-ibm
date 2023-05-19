---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_ingress_secret_opaque"
description: |-
  Get information about a managed opaque kubernetes secret backed by Secrets Manager secrets
---

# ibm_container_ingress_tls_secret
Get details about a managed opaque secret that is stored as a Kubernetes secret.

## Example usage
The following example retrieves information about the managed opaque secret that is named `mysecret` in the namespace `mynamespace` of a cluster that is named `mycluster`. 

```terraform
data ibm_container_ingress_secret_tls secret {
    cluster ="mycluster"
    secret_name = "mysecret"
    secret_namespace = "mynamespace"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cluster_id` - (Required, String) The cluster ID.
- `secret_name` - (Required, String) The name of the kubernetes secret.
- `secret_namespace` - (Required, string) The namespace of the kubernetes secret.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `secret_type` - (String) The type of Kubernetes secret (opaque).
- `persistence`  - (Bool) Persist the secret data in your cluster. If the secret is later deleted from the command line or OpenShift web console, the secret is automatically re-created in your cluster.
- `status` - (String) The status of the secret.
- `user_managed` - (Bool) Indicates whether the secret was created by a user.
- `fields` - (String) The fields of the opaque secret.
  - `crn` - (String) Secrets manager secret crn
  - `name` - (String) Field name
  - `expires_on` - (String) Expiration date of the secret
  - `type` - (String) Secrets manager secret type
  - `last_updated_timestamp` - (String) The most recent time the kubernetes secret was updated
