---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_ingress_secret_tls"
description: |-
  Get information about a managed TLS kubernetes secret backed by Secrets Manager secrets
---

# ibm_container_ingress_secret_tls
Get details about a managed TLS certificate that is stored as a Kubernetes TLS secret.

## Example usage
The following example retrieves information about the registered Secrets Manager TLS secret that is named `mysecret` in the namespace `mynamespace` of a cluster that is named `mycluster`. 

```terraform
data ibm_container_ingress_secret_tls secret {
    cluster ="mycluster"
    secret_name = "mysecret"
    secret_namespace = "mynamespace"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cluster` - (Required, String) The cluster ID.
- `secret_name` - (Required, String) The name of the kubernetes secret.
- `secret_namespace` - (Required, string) The namespace of the kubernetes secret.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `cert_crn` - (String) The backing IBM Cloud Secrets Manager Secret CRN.
- `type` - (String) The type of Kubernetes secret (TLS).
- `status` - (String) The Status of the secret.
- `user_managed` - (Bool) Indicates whether the secret was created by a user.
- `persistence`  - (Bool) Persist the secret data in your cluster. If the secret is later deleted from the command line or OpenShift web console, the secret is automatically re-created in your cluster.
- `domain_name` - (String) Domain name.
- `expires_on` - (String) Certificate expires on date.
- `last_updated_timestamp` - (String) Timestamp secret was last updated in cluster.