---
layout: "ibm"
page_title: "IBM: container_alb_cert"
sidebar_current: "docs-ibm-resource-container-alb-cert"
description: |-
  Get information about a IBM container alb cert.
---

# ibm\_container_alb_cert

Import the details of a Kubernetes cluster ALB certificate on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

In the following example, you can retrive a alb cert :

```hcl
data "ibm_container_alb_cert" "cert" {
  secret_name = "test-sec"
  cluster_id  = "myCluster"
}

```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, string)  The cluster ID.
* `secret_name` - (Required, string) The name of the ALB certificate secret. 
* `namespace` - (Optional, string) The namespace in which the secret has to be created.Default: `ibm-cert-store`


## Attribute Reference

The following attributes are exported:

* `id` - The ALB cert ID. The id is composed of \<cluster_name_id\>/\<secret_name\>.<br/>
* `cert_crn` - The certificate CRN. 
* `domain_name` - The Domain name of the certificate.
* `expires_on` - The Expiry date of the certificate.
* `issuer_name` - The Issuer name of the certificate.
* `cluster_crn` - The cluster crn.
* `cloud_cert_instance_id` - Cloud Certificate instance ID from which certificate is downloaded.
* `persistence`  - (Optional, bool) Persist the secret data in your cluster. If the secret is later deleted from the CLI or OpenShift web console, the secret is automatically re-created in your cluster.
* `status` - The Status of the secret.
