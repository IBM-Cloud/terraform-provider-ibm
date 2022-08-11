---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_alb_cert"
description: |-
  Get information about a IBM container ALB certificate.
---

# ibm_container_alb_cert
Retrieve information about all the Kubernetes cluster ALB certificate on IBM Cloud as a read-only data source.

## Example usage
The following example retrieves information of an ALB certificate.

```terraform
data "ibm_container_alb_cert" "cert" {
  secret_name = "test-sec"
  cluster_id  = "myCluster"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cluster_id` - (Required, String) The cluster ID.
- `namespace` - (Optional, string) The namespace in which the secret has to be **created.Default** `ibm-cert-store`
- `secret_name` - (Required, String) The name of the ALB certificate secret.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `cert_crn` - (String) The certificate CRN.
- `cloud_cert_instance_id` - (String) The Cloud certificate instance ID from which certificate is downloaded.
- `cluster_crn` - (String) The cluster CRN.
- `domain_name` - (String) The domain name of the certificate.
- `expires_on` - (String) The expiry date of the certificate.
- `id` - (String) The ALB cert ID. The ID is composed of `<cluster_name_id>/<secret_name>`.
- `issuer_name` - (String) The issuer name of the certificate.
- `persistence`  - (Bool) Persist the secret data in your cluster. If the secret is later deleted from the command line or OpenShift web console, the secret is automatically re-created in your cluster.
- `status` - (String) The Status of the secret.
